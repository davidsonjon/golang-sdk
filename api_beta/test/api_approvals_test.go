/*
Identity Security Cloud Beta API

Testing ApprovalsAPIService

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

func Test_api_beta_ApprovalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApprovalsAPIService GetApproval", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ApprovalsAPI.GetApproval(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApprovalsAPIService GetApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApprovalsAPI.GetApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
