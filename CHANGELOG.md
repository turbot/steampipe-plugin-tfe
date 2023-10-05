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
