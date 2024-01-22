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

// checks if the PatchServiceDeskIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchServiceDeskIntegrationRequest{}

// PatchServiceDeskIntegrationRequest A JSONPatch document as defined by [RFC 6902 - JSON Patch](https://tools.ietf.org/html/rfc6902).  Only `replace` operations are accepted by this endpoint.
type PatchServiceDeskIntegrationRequest struct {
	// Operations to be applied
	Operations []JsonPatchOperation `json:"operations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchServiceDeskIntegrationRequest PatchServiceDeskIntegrationRequest

// NewPatchServiceDeskIntegrationRequest instantiates a new PatchServiceDeskIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchServiceDeskIntegrationRequest() *PatchServiceDeskIntegrationRequest {
	this := PatchServiceDeskIntegrationRequest{}
	return &this
}

// NewPatchServiceDeskIntegrationRequestWithDefaults instantiates a new PatchServiceDeskIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchServiceDeskIntegrationRequestWithDefaults() *PatchServiceDeskIntegrationRequest {
	this := PatchServiceDeskIntegrationRequest{}
	return &this
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *PatchServiceDeskIntegrationRequest) GetOperations() []JsonPatchOperation {
	if o == nil || isNil(o.Operations) {
		var ret []JsonPatchOperation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchServiceDeskIntegrationRequest) GetOperationsOk() ([]JsonPatchOperation, bool) {
	if o == nil || isNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *PatchServiceDeskIntegrationRequest) HasOperations() bool {
	if o != nil && !isNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []JsonPatchOperation and assigns it to the Operations field.
func (o *PatchServiceDeskIntegrationRequest) SetOperations(v []JsonPatchOperation) {
	o.Operations = v
}

func (o PatchServiceDeskIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchServiceDeskIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchServiceDeskIntegrationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchServiceDeskIntegrationRequest := _PatchServiceDeskIntegrationRequest{}

	if err = json.Unmarshal(bytes, &varPatchServiceDeskIntegrationRequest); err == nil {
	*o = PatchServiceDeskIntegrationRequest(varPatchServiceDeskIntegrationRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchServiceDeskIntegrationRequest struct {
	value *PatchServiceDeskIntegrationRequest
	isSet bool
}

func (v NullablePatchServiceDeskIntegrationRequest) Get() *PatchServiceDeskIntegrationRequest {
	return v.value
}

func (v *NullablePatchServiceDeskIntegrationRequest) Set(val *PatchServiceDeskIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchServiceDeskIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchServiceDeskIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchServiceDeskIntegrationRequest(val *PatchServiceDeskIntegrationRequest) *NullablePatchServiceDeskIntegrationRequest {
	return &NullablePatchServiceDeskIntegrationRequest{value: val, isSet: true}
}

func (v NullablePatchServiceDeskIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchServiceDeskIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


