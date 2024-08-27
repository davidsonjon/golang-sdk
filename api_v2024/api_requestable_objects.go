/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// RequestableObjectsAPIService RequestableObjectsAPI service
type RequestableObjectsAPIService service

type ApiListRequestableObjectsRequest struct {
	ctx context.Context
	ApiService *RequestableObjectsAPIService
	identityId *string
	types *[]RequestableObjectType
	term *string
	statuses *[]RequestableObjectRequestStatus
	limit *int32
	offset *int32
	count *bool
	filters *string
	sorters *string
}

// If present, the value returns only requestable objects for the specified identity.  * Admin users can call this with any identity ID value.  * Non-admin users can only specify *me* or pass their own identity ID value.  * If absent, returns a list of all requestable objects for the tenant. Only admin users can make such a call. In this case, the available, pending, assigned accesses will not be annotated in the result.
func (r ApiListRequestableObjectsRequest) IdentityId(identityId string) ApiListRequestableObjectsRequest {
	r.identityId = &identityId
	return r
}

// Filters the results to the specified type/types, where each type is one of ROLE or ACCESS_PROFILE. If absent, all types are returned. Support for additional types may be added in the future without notice.
func (r ApiListRequestableObjectsRequest) Types(types []RequestableObjectType) ApiListRequestableObjectsRequest {
	r.types = &types
	return r
}

// It allows searching requestable access items with a partial match on the name or description. If term is provided, then the *filter* query parameter will be ignored.
func (r ApiListRequestableObjectsRequest) Term(term string) ApiListRequestableObjectsRequest {
	r.term = &term
	return r
}

// Filters the result to the specified status/statuses, where each status is one of AVAILABLE, ASSIGNED, or PENDING. It is an error to specify this parameter without also specifying an *identity-id* parameter. Additional statuses may be added in the future without notice.
func (r ApiListRequestableObjectsRequest) Statuses(statuses []RequestableObjectRequestStatus) ApiListRequestableObjectsRequest {
	r.statuses = &statuses
	return r
}

// Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
func (r ApiListRequestableObjectsRequest) Limit(limit int32) ApiListRequestableObjectsRequest {
	r.limit = &limit
	return r
}

// Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
func (r ApiListRequestableObjectsRequest) Offset(offset int32) ApiListRequestableObjectsRequest {
	r.offset = &offset
	return r
}

// If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
func (r ApiListRequestableObjectsRequest) Count(count bool) ApiListRequestableObjectsRequest {
	r.count = &count
	return r
}

// Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw* 
func (r ApiListRequestableObjectsRequest) Filters(filters string) ApiListRequestableObjectsRequest {
	r.filters = &filters
	return r
}

// Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** 
func (r ApiListRequestableObjectsRequest) Sorters(sorters string) ApiListRequestableObjectsRequest {
	r.sorters = &sorters
	return r
}

func (r ApiListRequestableObjectsRequest) Execute() ([]RequestableObject, *http.Response, error) {
	return r.ApiService.ListRequestableObjectsExecute(r)
}

/*
ListRequestableObjects Requestable Objects List

This endpoint returns a list of acccess items that that can be requested through the Access Request endpoints. Access items are marked with AVAILABLE, PENDING or ASSIGNED with respect to the identity provided using *identity-id* query param.
Any authenticated token can call this endpoint to see their requestable access items. A token with ORG_ADMIN authority is required to call this endpoint to return a list of all of the requestable access items for the org or for another identity.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListRequestableObjectsRequest
*/
func (a *RequestableObjectsAPIService) ListRequestableObjects(ctx context.Context) ApiListRequestableObjectsRequest {
	return ApiListRequestableObjectsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RequestableObject
func (a *RequestableObjectsAPIService) ListRequestableObjectsExecute(r ApiListRequestableObjectsRequest) ([]RequestableObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RequestableObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestableObjectsAPIService.ListRequestableObjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requestable-objects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.identityId != nil {
		parameterAddToQuery(localVarQueryParams, "identity-id", r.identityId, "")
	}
	if r.types != nil {
		parameterAddToQuery(localVarQueryParams, "types", r.types, "csv")
	}
	if r.term != nil {
		parameterAddToQuery(localVarQueryParams, "term", r.term, "")
	}
	if r.statuses != nil {
		parameterAddToQuery(localVarQueryParams, "statuses", r.statuses, "csv")
	}
	if r.limit != nil {
		parameterAddToQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 250
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.count != nil {
		parameterAddToQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue bool = false
		r.count = &defaultValue
	}
	if r.filters != nil {
		parameterAddToQuery(localVarQueryParams, "filters", r.filters, "")
	}
	if r.sorters != nil {
		parameterAddToQuery(localVarQueryParams, "sorters", r.sorters, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ListAccessProfiles401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ListAccessProfiles429Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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