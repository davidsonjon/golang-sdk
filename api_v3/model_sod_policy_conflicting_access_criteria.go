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

// checks if the SodPolicyConflictingAccessCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodPolicyConflictingAccessCriteria{}

// SodPolicyConflictingAccessCriteria struct for SodPolicyConflictingAccessCriteria
type SodPolicyConflictingAccessCriteria struct {
	LeftCriteria *AccessCriteria `json:"leftCriteria,omitempty"`
	RightCriteria *AccessCriteria `json:"rightCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodPolicyConflictingAccessCriteria SodPolicyConflictingAccessCriteria

// NewSodPolicyConflictingAccessCriteria instantiates a new SodPolicyConflictingAccessCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodPolicyConflictingAccessCriteria() *SodPolicyConflictingAccessCriteria {
	this := SodPolicyConflictingAccessCriteria{}
	return &this
}

// NewSodPolicyConflictingAccessCriteriaWithDefaults instantiates a new SodPolicyConflictingAccessCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodPolicyConflictingAccessCriteriaWithDefaults() *SodPolicyConflictingAccessCriteria {
	this := SodPolicyConflictingAccessCriteria{}
	return &this
}

// GetLeftCriteria returns the LeftCriteria field value if set, zero value otherwise.
func (o *SodPolicyConflictingAccessCriteria) GetLeftCriteria() AccessCriteria {
	if o == nil || isNil(o.LeftCriteria) {
		var ret AccessCriteria
		return ret
	}
	return *o.LeftCriteria
}

// GetLeftCriteriaOk returns a tuple with the LeftCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicyConflictingAccessCriteria) GetLeftCriteriaOk() (*AccessCriteria, bool) {
	if o == nil || isNil(o.LeftCriteria) {
		return nil, false
	}
	return o.LeftCriteria, true
}

// HasLeftCriteria returns a boolean if a field has been set.
func (o *SodPolicyConflictingAccessCriteria) HasLeftCriteria() bool {
	if o != nil && !isNil(o.LeftCriteria) {
		return true
	}

	return false
}

// SetLeftCriteria gets a reference to the given AccessCriteria and assigns it to the LeftCriteria field.
func (o *SodPolicyConflictingAccessCriteria) SetLeftCriteria(v AccessCriteria) {
	o.LeftCriteria = &v
}

// GetRightCriteria returns the RightCriteria field value if set, zero value otherwise.
func (o *SodPolicyConflictingAccessCriteria) GetRightCriteria() AccessCriteria {
	if o == nil || isNil(o.RightCriteria) {
		var ret AccessCriteria
		return ret
	}
	return *o.RightCriteria
}

// GetRightCriteriaOk returns a tuple with the RightCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicyConflictingAccessCriteria) GetRightCriteriaOk() (*AccessCriteria, bool) {
	if o == nil || isNil(o.RightCriteria) {
		return nil, false
	}
	return o.RightCriteria, true
}

// HasRightCriteria returns a boolean if a field has been set.
func (o *SodPolicyConflictingAccessCriteria) HasRightCriteria() bool {
	if o != nil && !isNil(o.RightCriteria) {
		return true
	}

	return false
}

// SetRightCriteria gets a reference to the given AccessCriteria and assigns it to the RightCriteria field.
func (o *SodPolicyConflictingAccessCriteria) SetRightCriteria(v AccessCriteria) {
	o.RightCriteria = &v
}

func (o SodPolicyConflictingAccessCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodPolicyConflictingAccessCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LeftCriteria) {
		toSerialize["leftCriteria"] = o.LeftCriteria
	}
	if !isNil(o.RightCriteria) {
		toSerialize["rightCriteria"] = o.RightCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodPolicyConflictingAccessCriteria) UnmarshalJSON(bytes []byte) (err error) {
	varSodPolicyConflictingAccessCriteria := _SodPolicyConflictingAccessCriteria{}

	if err = json.Unmarshal(bytes, &varSodPolicyConflictingAccessCriteria); err == nil {
	*o = SodPolicyConflictingAccessCriteria(varSodPolicyConflictingAccessCriteria)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "leftCriteria")
		delete(additionalProperties, "rightCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodPolicyConflictingAccessCriteria struct {
	value *SodPolicyConflictingAccessCriteria
	isSet bool
}

func (v NullableSodPolicyConflictingAccessCriteria) Get() *SodPolicyConflictingAccessCriteria {
	return v.value
}

func (v *NullableSodPolicyConflictingAccessCriteria) Set(val *SodPolicyConflictingAccessCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSodPolicyConflictingAccessCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSodPolicyConflictingAccessCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodPolicyConflictingAccessCriteria(val *SodPolicyConflictingAccessCriteria) *NullableSodPolicyConflictingAccessCriteria {
	return &NullableSodPolicyConflictingAccessCriteria{value: val, isSet: true}
}

func (v NullableSodPolicyConflictingAccessCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodPolicyConflictingAccessCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


