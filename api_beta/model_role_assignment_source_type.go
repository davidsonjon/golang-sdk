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

// RoleAssignmentSourceType Type which indicates how a particular Identity obtained a particular Role
type RoleAssignmentSourceType string

// List of RoleAssignmentSourceType
const (
	ROLEASSIGNMENTSOURCETYPE_ACCESS_REQUEST RoleAssignmentSourceType = "ACCESS_REQUEST"
	ROLEASSIGNMENTSOURCETYPE_ROLE_MEMBERSHIP RoleAssignmentSourceType = "ROLE_MEMBERSHIP"
)

// All allowed values of RoleAssignmentSourceType enum
var AllowedRoleAssignmentSourceTypeEnumValues = []RoleAssignmentSourceType{
	"ACCESS_REQUEST",
	"ROLE_MEMBERSHIP",
}

func (v *RoleAssignmentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleAssignmentSourceType(value)
	for _, existing := range AllowedRoleAssignmentSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleAssignmentSourceType", value)
}

// NewRoleAssignmentSourceTypeFromValue returns a pointer to a valid RoleAssignmentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleAssignmentSourceTypeFromValue(v string) (*RoleAssignmentSourceType, error) {
	ev := RoleAssignmentSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoleAssignmentSourceType: valid values are %v", v, AllowedRoleAssignmentSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleAssignmentSourceType) IsValid() bool {
	for _, existing := range AllowedRoleAssignmentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleAssignmentSourceType value
func (v RoleAssignmentSourceType) Ptr() *RoleAssignmentSourceType {
	return &v
}

type NullableRoleAssignmentSourceType struct {
	value *RoleAssignmentSourceType
	isSet bool
}

func (v NullableRoleAssignmentSourceType) Get() *RoleAssignmentSourceType {
	return v.value
}

func (v *NullableRoleAssignmentSourceType) Set(val *RoleAssignmentSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentSourceType(val *RoleAssignmentSourceType) *NullableRoleAssignmentSourceType {
	return &NullableRoleAssignmentSourceType{value: val, isSet: true}
}

func (v NullableRoleAssignmentSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

