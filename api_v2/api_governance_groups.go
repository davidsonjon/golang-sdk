/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// GovernanceGroupsAPIService GovernanceGroupsAPI service
type GovernanceGroupsAPIService service

type ApiBulkDeleteWorkGroupsRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	bulkDeleteWorkGroupsRequest *BulkDeleteWorkGroupsRequest
}

// Work group ids to delete
func (r ApiBulkDeleteWorkGroupsRequest) BulkDeleteWorkGroupsRequest(bulkDeleteWorkGroupsRequest BulkDeleteWorkGroupsRequest) ApiBulkDeleteWorkGroupsRequest {
	r.bulkDeleteWorkGroupsRequest = &bulkDeleteWorkGroupsRequest
	return r
}

func (r ApiBulkDeleteWorkGroupsRequest) Execute() (*BulkDeleteWorkGroups200Response, *http.Response, error) {
	return r.ApiService.BulkDeleteWorkGroupsExecute(r)
}

/*
BulkDeleteWorkGroups Bulk delete work groups

This API allows you to bulk-delete work groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkDeleteWorkGroupsRequest
*/
func (a *GovernanceGroupsAPIService) BulkDeleteWorkGroups(ctx context.Context) ApiBulkDeleteWorkGroupsRequest {
	return ApiBulkDeleteWorkGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BulkDeleteWorkGroups200Response
func (a *GovernanceGroupsAPIService) BulkDeleteWorkGroupsExecute(r ApiBulkDeleteWorkGroupsRequest) (*BulkDeleteWorkGroups200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BulkDeleteWorkGroups200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.BulkDeleteWorkGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/bulk-delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bulkDeleteWorkGroupsRequest == nil {
		return localVarReturnValue, nil, reportError("bulkDeleteWorkGroupsRequest is required and must be specified")
	}

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
	// body params
	localVarPostBody = r.bulkDeleteWorkGroupsRequest
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

type ApiCreateWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	createWorkgroupRequest *CreateWorkgroupRequest
}

// Work group to create.
func (r ApiCreateWorkgroupRequest) CreateWorkgroupRequest(createWorkgroupRequest CreateWorkgroupRequest) ApiCreateWorkgroupRequest {
	r.createWorkgroupRequest = &createWorkgroupRequest
	return r
}

func (r ApiCreateWorkgroupRequest) Execute() ([]ListWorkgroups200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateWorkgroupExecute(r)
}

/*
CreateWorkgroup Create Work Group

This API allows you to create a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWorkgroupRequest
*/
func (a *GovernanceGroupsAPIService) CreateWorkgroup(ctx context.Context) ApiCreateWorkgroupRequest {
	return ApiCreateWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ListWorkgroups200ResponseInner
func (a *GovernanceGroupsAPIService) CreateWorkgroupExecute(r ApiCreateWorkgroupRequest) ([]ListWorkgroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListWorkgroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.CreateWorkgroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createWorkgroupRequest == nil {
		return localVarReturnValue, nil, reportError("createWorkgroupRequest is required and must be specified")
	}

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
	// body params
	localVarPostBody = r.createWorkgroupRequest
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

type ApiDeleteWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	workgroupId string
}

func (r ApiDeleteWorkgroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWorkgroupExecute(r)
}

/*
DeleteWorkgroup Delete Work Group By Id

This API deletes a single workgroup based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workgroupId The workgroup ID
 @return ApiDeleteWorkgroupRequest
*/
func (a *GovernanceGroupsAPIService) DeleteWorkgroup(ctx context.Context, workgroupId string) ApiDeleteWorkgroupRequest {
	return ApiDeleteWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
		workgroupId: workgroupId,
	}
}

// Execute executes the request
func (a *GovernanceGroupsAPIService) DeleteWorkgroupExecute(r ApiDeleteWorkgroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.DeleteWorkgroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workgroupId"+"}", url.PathEscape(parameterValueToString(r.workgroupId, "workgroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	workgroupId string
}

func (r ApiGetWorkgroupRequest) Execute() (*ListWorkgroups200ResponseInner, *http.Response, error) {
	return r.ApiService.GetWorkgroupExecute(r)
}

