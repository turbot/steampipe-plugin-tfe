---
title: "Steampipe Table: tfe_team - Query Terraform Enterprise Teams using SQL"
description: "Allows users to query Teams in Terraform Enterprise, specifically the team's ID, name, organization, visibility, and user count, providing insights into team configurations and structures."
---

# Table: tfe_team - Query Terraform Enterprise Teams using SQL

Terraform Enterprise Teams are a resource within Terraform that allows you to manage and organize users into different groups. Teams provide a way to assign specific permissions and roles to a group of users, simplifying access management across the organization. They are particularly useful in large organizations where managing individual user permissions can be complex and time-consuming.

## Table Usage Guide

The `tfe_team` table provides insights into Teams within Terraform Enterprise. As an Infrastructure Engineer, explore team-specific details through this table, including team names, visibility settings, and user counts. Utilize it to understand team structures, manage access permissions, and streamline the organization of users within your Terraform Enterprise setup.

## Examples

### List teams
Explore which teams are present within your organization to understand the distribution of resources and tasks. This can help in managing your resources more efficiently by identifying any gaps or overlaps.

```sql+postgres
select
  *
from
  tfe_team;
```

```sql+sqlite
select
  *
from
  tfe_team;
```

### Teams with the most users
Discover the teams that have the highest number of users. This is useful for understanding which teams are the largest and may need additional resources or management.

```sql+postgres
select
  name,
  user_count
from
  tfe_team
order by
  user_count desc
limit 5;
```

```sql+sqlite
select
  name,
  user_count
from
  tfe_team
order by
  user_count desc
limit 5;
```