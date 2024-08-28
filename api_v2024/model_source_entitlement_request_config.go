/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the SourceEntitlementRequestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceEntitlementRequestConfig{}

// SourceEntitlementRequestConfig Entitlement Request Configuration
type SourceEntitlementRequestConfig struct {
	AccessRequestConfig *EntitlementAccessRequestConfig `json:"accessRequestConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceEntitlementRequestConfig SourceEntitlementRequestConfig

// NewSourceEntitlementRequestConfig instantiates a new SourceEntitlementRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceEntitlementRequestConfig() *SourceEntitlementRequestConfig {
	this := SourceEntitlementRequestConfig{}
	return &this
}

// NewSourceEntitlementRequestConfigWithDefaults instantiates a new SourceEntitlementRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceEntitlementRequestConfigWithDefaults() *SourceEntitlementRequestConfig {
	this := SourceEntitlementRequestConfig{}
	return &this
}

// GetAccessRequestConfig returns the AccessRequestConfig field value if set, zero value otherwise.
func (o *SourceEntitlementRequestConfig) GetAccessRequestConfig() EntitlementAccessRequestConfig {
	if o == nil || IsNil(o.AccessRequestConfig) {
		var ret EntitlementAccessRequestConfig
		return ret
	}
	return *o.AccessRequestConfig
}

// GetAccessRequestConfigOk returns a tuple with the AccessRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceEntitlementRequestConfig) GetAccessRequestConfigOk() (*EntitlementAccessRequestConfig, bool) {
	if o == nil || IsNil(o.AccessRequestConfig) {
		return nil, false
	}
	return o.AccessRequestConfig, true
}

// HasAccessRequestConfig returns a boolean if a field has been set.
func (o *SourceEntitlementRequestConfig) HasAccessRequestConfig() bool {
	if o != nil && !IsNil(o.AccessRequestConfig) {
		return true
	}

	return false
}

// SetAccessRequestConfig gets a reference to the given EntitlementAccessRequestConfig and assigns it to the AccessRequestConfig field.
func (o *SourceEntitlementRequestConfig) SetAccessRequestConfig(v EntitlementAccessRequestConfig) {
	o.AccessRequestConfig = &v
}

func (o SourceEntitlementRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceEntitlementRequestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessRequestConfig) {
		toSerialize["accessRequestConfig"] = o.AccessRequestConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceEntitlementRequestConfig) UnmarshalJSON(data []byte) (err error) {
	varSourceEntitlementRequestConfig := _SourceEntitlementRequestConfig{}

	err = json.Unmarshal(data, &varSourceEntitlementRequestConfig)

	if err != nil {
		return err
	}

	*o = SourceEntitlementRequestConfig(varSourceEntitlementRequestConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessRequestConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceEntitlementRequestConfig struct {
	value *SourceEntitlementRequestConfig
	isSet bool
}

func (v NullableSourceEntitlementRequestConfig) Get() *SourceEntitlementRequestConfig {
	return v.value
}

func (v *NullableSourceEntitlementRequestConfig) Set(val *SourceEntitlementRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceEntitlementRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceEntitlementRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceEntitlementRequestConfig(val *SourceEntitlementRequestConfig) *NullableSourceEntitlementRequestConfig {
	return &NullableSourceEntitlementRequestConfig{value: val, isSet: true}
}

func (v NullableSourceEntitlementRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceEntitlementRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


