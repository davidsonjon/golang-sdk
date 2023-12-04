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

// checks if the IdentityAttributeTransform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAttributeTransform{}

// IdentityAttributeTransform struct for IdentityAttributeTransform
type IdentityAttributeTransform struct {
	// Name of the identity attribute
	IdentityAttributeName *string `json:"identityAttributeName,omitempty"`
	TransformDefinition *TransformDefinition `json:"transformDefinition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttributeTransform IdentityAttributeTransform

// NewIdentityAttributeTransform instantiates a new IdentityAttributeTransform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttributeTransform() *IdentityAttributeTransform {
	this := IdentityAttributeTransform{}
	return &this
}

// NewIdentityAttributeTransformWithDefaults instantiates a new IdentityAttributeTransform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttributeTransformWithDefaults() *IdentityAttributeTransform {
	this := IdentityAttributeTransform{}
	return &this
}

// GetIdentityAttributeName returns the IdentityAttributeName field value if set, zero value otherwise.
func (o *IdentityAttributeTransform) GetIdentityAttributeName() string {
	if o == nil || isNil(o.IdentityAttributeName) {
		var ret string
		return ret
	}
	return *o.IdentityAttributeName
}

// GetIdentityAttributeNameOk returns a tuple with the IdentityAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributeTransform) GetIdentityAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.IdentityAttributeName) {
		return nil, false
	}
	return o.IdentityAttributeName, true
}

// HasIdentityAttributeName returns a boolean if a field has been set.
func (o *IdentityAttributeTransform) HasIdentityAttributeName() bool {
	if o != nil && !isNil(o.IdentityAttributeName) {
		return true
	}

	return false
}

// SetIdentityAttributeName gets a reference to the given string and assigns it to the IdentityAttributeName field.
func (o *IdentityAttributeTransform) SetIdentityAttributeName(v string) {
	o.IdentityAttributeName = &v
}

// GetTransformDefinition returns the TransformDefinition field value if set, zero value otherwise.
func (o *IdentityAttributeTransform) GetTransformDefinition() TransformDefinition {
	if o == nil || isNil(o.TransformDefinition) {
		var ret TransformDefinition
		return ret
	}
	return *o.TransformDefinition
}

// GetTransformDefinitionOk returns a tuple with the TransformDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributeTransform) GetTransformDefinitionOk() (*TransformDefinition, bool) {
	if o == nil || isNil(o.TransformDefinition) {
		return nil, false
	}
	return o.TransformDefinition, true
}

// HasTransformDefinition returns a boolean if a field has been set.
func (o *IdentityAttributeTransform) HasTransformDefinition() bool {
	if o != nil && !isNil(o.TransformDefinition) {
		return true
	}

	return false
}

// SetTransformDefinition gets a reference to the given TransformDefinition and assigns it to the TransformDefinition field.
func (o *IdentityAttributeTransform) SetTransformDefinition(v TransformDefinition) {
	o.TransformDefinition = &v
}

func (o IdentityAttributeTransform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAttributeTransform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityAttributeName) {
		toSerialize["identityAttributeName"] = o.IdentityAttributeName
	}
	if !isNil(o.TransformDefinition) {
		toSerialize["transformDefinition"] = o.TransformDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityAttributeTransform) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityAttributeTransform := _IdentityAttributeTransform{}

	if err = json.Unmarshal(bytes, &varIdentityAttributeTransform); err == nil {
	*o = IdentityAttributeTransform(varIdentityAttributeTransform)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityAttributeName")
		delete(additionalProperties, "transformDefinition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAttributeTransform struct {
	value *IdentityAttributeTransform
	isSet bool
}

func (v NullableIdentityAttributeTransform) Get() *IdentityAttributeTransform {
	return v.value
}

func (v *NullableIdentityAttributeTransform) Set(val *IdentityAttributeTransform) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributeTransform) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributeTransform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributeTransform(val *IdentityAttributeTransform) *NullableIdentityAttributeTransform {
	return &NullableIdentityAttributeTransform{value: val, isSet: true}
}

func (v NullableIdentityAttributeTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributeTransform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


