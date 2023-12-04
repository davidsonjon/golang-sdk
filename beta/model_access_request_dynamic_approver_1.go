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

// checks if the AccessRequestDynamicApprover1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestDynamicApprover1{}

// AccessRequestDynamicApprover1 struct for AccessRequestDynamicApprover1
type AccessRequestDynamicApprover1 struct {
	// The unique ID of the identity to add to the approver list for the access request.
	Id string `json:"id"`
	// The name of the identity to add to the approver list for the access request.
	Name string `json:"name"`
	// The type of object being referenced.
	Type map[string]interface{} `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestDynamicApprover1 AccessRequestDynamicApprover1

// NewAccessRequestDynamicApprover1 instantiates a new AccessRequestDynamicApprover1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestDynamicApprover1(id string, name string, type_ map[string]interface{}) *AccessRequestDynamicApprover1 {
	this := AccessRequestDynamicApprover1{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewAccessRequestDynamicApprover1WithDefaults instantiates a new AccessRequestDynamicApprover1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestDynamicApprover1WithDefaults() *AccessRequestDynamicApprover1 {
	this := AccessRequestDynamicApprover1{}
	return &this
}

// GetId returns the Id field value
func (o *AccessRequestDynamicApprover1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestDynamicApprover1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccessRequestDynamicApprover1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessRequestDynamicApprover1) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AccessRequestDynamicApprover1) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover1) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *AccessRequestDynamicApprover1) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o AccessRequestDynamicApprover1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestDynamicApprover1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestDynamicApprover1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
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

	varAccessRequestDynamicApprover1 := _AccessRequestDynamicApprover1{}

	if err = json.Unmarshal(bytes, &varAccessRequestDynamicApprover1); err == nil {
	*o = AccessRequestDynamicApprover1(varAccessRequestDynamicApprover1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestDynamicApprover1 struct {
	value *AccessRequestDynamicApprover1
	isSet bool
}

func (v NullableAccessRequestDynamicApprover1) Get() *AccessRequestDynamicApprover1 {
	return v.value
}

func (v *NullableAccessRequestDynamicApprover1) Set(val *AccessRequestDynamicApprover1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestDynamicApprover1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestDynamicApprover1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestDynamicApprover1(val *AccessRequestDynamicApprover1) *NullableAccessRequestDynamicApprover1 {
	return &NullableAccessRequestDynamicApprover1{value: val, isSet: true}
}

func (v NullableAccessRequestDynamicApprover1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestDynamicApprover1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


