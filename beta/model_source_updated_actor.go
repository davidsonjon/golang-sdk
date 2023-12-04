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

// checks if the SourceUpdatedActor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceUpdatedActor{}

// SourceUpdatedActor Identity who updated the source.
type SourceUpdatedActor struct {
	// DTO type of identity who updated the source.
	Type string `json:"type"`
	// ID of identity who updated the source.
	Id *string `json:"id,omitempty"`
	// Display name of identity who updated the source.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _SourceUpdatedActor SourceUpdatedActor

// NewSourceUpdatedActor instantiates a new SourceUpdatedActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceUpdatedActor(type_ string, name string) *SourceUpdatedActor {
	this := SourceUpdatedActor{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewSourceUpdatedActorWithDefaults instantiates a new SourceUpdatedActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceUpdatedActorWithDefaults() *SourceUpdatedActor {
	this := SourceUpdatedActor{}
	return &this
}

// GetType returns the Type field value
func (o *SourceUpdatedActor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceUpdatedActor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceUpdatedActor) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceUpdatedActor) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUpdatedActor) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceUpdatedActor) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceUpdatedActor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SourceUpdatedActor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceUpdatedActor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceUpdatedActor) SetName(v string) {
	o.Name = v
}

func (o SourceUpdatedActor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceUpdatedActor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceUpdatedActor) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSourceUpdatedActor := _SourceUpdatedActor{}

	if err = json.Unmarshal(bytes, &varSourceUpdatedActor); err == nil {
	*o = SourceUpdatedActor(varSourceUpdatedActor)
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

type NullableSourceUpdatedActor struct {
	value *SourceUpdatedActor
	isSet bool
}

func (v NullableSourceUpdatedActor) Get() *SourceUpdatedActor {
	return v.value
}

func (v *NullableSourceUpdatedActor) Set(val *SourceUpdatedActor) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceUpdatedActor) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceUpdatedActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceUpdatedActor(val *SourceUpdatedActor) *NullableSourceUpdatedActor {
	return &NullableSourceUpdatedActor{value: val, isSet: true}
}

func (v NullableSourceUpdatedActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceUpdatedActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


