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

// checks if the ConflictingAccessCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConflictingAccessCriteria{}

// ConflictingAccessCriteria struct for ConflictingAccessCriteria
type ConflictingAccessCriteria struct {
	LeftCriteria *AccessCriteria `json:"leftCriteria,omitempty"`
	RightCriteria *AccessCriteria `json:"rightCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConflictingAccessCriteria ConflictingAccessCriteria

// NewConflictingAccessCriteria instantiates a new ConflictingAccessCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictingAccessCriteria() *ConflictingAccessCriteria {
	this := ConflictingAccessCriteria{}
	return &this
}

// NewConflictingAccessCriteriaWithDefaults instantiates a new ConflictingAccessCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictingAccessCriteriaWithDefaults() *ConflictingAccessCriteria {
	this := ConflictingAccessCriteria{}
	return &this
}

// GetLeftCriteria returns the LeftCriteria field value if set, zero value otherwise.
func (o *ConflictingAccessCriteria) GetLeftCriteria() AccessCriteria {
	if o == nil || IsNil(o.LeftCriteria) {
		var ret AccessCriteria
		return ret
	}
	return *o.LeftCriteria
}

// GetLeftCriteriaOk returns a tuple with the LeftCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictingAccessCriteria) GetLeftCriteriaOk() (*AccessCriteria, bool) {
	if o == nil || IsNil(o.LeftCriteria) {
		return nil, false
	}
	return o.LeftCriteria, true
}

// HasLeftCriteria returns a boolean if a field has been set.
func (o *ConflictingAccessCriteria) HasLeftCriteria() bool {
	if o != nil && !IsNil(o.LeftCriteria) {
		return true
	}

	return false
}

// SetLeftCriteria gets a reference to the given AccessCriteria and assigns it to the LeftCriteria field.
func (o *ConflictingAccessCriteria) SetLeftCriteria(v AccessCriteria) {
	o.LeftCriteria = &v
}

// GetRightCriteria returns the RightCriteria field value if set, zero value otherwise.
func (o *ConflictingAccessCriteria) GetRightCriteria() AccessCriteria {
	if o == nil || IsNil(o.RightCriteria) {
		var ret AccessCriteria
		return ret
	}
	return *o.RightCriteria
}

// GetRightCriteriaOk returns a tuple with the RightCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictingAccessCriteria) GetRightCriteriaOk() (*AccessCriteria, bool) {
	if o == nil || IsNil(o.RightCriteria) {
		return nil, false
	}
	return o.RightCriteria, true
}

// HasRightCriteria returns a boolean if a field has been set.
func (o *ConflictingAccessCriteria) HasRightCriteria() bool {
	if o != nil && !IsNil(o.RightCriteria) {
		return true
	}

	return false
}

// SetRightCriteria gets a reference to the given AccessCriteria and assigns it to the RightCriteria field.
func (o *ConflictingAccessCriteria) SetRightCriteria(v AccessCriteria) {
	o.RightCriteria = &v
}

func (o ConflictingAccessCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConflictingAccessCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LeftCriteria) {
		toSerialize["leftCriteria"] = o.LeftCriteria
	}
	if !IsNil(o.RightCriteria) {
		toSerialize["rightCriteria"] = o.RightCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConflictingAccessCriteria) UnmarshalJSON(data []byte) (err error) {
	varConflictingAccessCriteria := _ConflictingAccessCriteria{}

	err = json.Unmarshal(data, &varConflictingAccessCriteria)

	if err != nil {
		return err
	}

	*o = ConflictingAccessCriteria(varConflictingAccessCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftCriteria")
		delete(additionalProperties, "rightCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConflictingAccessCriteria struct {
	value *ConflictingAccessCriteria
	isSet bool
}

func (v NullableConflictingAccessCriteria) Get() *ConflictingAccessCriteria {
	return v.value
}

func (v *NullableConflictingAccessCriteria) Set(val *ConflictingAccessCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictingAccessCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictingAccessCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictingAccessCriteria(val *ConflictingAccessCriteria) *NullableConflictingAccessCriteria {
	return &NullableConflictingAccessCriteria{value: val, isSet: true}
}

func (v NullableConflictingAccessCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictingAccessCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


