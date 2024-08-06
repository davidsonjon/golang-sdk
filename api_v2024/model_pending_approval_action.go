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

// PendingApprovalAction Enum represents action that is being processed on an approval.
type PendingApprovalAction string

// List of PendingApprovalAction
const (
	PENDINGAPPROVALACTION_APPROVED PendingApprovalAction = "APPROVED"
	PENDINGAPPROVALACTION_REJECTED PendingApprovalAction = "REJECTED"
	PENDINGAPPROVALACTION_FORWARDED PendingApprovalAction = "FORWARDED"
)

// All allowed values of PendingApprovalAction enum
var AllowedPendingApprovalActionEnumValues = []PendingApprovalAction{
	"APPROVED",
	"REJECTED",
	"FORWARDED",
}

func (v *PendingApprovalAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PendingApprovalAction(value)
	for _, existing := range AllowedPendingApprovalActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PendingApprovalAction", value)
}

// NewPendingApprovalActionFromValue returns a pointer to a valid PendingApprovalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPendingApprovalActionFromValue(v string) (*PendingApprovalAction, error) {
	ev := PendingApprovalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PendingApprovalAction: valid values are %v", v, AllowedPendingApprovalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PendingApprovalAction) IsValid() bool {
	for _, existing := range AllowedPendingApprovalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PendingApprovalAction value
func (v PendingApprovalAction) Ptr() *PendingApprovalAction {
	return &v
}

type NullablePendingApprovalAction struct {
	value *PendingApprovalAction
	isSet bool
}

func (v NullablePendingApprovalAction) Get() *PendingApprovalAction {
	return v.value
}

func (v *NullablePendingApprovalAction) Set(val *PendingApprovalAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingApprovalAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingApprovalAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingApprovalAction(val *PendingApprovalAction) *NullablePendingApprovalAction {
	return &NullablePendingApprovalAction{value: val, isSet: true}
}

func (v NullablePendingApprovalAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingApprovalAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

