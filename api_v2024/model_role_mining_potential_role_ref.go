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

// checks if the RoleMiningPotentialRoleRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRoleRef{}

// RoleMiningPotentialRoleRef struct for RoleMiningPotentialRoleRef
type RoleMiningPotentialRoleRef struct {
	// Id of the potential role
	Id *string `json:"id,omitempty"`
	// Name of the potential role
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleRef RoleMiningPotentialRoleRef

// NewRoleMiningPotentialRoleRef instantiates a new RoleMiningPotentialRoleRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleRef() *RoleMiningPotentialRoleRef {
	this := RoleMiningPotentialRoleRef{}
	return &this
}

// NewRoleMiningPotentialRoleRefWithDefaults instantiates a new RoleMiningPotentialRoleRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleRefWithDefaults() *RoleMiningPotentialRoleRef {
	this := RoleMiningPotentialRoleRef{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleRef) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleRef) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleRef) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMiningPotentialRoleRef) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleRef) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleRef) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleRef) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleMiningPotentialRoleRef) SetName(v string) {
	o.Name = &v
}

func (o RoleMiningPotentialRoleRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRoleRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRoleRef) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRoleRef := _RoleMiningPotentialRoleRef{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRoleRef); err == nil {
	*o = RoleMiningPotentialRoleRef(varRoleMiningPotentialRoleRef)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleRef struct {
	value *RoleMiningPotentialRoleRef
	isSet bool
}

func (v NullableRoleMiningPotentialRoleRef) Get() *RoleMiningPotentialRoleRef {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleRef) Set(val *RoleMiningPotentialRoleRef) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleRef) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleRef(val *RoleMiningPotentialRoleRef) *NullableRoleMiningPotentialRoleRef {
	return &NullableRoleMiningPotentialRoleRef{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

