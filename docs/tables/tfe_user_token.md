# Table: tfe_user_token

List tokens for the current user.

## Examples

### List user tokens

```sql
select
  *
from
  tfe_user_token
```

### Tokens by age in days

```sql
select
  id,
  description,
  created_at,
  date_part('day', age(current_timestamp, created_at)) as age_days
from
  tfe_user_token
order by
  age_days desc
```

### Tokens not used in the last 30 days

```sql
select
  id,
  description,
  last_used_at
from
  tfe_user_token
where
  last_used_at < current_date - interval '30 days'
```
