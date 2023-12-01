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

// checks if the EntitlementRequestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementRequestConfig{}

// EntitlementRequestConfig struct for EntitlementRequestConfig
type EntitlementRequestConfig struct {
	AccessRequestConfig *EntitlementAccessRequestConfig `json:"accessRequestConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementRequestConfig EntitlementRequestConfig

// NewEntitlementRequestConfig instantiates a new EntitlementRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementRequestConfig() *EntitlementRequestConfig {
	this := EntitlementRequestConfig{}
	return &this
}

// NewEntitlementRequestConfigWithDefaults instantiates a new EntitlementRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementRequestConfigWithDefaults() *EntitlementRequestConfig {
	this := EntitlementRequestConfig{}
	return &this
}

// GetAccessRequestConfig returns the AccessRequestConfig field value if set, zero value otherwise.
func (o *EntitlementRequestConfig) GetAccessRequestConfig() EntitlementAccessRequestConfig {
	if o == nil || isNil(o.AccessRequestConfig) {
		var ret EntitlementAccessRequestConfig
		return ret
	}
	return *o.AccessRequestConfig
}

// GetAccessRequestConfigOk returns a tuple with the AccessRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig) GetAccessRequestConfigOk() (*EntitlementAccessRequestConfig, bool) {
	if o == nil || isNil(o.AccessRequestConfig) {
		return nil, false
	}
	return o.AccessRequestConfig, true
}

// HasAccessRequestConfig returns a boolean if a field has been set.
func (o *EntitlementRequestConfig) HasAccessRequestConfig() bool {
	if o != nil && !isNil(o.AccessRequestConfig) {
		return true
	}

	return false
}

// SetAccessRequestConfig gets a reference to the given EntitlementAccessRequestConfig and assigns it to the AccessRequestConfig field.
func (o *EntitlementRequestConfig) SetAccessRequestConfig(v EntitlementAccessRequestConfig) {
	o.AccessRequestConfig = &v
}

func (o EntitlementRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementRequestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessRequestConfig) {
		toSerialize["accessRequestConfig"] = o.AccessRequestConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementRequestConfig) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementRequestConfig := _EntitlementRequestConfig{}

	if err = json.Unmarshal(bytes, &varEntitlementRequestConfig); err == nil {
	*o = EntitlementRequestConfig(varEntitlementRequestConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessRequestConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementRequestConfig struct {
	value *EntitlementRequestConfig
	isSet bool
}

func (v NullableEntitlementRequestConfig) Get() *EntitlementRequestConfig {
	return v.value
}

func (v *NullableEntitlementRequestConfig) Set(val *EntitlementRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementRequestConfig(val *EntitlementRequestConfig) *NullableEntitlementRequestConfig {
	return &NullableEntitlementRequestConfig{value: val, isSet: true}
}

func (v NullableEntitlementRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


