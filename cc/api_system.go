/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// SystemApiService SystemApi service
type SystemApiService service

type ApiRefreshIdentitiesRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	contentType *string
	refreshIdentitiesRequest *RefreshIdentitiesRequest
}

func (r ApiRefreshIdentitiesRequest) ContentType(contentType string) ApiRefreshIdentitiesRequest {
	r.contentType = &contentType
	return r
}

func (r ApiRefreshIdentitiesRequest) RefreshIdentitiesRequest(refreshIdentitiesRequest RefreshIdentitiesRequest) ApiRefreshIdentitiesRequest {
	r.refreshIdentitiesRequest = &refreshIdentitiesRequest
	return r
}

func (r ApiRefreshIdentitiesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RefreshIdentitiesExecute(r)
}

/*
RefreshIdentities Refresh Identities

This kicks off an identity refresh for a specified set of identity attributes.  This can be a long running process.  IdentityNow has pre-scheduled versions of this task at set intervals and events already, so only run this when directed by SailPoint.

_Note: If the identities specified by the filter do not exist, a full identity refresh will be run.  Use with caution._

Refresh Arguments:

| Key                   | Description                                        |
|-----------------------|----------------------------------------------------|
| correlateEntitlements | Analyzes entitlements, access profiles, and roles. |
| promoteAttributes     | Calculates identity attributes.                    |
| refreshManagerStatus  | Calculates manager correlation and manager status. |
| synchronizeAttributes | Performs attribute sync provisioning.              |
| pruneIdentities       | Removes any identities which don't have accounts.  |
| provision             | Provisions any assigned roles or access profiles.  |

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRefreshIdentitiesRequest
*/
func (a *SystemApiService) RefreshIdentities(ctx context.Context) ApiRefreshIdentitiesRequest {
	return ApiRefreshIdentitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SystemApiService) RefreshIdentitiesExecute(r ApiRefreshIdentitiesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.RefreshIdentities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cc/api/system/refreshIdentities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.contentType != nil {
		parameterAddToQuery(localVarQueryParams, "Content-Type", r.contentType, "")
	}
	// body params
	localVarPostBody = r.refreshIdentitiesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
