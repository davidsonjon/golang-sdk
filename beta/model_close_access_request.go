/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// CloseAccessRequest Request body payload for close access requests endpoint.
type CloseAccessRequest struct {
	// Access Request IDs for the requests to be closed. Accepts 1-500 Identity Request IDs per request.
	AccessRequestIds []string `json:"accessRequestIds"`
	// Reason for closing the access request. Displayed under Warnings in IdentityNow.
	Message *string `json:"message,omitempty"`
	// The request's provisioning status. Displayed as Stage in IdentityNow.
	ExecutionStatus *string `json:"executionStatus,omitempty"`
	// The request's overall status. Displayed as Status in IdentityNow.
	CompletionStatus *string `json:"completionStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloseAccessRequest CloseAccessRequest

// NewCloseAccessRequest instantiates a new CloseAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseAccessRequest(accessRequestIds []string) *CloseAccessRequest {
	this := CloseAccessRequest{}
	this.AccessRequestIds = accessRequestIds
	var message string = "The IdentityNow Administrator manually closed this request."
	this.Message = &message
	var executionStatus string = "Terminated"
	this.ExecutionStatus = &executionStatus
	var completionStatus string = "Failure"
	this.CompletionStatus = &completionStatus
	return &this
}

// NewCloseAccessRequestWithDefaults instantiates a new CloseAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseAccessRequestWithDefaults() *CloseAccessRequest {
	this := CloseAccessRequest{}
	var message string = "The IdentityNow Administrator manually closed this request."
	this.Message = &message
	var executionStatus string = "Terminated"
	this.ExecutionStatus = &executionStatus
	var completionStatus string = "Failure"
	this.CompletionStatus = &completionStatus
	return &this
}

// GetAccessRequestIds returns the AccessRequestIds field value
func (o *CloseAccessRequest) GetAccessRequestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccessRequestIds
}

// GetAccessRequestIdsOk returns a tuple with the AccessRequestIds field value
// and a boolean to check if the value has been set.
func (o *CloseAccessRequest) GetAccessRequestIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessRequestIds, true
}

// SetAccessRequestIds sets field value
func (o *CloseAccessRequest) SetAccessRequestIds(v []string) {
	o.AccessRequestIds = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CloseAccessRequest) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccessRequest) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CloseAccessRequest) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CloseAccessRequest) SetMessage(v string) {
	o.Message = &v
}

// GetExecutionStatus returns the ExecutionStatus field value if set, zero value otherwise.
func (o *CloseAccessRequest) GetExecutionStatus() string {
	if o == nil || isNil(o.ExecutionStatus) {
		var ret string
		return ret
	}
	return *o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccessRequest) GetExecutionStatusOk() (*string, bool) {
	if o == nil || isNil(o.ExecutionStatus) {
		return nil, false
	}
	return o.ExecutionStatus, true
}

// HasExecutionStatus returns a boolean if a field has been set.
func (o *CloseAccessRequest) HasExecutionStatus() bool {
	if o != nil && !isNil(o.ExecutionStatus) {
		return true
	}

	return false
}

// SetExecutionStatus gets a reference to the given string and assigns it to the ExecutionStatus field.
func (o *CloseAccessRequest) SetExecutionStatus(v string) {
	o.ExecutionStatus = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *CloseAccessRequest) GetCompletionStatus() string {
	if o == nil || isNil(o.CompletionStatus) {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccessRequest) GetCompletionStatusOk() (*string, bool) {
	if o == nil || isNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *CloseAccessRequest) HasCompletionStatus() bool {
	if o != nil && !isNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *CloseAccessRequest) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

func (o CloseAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessRequestIds"] = o.AccessRequestIds
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ExecutionStatus) {
		toSerialize["executionStatus"] = o.ExecutionStatus
	}
	if !isNil(o.CompletionStatus) {
		toSerialize["completionStatus"] = o.CompletionStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloseAccessRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCloseAccessRequest := _CloseAccessRequest{}

	if err = json.Unmarshal(bytes, &varCloseAccessRequest); err == nil {
		*o = CloseAccessRequest(varCloseAccessRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessRequestIds")
		delete(additionalProperties, "message")
		delete(additionalProperties, "executionStatus")
		delete(additionalProperties, "completionStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloseAccessRequest struct {
	value *CloseAccessRequest
	isSet bool
}

func (v NullableCloseAccessRequest) Get() *CloseAccessRequest {
	return v.value
}

func (v *NullableCloseAccessRequest) Set(val *CloseAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseAccessRequest(val *CloseAccessRequest) *NullableCloseAccessRequest {
	return &NullableCloseAccessRequest{value: val, isSet: true}
}

func (v NullableCloseAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


