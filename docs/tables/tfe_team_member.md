# Table: tfe_team_member

List all members of all teams for the Terraform Enterprise organization.

## Examples

### List all teams and members

```sql
select
  *
from
  tfe_team_member;
```

### List all members of given team

```sql
select
  *
from
  tfe_team_member
where
  team_id = 'team-ym4653V1jk9V9FCr';
```

### List teams for a user

```sql
select
  *
from
  tfe_team_member
where
  user_id = 'user-hHKqSi2HyqZ4iJZs';
```

### List all service accounts

```sql
select
  *
from
  tfe_team_member
where
  is_service_account;
```
