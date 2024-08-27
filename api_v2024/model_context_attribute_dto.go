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

// checks if the ContextAttributeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextAttributeDto{}

// ContextAttributeDto struct for ContextAttributeDto
type ContextAttributeDto struct {
	// The name of the attribute
	Attribute *string `json:"attribute,omitempty"`
	Value *ContextAttributeDtoValue `json:"value,omitempty"`
	// True if the attribute was derived.
	Derived *bool `json:"derived,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContextAttributeDto ContextAttributeDto

// NewContextAttributeDto instantiates a new ContextAttributeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextAttributeDto() *ContextAttributeDto {
	this := ContextAttributeDto{}
	var derived bool = false
	this.Derived = &derived
	return &this
}

// NewContextAttributeDtoWithDefaults instantiates a new ContextAttributeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextAttributeDtoWithDefaults() *ContextAttributeDto {
	this := ContextAttributeDto{}
	var derived bool = false
	this.Derived = &derived
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *ContextAttributeDto) GetAttribute() string {
	if o == nil || isNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextAttributeDto) GetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *ContextAttributeDto) HasAttribute() bool {
	if o != nil && !isNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *ContextAttributeDto) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContextAttributeDto) GetValue() ContextAttributeDtoValue {
	if o == nil || isNil(o.Value) {
		var ret ContextAttributeDtoValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextAttributeDto) GetValueOk() (*ContextAttributeDtoValue, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContextAttributeDto) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ContextAttributeDtoValue and assigns it to the Value field.
func (o *ContextAttributeDto) SetValue(v ContextAttributeDtoValue) {
	o.Value = &v
}

// GetDerived returns the Derived field value if set, zero value otherwise.
func (o *ContextAttributeDto) GetDerived() bool {
	if o == nil || isNil(o.Derived) {
		var ret bool
		return ret
	}
	return *o.Derived
}

// GetDerivedOk returns a tuple with the Derived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextAttributeDto) GetDerivedOk() (*bool, bool) {
	if o == nil || isNil(o.Derived) {
		return nil, false
	}
	return o.Derived, true
}

// HasDerived returns a boolean if a field has been set.
func (o *ContextAttributeDto) HasDerived() bool {
	if o != nil && !isNil(o.Derived) {
		return true
	}

	return false
}

// SetDerived gets a reference to the given bool and assigns it to the Derived field.
func (o *ContextAttributeDto) SetDerived(v bool) {
	o.Derived = &v
}

func (o ContextAttributeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextAttributeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Derived) {
		toSerialize["derived"] = o.Derived
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContextAttributeDto) UnmarshalJSON(bytes []byte) (err error) {
	varContextAttributeDto := _ContextAttributeDto{}

	if err = json.Unmarshal(bytes, &varContextAttributeDto); err == nil {
	*o = ContextAttributeDto(varContextAttributeDto)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "derived")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContextAttributeDto struct {
	value *ContextAttributeDto
	isSet bool
}

func (v NullableContextAttributeDto) Get() *ContextAttributeDto {
	return v.value
}

func (v *NullableContextAttributeDto) Set(val *ContextAttributeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableContextAttributeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableContextAttributeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextAttributeDto(val *ContextAttributeDto) *NullableContextAttributeDto {
	return &NullableContextAttributeDto{value: val, isSet: true}
}

func (v NullableContextAttributeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextAttributeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

