---
title: "Steampipe Table: tfe_team_member - Query Terraform Enterprise Team Members using SQL"
description: "Allows users to query Team Members in Terraform Enterprise, specifically the membership details, providing insights into team structure and access levels."
---

# Table: tfe_team_member - Query Terraform Enterprise Team Members using SQL

Terraform Enterprise is a collaborative, scalable, and enterprise-friendly service for managing infrastructure as code. It provides teams with a centralized platform to collaborate on infrastructure and manage access to resources. Team Members in Terraform Enterprise are the users assigned to a specific team, with permissions and access levels defined by their team's settings.

## Table Usage Guide

The `tfe_team_member` table provides insights into Team Members within Terraform Enterprise. As an Infrastructure Manager, explore member-specific details through this table, including team associations, user IDs, and access permissions. Utilize it to uncover information about team members, such as their roles, the teams they're part of, and their access levels within those teams.

## Examples

### List all teams and members
Discover the segments that include all teams and their respective members. This can be useful to gain an overview of team composition and structure in your organization.

```sql+postgres
select
  *
from
  tfe_team_member;
```

```sql+sqlite
select
  *
from
  tfe_team_member;
```

### List all members of given team
Explore which individuals are part of a specific team, assisting in team management and understanding team composition. This can be particularly useful when auditing team membership or planning resource allocation.

```sql+postgres
select
  *
from
  tfe_team_member
where
  team_id = 'team-ym4653V1jk9V9FCr';
```

```sql+sqlite
select
  *
from
  tfe_team_member
where
  team_id = 'team-ym4653V1jk9V9FCr';
```

### List teams for a user
Discover the teams a specific user is a part of, which can be useful for assessing their roles and responsibilities within the organization.

```sql+postgres
select
  *
from
  tfe_team_member
where
  user_id = 'user-hHKqSi2HyqZ4iJZs';
```

```sql+sqlite
select
  *
from
  tfe_team_member
where
  user_id = 'user-hHKqSi2HyqZ4iJZs';
```

### List all service accounts
Explore which team members are associated with service accounts. This can be useful for managing access permissions and ensuring security protocols are being followed.

```sql+postgres
select
  *
from
  tfe_team_member
where
  is_service_account;
```

```sql+sqlite
select
  *
from
  tfe_team_member
where
  is_service_account = 1;
```