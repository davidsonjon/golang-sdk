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

// AccountActivityApprovalStatus The state of an approval status
type AccountActivityApprovalStatus string

// List of AccountActivityApprovalStatus
const (
	ACCOUNTACTIVITYAPPROVALSTATUS_FINISHED AccountActivityApprovalStatus = "FINISHED"
	ACCOUNTACTIVITYAPPROVALSTATUS_REJECTED AccountActivityApprovalStatus = "REJECTED"
	ACCOUNTACTIVITYAPPROVALSTATUS_RETURNED AccountActivityApprovalStatus = "RETURNED"
	ACCOUNTACTIVITYAPPROVALSTATUS_EXPIRED AccountActivityApprovalStatus = "EXPIRED"
	ACCOUNTACTIVITYAPPROVALSTATUS_PENDING AccountActivityApprovalStatus = "PENDING"
	ACCOUNTACTIVITYAPPROVALSTATUS_CANCELED AccountActivityApprovalStatus = "CANCELED"
	ACCOUNTACTIVITYAPPROVALSTATUS_NULL AccountActivityApprovalStatus = "null"
)

// All allowed values of AccountActivityApprovalStatus enum
var AllowedAccountActivityApprovalStatusEnumValues = []AccountActivityApprovalStatus{
	"FINISHED",
	"REJECTED",
	"RETURNED",
	"EXPIRED",
	"PENDING",
	"CANCELED",
	"null",
}

func (v *AccountActivityApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountActivityApprovalStatus(value)
	for _, existing := range AllowedAccountActivityApprovalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountActivityApprovalStatus", value)
}

// NewAccountActivityApprovalStatusFromValue returns a pointer to a valid AccountActivityApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountActivityApprovalStatusFromValue(v string) (*AccountActivityApprovalStatus, error) {
	ev := AccountActivityApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountActivityApprovalStatus: valid values are %v", v, AllowedAccountActivityApprovalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountActivityApprovalStatus) IsValid() bool {
	for _, existing := range AllowedAccountActivityApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountActivityApprovalStatus value
func (v AccountActivityApprovalStatus) Ptr() *AccountActivityApprovalStatus {
	return &v
}

type NullableAccountActivityApprovalStatus struct {
	value *AccountActivityApprovalStatus
	isSet bool
}

func (v NullableAccountActivityApprovalStatus) Get() *AccountActivityApprovalStatus {
	return v.value
}

func (v *NullableAccountActivityApprovalStatus) Set(val *AccountActivityApprovalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountActivityApprovalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountActivityApprovalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountActivityApprovalStatus(val *AccountActivityApprovalStatus) *NullableAccountActivityApprovalStatus {
	return &NullableAccountActivityApprovalStatus{value: val, isSet: true}
}

func (v NullableAccountActivityApprovalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountActivityApprovalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

