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


// ListCompleteWorkflowLibrary200ResponseInner struct for ListCompleteWorkflowLibrary200ResponseInner
type ListCompleteWorkflowLibrary200ResponseInner struct {
	WorkflowLibraryAction *WorkflowLibraryAction
	WorkflowLibraryOperator *WorkflowLibraryOperator
	WorkflowLibraryTrigger *WorkflowLibraryTrigger
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ListCompleteWorkflowLibrary200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into WorkflowLibraryAction
	err = json.Unmarshal(data, &dst.WorkflowLibraryAction);
	if err == nil {
		jsonWorkflowLibraryAction, _ := json.Marshal(dst.WorkflowLibraryAction)
		if string(jsonWorkflowLibraryAction) == "{}" { // empty struct
			dst.WorkflowLibraryAction = nil
		} else {
			return nil // data stored in dst.WorkflowLibraryAction, return on the first match
		}
	} else {
		dst.WorkflowLibraryAction = nil
	}

	// try to unmarshal JSON data into WorkflowLibraryOperator
	err = json.Unmarshal(data, &dst.WorkflowLibraryOperator);
	if err == nil {
		jsonWorkflowLibraryOperator, _ := json.Marshal(dst.WorkflowLibraryOperator)
		if string(jsonWorkflowLibraryOperator) == "{}" { // empty struct
			dst.WorkflowLibraryOperator = nil
		} else {
			return nil // data stored in dst.WorkflowLibraryOperator, return on the first match
		}
	} else {
		dst.WorkflowLibraryOperator = nil
	}

	// try to unmarshal JSON data into WorkflowLibraryTrigger
	err = json.Unmarshal(data, &dst.WorkflowLibraryTrigger);
	if err == nil {
		jsonWorkflowLibraryTrigger, _ := json.Marshal(dst.WorkflowLibraryTrigger)
		if string(jsonWorkflowLibraryTrigger) == "{}" { // empty struct
			dst.WorkflowLibraryTrigger = nil
		} else {
			return nil // data stored in dst.WorkflowLibraryTrigger, return on the first match
		}
	} else {
		dst.WorkflowLibraryTrigger = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ListCompleteWorkflowLibrary200ResponseInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ListCompleteWorkflowLibrary200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.WorkflowLibraryAction != nil {
		return json.Marshal(&src.WorkflowLibraryAction)
	}

	if src.WorkflowLibraryOperator != nil {
		return json.Marshal(&src.WorkflowLibraryOperator)
	}

	if src.WorkflowLibraryTrigger != nil {
		return json.Marshal(&src.WorkflowLibraryTrigger)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableListCompleteWorkflowLibrary200ResponseInner struct {
	value *ListCompleteWorkflowLibrary200ResponseInner
	isSet bool
}

func (v NullableListCompleteWorkflowLibrary200ResponseInner) Get() *ListCompleteWorkflowLibrary200ResponseInner {
	return v.value
}

func (v *NullableListCompleteWorkflowLibrary200ResponseInner) Set(val *ListCompleteWorkflowLibrary200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListCompleteWorkflowLibrary200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListCompleteWorkflowLibrary200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCompleteWorkflowLibrary200ResponseInner(val *ListCompleteWorkflowLibrary200ResponseInner) *NullableListCompleteWorkflowLibrary200ResponseInner {
	return &NullableListCompleteWorkflowLibrary200ResponseInner{value: val, isSet: true}
}

func (v NullableListCompleteWorkflowLibrary200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCompleteWorkflowLibrary200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


