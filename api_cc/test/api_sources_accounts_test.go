/*
IdentityNow cc (private) APIs

Testing SourcesAccountsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_cc

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_cc_SourcesAccountsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourcesAccountsAPIService ExportAccountFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_CC.SourcesAccountsAPI.ExportAccountFeed(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
