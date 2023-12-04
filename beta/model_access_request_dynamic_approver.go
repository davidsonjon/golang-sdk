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

// checks if the AccessRequestDynamicApprover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestDynamicApprover{}

// AccessRequestDynamicApprover struct for AccessRequestDynamicApprover
type AccessRequestDynamicApprover struct {
	// The unique ID of the access request object. Can be used with the [access request status endpoint](https://developer.sailpoint.com/idn/api/beta/list-access-request-status) to get the status of the request. 
	AccessRequestId string `json:"accessRequestId"`
	RequestedFor AccessItemRequestedForDto `json:"requestedFor"`
	// The access items that are being requested.
	RequestedItems []AccessRequestDynamicApproverRequestedItemsInner `json:"requestedItems"`
	RequestedBy AccessItemRequesterDto `json:"requestedBy"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestDynamicApprover AccessRequestDynamicApprover

// NewAccessRequestDynamicApprover instantiates a new AccessRequestDynamicApprover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestDynamicApprover(accessRequestId string, requestedFor AccessItemRequestedForDto, requestedItems []AccessRequestDynamicApproverRequestedItemsInner, requestedBy AccessItemRequesterDto) *AccessRequestDynamicApprover {
	this := AccessRequestDynamicApprover{}
	this.AccessRequestId = accessRequestId
	this.RequestedFor = requestedFor
	this.RequestedItems = requestedItems
	this.RequestedBy = requestedBy
	return &this
}

// NewAccessRequestDynamicApproverWithDefaults instantiates a new AccessRequestDynamicApprover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestDynamicApproverWithDefaults() *AccessRequestDynamicApprover {
	this := AccessRequestDynamicApprover{}
	return &this
}

// GetAccessRequestId returns the AccessRequestId field value
func (o *AccessRequestDynamicApprover) GetAccessRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessRequestId
}

// GetAccessRequestIdOk returns a tuple with the AccessRequestId field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover) GetAccessRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessRequestId, true
}

// SetAccessRequestId sets field value
func (o *AccessRequestDynamicApprover) SetAccessRequestId(v string) {
	o.AccessRequestId = v
}

// GetRequestedFor returns the RequestedFor field value
func (o *AccessRequestDynamicApprover) GetRequestedFor() AccessItemRequestedForDto {
	if o == nil {
		var ret AccessItemRequestedForDto
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover) GetRequestedForOk() (*AccessItemRequestedForDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *AccessRequestDynamicApprover) SetRequestedFor(v AccessItemRequestedForDto) {
	o.RequestedFor = v
}

// GetRequestedItems returns the RequestedItems field value
func (o *AccessRequestDynamicApprover) GetRequestedItems() []AccessRequestDynamicApproverRequestedItemsInner {
	if o == nil {
		var ret []AccessRequestDynamicApproverRequestedItemsInner
		return ret
	}

	return o.RequestedItems
}

// GetRequestedItemsOk returns a tuple with the RequestedItems field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover) GetRequestedItemsOk() ([]AccessRequestDynamicApproverRequestedItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedItems, true
}

// SetRequestedItems sets field value
func (o *AccessRequestDynamicApprover) SetRequestedItems(v []AccessRequestDynamicApproverRequestedItemsInner) {
	o.RequestedItems = v
}

// GetRequestedBy returns the RequestedBy field value
func (o *AccessRequestDynamicApprover) GetRequestedBy() AccessItemRequesterDto {
	if o == nil {
		var ret AccessItemRequesterDto
		return ret
	}

	return o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value
// and a boolean to check if the value has been set.
func (o *AccessRequestDynamicApprover) GetRequestedByOk() (*AccessItemRequesterDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedBy, true
}

// SetRequestedBy sets field value
func (o *AccessRequestDynamicApprover) SetRequestedBy(v AccessItemRequesterDto) {
	o.RequestedBy = v
}

func (o AccessRequestDynamicApprover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestDynamicApprover) ToMap() (map[string]interface{}, error) {
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

func (o *AccessRequestDynamicApprover) UnmarshalJSON(bytes []byte) (err error) {
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

	varAccessRequestDynamicApprover := _AccessRequestDynamicApprover{}

	if err = json.Unmarshal(bytes, &varAccessRequestDynamicApprover); err == nil {
	*o = AccessRequestDynamicApprover(varAccessRequestDynamicApprover)
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

type NullableAccessRequestDynamicApprover struct {
	value *AccessRequestDynamicApprover
	isSet bool
}

func (v NullableAccessRequestDynamicApprover) Get() *AccessRequestDynamicApprover {
	return v.value
}

func (v *NullableAccessRequestDynamicApprover) Set(val *AccessRequestDynamicApprover) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestDynamicApprover) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestDynamicApprover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestDynamicApprover(val *AccessRequestDynamicApprover) *NullableAccessRequestDynamicApprover {
	return &NullableAccessRequestDynamicApprover{value: val, isSet: true}
}

func (v NullableAccessRequestDynamicApprover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestDynamicApprover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


