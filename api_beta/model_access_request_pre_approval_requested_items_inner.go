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

// checks if the AccessRequestPreApprovalRequestedItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPreApprovalRequestedItemsInner{}

// AccessRequestPreApprovalRequestedItemsInner struct for AccessRequestPreApprovalRequestedItemsInner
type AccessRequestPreApprovalRequestedItemsInner struct {
	// The unique ID of the access item being requested.
	Id string `json:"id"`
	// The human friendly name of the access item.
	Name string `json:"name"`
	// Detailed description of the access item.
	Description NullableString `json:"description,omitempty"`
	// The type of access item.
	Type map[string]interface{} `json:"type"`
	// The action to perform on the access item.
	Operation map[string]interface{} `json:"operation"`
	// A comment from the identity requesting the access.
	Comment NullableString `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPreApprovalRequestedItemsInner AccessRequestPreApprovalRequestedItemsInner

// NewAccessRequestPreApprovalRequestedItemsInner instantiates a new AccessRequestPreApprovalRequestedItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPreApprovalRequestedItemsInner(id string, name string, type_ map[string]interface{}, operation map[string]interface{}) *AccessRequestPreApprovalRequestedItemsInner {
	this := AccessRequestPreApprovalRequestedItemsInner{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Operation = operation
	return &this
}

// NewAccessRequestPreApprovalRequestedItemsInnerWithDefaults instantiates a new AccessRequestPreApprovalRequestedItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPreApprovalRequestedItemsInnerWithDefaults() *AccessRequestPreApprovalRequestedItemsInner {
	this := AccessRequestPreApprovalRequestedItemsInner{}
	return &this
}

// GetId returns the Id field value
func (o *AccessRequestPreApprovalRequestedItemsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApprovalRequestedItemsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestPreApprovalRequestedItemsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccessRequestPreApprovalRequestedItemsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApprovalRequestedItemsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessRequestPreApprovalRequestedItemsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPreApprovalRequestedItemsInner) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPreApprovalRequestedItemsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessRequestPreApprovalRequestedItemsInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessRequestPreApprovalRequestedItemsInner) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessRequestPreApprovalRequestedItemsInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessRequestPreApprovalRequestedItemsInner) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value
func (o *AccessRequestPreApprovalRequestedItemsInner) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApprovalRequestedItemsInner) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *AccessRequestPreApprovalRequestedItemsInner) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetOperation returns the Operation field value
func (o *AccessRequestPreApprovalRequestedItemsInner) GetOperation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApprovalRequestedItemsInner) GetOperationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Operation, true
}

// SetOperation sets field value
func (o *AccessRequestPreApprovalRequestedItemsInner) SetOperation(v map[string]interface{}) {
	o.Operation = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPreApprovalRequestedItemsInner) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPreApprovalRequestedItemsInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *AccessRequestPreApprovalRequestedItemsInner) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *AccessRequestPreApprovalRequestedItemsInner) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *AccessRequestPreApprovalRequestedItemsInner) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *AccessRequestPreApprovalRequestedItemsInner) UnsetComment() {
	o.Comment.Unset()
}

func (o AccessRequestPreApprovalRequestedItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPreApprovalRequestedItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["operation"] = o.Operation
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPreApprovalRequestedItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"operation",
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

	varAccessRequestPreApprovalRequestedItemsInner := _AccessRequestPreApprovalRequestedItemsInner{}

	err = json.Unmarshal(data, &varAccessRequestPreApprovalRequestedItemsInner)

	if err != nil {
		return err
	}

	*o = AccessRequestPreApprovalRequestedItemsInner(varAccessRequestPreApprovalRequestedItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPreApprovalRequestedItemsInner struct {
	value *AccessRequestPreApprovalRequestedItemsInner
	isSet bool
}

func (v NullableAccessRequestPreApprovalRequestedItemsInner) Get() *AccessRequestPreApprovalRequestedItemsInner {
	return v.value
}

func (v *NullableAccessRequestPreApprovalRequestedItemsInner) Set(val *AccessRequestPreApprovalRequestedItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPreApprovalRequestedItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPreApprovalRequestedItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPreApprovalRequestedItemsInner(val *AccessRequestPreApprovalRequestedItemsInner) *NullableAccessRequestPreApprovalRequestedItemsInner {
	return &NullableAccessRequestPreApprovalRequestedItemsInner{value: val, isSet: true}
}

func (v NullableAccessRequestPreApprovalRequestedItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPreApprovalRequestedItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


