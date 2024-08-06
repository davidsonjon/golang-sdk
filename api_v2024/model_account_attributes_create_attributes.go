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

// checks if the AccountAttributesCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAttributesCreateAttributes{}

// AccountAttributesCreateAttributes The schema attribute values for the account
type AccountAttributesCreateAttributes struct {
	// Target source to create an account
	SourceId string `json:"sourceId"`
	AdditionalProperties map[string]interface{}
}

type _AccountAttributesCreateAttributes AccountAttributesCreateAttributes

// NewAccountAttributesCreateAttributes instantiates a new AccountAttributesCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAttributesCreateAttributes(sourceId string) *AccountAttributesCreateAttributes {
	this := AccountAttributesCreateAttributes{}
	this.SourceId = sourceId
	return &this
}

// NewAccountAttributesCreateAttributesWithDefaults instantiates a new AccountAttributesCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAttributesCreateAttributesWithDefaults() *AccountAttributesCreateAttributes {
	this := AccountAttributesCreateAttributes{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *AccountAttributesCreateAttributes) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *AccountAttributesCreateAttributes) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *AccountAttributesCreateAttributes) SetSourceId(v string) {
	o.SourceId = v
}

func (o AccountAttributesCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAttributesCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountAttributesCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountAttributesCreateAttributes := _AccountAttributesCreateAttributes{}

	if err = json.Unmarshal(bytes, &varAccountAttributesCreateAttributes); err == nil {
	*o = AccountAttributesCreateAttributes(varAccountAttributesCreateAttributes)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sourceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAttributesCreateAttributes struct {
	value *AccountAttributesCreateAttributes
	isSet bool
}

func (v NullableAccountAttributesCreateAttributes) Get() *AccountAttributesCreateAttributes {
	return v.value
}

func (v *NullableAccountAttributesCreateAttributes) Set(val *AccountAttributesCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAttributesCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAttributesCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAttributesCreateAttributes(val *AccountAttributesCreateAttributes) *NullableAccountAttributesCreateAttributes {
	return &NullableAccountAttributesCreateAttributes{value: val, isSet: true}
}

func (v NullableAccountAttributesCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAttributesCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


