## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#105](https://github.com/turbot/steampipe-plugin-tfe/pull/105))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#105](https://github.com/turbot/steampipe-plugin-tfe/pull/105))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#82](https://github.com/turbot/steampipe-plugin-tfe/pull/82))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#82](https://github.com/turbot/steampipe-plugin-tfe/pull/82))

## v0.7.0 [2024-02-13]

_What's new?_

- New table added
  - [tfe_project](https://hub.steampipe.io/plugins/turbot/tfe/tables/tfe_project) ([#42](https://github.com/turbot/steampipe-plugin-tfe/pull/42)) (Thanks [@edebrye](https://github.com/edebrye) for the contribution!)

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#48](https://github.com/turbot/steampipe-plugin-tfe/pull/48))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#48](https://github.com/turbot/steampipe-plugin-tfe/pull/48))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-tfe/blob/main/docs/LICENSE). ([#48](https://github.com/turbot/steampipe-plugin-tfe/pull/48))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#47](https://github.com/turbot/steampipe-plugin-tfe/pull/47))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#30](https://github.com/turbot/steampipe-plugin-tfe/pull/30))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#28](https://github.com/turbot/steampipe-plugin-tfe/pull/28))
- Recompiled plugin with Go version `1.21`. ([#28](https://github.com/turbot/steampipe-plugin-tfe/pull/28))

## v0.4.1 [2023-06-12]

_Bug fixes_

- Fixed the `vcs_repo` column in `tfe_workspace` table to correctly return data instead of `null`. ([#22](https://github.com/turbot/steampipe-plugin-tfe/pull/22))

## v0.4.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#20](https://github.com/turbot/steampipe-plugin-tfe/pull/20))

## v0.3.1 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#15](https://github.com/turbot/steampipe-plugin-tfe/pull/15))

## v0.3.0 [2022-08-30]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v414-2022-08-26) which includes several caching and memory management improvements.
- Recompiled plugin with Go version `1.19`.

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#10](https://github.com/turbot/steampipe-plugin-tfe/pull/10))

## v0.2.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#8](https://github.com/turbot/steampipe-plugin-tfe/pull/8))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#7](https://github.com/turbot/steampipe-plugin-tfe/pull/7))

## v0.1.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#4](https://github.com/turbot/steampipe-plugin-tfe/pull/4))
- Recompiled plugin with Go version 1.17 ([#4](https://github.com/turbot/steampipe-plugin-tfe/pull/4))

## v0.0.2 [2021-09-17]

_Bug fixes_

- Fixed: Brand color is now correct

## v0.0.1 [2021-09-17]

_What's new?_

- New tables added
  - [tfe_current_user](https://hub.steampipe.io/plugins/turbot/tfe/tables/tfe_current_user)
  - [tfe_oauth_client](https://hub.soauth_clientpipe.io/plugins/turbot/tfe/tables/tfe_oauth_client)
  - [tfe_organization](https://hub.sorganizationpipe.io/plugins/turbot/tfe/tables/tfe_organization)
  - [tfe_organization_member](https://hub.steampipe.io/plugins/turbot/tfe/tables/tfe_organization_member)
  - [tfe_run](https://hub.srunpipe.io/plugins/turbot/tfe/tables/tfe_run)
  - [tfe_sentinel_policy](https://hub.ssentinel_policypipe.io/plugins/turbot/tfe/tables/tfe_sentinel_policy)
  - [tfe_ssh_key](https://hub.sssh_keypipe.io/plugins/turbot/tfe/tables/tfe_ssh_key)
  - [tfe_team](https://hub.steampipe.io/plugins/turbot/tfe/tables/tfe_team)
  - [tfe_team_member](https://hub.steam_memberpipe.io/plugins/turbot/tfe/tables/tfe_team_member)
  - [tfe_user_token](https://hub.suser_tokenpipe.io/plugins/turbot/tfe/tables/tfe_user_token)
  - [tfe_variable](https://hub.svariablepipe.io/plugins/turbot/tfe/tables/tfe_variable)
  - [tfe_workspace](https://hub.sworkspacepipe.io/plugins/turbot/tfe/tables/tfe_workspace)
