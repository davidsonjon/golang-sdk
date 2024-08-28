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

// JsonPatchOperationValue - The value to be used for the operation, required for \"add\" and \"replace\" operations
type JsonPatchOperationValue struct {
	ArrayOfArrayInner *[]ArrayInner
	Bool *bool
	Int32 *int32
	MapmapOfStringAny *map[string]interface{}
	String *string
}

// []ArrayInnerAsJsonPatchOperationValue is a convenience function that returns []ArrayInner wrapped in JsonPatchOperationValue
func ArrayOfArrayInnerAsJsonPatchOperationValue(v *[]ArrayInner) JsonPatchOperationValue {
	return JsonPatchOperationValue{
		ArrayOfArrayInner: v,
	}
}

// boolAsJsonPatchOperationValue is a convenience function that returns bool wrapped in JsonPatchOperationValue
func BoolAsJsonPatchOperationValue(v *bool) JsonPatchOperationValue {
	return JsonPatchOperationValue{
		Bool: v,
	}
}

// int32AsJsonPatchOperationValue is a convenience function that returns int32 wrapped in JsonPatchOperationValue
func Int32AsJsonPatchOperationValue(v *int32) JsonPatchOperationValue {
	return JsonPatchOperationValue{
		Int32: v,
	}
}

// map[string]interface{}AsJsonPatchOperationValue is a convenience function that returns map[string]interface{} wrapped in JsonPatchOperationValue
func MapmapOfStringAnyAsJsonPatchOperationValue(v *map[string]interface{}) JsonPatchOperationValue {
	return JsonPatchOperationValue{
		MapmapOfStringAny: v,
	}
}

// stringAsJsonPatchOperationValue is a convenience function that returns string wrapped in JsonPatchOperationValue
func StringAsJsonPatchOperationValue(v *string) JsonPatchOperationValue {
	return JsonPatchOperationValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *JsonPatchOperationValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfArrayInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfArrayInner)
	if err == nil {
		jsonArrayOfArrayInner, _ := json.Marshal(dst.ArrayOfArrayInner)
		if string(jsonArrayOfArrayInner) == "{}" { // empty struct
			dst.ArrayOfArrayInner = nil
		} else {
			if err = validator.Validate(dst.ArrayOfArrayInner); err != nil {
				dst.ArrayOfArrayInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfArrayInner = nil
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
		dst.ArrayOfArrayInner = nil
		dst.Bool = nil
		dst.Int32 = nil
		dst.MapmapOfStringAny = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(JsonPatchOperationValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(JsonPatchOperationValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src JsonPatchOperationValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfArrayInner != nil {
		return json.Marshal(&src.ArrayOfArrayInner)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *JsonPatchOperationValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfArrayInner != nil {
		return obj.ArrayOfArrayInner
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableJsonPatchOperationValue struct {
	value *JsonPatchOperationValue
	isSet bool
}

func (v NullableJsonPatchOperationValue) Get() *JsonPatchOperationValue {
	return v.value
}

func (v *NullableJsonPatchOperationValue) Set(val *JsonPatchOperationValue) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchOperationValue) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchOperationValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchOperationValue(val *JsonPatchOperationValue) *NullableJsonPatchOperationValue {
	return &NullableJsonPatchOperationValue{value: val, isSet: true}
}

func (v NullableJsonPatchOperationValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchOperationValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