/*
GetWorkgroup Get Work Group By Id

This API returns the details for a single workgroup based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workgroupId The workgroup ID
 @return ApiGetWorkgroupRequest
*/
func (a *GovernanceGroupsAPIService) GetWorkgroup(ctx context.Context, workgroupId string) ApiGetWorkgroupRequest {
	return ApiGetWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
		workgroupId: workgroupId,
	}
}

// Execute executes the request
//  @return ListWorkgroups200ResponseInner
func (a *GovernanceGroupsAPIService) GetWorkgroupExecute(r ApiGetWorkgroupRequest) (*ListWorkgroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListWorkgroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.GetWorkgroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workgroupId"+"}", url.PathEscape(parameterValueToString(r.workgroupId, "workgroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListWorkgroupConnectionsRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	workgroupId string
}

func (r ApiListWorkgroupConnectionsRequest) Execute() ([]ListWorkgroupConnections200ResponseInner, *http.Response, error) {
	return r.ApiService.ListWorkgroupConnectionsExecute(r)
}

/*
ListWorkgroupConnections List Work Group Connections

This API returns the connections of a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workgroupId The workgroup ID
 @return ApiListWorkgroupConnectionsRequest
*/
func (a *GovernanceGroupsAPIService) ListWorkgroupConnections(ctx context.Context, workgroupId string) ApiListWorkgroupConnectionsRequest {
	return ApiListWorkgroupConnectionsRequest{
		ApiService: a,
		ctx: ctx,
		workgroupId: workgroupId,
	}
}

