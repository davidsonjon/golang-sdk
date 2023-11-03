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

// checks if the AccessItemRequesterDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemRequesterDto{}

// AccessItemRequesterDto Access item requester's identity.
type AccessItemRequesterDto struct {
	// Access item requester's DTO type.
	Type *string `json:"type,omitempty"`
	// Access item requester's identity ID.
	Id *string `json:"id,omitempty"`
	// Access item owner's human-readable display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRequesterDto AccessItemRequesterDto

// NewAccessItemRequesterDto instantiates a new AccessItemRequesterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRequesterDto() *AccessItemRequesterDto {
	this := AccessItemRequesterDto{}
	return &this
}

// NewAccessItemRequesterDtoWithDefaults instantiates a new AccessItemRequesterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemRequesterDtoWithDefaults() *AccessItemRequesterDto {
	this := AccessItemRequesterDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessItemRequesterDto) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequesterDto) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessItemRequesterDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessItemRequesterDto) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemRequesterDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequesterDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemRequesterDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemRequesterDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessItemRequesterDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequesterDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessItemRequesterDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessItemRequesterDto) SetName(v string) {
	o.Name = &v
}

func (o AccessItemRequesterDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemRequesterDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
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

func (o *AccessItemRequesterDto) UnmarshalJSON(bytes []byte) (err error) {
	varAccessItemRequesterDto := _AccessItemRequesterDto{}

	if err = json.Unmarshal(bytes, &varAccessItemRequesterDto); err == nil {
		*o = AccessItemRequesterDto(varAccessItemRequesterDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemRequesterDto struct {
	value *AccessItemRequesterDto
	isSet bool
}

func (v NullableAccessItemRequesterDto) Get() *AccessItemRequesterDto {
	return v.value
}

func (v *NullableAccessItemRequesterDto) Set(val *AccessItemRequesterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemRequesterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemRequesterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemRequesterDto(val *AccessItemRequesterDto) *NullableAccessItemRequesterDto {
	return &NullableAccessItemRequesterDto{value: val, isSet: true}
}

func (v NullableAccessItemRequesterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemRequesterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


