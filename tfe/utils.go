package tfe

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/go-tfe"
	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*tfe.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "tfe"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*tfe.Client), nil
	}

	var token, organization string
	// setting default hostname
	hostname := "app.terraform.io"

	if os.Getenv("TFE_HOSTNAME") != "" {
		hostname = os.Getenv("TFE_HOSTNAME")
	}

	configFilePath := os.Getenv("TF_CLI_CONFIG_FILE")

	if configFilePath == "" {
		configFilePath = os.Getenv("TERRAFORM_CONFIG")
	}
	if configFilePath != "" {
		config, err := getCliConfig(configFilePath)
		if err != nil {
			return nil, err
		}
		// The hostname should be in format of "app.terraform.io" to match it in credential file
		formatedHostname := hostname
		if strings.HasPrefix(hostname, "https://") {
			formatedHostname = strings.Split(hostname,"/")[2]
		}
		for k, v := range config.Credentials {
			if k == formatedHostname {
				token = v["token"].(string)
			}
		}
	}

	// Default to the env var settings
	if os.Getenv("TFE_TOKEN") != "" {
		token = os.Getenv("TFE_TOKEN")
	}

	sslSkipVerify := strings.ToLower(os.Getenv("TFE_SSL_SKIP_VERIFY")) == "true"

	// Prefer config settings
	tfeConfig := GetConfig(d.Connection)
	if tfeConfig.Hostname != nil {
		hostname = *tfeConfig.Hostname
	}
	if tfeConfig.Token != nil {
		token = *tfeConfig.Token
	}
	if tfeConfig.SSLSkipVerify != nil {
		sslSkipVerify = *tfeConfig.SSLSkipVerify
	}
	if tfeConfig.Organization != nil {
		organization = *tfeConfig.Organization
	}

	// Error if the minimum config is not set
	if token == "" {
		return nil, errors.New("token must be configured")
	}
	if organization == "" {
		return nil, errors.New("organization must be configured")
	}

	// formatting hostname to pass the client
	fullHostName := hostname
	if !strings.HasPrefix(hostname, "https://") {
			fullHostName =  "https://"+hostname+"/"
		}

	// HTTP client and TLS config
	httpClient := tfe.DefaultConfig().HTTPClient
	transport := httpClient.Transport.(*http.Transport)
	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}
	transport.TLSClientConfig.InsecureSkipVerify = sslSkipVerify

	// Create a new TFE client config
	cfg := &tfe.Config{
		Address:    fullHostName,
		Token:      token,
		HTTPClient: httpClient,
	}

	// Create a new TFE client.
	conn, err := tfe.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	conn.RetryServerErrors(true)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func GetOrganizationName(_ context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cacheKey := "GetOrganization"

	// if found in cache, return the result
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(string), nil
	}
	tfeConfig := GetConfig(d.Connection)
	var organization string
	if tfeConfig.Organization != nil {
		organization = *tfeConfig.Organization
	}

	// save to extension cache
	d.ConnectionManager.Cache.Set(cacheKey, organization)
	return organization, nil

}

func isNotFoundError(err error) bool {
	return err.Error() == "resource not found"
}

func getCliConfig(filePath string) (*Config, error) {
	config := &Config{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error reading the CLI config file")
	}

	// Parse the CLI config file content.
	obj, err := hcl.Parse(string(content))
	if err != nil {
		return nil, errors.New("error parsing the CLI config file")
	}

	// Decode the CLI config file content.
	if err := hcl.DecodeObject(&config, obj); err != nil {
		return nil, errors.New("error decoding the CLI config file")
	}
	return config, nil
}

type Config struct {
	Hosts       map[string]*ConfigHost            `hcl:"host"`
	Credentials map[string]map[string]interface{} `hcl:"credentials"`
}

type ConfigHost struct {
	Services map[string]interface{} `hcl:"services"`
}
