/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// ApprovalScheme1 Describes the individual or group that is responsible for an approval step.
type ApprovalScheme1 string

// List of ApprovalScheme_1
const (
	APPROVALSCHEME1_APP_OWNER ApprovalScheme1 = "APP_OWNER"
	APPROVALSCHEME1_SOURCE_OWNER ApprovalScheme1 = "SOURCE_OWNER"
	APPROVALSCHEME1_MANAGER ApprovalScheme1 = "MANAGER"
	APPROVALSCHEME1_ROLE_OWNER ApprovalScheme1 = "ROLE_OWNER"
	APPROVALSCHEME1_ACCESS_PROFILE_OWNER ApprovalScheme1 = "ACCESS_PROFILE_OWNER"
	APPROVALSCHEME1_GOVERNANCE_GROUP ApprovalScheme1 = "GOVERNANCE_GROUP"
)

// All allowed values of ApprovalScheme1 enum
var AllowedApprovalScheme1EnumValues = []ApprovalScheme1{
	"APP_OWNER",
	"SOURCE_OWNER",
	"MANAGER",
	"ROLE_OWNER",
	"ACCESS_PROFILE_OWNER",
	"GOVERNANCE_GROUP",
}

func (v *ApprovalScheme1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovalScheme1(value)
	for _, existing := range AllowedApprovalScheme1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApprovalScheme1", value)
}

// NewApprovalScheme1FromValue returns a pointer to a valid ApprovalScheme1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovalScheme1FromValue(v string) (*ApprovalScheme1, error) {
	ev := ApprovalScheme1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovalScheme1: valid values are %v", v, AllowedApprovalScheme1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovalScheme1) IsValid() bool {
	for _, existing := range AllowedApprovalScheme1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApprovalScheme_1 value
func (v ApprovalScheme1) Ptr() *ApprovalScheme1 {
	return &v
}

type NullableApprovalScheme1 struct {
	value *ApprovalScheme1
	isSet bool
}

func (v NullableApprovalScheme1) Get() *ApprovalScheme1 {
	return v.value
}

func (v *NullableApprovalScheme1) Set(val *ApprovalScheme1) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalScheme1) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalScheme1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalScheme1(val *ApprovalScheme1) *NullableApprovalScheme1 {
	return &NullableApprovalScheme1{value: val, isSet: true}
}

func (v NullableApprovalScheme1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalScheme1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

