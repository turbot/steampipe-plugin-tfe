package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfeTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_team",
		Description: "Teams in the organization.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("organization_name"),
			Hydrate:    listTeam,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTeam,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the team."},
			{Name: "user_count", Type: proto.ColumnType_INT, Transform: transform.FromField("UserCount"), Description: "Number of users in the team."},
			// Others columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the team."},
			{Name: "organization_access", Type: proto.ColumnType_JSON, Description: "Organization access granted to the team."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("organization_name"), Description: "Name of the organization containing the team."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Description: "Permissions granted to the team."},
			{Name: "visibility", Type: proto.ColumnType_STRING, Description: "The team's visibility: secret, organization."},
		},
	}
}

func listTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_team.listTeam", "connection_error", err)
		return nil, err
	}
	result, err := conn.Teams.List(ctx, d.KeyColumnQuals["organization_name"].GetStringValue(), tfe.TeamListOptions{})
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("tfe_team.listTeam", "query_error", err)
		return nil, err
	}
	for _, i := range result.Items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.getTeam", "connection_error", err)
		return nil, err
	}
	result, err := conn.Teams.Read(ctx, d.KeyColumnQuals["id"].GetStringValue())
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.getTeam", "query_error", err)
		return nil, err
	}
	return result, nil
}
