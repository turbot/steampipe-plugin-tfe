![image](https://hub.steampipe.io/images/plugins/turbot/tfe-social-graphic.png)

# Terraform Enterprise Plugin for Steampipe

Use SQL to query workspaces, runs, users and more from Terraform Enterprise.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/tfe)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/tfe/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-tfe/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install tfe
```

Run a query:

```sql
select
  name,
  execution_mode,
  resource_count
from
  tfe_workspace;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-tfe.git
cd steampipe-plugin-tfe
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/tfe.spc
```

Try it!

```
steampipe query
> .inspect tfe
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-tfe/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-tfe/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Terraform Enterprise Plugin](https://github.com/turbot/steampipe-plugin-tfe/labels/help%20wanted)
