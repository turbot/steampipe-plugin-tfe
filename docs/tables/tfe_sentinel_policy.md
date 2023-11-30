---
title: "Steampipe Table: tfe_sentinel_policy - Query Terraform Enterprise Sentinel Policies using SQL"
description: "Allows users to query Sentinel Policies in Terraform Enterprise, specifically the enforcement levels, policy code and associated metadata, providing insights into policy management and potential security configurations."
---

# Table: tfe_sentinel_policy - Query Terraform Enterprise Sentinel Policies using SQL

Sentinel Policies in Terraform Enterprise (TFE) are a set of rules that define the behavior of resources and modules in a Terraform workspace. They provide a means to enforce certain standards and best practices across your organization. This includes rules for security, compliance, and cost management that are enforced when making changes to infrastructure.

## Table Usage Guide

The `tfe_sentinel_policy` table provides insights into Sentinel Policies within Terraform Enterprise. As a DevOps engineer or security analyst, explore policy-specific details through this table, including enforcement levels, policy code, and associated metadata. Utilize it to uncover information about policies, such as those with strict enforcement levels, the specific rules defined in the policy code, and the overall management of policies within your Terraform workspace.

## Examples

### Basic info
Explore the policies in your Sentinel infrastructure to understand the rules that are currently in place. This can help in assessing your security posture and identifying areas for improvement.

```sql
select
  *
from
  tfe_sentinel_policy;
```

### List policies that have policy sets
Discover the Sentinel policies that are associated with one or more policy sets. This can be useful to understand the application of these policies across different sets, helping to manage and optimize policy usage.

```sql
select
  id,
  name,
  policy_set_count
from
  tfe_sentinel_policy
where
  policy_set_count > 0;
```

### Get policy by ID
Explore the specific details of a policy by using its unique identifier. This is particularly useful when you need to quickly assess the characteristics of a single policy in your Terraform Enterprise environment.

```sql
select
  *
from
  tfe_sentinel_policy
where
  id = 'pol-vjgEm4UE6hCsU6a2';
```