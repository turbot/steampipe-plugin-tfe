# Table: tfe_current_user

Get information about the current user making the request.

## Examples

### Get user information

```sql
select
  *
from
  tfe_current_user
```

### Check if this is a service account

```sql
select
  username,
  is_service_account
from
  tfe_current_user
```
