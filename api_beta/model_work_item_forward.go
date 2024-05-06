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

// checks if the WorkItemForward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemForward{}

// WorkItemForward struct for WorkItemForward
type WorkItemForward struct {
	// The ID of the identity to forward this work item to.
	TargetOwnerId string `json:"targetOwnerId"`
	// Comments to send to the target owner
	Comment string `json:"comment"`
	// If true, send a notification to the target owner.
	SendNotifications *bool `json:"sendNotifications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkItemForward WorkItemForward

// NewWorkItemForward instantiates a new WorkItemForward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemForward(targetOwnerId string, comment string) *WorkItemForward {
	this := WorkItemForward{}
	this.TargetOwnerId = targetOwnerId
	this.Comment = comment
	var sendNotifications bool = true
	this.SendNotifications = &sendNotifications
	return &this
}

// NewWorkItemForwardWithDefaults instantiates a new WorkItemForward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemForwardWithDefaults() *WorkItemForward {
	this := WorkItemForward{}
	var sendNotifications bool = true
	this.SendNotifications = &sendNotifications
	return &this
}

// GetTargetOwnerId returns the TargetOwnerId field value
func (o *WorkItemForward) GetTargetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetOwnerId
}

// GetTargetOwnerIdOk returns a tuple with the TargetOwnerId field value
// and a boolean to check if the value has been set.
func (o *WorkItemForward) GetTargetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetOwnerId, true
}

// SetTargetOwnerId sets field value
func (o *WorkItemForward) SetTargetOwnerId(v string) {
	o.TargetOwnerId = v
}

// GetComment returns the Comment field value
func (o *WorkItemForward) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *WorkItemForward) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *WorkItemForward) SetComment(v string) {
	o.Comment = v
}

// GetSendNotifications returns the SendNotifications field value if set, zero value otherwise.
func (o *WorkItemForward) GetSendNotifications() bool {
	if o == nil || isNil(o.SendNotifications) {
		var ret bool
		return ret
	}
	return *o.SendNotifications
}

// GetSendNotificationsOk returns a tuple with the SendNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemForward) GetSendNotificationsOk() (*bool, bool) {
	if o == nil || isNil(o.SendNotifications) {
		return nil, false
	}
	return o.SendNotifications, true
}

// HasSendNotifications returns a boolean if a field has been set.
func (o *WorkItemForward) HasSendNotifications() bool {
	if o != nil && !isNil(o.SendNotifications) {
		return true
	}

	return false
}

// SetSendNotifications gets a reference to the given bool and assigns it to the SendNotifications field.
func (o *WorkItemForward) SetSendNotifications(v bool) {
	o.SendNotifications = &v
}

func (o WorkItemForward) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemForward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetOwnerId"] = o.TargetOwnerId
	toSerialize["comment"] = o.Comment
	if !isNil(o.SendNotifications) {
		toSerialize["sendNotifications"] = o.SendNotifications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkItemForward) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetOwnerId",
		"comment",
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

	varWorkItemForward := _WorkItemForward{}

	if err = json.Unmarshal(bytes, &varWorkItemForward); err == nil {
	*o = WorkItemForward(varWorkItemForward)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "targetOwnerId")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "sendNotifications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkItemForward struct {
	value *WorkItemForward
	isSet bool
}

func (v NullableWorkItemForward) Get() *WorkItemForward {
	return v.value
}

func (v *NullableWorkItemForward) Set(val *WorkItemForward) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemForward) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemForward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemForward(val *WorkItemForward) *NullableWorkItemForward {
	return &NullableWorkItemForward{value: val, isSet: true}
}

func (v NullableWorkItemForward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemForward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


