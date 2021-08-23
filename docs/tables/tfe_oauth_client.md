# Table: tfe_oauth_client

List oauth clients in a Terraform Enterprise organization.

## Examples

### List oauth clients

```sql
select
  *
from
  tfe_oauth_client;
```

### Get a oauth client by ID

```sql
select
  *
from
  tfe_oauth_client
where
  id = 'oc-JM8tnPzgdo1wM3jy';
```

### Oauth client by age in days

```sql
select
  id,
  created_at,
  date_part('day', age(current_timestamp, created_at)) as age_days
from
  tfe_oauth_client
order by
  age_days desc;
```
