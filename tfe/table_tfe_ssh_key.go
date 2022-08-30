package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableTfeSshKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_ssh_key",
		Description: "SSH keys in the organization.",
		List: &plugin.ListConfig{
			Hydrate: listSshKey,
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
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the ssh key."},
		},
	}
}

func listSshKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_ssh_key.listSshKey", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)
	limit := d.QueryContext.Limit
	options := tfe.SSHKeyListOptions{
		ListOptions: tfe.ListOptions{
			// https://www.terraform.io/docs/cloud/api/index.html#pagination
			PageSize: 100,
		},
	}
	if limit != nil {
		if *limit < int64(100) {
			options.PageSize = int(*limit)
		}
	}

	pagesLeft := true
	for pagesLeft {
		result, err := conn.SSHKeys.List(ctx, organizationName, options)
		if err != nil {
			if isNotFoundError(err) {
				return nil, nil
			}
			plugin.Logger(ctx).Error("tfe_ssh_key.listSshKey", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		// Pagination
		if result.Pagination.CurrentPage < result.Pagination.TotalPages {
			options.PageNumber = result.Pagination.NextPage
		} else {
			pagesLeft = false
		}
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
