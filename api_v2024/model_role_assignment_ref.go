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

// checks if the RoleAssignmentRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAssignmentRef{}

// RoleAssignmentRef struct for RoleAssignmentRef
type RoleAssignmentRef struct {
	// Assignment Id
	Id *string `json:"id,omitempty"`
	Role *BaseReferenceDto1 `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleAssignmentRef RoleAssignmentRef

// NewRoleAssignmentRef instantiates a new RoleAssignmentRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentRef() *RoleAssignmentRef {
	this := RoleAssignmentRef{}
	return &this
}

// NewRoleAssignmentRefWithDefaults instantiates a new RoleAssignmentRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentRefWithDefaults() *RoleAssignmentRef {
	this := RoleAssignmentRef{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAssignmentRef) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRef) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAssignmentRef) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleAssignmentRef) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleAssignmentRef) GetRole() BaseReferenceDto1 {
	if o == nil || IsNil(o.Role) {
		var ret BaseReferenceDto1
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRef) GetRoleOk() (*BaseReferenceDto1, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleAssignmentRef) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given BaseReferenceDto1 and assigns it to the Role field.
func (o *RoleAssignmentRef) SetRole(v BaseReferenceDto1) {
	o.Role = &v
}

func (o RoleAssignmentRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAssignmentRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleAssignmentRef) UnmarshalJSON(data []byte) (err error) {
	varRoleAssignmentRef := _RoleAssignmentRef{}

	err = json.Unmarshal(data, &varRoleAssignmentRef)

	if err != nil {
		return err
	}

	*o = RoleAssignmentRef(varRoleAssignmentRef)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleAssignmentRef struct {
	value *RoleAssignmentRef
	isSet bool
}

func (v NullableRoleAssignmentRef) Get() *RoleAssignmentRef {
	return v.value
}

func (v *NullableRoleAssignmentRef) Set(val *RoleAssignmentRef) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentRef) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentRef(val *RoleAssignmentRef) *NullableRoleAssignmentRef {
	return &NullableRoleAssignmentRef{value: val, isSet: true}
}

func (v NullableRoleAssignmentRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


