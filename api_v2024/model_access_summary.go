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

// checks if the AccessSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessSummary{}

// AccessSummary An object holding the access that is being reviewed
type AccessSummary struct {
	Access *AccessSummaryAccess `json:"access,omitempty"`
	Entitlement NullableReviewableEntitlement `json:"entitlement,omitempty"`
	AccessProfile *ReviewableAccessProfile `json:"accessProfile,omitempty"`
	Role NullableReviewableRole `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessSummary AccessSummary

// NewAccessSummary instantiates a new AccessSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessSummary() *AccessSummary {
	this := AccessSummary{}
	return &this
}

// NewAccessSummaryWithDefaults instantiates a new AccessSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessSummaryWithDefaults() *AccessSummary {
	this := AccessSummary{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AccessSummary) GetAccess() AccessSummaryAccess {
	if o == nil || isNil(o.Access) {
		var ret AccessSummaryAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSummary) GetAccessOk() (*AccessSummaryAccess, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AccessSummary) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AccessSummaryAccess and assigns it to the Access field.
func (o *AccessSummary) SetAccess(v AccessSummaryAccess) {
	o.Access = &v
}

// GetEntitlement returns the Entitlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessSummary) GetEntitlement() ReviewableEntitlement {
	if o == nil || isNil(o.Entitlement.Get()) {
		var ret ReviewableEntitlement
		return ret
	}
	return *o.Entitlement.Get()
}

// GetEntitlementOk returns a tuple with the Entitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessSummary) GetEntitlementOk() (*ReviewableEntitlement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlement.Get(), o.Entitlement.IsSet()
}

// HasEntitlement returns a boolean if a field has been set.
func (o *AccessSummary) HasEntitlement() bool {
	if o != nil && o.Entitlement.IsSet() {
		return true
	}

	return false
}

// SetEntitlement gets a reference to the given NullableReviewableEntitlement and assigns it to the Entitlement field.
func (o *AccessSummary) SetEntitlement(v ReviewableEntitlement) {
	o.Entitlement.Set(&v)
}
// SetEntitlementNil sets the value for Entitlement to be an explicit nil
func (o *AccessSummary) SetEntitlementNil() {
	o.Entitlement.Set(nil)
}

// UnsetEntitlement ensures that no value is present for Entitlement, not even an explicit nil
func (o *AccessSummary) UnsetEntitlement() {
	o.Entitlement.Unset()
}

// GetAccessProfile returns the AccessProfile field value if set, zero value otherwise.
func (o *AccessSummary) GetAccessProfile() ReviewableAccessProfile {
	if o == nil || isNil(o.AccessProfile) {
		var ret ReviewableAccessProfile
		return ret
	}
	return *o.AccessProfile
}

// GetAccessProfileOk returns a tuple with the AccessProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSummary) GetAccessProfileOk() (*ReviewableAccessProfile, bool) {
	if o == nil || isNil(o.AccessProfile) {
		return nil, false
	}
	return o.AccessProfile, true
}

// HasAccessProfile returns a boolean if a field has been set.
func (o *AccessSummary) HasAccessProfile() bool {
	if o != nil && !isNil(o.AccessProfile) {
		return true
	}

	return false
}

// SetAccessProfile gets a reference to the given ReviewableAccessProfile and assigns it to the AccessProfile field.
func (o *AccessSummary) SetAccessProfile(v ReviewableAccessProfile) {
	o.AccessProfile = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessSummary) GetRole() ReviewableRole {
	if o == nil || isNil(o.Role.Get()) {
		var ret ReviewableRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessSummary) GetRoleOk() (*ReviewableRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *AccessSummary) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableReviewableRole and assigns it to the Role field.
func (o *AccessSummary) SetRole(v ReviewableRole) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *AccessSummary) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *AccessSummary) UnsetRole() {
	o.Role.Unset()
}

func (o AccessSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if o.Entitlement.IsSet() {
		toSerialize["entitlement"] = o.Entitlement.Get()
	}
	if !isNil(o.AccessProfile) {
		toSerialize["accessProfile"] = o.AccessProfile
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessSummary) UnmarshalJSON(bytes []byte) (err error) {
	varAccessSummary := _AccessSummary{}

	if err = json.Unmarshal(bytes, &varAccessSummary); err == nil {
	*o = AccessSummary(varAccessSummary)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "entitlement")
		delete(additionalProperties, "accessProfile")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessSummary struct {
	value *AccessSummary
	isSet bool
}

func (v NullableAccessSummary) Get() *AccessSummary {
	return v.value
}

func (v *NullableAccessSummary) Set(val *AccessSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessSummary(val *AccessSummary) *NullableAccessSummary {
	return &NullableAccessSummary{value: val, isSet: true}
}

func (v NullableAccessSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


