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

// OutlierContributingFeatureValue - The feature value
type OutlierContributingFeatureValue struct {
	Float32 *float32
	Int64 *int64
}

// float32AsOutlierContributingFeatureValue is a convenience function that returns float32 wrapped in OutlierContributingFeatureValue
func Float32AsOutlierContributingFeatureValue(v *float32) OutlierContributingFeatureValue {
	return OutlierContributingFeatureValue{
		Float32: v,
	}
}

// int64AsOutlierContributingFeatureValue is a convenience function that returns int64 wrapped in OutlierContributingFeatureValue
func Int64AsOutlierContributingFeatureValue(v *int64) OutlierContributingFeatureValue {
	return OutlierContributingFeatureValue{
		Int64: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OutlierContributingFeatureValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			match++
		}
	} else {
		dst.Int64 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.Int64 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OutlierContributingFeatureValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OutlierContributingFeatureValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OutlierContributingFeatureValue) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OutlierContributingFeatureValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.Int64 != nil {
		return obj.Int64
	}

	// all schemas are nil
	return nil
}

type NullableOutlierContributingFeatureValue struct {
	value *OutlierContributingFeatureValue
	isSet bool
}

func (v NullableOutlierContributingFeatureValue) Get() *OutlierContributingFeatureValue {
	return v.value
}

func (v *NullableOutlierContributingFeatureValue) Set(val *OutlierContributingFeatureValue) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlierContributingFeatureValue) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlierContributingFeatureValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlierContributingFeatureValue(val *OutlierContributingFeatureValue) *NullableOutlierContributingFeatureValue {
	return &NullableOutlierContributingFeatureValue{value: val, isSet: true}
}

func (v NullableOutlierContributingFeatureValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlierContributingFeatureValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


