/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the SourceCreatedActor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceCreatedActor{}

// SourceCreatedActor Identity who created the source.
type SourceCreatedActor struct {
	// DTO type of identity who created the source.
	Type string `json:"type"`
	// ID of identity who created the source.
	Id string `json:"id"`
	// Display name of identity who created the source.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _SourceCreatedActor SourceCreatedActor

// NewSourceCreatedActor instantiates a new SourceCreatedActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceCreatedActor(type_ string, id string, name string) *SourceCreatedActor {
	this := SourceCreatedActor{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewSourceCreatedActorWithDefaults instantiates a new SourceCreatedActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceCreatedActorWithDefaults() *SourceCreatedActor {
	this := SourceCreatedActor{}
	return &this
}

// GetType returns the Type field value
func (o *SourceCreatedActor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceCreatedActor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceCreatedActor) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SourceCreatedActor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceCreatedActor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceCreatedActor) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceCreatedActor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceCreatedActor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceCreatedActor) SetName(v string) {
	o.Name = v
}

func (o SourceCreatedActor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceCreatedActor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceCreatedActor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSourceCreatedActor := _SourceCreatedActor{}

	err = json.Unmarshal(data, &varSourceCreatedActor)

	if err != nil {
		return err
	}

	*o = SourceCreatedActor(varSourceCreatedActor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceCreatedActor struct {
	value *SourceCreatedActor
	isSet bool
}

func (v NullableSourceCreatedActor) Get() *SourceCreatedActor {
	return v.value
}

func (v *NullableSourceCreatedActor) Set(val *SourceCreatedActor) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceCreatedActor) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceCreatedActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceCreatedActor(val *SourceCreatedActor) *NullableSourceCreatedActor {
	return &NullableSourceCreatedActor{value: val, isSet: true}
}

func (v NullableSourceCreatedActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceCreatedActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


