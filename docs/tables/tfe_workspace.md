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
