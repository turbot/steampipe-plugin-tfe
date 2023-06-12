# Table: tfe_workspace

List workspaces in the Terraform Enterprise organization.

## Examples

### List workspaces

```sql
select
  *
from
  tfe_workspace;
```

### Get a workspace by ID

```sql
select
  *
from
  tfe_workspace
where
  id = 'ws-ocYGM1ouZNZWZoUy';
```

### Get VCS repository settings for workspaces

```sql
select
  id,
  name,
  vcs_repo ->> 'Identifier' as vcs_repo_identifier,
  vcs_repo ->> 'OAuthTokenID' as vcs_repo_oauth_token_id,
  vcs_repo ->> 'Branch' as vcs_repo_branch,
  vcs_repo ->> 'DisplayIdentifier' as vcs_repo_display_identifier,
  vcs_repo ->> 'IngressSubmodules' as vcs_repo_ingress_submodules,
  vcs_repo ->> 'RepositoryHTTPURL' as vcs_repo_repository_http_url,
  vcs_repo ->> 'ServiceProvider' as vcs_repo_service_provider
from
  tfe_workspace;
```
