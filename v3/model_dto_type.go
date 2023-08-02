/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// DtoType An enumeration of the types of DTOs supported within the IdentityNow infrastructure.
type DtoType string

// List of DtoType
const (
	DTOTYPE_ACCOUNT_CORRELATION_CONFIG DtoType = "ACCOUNT_CORRELATION_CONFIG"
	DTOTYPE_ACCESS_PROFILE DtoType = "ACCESS_PROFILE"
	DTOTYPE_ACCESS_REQUEST_APPROVAL DtoType = "ACCESS_REQUEST_APPROVAL"
	DTOTYPE_ACCOUNT DtoType = "ACCOUNT"
	DTOTYPE_APPLICATION DtoType = "APPLICATION"
	DTOTYPE_CAMPAIGN DtoType = "CAMPAIGN"
	DTOTYPE_CAMPAIGN_FILTER DtoType = "CAMPAIGN_FILTER"
	DTOTYPE_CERTIFICATION DtoType = "CERTIFICATION"
	DTOTYPE_CLUSTER DtoType = "CLUSTER"
	DTOTYPE_CONNECTOR_SCHEMA DtoType = "CONNECTOR_SCHEMA"
	DTOTYPE_ENTITLEMENT DtoType = "ENTITLEMENT"
	DTOTYPE_GOVERNANCE_GROUP DtoType = "GOVERNANCE_GROUP"
	DTOTYPE_IDENTITY DtoType = "IDENTITY"
	DTOTYPE_IDENTITY_PROFILE DtoType = "IDENTITY_PROFILE"
	DTOTYPE_IDENTITY_REQUEST DtoType = "IDENTITY_REQUEST"
	DTOTYPE_LIFECYCLE_STATE DtoType = "LIFECYCLE_STATE"
	DTOTYPE_PASSWORD_POLICY DtoType = "PASSWORD_POLICY"
	DTOTYPE_ROLE DtoType = "ROLE"
	DTOTYPE_RULE DtoType = "RULE"
	DTOTYPE_SOD_POLICY DtoType = "SOD_POLICY"
	DTOTYPE_SOURCE DtoType = "SOURCE"
	DTOTYPE_TAG DtoType = "TAG"
	DTOTYPE_TAG_CATEGORY DtoType = "TAG_CATEGORY"
	DTOTYPE_TASK_RESULT DtoType = "TASK_RESULT"
	DTOTYPE_REPORT_RESULT DtoType = "REPORT_RESULT"
	DTOTYPE_SOD_VIOLATION DtoType = "SOD_VIOLATION"
	DTOTYPE_ACCOUNT_ACTIVITY DtoType = "ACCOUNT_ACTIVITY"
	DTOTYPE_WORKGROUP DtoType = "WORKGROUP"
)

// All allowed values of DtoType enum
var AllowedDtoTypeEnumValues = []DtoType{
	"ACCOUNT_CORRELATION_CONFIG",
	"ACCESS_PROFILE",
	"ACCESS_REQUEST_APPROVAL",
	"ACCOUNT",
	"APPLICATION",
	"CAMPAIGN",
	"CAMPAIGN_FILTER",
	"CERTIFICATION",
	"CLUSTER",
	"CONNECTOR_SCHEMA",
	"ENTITLEMENT",
	"GOVERNANCE_GROUP",
	"IDENTITY",
	"IDENTITY_PROFILE",
	"IDENTITY_REQUEST",
	"LIFECYCLE_STATE",
	"PASSWORD_POLICY",
	"ROLE",
	"RULE",
	"SOD_POLICY",
	"SOURCE",
	"TAG",
	"TAG_CATEGORY",
	"TASK_RESULT",
	"REPORT_RESULT",
	"SOD_VIOLATION",
	"ACCOUNT_ACTIVITY",
	"WORKGROUP",
}

func (v *DtoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DtoType(value)
	for _, existing := range AllowedDtoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DtoType", value)
}

// NewDtoTypeFromValue returns a pointer to a valid DtoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDtoTypeFromValue(v string) (*DtoType, error) {
	ev := DtoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DtoType: valid values are %v", v, AllowedDtoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DtoType) IsValid() bool {
	for _, existing := range AllowedDtoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DtoType value
func (v DtoType) Ptr() *DtoType {
	return &v
}

type NullableDtoType struct {
	value *DtoType
	isSet bool
}

func (v NullableDtoType) Get() *DtoType {
	return v.value
}

func (v *NullableDtoType) Set(val *DtoType) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoType) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoType(val *DtoType) *NullableDtoType {
	return &NullableDtoType{value: val, isSet: true}
}

func (v NullableDtoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

