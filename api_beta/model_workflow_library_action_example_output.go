/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// WorkflowLibraryActionExampleOutput - struct for WorkflowLibraryActionExampleOutput
type WorkflowLibraryActionExampleOutput struct {
	ArrayOfMapmapOfStringAny *[]map[string]interface{}
	MapmapOfStringAny *map[string]interface{}
}

// []map[string]interface{}AsWorkflowLibraryActionExampleOutput is a convenience function that returns []map[string]interface{} wrapped in WorkflowLibraryActionExampleOutput
func ArrayOfMapmapOfStringAnyAsWorkflowLibraryActionExampleOutput(v *[]map[string]interface{}) WorkflowLibraryActionExampleOutput {
	return WorkflowLibraryActionExampleOutput{
		ArrayOfMapmapOfStringAny: v,
	}
}

// map[string]interface{}AsWorkflowLibraryActionExampleOutput is a convenience function that returns map[string]interface{} wrapped in WorkflowLibraryActionExampleOutput
func MapmapOfStringAnyAsWorkflowLibraryActionExampleOutput(v *map[string]interface{}) WorkflowLibraryActionExampleOutput {
	return WorkflowLibraryActionExampleOutput{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowLibraryActionExampleOutput) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfMapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.ArrayOfMapmapOfStringAny)
	if err == nil {
		jsonArrayOfMapmapOfStringAny, _ := json.Marshal(dst.ArrayOfMapmapOfStringAny)
		if string(jsonArrayOfMapmapOfStringAny) == "{}" { // empty struct
			dst.ArrayOfMapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.ArrayOfMapmapOfStringAny); err != nil {
				dst.ArrayOfMapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfMapmapOfStringAny = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfMapmapOfStringAny = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WorkflowLibraryActionExampleOutput)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WorkflowLibraryActionExampleOutput)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowLibraryActionExampleOutput) MarshalJSON() ([]byte, error) {
	if src.ArrayOfMapmapOfStringAny != nil {
		return json.Marshal(&src.ArrayOfMapmapOfStringAny)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowLibraryActionExampleOutput) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfMapmapOfStringAny != nil {
		return obj.ArrayOfMapmapOfStringAny
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowLibraryActionExampleOutput struct {
	value *WorkflowLibraryActionExampleOutput
	isSet bool
}

func (v NullableWorkflowLibraryActionExampleOutput) Get() *WorkflowLibraryActionExampleOutput {
	return v.value
}

func (v *NullableWorkflowLibraryActionExampleOutput) Set(val *WorkflowLibraryActionExampleOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryActionExampleOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryActionExampleOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryActionExampleOutput(val *WorkflowLibraryActionExampleOutput) *NullableWorkflowLibraryActionExampleOutput {
	return &NullableWorkflowLibraryActionExampleOutput{value: val, isSet: true}
}

func (v NullableWorkflowLibraryActionExampleOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryActionExampleOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


