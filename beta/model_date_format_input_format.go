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

// DateFormatInputFormat - A string value indicating either the explicit SimpleDateFormat or the built-in named format that the data is coming in as.  *If no inputFormat is provided, the transform assumes that it is in ISO8601 format*
type DateFormatInputFormat struct {
	NamedConstructs *NamedConstructs
	String *string
}

// NamedConstructsAsDateFormatInputFormat is a convenience function that returns NamedConstructs wrapped in DateFormatInputFormat
func NamedConstructsAsDateFormatInputFormat(v *NamedConstructs) DateFormatInputFormat {
	return DateFormatInputFormat{
		NamedConstructs: v,
	}
}

// stringAsDateFormatInputFormat is a convenience function that returns string wrapped in DateFormatInputFormat
func StringAsDateFormatInputFormat(v *string) DateFormatInputFormat {
	return DateFormatInputFormat{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DateFormatInputFormat) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NamedConstructs
	err = newStrictDecoder(data).Decode(&dst.NamedConstructs)
	if err == nil {
		jsonNamedConstructs, _ := json.Marshal(dst.NamedConstructs)
		if string(jsonNamedConstructs) == "{}" { // empty struct
			dst.NamedConstructs = nil
		} else {
			match++
		}
	} else {
		dst.NamedConstructs = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NamedConstructs = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DateFormatInputFormat)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DateFormatInputFormat)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DateFormatInputFormat) MarshalJSON() ([]byte, error) {
	if src.NamedConstructs != nil {
		return json.Marshal(&src.NamedConstructs)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DateFormatInputFormat) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NamedConstructs != nil {
		return obj.NamedConstructs
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDateFormatInputFormat struct {
	value *DateFormatInputFormat
	isSet bool
}

func (v NullableDateFormatInputFormat) Get() *DateFormatInputFormat {
	return v.value
}

func (v *NullableDateFormatInputFormat) Set(val *DateFormatInputFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableDateFormatInputFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableDateFormatInputFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateFormatInputFormat(val *DateFormatInputFormat) *NullableDateFormatInputFormat {
	return &NullableDateFormatInputFormat{value: val, isSet: true}
}

func (v NullableDateFormatInputFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateFormatInputFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


