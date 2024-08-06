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

// checks if the Entitlement1AccessModelMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entitlement1AccessModelMetadata{}

// Entitlement1AccessModelMetadata struct for Entitlement1AccessModelMetadata
type Entitlement1AccessModelMetadata struct {
	Attributes []AttributeDTO `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Entitlement1AccessModelMetadata Entitlement1AccessModelMetadata

// NewEntitlement1AccessModelMetadata instantiates a new Entitlement1AccessModelMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement1AccessModelMetadata() *Entitlement1AccessModelMetadata {
	this := Entitlement1AccessModelMetadata{}
	return &this
}

// NewEntitlement1AccessModelMetadataWithDefaults instantiates a new Entitlement1AccessModelMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlement1AccessModelMetadataWithDefaults() *Entitlement1AccessModelMetadata {
	this := Entitlement1AccessModelMetadata{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entitlement1AccessModelMetadata) GetAttributes() []AttributeDTO {
	if o == nil {
		var ret []AttributeDTO
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlement1AccessModelMetadata) GetAttributesOk() ([]AttributeDTO, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Entitlement1AccessModelMetadata) HasAttributes() bool {
	if o != nil && isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []AttributeDTO and assigns it to the Attributes field.
func (o *Entitlement1AccessModelMetadata) SetAttributes(v []AttributeDTO) {
	o.Attributes = v
}

func (o Entitlement1AccessModelMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entitlement1AccessModelMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Entitlement1AccessModelMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlement1AccessModelMetadata := _Entitlement1AccessModelMetadata{}

	if err = json.Unmarshal(bytes, &varEntitlement1AccessModelMetadata); err == nil {
	*o = Entitlement1AccessModelMetadata(varEntitlement1AccessModelMetadata)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlement1AccessModelMetadata struct {
	value *Entitlement1AccessModelMetadata
	isSet bool
}

func (v NullableEntitlement1AccessModelMetadata) Get() *Entitlement1AccessModelMetadata {
	return v.value
}

func (v *NullableEntitlement1AccessModelMetadata) Set(val *Entitlement1AccessModelMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement1AccessModelMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement1AccessModelMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement1AccessModelMetadata(val *Entitlement1AccessModelMetadata) *NullableEntitlement1AccessModelMetadata {
	return &NullableEntitlement1AccessModelMetadata{value: val, isSet: true}
}

func (v NullableEntitlement1AccessModelMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement1AccessModelMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


