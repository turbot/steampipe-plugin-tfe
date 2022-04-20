package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableTfeTeamMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_team_member",
		Description: "Team members in the organization.",
		List: &plugin.ListConfig{
			ParentHydrate: listTeam,
			Hydrate:       listTeamMember,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "team_id", Type: proto.ColumnType_STRING, Hydrate: getTeamIDQual, Transform: transform.FromField("ID"), Description: "ID of the team."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username of the team member."},
			// Other columns
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "URL of the user avatar."},
			{Name: "is_service_account", Type: proto.ColumnType_BOOL, Description: "True if the user is a service account."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "ID of the organization containing the team."},
			{Name: "two_factor", Type: proto.ColumnType_JSON, Description: "Details of two factor authentication for the user."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID"), Description: "ID of the user."},
			{Name: "v2_only", Type: proto.ColumnType_BOOL, Description: "If true, the user can only use v2."},
		},
	}
}

func getTeamIDQual(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	item := h.ParentItem.(*tfe.Team)
	return item, nil
}

func listTeamMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_team.listTeam", "connection_error", err)
		return nil, err
	}
	team := h.Item.(*tfe.Team)
	items, err := conn.TeamMembers.ListUsers(ctx, team.ID)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_team.listTeam", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
