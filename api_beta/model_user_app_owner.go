/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the UserAppOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAppOwner{}

// UserAppOwner struct for UserAppOwner
type UserAppOwner struct {
	// The identity ID
	Id *string `json:"id,omitempty"`
	// It will always be \"IDENTITY\"
	Type *string `json:"type,omitempty"`
	// The identity name
	Name *string `json:"name,omitempty"`
	// The identity alias
	Alias *string `json:"alias,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserAppOwner UserAppOwner

// NewUserAppOwner instantiates a new UserAppOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAppOwner() *UserAppOwner {
	this := UserAppOwner{}
	return &this
}

// NewUserAppOwnerWithDefaults instantiates a new UserAppOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAppOwnerWithDefaults() *UserAppOwner {
	this := UserAppOwner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserAppOwner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAppOwner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserAppOwner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserAppOwner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserAppOwner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAppOwner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserAppOwner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserAppOwner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserAppOwner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAppOwner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserAppOwner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserAppOwner) SetName(v string) {
	o.Name = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *UserAppOwner) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAppOwner) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *UserAppOwner) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *UserAppOwner) SetAlias(v string) {
	o.Alias = &v
}

func (o UserAppOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAppOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserAppOwner) UnmarshalJSON(data []byte) (err error) {
	varUserAppOwner := _UserAppOwner{}

	err = json.Unmarshal(data, &varUserAppOwner)

	if err != nil {
		return err
	}

	*o = UserAppOwner(varUserAppOwner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "alias")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAppOwner struct {
	value *UserAppOwner
	isSet bool
}

func (v NullableUserAppOwner) Get() *UserAppOwner {
	return v.value
}

func (v *NullableUserAppOwner) Set(val *UserAppOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAppOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAppOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAppOwner(val *UserAppOwner) *NullableUserAppOwner {
	return &NullableUserAppOwner{value: val, isSet: true}
}

func (v NullableUserAppOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAppOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


