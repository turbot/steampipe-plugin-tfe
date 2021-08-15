# Table: tfe_run

List runs for the workspace.

Notes:
* List queries require a `workspace_id`.
* Get queries require a run `id`.

## Examples

### List runs

```sql
select
  *
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
```

### Runs that errored in the last 24 hrs

```sql
select
  id,
  created_at,
  status
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
  and status = 'errored'
  and created_at > current_timestamp - interval '24 hrs'
```

### Runs that errored in the last 24 hrs

```sql
select
  id,
  created_at,
  status
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
  and status = 'errored'
  and created_at > current_timestamp - interval '24 hrs'
```

### Which users created the most runs?

```sql
select
  created_by ->> 'Username' as username,
  count(*)
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
group by
  username
order by
  count desc
```
