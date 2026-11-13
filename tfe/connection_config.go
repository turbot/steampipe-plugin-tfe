package tfe

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type tfeConfig struct {
	Hostname      *string `hcl:"hostname"`
	Token         *string `hcl:"token"`
	SSLSkipVerify *bool   `hcl:"ssl_skip_verify"`
	Organization  *string `hcl:"organization"`
}

func ConfigInstance() interface{} {
	return &tfeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) tfeConfig {
	if connection == nil || connection.GetConfig() == nil {
		return tfeConfig{}
	}
	config, _ := connection.GetConfig().(tfeConfig)
	return config
}
