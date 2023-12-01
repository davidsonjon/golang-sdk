/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the TypedReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypedReference{}

// TypedReference A typed reference to the object. 
type TypedReference struct {
	Type DtoType `json:"type"`
	// The id of the object. 
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _TypedReference TypedReference

// NewTypedReference instantiates a new TypedReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypedReference(type_ DtoType, id string) *TypedReference {
	this := TypedReference{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewTypedReferenceWithDefaults instantiates a new TypedReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypedReferenceWithDefaults() *TypedReference {
	this := TypedReference{}
	return &this
}

// GetType returns the Type field value
func (o *TypedReference) GetType() DtoType {
	if o == nil {
		var ret DtoType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TypedReference) GetTypeOk() (*DtoType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TypedReference) SetType(v DtoType) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TypedReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TypedReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TypedReference) SetId(v string) {
	o.Id = v
}

func (o TypedReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypedReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TypedReference) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varTypedReference := _TypedReference{}

	if err = json.Unmarshal(bytes, &varTypedReference); err == nil {
	*o = TypedReference(varTypedReference)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTypedReference struct {
	value *TypedReference
	isSet bool
}

func (v NullableTypedReference) Get() *TypedReference {
	return v.value
}

func (v *NullableTypedReference) Set(val *TypedReference) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedReference) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedReference(val *TypedReference) *NullableTypedReference {
	return &NullableTypedReference{value: val, isSet: true}
}

func (v NullableTypedReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


