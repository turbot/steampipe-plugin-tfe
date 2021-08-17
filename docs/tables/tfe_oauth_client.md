# Table: tfe_oauth_client

List oauth clients in a Terraform Enterprise organization.

Notes:
* List queries require an `organization_name`.
* Get queries require a oauth client `id`.

## Examples

### List oauth clients

```sql
select
  *
from
  tfe_oauth_client
where
  organization_name = 'example-org-8a362a';
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
