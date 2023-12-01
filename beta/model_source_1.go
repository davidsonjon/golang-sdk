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

// checks if the Source1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Source1{}

// Source1 struct for Source1
type Source1 struct {
	// The type of the source
	Type *string `json:"type,omitempty"`
	// The source properties
	Properties map[string]interface{} `json:"properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Source1 Source1

// NewSource1 instantiates a new Source1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource1() *Source1 {
	this := Source1{}
	return &this
}

// NewSource1WithDefaults instantiates a new Source1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSource1WithDefaults() *Source1 {
	this := Source1{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Source1) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source1) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Source1) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Source1) SetType(v string) {
	o.Type = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Source1) GetProperties() map[string]interface{} {
	if o == nil || isNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source1) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Source1) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *Source1) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o Source1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Source1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Source1) UnmarshalJSON(bytes []byte) (err error) {
	varSource1 := _Source1{}

	if err = json.Unmarshal(bytes, &varSource1); err == nil {
	*o = Source1(varSource1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSource1 struct {
	value *Source1
	isSet bool
}

func (v NullableSource1) Get() *Source1 {
	return v.value
}

func (v *NullableSource1) Set(val *Source1) {
	v.value = val
	v.isSet = true
}

func (v NullableSource1) IsSet() bool {
	return v.isSet
}

func (v *NullableSource1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource1(val *Source1) *NullableSource1 {
	return &NullableSource1{value: val, isSet: true}
}

func (v NullableSource1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


