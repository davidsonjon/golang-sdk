/*
IdentityNow V3 API

Testing RequestableObjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_v3_RequestableObjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RequestableObjectsAPIService ListRequestableObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.RequestableObjectsAPI.ListRequestableObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
