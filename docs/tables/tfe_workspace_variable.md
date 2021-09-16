# Table: tfe_workspace_variable

List variables in a workspace.

Notes:
* List queries require a `workspace_id`.
* Get queries require a `workspace_id` and variable `id`.

## Examples

### Basic info

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

### List environment variables

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

### List sensitive variables

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
