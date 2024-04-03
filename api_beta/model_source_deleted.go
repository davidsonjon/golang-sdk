/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SourceDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceDeleted{}

// SourceDeleted struct for SourceDeleted
type SourceDeleted struct {
	// The unique ID of the source.
	Id string `json:"id"`
	// Human friendly name of the source.
	Name string `json:"name"`
	// The connection type.
	Type string `json:"type"`
	// The date and time the source was deleted.
	Deleted time.Time `json:"deleted"`
	// The connector type used to connect to the source.
	Connector string `json:"connector"`
	Actor SourceDeletedActor `json:"actor"`
	AdditionalProperties map[string]interface{}
}

type _SourceDeleted SourceDeleted

// NewSourceDeleted instantiates a new SourceDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceDeleted(id string, name string, type_ string, deleted time.Time, connector string, actor SourceDeletedActor) *SourceDeleted {
	this := SourceDeleted{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Deleted = deleted
	this.Connector = connector
	this.Actor = actor
	return &this
}

// NewSourceDeletedWithDefaults instantiates a new SourceDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceDeletedWithDefaults() *SourceDeleted {
	this := SourceDeleted{}
	return &this
}

// GetId returns the Id field value
func (o *SourceDeleted) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceDeleted) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceDeleted) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceDeleted) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceDeleted) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceDeleted) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SourceDeleted) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceDeleted) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceDeleted) SetType(v string) {
	o.Type = v
}

// GetDeleted returns the Deleted field value
func (o *SourceDeleted) GetDeleted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *SourceDeleted) GetDeletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *SourceDeleted) SetDeleted(v time.Time) {
	o.Deleted = v
}

// GetConnector returns the Connector field value
func (o *SourceDeleted) GetConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *SourceDeleted) GetConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *SourceDeleted) SetConnector(v string) {
	o.Connector = v
}

// GetActor returns the Actor field value
func (o *SourceDeleted) GetActor() SourceDeletedActor {
	if o == nil {
		var ret SourceDeletedActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *SourceDeleted) GetActorOk() (*SourceDeletedActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *SourceDeleted) SetActor(v SourceDeletedActor) {
	o.Actor = v
}

func (o SourceDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["deleted"] = o.Deleted
	toSerialize["connector"] = o.Connector
	toSerialize["actor"] = o.Actor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceDeleted) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"deleted",
		"connector",
		"actor",
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

	varSourceDeleted := _SourceDeleted{}

	if err = json.Unmarshal(bytes, &varSourceDeleted); err == nil {
	*o = SourceDeleted(varSourceDeleted)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "deleted")
		delete(additionalProperties, "connector")
		delete(additionalProperties, "actor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceDeleted struct {
	value *SourceDeleted
	isSet bool
}

func (v NullableSourceDeleted) Get() *SourceDeleted {
	return v.value
}

func (v *NullableSourceDeleted) Set(val *SourceDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceDeleted(val *SourceDeleted) *NullableSourceDeleted {
	return &NullableSourceDeleted{value: val, isSet: true}
}

func (v NullableSourceDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


