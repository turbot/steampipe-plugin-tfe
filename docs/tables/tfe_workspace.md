---
title: "Steampipe Table: tfe_workspace - Query Terraform Enterprise Workspaces using SQL"
description: "Allows users to query Workspaces in Terraform Enterprise, specifically the details related to each workspace including its ID, name, organization, and other associated metadata."
---

# Table: tfe_workspace - Query Terraform Enterprise Workspaces using SQL

Terraform Enterprise is a collaborative, scalable, and enterprise-friendly service provided by HashiCorp that enables teams to use Terraform together. It's designed to suit the collaboration and governance needs of large teams and organizations. Workspaces in Terraform Enterprise are used to manage and track infrastructure deployments.

## Table Usage Guide

The `tfe_workspace` table provides insights into Workspaces within Terraform Enterprise. As a DevOps engineer or a system administrator, explore workspace-specific details through this table, including its ID, name, organization, and other associated metadata. Utilize it to uncover information about workspaces, such as those related to specific organizations, the status of the workspace, and the verification of associated metadata.

## Examples

### List workspaces
Explore all the workspaces available in your Terraform Enterprise setup to better manage and organize your infrastructure as code projects. This is useful for gaining a holistic view of your current workspaces, identifying potential areas for consolidation or reorganization.

```sql+postgres
select
  *
from
  tfe_workspace;
```

```sql+sqlite
select
  *
from
  tfe_workspace;
```

### Get a workspace by ID
Explore the details of a specific workspace in your infrastructure by using its unique identifier. This can help you understand the workspace's current state and configuration, which is useful for troubleshooting or auditing purposes.

```sql+postgres
select
  *
from
  tfe_workspace
where
  id = 'ws-ocYGM1ouZNZWZoUy';
```

```sql+sqlite
select
  *
from
  tfe_workspace
where
  id = 'ws-ocYGM1ouZNZWZoUy';
```

### Get VCS repository settings for workspaces
Explore the configuration of your Version Control System (VCS) repositories linked to your workspaces. This can aid in understanding the specific settings for each repository, such as the associated OAuth token, branch details, and service providers.

```sql+postgres
select
  id,
  name,
  vcs_repo ->> 'Identifier' as vcs_repo_identifier,
  vcs_repo ->> 'OAuthTokenID' as vcs_repo_oauth_token_id,
  vcs_repo ->> 'Branch' as vcs_repo_branch,
  vcs_repo ->> 'DisplayIdentifier' as vcs_repo_display_identifier,
  vcs_repo ->> 'IngressSubmodules' as vcs_repo_ingress_submodules,
  vcs_repo ->> 'RepositoryHTTPURL' as vcs_repo_repository_http_url,
  vcs_repo ->> 'ServiceProvider' as vcs_repo_service_provider
from
  tfe_workspace;
```

```sql+sqlite
select
  id,
  name,
  json_extract(vcs_repo, '$.Identifier') as vcs_repo_identifier,
  json_extract(vcs_repo, '$.OAuthTokenID') as vcs_repo_oauth_token_id,
  json_extract(vcs_repo, '$.Branch') as vcs_repo_branch,
  json_extract(vcs_repo, '$.DisplayIdentifier') as vcs_repo_display_identifier,
  json_extract(vcs_repo, '$.IngressSubmodules') as vcs_repo_ingress_submodules,
  json_extract(vcs_repo, '$.RepositoryHTTPURL') as vcs_repo_repository_http_url,
  json_extract(vcs_repo, '$.ServiceProvider') as vcs_repo_service_provider
from
  tfe_workspace;
```