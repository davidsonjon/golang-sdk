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

// checks if the OwnerReferenceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerReferenceDto{}

// OwnerReferenceDto Simplified DTO for the owner object of the entitlement
type OwnerReferenceDto struct {
	// The owner id for the entitlement
	Id *string `json:"id,omitempty"`
	// The owner name for the entitlement
	Name *string `json:"name,omitempty"`
	// The type of the owner. Initially only type IDENTITY is supported
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OwnerReferenceDto OwnerReferenceDto

// NewOwnerReferenceDto instantiates a new OwnerReferenceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerReferenceDto() *OwnerReferenceDto {
	this := OwnerReferenceDto{}
	return &this
}

// NewOwnerReferenceDtoWithDefaults instantiates a new OwnerReferenceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerReferenceDtoWithDefaults() *OwnerReferenceDto {
	this := OwnerReferenceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OwnerReferenceDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReferenceDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OwnerReferenceDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OwnerReferenceDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OwnerReferenceDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReferenceDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OwnerReferenceDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OwnerReferenceDto) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OwnerReferenceDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReferenceDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OwnerReferenceDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OwnerReferenceDto) SetType(v string) {
	o.Type = &v
}

func (o OwnerReferenceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerReferenceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OwnerReferenceDto) UnmarshalJSON(data []byte) (err error) {
	varOwnerReferenceDto := _OwnerReferenceDto{}

	err = json.Unmarshal(data, &varOwnerReferenceDto)

	if err != nil {
		return err
	}

	*o = OwnerReferenceDto(varOwnerReferenceDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOwnerReferenceDto struct {
	value *OwnerReferenceDto
	isSet bool
}

func (v NullableOwnerReferenceDto) Get() *OwnerReferenceDto {
	return v.value
}

func (v *NullableOwnerReferenceDto) Set(val *OwnerReferenceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerReferenceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerReferenceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerReferenceDto(val *OwnerReferenceDto) *NullableOwnerReferenceDto {
	return &NullableOwnerReferenceDto{value: val, isSet: true}
}

func (v NullableOwnerReferenceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerReferenceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


