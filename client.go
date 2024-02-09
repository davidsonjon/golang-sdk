/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

package sailpoint

import (
	"regexp"

	"github.com/hashicorp/go-retryablehttp"
	beta "github.com/sailpoint-oss/golang-sdk/v2/api_beta"
	v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the IdentityNow V3 API API v3.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	V3    *v3.APIClient
	Beta  *beta.APIClient
	token string
}

type service struct {
	client     *v3.APIClient
	betaClient *beta.APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = retryablehttp.NewClient()
	}

	c := &APIClient{}

	CV3 := v3.NewConfiguration(cfg.ClientConfiguration.ClientId, cfg.ClientConfiguration.ClientSecret, cfg.ClientConfiguration.BaseURL+"/v3", cfg.ClientConfiguration.TokenURL, cfg.ClientConfiguration.Token)
	CBeta := beta.NewConfiguration(cfg.ClientConfiguration.ClientId, cfg.ClientConfiguration.ClientSecret, cfg.ClientConfiguration.BaseURL+"/beta", cfg.ClientConfiguration.TokenURL, cfg.ClientConfiguration.Token)

	CV3.HTTPClient = cfg.HTTPClient
	CBeta.HTTPClient = cfg.HTTPClient

	c.V3 = v3.NewAPIClient(CV3)
	c.Beta = beta.NewAPIClient(CBeta)

	// API Services

	return c
}
