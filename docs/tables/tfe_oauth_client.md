---
title: "Steampipe Table: tfe_oauth_client - Query Terraform Enterprise OAuth Clients using SQL"
description: "Allows users to query OAuth Clients in Terraform Enterprise, specifically the client credentials, providing insights into OAuth client configurations and associated details."
---

# Table: tfe_oauth_client - Query Terraform Enterprise OAuth Clients using SQL

An OAuth Client in Terraform Enterprise is a fundamental building block when working with OAuth-based connections in the platform. It represents the applications that are registered to use the OAuth 2.0 authorization framework. OAuth Clients are responsible for managing the OAuth tokens and their associated scopes, providing a secure way to authenticate users and interact with the Terraform Enterprise API.

## Table Usage Guide

The `tfe_oauth_client` table provides insights into OAuth Clients within Terraform Enterprise. As a security administrator, explore OAuth client-specific details through this table, including client credentials, token details, and associated metadata. Utilize it to uncover information about OAuth clients, such as those with specific permissions, the relationships between clients, and the verification of token scopes.

## Examples

### Basic info
Explore the OAuth client details within your infrastructure to gain insights into their configuration and usage. This can be useful in understanding the client's behavior and identifying any potential issues or areas for improvement.

```sql+postgres
select
  *
from
  tfe_oauth_client;
```

```sql+sqlite
select
  *
from
  tfe_oauth_client;
```

### Get OAuth client by ID
Explore which OAuth client corresponds to a specific ID to manage access and permissions more effectively. This can be useful in scenarios where you need to understand the access granted to a particular client or troubleshoot issues related to client permissions.

```sql+postgres
select
  *
from
  tfe_oauth_client
where
  id = 'oc-JM8tnPzgdo1wM3jy';
```

```sql+sqlite
select
  *
from
  tfe_oauth_client
where
  id = 'oc-JM8tnPzgdo1wM3jy';
```

### List OAuth clients sorted by age
Analyze the settings to understand the age of your OAuth clients, allowing you to prioritize updates or maintenance based on their age. This can be useful in managing the lifecycle of your OAuth clients and ensuring older clients are still functioning properly.

```sql+postgres
select
  id,
  created_at,
  date_part('day', age(current_timestamp, created_at)) as age_days
from
  tfe_oauth_client
order by
  age_days desc;
```

```sql+sqlite
select
  id,
  created_at,
  julianday('now') - julianday(created_at) as age_days
from
  tfe_oauth_client
order by
  age_days desc;
```