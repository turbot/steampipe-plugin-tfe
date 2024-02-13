---
title: "Steampipe Table: tfe_project - Query Terraform Enterprise Projects using SQL"
description: "Allows users to query Projects in Terraform Enterprise, specifically the ID and name association."
---

# Table: tfe_project - Query Terraform Enterprise Projects using SQL

Terraform Cloud projects let you organize your workspaces into groups. You can structure your projects based on your
organization's resource usage and ownership patterns, such as teams, business units, or services. With Terraform Cloud
Standard Edition, you can give teams access to groups of workspaces using projects.

## Table Usage Guide
The `tfe_project` table provides information about Projects within Terraform Enterprise organization. As a DevOps 
engineer or a system administrator, explore project's details through this table, including its ID, name and organization.
Utilize it in conjunction with `tfe_workspace` table to improve grouping and filtering on workspaces insights.

**Important Notes**
- You must specify the `organization` in the `tfe.spc` file to query this table.

## Examples

### List projects
Explore which projects are in the Terraform Enterprise organization.

```sql+postgres
select
  id,
  name,
  organization,
  organization_name
from
  tfe_project;
```

```sql+sqlite
select
  id,
  name,
  organization,
  organization_name
from
  tfe_project;
```

### List workspace in a specific project
Explore which workspaces belong to a specific project. This can provide an additional filtering layer to analyse 
relative workspaces.

```sql+postgres
select
  w.name 
from
  tfe_workspace as w 
  join
    tfe_project as p 
    on p.id = w.project_id 
where
  p.name = 'my-project';
```

```sql+sqlite
select
  w.name 
from
  tfe_workspace as w 
  join
    tfe_project as p 
    on p.id = w.project_id 
where
  p.name = 'my-project';
```
