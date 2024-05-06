/*
Identity Security Cloud Beta API

Testing CustomFormsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_beta_CustomFormsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomFormsAPIService CreateFormDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.CreateFormDefinition(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService CreateFormDefinitionDynamicSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.CreateFormDefinitionDynamicSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService CreateFormDefinitionFileRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formDefinitionID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.CreateFormDefinitionFileRequest(context.Background(), formDefinitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService CreateFormInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.CreateFormInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService DeleteFormDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formDefinitionID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.DeleteFormDefinition(context.Background(), formDefinitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService ExportFormDefinitionsByTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.ExportFormDefinitionsByTenant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService GetFileFromS3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formDefinitionID string
		var fileID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.GetFileFromS3(context.Background(), formDefinitionID, fileID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService GetFormDefinitionByKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formDefinitionID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.GetFormDefinitionByKey(context.Background(), formDefinitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService GetFormInstanceByKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formInstanceID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.GetFormInstanceByKey(context.Background(), formInstanceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService GetFormInstanceFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formInstanceID string
		var fileID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.GetFormInstanceFile(context.Background(), formInstanceID, fileID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService ImportFormDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.ImportFormDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService PatchFormDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formDefinitionID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.PatchFormDefinition(context.Background(), formDefinitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService PatchFormInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formInstanceID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.PatchFormInstance(context.Background(), formInstanceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService SearchFormDefinitionsByTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.SearchFormDefinitionsByTenant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService SearchFormElementDataByElementID", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formInstanceID string
		var formElementID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.SearchFormElementDataByElementID(context.Background(), formInstanceID, formElementID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService SearchFormInstancesByTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.SearchFormInstancesByTenant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService SearchPreDefinedSelectOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.SearchPreDefinedSelectOptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomFormsAPIService ShowPreviewDataSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formDefinitionID string

		resp, httpRes, err := apiClient.API_BETA.CustomFormsAPI.ShowPreviewDataSource(context.Background(), formDefinitionID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
