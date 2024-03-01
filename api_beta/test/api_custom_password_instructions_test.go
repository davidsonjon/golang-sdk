/*
IdentityNow Beta API

Testing CustomPasswordInstructionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_beta_CustomPasswordInstructionsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomPasswordInstructionsAPIService CreateCustomPasswordInstructions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomPasswordInstructionsAPI.CreateCustomPasswordInstructions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPasswordInstructionsAPIService DeleteCustomPasswordInstructions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pageId string

		httpRes, err := apiClient.API_BETA.CustomPasswordInstructionsAPI.DeleteCustomPasswordInstructions(context.Background(), pageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPasswordInstructionsAPIService GetCustomPasswordInstructions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pageId string

		resp, httpRes, err := apiClient.API_BETA.CustomPasswordInstructionsAPI.GetCustomPasswordInstructions(context.Background(), pageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
