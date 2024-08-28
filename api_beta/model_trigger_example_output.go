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

// TriggerExampleOutput - An example of the JSON payload that will be sent by the subscribed service to the trigger in response to an event.
type TriggerExampleOutput struct {
	AccessRequestDynamicApprover1 *AccessRequestDynamicApprover1
	AccessRequestPreApproval1 *AccessRequestPreApproval1
}

// AccessRequestDynamicApprover1AsTriggerExampleOutput is a convenience function that returns AccessRequestDynamicApprover1 wrapped in TriggerExampleOutput
func AccessRequestDynamicApprover1AsTriggerExampleOutput(v *AccessRequestDynamicApprover1) TriggerExampleOutput {
	return TriggerExampleOutput{
		AccessRequestDynamicApprover1: v,
	}
}

// AccessRequestPreApproval1AsTriggerExampleOutput is a convenience function that returns AccessRequestPreApproval1 wrapped in TriggerExampleOutput
func AccessRequestPreApproval1AsTriggerExampleOutput(v *AccessRequestPreApproval1) TriggerExampleOutput {
	return TriggerExampleOutput{
		AccessRequestPreApproval1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TriggerExampleOutput) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into AccessRequestDynamicApprover1
	err = newStrictDecoder(data).Decode(&dst.AccessRequestDynamicApprover1)
	if err == nil {
		jsonAccessRequestDynamicApprover1, _ := json.Marshal(dst.AccessRequestDynamicApprover1)
		if string(jsonAccessRequestDynamicApprover1) == "{}" { // empty struct
			dst.AccessRequestDynamicApprover1 = nil
		} else {
			if err = validator.Validate(dst.AccessRequestDynamicApprover1); err != nil {
				dst.AccessRequestDynamicApprover1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AccessRequestDynamicApprover1 = nil
	}

	// try to unmarshal data into AccessRequestPreApproval1
	err = newStrictDecoder(data).Decode(&dst.AccessRequestPreApproval1)
	if err == nil {
		jsonAccessRequestPreApproval1, _ := json.Marshal(dst.AccessRequestPreApproval1)
		if string(jsonAccessRequestPreApproval1) == "{}" { // empty struct
			dst.AccessRequestPreApproval1 = nil
		} else {
			if err = validator.Validate(dst.AccessRequestPreApproval1); err != nil {
				dst.AccessRequestPreApproval1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AccessRequestPreApproval1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessRequestDynamicApprover1 = nil
		dst.AccessRequestPreApproval1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TriggerExampleOutput)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TriggerExampleOutput)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TriggerExampleOutput) MarshalJSON() ([]byte, error) {
	if src.AccessRequestDynamicApprover1 != nil {
		return json.Marshal(&src.AccessRequestDynamicApprover1)
	}

	if src.AccessRequestPreApproval1 != nil {
		return json.Marshal(&src.AccessRequestPreApproval1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TriggerExampleOutput) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessRequestDynamicApprover1 != nil {
		return obj.AccessRequestDynamicApprover1
	}

	if obj.AccessRequestPreApproval1 != nil {
		return obj.AccessRequestPreApproval1
	}

	// all schemas are nil
	return nil
}

type NullableTriggerExampleOutput struct {
	value *TriggerExampleOutput
	isSet bool
}

func (v NullableTriggerExampleOutput) Get() *TriggerExampleOutput {
	return v.value
}

func (v *NullableTriggerExampleOutput) Set(val *TriggerExampleOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerExampleOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerExampleOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerExampleOutput(val *TriggerExampleOutput) *NullableTriggerExampleOutput {
	return &NullableTriggerExampleOutput{value: val, isSet: true}
}

func (v NullableTriggerExampleOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerExampleOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


