/*
IdentityNow Beta API

Testing SourceUsagesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_beta_SourceUsagesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourceUsagesAPIService GetStatusBySourceId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourceUsagesAPI.GetStatusBySourceId(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourceUsagesAPIService GetUsagesBySourceId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourceUsagesAPI.GetUsagesBySourceId(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
