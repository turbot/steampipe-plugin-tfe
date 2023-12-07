---
title: "Steampipe Table: tfe_organization - Query Terraform Enterprise Organizations using SQL"
description: "Allows users to query Terraform Enterprise Organizations, specifically retrieving details about each organization, including its name, email, and session timeout."
---

# Table: tfe_organization - Query Terraform Enterprise Organizations using SQL

A Terraform Enterprise Organization represents a collection of users, teams, and workspaces. Within an organization, users can collaborate and manage workspaces together. The organization also provides a context in which resources such as Sentinel policies and SSH keys can be shared.

## Table Usage Guide

The `tfe_organization` table provides insights into organizations within Terraform Enterprise. As a DevOps engineer, explore organization-specific details through this table, including membership, collaboration status, and associated workspaces. Utilize it to uncover information about organizations, such as those with specific team access, the collaboration status between teams, and the management of workspaces.

## Examples

### List organizations
Explore the different organizations within your network. This allows for better management and understanding of the various groups interacting with your systems.

```sql+postgres
select
  *
from
  tfe_organization;
```

```sql+sqlite
select
  *
from
  tfe_organization;
```

### Organizations that do not require two factor
Discover the organizations that do not comply with two-factor authentication. This can be useful for assessing security measures and identifying potential vulnerabilities within your organization.

```sql+postgres
select
  name,
  two_factor_conformant
from
  tfe_organization
where
  not two_factor_conformant;
```

```sql+sqlite
select
  name,
  two_factor_conformant
from
  tfe_organization
where
  not two_factor_conformant;
```