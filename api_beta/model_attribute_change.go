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

// checks if the AttributeChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeChange{}

// AttributeChange struct for AttributeChange
type AttributeChange struct {
	// the attribute name
	Name *string `json:"name,omitempty"`
	// the old value of attribute
	PreviousValue *string `json:"previousValue,omitempty"`
	// the new value of attribute
	NewValue *string `json:"newValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttributeChange AttributeChange

// NewAttributeChange instantiates a new AttributeChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeChange() *AttributeChange {
	this := AttributeChange{}
	return &this
}

// NewAttributeChangeWithDefaults instantiates a new AttributeChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeChangeWithDefaults() *AttributeChange {
	this := AttributeChange{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttributeChange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeChange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttributeChange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttributeChange) SetName(v string) {
	o.Name = &v
}

// GetPreviousValue returns the PreviousValue field value if set, zero value otherwise.
func (o *AttributeChange) GetPreviousValue() string {
	if o == nil || IsNil(o.PreviousValue) {
		var ret string
		return ret
	}
	return *o.PreviousValue
}

// GetPreviousValueOk returns a tuple with the PreviousValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeChange) GetPreviousValueOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousValue) {
		return nil, false
	}
	return o.PreviousValue, true
}

// HasPreviousValue returns a boolean if a field has been set.
func (o *AttributeChange) HasPreviousValue() bool {
	if o != nil && !IsNil(o.PreviousValue) {
		return true
	}

	return false
}

// SetPreviousValue gets a reference to the given string and assigns it to the PreviousValue field.
func (o *AttributeChange) SetPreviousValue(v string) {
	o.PreviousValue = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *AttributeChange) GetNewValue() string {
	if o == nil || IsNil(o.NewValue) {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeChange) GetNewValueOk() (*string, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *AttributeChange) HasNewValue() bool {
	if o != nil && !IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *AttributeChange) SetNewValue(v string) {
	o.NewValue = &v
}

func (o AttributeChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PreviousValue) {
		toSerialize["previousValue"] = o.PreviousValue
	}
	if !IsNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttributeChange) UnmarshalJSON(data []byte) (err error) {
	varAttributeChange := _AttributeChange{}

	err = json.Unmarshal(data, &varAttributeChange)

	if err != nil {
		return err
	}

	*o = AttributeChange(varAttributeChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "previousValue")
		delete(additionalProperties, "newValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttributeChange struct {
	value *AttributeChange
	isSet bool
}

func (v NullableAttributeChange) Get() *AttributeChange {
	return v.value
}

func (v *NullableAttributeChange) Set(val *AttributeChange) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeChange) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeChange(val *AttributeChange) *NullableAttributeChange {
	return &NullableAttributeChange{value: val, isSet: true}
}

func (v NullableAttributeChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


