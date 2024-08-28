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

// checks if the IdentityPreviewResponseIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityPreviewResponseIdentity{}

// IdentityPreviewResponseIdentity Identity's basic details.
type IdentityPreviewResponseIdentity struct {
	// Identity's DTO type.
	Type *string `json:"type,omitempty"`
	// Identity ID.
	Id *string `json:"id,omitempty"`
	// Identity's display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityPreviewResponseIdentity IdentityPreviewResponseIdentity

// NewIdentityPreviewResponseIdentity instantiates a new IdentityPreviewResponseIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPreviewResponseIdentity() *IdentityPreviewResponseIdentity {
	this := IdentityPreviewResponseIdentity{}
	return &this
}

// NewIdentityPreviewResponseIdentityWithDefaults instantiates a new IdentityPreviewResponseIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPreviewResponseIdentityWithDefaults() *IdentityPreviewResponseIdentity {
	this := IdentityPreviewResponseIdentity{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityPreviewResponseIdentity) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewResponseIdentity) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityPreviewResponseIdentity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityPreviewResponseIdentity) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityPreviewResponseIdentity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewResponseIdentity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityPreviewResponseIdentity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityPreviewResponseIdentity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityPreviewResponseIdentity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewResponseIdentity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityPreviewResponseIdentity) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityPreviewResponseIdentity) SetName(v string) {
	o.Name = &v
}

func (o IdentityPreviewResponseIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityPreviewResponseIdentity) ToMap() (map[string]interface{}, error) {
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

func (o *IdentityPreviewResponseIdentity) UnmarshalJSON(data []byte) (err error) {
	varIdentityPreviewResponseIdentity := _IdentityPreviewResponseIdentity{}

	err = json.Unmarshal(data, &varIdentityPreviewResponseIdentity)

	if err != nil {
		return err
	}

	*o = IdentityPreviewResponseIdentity(varIdentityPreviewResponseIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityPreviewResponseIdentity struct {
	value *IdentityPreviewResponseIdentity
	isSet bool
}

func (v NullableIdentityPreviewResponseIdentity) Get() *IdentityPreviewResponseIdentity {
	return v.value
}

func (v *NullableIdentityPreviewResponseIdentity) Set(val *IdentityPreviewResponseIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPreviewResponseIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPreviewResponseIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPreviewResponseIdentity(val *IdentityPreviewResponseIdentity) *NullableIdentityPreviewResponseIdentity {
	return &NullableIdentityPreviewResponseIdentity{value: val, isSet: true}
}

func (v NullableIdentityPreviewResponseIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPreviewResponseIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


