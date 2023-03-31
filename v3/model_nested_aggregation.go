/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the NestedAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedAggregation{}

// NestedAggregation The nested aggregation object.
type NestedAggregation struct {
	// The name of the nested aggregate to be included in the result.
	Name string `json:"name"`
	// The type of the nested object.
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _NestedAggregation NestedAggregation

// NewNestedAggregation instantiates a new NestedAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedAggregation(name string, type_ string) *NestedAggregation {
	this := NestedAggregation{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewNestedAggregationWithDefaults instantiates a new NestedAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedAggregationWithDefaults() *NestedAggregation {
	this := NestedAggregation{}
	return &this
}

// GetName returns the Name field value
func (o *NestedAggregation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedAggregation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedAggregation) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *NestedAggregation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NestedAggregation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NestedAggregation) SetType(v string) {
	o.Type = v
}

func (o NestedAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedAggregation) UnmarshalJSON(bytes []byte) (err error) {
	varNestedAggregation := _NestedAggregation{}

	if err = json.Unmarshal(bytes, &varNestedAggregation); err == nil {
		*o = NestedAggregation(varNestedAggregation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedAggregation struct {
	value *NestedAggregation
	isSet bool
}

func (v NullableNestedAggregation) Get() *NestedAggregation {
	return v.value
}

func (v *NullableNestedAggregation) Set(val *NestedAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedAggregation(val *NestedAggregation) *NullableNestedAggregation {
	return &NullableNestedAggregation{value: val, isSet: true}
}

func (v NullableNestedAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


