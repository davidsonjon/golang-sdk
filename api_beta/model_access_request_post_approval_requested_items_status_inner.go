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

// checks if the AccessRequestPostApprovalRequestedItemsStatusInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPostApprovalRequestedItemsStatusInner{}

// AccessRequestPostApprovalRequestedItemsStatusInner struct for AccessRequestPostApprovalRequestedItemsStatusInner
type AccessRequestPostApprovalRequestedItemsStatusInner struct {
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
	// Additional customer defined metadata about the access item.
	ClientMetadata map[string]interface{} `json:"clientMetadata,omitempty"`
	// A list of one or more approvers for the access request.
	ApprovalInfo []AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner `json:"approvalInfo"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPostApprovalRequestedItemsStatusInner AccessRequestPostApprovalRequestedItemsStatusInner

// NewAccessRequestPostApprovalRequestedItemsStatusInner instantiates a new AccessRequestPostApprovalRequestedItemsStatusInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPostApprovalRequestedItemsStatusInner(id string, name string, type_ map[string]interface{}, operation map[string]interface{}, approvalInfo []AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) *AccessRequestPostApprovalRequestedItemsStatusInner {
	this := AccessRequestPostApprovalRequestedItemsStatusInner{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Operation = operation
	this.ApprovalInfo = approvalInfo
	return &this
}

// NewAccessRequestPostApprovalRequestedItemsStatusInnerWithDefaults instantiates a new AccessRequestPostApprovalRequestedItemsStatusInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPostApprovalRequestedItemsStatusInnerWithDefaults() *AccessRequestPostApprovalRequestedItemsStatusInner {
	this := AccessRequestPostApprovalRequestedItemsStatusInner{}
	return &this
}

// GetId returns the Id field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetOperation returns the Operation field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetOperation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetOperationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Operation, true
}

// SetOperation sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetOperation(v map[string]interface{}) {
	o.Operation = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) UnsetComment() {
	o.Comment.Unset()
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetClientMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetClientMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ClientMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) HasClientMetadata() bool {
	if o != nil && !IsNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]interface{} and assigns it to the ClientMetadata field.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetClientMetadata(v map[string]interface{}) {
	o.ClientMetadata = v
}

// GetApprovalInfo returns the ApprovalInfo field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetApprovalInfo() []AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner {
	if o == nil {
		var ret []AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner
		return ret
	}

	return o.ApprovalInfo
}

// GetApprovalInfoOk returns a tuple with the ApprovalInfo field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) GetApprovalInfoOk() ([]AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalInfo, true
}

// SetApprovalInfo sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInner) SetApprovalInfo(v []AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) {
	o.ApprovalInfo = v
}

func (o AccessRequestPostApprovalRequestedItemsStatusInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPostApprovalRequestedItemsStatusInner) ToMap() (map[string]interface{}, error) {
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
	if o.ClientMetadata != nil {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	toSerialize["approvalInfo"] = o.ApprovalInfo

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPostApprovalRequestedItemsStatusInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"operation",
		"approvalInfo",
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

	varAccessRequestPostApprovalRequestedItemsStatusInner := _AccessRequestPostApprovalRequestedItemsStatusInner{}

	err = json.Unmarshal(data, &varAccessRequestPostApprovalRequestedItemsStatusInner)

	if err != nil {
		return err
	}

	*o = AccessRequestPostApprovalRequestedItemsStatusInner(varAccessRequestPostApprovalRequestedItemsStatusInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "approvalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPostApprovalRequestedItemsStatusInner struct {
	value *AccessRequestPostApprovalRequestedItemsStatusInner
	isSet bool
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInner) Get() *AccessRequestPostApprovalRequestedItemsStatusInner {
	return v.value
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInner) Set(val *AccessRequestPostApprovalRequestedItemsStatusInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPostApprovalRequestedItemsStatusInner(val *AccessRequestPostApprovalRequestedItemsStatusInner) *NullableAccessRequestPostApprovalRequestedItemsStatusInner {
	return &NullableAccessRequestPostApprovalRequestedItemsStatusInner{value: val, isSet: true}
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


