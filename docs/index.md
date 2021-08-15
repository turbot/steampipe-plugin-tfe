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
where
  organization_name = 'my-org'
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

### Configuration

Installing the latest tfe plugin will create a config file (`~/.steampipe/config/tfe.spc`) with a single connection named `tfe`:

```hcl
connection "tfe" {
  plugin = "tfe"
  token  = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fbd"
}
```

- `token` - API token from Terraform Enterprise.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-tfe
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
