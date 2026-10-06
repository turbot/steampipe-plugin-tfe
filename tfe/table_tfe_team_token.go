package tfe

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTfeTeamToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_team_token",
		Description: "Team tokens in the organization.",
		List: &plugin.ListConfig{
			Hydrate: listTeamToken,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the token."},
			{Name: "team_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.ID"), Description: "ID of the team."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the token was created."},
			{Name: "last_used_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastUsedAt").NullIfZero(), Description: "Time when the token was last used, or null if it has never been used."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the token."},
			{Name: "expired_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ExpiredAt").NullIfZero(), Description: "Time when the token expires, or null if it does not expire or the server does not support token expiration."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the team."},
		},
	}
}

func listTeamToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_team_token.listTeamToken", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)
	options := tfe.TeamTokenListOptions{
		ListOptions: tfe.ListOptions{PageSize: 100},
	}
	if limit := d.QueryContext.Limit; limit != nil && *limit < 100 {
		options.PageSize = int(*limit)
	}
	for {
		result, err := conn.TeamTokens.List(ctx, organizationName, &options)
		if err != nil {
			if options.PageNumber == 0 && isNotFoundError(err) {
				plugin.Logger(ctx).Debug("tfe_team_token.listTeamToken", "legacy_fallback", "Team token listing is unavailable; reading tokens by team.")
				return listTeams(ctx, d, h, func(team *tfe.Team) error {
					token, err := getLegacyTeamToken(ctx, d, team)
					if err != nil {
						return err
					}
					if token != nil {
						d.StreamListItem(ctx, token)
					}
					return nil
				})
			}
			plugin.Logger(ctx).Error("tfe_team_token.listTeamToken", "query_error", err)
			return nil, err
		}
		for _, token := range result.Items {
			if token.Team == nil || token.Team.ID == "" {
				err := fmt.Errorf("team token %s has no team relationship", token.ID)
				plugin.Logger(ctx).Error("tfe_team_token.listTeamToken", "query_error", err)
				return nil, err
			}
			d.StreamListItem(ctx, token)
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if result.Pagination.CurrentPage >= result.Pagination.TotalPages {
			return nil, nil
		}
		options.PageNumber = result.Pagination.NextPage
	}
}

func getLegacyTeamToken(ctx context.Context, d *plugin.QueryData, team *tfe.Team) (*tfe.TeamToken, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_team_token.getLegacyTeamToken", "connection_error", err)
		return nil, err
	}
	result, err := conn.TeamTokens.Read(ctx, team.ID)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("tfe_team_token.getLegacyTeamToken", "query_error", err)
		return nil, err
	}
	result.Team = team
	return result, nil
}
