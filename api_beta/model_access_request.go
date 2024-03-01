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

// checks if the AccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequest{}

// AccessRequest struct for AccessRequest
type AccessRequest struct {
	// A list of Identity IDs for whom the Access is requested. If it's a Revoke request, there can only be one Identity ID.
	RequestedFor []string `json:"requestedFor"`
	RequestType *AccessRequestType `json:"requestType,omitempty"`
	RequestedItems []AccessRequestItem `json:"requestedItems"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on associated APIs such as /account-activities.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequest AccessRequest

// NewAccessRequest instantiates a new AccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequest(requestedFor []string, requestedItems []AccessRequestItem) *AccessRequest {
	this := AccessRequest{}
	this.RequestedFor = requestedFor
	this.RequestedItems = requestedItems
	return &this
}

// NewAccessRequestWithDefaults instantiates a new AccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestWithDefaults() *AccessRequest {
	this := AccessRequest{}
	return &this
}

// GetRequestedFor returns the RequestedFor field value
func (o *AccessRequest) GetRequestedFor() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestedForOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *AccessRequest) SetRequestedFor(v []string) {
	o.RequestedFor = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *AccessRequest) GetRequestType() AccessRequestType {
	if o == nil || isNil(o.RequestType) {
		var ret AccessRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestTypeOk() (*AccessRequestType, bool) {
	if o == nil || isNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *AccessRequest) HasRequestType() bool {
	if o != nil && !isNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given AccessRequestType and assigns it to the RequestType field.
func (o *AccessRequest) SetRequestType(v AccessRequestType) {
	o.RequestType = &v
}

// GetRequestedItems returns the RequestedItems field value
func (o *AccessRequest) GetRequestedItems() []AccessRequestItem {
	if o == nil {
		var ret []AccessRequestItem
		return ret
	}

	return o.RequestedItems
}

// GetRequestedItemsOk returns a tuple with the RequestedItems field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestedItemsOk() ([]AccessRequestItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedItems, true
}

// SetRequestedItems sets field value
func (o *AccessRequest) SetRequestedItems(v []AccessRequestItem) {
	o.RequestedItems = v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *AccessRequest) GetClientMetadata() map[string]string {
	if o == nil || isNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *AccessRequest) HasClientMetadata() bool {
	if o != nil && !isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *AccessRequest) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

func (o AccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestedFor"] = o.RequestedFor
	if !isNil(o.RequestType) {
		toSerialize["requestType"] = o.RequestType
	}
	toSerialize["requestedItems"] = o.RequestedItems
	if !isNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestedFor",
		"requestedItems",
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

	varAccessRequest := _AccessRequest{}

	if err = json.Unmarshal(bytes, &varAccessRequest); err == nil {
	*o = AccessRequest(varAccessRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "requestedItems")
		delete(additionalProperties, "clientMetadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequest struct {
	value *AccessRequest
	isSet bool
}

func (v NullableAccessRequest) Get() *AccessRequest {
	return v.value
}

func (v *NullableAccessRequest) Set(val *AccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequest(val *AccessRequest) *NullableAccessRequest {
	return &NullableAccessRequest{value: val, isSet: true}
}

func (v NullableAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

