package tfe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-tfe",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "organization_name",
				Hydrate: GetOrganizationName,
			},
		},
		//DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultTransform: transform.FromGo(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: func(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
					return isNotFoundError(err)
				},
			},
		},
		TableMap: map[string]*plugin.Table{
			"tfe_current_user":        tableTfeCurrentUser(ctx),
			"tfe_oauth_client":        tableTfeOauthClient(ctx),
			"tfe_organization":        tableTfeOrganization(ctx),
			"tfe_organization_member": tableTfeOrganizationMember(ctx),
			"tfe_project":             tableTfeProject(ctx),
			"tfe_run":                 tableTfeRun(ctx),
			"tfe_sentinel_policy":     tableTfeSentinelPolicy(ctx),
			"tfe_ssh_key":             tableTfeSshKey(ctx),
			"tfe_team":                tableTfeTeam(ctx),
			"tfe_team_member":         tableTfeTeamMember(ctx),
			"tfe_user_token":          tableTfeUserToken(ctx),
			"tfe_variable":            tableTfeVariable(ctx),
			"tfe_workspace":           tableTfeWorkspace(ctx),
		},
	}
	return p
}
