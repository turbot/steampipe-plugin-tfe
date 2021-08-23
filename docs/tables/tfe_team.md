# Table: tfe_team

List teams for the Terraform Enterprise organization.

## Examples

### List teams

```sql
select
  *
from
  tfe_team;
```

### Teams with the most users

```sql
select
  name,
  user_count
from
  tfe_team
order by
  user_count desc
limit 5;
```
