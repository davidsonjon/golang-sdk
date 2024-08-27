/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the LifecycleStateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleStateDto{}

// LifecycleStateDto struct for LifecycleStateDto
type LifecycleStateDto struct {
	// The name of the lifecycle state
	StateName string `json:"stateName"`
	// Whether the lifecycle state has been manually or automatically set
	ManuallyUpdated bool `json:"manuallyUpdated"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleStateDto LifecycleStateDto

// NewLifecycleStateDto instantiates a new LifecycleStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleStateDto(stateName string, manuallyUpdated bool) *LifecycleStateDto {
	this := LifecycleStateDto{}
	this.StateName = stateName
	this.ManuallyUpdated = manuallyUpdated
	return &this
}

// NewLifecycleStateDtoWithDefaults instantiates a new LifecycleStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleStateDtoWithDefaults() *LifecycleStateDto {
	this := LifecycleStateDto{}
	return &this
}

// GetStateName returns the StateName field value
func (o *LifecycleStateDto) GetStateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value
// and a boolean to check if the value has been set.
func (o *LifecycleStateDto) GetStateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateName, true
}

// SetStateName sets field value
func (o *LifecycleStateDto) SetStateName(v string) {
	o.StateName = v
}

// GetManuallyUpdated returns the ManuallyUpdated field value
func (o *LifecycleStateDto) GetManuallyUpdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ManuallyUpdated
}

// GetManuallyUpdatedOk returns a tuple with the ManuallyUpdated field value
// and a boolean to check if the value has been set.
func (o *LifecycleStateDto) GetManuallyUpdatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManuallyUpdated, true
}

// SetManuallyUpdated sets field value
func (o *LifecycleStateDto) SetManuallyUpdated(v bool) {
	o.ManuallyUpdated = v
}

func (o LifecycleStateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleStateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stateName"] = o.StateName
	toSerialize["manuallyUpdated"] = o.ManuallyUpdated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleStateDto) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stateName",
		"manuallyUpdated",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLifecycleStateDto := _LifecycleStateDto{}

	if err = json.Unmarshal(bytes, &varLifecycleStateDto); err == nil {
	*o = LifecycleStateDto(varLifecycleStateDto)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "stateName")
		delete(additionalProperties, "manuallyUpdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleStateDto struct {
	value *LifecycleStateDto
	isSet bool
}

func (v NullableLifecycleStateDto) Get() *LifecycleStateDto {
	return v.value
}

func (v *NullableLifecycleStateDto) Set(val *LifecycleStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleStateDto(val *LifecycleStateDto) *NullableLifecycleStateDto {
	return &NullableLifecycleStateDto{value: val, isSet: true}
}

func (v NullableLifecycleStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

