/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the Range type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Range{}

// Range The range of values to be filtered.
type Range struct {
	Lower *Bound `json:"lower,omitempty"`
	Upper *Bound `json:"upper,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Range Range

// NewRange instantiates a new Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRange() *Range {
	this := Range{}
	return &this
}

// NewRangeWithDefaults instantiates a new Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeWithDefaults() *Range {
	this := Range{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *Range) GetLower() Bound {
	if o == nil || IsNil(o.Lower) {
		var ret Bound
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetLowerOk() (*Bound, bool) {
	if o == nil || IsNil(o.Lower) {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *Range) HasLower() bool {
	if o != nil && !IsNil(o.Lower) {
		return true
	}

	return false
}

// SetLower gets a reference to the given Bound and assigns it to the Lower field.
func (o *Range) SetLower(v Bound) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *Range) GetUpper() Bound {
	if o == nil || IsNil(o.Upper) {
		var ret Bound
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetUpperOk() (*Bound, bool) {
	if o == nil || IsNil(o.Upper) {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *Range) HasUpper() bool {
	if o != nil && !IsNil(o.Upper) {
		return true
	}

	return false
}

// SetUpper gets a reference to the given Bound and assigns it to the Upper field.
func (o *Range) SetUpper(v Bound) {
	o.Upper = &v
}

func (o Range) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Range) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lower) {
		toSerialize["lower"] = o.Lower
	}
	if !IsNil(o.Upper) {
		toSerialize["upper"] = o.Upper
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Range) UnmarshalJSON(data []byte) (err error) {
	varRange := _Range{}

	err = json.Unmarshal(data, &varRange)

	if err != nil {
		return err
	}

	*o = Range(varRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lower")
		delete(additionalProperties, "upper")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRange struct {
	value *Range
	isSet bool
}

func (v NullableRange) Get() *Range {
	return v.value
}

func (v *NullableRange) Set(val *Range) {
	v.value = val
	v.isSet = true
}

func (v NullableRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRange(val *Range) *NullableRange {
	return &NullableRange{value: val, isSet: true}
}

func (v NullableRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


