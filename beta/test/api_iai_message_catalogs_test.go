/*
IdentityNow Beta API

Testing IAIMessageCatalogsApiService

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

func Test_beta_IAIMessageCatalogsApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIMessageCatalogsApiService GetMessageCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var catalogId string

		resp, httpRes, err := apiClient.BETA.IAIMessageCatalogsApi.GetMessageCatalogs(context.Background(), catalogId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
