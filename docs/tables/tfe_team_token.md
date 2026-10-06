---
title: "Steampipe Table: tfe_team_token - Query Terraform Cloud and Enterprise Team Tokens using SQL"
description: "Allows users to query team token metadata in Terraform Cloud and Enterprise, including the team, organization, creation time, last usage time, and expiration time."
---

# Table: tfe_team_token - Query Terraform Cloud and Enterprise Team Tokens using SQL

Team tokens authenticate API requests on behalf of a team. The `tfe_team_token` table lists token metadata for teams in the organization configured for the connection, including multiple tokens per team on servers that support them.

## Table Usage Guide

This table uses the same connection settings as `tfe_team`. It supports Terraform Cloud and Terraform Enterprise through the configured `hostname`. Use credentials with permission to list the organization's teams and read their team tokens.

The table first uses the paginated organization team-token listing API. If the first request returns not found, it falls back to listing teams and reading each team's legacy token, supporting older Terraform Enterprise versions. A not-found response can also indicate that the API hides resources from the caller; use credentials authorized to read team tokens for complete results.

On the legacy API, teams whose token endpoint returns not found are omitted, including teams without a token. Other API errors are returned to the query rather than triggering a fallback. The table is read-only and does not expose token secrets.

The `team_id` column identifies the associated team, while `organization_name` identifies the configured organization. The `last_used_at` column is null for tokens that have never been used. The `expired_at` column is null for tokens without an expiration date or on Terraform Enterprise versions that do not support token expiration.

## Examples

### List team tokens

```sql+postgres
select
  id,
  team_id,
  organization_name,
  description,
  created_at,
  last_used_at,
  expired_at
from
  tfe_team_token;
```

```sql+sqlite
select
  id,
  team_id,
  organization_name,
  description,
  created_at,
  last_used_at,
  expired_at
from
  tfe_team_token;
```

### List tokens with team names

```sql+postgres
select
  token.id,
  team.name as team_name,
  token.organization_name,
  token.created_at,
  token.last_used_at
from
  tfe_team_token as token
  join tfe_team as team on token.team_id = team.id
    and token.organization_name = team.organization_name;
```

```sql+sqlite
select
  token.id,
  team.name as team_name,
  token.organization_name,
  token.created_at,
  token.last_used_at
from
  tfe_team_token as token
  join tfe_team as team on token.team_id = team.id
    and token.organization_name = team.organization_name;
```

### Find tokens that have never been used or have not been used in 30 days

```sql+postgres
select
  id,
  team_id,
  organization_name,
  last_used_at
from
  tfe_team_token
where
  last_used_at is null
  or last_used_at < current_timestamp - interval '30 days';
```

```sql+sqlite
select
  id,
  team_id,
  organization_name,
  last_used_at
from
  tfe_team_token
where
  last_used_at is null
  or datetime(last_used_at) < datetime('now', '-30 days');
```
