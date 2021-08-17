package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfeSshKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_ssh_key",
		Description: "SSH keys in the organization.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("organization_name"),
			Hydrate:    listSshKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSshKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the ssh key."},
			// Others columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the ssh key."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("organization_name"), Description: "Name of the organization containing the ssh key."},
		},
	}
}

func listSshKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_ssh_key.listSshKey", "connection_error", err)
		return nil, err
	}
	result, err := conn.SSHKeys.List(ctx, d.KeyColumnQuals["organization_name"].GetStringValue(), tfe.SSHKeyListOptions{})
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("tfe_ssh_key.listSshKey", "query_error", err)
		return nil, err
	}
	for _, i := range result.Items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getSshKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_ssh_key.getSshKey", "connection_error", err)
		return nil, err
	}
	result, err := conn.SSHKeys.Read(ctx, d.KeyColumnQuals["id"].GetStringValue())
	if err != nil {
		plugin.Logger(ctx).Error("tfe_ssh_key.getSshKey", "query_error", err)
		return nil, err
	}
	return result, nil
}
