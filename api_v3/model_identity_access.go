/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// IdentityAccess - struct for IdentityAccess
type IdentityAccess struct {
	AccessProfileEntitlement *AccessProfileEntitlement
	AccessProfileRole *AccessProfileRole
	AccessProfileSummary *AccessProfileSummary
}

// AccessProfileEntitlementAsIdentityAccess is a convenience function that returns AccessProfileEntitlement wrapped in IdentityAccess
func AccessProfileEntitlementAsIdentityAccess(v *AccessProfileEntitlement) IdentityAccess {
	return IdentityAccess{
		AccessProfileEntitlement: v,
	}
}

// AccessProfileRoleAsIdentityAccess is a convenience function that returns AccessProfileRole wrapped in IdentityAccess
func AccessProfileRoleAsIdentityAccess(v *AccessProfileRole) IdentityAccess {
	return IdentityAccess{
		AccessProfileRole: v,
	}
}

// AccessProfileSummaryAsIdentityAccess is a convenience function that returns AccessProfileSummary wrapped in IdentityAccess
func AccessProfileSummaryAsIdentityAccess(v *AccessProfileSummary) IdentityAccess {
	return IdentityAccess{
		AccessProfileSummary: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityAccess) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessProfileEntitlement
	err = newStrictDecoder(data).Decode(&dst.AccessProfileEntitlement)
	if err == nil {
		jsonAccessProfileEntitlement, _ := json.Marshal(dst.AccessProfileEntitlement)
		if string(jsonAccessProfileEntitlement) == "{}" { // empty struct
			dst.AccessProfileEntitlement = nil
		} else {
			if err = validator.Validate(dst.AccessProfileEntitlement); err != nil {
				dst.AccessProfileEntitlement = nil
			} else {
				match++
			}
		}
	} else {
		dst.AccessProfileEntitlement = nil
	}

	// try to unmarshal data into AccessProfileRole
	err = newStrictDecoder(data).Decode(&dst.AccessProfileRole)
	if err == nil {
		jsonAccessProfileRole, _ := json.Marshal(dst.AccessProfileRole)
		if string(jsonAccessProfileRole) == "{}" { // empty struct
			dst.AccessProfileRole = nil
		} else {
			if err = validator.Validate(dst.AccessProfileRole); err != nil {
				dst.AccessProfileRole = nil
			} else {
				match++
			}
		}
	} else {
		dst.AccessProfileRole = nil
	}

	// try to unmarshal data into AccessProfileSummary
	err = newStrictDecoder(data).Decode(&dst.AccessProfileSummary)
	if err == nil {
		jsonAccessProfileSummary, _ := json.Marshal(dst.AccessProfileSummary)
		if string(jsonAccessProfileSummary) == "{}" { // empty struct
			dst.AccessProfileSummary = nil
		} else {
			if err = validator.Validate(dst.AccessProfileSummary); err != nil {
				dst.AccessProfileSummary = nil
			} else {
				match++
			}
		}
	} else {
		dst.AccessProfileSummary = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessProfileEntitlement = nil
		dst.AccessProfileRole = nil
		dst.AccessProfileSummary = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IdentityAccess)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IdentityAccess)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityAccess) MarshalJSON() ([]byte, error) {
	if src.AccessProfileEntitlement != nil {
		return json.Marshal(&src.AccessProfileEntitlement)
	}

	if src.AccessProfileRole != nil {
		return json.Marshal(&src.AccessProfileRole)
	}

	if src.AccessProfileSummary != nil {
		return json.Marshal(&src.AccessProfileSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityAccess) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessProfileEntitlement != nil {
		return obj.AccessProfileEntitlement
	}

	if obj.AccessProfileRole != nil {
		return obj.AccessProfileRole
	}

	if obj.AccessProfileSummary != nil {
		return obj.AccessProfileSummary
	}

	// all schemas are nil
	return nil
}

type NullableIdentityAccess struct {
	value *IdentityAccess
	isSet bool
}

func (v NullableIdentityAccess) Get() *IdentityAccess {
	return v.value
}

func (v *NullableIdentityAccess) Set(val *IdentityAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAccess(val *IdentityAccess) *NullableIdentityAccess {
	return &NullableIdentityAccess{value: val, isSet: true}
}

func (v NullableIdentityAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