// Execute executes the request
//  @return []ListWorkgroupConnections200ResponseInner
func (a *GovernanceGroupsAPIService) ListWorkgroupConnectionsExecute(r ApiListWorkgroupConnectionsRequest) ([]ListWorkgroupConnections200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListWorkgroupConnections200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.ListWorkgroupConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}/connections"
	localVarPath = strings.Replace(localVarPath, "{"+"workgroupId"+"}", url.PathEscape(parameterValueToString(r.workgroupId, "workgroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListWorkgroupMembersRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	workgroupId string
	limit *int32
	offset *int32
	filters *string
}

// Max number of results to return
func (r ApiListWorkgroupMembersRequest) Limit(limit int32) ApiListWorkgroupMembersRequest {
	r.limit = &limit
	return r
}

// Offset into the full result set. Usually specified with *limit* to paginate through the results.
func (r ApiListWorkgroupMembersRequest) Offset(offset int32) ApiListWorkgroupMembersRequest {
	r.offset = &offset
	return r
}

// Filter results using the following syntax. [{property:name, value: \&quot;Tyler\&quot;, operation: EQ}]
func (r ApiListWorkgroupMembersRequest) Filters(filters string) ApiListWorkgroupMembersRequest {
	r.filters = &filters
	return r
}

func (r ApiListWorkgroupMembersRequest) Execute() ([]ListWorkgroupMembers200ResponseInner, *http.Response, error) {
	return r.ApiService.ListWorkgroupMembersExecute(r)
}

/*
ListWorkgroupMembers List Work Group Members

This API returns the members of a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workgroupId The workgroup ID
 @return ApiListWorkgroupMembersRequest
*/
func (a *GovernanceGroupsAPIService) ListWorkgroupMembers(ctx context.Context, workgroupId string) ApiListWorkgroupMembersRequest {
	return ApiListWorkgroupMembersRequest{
		ApiService: a,
		ctx: ctx,
		workgroupId: workgroupId,
	}
}

// Execute executes the request
//  @return []ListWorkgroupMembers200ResponseInner
func (a *GovernanceGroupsAPIService) ListWorkgroupMembersExecute(r ApiListWorkgroupMembersRequest) ([]ListWorkgroupMembers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListWorkgroupMembers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.ListWorkgroupMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"workgroupId"+"}", url.PathEscape(parameterValueToString(r.workgroupId, "workgroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.filters != nil {
		parameterAddToQuery(localVarQueryParams, "filters", r.filters, "")
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

type ApiListWorkgroupsRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	limit *int32
	offset *int32
	filters *string
}

// Max number of results to return
func (r ApiListWorkgroupsRequest) Limit(limit int32) ApiListWorkgroupsRequest {
	r.limit = &limit
	return r
}

// Offset into the full result set. Usually specified with *limit* to paginate through the results.
func (r ApiListWorkgroupsRequest) Offset(offset int32) ApiListWorkgroupsRequest {
	r.offset = &offset
	return r
}

// Filter results using the following syntax. [{property:name, value: \&quot;Tyler\&quot;, operation: EQ}]
func (r ApiListWorkgroupsRequest) Filters(filters string) ApiListWorkgroupsRequest {
	r.filters = &filters
	return r
}

func (r ApiListWorkgroupsRequest) Execute() ([]ListWorkgroups200ResponseInner, *http.Response, error) {
	return r.ApiService.ListWorkgroupsExecute(r)
}

/*
ListWorkgroups List Work Groups

This API returns a list of work groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWorkgroupsRequest
*/
func (a *GovernanceGroupsAPIService) ListWorkgroups(ctx context.Context) ApiListWorkgroupsRequest {
	return ApiListWorkgroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ListWorkgroups200ResponseInner
func (a *GovernanceGroupsAPIService) ListWorkgroupsExecute(r ApiListWorkgroupsRequest) ([]ListWorkgroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListWorkgroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.ListWorkgroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.filters != nil {
		parameterAddToQuery(localVarQueryParams, "filters", r.filters, "")
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

type ApiModifyWorkgroupMembersRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	workgroupId string
	modifyWorkgroupMembersRequest *ModifyWorkgroupMembersRequest
}

// Add/Remove workgroup member ids.
func (r ApiModifyWorkgroupMembersRequest) ModifyWorkgroupMembersRequest(modifyWorkgroupMembersRequest ModifyWorkgroupMembersRequest) ApiModifyWorkgroupMembersRequest {
	r.modifyWorkgroupMembersRequest = &modifyWorkgroupMembersRequest
	return r
}

func (r ApiModifyWorkgroupMembersRequest) Execute() (*http.Response, error) {
	return r.ApiService.ModifyWorkgroupMembersExecute(r)
}

/*
ModifyWorkgroupMembers Modify Work Group Members

This API allows you to modify the members of a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workgroupId The workgroup ID
 @return ApiModifyWorkgroupMembersRequest
*/
func (a *GovernanceGroupsAPIService) ModifyWorkgroupMembers(ctx context.Context, workgroupId string) ApiModifyWorkgroupMembersRequest {
	return ApiModifyWorkgroupMembersRequest{
		ApiService: a,
		ctx: ctx,
		workgroupId: workgroupId,
	}
}

// Execute executes the request
func (a *GovernanceGroupsAPIService) ModifyWorkgroupMembersExecute(r ApiModifyWorkgroupMembersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.ModifyWorkgroupMembers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"workgroupId"+"}", url.PathEscape(parameterValueToString(r.workgroupId, "workgroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modifyWorkgroupMembersRequest == nil {
		return nil, reportError("modifyWorkgroupMembersRequest is required and must be specified")
	}

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
	// body params
	localVarPostBody = r.modifyWorkgroupMembersRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsAPIService
	workgroupId string
	createWorkgroupRequest *CreateWorkgroupRequest
}

// Work group to modify.
func (r ApiUpdateWorkgroupRequest) CreateWorkgroupRequest(createWorkgroupRequest CreateWorkgroupRequest) ApiUpdateWorkgroupRequest {
	r.createWorkgroupRequest = &createWorkgroupRequest
	return r
}

func (r ApiUpdateWorkgroupRequest) Execute() (*ListWorkgroups200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateWorkgroupExecute(r)
}

/*
UpdateWorkgroup Update Work Group By Id

This API updates and returns the details for a single workgroup based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workgroupId The workgroup ID
 @return ApiUpdateWorkgroupRequest
*/
func (a *GovernanceGroupsAPIService) UpdateWorkgroup(ctx context.Context, workgroupId string) ApiUpdateWorkgroupRequest {
	return ApiUpdateWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
		workgroupId: workgroupId,
	}
}

// Execute executes the request
//  @return ListWorkgroups200ResponseInner
func (a *GovernanceGroupsAPIService) UpdateWorkgroupExecute(r ApiUpdateWorkgroupRequest) (*ListWorkgroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListWorkgroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsAPIService.UpdateWorkgroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workgroupId"+"}", url.PathEscape(parameterValueToString(r.workgroupId, "workgroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createWorkgroupRequest == nil {
		return localVarReturnValue, nil, reportError("createWorkgroupRequest is required and must be specified")
	}

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
	// body params
	localVarPostBody = r.createWorkgroupRequest
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
