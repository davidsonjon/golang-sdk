# \OrgConfigAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgConfig**](OrgConfigAPI.md#GetOrgConfig) | **Get** /org-config | Get Org configuration settings
[**GetValidTimeZones**](OrgConfigAPI.md#GetValidTimeZones) | **Get** /org-config/valid-time-zones | Get list of time zones
[**PatchOrgConfig**](OrgConfigAPI.md#PatchOrgConfig) | **Patch** /org-config | Patch an Org configuration property



## GetOrgConfig

> OrgConfig GetOrgConfig(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

Get Org configuration settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgConfigAPI.GetOrgConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgConfigAPI.GetOrgConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgConfig`: OrgConfig
	fmt.Fprintf(os.Stdout, "Response from `OrgConfigAPI.GetOrgConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**OrgConfig**](OrgConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidTimeZones

> []string GetValidTimeZones(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

Get list of time zones



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgConfigAPI.GetValidTimeZones(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgConfigAPI.GetValidTimeZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidTimeZones`: []string
	fmt.Fprintf(os.Stdout, "Response from `OrgConfigAPI.GetValidTimeZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidTimeZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

**[]string**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrgConfig

> OrgConfig PatchOrgConfig(ctx).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()

Patch an Org configuration property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of schema attribute update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgConfigAPI.PatchOrgConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgConfigAPI.PatchOrgConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOrgConfig`: OrgConfig
	fmt.Fprintf(os.Stdout, "Response from `OrgConfigAPI.PatchOrgConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of schema attribute update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**OrgConfig**](OrgConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

