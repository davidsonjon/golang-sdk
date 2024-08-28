/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// IdentityAttributesChangedChangesInnerOldValue - The value of the identity attribute before it changed.
type IdentityAttributesChangedChangesInnerOldValue struct {
	ArrayOfString *[]string
	Bool *bool
	MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue *map[string]IdentityAttributesChangedChangesInnerOldValueOneOfValue
	String *string
}

// []stringAsIdentityAttributesChangedChangesInnerOldValue is a convenience function that returns []string wrapped in IdentityAttributesChangedChangesInnerOldValue
func ArrayOfStringAsIdentityAttributesChangedChangesInnerOldValue(v *[]string) IdentityAttributesChangedChangesInnerOldValue {
	return IdentityAttributesChangedChangesInnerOldValue{
		ArrayOfString: v,
	}
}

// boolAsIdentityAttributesChangedChangesInnerOldValue is a convenience function that returns bool wrapped in IdentityAttributesChangedChangesInnerOldValue
func BoolAsIdentityAttributesChangedChangesInnerOldValue(v *bool) IdentityAttributesChangedChangesInnerOldValue {
	return IdentityAttributesChangedChangesInnerOldValue{
		Bool: v,
	}
}

// map[string]IdentityAttributesChangedChangesInnerOldValueOneOfValueAsIdentityAttributesChangedChangesInnerOldValue is a convenience function that returns map[string]IdentityAttributesChangedChangesInnerOldValueOneOfValue wrapped in IdentityAttributesChangedChangesInnerOldValue
func MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValueAsIdentityAttributesChangedChangesInnerOldValue(v *map[string]IdentityAttributesChangedChangesInnerOldValueOneOfValue) IdentityAttributesChangedChangesInnerOldValue {
	return IdentityAttributesChangedChangesInnerOldValue{
		MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue: v,
	}
}

// stringAsIdentityAttributesChangedChangesInnerOldValue is a convenience function that returns string wrapped in IdentityAttributesChangedChangesInnerOldValue
func StringAsIdentityAttributesChangedChangesInnerOldValue(v *string) IdentityAttributesChangedChangesInnerOldValue {
	return IdentityAttributesChangedChangesInnerOldValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityAttributesChangedChangesInnerOldValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

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

	// try to unmarshal data into MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue)
	if err == nil {
		jsonMapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue, _ := json.Marshal(dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue)
		if string(jsonMapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue) == "{}" { // empty struct
			dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue); err != nil {
				dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue = nil
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
		dst.ArrayOfString = nil
		dst.Bool = nil
		dst.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IdentityAttributesChangedChangesInnerOldValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IdentityAttributesChangedChangesInnerOldValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityAttributesChangedChangesInnerOldValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue != nil {
		return json.Marshal(&src.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityAttributesChangedChangesInnerOldValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue != nil {
		return obj.MapmapOfStringIdentityAttributesChangedChangesInnerOldValueOneOfValue
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableIdentityAttributesChangedChangesInnerOldValue struct {
	value *IdentityAttributesChangedChangesInnerOldValue
	isSet bool
}

func (v NullableIdentityAttributesChangedChangesInnerOldValue) Get() *IdentityAttributesChangedChangesInnerOldValue {
	return v.value
}

func (v *NullableIdentityAttributesChangedChangesInnerOldValue) Set(val *IdentityAttributesChangedChangesInnerOldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributesChangedChangesInnerOldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributesChangedChangesInnerOldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributesChangedChangesInnerOldValue(val *IdentityAttributesChangedChangesInnerOldValue) *NullableIdentityAttributesChangedChangesInnerOldValue {
	return &NullableIdentityAttributesChangedChangesInnerOldValue{value: val, isSet: true}
}

func (v NullableIdentityAttributesChangedChangesInnerOldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributesChangedChangesInnerOldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


