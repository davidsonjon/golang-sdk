/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GetDiscoveredApplications200ResponseInner - struct for GetDiscoveredApplications200ResponseInner
type GetDiscoveredApplications200ResponseInner struct {
	FullDiscoveredApplications *FullDiscoveredApplications
	SlimDiscoveredApplications *SlimDiscoveredApplications
}

// FullDiscoveredApplicationsAsGetDiscoveredApplications200ResponseInner is a convenience function that returns FullDiscoveredApplications wrapped in GetDiscoveredApplications200ResponseInner
func FullDiscoveredApplicationsAsGetDiscoveredApplications200ResponseInner(v *FullDiscoveredApplications) GetDiscoveredApplications200ResponseInner {
	return GetDiscoveredApplications200ResponseInner{
		FullDiscoveredApplications: v,
	}
}

// SlimDiscoveredApplicationsAsGetDiscoveredApplications200ResponseInner is a convenience function that returns SlimDiscoveredApplications wrapped in GetDiscoveredApplications200ResponseInner
func SlimDiscoveredApplicationsAsGetDiscoveredApplications200ResponseInner(v *SlimDiscoveredApplications) GetDiscoveredApplications200ResponseInner {
	return GetDiscoveredApplications200ResponseInner{
		SlimDiscoveredApplications: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDiscoveredApplications200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FullDiscoveredApplications
	err = newStrictDecoder(data).Decode(&dst.FullDiscoveredApplications)
	if err == nil {
		jsonFullDiscoveredApplications, _ := json.Marshal(dst.FullDiscoveredApplications)
		if string(jsonFullDiscoveredApplications) == "{}" { // empty struct
			dst.FullDiscoveredApplications = nil
		} else {
			if err = validator.Validate(dst.FullDiscoveredApplications); err != nil {
				dst.FullDiscoveredApplications = nil
			} else {
				match++
			}
		}
	} else {
		dst.FullDiscoveredApplications = nil
	}

	// try to unmarshal data into SlimDiscoveredApplications
	err = newStrictDecoder(data).Decode(&dst.SlimDiscoveredApplications)
	if err == nil {
		jsonSlimDiscoveredApplications, _ := json.Marshal(dst.SlimDiscoveredApplications)
		if string(jsonSlimDiscoveredApplications) == "{}" { // empty struct
			dst.SlimDiscoveredApplications = nil
		} else {
			if err = validator.Validate(dst.SlimDiscoveredApplications); err != nil {
				dst.SlimDiscoveredApplications = nil
			} else {
				match++
			}
		}
	} else {
		dst.SlimDiscoveredApplications = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FullDiscoveredApplications = nil
		dst.SlimDiscoveredApplications = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDiscoveredApplications200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDiscoveredApplications200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDiscoveredApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.FullDiscoveredApplications != nil {
		return json.Marshal(&src.FullDiscoveredApplications)
	}

	if src.SlimDiscoveredApplications != nil {
		return json.Marshal(&src.SlimDiscoveredApplications)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDiscoveredApplications200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FullDiscoveredApplications != nil {
		return obj.FullDiscoveredApplications
	}

	if obj.SlimDiscoveredApplications != nil {
		return obj.SlimDiscoveredApplications
	}

	// all schemas are nil
	return nil
}

type NullableGetDiscoveredApplications200ResponseInner struct {
	value *GetDiscoveredApplications200ResponseInner
	isSet bool
}

func (v NullableGetDiscoveredApplications200ResponseInner) Get() *GetDiscoveredApplications200ResponseInner {
	return v.value
}

func (v *NullableGetDiscoveredApplications200ResponseInner) Set(val *GetDiscoveredApplications200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoveredApplications200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoveredApplications200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoveredApplications200ResponseInner(val *GetDiscoveredApplications200ResponseInner) *NullableGetDiscoveredApplications200ResponseInner {
	return &NullableGetDiscoveredApplications200ResponseInner{value: val, isSet: true}
}

func (v NullableGetDiscoveredApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoveredApplications200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


