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

// IdentityAttributesChangedChangesInnerOldValueOneOfValue - struct for IdentityAttributesChangedChangesInnerOldValueOneOfValue
type IdentityAttributesChangedChangesInnerOldValueOneOfValue struct {
	Bool *bool
	Float32 *float32
	Int32 *int32
	String *string
}

// boolAsIdentityAttributesChangedChangesInnerOldValueOneOfValue is a convenience function that returns bool wrapped in IdentityAttributesChangedChangesInnerOldValueOneOfValue
func BoolAsIdentityAttributesChangedChangesInnerOldValueOneOfValue(v *bool) IdentityAttributesChangedChangesInnerOldValueOneOfValue {
	return IdentityAttributesChangedChangesInnerOldValueOneOfValue{
		Bool: v,
	}
}

// float32AsIdentityAttributesChangedChangesInnerOldValueOneOfValue is a convenience function that returns float32 wrapped in IdentityAttributesChangedChangesInnerOldValueOneOfValue
func Float32AsIdentityAttributesChangedChangesInnerOldValueOneOfValue(v *float32) IdentityAttributesChangedChangesInnerOldValueOneOfValue {
	return IdentityAttributesChangedChangesInnerOldValueOneOfValue{
		Float32: v,
	}
}

// int32AsIdentityAttributesChangedChangesInnerOldValueOneOfValue is a convenience function that returns int32 wrapped in IdentityAttributesChangedChangesInnerOldValueOneOfValue
func Int32AsIdentityAttributesChangedChangesInnerOldValueOneOfValue(v *int32) IdentityAttributesChangedChangesInnerOldValueOneOfValue {
	return IdentityAttributesChangedChangesInnerOldValueOneOfValue{
		Int32: v,
	}
}

// stringAsIdentityAttributesChangedChangesInnerOldValueOneOfValue is a convenience function that returns string wrapped in IdentityAttributesChangedChangesInnerOldValueOneOfValue
func StringAsIdentityAttributesChangedChangesInnerOldValueOneOfValue(v *string) IdentityAttributesChangedChangesInnerOldValueOneOfValue {
	return IdentityAttributesChangedChangesInnerOldValueOneOfValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityAttributesChangedChangesInnerOldValueOneOfValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			if err = validator.Validate(dst.Float32); err != nil {
				dst.Float32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IdentityAttributesChangedChangesInnerOldValueOneOfValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IdentityAttributesChangedChangesInnerOldValueOneOfValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityAttributesChangedChangesInnerOldValueOneOfValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityAttributesChangedChangesInnerOldValueOneOfValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue struct {
	value *IdentityAttributesChangedChangesInnerOldValueOneOfValue
	isSet bool
}

func (v NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue) Get() *IdentityAttributesChangedChangesInnerOldValueOneOfValue {
	return v.value
}

func (v *NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue) Set(val *IdentityAttributesChangedChangesInnerOldValueOneOfValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributesChangedChangesInnerOldValueOneOfValue(val *IdentityAttributesChangedChangesInnerOldValueOneOfValue) *NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue {
	return &NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue{value: val, isSet: true}
}

func (v NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributesChangedChangesInnerOldValueOneOfValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


