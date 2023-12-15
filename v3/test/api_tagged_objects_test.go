/*
IdentityNow V3 API

Testing TaggedObjectsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk"
)

func Test_v3_TaggedObjectsApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaggedObjectsApiService DeleteTaggedObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		httpRes, err := apiClient.V3.TaggedObjectsApi.DeleteTaggedObject(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService DeleteTagsToManyObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.V3.TaggedObjectsApi.DeleteTagsToManyObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService GetTaggedObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.V3.TaggedObjectsApi.GetTaggedObject(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService ListTaggedObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.TaggedObjectsApi.ListTaggedObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService ListTaggedObjectsByType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.V3.TaggedObjectsApi.ListTaggedObjectsByType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService PutTaggedObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.V3.TaggedObjectsApi.PutTaggedObject(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService SetTagToObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.V3.TaggedObjectsApi.SetTagToObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsApiService SetTagsToManyObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.TaggedObjectsApi.SetTagsToManyObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
