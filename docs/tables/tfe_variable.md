---
title: "Steampipe Table: tfe_variable - Query Terraform Enterprise Variables using SQL"
description: "Allows users to query Variables in Terraform Enterprise, specifically the key-value pairs that define the infrastructure and configurations."
---

# Table: tfe_variable - Query Terraform Enterprise Variables using SQL

A Terraform Enterprise Variable is a key-value pair that defines the infrastructure and configurations. Variables in Terraform Enterprise provide a flexible and dynamic way of managing and organizing resources. They are used to customize aspects of the execution plan and to store sensitive information securely.

## Table Usage Guide

The `tfe_variable` table provides insights into Variables within Terraform Enterprise. As a DevOps engineer, explore variable-specific details through this table, including key-value pairs, category, and associated workspace. Utilize it to manage and organize your resources, customize your execution plan, and securely store sensitive information.

**Important Notes**
- You must specify the `workspace_id` in the `where` clause to query this table.

## Examples

### Basic info
Discover the segments that contain sensitive information within a specific workspace. This is beneficial in maintaining data security and ensuring only authorized personnel have access to sensitive data.

```sql
select
  id,
  key,
  value,
  category,
  sensitive
from
  tfe_variable
where
  workspace_id = 'ws-1SWwYqrgF3Aeazmn';
```

### List environment variables
Explore which environment variables are associated with a specific workspace, allowing you to understand and manage the settings and configurations for that workspace.

```sql
select
  id,
  key,
  value,
  category
from
  tfe_variable
where
  workspace_id = 'ws-1SWwYqrgF3Aeazmn' and category = 'env';
```

### List sensitive variables
Analyze the settings to understand which variables within a specific workspace are sensitive. This can aid in maintaining security and confidentiality within your system.

```sql
select
  id,
  key,
  value,
  sensitive
from
  tfe_variable
where
  workspace_id = 'ws-1SWwYqrgF3Aeazmn' and sensitive;
```