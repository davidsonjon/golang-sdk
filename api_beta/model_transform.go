/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the Transform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transform{}

// Transform The representation of an internally- or customer-defined transform.
type Transform struct {
	// Unique name of this transform
	Name string `json:"name"`
	// The type of transform operation
	Type string `json:"type"`
	// Meta-data about the transform. Values in this list are specific to the type of transform to be executed.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _Transform Transform

// NewTransform instantiates a new Transform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransform(name string, type_ string, attributes map[string]interface{}) *Transform {
	this := Transform{}
	this.Name = name
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTransformWithDefaults instantiates a new Transform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformWithDefaults() *Transform {
	this := Transform{}
	return &this
}

// GetName returns the Name field value
func (o *Transform) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Transform) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Transform) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Transform) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transform) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transform) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Transform) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transform) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *Transform) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o Transform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Transform) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransform := _Transform{}

	if err = json.Unmarshal(bytes, &varTransform); err == nil {
	*o = Transform(varTransform)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransform struct {
	value *Transform
	isSet bool
}

func (v NullableTransform) Get() *Transform {
	return v.value
}

func (v *NullableTransform) Set(val *Transform) {
	v.value = val
	v.isSet = true
}

func (v NullableTransform) IsSet() bool {
	return v.isSet
}

func (v *NullableTransform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransform(val *Transform) *NullableTransform {
	return &NullableTransform{value: val, isSet: true}
}

func (v NullableTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


