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

// checks if the DataAccessPoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAccessPoliciesInner{}

// DataAccessPoliciesInner struct for DataAccessPoliciesInner
type DataAccessPoliciesInner struct {
	// Value of the policy
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataAccessPoliciesInner DataAccessPoliciesInner

// NewDataAccessPoliciesInner instantiates a new DataAccessPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAccessPoliciesInner() *DataAccessPoliciesInner {
	this := DataAccessPoliciesInner{}
	return &this
}

// NewDataAccessPoliciesInnerWithDefaults instantiates a new DataAccessPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAccessPoliciesInnerWithDefaults() *DataAccessPoliciesInner {
	this := DataAccessPoliciesInner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataAccessPoliciesInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccessPoliciesInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataAccessPoliciesInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataAccessPoliciesInner) SetValue(v string) {
	o.Value = &v
}

func (o DataAccessPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAccessPoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataAccessPoliciesInner) UnmarshalJSON(bytes []byte) (err error) {
	varDataAccessPoliciesInner := _DataAccessPoliciesInner{}

	if err = json.Unmarshal(bytes, &varDataAccessPoliciesInner); err == nil {
	*o = DataAccessPoliciesInner(varDataAccessPoliciesInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataAccessPoliciesInner struct {
	value *DataAccessPoliciesInner
	isSet bool
}

func (v NullableDataAccessPoliciesInner) Get() *DataAccessPoliciesInner {
	return v.value
}

func (v *NullableDataAccessPoliciesInner) Set(val *DataAccessPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessPoliciesInner(val *DataAccessPoliciesInner) *NullableDataAccessPoliciesInner {
	return &NullableDataAccessPoliciesInner{value: val, isSet: true}
}

func (v NullableDataAccessPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


