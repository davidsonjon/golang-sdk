/*
IdentityNow Beta API

Testing CertificationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_beta_CertificationsApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificationsApiService GetIdentityCertificationItemPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var itemId string

		resp, httpRes, err := apiClient.BETA.CertificationsApi.GetIdentityCertificationItemPermissions(context.Background(), certificationId, itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsApiService GetIdentityCertificationPendingTasks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.CertificationsApi.GetIdentityCertificationPendingTasks(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsApiService GetIdentityCertificationTaskStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var taskId string

		resp, httpRes, err := apiClient.BETA.CertificationsApi.GetIdentityCertificationTaskStatus(context.Background(), id, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsApiService ListCertificationReviewers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.CertificationsApi.ListCertificationReviewers(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationsApiService SubmitReassignCertsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.CertificationsApi.SubmitReassignCertsAsync(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
