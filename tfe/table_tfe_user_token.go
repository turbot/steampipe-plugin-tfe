package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
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
	}
	return nil, nil
}
