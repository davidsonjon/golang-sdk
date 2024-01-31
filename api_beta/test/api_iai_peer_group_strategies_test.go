/*
IdentityNow Beta API

Testing IAIPeerGroupStrategiesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_beta_IAIPeerGroupStrategiesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIPeerGroupStrategiesAPIService GetPeerGroupOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var strategy string

		resp, httpRes, err := apiClient.API_BETA.IAIPeerGroupStrategiesAPI.GetPeerGroupOutliers(context.Background(), strategy).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
