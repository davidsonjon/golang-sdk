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

// checks if the ExceptionCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionCriteria{}

// ExceptionCriteria struct for ExceptionCriteria
type ExceptionCriteria struct {
	// List of exception criteria. There is a min of 1 and max of 50 items in the list.
	CriteriaList []ExceptionCriteriaCriteriaListInner `json:"criteriaList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExceptionCriteria ExceptionCriteria

// NewExceptionCriteria instantiates a new ExceptionCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionCriteria() *ExceptionCriteria {
	this := ExceptionCriteria{}
	return &this
}

// NewExceptionCriteriaWithDefaults instantiates a new ExceptionCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionCriteriaWithDefaults() *ExceptionCriteria {
	this := ExceptionCriteria{}
	return &this
}

// GetCriteriaList returns the CriteriaList field value if set, zero value otherwise.
func (o *ExceptionCriteria) GetCriteriaList() []ExceptionCriteriaCriteriaListInner {
	if o == nil || IsNil(o.CriteriaList) {
		var ret []ExceptionCriteriaCriteriaListInner
		return ret
	}
	return o.CriteriaList
}

// GetCriteriaListOk returns a tuple with the CriteriaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteria) GetCriteriaListOk() ([]ExceptionCriteriaCriteriaListInner, bool) {
	if o == nil || IsNil(o.CriteriaList) {
		return nil, false
	}
	return o.CriteriaList, true
}

// HasCriteriaList returns a boolean if a field has been set.
func (o *ExceptionCriteria) HasCriteriaList() bool {
	if o != nil && !IsNil(o.CriteriaList) {
		return true
	}

	return false
}

// SetCriteriaList gets a reference to the given []ExceptionCriteriaCriteriaListInner and assigns it to the CriteriaList field.
func (o *ExceptionCriteria) SetCriteriaList(v []ExceptionCriteriaCriteriaListInner) {
	o.CriteriaList = v
}

func (o ExceptionCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CriteriaList) {
		toSerialize["criteriaList"] = o.CriteriaList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExceptionCriteria) UnmarshalJSON(data []byte) (err error) {
	varExceptionCriteria := _ExceptionCriteria{}

	err = json.Unmarshal(data, &varExceptionCriteria)

	if err != nil {
		return err
	}

	*o = ExceptionCriteria(varExceptionCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "criteriaList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExceptionCriteria struct {
	value *ExceptionCriteria
	isSet bool
}

func (v NullableExceptionCriteria) Get() *ExceptionCriteria {
	return v.value
}

func (v *NullableExceptionCriteria) Set(val *ExceptionCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionCriteria(val *ExceptionCriteria) *NullableExceptionCriteria {
	return &NullableExceptionCriteria{value: val, isSet: true}
}

func (v NullableExceptionCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


