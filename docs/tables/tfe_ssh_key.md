# Table: tfe_ssh_key

List ssh keys in a Terraform Enterprise organization.

Notes:
* List queries require an `organization_name`.
* Get queries require a workspace `id`.

## Examples

### List ssh keys in a organization

```sql
select
  *
from
  tfe_ssh_key
where
  organization_name = 'example-org-872e34';
```

### Get a ssh keys by id

```sql
select
  *
from
  tfe_ssh_key
where
  id = 'sshkey-1NSDCvowf3WtbStu';
```
