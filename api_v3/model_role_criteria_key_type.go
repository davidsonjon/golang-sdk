/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// RoleCriteriaKeyType Indicates whether the associated criteria represents an expression on identity attributes, account attributes, or entitlements, respectively.
type RoleCriteriaKeyType string

// List of RoleCriteriaKeyType
const (
	ROLECRITERIAKEYTYPE_IDENTITY RoleCriteriaKeyType = "IDENTITY"
	ROLECRITERIAKEYTYPE_ACCOUNT RoleCriteriaKeyType = "ACCOUNT"
	ROLECRITERIAKEYTYPE_ENTITLEMENT RoleCriteriaKeyType = "ENTITLEMENT"
)

// All allowed values of RoleCriteriaKeyType enum
var AllowedRoleCriteriaKeyTypeEnumValues = []RoleCriteriaKeyType{
	"IDENTITY",
	"ACCOUNT",
	"ENTITLEMENT",
}

func (v *RoleCriteriaKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleCriteriaKeyType(value)
	for _, existing := range AllowedRoleCriteriaKeyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleCriteriaKeyType", value)
}

// NewRoleCriteriaKeyTypeFromValue returns a pointer to a valid RoleCriteriaKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleCriteriaKeyTypeFromValue(v string) (*RoleCriteriaKeyType, error) {
	ev := RoleCriteriaKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoleCriteriaKeyType: valid values are %v", v, AllowedRoleCriteriaKeyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleCriteriaKeyType) IsValid() bool {
	for _, existing := range AllowedRoleCriteriaKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleCriteriaKeyType value
func (v RoleCriteriaKeyType) Ptr() *RoleCriteriaKeyType {
	return &v
}

type NullableRoleCriteriaKeyType struct {
	value *RoleCriteriaKeyType
	isSet bool
}

func (v NullableRoleCriteriaKeyType) Get() *RoleCriteriaKeyType {
	return v.value
}

func (v *NullableRoleCriteriaKeyType) Set(val *RoleCriteriaKeyType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCriteriaKeyType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCriteriaKeyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCriteriaKeyType(val *RoleCriteriaKeyType) *NullableRoleCriteriaKeyType {
	return &NullableRoleCriteriaKeyType{value: val, isSet: true}
}

func (v NullableRoleCriteriaKeyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCriteriaKeyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

