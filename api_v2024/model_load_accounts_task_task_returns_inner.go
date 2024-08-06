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

// checks if the LoadAccountsTaskTaskReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadAccountsTaskTaskReturnsInner{}

// LoadAccountsTaskTaskReturnsInner struct for LoadAccountsTaskTaskReturnsInner
type LoadAccountsTaskTaskReturnsInner struct {
	// The display label of the return value
	DisplayLabel *string `json:"displayLabel,omitempty"`
	// The attribute name of the return value
	AttributeName *string `json:"attributeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadAccountsTaskTaskReturnsInner LoadAccountsTaskTaskReturnsInner

// NewLoadAccountsTaskTaskReturnsInner instantiates a new LoadAccountsTaskTaskReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadAccountsTaskTaskReturnsInner() *LoadAccountsTaskTaskReturnsInner {
	this := LoadAccountsTaskTaskReturnsInner{}
	return &this
}

// NewLoadAccountsTaskTaskReturnsInnerWithDefaults instantiates a new LoadAccountsTaskTaskReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadAccountsTaskTaskReturnsInnerWithDefaults() *LoadAccountsTaskTaskReturnsInner {
	this := LoadAccountsTaskTaskReturnsInner{}
	return &this
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskReturnsInner) GetDisplayLabel() string {
	if o == nil || isNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskReturnsInner) GetDisplayLabelOk() (*string, bool) {
	if o == nil || isNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskReturnsInner) HasDisplayLabel() bool {
	if o != nil && !isNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *LoadAccountsTaskTaskReturnsInner) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskReturnsInner) GetAttributeName() string {
	if o == nil || isNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskReturnsInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskReturnsInner) HasAttributeName() bool {
	if o != nil && !isNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *LoadAccountsTaskTaskReturnsInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

func (o LoadAccountsTaskTaskReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadAccountsTaskTaskReturnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if !isNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadAccountsTaskTaskReturnsInner) UnmarshalJSON(bytes []byte) (err error) {
	varLoadAccountsTaskTaskReturnsInner := _LoadAccountsTaskTaskReturnsInner{}

	if err = json.Unmarshal(bytes, &varLoadAccountsTaskTaskReturnsInner); err == nil {
	*o = LoadAccountsTaskTaskReturnsInner(varLoadAccountsTaskTaskReturnsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayLabel")
		delete(additionalProperties, "attributeName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadAccountsTaskTaskReturnsInner struct {
	value *LoadAccountsTaskTaskReturnsInner
	isSet bool
}

func (v NullableLoadAccountsTaskTaskReturnsInner) Get() *LoadAccountsTaskTaskReturnsInner {
	return v.value
}

func (v *NullableLoadAccountsTaskTaskReturnsInner) Set(val *LoadAccountsTaskTaskReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadAccountsTaskTaskReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadAccountsTaskTaskReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadAccountsTaskTaskReturnsInner(val *LoadAccountsTaskTaskReturnsInner) *NullableLoadAccountsTaskTaskReturnsInner {
	return &NullableLoadAccountsTaskTaskReturnsInner{value: val, isSet: true}
}

func (v NullableLoadAccountsTaskTaskReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadAccountsTaskTaskReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


