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

// checks if the SodViolationContextConflictingAccessCriteriaLeftCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationContextConflictingAccessCriteriaLeftCriteria{}

// SodViolationContextConflictingAccessCriteriaLeftCriteria struct for SodViolationContextConflictingAccessCriteriaLeftCriteria
type SodViolationContextConflictingAccessCriteriaLeftCriteria struct {
	CriteriaList []SodExemptCriteria `json:"criteriaList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContextConflictingAccessCriteriaLeftCriteria SodViolationContextConflictingAccessCriteriaLeftCriteria

// NewSodViolationContextConflictingAccessCriteriaLeftCriteria instantiates a new SodViolationContextConflictingAccessCriteriaLeftCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContextConflictingAccessCriteriaLeftCriteria() *SodViolationContextConflictingAccessCriteriaLeftCriteria {
	this := SodViolationContextConflictingAccessCriteriaLeftCriteria{}
	return &this
}

// NewSodViolationContextConflictingAccessCriteriaLeftCriteriaWithDefaults instantiates a new SodViolationContextConflictingAccessCriteriaLeftCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContextConflictingAccessCriteriaLeftCriteriaWithDefaults() *SodViolationContextConflictingAccessCriteriaLeftCriteria {
	this := SodViolationContextConflictingAccessCriteriaLeftCriteria{}
	return &this
}

// GetCriteriaList returns the CriteriaList field value if set, zero value otherwise.
func (o *SodViolationContextConflictingAccessCriteriaLeftCriteria) GetCriteriaList() []SodExemptCriteria {
	if o == nil || IsNil(o.CriteriaList) {
		var ret []SodExemptCriteria
		return ret
	}
	return o.CriteriaList
}

// GetCriteriaListOk returns a tuple with the CriteriaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextConflictingAccessCriteriaLeftCriteria) GetCriteriaListOk() ([]SodExemptCriteria, bool) {
	if o == nil || IsNil(o.CriteriaList) {
		return nil, false
	}
	return o.CriteriaList, true
}

// HasCriteriaList returns a boolean if a field has been set.
func (o *SodViolationContextConflictingAccessCriteriaLeftCriteria) HasCriteriaList() bool {
	if o != nil && !IsNil(o.CriteriaList) {
		return true
	}

	return false
}

// SetCriteriaList gets a reference to the given []SodExemptCriteria and assigns it to the CriteriaList field.
func (o *SodViolationContextConflictingAccessCriteriaLeftCriteria) SetCriteriaList(v []SodExemptCriteria) {
	o.CriteriaList = v
}

func (o SodViolationContextConflictingAccessCriteriaLeftCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationContextConflictingAccessCriteriaLeftCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CriteriaList) {
		toSerialize["criteriaList"] = o.CriteriaList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationContextConflictingAccessCriteriaLeftCriteria) UnmarshalJSON(data []byte) (err error) {
	varSodViolationContextConflictingAccessCriteriaLeftCriteria := _SodViolationContextConflictingAccessCriteriaLeftCriteria{}

	err = json.Unmarshal(data, &varSodViolationContextConflictingAccessCriteriaLeftCriteria)

	if err != nil {
		return err
	}

	*o = SodViolationContextConflictingAccessCriteriaLeftCriteria(varSodViolationContextConflictingAccessCriteriaLeftCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "criteriaList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContextConflictingAccessCriteriaLeftCriteria struct {
	value *SodViolationContextConflictingAccessCriteriaLeftCriteria
	isSet bool
}

func (v NullableSodViolationContextConflictingAccessCriteriaLeftCriteria) Get() *SodViolationContextConflictingAccessCriteriaLeftCriteria {
	return v.value
}

func (v *NullableSodViolationContextConflictingAccessCriteriaLeftCriteria) Set(val *SodViolationContextConflictingAccessCriteriaLeftCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContextConflictingAccessCriteriaLeftCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContextConflictingAccessCriteriaLeftCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContextConflictingAccessCriteriaLeftCriteria(val *SodViolationContextConflictingAccessCriteriaLeftCriteria) *NullableSodViolationContextConflictingAccessCriteriaLeftCriteria {
	return &NullableSodViolationContextConflictingAccessCriteriaLeftCriteria{value: val, isSet: true}
}

func (v NullableSodViolationContextConflictingAccessCriteriaLeftCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContextConflictingAccessCriteriaLeftCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


