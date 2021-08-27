package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfeRun(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_run",
		Description: "Runs in the workspace.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("workspace_id"),
			Hydrate:    listRun,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRun,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the run."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the run was created."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the run, e.g. pending, planning, applying, applied, errored, etc."},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "Message associated with the run."},
			// Other columns
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "Actions for the run."},
			{Name: "apply", Type: proto.ColumnType_JSON, Description: "Apply phase information from the run."},
			{Name: "configuration_version", Type: proto.ColumnType_JSON, Description: "Configuration record used in the run."},
			{Name: "cost_estimate", Type: proto.ColumnType_JSON, Description: "Cost estimate for the resources in this run."},
			{Name: "created_by", Type: proto.ColumnType_JSON, Description: "Basic information about the user."},
			{Name: "force_cancel_available_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when force cancel will be available for the run."},
			{Name: "has_changes", Type: proto.ColumnType_BOOL, Description: "True if the run has changes."},
			{Name: "is_destroy", Type: proto.ColumnType_BOOL, Description: "Specifies if this plan is a destroy plan, which will destroy all provisioned resources."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Description: "Permissions for the run."},
			{Name: "plan", Type: proto.ColumnType_JSON, Description: "Plan phase information from the run."},
			{Name: "policy_checks", Type: proto.ColumnType_JSON, Description: "Policy check information from the run."},
			{Name: "position_in_queue", Type: proto.ColumnType_INT, Description: "Position in the queue for this run."},
			{Name: "refresh", Type: proto.ColumnType_BOOL, Description: "Whether or not to refresh the state before a plan."},
			{Name: "refresh_only", Type: proto.ColumnType_BOOL, Description: "Whether this run should use the refresh-only plan mode, which will refresh the state without modifying any resources."},
			{Name: "replace_addrs", Type: proto.ColumnType_JSON, Description: "Optional list of resource addresses to be passed to the -replace flag."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "Source of the run request, e.g. tfe-api."},
			{Name: "status_timestamps", Type: proto.ColumnType_JSON, Description: "Timestamps for status changes in the run."},
			{Name: "target_addrs", Type: proto.ColumnType_JSON, Description: "Optional list of resource addresses to be passed to the -target flag."},
			{Name: "workspace_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Workspace.ID"), Description: "Workspace ID that contains the run."},
		},
	}
}

func listRun(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_run.listRun", "connection_error", err)
		return nil, err
	}
	include := "plan,apply,created_by,cost_estimate,configuration_version,configuration_version.ingress_attributes"
	options := tfe.RunListOptions{
		Include: &include,
	}
	pagesLeft := true
	for pagesLeft {
		result, err := conn.Runs.List(ctx, d.KeyColumnQuals["workspace_id"].GetStringValue(), options)
		if err != nil {
			plugin.Logger(ctx).Error("tfe_run.listRun", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
		}
		if result.Pagination.CurrentPage < result.Pagination.TotalPages {
			options.PageNumber = result.Pagination.NextPage
		} else {
			pagesLeft = false
		}
	}
	return nil, nil
}

func getRun(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_run.getRun", "connection_error", err)
		return nil, err
	}
	include := "plan,apply,created_by,cost_estimate,configuration_version,configuration_version.ingress_attributes,workspace"
	result, err := conn.Runs.ReadWithOptions(ctx, d.KeyColumnQuals["id"].GetStringValue(), &tfe.RunReadOptions{Include: include})
	if err != nil {
		plugin.Logger(ctx).Error("tfe_run.getRun", "query_error", err)
		return nil, err
	}
	return result, nil
}
