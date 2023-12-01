/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// checks if the SourceDeletedActor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceDeletedActor{}

// SourceDeletedActor Identity who deleted the source.
type SourceDeletedActor struct {
	// DTO type of identity who deleted the source.
	Type string `json:"type"`
	// ID of identity who deleted the source.
	Id string `json:"id"`
	// Display name of identity who deleted the source.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _SourceDeletedActor SourceDeletedActor

// NewSourceDeletedActor instantiates a new SourceDeletedActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceDeletedActor(type_ string, id string, name string) *SourceDeletedActor {
	this := SourceDeletedActor{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewSourceDeletedActorWithDefaults instantiates a new SourceDeletedActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceDeletedActorWithDefaults() *SourceDeletedActor {
	this := SourceDeletedActor{}
	return &this
}

// GetType returns the Type field value
func (o *SourceDeletedActor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceDeletedActor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceDeletedActor) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SourceDeletedActor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceDeletedActor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceDeletedActor) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceDeletedActor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceDeletedActor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceDeletedActor) SetName(v string) {
	o.Name = v
}

func (o SourceDeletedActor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceDeletedActor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceDeletedActor) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
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

	varSourceDeletedActor := _SourceDeletedActor{}

	if err = json.Unmarshal(bytes, &varSourceDeletedActor); err == nil {
	*o = SourceDeletedActor(varSourceDeletedActor)
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

type NullableSourceDeletedActor struct {
	value *SourceDeletedActor
	isSet bool
}

func (v NullableSourceDeletedActor) Get() *SourceDeletedActor {
	return v.value
}

func (v *NullableSourceDeletedActor) Set(val *SourceDeletedActor) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceDeletedActor) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceDeletedActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceDeletedActor(val *SourceDeletedActor) *NullableSourceDeletedActor {
	return &NullableSourceDeletedActor{value: val, isSet: true}
}

func (v NullableSourceDeletedActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceDeletedActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


