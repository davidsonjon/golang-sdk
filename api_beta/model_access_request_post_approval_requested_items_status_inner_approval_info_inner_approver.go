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

// checks if the AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}

// AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover The identity of the approver.
type AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type"`
	// ID of identity who approved the access item request.
	Id string `json:"id"`
	// Human-readable display name of identity who approved the access item request.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover

// NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover instantiates a new AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover(type_ map[string]interface{}, id string, name string) *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	this := AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApproverWithDefaults instantiates a new AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApproverWithDefaults() *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	this := AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}
	return &this
}

// GetType returns the Type field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) SetName(v string) {
	o.Name = v
}

func (o AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) UnmarshalJSON(data []byte) (err error) {
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

	varAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover := _AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}

	err = json.Unmarshal(data, &varAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover)

	if err != nil {
		return err
	}

	*o = AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover(varAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover struct {
	value *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover
	isSet bool
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) Get() *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	return v.value
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) Set(val *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover(val *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	return &NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{value: val, isSet: true}
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


