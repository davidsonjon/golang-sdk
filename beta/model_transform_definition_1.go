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

// checks if the TransformDefinition1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformDefinition1{}

// TransformDefinition1 struct for TransformDefinition1
type TransformDefinition1 struct {
	// The type of the transform definition.
	Type *string `json:"type,omitempty"`
	// Arbitrary key-value pairs to store any metadata for the object
	Attributes *map[string]TransformDefinition1AttributesValue `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransformDefinition1 TransformDefinition1

// NewTransformDefinition1 instantiates a new TransformDefinition1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformDefinition1() *TransformDefinition1 {
	this := TransformDefinition1{}
	return &this
}

// NewTransformDefinition1WithDefaults instantiates a new TransformDefinition1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformDefinition1WithDefaults() *TransformDefinition1 {
	this := TransformDefinition1{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransformDefinition1) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformDefinition1) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransformDefinition1) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransformDefinition1) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TransformDefinition1) GetAttributes() map[string]TransformDefinition1AttributesValue {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]TransformDefinition1AttributesValue
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformDefinition1) GetAttributesOk() (*map[string]TransformDefinition1AttributesValue, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TransformDefinition1) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]TransformDefinition1AttributesValue and assigns it to the Attributes field.
func (o *TransformDefinition1) SetAttributes(v map[string]TransformDefinition1AttributesValue) {
	o.Attributes = &v
}

func (o TransformDefinition1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformDefinition1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransformDefinition1) UnmarshalJSON(bytes []byte) (err error) {
	varTransformDefinition1 := _TransformDefinition1{}

	if err = json.Unmarshal(bytes, &varTransformDefinition1); err == nil {
	*o = TransformDefinition1(varTransformDefinition1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransformDefinition1 struct {
	value *TransformDefinition1
	isSet bool
}

func (v NullableTransformDefinition1) Get() *TransformDefinition1 {
	return v.value
}

func (v *NullableTransformDefinition1) Set(val *TransformDefinition1) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformDefinition1) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformDefinition1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformDefinition1(val *TransformDefinition1) *NullableTransformDefinition1 {
	return &NullableTransformDefinition1{value: val, isSet: true}
}

func (v NullableTransformDefinition1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformDefinition1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


