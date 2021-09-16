package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfeVariable(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_variable",
		Description: "Workspace variables for the user.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("workspace_id"),
			Hydrate:    listVariable,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"workspace_id", "id"}),
			Hydrate:    getVariable,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the variable."},
			{Name: "workspace_id", Type: proto.ColumnType_STRING, Description: "ID of the workspace.", Transform: transform.FromField("Workspace.ID")},
			{Name: "key", Type: proto.ColumnType_STRING, Description: "Name of the variable."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of the variable."},
			{Name: "category", Type: proto.ColumnType_STRING, Description: "Whether this is a Terraform or environment variable. Valid values are terraform or env."},
			{Name: "sensitive", Type: proto.ColumnType_BOOL, Description: "Whether the value is sensitive. If true then the variable is written once and not visible thereafter. Defaults to false."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the variable."},
			// Other columns
			{Name: "workspace", Type: proto.ColumnType_JSON, Description: "Workspace information."},
		},
	}
}

func listVariable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_variable.listVariable", "connection_error", err)
		return nil, err
	}
	limit := d.QueryContext.Limit
	options := tfe.VariableListOptions{
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
		result, err := conn.Variables.List(ctx, d.KeyColumnQuals["workspace_id"].GetStringValue(), options)
		if err != nil {
			plugin.Logger(ctx).Error("tfe_variable.listVariable", "query_error", err)
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

func getVariable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	workspaceId := d.KeyColumnQuals["workspace_id"].GetStringValue()
	variableId := d.KeyColumnQuals["id"].GetStringValue()

	if workspaceId == "" ||  variableId == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_variable.getVariable", "connection_error", err)
		return nil, err
	}

	result, err := conn.Variables.Read(ctx, workspaceId, variableId)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_variable.getVariable", "query_error", err)
		return nil, err
	}
	return result, nil
}
