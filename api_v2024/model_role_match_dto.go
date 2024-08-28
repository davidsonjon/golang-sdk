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

// checks if the RoleMatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMatchDto{}

// RoleMatchDto struct for RoleMatchDto
type RoleMatchDto struct {
	RoleRef *BaseReferenceDto1 `json:"roleRef,omitempty"`
	MatchedAttributes []ContextAttributeDto `json:"matchedAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMatchDto RoleMatchDto

// NewRoleMatchDto instantiates a new RoleMatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMatchDto() *RoleMatchDto {
	this := RoleMatchDto{}
	return &this
}

// NewRoleMatchDtoWithDefaults instantiates a new RoleMatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMatchDtoWithDefaults() *RoleMatchDto {
	this := RoleMatchDto{}
	return &this
}

// GetRoleRef returns the RoleRef field value if set, zero value otherwise.
func (o *RoleMatchDto) GetRoleRef() BaseReferenceDto1 {
	if o == nil || IsNil(o.RoleRef) {
		var ret BaseReferenceDto1
		return ret
	}
	return *o.RoleRef
}

// GetRoleRefOk returns a tuple with the RoleRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMatchDto) GetRoleRefOk() (*BaseReferenceDto1, bool) {
	if o == nil || IsNil(o.RoleRef) {
		return nil, false
	}
	return o.RoleRef, true
}

// HasRoleRef returns a boolean if a field has been set.
func (o *RoleMatchDto) HasRoleRef() bool {
	if o != nil && !IsNil(o.RoleRef) {
		return true
	}

	return false
}

// SetRoleRef gets a reference to the given BaseReferenceDto1 and assigns it to the RoleRef field.
func (o *RoleMatchDto) SetRoleRef(v BaseReferenceDto1) {
	o.RoleRef = &v
}

// GetMatchedAttributes returns the MatchedAttributes field value if set, zero value otherwise.
func (o *RoleMatchDto) GetMatchedAttributes() []ContextAttributeDto {
	if o == nil || IsNil(o.MatchedAttributes) {
		var ret []ContextAttributeDto
		return ret
	}
	return o.MatchedAttributes
}

// GetMatchedAttributesOk returns a tuple with the MatchedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMatchDto) GetMatchedAttributesOk() ([]ContextAttributeDto, bool) {
	if o == nil || IsNil(o.MatchedAttributes) {
		return nil, false
	}
	return o.MatchedAttributes, true
}

// HasMatchedAttributes returns a boolean if a field has been set.
func (o *RoleMatchDto) HasMatchedAttributes() bool {
	if o != nil && !IsNil(o.MatchedAttributes) {
		return true
	}

	return false
}

// SetMatchedAttributes gets a reference to the given []ContextAttributeDto and assigns it to the MatchedAttributes field.
func (o *RoleMatchDto) SetMatchedAttributes(v []ContextAttributeDto) {
	o.MatchedAttributes = v
}

func (o RoleMatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleRef) {
		toSerialize["roleRef"] = o.RoleRef
	}
	if !IsNil(o.MatchedAttributes) {
		toSerialize["matchedAttributes"] = o.MatchedAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMatchDto) UnmarshalJSON(data []byte) (err error) {
	varRoleMatchDto := _RoleMatchDto{}

	err = json.Unmarshal(data, &varRoleMatchDto)

	if err != nil {
		return err
	}

	*o = RoleMatchDto(varRoleMatchDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "roleRef")
		delete(additionalProperties, "matchedAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMatchDto struct {
	value *RoleMatchDto
	isSet bool
}

func (v NullableRoleMatchDto) Get() *RoleMatchDto {
	return v.value
}

func (v *NullableRoleMatchDto) Set(val *RoleMatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMatchDto(val *RoleMatchDto) *NullableRoleMatchDto {
	return &NullableRoleMatchDto{value: val, isSet: true}
}

func (v NullableRoleMatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


