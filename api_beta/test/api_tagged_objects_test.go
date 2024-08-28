/*
Identity Security Cloud Beta API

Testing TaggedObjectsAPIService

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

func Test_api_beta_TaggedObjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaggedObjectsAPIService DeleteTaggedObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		httpRes, err := apiClient.TaggedObjectsAPI.DeleteTaggedObject(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService DeleteTagsToManyObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TaggedObjectsAPI.DeleteTagsToManyObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService GetTaggedObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.TaggedObjectsAPI.GetTaggedObject(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService ListTaggedObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaggedObjectsAPI.ListTaggedObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService ListTaggedObjectsByType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.TaggedObjectsAPI.ListTaggedObjectsByType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService PutTaggedObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var id string

		resp, httpRes, err := apiClient.TaggedObjectsAPI.PutTaggedObject(context.Background(), type_, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService SetTagToObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TaggedObjectsAPI.SetTagToObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaggedObjectsAPIService SetTagsToManyObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaggedObjectsAPI.SetTagsToManyObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
