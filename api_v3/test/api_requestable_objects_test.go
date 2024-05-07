/*
Identity Security Cloud V3 API

Testing RequestableObjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_v3_RequestableObjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RequestableObjectsAPIService ListRequestableObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.RequestableObjectsAPI.ListRequestableObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
