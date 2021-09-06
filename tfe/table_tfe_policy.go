package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfePolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_policy",
		Description: "Policies in the organization.",
		List: &plugin.ListConfig{
			Hydrate: listPolicy,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPolicy,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the policy."},
			// Others columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the policy."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the policy."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the policy's purpose. This field supports Markdown and will be rendered in the Terraform Cloud UI."},
			{Name: "policy_set_count", Type: proto.ColumnType_INT, Description: "The number of policy sets in the policy"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The update timestamp of the policy."},
			{Name: "enforce", Type: proto.ColumnType_JSON, Description: "An array of enforcement configurations which map Sentinel file paths to their enforcement modes. Currently policies only support a single file, so this array will consist of a single element. If the path in the enforcement map does not match the Sentinel policy (<NAME>.sentinel), then the default hard-mandatory will be used."},
			{Name: "organization", Type: proto.ColumnType_JSON, Description: "The organization information."},
		},
	}
}

func listPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_policy.listPolicy", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)
	limit := d.QueryContext.Limit
	options := tfe.PolicyListOptions{
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
		result, err := conn.Policies.List(ctx, organizationName, options)
		if err != nil {
			if isNotFoundError(err) {
				return nil, nil
			}
			plugin.Logger(ctx).Error("tfe_policy.listPolicy", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if plugin.IsCancelled(ctx) {
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

func getPolicy(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_policy.getPolicy", "connection_error", err)
		return nil, err
	}
	result, err := conn.Policies.Read(ctx, d.KeyColumnQuals["id"].GetStringValue())
	if err != nil {
		plugin.Logger(ctx).Error("tfe_policy.getPolicy", "query_error", err)
		return nil, err
	}
	return result, nil
}
