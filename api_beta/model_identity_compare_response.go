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

// checks if the IdentityCompareResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityCompareResponse{}

// IdentityCompareResponse struct for IdentityCompareResponse
type IdentityCompareResponse struct {
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on completion of the violation check.
	AccessItemDiff map[string]map[string]interface{} `json:"accessItemDiff,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityCompareResponse IdentityCompareResponse

// NewIdentityCompareResponse instantiates a new IdentityCompareResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCompareResponse() *IdentityCompareResponse {
	this := IdentityCompareResponse{}
	return &this
}

// NewIdentityCompareResponseWithDefaults instantiates a new IdentityCompareResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCompareResponseWithDefaults() *IdentityCompareResponse {
	this := IdentityCompareResponse{}
	return &this
}

// GetAccessItemDiff returns the AccessItemDiff field value if set, zero value otherwise.
func (o *IdentityCompareResponse) GetAccessItemDiff() map[string]map[string]interface{} {
	if o == nil || IsNil(o.AccessItemDiff) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.AccessItemDiff
}

// GetAccessItemDiffOk returns a tuple with the AccessItemDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCompareResponse) GetAccessItemDiffOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AccessItemDiff) {
		return map[string]map[string]interface{}{}, false
	}
	return o.AccessItemDiff, true
}

// HasAccessItemDiff returns a boolean if a field has been set.
func (o *IdentityCompareResponse) HasAccessItemDiff() bool {
	if o != nil && !IsNil(o.AccessItemDiff) {
		return true
	}

	return false
}

// SetAccessItemDiff gets a reference to the given map[string]map[string]interface{} and assigns it to the AccessItemDiff field.
func (o *IdentityCompareResponse) SetAccessItemDiff(v map[string]map[string]interface{}) {
	o.AccessItemDiff = v
}

func (o IdentityCompareResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityCompareResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessItemDiff) {
		toSerialize["accessItemDiff"] = o.AccessItemDiff
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityCompareResponse) UnmarshalJSON(data []byte) (err error) {
	varIdentityCompareResponse := _IdentityCompareResponse{}

	err = json.Unmarshal(data, &varIdentityCompareResponse)

	if err != nil {
		return err
	}

	*o = IdentityCompareResponse(varIdentityCompareResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessItemDiff")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityCompareResponse struct {
	value *IdentityCompareResponse
	isSet bool
}

func (v NullableIdentityCompareResponse) Get() *IdentityCompareResponse {
	return v.value
}

func (v *NullableIdentityCompareResponse) Set(val *IdentityCompareResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCompareResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCompareResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCompareResponse(val *IdentityCompareResponse) *NullableIdentityCompareResponse {
	return &NullableIdentityCompareResponse{value: val, isSet: true}
}

func (v NullableIdentityCompareResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCompareResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


