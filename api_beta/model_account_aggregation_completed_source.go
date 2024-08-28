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

// checks if the AccountAggregationCompletedSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAggregationCompletedSource{}

// AccountAggregationCompletedSource The source the accounts are being aggregated from.
type AccountAggregationCompletedSource struct {
	// The DTO type of the source the accounts are being aggregated from.
	Type string `json:"type"`
	// The ID of the source the accounts are being aggregated from.
	Id string `json:"id"`
	// Display name of the source the accounts are being aggregated from.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AccountAggregationCompletedSource AccountAggregationCompletedSource

// NewAccountAggregationCompletedSource instantiates a new AccountAggregationCompletedSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAggregationCompletedSource(type_ string, id string, name string) *AccountAggregationCompletedSource {
	this := AccountAggregationCompletedSource{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewAccountAggregationCompletedSourceWithDefaults instantiates a new AccountAggregationCompletedSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAggregationCompletedSourceWithDefaults() *AccountAggregationCompletedSource {
	this := AccountAggregationCompletedSource{}
	return &this
}

// GetType returns the Type field value
func (o *AccountAggregationCompletedSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountAggregationCompletedSource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AccountAggregationCompletedSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountAggregationCompletedSource) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccountAggregationCompletedSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountAggregationCompletedSource) SetName(v string) {
	o.Name = v
}

func (o AccountAggregationCompletedSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAggregationCompletedSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountAggregationCompletedSource) UnmarshalJSON(data []byte) (err error) {
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

	varAccountAggregationCompletedSource := _AccountAggregationCompletedSource{}

	err = json.Unmarshal(data, &varAccountAggregationCompletedSource)

	if err != nil {
		return err
	}

	*o = AccountAggregationCompletedSource(varAccountAggregationCompletedSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAggregationCompletedSource struct {
	value *AccountAggregationCompletedSource
	isSet bool
}

func (v NullableAccountAggregationCompletedSource) Get() *AccountAggregationCompletedSource {
	return v.value
}

func (v *NullableAccountAggregationCompletedSource) Set(val *AccountAggregationCompletedSource) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAggregationCompletedSource) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAggregationCompletedSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAggregationCompletedSource(val *AccountAggregationCompletedSource) *NullableAccountAggregationCompletedSource {
	return &NullableAccountAggregationCompletedSource{value: val, isSet: true}
}

func (v NullableAccountAggregationCompletedSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAggregationCompletedSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


