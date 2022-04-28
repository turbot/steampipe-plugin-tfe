package tfe

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/go-tfe"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*tfe.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "tfe"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*tfe.Client), nil
	}

	var organization string
	// Default to the env var settings
	hostname := os.Getenv("TFE_HOSTNAME")
	token := os.Getenv("TFE_TOKEN")
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

	// HTTP client and TLS config
	httpClient := tfe.DefaultConfig().HTTPClient
	transport := httpClient.Transport.(*http.Transport)
	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}
	transport.TLSClientConfig.InsecureSkipVerify = sslSkipVerify

	// Default to Terraform Cloud host service
	if hostname == "" {
		hostname = "https://app.terraform.io"
	}

	if !strings.HasPrefix(hostname, "https://") {
		hostname = "https://" + hostname
	}

	// Create a new TFE client config
	cfg := &tfe.Config{
		Address:    hostname,
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
