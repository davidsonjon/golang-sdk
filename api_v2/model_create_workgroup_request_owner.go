/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2

import (
	"encoding/json"
)

// checks if the CreateWorkgroupRequestOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkgroupRequestOwner{}

// CreateWorkgroupRequestOwner struct for CreateWorkgroupRequestOwner
type CreateWorkgroupRequestOwner struct {
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateWorkgroupRequestOwner CreateWorkgroupRequestOwner

// NewCreateWorkgroupRequestOwner instantiates a new CreateWorkgroupRequestOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkgroupRequestOwner() *CreateWorkgroupRequestOwner {
	this := CreateWorkgroupRequestOwner{}
	return &this
}

// NewCreateWorkgroupRequestOwnerWithDefaults instantiates a new CreateWorkgroupRequestOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkgroupRequestOwnerWithDefaults() *CreateWorkgroupRequestOwner {
	this := CreateWorkgroupRequestOwner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateWorkgroupRequestOwner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkgroupRequestOwner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateWorkgroupRequestOwner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateWorkgroupRequestOwner) SetId(v string) {
	o.Id = &v
}

func (o CreateWorkgroupRequestOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkgroupRequestOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateWorkgroupRequestOwner) UnmarshalJSON(bytes []byte) (err error) {
	varCreateWorkgroupRequestOwner := _CreateWorkgroupRequestOwner{}

	if err = json.Unmarshal(bytes, &varCreateWorkgroupRequestOwner); err == nil {
	*o = CreateWorkgroupRequestOwner(varCreateWorkgroupRequestOwner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateWorkgroupRequestOwner struct {
	value *CreateWorkgroupRequestOwner
	isSet bool
}

func (v NullableCreateWorkgroupRequestOwner) Get() *CreateWorkgroupRequestOwner {
	return v.value
}

func (v *NullableCreateWorkgroupRequestOwner) Set(val *CreateWorkgroupRequestOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkgroupRequestOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkgroupRequestOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkgroupRequestOwner(val *CreateWorkgroupRequestOwner) *NullableCreateWorkgroupRequestOwner {
	return &NullableCreateWorkgroupRequestOwner{value: val, isSet: true}
}

func (v NullableCreateWorkgroupRequestOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkgroupRequestOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

