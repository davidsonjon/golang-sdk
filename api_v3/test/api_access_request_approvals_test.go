/*
Identity Security Cloud V3 API

Testing AccessRequestApprovalsAPIService

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

func Test_api_v3_AccessRequestApprovalsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessRequestApprovalsAPIService ApproveAccessRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		resp, httpRes, err := apiClient.API_V3.AccessRequestApprovalsAPI.ApproveAccessRequest(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsAPIService ForwardAccessRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		resp, httpRes, err := apiClient.API_V3.AccessRequestApprovalsAPI.ForwardAccessRequest(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsAPIService GetAccessRequestApprovalSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.AccessRequestApprovalsAPI.GetAccessRequestApprovalSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsAPIService ListCompletedApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.AccessRequestApprovalsAPI.ListCompletedApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsAPIService ListPendingApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.AccessRequestApprovalsAPI.ListPendingApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsAPIService RejectAccessRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		resp, httpRes, err := apiClient.API_V3.AccessRequestApprovalsAPI.RejectAccessRequest(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
