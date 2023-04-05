/*
IdentityNow V3 API

Testing CertificationSummariesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_v3_CertificationSummariesApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificationSummariesApiService GetIdentityAccessSummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var type_ string

		resp, httpRes, err := apiClient.V3.CertificationSummariesApi.GetIdentityAccessSummaries(context.Background(), id, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationSummariesApiService GetIdentityDecisionSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.V3.CertificationSummariesApi.GetIdentityDecisionSummary(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationSummariesApiService GetIdentitySummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.V3.CertificationSummariesApi.GetIdentitySummaries(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationSummariesApiService GetIdentitySummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var identitySummaryId string

		resp, httpRes, err := apiClient.V3.CertificationSummariesApi.GetIdentitySummary(context.Background(), id, identitySummaryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
