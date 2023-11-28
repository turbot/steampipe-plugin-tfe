# Table: tfe_project

List projects in the Terraform Enterprise organization.

## Examples

### List projects

```sql
select
  *
from
  tfe_project;
```

### Get a project by id

```sql
select
  *
from
  tfe_project
where
  id = 'prj-abcdefgh';
```
