# Table: tfe_team_member

List all members of all teams for the Terraform Enterprise organization.

Note: An `organization_name` must be provided in all queries to this table.

## Examples

### List all teams and members

```sql
select
  *
from
  tfe_team_member
where
  organization_name = 'example-org-8a362a'
```

### List all members of given team

```sql
select
  *
from
  tfe_team_member
where
  organization_name = 'example-org-8a362a'
  and team_id = 'team-ym4653V1jk9V9FCr'
```

### List teams for a user

```sql
select
  *
from
  tfe_team_member
where
  organization_name = 'example-org-8a362a'
  and user_id = 'user-hHKqSi2HyqZ4iJZs'
```
