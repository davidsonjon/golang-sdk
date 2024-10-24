/*
Identity Security Cloud V2024 API

Testing TenantAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v2024

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk/v2"
)

func Test_api_v2024_TenantAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TenantAPIService GetTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TenantAPI.GetTenant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
