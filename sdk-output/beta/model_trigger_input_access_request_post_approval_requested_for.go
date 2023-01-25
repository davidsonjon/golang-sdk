/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// TriggerInputAccessRequestPostApprovalRequestedFor The identity who the access request is for.
type TriggerInputAccessRequestPostApprovalRequestedFor struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type"`
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccessRequestPostApprovalRequestedFor TriggerInputAccessRequestPostApprovalRequestedFor

// NewTriggerInputAccessRequestPostApprovalRequestedFor instantiates a new TriggerInputAccessRequestPostApprovalRequestedFor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccessRequestPostApprovalRequestedFor(type_ map[string]interface{}, id string, name string) *TriggerInputAccessRequestPostApprovalRequestedFor {
	this := TriggerInputAccessRequestPostApprovalRequestedFor{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewTriggerInputAccessRequestPostApprovalRequestedForWithDefaults instantiates a new TriggerInputAccessRequestPostApprovalRequestedFor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccessRequestPostApprovalRequestedForWithDefaults() *TriggerInputAccessRequestPostApprovalRequestedFor {
	this := TriggerInputAccessRequestPostApprovalRequestedFor{}
	return &this
}

// GetType returns the Type field value
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedFor) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccessRequestPostApprovalRequestedFor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccessRequestPostApprovalRequestedFor) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccessRequestPostApprovalRequestedFor := _TriggerInputAccessRequestPostApprovalRequestedFor{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccessRequestPostApprovalRequestedFor); err == nil {
		*o = TriggerInputAccessRequestPostApprovalRequestedFor(varTriggerInputAccessRequestPostApprovalRequestedFor)
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

type NullableTriggerInputAccessRequestPostApprovalRequestedFor struct {
	value *TriggerInputAccessRequestPostApprovalRequestedFor
	isSet bool
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedFor) Get() *TriggerInputAccessRequestPostApprovalRequestedFor {
	return v.value
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedFor) Set(val *TriggerInputAccessRequestPostApprovalRequestedFor) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedFor) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedFor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccessRequestPostApprovalRequestedFor(val *TriggerInputAccessRequestPostApprovalRequestedFor) *NullableTriggerInputAccessRequestPostApprovalRequestedFor {
	return &NullableTriggerInputAccessRequestPostApprovalRequestedFor{value: val, isSet: true}
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedFor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedFor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


