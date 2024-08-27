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

// LocaleOrigin An indicator of how the locale was selected. *DEFAULT* means the locale is the system default. *REQUEST* means the locale was selected from the request context (i.e., best match based on the *Accept-Language* header). Additional values may be added in the future without notice.
type LocaleOrigin string

// List of LocaleOrigin
const (
	LOCALEORIGIN_DEFAULT LocaleOrigin = "DEFAULT"
	LOCALEORIGIN_REQUEST LocaleOrigin = "REQUEST"
	LOCALEORIGIN_NULL LocaleOrigin = "null"
)

// All allowed values of LocaleOrigin enum
var AllowedLocaleOriginEnumValues = []LocaleOrigin{
	"DEFAULT",
	"REQUEST",
	"null",
}

func (v *LocaleOrigin) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocaleOrigin(value)
	for _, existing := range AllowedLocaleOriginEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocaleOrigin", value)
}

// NewLocaleOriginFromValue returns a pointer to a valid LocaleOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocaleOriginFromValue(v string) (*LocaleOrigin, error) {
	ev := LocaleOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocaleOrigin: valid values are %v", v, AllowedLocaleOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocaleOrigin) IsValid() bool {
	for _, existing := range AllowedLocaleOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocaleOrigin value
func (v LocaleOrigin) Ptr() *LocaleOrigin {
	return &v
}

type NullableLocaleOrigin struct {
	value *LocaleOrigin
	isSet bool
}

func (v NullableLocaleOrigin) Get() *LocaleOrigin {
	return v.value
}

func (v *NullableLocaleOrigin) Set(val *LocaleOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableLocaleOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableLocaleOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocaleOrigin(val *LocaleOrigin) *NullableLocaleOrigin {
	return &NullableLocaleOrigin{value: val, isSet: true}
}

func (v NullableLocaleOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocaleOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
