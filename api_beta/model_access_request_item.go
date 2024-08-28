/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the AccessRequestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestItem{}

// AccessRequestItem struct for AccessRequestItem
type AccessRequestItem struct {
	// The type of the item being requested.
	Type string `json:"type"`
	// ID of Role, Access Profile or Entitlement being requested.
	Id string `json:"id"`
	// Comment provided by requester. * Comment is required when the request is of type Revoke Access. 
	Comment *string `json:"comment,omitempty"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on associated APIs such as /account-activities and /access-request-status.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	// The date the role or access profile or entitlement is no longer assigned to the specified identity. Also known as the expiration date. * Specify a date in the future. * The current SLA for the deprovisioning is 24 hours. * This date can be modified to either extend or decrease the duration of access item assignments for the specified identity. You can change the expiration date for requests for yourself or direct reports, but you cannot remove an expiration date on an already approved item. If the access request has not been approved, you can cancel it and submit a new one without the expiration. If it has already been approved, then you have to revoke the access and then re-request without the expiration. 
	RemoveDate *time.Time `json:"removeDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestItem AccessRequestItem

// NewAccessRequestItem instantiates a new AccessRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestItem(type_ string, id string) *AccessRequestItem {
	this := AccessRequestItem{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAccessRequestItemWithDefaults instantiates a new AccessRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestItemWithDefaults() *AccessRequestItem {
	this := AccessRequestItem{}
	return &this
}

// GetType returns the Type field value
func (o *AccessRequestItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessRequestItem) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AccessRequestItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestItem) SetId(v string) {
	o.Id = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AccessRequestItem) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AccessRequestItem) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AccessRequestItem) SetComment(v string) {
	o.Comment = &v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *AccessRequestItem) GetClientMetadata() map[string]string {
	if o == nil || IsNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *AccessRequestItem) HasClientMetadata() bool {
	if o != nil && !IsNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *AccessRequestItem) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise.
func (o *AccessRequestItem) GetRemoveDate() time.Time {
	if o == nil || IsNil(o.RemoveDate) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RemoveDate) {
		return nil, false
	}
	return o.RemoveDate, true
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *AccessRequestItem) HasRemoveDate() bool {
	if o != nil && !IsNil(o.RemoveDate) {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given time.Time and assigns it to the RemoveDate field.
func (o *AccessRequestItem) SetRemoveDate(v time.Time) {
	o.RemoveDate = &v
}

func (o AccessRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if !IsNil(o.RemoveDate) {
		toSerialize["removeDate"] = o.RemoveDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varAccessRequestItem := _AccessRequestItem{}

	err = json.Unmarshal(data, &varAccessRequestItem)

	if err != nil {
		return err
	}

	*o = AccessRequestItem(varAccessRequestItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "removeDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestItem struct {
	value *AccessRequestItem
	isSet bool
}

func (v NullableAccessRequestItem) Get() *AccessRequestItem {
	return v.value
}

func (v *NullableAccessRequestItem) Set(val *AccessRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestItem(val *AccessRequestItem) *NullableAccessRequestItem {
	return &NullableAccessRequestItem{value: val, isSet: true}
}

func (v NullableAccessRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


