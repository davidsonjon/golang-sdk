/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_cc

import (
	"encoding/json"
	"os"
)

// checks if the LoadEntitlementsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadEntitlementsRequest{}

// LoadEntitlementsRequest struct for LoadEntitlementsRequest
type LoadEntitlementsRequest struct {
	File **os.File `json:"file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadEntitlementsRequest LoadEntitlementsRequest

// NewLoadEntitlementsRequest instantiates a new LoadEntitlementsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadEntitlementsRequest() *LoadEntitlementsRequest {
	this := LoadEntitlementsRequest{}
	return &this
}

// NewLoadEntitlementsRequestWithDefaults instantiates a new LoadEntitlementsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadEntitlementsRequestWithDefaults() *LoadEntitlementsRequest {
	this := LoadEntitlementsRequest{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *LoadEntitlementsRequest) GetFile() *os.File {
	if o == nil || isNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementsRequest) GetFileOk() (**os.File, bool) {
	if o == nil || isNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *LoadEntitlementsRequest) HasFile() bool {
	if o != nil && !isNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *LoadEntitlementsRequest) SetFile(v *os.File) {
	o.File = &v
}

func (o LoadEntitlementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadEntitlementsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.File) {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadEntitlementsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varLoadEntitlementsRequest := _LoadEntitlementsRequest{}

	if err = json.Unmarshal(bytes, &varLoadEntitlementsRequest); err == nil {
	*o = LoadEntitlementsRequest(varLoadEntitlementsRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadEntitlementsRequest struct {
	value *LoadEntitlementsRequest
	isSet bool
}

func (v NullableLoadEntitlementsRequest) Get() *LoadEntitlementsRequest {
	return v.value
}

func (v *NullableLoadEntitlementsRequest) Set(val *LoadEntitlementsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadEntitlementsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadEntitlementsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadEntitlementsRequest(val *LoadEntitlementsRequest) *NullableLoadEntitlementsRequest {
	return &NullableLoadEntitlementsRequest{value: val, isSet: true}
}

func (v NullableLoadEntitlementsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadEntitlementsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


