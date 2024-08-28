/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the BaseDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDocument{}

// BaseDocument struct for BaseDocument
type BaseDocument struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type DocumentType `json:"_type"`
	AdditionalProperties map[string]interface{}
}

type _BaseDocument BaseDocument

// NewBaseDocument instantiates a new BaseDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDocument(id string, name string, type_ DocumentType) *BaseDocument {
	this := BaseDocument{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewBaseDocumentWithDefaults instantiates a new BaseDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDocumentWithDefaults() *BaseDocument {
	this := BaseDocument{}
	return &this
}

// GetId returns the Id field value
func (o *BaseDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseDocument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseDocument) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BaseDocument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseDocument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseDocument) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BaseDocument) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseDocument) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseDocument) SetType(v DocumentType) {
	o.Type = v
}

func (o BaseDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["_type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseDocument) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"_type",
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

	varBaseDocument := _BaseDocument{}

	err = json.Unmarshal(data, &varBaseDocument)

	if err != nil {
		return err
	}

	*o = BaseDocument(varBaseDocument)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseDocument struct {
	value *BaseDocument
	isSet bool
}

func (v NullableBaseDocument) Get() *BaseDocument {
	return v.value
}

func (v *NullableBaseDocument) Set(val *BaseDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDocument(val *BaseDocument) *NullableBaseDocument {
	return &NullableBaseDocument{value: val, isSet: true}
}

func (v NullableBaseDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


