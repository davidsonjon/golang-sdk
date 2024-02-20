/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// RequestableObjectType The currently supported requestable object types. 
type RequestableObjectType string

// List of RequestableObjectType
const (
	REQUESTABLEOBJECTTYPE_ACCESS_PROFILE RequestableObjectType = "ACCESS_PROFILE"
	REQUESTABLEOBJECTTYPE_ROLE RequestableObjectType = "ROLE"
	REQUESTABLEOBJECTTYPE_ENTITLEMENT RequestableObjectType = "ENTITLEMENT"
)

// All allowed values of RequestableObjectType enum
var AllowedRequestableObjectTypeEnumValues = []RequestableObjectType{
	"ACCESS_PROFILE",
	"ROLE",
	"ENTITLEMENT",
}

func (v *RequestableObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestableObjectType(value)
	for _, existing := range AllowedRequestableObjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestableObjectType", value)
}

// NewRequestableObjectTypeFromValue returns a pointer to a valid RequestableObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestableObjectTypeFromValue(v string) (*RequestableObjectType, error) {
	ev := RequestableObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestableObjectType: valid values are %v", v, AllowedRequestableObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestableObjectType) IsValid() bool {
	for _, existing := range AllowedRequestableObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestableObjectType value
func (v RequestableObjectType) Ptr() *RequestableObjectType {
	return &v
}

type NullableRequestableObjectType struct {
	value *RequestableObjectType
	isSet bool
}

func (v NullableRequestableObjectType) Get() *RequestableObjectType {
	return v.value
}

func (v *NullableRequestableObjectType) Set(val *RequestableObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestableObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestableObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestableObjectType(val *RequestableObjectType) *NullableRequestableObjectType {
	return &NullableRequestableObjectType{value: val, isSet: true}
}

func (v NullableRequestableObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestableObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

