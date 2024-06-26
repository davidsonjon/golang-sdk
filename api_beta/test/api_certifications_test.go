/*
Identity Security Cloud Beta API

Testing CertificationsAPIService

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

func Test_api_beta_CertificationsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificationsAPIService GetIdentityCertificationItemPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var itemId string

		resp, httpRes, err := apiClient.API_BETA.CertificationsAPI.GetIdentityCertificationItemPermissions(context.Background(), certificationId, itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsAPIService GetIdentityCertificationPendingTasks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.CertificationsAPI.GetIdentityCertificationPendingTasks(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsAPIService GetIdentityCertificationTaskStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var taskId string

		resp, httpRes, err := apiClient.API_BETA.CertificationsAPI.GetIdentityCertificationTaskStatus(context.Background(), id, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsAPIService ListCertificationReviewers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.CertificationsAPI.ListCertificationReviewers(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsAPIService SubmitReassignCertsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.CertificationsAPI.SubmitReassignCertsAsync(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
