# Table: tfe_organization_member

List organizations the user has permission to see.

Note: An `organization_name` must be provided in all queries to this table.

## Examples

### List all users who are members of the organization

```sql
select
  *
from
  tfe_organization_member
where
  organization_name = 'example-org-6a268a'
```

### Check two factor authentication status for each org member

```sql
select
  username,
  (member -> 'TwoFactor' ->> 'Enabled')::bool as two_factor_enabled
from
  tfe_organization_member
where
  organization_name = 'example-org-6a268a'
```
