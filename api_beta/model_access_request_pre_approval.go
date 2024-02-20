/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the AccessRequestPreApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPreApproval{}

// AccessRequestPreApproval struct for AccessRequestPreApproval
type AccessRequestPreApproval struct {
	// The unique ID of the access request.
	AccessRequestId string `json:"accessRequestId"`
	// Identities access was requested for.
	RequestedFor []AccessItemRequestedForDto `json:"requestedFor"`
	// Details of the access items being requested.
	RequestedItems []AccessRequestPreApprovalRequestedItemsInner `json:"requestedItems"`
	RequestedBy AccessItemRequesterDto `json:"requestedBy"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPreApproval AccessRequestPreApproval

// NewAccessRequestPreApproval instantiates a new AccessRequestPreApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPreApproval(accessRequestId string, requestedFor []AccessItemRequestedForDto, requestedItems []AccessRequestPreApprovalRequestedItemsInner, requestedBy AccessItemRequesterDto) *AccessRequestPreApproval {
	this := AccessRequestPreApproval{}
	this.AccessRequestId = accessRequestId
	this.RequestedFor = requestedFor
	this.RequestedItems = requestedItems
	this.RequestedBy = requestedBy
	return &this
}

// NewAccessRequestPreApprovalWithDefaults instantiates a new AccessRequestPreApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPreApprovalWithDefaults() *AccessRequestPreApproval {
	this := AccessRequestPreApproval{}
	return &this
}

// GetAccessRequestId returns the AccessRequestId field value
func (o *AccessRequestPreApproval) GetAccessRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessRequestId
}

// GetAccessRequestIdOk returns a tuple with the AccessRequestId field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval) GetAccessRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessRequestId, true
}

// SetAccessRequestId sets field value
func (o *AccessRequestPreApproval) SetAccessRequestId(v string) {
	o.AccessRequestId = v
}

// GetRequestedFor returns the RequestedFor field value
func (o *AccessRequestPreApproval) GetRequestedFor() []AccessItemRequestedForDto {
	if o == nil {
		var ret []AccessItemRequestedForDto
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval) GetRequestedForOk() ([]AccessItemRequestedForDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *AccessRequestPreApproval) SetRequestedFor(v []AccessItemRequestedForDto) {
	o.RequestedFor = v
}

// GetRequestedItems returns the RequestedItems field value
func (o *AccessRequestPreApproval) GetRequestedItems() []AccessRequestPreApprovalRequestedItemsInner {
	if o == nil {
		var ret []AccessRequestPreApprovalRequestedItemsInner
		return ret
	}

	return o.RequestedItems
}

// GetRequestedItemsOk returns a tuple with the RequestedItems field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval) GetRequestedItemsOk() ([]AccessRequestPreApprovalRequestedItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedItems, true
}

// SetRequestedItems sets field value
func (o *AccessRequestPreApproval) SetRequestedItems(v []AccessRequestPreApprovalRequestedItemsInner) {
	o.RequestedItems = v
}

// GetRequestedBy returns the RequestedBy field value
func (o *AccessRequestPreApproval) GetRequestedBy() AccessItemRequesterDto {
	if o == nil {
		var ret AccessItemRequesterDto
		return ret
	}

	return o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval) GetRequestedByOk() (*AccessItemRequesterDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedBy, true
}

// SetRequestedBy sets field value
func (o *AccessRequestPreApproval) SetRequestedBy(v AccessItemRequesterDto) {
	o.RequestedBy = v
}

func (o AccessRequestPreApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPreApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessRequestId"] = o.AccessRequestId
	toSerialize["requestedFor"] = o.RequestedFor
	toSerialize["requestedItems"] = o.RequestedItems
	toSerialize["requestedBy"] = o.RequestedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPreApproval) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessRequestId",
		"requestedFor",
		"requestedItems",
		"requestedBy",
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

	varAccessRequestPreApproval := _AccessRequestPreApproval{}

	if err = json.Unmarshal(bytes, &varAccessRequestPreApproval); err == nil {
	*o = AccessRequestPreApproval(varAccessRequestPreApproval)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessRequestId")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requestedItems")
		delete(additionalProperties, "requestedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPreApproval struct {
	value *AccessRequestPreApproval
	isSet bool
}

func (v NullableAccessRequestPreApproval) Get() *AccessRequestPreApproval {
	return v.value
}

func (v *NullableAccessRequestPreApproval) Set(val *AccessRequestPreApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPreApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPreApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPreApproval(val *AccessRequestPreApproval) *NullableAccessRequestPreApproval {
	return &NullableAccessRequestPreApproval{value: val, isSet: true}
}

func (v NullableAccessRequestPreApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPreApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


