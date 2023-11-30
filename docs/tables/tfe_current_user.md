---
title: "Steampipe Table: tfe_current_user - Query Terraform Enterprise Current User using SQL"
description: "Allows users to query Current Users in Terraform Enterprise, specifically to retrieve detailed information about the authenticated user."
---

# Table: tfe_current_user - Query Terraform Enterprise Current User using SQL

Terraform Enterprise is an advanced version of Terraform that provides collaboration and governance features. It acts as a shared platform for teams to collaborate on infrastructure as code, providing a workspace for sharing and storing Terraform configurations. It also includes a private registry for sharing Terraform modules.

## Table Usage Guide

The `tfe_current_user` table provides insights into the authenticated users within Terraform Enterprise. As a system administrator, you can explore user-specific details through this table, including user's ID, username, email, and other associated metadata. Utilize it to uncover information about users, such as their verification status, two-factor authentication status, and whether they are a site administrator.

## Examples

### Get user information
Explore your current user profile details in Terraform Enterprise. This could be useful for auditing or troubleshooting purposes.

```sql
select
  *
from
  tfe_current_user
```

### Check if this is a service account
Determine if the current user is a service account. This is useful for managing user access and identifying potential security risks.

```sql
select
  username,
  is_service_account
from
  tfe_current_user
```