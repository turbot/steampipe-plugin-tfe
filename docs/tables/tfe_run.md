---
title: "Steampipe Table: tfe_run - Query Terraform Cloud Runs using SQL"
description: "Allows users to query Runs in Terraform Cloud, specifically providing details about the run's status, timestamps, and associated resources."
---

# Table: tfe_run - Query Terraform Cloud Runs using SQL

A Run in Terraform Cloud represents a single run of Terraform in a workspace. It includes details about the run's status, timestamps, and associated resources such as workspace, configuration version, and plan. Runs are a key component of Terraform Cloud's collaboration features, as they allow users to propose, review, and approve infrastructure changes.

## Table Usage Guide

The `tfe_run` table provides insights into Runs within Terraform Cloud. As a DevOps engineer, explore run-specific details through this table, including status, timestamps, and associated resources. Utilize it to understand the state of infrastructure changes, the progress of runs, and the details of associated resources.

**Important Notes**
- You must specify the `workspace_id` in the `where` clause to query this table.

## Examples

### List runs
Uncover the details of all the runs associated with a specific workspace. This can be useful when you need to monitor or review the progress and status of all tasks within that workspace.

```sql+postgres
select
  *
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx';
```

```sql+sqlite
select
  *
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx';
```

### Runs that errored in the last 24 hrs
Identify instances where workflow runs have encountered errors in the past day. This is useful for troubleshooting recent issues and understanding the frequency of errors in your workflows.

```sql+postgres
select
  id,
  created_at,
  status
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
  and status = 'errored'
  and created_at > current_timestamp - interval '24 hrs';
```

```sql+sqlite
select
  id,
  created_at,
  status
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
  and status = 'errored'
  and created_at > datetime('now','-24 hours');
```

### Which users created the most runs?
Discover the users who have initiated the most operations within a specific workspace. This can help identify the most active users and understand usage patterns.

```sql+postgres
select
  created_by ->> 'Username' as username,
  count(*)
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
group by
  username
order by
  count desc;
```

```sql+sqlite
select
  json_extract(created_by, '$.Username') as username,
  count(*)
from
  tfe_run
where
  workspace_id = 'ws-ocKJU1ouZNZWZoUx'
group by
  username
order by
  count(*) desc;
```