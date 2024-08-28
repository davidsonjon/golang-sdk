/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the AccountAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAttributes{}

// AccountAttributes struct for AccountAttributes
type AccountAttributes struct {
	// The schema attribute values for the account
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _AccountAttributes AccountAttributes

// NewAccountAttributes instantiates a new AccountAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAttributes(attributes map[string]interface{}) *AccountAttributes {
	this := AccountAttributes{}
	this.Attributes = attributes
	return &this
}

// NewAccountAttributesWithDefaults instantiates a new AccountAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAttributesWithDefaults() *AccountAttributes {
	this := AccountAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *AccountAttributes) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AccountAttributes) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *AccountAttributes) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o AccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountAttributes := _AccountAttributes{}

	err = json.Unmarshal(data, &varAccountAttributes)

	if err != nil {
		return err
	}

	*o = AccountAttributes(varAccountAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAttributes struct {
	value *AccountAttributes
	isSet bool
}

func (v NullableAccountAttributes) Get() *AccountAttributes {
	return v.value
}

func (v *NullableAccountAttributes) Set(val *AccountAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAttributes(val *AccountAttributes) *NullableAccountAttributes {
	return &NullableAccountAttributes{value: val, isSet: true}
}

func (v NullableAccountAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


