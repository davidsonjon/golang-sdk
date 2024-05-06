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

// checks if the Identity1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Identity1{}

// Identity1 The definition of an Identity according to the Reassignment Configuration service
type Identity1 struct {
	// The ID of the object
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Identity1 Identity1

// NewIdentity1 instantiates a new Identity1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity1() *Identity1 {
	this := Identity1{}
	return &this
}

// NewIdentity1WithDefaults instantiates a new Identity1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentity1WithDefaults() *Identity1 {
	this := Identity1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Identity1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Identity1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Identity1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Identity1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Identity1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Identity1) SetName(v string) {
	o.Name = &v
}

func (o Identity1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Identity1) ToMap() (map[string]interface{}, error) {
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

func (o *Identity1) UnmarshalJSON(bytes []byte) (err error) {
	varIdentity1 := _Identity1{}

	if err = json.Unmarshal(bytes, &varIdentity1); err == nil {
	*o = Identity1(varIdentity1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentity1 struct {
	value *Identity1
	isSet bool
}

func (v NullableIdentity1) Get() *Identity1 {
	return v.value
}

func (v *NullableIdentity1) Set(val *Identity1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity1(val *Identity1) *NullableIdentity1 {
	return &NullableIdentity1{value: val, isSet: true}
}

func (v NullableIdentity1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


