package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfeWorkspace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_workspace",
		Description: "Workspaces for the user.",
		List: &plugin.ListConfig{
			Hydrate:    listWorkspace,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getWorkspace,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the workspace."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the workspace."},
			// Other columns
			{Name: "actions", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "agent_pool", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "agent_pool_id", Type: proto.ColumnType_STRING, Description: "The ID of the agent pool belonging to the workspace's organization."},
			{Name: "allow_destroy_plan", Type: proto.ColumnType_BOOL, Description: "Whether destroy plans can be queued on the workspace."},
			{Name: "apply_duration_average", Type: proto.ColumnType_STRING, Description: "This is the average time runs spend in the apply phase, represented in milliseconds."},
			{Name: "can_queue_destroy_plan", Type: proto.ColumnType_BOOL, Description: "True if the destroy plan can be queued."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the workspace was created."},
			{Name: "current_run", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description for the workspace."},
			{Name: "environment", Type: proto.ColumnType_STRING, Description: "Environment for the workspace runs."},
			{Name: "execution_mode", Type: proto.ColumnType_STRING, Description: "Which execution mode to use. Valid values are remote, local, and agent. When set to local, the workspace will be used for state storage only."},
			{Name: "file_triggers_enabled", Type: proto.ColumnType_BOOL, Description: "Whether to filter runs based on the changed files in a VCS push. If enabled, the working-directory and trigger-prefixes describe a set of paths which must contain changes for a VCS push to trigger a run. If disabled, any push will trigger a run."},
			{Name: "global_remote_state", Type: proto.ColumnType_BOOL, Description: "Whether the workspace should allow all workspaces in the organization to access its state data during runs. If false, then only specifically approved workspaces can access its state."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "True if the workspace is locked."},
			{Name: "migration_environment", Type: proto.ColumnType_STRING, Description: ""},
			// Deprecated - {Name: "operations", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "organization", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the workspace."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "plan_duration_average", Type: proto.ColumnType_STRING, Description: "This is the average time runs spend in the plan phase, represented in milliseconds."},
			{Name: "policy_check_failures", Type: proto.ColumnType_INT, Description: "Reports the number of run failures resulting from a policy check failure."},
			{Name: "queue_all_runs", Type: proto.ColumnType_BOOL, Description: "Whether runs should be queued immediately after workspace creation. When set to false, runs triggered by a VCS change will not be queued until at least one run is manually queued."},
			{Name: "resource_count", Type: proto.ColumnType_INT, Description: "Number of resources in the workspace."},
			{Name: "run_failures", Type: proto.ColumnType_INT, Description: "Reports the number of failed runs."},
			// What is this? {Name: "runs_count", Type: proto.ColumnType_INT, Description: ""},
			{Name: "source_name", Type: proto.ColumnType_STRING, Description: "A friendly name for the application or client creating this workspace."},
			{Name: "source_url", Type: proto.ColumnType_STRING, Description: "A URL for the application or client creating this workspace. This can be the URL of a related resource in another app, or a link to documentation or other info about the client."},
			{Name: "speculative_enabled", Type: proto.ColumnType_BOOL, Description: "Whether this workspace allows automatic speculative plans. Setting this to false prevents Terraform Cloud from running plans on pull requests, which can improve security if the VCS repository is public or includes untrusted contributors."},
			{Name: "ssh_key", Type: proto.ColumnType_JSON, Description: "SSH key assigned to the workspace."},
			// What is this? {Name: "structured_run_output_enabled", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "terraform_version", Type: proto.ColumnType_STRING, Description: "The version of Terraform to use for this workspace. Upon creating a workspace, the latest version is selected unless otherwise specified (e.g. 0.11.1)."},
			{Name: "trigger_prefixes", Type: proto.ColumnType_JSON, Description: "List of repository-root-relative paths which should be tracked for changes, in addition to the working directory."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the workspace was last updated."},
			{Name: "vcs_repo", Type: proto.ColumnType_JSON, Description: "Settings for the workspace's VCS repository. If omitted, the workspace is created without a VCS repo."},
			{Name: "working_directory", Type: proto.ColumnType_STRING, Description: "A relative path that Terraform will execute within. This defaults to the root of your repository and is typically set to a subdirectory matching the environment when multiple environments exist within the same repository."},
		},
	}
}

func listWorkspace(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_workspace.listWorkspace", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)
	include := "current_run"
	limit := d.QueryContext.Limit
	options := tfe.WorkspaceListOptions{
		Include: &include,
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
		result, err := conn.Workspaces.List(ctx, organizationName, options)
		if err != nil {
			plugin.Logger(ctx).Error("tfe_workspace.listWorkspace", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
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

func getWorkspace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.getWorkspace", "connection_error", err)
		return nil, err
	}
	opts := tfe.WorkspaceReadOptions{Include: "current_run"}
	result, err := conn.Workspaces.ReadByIDWithOptions(ctx, d.KeyColumnQuals["id"].GetStringValue(), &opts)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.getWorkspace", "query_error", err)
		return nil, err
	}
	return result, nil
}
