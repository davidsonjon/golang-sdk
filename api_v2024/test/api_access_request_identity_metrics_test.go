/*
Identity Security Cloud V2024 API

Testing AccessRequestIdentityMetricsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v2024

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_v2024_AccessRequestIdentityMetricsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessRequestIdentityMetricsAPIService GetAccessRequestIdentityMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string
		var requestedObjectId string
		var type_ string

		resp, httpRes, err := apiClient.API_V2024.AccessRequestIdentityMetricsAPI.GetAccessRequestIdentityMetrics(context.Background(), identityId, requestedObjectId, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}