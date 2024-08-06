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

// CommonAccessType The type of access item.
type CommonAccessType string

// List of CommonAccessType
const (
	COMMONACCESSTYPE_ACCESS_PROFILE CommonAccessType = "ACCESS_PROFILE"
	COMMONACCESSTYPE_ROLE CommonAccessType = "ROLE"
)

// All allowed values of CommonAccessType enum
var AllowedCommonAccessTypeEnumValues = []CommonAccessType{
	"ACCESS_PROFILE",
	"ROLE",
}

func (v *CommonAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommonAccessType(value)
	for _, existing := range AllowedCommonAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommonAccessType", value)
}

// NewCommonAccessTypeFromValue returns a pointer to a valid CommonAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommonAccessTypeFromValue(v string) (*CommonAccessType, error) {
	ev := CommonAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommonAccessType: valid values are %v", v, AllowedCommonAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommonAccessType) IsValid() bool {
	for _, existing := range AllowedCommonAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommonAccessType value
func (v CommonAccessType) Ptr() *CommonAccessType {
	return &v
}

type NullableCommonAccessType struct {
	value *CommonAccessType
	isSet bool
}

func (v NullableCommonAccessType) Get() *CommonAccessType {
	return v.value
}

func (v *NullableCommonAccessType) Set(val *CommonAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAccessType(val *CommonAccessType) *NullableCommonAccessType {
	return &NullableCommonAccessType{value: val, isSet: true}
}

func (v NullableCommonAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

