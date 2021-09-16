# Table: tfe_sentinel_policy

List policies in a Terraform Enterprise organization.

## Examples

### Basic info

```sql
select
  *
from
  tfe_sentinel_policy;
```

### List policies that have policy sets

```sql
select
  id,
  name,
  policy_set_count
from
  tfe_sentinel_policy
where
  policy_set_count > 0;
```

### Get policy by ID

```sql
select
  *
from
  tfe_sentinel_policy
where
  id = 'pol-vjgEm4UE6hCsU6a2';
```
