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

// ManualWorkItemState Indicates the state of the request processing for this item: * PENDING: The request for this item is awaiting processing. * APPROVED: The request for this item has been approved. * REJECTED: The request for this item was rejected. * EXPIRED: The request for this item expired with no action taken. * CANCELLED: The request for this item was cancelled with no user action. * ARCHIVED: The request for this item has been archived after completion.
type ManualWorkItemState string

// List of ManualWorkItemState
const (
	MANUALWORKITEMSTATE_PENDING ManualWorkItemState = "PENDING"
	MANUALWORKITEMSTATE_APPROVED ManualWorkItemState = "APPROVED"
	MANUALWORKITEMSTATE_REJECTED ManualWorkItemState = "REJECTED"
	MANUALWORKITEMSTATE_EXPIRED ManualWorkItemState = "EXPIRED"
	MANUALWORKITEMSTATE_CANCELLED ManualWorkItemState = "CANCELLED"
	MANUALWORKITEMSTATE_ARCHIVED ManualWorkItemState = "ARCHIVED"
)

// All allowed values of ManualWorkItemState enum
var AllowedManualWorkItemStateEnumValues = []ManualWorkItemState{
	"PENDING",
	"APPROVED",
	"REJECTED",
	"EXPIRED",
	"CANCELLED",
	"ARCHIVED",
}

func (v *ManualWorkItemState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ManualWorkItemState(value)
	for _, existing := range AllowedManualWorkItemStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ManualWorkItemState", value)
}

// NewManualWorkItemStateFromValue returns a pointer to a valid ManualWorkItemState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManualWorkItemStateFromValue(v string) (*ManualWorkItemState, error) {
	ev := ManualWorkItemState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ManualWorkItemState: valid values are %v", v, AllowedManualWorkItemStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManualWorkItemState) IsValid() bool {
	for _, existing := range AllowedManualWorkItemStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManualWorkItemState value
func (v ManualWorkItemState) Ptr() *ManualWorkItemState {
	return &v
}

type NullableManualWorkItemState struct {
	value *ManualWorkItemState
	isSet bool
}

func (v NullableManualWorkItemState) Get() *ManualWorkItemState {
	return v.value
}

func (v *NullableManualWorkItemState) Set(val *ManualWorkItemState) {
	v.value = val
	v.isSet = true
}

func (v NullableManualWorkItemState) IsSet() bool {
	return v.isSet
}

func (v *NullableManualWorkItemState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualWorkItemState(val *ManualWorkItemState) *NullableManualWorkItemState {
	return &NullableManualWorkItemState{value: val, isSet: true}
}

func (v NullableManualWorkItemState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualWorkItemState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

