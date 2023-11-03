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

// RoleMiningSessionDtoCreatedBy - The session created by details
type RoleMiningSessionDtoCreatedBy struct {
	EntityCreatedByDTO *EntityCreatedByDTO
	String *string
}

// EntityCreatedByDTOAsRoleMiningSessionDtoCreatedBy is a convenience function that returns EntityCreatedByDTO wrapped in RoleMiningSessionDtoCreatedBy
func EntityCreatedByDTOAsRoleMiningSessionDtoCreatedBy(v *EntityCreatedByDTO) RoleMiningSessionDtoCreatedBy {
	return RoleMiningSessionDtoCreatedBy{
		EntityCreatedByDTO: v,
	}
}

// stringAsRoleMiningSessionDtoCreatedBy is a convenience function that returns string wrapped in RoleMiningSessionDtoCreatedBy
func StringAsRoleMiningSessionDtoCreatedBy(v *string) RoleMiningSessionDtoCreatedBy {
	return RoleMiningSessionDtoCreatedBy{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RoleMiningSessionDtoCreatedBy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntityCreatedByDTO
	err = newStrictDecoder(data).Decode(&dst.EntityCreatedByDTO)
	if err == nil {
		jsonEntityCreatedByDTO, _ := json.Marshal(dst.EntityCreatedByDTO)
		if string(jsonEntityCreatedByDTO) == "{}" { // empty struct
			dst.EntityCreatedByDTO = nil
		} else {
			match++
		}
	} else {
		dst.EntityCreatedByDTO = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EntityCreatedByDTO = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RoleMiningSessionDtoCreatedBy)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RoleMiningSessionDtoCreatedBy)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RoleMiningSessionDtoCreatedBy) MarshalJSON() ([]byte, error) {
	if src.EntityCreatedByDTO != nil {
		return json.Marshal(&src.EntityCreatedByDTO)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RoleMiningSessionDtoCreatedBy) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EntityCreatedByDTO != nil {
		return obj.EntityCreatedByDTO
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRoleMiningSessionDtoCreatedBy struct {
	value *RoleMiningSessionDtoCreatedBy
	isSet bool
}

func (v NullableRoleMiningSessionDtoCreatedBy) Get() *RoleMiningSessionDtoCreatedBy {
	return v.value
}

func (v *NullableRoleMiningSessionDtoCreatedBy) Set(val *RoleMiningSessionDtoCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningSessionDtoCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningSessionDtoCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningSessionDtoCreatedBy(val *RoleMiningSessionDtoCreatedBy) *NullableRoleMiningSessionDtoCreatedBy {
	return &NullableRoleMiningSessionDtoCreatedBy{value: val, isSet: true}
}

func (v NullableRoleMiningSessionDtoCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningSessionDtoCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

