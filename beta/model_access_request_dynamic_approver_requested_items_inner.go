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

// checks if the AccessRequestDynamicApproverRequestedItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestDynamicApproverRequestedItemsInner{}

// AccessRequestDynamicApproverRequestedItemsInner struct for AccessRequestDynamicApproverRequestedItemsInner
type AccessRequestDynamicApproverRequestedItemsInner struct {
	// The unique ID of the access item.
	Id string `json:"id"`
	// Human friendly name of the access item.
	Name string `json:"name"`
	// Extended description of the access item.
	Description NullableString `json:"description,omitempty"`
	// The type of access item being requested.
	Type map[string]interface{} `json:"type"`
	// Grant or revoke the access item
	Operation map[string]interface{} `json:"operation"`
	// A comment from the requestor on why the access is needed.
	Comment NullableString `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestDynamicApproverRequestedItemsInner AccessRequestDynamicApproverRequestedItemsInner

// NewAccessRequestDynamicApproverRequestedItemsInner instantiates a new AccessRequestDynamicApproverRequestedItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestDynamicApproverRequestedItemsInner(id string, name string, type_ map[string]interface{}, operation map[string]interface{}) *AccessRequestDynamicApproverRequestedItemsInner {
	this := AccessRequestDynamicApproverRequestedItemsInner{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Operation = operation
	return &this
}

// NewAccessRequestDynamicApproverRequestedItemsInnerWithDefaults instantiates a new AccessRequestDynamicApproverRequestedItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestDynamicApproverRequestedItemsInnerWithDefaults() *AccessRequestDynamicApproverRequestedItemsInner {
	this := AccessRequestDynamicApproverRequestedItemsInner{}
	return &this
}

// GetId returns the Id field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessRequestDynamicApproverRequestedItemsInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessRequestDynamicApproverRequestedItemsInner) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetOperation returns the Operation field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetOperation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetOperationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Operation, true
}

// SetOperation sets field value
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetOperation(v map[string]interface{}) {
	o.Operation = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetComment() string {
	if o == nil || isNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestDynamicApproverRequestedItemsInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *AccessRequestDynamicApproverRequestedItemsInner) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *AccessRequestDynamicApproverRequestedItemsInner) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *AccessRequestDynamicApproverRequestedItemsInner) UnsetComment() {
	o.Comment.Unset()
}

func (o AccessRequestDynamicApproverRequestedItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestDynamicApproverRequestedItemsInner) ToMap() (map[string]interface{}, error) {
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

func (o *AccessRequestDynamicApproverRequestedItemsInner) UnmarshalJSON(bytes []byte) (err error) {
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

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessRequestDynamicApproverRequestedItemsInner := _AccessRequestDynamicApproverRequestedItemsInner{}

	if err = json.Unmarshal(bytes, &varAccessRequestDynamicApproverRequestedItemsInner); err == nil {
	*o = AccessRequestDynamicApproverRequestedItemsInner(varAccessRequestDynamicApproverRequestedItemsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

type NullableAccessRequestDynamicApproverRequestedItemsInner struct {
	value *AccessRequestDynamicApproverRequestedItemsInner
	isSet bool
}

func (v NullableAccessRequestDynamicApproverRequestedItemsInner) Get() *AccessRequestDynamicApproverRequestedItemsInner {
	return v.value
}

func (v *NullableAccessRequestDynamicApproverRequestedItemsInner) Set(val *AccessRequestDynamicApproverRequestedItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestDynamicApproverRequestedItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestDynamicApproverRequestedItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestDynamicApproverRequestedItemsInner(val *AccessRequestDynamicApproverRequestedItemsInner) *NullableAccessRequestDynamicApproverRequestedItemsInner {
	return &NullableAccessRequestDynamicApproverRequestedItemsInner{value: val, isSet: true}
}

func (v NullableAccessRequestDynamicApproverRequestedItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestDynamicApproverRequestedItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


