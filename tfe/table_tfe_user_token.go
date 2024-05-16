package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTfeUserToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_user_token",
		Description: "Tokens for the user.",
		List: &plugin.ListConfig{
			ParentHydrate: listCurrentUser,
			Hydrate:       listUserToken,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the token."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the token was created."},
			{Name: "last_used_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the token was last used."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the token."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the organization member."},
		},
	}
}

func listUserToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_user_token.listUserToken", "connection_error", err)
		return nil, err
	}
	user := h.Item.(*tfe.User)
	result, err := conn.UserTokens.List(ctx, user.ID)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_user_token.listUserToken", "query_error", err)
		return nil, err
	}
	for _, i := range result.Items {
		d.StreamListItem(ctx, i)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
