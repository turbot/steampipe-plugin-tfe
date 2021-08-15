# Table: tfe_workspace

List workspaces in the Terraform Enterprise organization.

Notes:
* List queries require an `organization_name`.
* Get queries require a workspace `id`.

## Examples

### List workspaces

```sql
select
  *
from
  tfe_workspace
where
  organization_name = 'example-org-8a362a'
```

### Get a workspace by ID

```sql
select
  *
from
  tfe_workspace
where
  id = 'ws-ocYGM1ouZNZWZoUy'
```
