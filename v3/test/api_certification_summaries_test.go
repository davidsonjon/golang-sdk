/*
IdentityNow V3 API

Testing CertificationSummariesAPIService

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

func Test_v3_CertificationSummariesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificationSummariesAPIService GetIdentityAccessSummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var type_ string

		resp, httpRes, err := apiClient.V3.CertificationSummariesAPI.GetIdentityAccessSummaries(context.Background(), id, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationSummariesAPIService GetIdentityDecisionSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.V3.CertificationSummariesAPI.GetIdentityDecisionSummary(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationSummariesAPIService GetIdentitySummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.V3.CertificationSummariesAPI.GetIdentitySummaries(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationSummariesAPIService GetIdentitySummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var identitySummaryId string

		resp, httpRes, err := apiClient.V3.CertificationSummariesAPI.GetIdentitySummary(context.Background(), id, identitySummaryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
