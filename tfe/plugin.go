package tfe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-tfe",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		//DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultTransform: transform.FromGo(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"tfe_current_user":        tableTfeCurrentUser(ctx),
			"tfe_oauth_client":        tableTfeOauthClient(ctx),
			"tfe_organization":        tableTfeOrganization(ctx),
			"tfe_organization_member": tableTfeOrganizationMember(ctx),
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
