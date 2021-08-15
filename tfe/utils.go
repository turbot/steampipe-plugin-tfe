package tfe

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/go-tfe"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*tfe.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "tfe"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*tfe.Client), nil
	}

	// Default to the env var settings
	hostname := os.Getenv("TFE_HOSTNAME")
	token := os.Getenv("TFE_TOKEN")
	sslSkipVerify := strings.ToLower(os.Getenv("TFE_SSL_SKIP_VERIFY")) == "true"

	// Prefer config settings
	tfeConfig := GetConfig(d.Connection)
	if &tfeConfig != nil {
		if tfeConfig.Hostname != nil {
			hostname = *tfeConfig.Hostname
		}
		if tfeConfig.Token != nil {
			token = *tfeConfig.Token
		}
		if tfeConfig.SSLSkipVerify != nil {
			sslSkipVerify = *tfeConfig.SSLSkipVerify
		}
	}

	// Error if the minimum config is not set
	if hostname == "" {
		hostname = "https://app.terraform.io/"
	}
	if token == "" {
		return nil, errors.New("token must be configured")
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

func isNotFoundError(err error) bool {
	return err.Error() == "resource not found"
}
