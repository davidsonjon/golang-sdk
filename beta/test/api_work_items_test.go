/*
IdentityNow Beta API

Testing WorkItemsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk"
)

func Test_beta_WorkItemsApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkItemsApiService ApproveApprovalItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var approvalItemId string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.ApproveApprovalItem(context.Background(), id, approvalItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService ApproveApprovalItemsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.ApproveApprovalItemsInBulk(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService CompleteWorkItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.CompleteWorkItem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService ForwardWorkItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BETA.WorkItemsApi.ForwardWorkItem(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService GetCompletedWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.GetCompletedWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService GetCountCompletedWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.GetCountCompletedWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService GetCountWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.GetCountWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService GetWorkItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.GetWorkItem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService GetWorkItemsSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.GetWorkItemsSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService ListWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.ListWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService RejectApprovalItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var approvalItemId string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.RejectApprovalItem(context.Background(), id, approvalItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService RejectApprovalItemsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.RejectApprovalItemsInBulk(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsApiService SubmitAccountSelection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.WorkItemsApi.SubmitAccountSelection(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
