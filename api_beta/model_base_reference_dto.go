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

// checks if the BaseReferenceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseReferenceDto{}

// BaseReferenceDto struct for BaseReferenceDto
type BaseReferenceDto struct {
	// the application ID
	Id *string `json:"id,omitempty"`
	// the application name
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseReferenceDto BaseReferenceDto

// NewBaseReferenceDto instantiates a new BaseReferenceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReferenceDto() *BaseReferenceDto {
	this := BaseReferenceDto{}
	return &this
}

// NewBaseReferenceDtoWithDefaults instantiates a new BaseReferenceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReferenceDtoWithDefaults() *BaseReferenceDto {
	this := BaseReferenceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseReferenceDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReferenceDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseReferenceDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseReferenceDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseReferenceDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReferenceDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseReferenceDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseReferenceDto) SetName(v string) {
	o.Name = &v
}

func (o BaseReferenceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseReferenceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *BaseReferenceDto) UnmarshalJSON(data []byte) (err error) {
	varBaseReferenceDto := _BaseReferenceDto{}

	err = json.Unmarshal(data, &varBaseReferenceDto)

	if err != nil {
		return err
	}

	*o = BaseReferenceDto(varBaseReferenceDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReferenceDto struct {
	value *BaseReferenceDto
	isSet bool
}

func (v NullableBaseReferenceDto) Get() *BaseReferenceDto {
	return v.value
}

func (v *NullableBaseReferenceDto) Set(val *BaseReferenceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReferenceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReferenceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReferenceDto(val *BaseReferenceDto) *NullableBaseReferenceDto {
	return &NullableBaseReferenceDto{value: val, isSet: true}
}

func (v NullableBaseReferenceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReferenceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


