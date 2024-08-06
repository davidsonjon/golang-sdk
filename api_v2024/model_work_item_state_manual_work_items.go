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

// WorkItemStateManualWorkItems The state of a work item
type WorkItemStateManualWorkItems string

// List of WorkItemStateManualWorkItems
const (
	WORKITEMSTATEMANUALWORKITEMS_FINISHED WorkItemStateManualWorkItems = "Finished"
	WORKITEMSTATEMANUALWORKITEMS_REJECTED WorkItemStateManualWorkItems = "Rejected"
	WORKITEMSTATEMANUALWORKITEMS_RETURNED WorkItemStateManualWorkItems = "Returned"
	WORKITEMSTATEMANUALWORKITEMS_EXPIRED WorkItemStateManualWorkItems = "Expired"
	WORKITEMSTATEMANUALWORKITEMS_PENDING WorkItemStateManualWorkItems = "Pending"
	WORKITEMSTATEMANUALWORKITEMS_CANCELED WorkItemStateManualWorkItems = "Canceled"
)

// All allowed values of WorkItemStateManualWorkItems enum
var AllowedWorkItemStateManualWorkItemsEnumValues = []WorkItemStateManualWorkItems{
	"Finished",
	"Rejected",
	"Returned",
	"Expired",
	"Pending",
	"Canceled",
}

func (v *WorkItemStateManualWorkItems) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkItemStateManualWorkItems(value)
	for _, existing := range AllowedWorkItemStateManualWorkItemsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkItemStateManualWorkItems", value)
}

// NewWorkItemStateManualWorkItemsFromValue returns a pointer to a valid WorkItemStateManualWorkItems
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkItemStateManualWorkItemsFromValue(v string) (*WorkItemStateManualWorkItems, error) {
	ev := WorkItemStateManualWorkItems(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkItemStateManualWorkItems: valid values are %v", v, AllowedWorkItemStateManualWorkItemsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkItemStateManualWorkItems) IsValid() bool {
	for _, existing := range AllowedWorkItemStateManualWorkItemsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkItemStateManualWorkItems value
func (v WorkItemStateManualWorkItems) Ptr() *WorkItemStateManualWorkItems {
	return &v
}

type NullableWorkItemStateManualWorkItems struct {
	value *WorkItemStateManualWorkItems
	isSet bool
}

func (v NullableWorkItemStateManualWorkItems) Get() *WorkItemStateManualWorkItems {
	return v.value
}

func (v *NullableWorkItemStateManualWorkItems) Set(val *WorkItemStateManualWorkItems) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemStateManualWorkItems) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemStateManualWorkItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemStateManualWorkItems(val *WorkItemStateManualWorkItems) *NullableWorkItemStateManualWorkItems {
	return &NullableWorkItemStateManualWorkItems{value: val, isSet: true}
}

func (v NullableWorkItemStateManualWorkItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemStateManualWorkItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

