/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// TriggerType The type of trigger.
type TriggerType string

// List of TriggerType
const (
	TRIGGERTYPE_REQUEST_RESPONSE TriggerType = "REQUEST_RESPONSE"
	TRIGGERTYPE_FIRE_AND_FORGET TriggerType = "FIRE_AND_FORGET"
)

// All allowed values of TriggerType enum
var AllowedTriggerTypeEnumValues = []TriggerType{
	"REQUEST_RESPONSE",
	"FIRE_AND_FORGET",
}

func (v *TriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerType(value)
	for _, existing := range AllowedTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerType", value)
}

// NewTriggerTypeFromValue returns a pointer to a valid TriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerTypeFromValue(v string) (*TriggerType, error) {
	ev := TriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerType: valid values are %v", v, AllowedTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerType) IsValid() bool {
	for _, existing := range AllowedTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerType value
func (v TriggerType) Ptr() *TriggerType {
	return &v
}

type NullableTriggerType struct {
	value *TriggerType
	isSet bool
}

func (v NullableTriggerType) Get() *TriggerType {
	return v.value
}

func (v *NullableTriggerType) Set(val *TriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerType(val *TriggerType) *NullableTriggerType {
	return &NullableTriggerType{value: val, isSet: true}
}

func (v NullableTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

