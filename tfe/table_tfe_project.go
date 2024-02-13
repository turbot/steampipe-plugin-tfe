package tfe

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTfeProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_project",
		Description: "Projects for the workspaces.",
		List: &plugin.ListConfig{
			Hydrate: listProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the project."},
			{Name: "organization", Type: proto.ColumnType_JSON, Description: "Organization details that the project belongs to."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the project."},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_project.listProject", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)
	limit := d.QueryContext.Limit
	options := tfe.ProjectListOptions{
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
		result, err := conn.Projects.List(ctx, organizationName, &options)
		if err != nil {
			plugin.Logger(ctx).Error("tfe_project.listProject", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
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

func getProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_project.getProject", "connection_error", err)
		return nil, err
	}
	result, err := conn.Projects.Read(ctx, d.EqualsQuals["id"].GetStringValue())
	
	if err != nil {
		plugin.Logger(ctx).Error("tfe_project.getProject", "query_error", err)
		return nil, err
	}
	return result, nil
}
