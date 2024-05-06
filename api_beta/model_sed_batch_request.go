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

// checks if the SedBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SedBatchRequest{}

// SedBatchRequest Sed Batch Request
type SedBatchRequest struct {
	// list of entitlement ids
	Entitlements []string `json:"entitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SedBatchRequest SedBatchRequest

// NewSedBatchRequest instantiates a new SedBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSedBatchRequest() *SedBatchRequest {
	this := SedBatchRequest{}
	return &this
}

// NewSedBatchRequestWithDefaults instantiates a new SedBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSedBatchRequestWithDefaults() *SedBatchRequest {
	this := SedBatchRequest{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *SedBatchRequest) GetEntitlements() []string {
	if o == nil || isNil(o.Entitlements) {
		var ret []string
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SedBatchRequest) GetEntitlementsOk() ([]string, bool) {
	if o == nil || isNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *SedBatchRequest) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []string and assigns it to the Entitlements field.
func (o *SedBatchRequest) SetEntitlements(v []string) {
	o.Entitlements = v
}

func (o SedBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SedBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SedBatchRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSedBatchRequest := _SedBatchRequest{}

	if err = json.Unmarshal(bytes, &varSedBatchRequest); err == nil {
	*o = SedBatchRequest(varSedBatchRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSedBatchRequest struct {
	value *SedBatchRequest
	isSet bool
}

func (v NullableSedBatchRequest) Get() *SedBatchRequest {
	return v.value
}

func (v *NullableSedBatchRequest) Set(val *SedBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSedBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSedBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSedBatchRequest(val *SedBatchRequest) *NullableSedBatchRequest {
	return &NullableSedBatchRequest{value: val, isSet: true}
}

func (v NullableSedBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSedBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


