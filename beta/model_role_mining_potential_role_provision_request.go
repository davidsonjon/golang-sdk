/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the RoleMiningPotentialRoleProvisionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRoleProvisionRequest{}

// RoleMiningPotentialRoleProvisionRequest struct for RoleMiningPotentialRoleProvisionRequest
type RoleMiningPotentialRoleProvisionRequest struct {
	// Name of the new role being created
	RoleName *string `json:"roleName,omitempty"`
	// Short description of the new role being created
	RoleDescription *string `json:"roleDescription,omitempty"`
	// ID of the identity that will own this role
	OwnerId *string `json:"ownerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleProvisionRequest RoleMiningPotentialRoleProvisionRequest

// NewRoleMiningPotentialRoleProvisionRequest instantiates a new RoleMiningPotentialRoleProvisionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleProvisionRequest() *RoleMiningPotentialRoleProvisionRequest {
	this := RoleMiningPotentialRoleProvisionRequest{}
	return &this
}

// NewRoleMiningPotentialRoleProvisionRequestWithDefaults instantiates a new RoleMiningPotentialRoleProvisionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleProvisionRequestWithDefaults() *RoleMiningPotentialRoleProvisionRequest {
	this := RoleMiningPotentialRoleProvisionRequest{}
	return &this
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleName() string {
	if o == nil || isNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || isNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasRoleName() bool {
	if o != nil && !isNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetRoleName(v string) {
	o.RoleName = &v
}

// GetRoleDescription returns the RoleDescription field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleDescription() string {
	if o == nil || isNil(o.RoleDescription) {
		var ret string
		return ret
	}
	return *o.RoleDescription
}

// GetRoleDescriptionOk returns a tuple with the RoleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.RoleDescription) {
		return nil, false
	}
	return o.RoleDescription, true
}

// HasRoleDescription returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasRoleDescription() bool {
	if o != nil && !isNil(o.RoleDescription) {
		return true
	}

	return false
}

// SetRoleDescription gets a reference to the given string and assigns it to the RoleDescription field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetRoleDescription(v string) {
	o.RoleDescription = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o RoleMiningPotentialRoleProvisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRoleProvisionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	if !isNil(o.RoleDescription) {
		toSerialize["roleDescription"] = o.RoleDescription
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRoleProvisionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRoleProvisionRequest := _RoleMiningPotentialRoleProvisionRequest{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRoleProvisionRequest); err == nil {
		*o = RoleMiningPotentialRoleProvisionRequest(varRoleMiningPotentialRoleProvisionRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "roleName")
		delete(additionalProperties, "roleDescription")
		delete(additionalProperties, "ownerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleProvisionRequest struct {
	value *RoleMiningPotentialRoleProvisionRequest
	isSet bool
}

func (v NullableRoleMiningPotentialRoleProvisionRequest) Get() *RoleMiningPotentialRoleProvisionRequest {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleProvisionRequest) Set(val *RoleMiningPotentialRoleProvisionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleProvisionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleProvisionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleProvisionRequest(val *RoleMiningPotentialRoleProvisionRequest) *NullableRoleMiningPotentialRoleProvisionRequest {
	return &NullableRoleMiningPotentialRoleProvisionRequest{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleProvisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleProvisionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


