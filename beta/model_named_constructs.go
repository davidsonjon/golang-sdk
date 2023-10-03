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

// NamedConstructs | Construct       | Date Time Pattern | Description | | ---------       | ----------------- | ----------- | | ISO8601         | `yyyy-MM-dd'T'HH:mm:ss.SSSX` | The ISO8601 standard. |           | LDAP            | `yyyyMMddHHmmss.Z`           | The LDAP standard.    | | PEOPLE_SOFT     | `MM/dd/yyyy`                 | The date format People Soft uses. | | EPOCH_TIME_JAVA | # ms from midnight, January 1st, 1970 | The incoming date value as elapsed time in milliseconds from midnight, January 1st, 1970. | | EPOCH_TIME_WIN32| # intervals of 100ns from midnight, January 1st, 1601 | The incoming date value as elapsed time in 100-nanosecond intervals from midnight, January 1st, 1601. | 
type NamedConstructs string

// List of namedConstructs
const (
	NAMEDCONSTRUCTS_ISO8601 NamedConstructs = "ISO8601"
	NAMEDCONSTRUCTS_LDAP NamedConstructs = "LDAP"
	NAMEDCONSTRUCTS_PEOPLE_SOFT NamedConstructs = "PEOPLE_SOFT"
	NAMEDCONSTRUCTS_EPOCH_TIME_JAVA NamedConstructs = "EPOCH_TIME_JAVA"
	NAMEDCONSTRUCTS_EPOCH_TIME_WIN32 NamedConstructs = "EPOCH_TIME_WIN32"
)

// All allowed values of NamedConstructs enum
var AllowedNamedConstructsEnumValues = []NamedConstructs{
	"ISO8601",
	"LDAP",
	"PEOPLE_SOFT",
	"EPOCH_TIME_JAVA",
	"EPOCH_TIME_WIN32",
}

func (v *NamedConstructs) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NamedConstructs(value)
	for _, existing := range AllowedNamedConstructsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NamedConstructs", value)
}

// NewNamedConstructsFromValue returns a pointer to a valid NamedConstructs
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNamedConstructsFromValue(v string) (*NamedConstructs, error) {
	ev := NamedConstructs(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NamedConstructs: valid values are %v", v, AllowedNamedConstructsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NamedConstructs) IsValid() bool {
	for _, existing := range AllowedNamedConstructsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to namedConstructs value
func (v NamedConstructs) Ptr() *NamedConstructs {
	return &v
}

type NullableNamedConstructs struct {
	value *NamedConstructs
	isSet bool
}

func (v NullableNamedConstructs) Get() *NamedConstructs {
	return v.value
}

func (v *NullableNamedConstructs) Set(val *NamedConstructs) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedConstructs) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedConstructs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedConstructs(val *NamedConstructs) *NullableNamedConstructs {
	return &NullableNamedConstructs{value: val, isSet: true}
}

func (v NullableNamedConstructs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedConstructs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

