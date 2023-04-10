package tfe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTfeCurrentUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_current_user",
		Description: "User making the request.",
		List: &plugin.ListConfig{
			Hydrate: listCurrentUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the user."},
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "URL of the user avatar."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "User email address."},
			{Name: "is_service_account", Type: proto.ColumnType_BOOL, Description: "True if the user is a service account."},
			{Name: "two_factor", Type: proto.ColumnType_JSON, Description: "Details of two factor authentication for the user."},
			{Name: "unconfirmed_email", Type: proto.ColumnType_STRING, Description: "Unconfirmed email address for the user."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username of the user."},
			{Name: "v2_only", Type: proto.ColumnType_BOOL, Description: "If true, the user can only use v2."},
		},
	}
}

func listCurrentUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_current_user.listCurrentUser", "connection_error", err)
		return nil, err
	}
	result, err := conn.Users.ReadCurrent(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_current_user.listCurrentUser", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
