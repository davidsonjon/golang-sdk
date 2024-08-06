/*
Identity Security Cloud V2024 API

Testing WorkItemsAPIService

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

func Test_api_v2024_WorkItemsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkItemsAPIService ApproveApprovalItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var approvalItemId string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.ApproveApprovalItem(context.Background(), id, approvalItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService ApproveApprovalItemsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.ApproveApprovalItemsInBulk(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService CompleteWorkItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.CompleteWorkItem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService GetCompletedWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.GetCompletedWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService GetCountCompletedWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.GetCountCompletedWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService GetCountWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.GetCountWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService GetWorkItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.GetWorkItem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService GetWorkItemsSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.GetWorkItemsSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService ListWorkItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.ListWorkItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService RejectApprovalItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var approvalItemId string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.RejectApprovalItem(context.Background(), id, approvalItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService RejectApprovalItemsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.RejectApprovalItemsInBulk(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService SendWorkItemForward", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_V2024.WorkItemsAPI.SendWorkItemForward(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkItemsAPIService SubmitAccountSelection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V2024.WorkItemsAPI.SubmitAccountSelection(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
