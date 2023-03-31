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

// checks if the RoleIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleIdentity{}

// RoleIdentity A subset of the fields of an Identity which is a member of a Role.
type RoleIdentity struct {
	// The ID of the Identity
	Id *string `json:"id,omitempty"`
	// The alias / username of the Identity
	AliasName *string `json:"aliasName,omitempty"`
	// The human-readable display name of the Identity
	Name *string `json:"name,omitempty"`
	// Email address of the Identity
	Email *string `json:"email,omitempty"`
	RoleAssignmentSource *RoleAssignmentSourceType `json:"roleAssignmentSource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleIdentity RoleIdentity

// NewRoleIdentity instantiates a new RoleIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleIdentity() *RoleIdentity {
	this := RoleIdentity{}
	return &this
}

// NewRoleIdentityWithDefaults instantiates a new RoleIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleIdentityWithDefaults() *RoleIdentity {
	this := RoleIdentity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleIdentity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleIdentity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleIdentity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleIdentity) SetId(v string) {
	o.Id = &v
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *RoleIdentity) GetAliasName() string {
	if o == nil || isNil(o.AliasName) {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleIdentity) GetAliasNameOk() (*string, bool) {
	if o == nil || isNil(o.AliasName) {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *RoleIdentity) HasAliasName() bool {
	if o != nil && !isNil(o.AliasName) {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *RoleIdentity) SetAliasName(v string) {
	o.AliasName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleIdentity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleIdentity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleIdentity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleIdentity) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RoleIdentity) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleIdentity) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RoleIdentity) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RoleIdentity) SetEmail(v string) {
	o.Email = &v
}

// GetRoleAssignmentSource returns the RoleAssignmentSource field value if set, zero value otherwise.
func (o *RoleIdentity) GetRoleAssignmentSource() RoleAssignmentSourceType {
	if o == nil || isNil(o.RoleAssignmentSource) {
		var ret RoleAssignmentSourceType
		return ret
	}
	return *o.RoleAssignmentSource
}

// GetRoleAssignmentSourceOk returns a tuple with the RoleAssignmentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleIdentity) GetRoleAssignmentSourceOk() (*RoleAssignmentSourceType, bool) {
	if o == nil || isNil(o.RoleAssignmentSource) {
		return nil, false
	}
	return o.RoleAssignmentSource, true
}

// HasRoleAssignmentSource returns a boolean if a field has been set.
func (o *RoleIdentity) HasRoleAssignmentSource() bool {
	if o != nil && !isNil(o.RoleAssignmentSource) {
		return true
	}

	return false
}

// SetRoleAssignmentSource gets a reference to the given RoleAssignmentSourceType and assigns it to the RoleAssignmentSource field.
func (o *RoleIdentity) SetRoleAssignmentSource(v RoleAssignmentSourceType) {
	o.RoleAssignmentSource = &v
}

func (o RoleIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AliasName) {
		toSerialize["aliasName"] = o.AliasName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.RoleAssignmentSource) {
		toSerialize["roleAssignmentSource"] = o.RoleAssignmentSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleIdentity) UnmarshalJSON(bytes []byte) (err error) {
	varRoleIdentity := _RoleIdentity{}

	if err = json.Unmarshal(bytes, &varRoleIdentity); err == nil {
		*o = RoleIdentity(varRoleIdentity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "aliasName")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "roleAssignmentSource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleIdentity struct {
	value *RoleIdentity
	isSet bool
}

func (v NullableRoleIdentity) Get() *RoleIdentity {
	return v.value
}

func (v *NullableRoleIdentity) Set(val *RoleIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleIdentity(val *RoleIdentity) *NullableRoleIdentity {
	return &NullableRoleIdentity{value: val, isSet: true}
}

func (v NullableRoleIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


