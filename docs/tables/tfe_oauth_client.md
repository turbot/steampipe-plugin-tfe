# Table: tfe_oauth_client

List OAuth clients in a Terraform Enterprise organization.

## Examples

### Basic info

```sql
select
  *
from
  tfe_oauth_client;
```

### Get OAuth client by ID

```sql
select
  *
from
  tfe_oauth_client
where
  id = 'oc-JM8tnPzgdo1wM3jy';
```

### List OAuth clients sorted by age

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
