# Table: tfe_organization_member

List organizations the user has permission to see.

## Examples

### List all users who are members of the organization

```sql
select
  *
from
  tfe_organization_member;
```

### Check two factor authentication status for each org member

```sql
select
  username,
  (member -> 'TwoFactor' ->> 'Enabled')::bool as two_factor_enabled
from
  tfe_organization_member;
```
