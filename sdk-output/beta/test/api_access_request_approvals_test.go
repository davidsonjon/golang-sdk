/*
IdentityNow Beta API

Testing AccessRequestApprovalsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_beta_AccessRequestApprovalsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessRequestApprovalsApiService ApprovalSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccessRequestApprovalsApi.ApprovalSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsApiService ApproveRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		resp, httpRes, err := apiClient.AccessRequestApprovalsApi.ApproveRequest(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsApiService ForwardRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		resp, httpRes, err := apiClient.AccessRequestApprovalsApi.ForwardRequest(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsApiService ListCompletedApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccessRequestApprovalsApi.ListCompletedApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsApiService ListPendingApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccessRequestApprovalsApi.ListPendingApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRequestApprovalsApiService RejectRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		resp, httpRes, err := apiClient.AccessRequestApprovalsApi.RejectRequest(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
