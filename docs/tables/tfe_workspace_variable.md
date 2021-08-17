# Table: tfe_workspace_variable

List variables in a workspace.

Notes:
* List queries require a `workspace_id`.
* Get queries require a `workspace_id` and variable `id`.

## Examples

### List variables in a workspaces

```sql
select
  id,
  key,
  value,
  category,
  sensitive
from
  tfe_workspace_variable
where
  workspace_id = 'ws-1SWwYqrgF3Aeazmn';
```

### List environment variables in a workspace

```sql
select
  id,
  key,
  value,
  category
from
  tfe_workspace_variable
where
  workspace_id = 'ws-1SWwYqrgF3Aeazmn' and category = 'env';
```

### List sensitive variables in a workspace

```sql
select
  id,
  key,
  value,
  sensitive
from
  tfe_workspace_variable
where
  workspace_id = 'ws-1SWwYqrgF3Aeazmn' and sensitive;
```
