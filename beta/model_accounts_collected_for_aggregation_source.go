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

// checks if the AccountsCollectedForAggregationSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsCollectedForAggregationSource{}

// AccountsCollectedForAggregationSource Reference to the source that has been aggregated.
type AccountsCollectedForAggregationSource struct {
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// The type of object that is referenced
	Type string `json:"type"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AccountsCollectedForAggregationSource AccountsCollectedForAggregationSource

// NewAccountsCollectedForAggregationSource instantiates a new AccountsCollectedForAggregationSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsCollectedForAggregationSource(id string, type_ string, name string) *AccountsCollectedForAggregationSource {
	this := AccountsCollectedForAggregationSource{}
	this.Id = id
	this.Type = type_
	this.Name = name
	return &this
}

// NewAccountsCollectedForAggregationSourceWithDefaults instantiates a new AccountsCollectedForAggregationSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsCollectedForAggregationSourceWithDefaults() *AccountsCollectedForAggregationSource {
	this := AccountsCollectedForAggregationSource{}
	return &this
}

// GetId returns the Id field value
func (o *AccountsCollectedForAggregationSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregationSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountsCollectedForAggregationSource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AccountsCollectedForAggregationSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregationSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountsCollectedForAggregationSource) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AccountsCollectedForAggregationSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregationSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountsCollectedForAggregationSource) SetName(v string) {
	o.Name = v
}

func (o AccountsCollectedForAggregationSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsCollectedForAggregationSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountsCollectedForAggregationSource) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAccountsCollectedForAggregationSource := _AccountsCollectedForAggregationSource{}

	if err = json.Unmarshal(bytes, &varAccountsCollectedForAggregationSource); err == nil {
	*o = AccountsCollectedForAggregationSource(varAccountsCollectedForAggregationSource)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountsCollectedForAggregationSource struct {
	value *AccountsCollectedForAggregationSource
	isSet bool
}

func (v NullableAccountsCollectedForAggregationSource) Get() *AccountsCollectedForAggregationSource {
	return v.value
}

func (v *NullableAccountsCollectedForAggregationSource) Set(val *AccountsCollectedForAggregationSource) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsCollectedForAggregationSource) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsCollectedForAggregationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsCollectedForAggregationSource(val *AccountsCollectedForAggregationSource) *NullableAccountsCollectedForAggregationSource {
	return &NullableAccountsCollectedForAggregationSource{value: val, isSet: true}
}

func (v NullableAccountsCollectedForAggregationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsCollectedForAggregationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


