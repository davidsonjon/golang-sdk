/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2

import (
	"encoding/json"
	"fmt"
)

// checks if the GetOrgSettings200ResponseApprovalConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrgSettings200ResponseApprovalConfig{}

// GetOrgSettings200ResponseApprovalConfig struct for GetOrgSettings200ResponseApprovalConfig
type GetOrgSettings200ResponseApprovalConfig struct {
	DaysTillEscalation int32 `json:"daysTillEscalation"`
	DaysBetweenReminders int32 `json:"daysBetweenReminders"`
	MaxReminders int32 `json:"maxReminders"`
	FallbackApprover string `json:"fallbackApprover"`
	AdditionalProperties map[string]interface{}
}

type _GetOrgSettings200ResponseApprovalConfig GetOrgSettings200ResponseApprovalConfig

// NewGetOrgSettings200ResponseApprovalConfig instantiates a new GetOrgSettings200ResponseApprovalConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrgSettings200ResponseApprovalConfig(daysTillEscalation int32, daysBetweenReminders int32, maxReminders int32, fallbackApprover string) *GetOrgSettings200ResponseApprovalConfig {
	this := GetOrgSettings200ResponseApprovalConfig{}
	this.DaysTillEscalation = daysTillEscalation
	this.DaysBetweenReminders = daysBetweenReminders
	this.MaxReminders = maxReminders
	this.FallbackApprover = fallbackApprover
	return &this
}

// NewGetOrgSettings200ResponseApprovalConfigWithDefaults instantiates a new GetOrgSettings200ResponseApprovalConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrgSettings200ResponseApprovalConfigWithDefaults() *GetOrgSettings200ResponseApprovalConfig {
	this := GetOrgSettings200ResponseApprovalConfig{}
	return &this
}

// GetDaysTillEscalation returns the DaysTillEscalation field value
func (o *GetOrgSettings200ResponseApprovalConfig) GetDaysTillEscalation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysTillEscalation
}

// GetDaysTillEscalationOk returns a tuple with the DaysTillEscalation field value
// and a boolean to check if the value has been set.
func (o *GetOrgSettings200ResponseApprovalConfig) GetDaysTillEscalationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysTillEscalation, true
}

// SetDaysTillEscalation sets field value
func (o *GetOrgSettings200ResponseApprovalConfig) SetDaysTillEscalation(v int32) {
	o.DaysTillEscalation = v
}

// GetDaysBetweenReminders returns the DaysBetweenReminders field value
func (o *GetOrgSettings200ResponseApprovalConfig) GetDaysBetweenReminders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysBetweenReminders
}

// GetDaysBetweenRemindersOk returns a tuple with the DaysBetweenReminders field value
// and a boolean to check if the value has been set.
func (o *GetOrgSettings200ResponseApprovalConfig) GetDaysBetweenRemindersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysBetweenReminders, true
}

// SetDaysBetweenReminders sets field value
func (o *GetOrgSettings200ResponseApprovalConfig) SetDaysBetweenReminders(v int32) {
	o.DaysBetweenReminders = v
}

// GetMaxReminders returns the MaxReminders field value
func (o *GetOrgSettings200ResponseApprovalConfig) GetMaxReminders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxReminders
}

// GetMaxRemindersOk returns a tuple with the MaxReminders field value
// and a boolean to check if the value has been set.
func (o *GetOrgSettings200ResponseApprovalConfig) GetMaxRemindersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxReminders, true
}

// SetMaxReminders sets field value
func (o *GetOrgSettings200ResponseApprovalConfig) SetMaxReminders(v int32) {
	o.MaxReminders = v
}

// GetFallbackApprover returns the FallbackApprover field value
func (o *GetOrgSettings200ResponseApprovalConfig) GetFallbackApprover() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FallbackApprover
}

// GetFallbackApproverOk returns a tuple with the FallbackApprover field value
// and a boolean to check if the value has been set.
func (o *GetOrgSettings200ResponseApprovalConfig) GetFallbackApproverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FallbackApprover, true
}

// SetFallbackApprover sets field value
func (o *GetOrgSettings200ResponseApprovalConfig) SetFallbackApprover(v string) {
	o.FallbackApprover = v
}

func (o GetOrgSettings200ResponseApprovalConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrgSettings200ResponseApprovalConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["daysTillEscalation"] = o.DaysTillEscalation
	toSerialize["daysBetweenReminders"] = o.DaysBetweenReminders
	toSerialize["maxReminders"] = o.MaxReminders
	toSerialize["fallbackApprover"] = o.FallbackApprover

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOrgSettings200ResponseApprovalConfig) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"daysTillEscalation",
		"daysBetweenReminders",
		"maxReminders",
		"fallbackApprover",
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

	varGetOrgSettings200ResponseApprovalConfig := _GetOrgSettings200ResponseApprovalConfig{}

	if err = json.Unmarshal(bytes, &varGetOrgSettings200ResponseApprovalConfig); err == nil {
	*o = GetOrgSettings200ResponseApprovalConfig(varGetOrgSettings200ResponseApprovalConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "daysTillEscalation")
		delete(additionalProperties, "daysBetweenReminders")
		delete(additionalProperties, "maxReminders")
		delete(additionalProperties, "fallbackApprover")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrgSettings200ResponseApprovalConfig struct {
	value *GetOrgSettings200ResponseApprovalConfig
	isSet bool
}

func (v NullableGetOrgSettings200ResponseApprovalConfig) Get() *GetOrgSettings200ResponseApprovalConfig {
	return v.value
}

func (v *NullableGetOrgSettings200ResponseApprovalConfig) Set(val *GetOrgSettings200ResponseApprovalConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrgSettings200ResponseApprovalConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrgSettings200ResponseApprovalConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrgSettings200ResponseApprovalConfig(val *GetOrgSettings200ResponseApprovalConfig) *NullableGetOrgSettings200ResponseApprovalConfig {
	return &NullableGetOrgSettings200ResponseApprovalConfig{value: val, isSet: true}
}

func (v NullableGetOrgSettings200ResponseApprovalConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrgSettings200ResponseApprovalConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

