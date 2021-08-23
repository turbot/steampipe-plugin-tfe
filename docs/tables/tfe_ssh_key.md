# Table: tfe_ssh_key

List ssh keys in a Terraform Enterprise organization.

## Examples

### List ssh keys in a organization

```sql
select
  *
from
  tfe_ssh_key;
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
