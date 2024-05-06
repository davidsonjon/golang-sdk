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

// AccessRequestRecommendationItemType The type of access item.
type AccessRequestRecommendationItemType string

// List of AccessRequestRecommendationItemType
const (
	ACCESSREQUESTRECOMMENDATIONITEMTYPE_ACCESS_PROFILE AccessRequestRecommendationItemType = "ACCESS_PROFILE"
	ACCESSREQUESTRECOMMENDATIONITEMTYPE_ROLE AccessRequestRecommendationItemType = "ROLE"
)

// All allowed values of AccessRequestRecommendationItemType enum
var AllowedAccessRequestRecommendationItemTypeEnumValues = []AccessRequestRecommendationItemType{
	"ACCESS_PROFILE",
	"ROLE",
}

func (v *AccessRequestRecommendationItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessRequestRecommendationItemType(value)
	for _, existing := range AllowedAccessRequestRecommendationItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessRequestRecommendationItemType", value)
}

// NewAccessRequestRecommendationItemTypeFromValue returns a pointer to a valid AccessRequestRecommendationItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessRequestRecommendationItemTypeFromValue(v string) (*AccessRequestRecommendationItemType, error) {
	ev := AccessRequestRecommendationItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessRequestRecommendationItemType: valid values are %v", v, AllowedAccessRequestRecommendationItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessRequestRecommendationItemType) IsValid() bool {
	for _, existing := range AllowedAccessRequestRecommendationItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessRequestRecommendationItemType value
func (v AccessRequestRecommendationItemType) Ptr() *AccessRequestRecommendationItemType {
	return &v
}

type NullableAccessRequestRecommendationItemType struct {
	value *AccessRequestRecommendationItemType
	isSet bool
}

func (v NullableAccessRequestRecommendationItemType) Get() *AccessRequestRecommendationItemType {
	return v.value
}

func (v *NullableAccessRequestRecommendationItemType) Set(val *AccessRequestRecommendationItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestRecommendationItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestRecommendationItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestRecommendationItemType(val *AccessRequestRecommendationItemType) *NullableAccessRequestRecommendationItemType {
	return &NullableAccessRequestRecommendationItemType{value: val, isSet: true}
}

func (v NullableAccessRequestRecommendationItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestRecommendationItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

