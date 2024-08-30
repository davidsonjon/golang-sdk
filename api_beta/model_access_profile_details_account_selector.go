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

// checks if the AccessProfileDetailsAccountSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfileDetailsAccountSelector{}

// AccessProfileDetailsAccountSelector How to select account when there are multiple accounts for the user
type AccessProfileDetailsAccountSelector struct {
	Selectors []Selector `json:"selectors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileDetailsAccountSelector AccessProfileDetailsAccountSelector

// NewAccessProfileDetailsAccountSelector instantiates a new AccessProfileDetailsAccountSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileDetailsAccountSelector() *AccessProfileDetailsAccountSelector {
	this := AccessProfileDetailsAccountSelector{}
	return &this
}

// NewAccessProfileDetailsAccountSelectorWithDefaults instantiates a new AccessProfileDetailsAccountSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileDetailsAccountSelectorWithDefaults() *AccessProfileDetailsAccountSelector {
	this := AccessProfileDetailsAccountSelector{}
	return &this
}

// GetSelectors returns the Selectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDetailsAccountSelector) GetSelectors() []Selector {
	if o == nil {
		var ret []Selector
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDetailsAccountSelector) GetSelectorsOk() ([]Selector, bool) {
	if o == nil || IsNil(o.Selectors) {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *AccessProfileDetailsAccountSelector) HasSelectors() bool {
	if o != nil && !IsNil(o.Selectors) {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []Selector and assigns it to the Selectors field.
func (o *AccessProfileDetailsAccountSelector) SetSelectors(v []Selector) {
	o.Selectors = v
}

func (o AccessProfileDetailsAccountSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfileDetailsAccountSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessProfileDetailsAccountSelector) UnmarshalJSON(data []byte) (err error) {
	varAccessProfileDetailsAccountSelector := _AccessProfileDetailsAccountSelector{}

	err = json.Unmarshal(data, &varAccessProfileDetailsAccountSelector)

	if err != nil {
		return err
	}

	*o = AccessProfileDetailsAccountSelector(varAccessProfileDetailsAccountSelector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selectors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileDetailsAccountSelector struct {
	value *AccessProfileDetailsAccountSelector
	isSet bool
}

func (v NullableAccessProfileDetailsAccountSelector) Get() *AccessProfileDetailsAccountSelector {
	return v.value
}

func (v *NullableAccessProfileDetailsAccountSelector) Set(val *AccessProfileDetailsAccountSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileDetailsAccountSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileDetailsAccountSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileDetailsAccountSelector(val *AccessProfileDetailsAccountSelector) *NullableAccessProfileDetailsAccountSelector {
	return &NullableAccessProfileDetailsAccountSelector{value: val, isSet: true}
}

func (v NullableAccessProfileDetailsAccountSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileDetailsAccountSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


