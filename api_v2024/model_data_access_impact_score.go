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

// checks if the DataAccessImpactScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAccessImpactScore{}

// DataAccessImpactScore struct for DataAccessImpactScore
type DataAccessImpactScore struct {
	// Impact Score for this data
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataAccessImpactScore DataAccessImpactScore

// NewDataAccessImpactScore instantiates a new DataAccessImpactScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAccessImpactScore() *DataAccessImpactScore {
	this := DataAccessImpactScore{}
	return &this
}

// NewDataAccessImpactScoreWithDefaults instantiates a new DataAccessImpactScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAccessImpactScoreWithDefaults() *DataAccessImpactScore {
	this := DataAccessImpactScore{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataAccessImpactScore) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccessImpactScore) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataAccessImpactScore) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataAccessImpactScore) SetValue(v string) {
	o.Value = &v
}

func (o DataAccessImpactScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAccessImpactScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataAccessImpactScore) UnmarshalJSON(bytes []byte) (err error) {
	varDataAccessImpactScore := _DataAccessImpactScore{}

	if err = json.Unmarshal(bytes, &varDataAccessImpactScore); err == nil {
	*o = DataAccessImpactScore(varDataAccessImpactScore)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataAccessImpactScore struct {
	value *DataAccessImpactScore
	isSet bool
}

func (v NullableDataAccessImpactScore) Get() *DataAccessImpactScore {
	return v.value
}

func (v *NullableDataAccessImpactScore) Set(val *DataAccessImpactScore) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessImpactScore) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessImpactScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessImpactScore(val *DataAccessImpactScore) *NullableDataAccessImpactScore {
	return &NullableDataAccessImpactScore{value: val, isSet: true}
}

func (v NullableDataAccessImpactScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessImpactScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


