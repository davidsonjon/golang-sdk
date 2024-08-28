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

// checks if the AccessItemRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemRef{}

// AccessItemRef struct for AccessItemRef
type AccessItemRef struct {
	// ID of the access item to retrieve the recommendation for.
	Id *string `json:"id,omitempty"`
	// Access item's type.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRef AccessItemRef

// NewAccessItemRef instantiates a new AccessItemRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRef() *AccessItemRef {
	this := AccessItemRef{}
	return &this
}

// NewAccessItemRefWithDefaults instantiates a new AccessItemRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemRefWithDefaults() *AccessItemRef {
	this := AccessItemRef{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemRef) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRef) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemRef) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemRef) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessItemRef) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRef) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessItemRef) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessItemRef) SetType(v string) {
	o.Type = &v
}

func (o AccessItemRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemRef) UnmarshalJSON(data []byte) (err error) {
	varAccessItemRef := _AccessItemRef{}

	err = json.Unmarshal(data, &varAccessItemRef)

	if err != nil {
		return err
	}

	*o = AccessItemRef(varAccessItemRef)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemRef struct {
	value *AccessItemRef
	isSet bool
}

func (v NullableAccessItemRef) Get() *AccessItemRef {
	return v.value
}

func (v *NullableAccessItemRef) Set(val *AccessItemRef) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemRef) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemRef(val *AccessItemRef) *NullableAccessItemRef {
	return &NullableAccessItemRef{value: val, isSet: true}
}

func (v NullableAccessItemRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


