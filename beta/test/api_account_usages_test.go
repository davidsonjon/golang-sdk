/*
IdentityNow Beta API

Testing AccountUsagesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk"
)

func Test_beta_AccountUsagesApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountUsagesApiService GetUsagesByAccountId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.BETA.AccountUsagesApi.GetUsagesByAccountId(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
