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

// checks if the IdentitiesReportArguments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitiesReportArguments{}

// IdentitiesReportArguments Arguments for Identities report (IDENTITIES)
type IdentitiesReportArguments struct {
	// Boolean FLAG to specify if only correlated identities should be used in report processing
	CorrelatedOnly *bool `json:"correlatedOnly,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitiesReportArguments IdentitiesReportArguments

// NewIdentitiesReportArguments instantiates a new IdentitiesReportArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitiesReportArguments() *IdentitiesReportArguments {
	this := IdentitiesReportArguments{}
	var correlatedOnly bool = false
	this.CorrelatedOnly = &correlatedOnly
	return &this
}

// NewIdentitiesReportArgumentsWithDefaults instantiates a new IdentitiesReportArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitiesReportArgumentsWithDefaults() *IdentitiesReportArguments {
	this := IdentitiesReportArguments{}
	var correlatedOnly bool = false
	this.CorrelatedOnly = &correlatedOnly
	return &this
}

// GetCorrelatedOnly returns the CorrelatedOnly field value if set, zero value otherwise.
func (o *IdentitiesReportArguments) GetCorrelatedOnly() bool {
	if o == nil || isNil(o.CorrelatedOnly) {
		var ret bool
		return ret
	}
	return *o.CorrelatedOnly
}

// GetCorrelatedOnlyOk returns a tuple with the CorrelatedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitiesReportArguments) GetCorrelatedOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.CorrelatedOnly) {
		return nil, false
	}
	return o.CorrelatedOnly, true
}

// HasCorrelatedOnly returns a boolean if a field has been set.
func (o *IdentitiesReportArguments) HasCorrelatedOnly() bool {
	if o != nil && !isNil(o.CorrelatedOnly) {
		return true
	}

	return false
}

// SetCorrelatedOnly gets a reference to the given bool and assigns it to the CorrelatedOnly field.
func (o *IdentitiesReportArguments) SetCorrelatedOnly(v bool) {
	o.CorrelatedOnly = &v
}

func (o IdentitiesReportArguments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitiesReportArguments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CorrelatedOnly) {
		toSerialize["correlatedOnly"] = o.CorrelatedOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitiesReportArguments) UnmarshalJSON(bytes []byte) (err error) {
	varIdentitiesReportArguments := _IdentitiesReportArguments{}

	if err = json.Unmarshal(bytes, &varIdentitiesReportArguments); err == nil {
	*o = IdentitiesReportArguments(varIdentitiesReportArguments)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "correlatedOnly")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitiesReportArguments struct {
	value *IdentitiesReportArguments
	isSet bool
}

func (v NullableIdentitiesReportArguments) Get() *IdentitiesReportArguments {
	return v.value
}

func (v *NullableIdentitiesReportArguments) Set(val *IdentitiesReportArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitiesReportArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitiesReportArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitiesReportArguments(val *IdentitiesReportArguments) *NullableIdentitiesReportArguments {
	return &NullableIdentitiesReportArguments{value: val, isSet: true}
}

func (v NullableIdentitiesReportArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitiesReportArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

