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

// ListIdentityAccessItems200ResponseInner - struct for ListIdentityAccessItems200ResponseInner
type ListIdentityAccessItems200ResponseInner struct {
	AccessItemAccessProfileResponse *AccessItemAccessProfileResponse
	AccessItemAccountResponse *AccessItemAccountResponse
	AccessItemAppResponse *AccessItemAppResponse
	AccessItemEntitlementResponse *AccessItemEntitlementResponse
	AccessItemRoleResponse *AccessItemRoleResponse
}

// AccessItemAccessProfileResponseAsListIdentityAccessItems200ResponseInner is a convenience function that returns AccessItemAccessProfileResponse wrapped in ListIdentityAccessItems200ResponseInner
func AccessItemAccessProfileResponseAsListIdentityAccessItems200ResponseInner(v *AccessItemAccessProfileResponse) ListIdentityAccessItems200ResponseInner {
	return ListIdentityAccessItems200ResponseInner{
		AccessItemAccessProfileResponse: v,
	}
}

// AccessItemAccountResponseAsListIdentityAccessItems200ResponseInner is a convenience function that returns AccessItemAccountResponse wrapped in ListIdentityAccessItems200ResponseInner
func AccessItemAccountResponseAsListIdentityAccessItems200ResponseInner(v *AccessItemAccountResponse) ListIdentityAccessItems200ResponseInner {
	return ListIdentityAccessItems200ResponseInner{
		AccessItemAccountResponse: v,
	}
}

// AccessItemAppResponseAsListIdentityAccessItems200ResponseInner is a convenience function that returns AccessItemAppResponse wrapped in ListIdentityAccessItems200ResponseInner
func AccessItemAppResponseAsListIdentityAccessItems200ResponseInner(v *AccessItemAppResponse) ListIdentityAccessItems200ResponseInner {
	return ListIdentityAccessItems200ResponseInner{
		AccessItemAppResponse: v,
	}
}

// AccessItemEntitlementResponseAsListIdentityAccessItems200ResponseInner is a convenience function that returns AccessItemEntitlementResponse wrapped in ListIdentityAccessItems200ResponseInner
func AccessItemEntitlementResponseAsListIdentityAccessItems200ResponseInner(v *AccessItemEntitlementResponse) ListIdentityAccessItems200ResponseInner {
	return ListIdentityAccessItems200ResponseInner{
		AccessItemEntitlementResponse: v,
	}
}

// AccessItemRoleResponseAsListIdentityAccessItems200ResponseInner is a convenience function that returns AccessItemRoleResponse wrapped in ListIdentityAccessItems200ResponseInner
func AccessItemRoleResponseAsListIdentityAccessItems200ResponseInner(v *AccessItemRoleResponse) ListIdentityAccessItems200ResponseInner {
	return ListIdentityAccessItems200ResponseInner{
		AccessItemRoleResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListIdentityAccessItems200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessItemAccessProfileResponse
	err = newStrictDecoder(data).Decode(&dst.AccessItemAccessProfileResponse)
	if err == nil {
		jsonAccessItemAccessProfileResponse, _ := json.Marshal(dst.AccessItemAccessProfileResponse)
		if string(jsonAccessItemAccessProfileResponse) == "{}" { // empty struct
			dst.AccessItemAccessProfileResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessItemAccessProfileResponse = nil
	}

	// try to unmarshal data into AccessItemAccountResponse
	err = newStrictDecoder(data).Decode(&dst.AccessItemAccountResponse)
	if err == nil {
		jsonAccessItemAccountResponse, _ := json.Marshal(dst.AccessItemAccountResponse)
		if string(jsonAccessItemAccountResponse) == "{}" { // empty struct
			dst.AccessItemAccountResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessItemAccountResponse = nil
	}

	// try to unmarshal data into AccessItemAppResponse
	err = newStrictDecoder(data).Decode(&dst.AccessItemAppResponse)
	if err == nil {
		jsonAccessItemAppResponse, _ := json.Marshal(dst.AccessItemAppResponse)
		if string(jsonAccessItemAppResponse) == "{}" { // empty struct
			dst.AccessItemAppResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessItemAppResponse = nil
	}

	// try to unmarshal data into AccessItemEntitlementResponse
	err = newStrictDecoder(data).Decode(&dst.AccessItemEntitlementResponse)
	if err == nil {
		jsonAccessItemEntitlementResponse, _ := json.Marshal(dst.AccessItemEntitlementResponse)
		if string(jsonAccessItemEntitlementResponse) == "{}" { // empty struct
			dst.AccessItemEntitlementResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessItemEntitlementResponse = nil
	}

	// try to unmarshal data into AccessItemRoleResponse
	err = newStrictDecoder(data).Decode(&dst.AccessItemRoleResponse)
	if err == nil {
		jsonAccessItemRoleResponse, _ := json.Marshal(dst.AccessItemRoleResponse)
		if string(jsonAccessItemRoleResponse) == "{}" { // empty struct
			dst.AccessItemRoleResponse = nil
		} else {
			match++
		}
	} else {
		dst.AccessItemRoleResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessItemAccessProfileResponse = nil
		dst.AccessItemAccountResponse = nil
		dst.AccessItemAppResponse = nil
		dst.AccessItemEntitlementResponse = nil
		dst.AccessItemRoleResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListIdentityAccessItems200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListIdentityAccessItems200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListIdentityAccessItems200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AccessItemAccessProfileResponse != nil {
		return json.Marshal(&src.AccessItemAccessProfileResponse)
	}

	if src.AccessItemAccountResponse != nil {
		return json.Marshal(&src.AccessItemAccountResponse)
	}

	if src.AccessItemAppResponse != nil {
		return json.Marshal(&src.AccessItemAppResponse)
	}

	if src.AccessItemEntitlementResponse != nil {
		return json.Marshal(&src.AccessItemEntitlementResponse)
	}

	if src.AccessItemRoleResponse != nil {
		return json.Marshal(&src.AccessItemRoleResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListIdentityAccessItems200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessItemAccessProfileResponse != nil {
		return obj.AccessItemAccessProfileResponse
	}

	if obj.AccessItemAccountResponse != nil {
		return obj.AccessItemAccountResponse
	}

	if obj.AccessItemAppResponse != nil {
		return obj.AccessItemAppResponse
	}

	if obj.AccessItemEntitlementResponse != nil {
		return obj.AccessItemEntitlementResponse
	}

	if obj.AccessItemRoleResponse != nil {
		return obj.AccessItemRoleResponse
	}

	// all schemas are nil
	return nil
}

type NullableListIdentityAccessItems200ResponseInner struct {
	value *ListIdentityAccessItems200ResponseInner
	isSet bool
}

func (v NullableListIdentityAccessItems200ResponseInner) Get() *ListIdentityAccessItems200ResponseInner {
	return v.value
}

func (v *NullableListIdentityAccessItems200ResponseInner) Set(val *ListIdentityAccessItems200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListIdentityAccessItems200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListIdentityAccessItems200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIdentityAccessItems200ResponseInner(val *ListIdentityAccessItems200ResponseInner) *NullableListIdentityAccessItems200ResponseInner {
	return &NullableListIdentityAccessItems200ResponseInner{value: val, isSet: true}
}

func (v NullableListIdentityAccessItems200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIdentityAccessItems200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


