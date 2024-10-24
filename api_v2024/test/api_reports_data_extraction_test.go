/*
Identity Security Cloud V2024 API

Testing ReportsDataExtractionAPIService

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

func Test_api_v2024_ReportsDataExtractionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReportsDataExtractionAPIService CancelReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ReportsDataExtractionAPI.CancelReport(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsDataExtractionAPIService GetReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskResultId string

		resp, httpRes, err := apiClient.ReportsDataExtractionAPI.GetReport(context.Background(), taskResultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsDataExtractionAPIService GetReportResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskResultId string

		resp, httpRes, err := apiClient.ReportsDataExtractionAPI.GetReportResult(context.Background(), taskResultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsDataExtractionAPIService StartReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReportsDataExtractionAPI.StartReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
