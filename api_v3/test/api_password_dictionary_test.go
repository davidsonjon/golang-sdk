/*
Identity Security Cloud V3 API

Testing PasswordDictionaryAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"testing"

	openapiclient "github.com/sailpoint-oss/davidsonjon/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_v3_PasswordDictionaryAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PasswordDictionaryAPIService GetPasswordDictionary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.PasswordDictionaryAPI.GetPasswordDictionary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordDictionaryAPIService PutPasswordDictionary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.API_V3.PasswordDictionaryAPI.PutPasswordDictionary(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
