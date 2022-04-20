package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-tfe/tfe"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: tfe.Plugin})
}
