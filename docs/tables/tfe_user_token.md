---
title: "Steampipe Table: tfe_user_token - Query Terraform Enterprise User Tokens using SQL"
description: "Allows users to query User Tokens in Terraform Enterprise, providing insights into token details, including the associated user, creation time, and last access time."
---

# Table: tfe_user_token - Query Terraform Enterprise User Tokens using SQL

A User Token in Terraform Enterprise represents the authentication details for a specific user. These tokens are used to access the Terraform Enterprise API, and they carry the security credentials for a user authentication request. User tokens are unique to each user and can be created, listed, and deleted through the API.

## Table Usage Guide

The `tfe_user_token` table provides insights into User Tokens within Terraform Enterprise. As a DevOps engineer, explore token-specific details through this table, including the associated user, creation time, and last access time. Utilize it to monitor user activity, manage access control, and ensure secure usage of Terraform Enterprise API.

## Examples

### List user tokens
Determine the areas in which user tokens are being utilized. This can provide insights into user activity and potential security risks, allowing for proactive management and prevention of unauthorized access.

```sql+postgres
select
  *
from
  tfe_user_token;
```

```sql+sqlite
select
  *
from
  tfe_user_token;
```

### Tokens by age in days
Analyze the age of user tokens to understand their longevity and usage patterns. This can aid in identifying outdated or rarely used tokens for potential clean-up or renewal.

```sql+postgres
select
  id,
  description,
  created_at,
  date_part('day', age(current_timestamp, created_at)) as age_days
from
  tfe_user_token
order by
  age_days desc;
```

```sql+sqlite
select
  id,
  description,
  created_at,
  julianday('now') - julianday(created_at) as age_days
from
  tfe_user_token
order by
  age_days desc;
```

### Tokens not used in the last 30 days
Explore which user tokens have been inactive for the past 30 days. This is useful for identifying potentially unused or forgotten tokens that may need to be reviewed or deleted for security purposes.

```sql+postgres
select
  id,
  description,
  last_used_at
from
  tfe_user_token
where
  last_used_at < current_date - interval '30 days';
```

```sql+sqlite
select
  id,
  description,
  last_used_at
from
  tfe_user_token
where
  last_used_at < date('now','-30 day');
```