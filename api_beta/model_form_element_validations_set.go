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

// checks if the FormElementValidationsSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormElementValidationsSet{}

// FormElementValidationsSet Set of FormElementValidation items.
type FormElementValidationsSet struct {
	ValidationType *string `json:"validationType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormElementValidationsSet FormElementValidationsSet

// NewFormElementValidationsSet instantiates a new FormElementValidationsSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormElementValidationsSet() *FormElementValidationsSet {
	this := FormElementValidationsSet{}
	return &this
}

// NewFormElementValidationsSetWithDefaults instantiates a new FormElementValidationsSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormElementValidationsSetWithDefaults() *FormElementValidationsSet {
	this := FormElementValidationsSet{}
	return &this
}

// GetValidationType returns the ValidationType field value if set, zero value otherwise.
func (o *FormElementValidationsSet) GetValidationType() string {
	if o == nil || IsNil(o.ValidationType) {
		var ret string
		return ret
	}
	return *o.ValidationType
}

// GetValidationTypeOk returns a tuple with the ValidationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElementValidationsSet) GetValidationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationType) {
		return nil, false
	}
	return o.ValidationType, true
}

// HasValidationType returns a boolean if a field has been set.
func (o *FormElementValidationsSet) HasValidationType() bool {
	if o != nil && !IsNil(o.ValidationType) {
		return true
	}

	return false
}

// SetValidationType gets a reference to the given string and assigns it to the ValidationType field.
func (o *FormElementValidationsSet) SetValidationType(v string) {
	o.ValidationType = &v
}

func (o FormElementValidationsSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormElementValidationsSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValidationType) {
		toSerialize["validationType"] = o.ValidationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormElementValidationsSet) UnmarshalJSON(data []byte) (err error) {
	varFormElementValidationsSet := _FormElementValidationsSet{}

	err = json.Unmarshal(data, &varFormElementValidationsSet)

	if err != nil {
		return err
	}

	*o = FormElementValidationsSet(varFormElementValidationsSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "validationType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormElementValidationsSet struct {
	value *FormElementValidationsSet
	isSet bool
}

func (v NullableFormElementValidationsSet) Get() *FormElementValidationsSet {
	return v.value
}

func (v *NullableFormElementValidationsSet) Set(val *FormElementValidationsSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFormElementValidationsSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFormElementValidationsSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormElementValidationsSet(val *FormElementValidationsSet) *NullableFormElementValidationsSet {
	return &NullableFormElementValidationsSet{value: val, isSet: true}
}

func (v NullableFormElementValidationsSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormElementValidationsSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


