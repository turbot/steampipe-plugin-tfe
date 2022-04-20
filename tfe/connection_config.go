package tfe

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type tfeConfig struct {
	Hostname      *string `cty:"hostname"`
	Token         *string `cty:"token"`
	SSLSkipVerify *bool   `cty:"ssl_skip_verify"`
	Organization  *string `cty:"organization"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"hostname": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
	"ssl_skip_verify": {
		Type: schema.TypeBool,
	},
	"organization": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &tfeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) tfeConfig {
	if connection == nil || connection.Config == nil {
		return tfeConfig{}
	}
	config, _ := connection.Config.(tfeConfig)
	return config
}
