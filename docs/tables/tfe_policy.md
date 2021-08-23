# Table: tfe_policy

List policies in a Terraform Enterprise organization.
## Examples

### List policies

```sql
select
  *
from
  tfe_policy;
```

### List policies which have policy sets

```sql
select
  id,
  name,
  policy_set_count
from
  tfe_policy
where
  policy_set_count > 0;
```

### Get a policy by ID

```sql
select
  *
from
  tfe_policy
where
  id = 'pol-vjgEm4UE6hCsU6a2';
```
