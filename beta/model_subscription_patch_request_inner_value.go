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

// SubscriptionPatchRequestInnerValue The value to be used for the operation, required for \"add\" and \"replace\" operations
type SubscriptionPatchRequestInnerValue struct {
	SubscriptionPatchRequestInnerValueAnyOfInnervar *[]SubscriptionPatchRequestInnerValueAnyOfInner
	int32var *int32
	mapvar *map[string]interface{}
	stringvar *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SubscriptionPatchRequestInnerValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SubscriptionPatchRequestInnerValueAnyOfInnervar
	err = json.Unmarshal(data, &dst.SubscriptionPatchRequestInnerValueAnyOfInnervar);
	if err == nil {
		jsonSubscriptionPatchRequestInnerValueAnyOfInnervar, _ := json.Marshal(dst.SubscriptionPatchRequestInnerValueAnyOfInnervar)
		if string(jsonSubscriptionPatchRequestInnerValueAnyOfInnervar) == "{}" { // empty struct
			dst.SubscriptionPatchRequestInnerValueAnyOfInnervar = nil
		} else {
			return nil // data stored in dst.SubscriptionPatchRequestInnerValueAnyOfInnervar, return on the first match
		}
	} else {
		dst.SubscriptionPatchRequestInnerValueAnyOfInnervar = nil
	}

	// try to unmarshal JSON data into int32var
	err = json.Unmarshal(data, &dst.int32var);
	if err == nil {
		jsonint32var, _ := json.Marshal(dst.int32var)
		if string(jsonint32var) == "{}" { // empty struct
			dst.int32var = nil
		} else {
			return nil // data stored in dst.int32var, return on the first match
		}
	} else {
		dst.int32var = nil
	}

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

	return fmt.Errorf("data failed to match schemas in anyOf(SubscriptionPatchRequestInnerValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SubscriptionPatchRequestInnerValue) MarshalJSON() ([]byte, error) {
	if src.SubscriptionPatchRequestInnerValueAnyOfInnervar != nil {
		return json.Marshal(&src.SubscriptionPatchRequestInnerValueAnyOfInnervar)
	}

	if src.int32var != nil {
		return json.Marshal(&src.int32var)
	}

	if src.mapvar != nil {
		return json.Marshal(&src.mapvar)
	}

	if src.stringvar != nil {
		return json.Marshal(&src.stringvar)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubscriptionPatchRequestInnerValue struct {
	value *SubscriptionPatchRequestInnerValue
	isSet bool
}

func (v NullableSubscriptionPatchRequestInnerValue) Get() *SubscriptionPatchRequestInnerValue {
	return v.value
}

func (v *NullableSubscriptionPatchRequestInnerValue) Set(val *SubscriptionPatchRequestInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPatchRequestInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPatchRequestInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPatchRequestInnerValue(val *SubscriptionPatchRequestInnerValue) *NullableSubscriptionPatchRequestInnerValue {
	return &NullableSubscriptionPatchRequestInnerValue{value: val, isSet: true}
}

func (v NullableSubscriptionPatchRequestInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPatchRequestInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


