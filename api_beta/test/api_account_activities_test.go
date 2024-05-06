/*
Identity Security Cloud Beta API

Testing AccountActivitiesAPIService

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

func Test_api_beta_AccountActivitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountActivitiesAPIService GetAccountActivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.AccountActivitiesAPI.GetAccountActivity(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountActivitiesAPIService ListAccountActivities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.AccountActivitiesAPI.ListAccountActivities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
