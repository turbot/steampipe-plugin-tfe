---
title: "Steampipe Table: tfe_organization_member - Query Terraform Enterprise Organization Members using SQL"
description: "Allows users to query Terraform Enterprise Organization Members, specifically their access levels, user details, and associated organization information."
---

# Table: tfe_organization_member - Query Terraform Enterprise Organization Members using SQL

Terraform Enterprise is a collaborative, scalable, and enterprise-grade version of Terraform that enables teams to work together on infrastructure as code. It provides a centralized workspace for managing Terraform runs, state, and modules, as well as access control and policy enforcement. An Organization Member in Terraform Enterprise refers to a user who is part of a specific organization, with assigned permissions and roles within that organization.

## Table Usage Guide

The `tfe_organization_member` table provides insights into the members of an organization within Terraform Enterprise. As a system administrator or DevOps engineer, explore member-specific details through this table, including their access levels, user details, and associated organization information. Utilize it to manage and monitor user roles and permissions within your organization, ensuring security and compliance.

## Examples

### List all users who are members of the organization
Discover the segments that involve all users who are part of an organization. This could be beneficial in understanding the distribution of users across different organizational structures.

```sql
select
  *
from
  tfe_organization_member;
```

### Check two factor authentication status for each org member
Determine the status of two-factor authentication for each member of an organization. This can help enhance security by identifying members who have not yet enabled this feature.

```sql
select
  username,
  (member -> 'TwoFactor' ->> 'Enabled')::bool as two_factor_enabled
from
  tfe_organization_member;
```