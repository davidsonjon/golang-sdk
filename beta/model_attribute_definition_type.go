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

// AttributeDefinitionType The underlying type of the value which an AttributeDefinition represents.
type AttributeDefinitionType string

// List of AttributeDefinitionType
const (
	ATTRIBUTEDEFINITIONTYPE_STRING AttributeDefinitionType = "STRING"
	ATTRIBUTEDEFINITIONTYPE_LONG AttributeDefinitionType = "LONG"
	ATTRIBUTEDEFINITIONTYPE_INT AttributeDefinitionType = "INT"
	ATTRIBUTEDEFINITIONTYPE_BOOLEAN AttributeDefinitionType = "BOOLEAN"
)

// All allowed values of AttributeDefinitionType enum
var AllowedAttributeDefinitionTypeEnumValues = []AttributeDefinitionType{
	"STRING",
	"LONG",
	"INT",
	"BOOLEAN",
}

func (v *AttributeDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttributeDefinitionType(value)
	for _, existing := range AllowedAttributeDefinitionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttributeDefinitionType", value)
}

// NewAttributeDefinitionTypeFromValue returns a pointer to a valid AttributeDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAttributeDefinitionTypeFromValue(v string) (*AttributeDefinitionType, error) {
	ev := AttributeDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AttributeDefinitionType: valid values are %v", v, AllowedAttributeDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AttributeDefinitionType) IsValid() bool {
	for _, existing := range AllowedAttributeDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttributeDefinitionType value
func (v AttributeDefinitionType) Ptr() *AttributeDefinitionType {
	return &v
}

type NullableAttributeDefinitionType struct {
	value *AttributeDefinitionType
	isSet bool
}

func (v NullableAttributeDefinitionType) Get() *AttributeDefinitionType {
	return v.value
}

func (v *NullableAttributeDefinitionType) Set(val *AttributeDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeDefinitionType(val *AttributeDefinitionType) *NullableAttributeDefinitionType {
	return &NullableAttributeDefinitionType{value: val, isSet: true}
}

func (v NullableAttributeDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

