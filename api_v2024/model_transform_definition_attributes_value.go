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

// TransformDefinitionAttributesValue struct for TransformDefinitionAttributesValue
type TransformDefinitionAttributesValue struct {
	mapvar *map[string]interface{}
	stringvar *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransformDefinitionAttributesValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into mapvar
	err = json.Unmarshal(data, &dst.mapvar);
	if err == nil {
		jsonmapvar, _ := json.Marshal(dst.mapvar)
		if string(jsonmapvar) == "{}" { // empty struct
			dst.mapvar = nil
		} else {
			return nil // data stored in dst.mapvar, return on the first match
		}
	} else {
		dst.mapvar = nil
	}

	// try to unmarshal JSON data into stringvar
	err = json.Unmarshal(data, &dst.stringvar);
	if err == nil {
		jsonstringvar, _ := json.Marshal(dst.stringvar)
		if string(jsonstringvar) == "{}" { // empty struct
			dst.stringvar = nil
		} else {
			return nil // data stored in dst.stringvar, return on the first match
		}
	} else {
		dst.stringvar = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransformDefinitionAttributesValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransformDefinitionAttributesValue) MarshalJSON() ([]byte, error) {
	if src.mapvar != nil {
		return json.Marshal(&src.mapvar)
	}

	if src.stringvar != nil {
		return json.Marshal(&src.stringvar)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransformDefinitionAttributesValue struct {
	value *TransformDefinitionAttributesValue
	isSet bool
}

func (v NullableTransformDefinitionAttributesValue) Get() *TransformDefinitionAttributesValue {
	return v.value
}

func (v *NullableTransformDefinitionAttributesValue) Set(val *TransformDefinitionAttributesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformDefinitionAttributesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformDefinitionAttributesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformDefinitionAttributesValue(val *TransformDefinitionAttributesValue) *NullableTransformDefinitionAttributesValue {
	return &NullableTransformDefinitionAttributesValue{value: val, isSet: true}
}

func (v NullableTransformDefinitionAttributesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformDefinitionAttributesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


