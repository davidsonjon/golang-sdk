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

// RoleMiningPotentialRoleSummaryCreatedBy - The potential role created by details
type RoleMiningPotentialRoleSummaryCreatedBy struct {
	EntityCreatedByDTO *EntityCreatedByDTO
	String *string
}

// EntityCreatedByDTOAsRoleMiningPotentialRoleSummaryCreatedBy is a convenience function that returns EntityCreatedByDTO wrapped in RoleMiningPotentialRoleSummaryCreatedBy
func EntityCreatedByDTOAsRoleMiningPotentialRoleSummaryCreatedBy(v *EntityCreatedByDTO) RoleMiningPotentialRoleSummaryCreatedBy {
	return RoleMiningPotentialRoleSummaryCreatedBy{
		EntityCreatedByDTO: v,
	}
}

// stringAsRoleMiningPotentialRoleSummaryCreatedBy is a convenience function that returns string wrapped in RoleMiningPotentialRoleSummaryCreatedBy
func StringAsRoleMiningPotentialRoleSummaryCreatedBy(v *string) RoleMiningPotentialRoleSummaryCreatedBy {
	return RoleMiningPotentialRoleSummaryCreatedBy{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RoleMiningPotentialRoleSummaryCreatedBy) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(RoleMiningPotentialRoleSummaryCreatedBy)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RoleMiningPotentialRoleSummaryCreatedBy)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RoleMiningPotentialRoleSummaryCreatedBy) MarshalJSON() ([]byte, error) {
	if src.EntityCreatedByDTO != nil {
		return json.Marshal(&src.EntityCreatedByDTO)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RoleMiningPotentialRoleSummaryCreatedBy) GetActualInstance() (interface{}) {
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

type NullableRoleMiningPotentialRoleSummaryCreatedBy struct {
	value *RoleMiningPotentialRoleSummaryCreatedBy
	isSet bool
}

func (v NullableRoleMiningPotentialRoleSummaryCreatedBy) Get() *RoleMiningPotentialRoleSummaryCreatedBy {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleSummaryCreatedBy) Set(val *RoleMiningPotentialRoleSummaryCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleSummaryCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleSummaryCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleSummaryCreatedBy(val *RoleMiningPotentialRoleSummaryCreatedBy) *NullableRoleMiningPotentialRoleSummaryCreatedBy {
	return &NullableRoleMiningPotentialRoleSummaryCreatedBy{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleSummaryCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleSummaryCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

