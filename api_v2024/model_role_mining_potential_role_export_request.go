/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the RoleMiningPotentialRoleExportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRoleExportRequest{}

// RoleMiningPotentialRoleExportRequest struct for RoleMiningPotentialRoleExportRequest
type RoleMiningPotentialRoleExportRequest struct {
	// The minimum popularity among identities in the role which an entitlement must have to be included in the report
	MinEntitlementPopularity *int32 `json:"minEntitlementPopularity,omitempty"`
	// If false, do not include entitlements that are highly popular among the entire orginization
	IncludeCommonAccess *bool `json:"includeCommonAccess,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleExportRequest RoleMiningPotentialRoleExportRequest

// NewRoleMiningPotentialRoleExportRequest instantiates a new RoleMiningPotentialRoleExportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleExportRequest() *RoleMiningPotentialRoleExportRequest {
	this := RoleMiningPotentialRoleExportRequest{}
	return &this
}

// NewRoleMiningPotentialRoleExportRequestWithDefaults instantiates a new RoleMiningPotentialRoleExportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleExportRequestWithDefaults() *RoleMiningPotentialRoleExportRequest {
	this := RoleMiningPotentialRoleExportRequest{}
	return &this
}

// GetMinEntitlementPopularity returns the MinEntitlementPopularity field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleExportRequest) GetMinEntitlementPopularity() int32 {
	if o == nil || isNil(o.MinEntitlementPopularity) {
		var ret int32
		return ret
	}
	return *o.MinEntitlementPopularity
}

// GetMinEntitlementPopularityOk returns a tuple with the MinEntitlementPopularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleExportRequest) GetMinEntitlementPopularityOk() (*int32, bool) {
	if o == nil || isNil(o.MinEntitlementPopularity) {
		return nil, false
	}
	return o.MinEntitlementPopularity, true
}

// HasMinEntitlementPopularity returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleExportRequest) HasMinEntitlementPopularity() bool {
	if o != nil && !isNil(o.MinEntitlementPopularity) {
		return true
	}

	return false
}

// SetMinEntitlementPopularity gets a reference to the given int32 and assigns it to the MinEntitlementPopularity field.
func (o *RoleMiningPotentialRoleExportRequest) SetMinEntitlementPopularity(v int32) {
	o.MinEntitlementPopularity = &v
}

// GetIncludeCommonAccess returns the IncludeCommonAccess field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleExportRequest) GetIncludeCommonAccess() bool {
	if o == nil || isNil(o.IncludeCommonAccess) {
		var ret bool
		return ret
	}
	return *o.IncludeCommonAccess
}

// GetIncludeCommonAccessOk returns a tuple with the IncludeCommonAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleExportRequest) GetIncludeCommonAccessOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeCommonAccess) {
		return nil, false
	}
	return o.IncludeCommonAccess, true
}

// HasIncludeCommonAccess returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleExportRequest) HasIncludeCommonAccess() bool {
	if o != nil && !isNil(o.IncludeCommonAccess) {
		return true
	}

	return false
}

// SetIncludeCommonAccess gets a reference to the given bool and assigns it to the IncludeCommonAccess field.
func (o *RoleMiningPotentialRoleExportRequest) SetIncludeCommonAccess(v bool) {
	o.IncludeCommonAccess = &v
}

func (o RoleMiningPotentialRoleExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRoleExportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinEntitlementPopularity) {
		toSerialize["minEntitlementPopularity"] = o.MinEntitlementPopularity
	}
	if !isNil(o.IncludeCommonAccess) {
		toSerialize["includeCommonAccess"] = o.IncludeCommonAccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRoleExportRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRoleExportRequest := _RoleMiningPotentialRoleExportRequest{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRoleExportRequest); err == nil {
	*o = RoleMiningPotentialRoleExportRequest(varRoleMiningPotentialRoleExportRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "minEntitlementPopularity")
		delete(additionalProperties, "includeCommonAccess")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleExportRequest struct {
	value *RoleMiningPotentialRoleExportRequest
	isSet bool
}

func (v NullableRoleMiningPotentialRoleExportRequest) Get() *RoleMiningPotentialRoleExportRequest {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleExportRequest) Set(val *RoleMiningPotentialRoleExportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleExportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleExportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleExportRequest(val *RoleMiningPotentialRoleExportRequest) *NullableRoleMiningPotentialRoleExportRequest {
	return &NullableRoleMiningPotentialRoleExportRequest{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleExportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleExportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


