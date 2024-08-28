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

// checks if the AccessItemRequestedFor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemRequestedFor{}

// AccessItemRequestedFor Identity the access item is requested for.
type AccessItemRequestedFor struct {
	// DTO type of identity the access item is requested for.
	Type *string `json:"type,omitempty"`
	// ID of identity the access item is requested for.
	Id *string `json:"id,omitempty"`
	// Human-readable display name of identity the access item is requested for.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRequestedFor AccessItemRequestedFor

// NewAccessItemRequestedFor instantiates a new AccessItemRequestedFor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRequestedFor() *AccessItemRequestedFor {
	this := AccessItemRequestedFor{}
	return &this
}

// NewAccessItemRequestedForWithDefaults instantiates a new AccessItemRequestedFor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemRequestedForWithDefaults() *AccessItemRequestedFor {
	this := AccessItemRequestedFor{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessItemRequestedFor) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequestedFor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessItemRequestedFor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessItemRequestedFor) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemRequestedFor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequestedFor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemRequestedFor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemRequestedFor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessItemRequestedFor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequestedFor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessItemRequestedFor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessItemRequestedFor) SetName(v string) {
	o.Name = &v
}

func (o AccessItemRequestedFor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemRequestedFor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemRequestedFor) UnmarshalJSON(data []byte) (err error) {
	varAccessItemRequestedFor := _AccessItemRequestedFor{}

	err = json.Unmarshal(data, &varAccessItemRequestedFor)

	if err != nil {
		return err
	}

	*o = AccessItemRequestedFor(varAccessItemRequestedFor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemRequestedFor struct {
	value *AccessItemRequestedFor
	isSet bool
}

func (v NullableAccessItemRequestedFor) Get() *AccessItemRequestedFor {
	return v.value
}

func (v *NullableAccessItemRequestedFor) Set(val *AccessItemRequestedFor) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemRequestedFor) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemRequestedFor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemRequestedFor(val *AccessItemRequestedFor) *NullableAccessItemRequestedFor {
	return &NullableAccessItemRequestedFor{value: val, isSet: true}
}

func (v NullableAccessItemRequestedFor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemRequestedFor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


