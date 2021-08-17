# Table: tfe_oauth_client

List policies in a Terraform Enterprise organization.

Notes:
* List queries require an `organization_name`.
* Get queries require a policy `id`.

## Examples

### List policies

```sql
select
  *
from
  tfe_policy
where
  organization_name = 'example-org-872e34';
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
  organization_name = 'example-org-872e34' and policy_set_count > 0;
```

### Get a oauth client by ID

```sql
select
  *
from
  tfe_policy
where
  id = 'pol-vjgEm4UE6hCsU6a2';
```
