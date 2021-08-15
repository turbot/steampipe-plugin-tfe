# Table: tfe_team

List teams for the Terraform Enterprise organization.

Notes:
* List queries require an `organization_name`.
* Get queries require a team `id`.

## Examples

### List teams

```sql
select
  *
from
  tfe_team
where
  organization_name = 'example-org-8a362a'
```

### Teams with the most users

```sql
select
  name,
  user_count
from
  tfe_team
where
  organization_name = 'example-org-8a362a'
order by
  user_count desc
limit 5
```
