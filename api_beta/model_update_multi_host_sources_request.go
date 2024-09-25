/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the UpdateMultiHostSourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMultiHostSourcesRequest{}

// UpdateMultiHostSourcesRequest A JSONPatch document as defined by [RFC 6902 - JSON Patch](https://tools.ietf.org/html/rfc6902).  Only `replace` operations are accepted by this endpoint.
type UpdateMultiHostSourcesRequest struct {
	// Operations to be applied.
	Operations []JsonPatchOperation `json:"operations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateMultiHostSourcesRequest UpdateMultiHostSourcesRequest

// NewUpdateMultiHostSourcesRequest instantiates a new UpdateMultiHostSourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMultiHostSourcesRequest() *UpdateMultiHostSourcesRequest {
	this := UpdateMultiHostSourcesRequest{}
	return &this
}

// NewUpdateMultiHostSourcesRequestWithDefaults instantiates a new UpdateMultiHostSourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMultiHostSourcesRequestWithDefaults() *UpdateMultiHostSourcesRequest {
	this := UpdateMultiHostSourcesRequest{}
	return &this
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *UpdateMultiHostSourcesRequest) GetOperations() []JsonPatchOperation {
	if o == nil || IsNil(o.Operations) {
		var ret []JsonPatchOperation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMultiHostSourcesRequest) GetOperationsOk() ([]JsonPatchOperation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *UpdateMultiHostSourcesRequest) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []JsonPatchOperation and assigns it to the Operations field.
func (o *UpdateMultiHostSourcesRequest) SetOperations(v []JsonPatchOperation) {
	o.Operations = v
}

func (o UpdateMultiHostSourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMultiHostSourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateMultiHostSourcesRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateMultiHostSourcesRequest := _UpdateMultiHostSourcesRequest{}

	err = json.Unmarshal(data, &varUpdateMultiHostSourcesRequest)

	if err != nil {
		return err
	}

	*o = UpdateMultiHostSourcesRequest(varUpdateMultiHostSourcesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateMultiHostSourcesRequest struct {
	value *UpdateMultiHostSourcesRequest
	isSet bool
}

func (v NullableUpdateMultiHostSourcesRequest) Get() *UpdateMultiHostSourcesRequest {
	return v.value
}

func (v *NullableUpdateMultiHostSourcesRequest) Set(val *UpdateMultiHostSourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMultiHostSourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMultiHostSourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMultiHostSourcesRequest(val *UpdateMultiHostSourcesRequest) *NullableUpdateMultiHostSourcesRequest {
	return &NullableUpdateMultiHostSourcesRequest{value: val, isSet: true}
}

func (v NullableUpdateMultiHostSourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMultiHostSourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


