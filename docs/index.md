---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/tfe.svg"
brand_color: "#4e7e14"
display_name: "Terraform Enterprise"
short_name: "tfe"
description: "Steampipe plugin to query resources, users and more from Terraform Enterprise."
og_description: "Query Terraform Enterprise with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/tfe-social-graphic.png"
---

# Terraform Enterprise + Steampipe

[Terraform Enterprise](https://www.terraform.io/cloud) is a cloud hosting company that provides virtual private servers and other infrastructure services.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List workspaces in your Terraform Enterprise organization:

```sql
select
  name,
  execution_mode,
  resource_count
from
  tfe_workspace
```

```
+-----------------+----------------+----------------+
| name            | execution_mode | resource_count |
+-----------------+----------------+----------------+
| getting-started | remote         | 5              |
+-----------------+----------------+----------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/tfe/tables)**

## Get started

### Install

Download and install the latest Terraform Enterprise plugin:

```bash
steampipe plugin install tfe
```

### Credentials

| Item        | Description                                                                  |
| :---------- | :--------------------------------------------------------------------------- |
| Credentials | Terraform Cloud/Enterprise requires a [token](https://www.terraform.io/docs/cloud/users-teams-organizations/api-tokens.html) for all requests. |
| Radius      | Each connection represents a single Terraform Cloud/Enterprise account. |
| Resolution  |  1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/tfe.spc`).<br />2. Credentials specified in environment variables e.g. `TFE_TOKEN`.<br />3. Credentials in the credential file (for example `~/.terraform.d/credentials.tfrc.json` on Linux systems) for the path specified in the `TF_CLI_CONFIG_FILE` or `TERRAFORM_CONFIG` environment variable.|

### Configuration

Installing the latest tfe plugin will create a config file (`~/.steampipe/config/tfe.spc`) with a single connection named `tfe`:

```hcl
connection "tfe" {
  plugin = "tfe"
  token  = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fbd"
  organization = "example-org-872e34"
}
```

- `token` - API token from Terraform Enterprise.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-tfe
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)

## Multi-Organization Connections


You may create multiple tfe connections:
```hcl
connection "tfe_01 {
  plugin        = "tfe"
  token         = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fbd"
  organization  = "example-org-872e34"
}

connection "tfe_02 {
  plugin      = "tfe"
  token         = "6a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fcd"
  organization  = "example-org-123f45"
}

connection "tfe_03 {
  plugin      = "tfe"
  token         = "7a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fef"
  organization  = "example-org-123f90"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html).  As such, you can use qualified table names to query a specific connection:

```sql
select * from tfe_02.tfe_organization_member
```

Alternatively, can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/using-steampipe/managing-connections#setting-the-search-path):
```sql
select * from tfe_organization_member
```


You can multi-organization connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators).  Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection:

```
connection "tfe_all {
  plugin      = "tfe"
  type        = "aggregator"
  connections = ["tfe_01", "tfe_02", "tfe_03"]
}
```

Querying tables from this connection will return results from the `tfe_01`, `tfe_02`, and `tfe_03` connections:
```sql
select * from tfe_all.tfe_workspace
```

Steampipe supports the `*` wildcard in the connection names.  For example, to aggregate all the TFE plugin connections whose names begin with `tfe_`:

```hcl
connection "tfe_all" {
  type        = "aggregator"
  plugin      = "tfe"
  connections = ["tfe_*"]
}
```

Aggregators are powerful, but they are not infinitely scalable. Like any other steampipe connection, they query APIs and are subject to API limits and throttling.
