/*
Identity Security Cloud Beta API

Testing AccountUsagesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk/v2"
)

func Test_api_beta_AccountUsagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountUsagesAPIService GetUsagesByAccountId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountUsagesAPI.GetUsagesByAccountId(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
