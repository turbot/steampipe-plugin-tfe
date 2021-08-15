# Table: tfe_organization

List organizations the user has permission to see.

## Examples

### List organizations

```sql
select
  *
from
  tfe_organization
```

### Organizations that do not require two factor

```sql
select
  name,
  two_factor_conformant
from
  tfe_organization
where
  not two_factor_conformant
```
