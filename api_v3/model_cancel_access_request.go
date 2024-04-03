/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the CancelAccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelAccessRequest{}

// CancelAccessRequest Request body payload for cancel access request endpoint.
type CancelAccessRequest struct {
	// This refers to the identityRequestId. To successfully cancel an access request, you must provide the identityRequestId.
	AccountActivityId string `json:"accountActivityId"`
	// Reason for cancelling the pending access request.
	Comment string `json:"comment"`
	AdditionalProperties map[string]interface{}
}

type _CancelAccessRequest CancelAccessRequest

// NewCancelAccessRequest instantiates a new CancelAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAccessRequest(accountActivityId string, comment string) *CancelAccessRequest {
	this := CancelAccessRequest{}
	this.AccountActivityId = accountActivityId
	this.Comment = comment
	return &this
}

// NewCancelAccessRequestWithDefaults instantiates a new CancelAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAccessRequestWithDefaults() *CancelAccessRequest {
	this := CancelAccessRequest{}
	return &this
}

// GetAccountActivityId returns the AccountActivityId field value
func (o *CancelAccessRequest) GetAccountActivityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountActivityId
}

// GetAccountActivityIdOk returns a tuple with the AccountActivityId field value
// and a boolean to check if the value has been set.
func (o *CancelAccessRequest) GetAccountActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountActivityId, true
}

// SetAccountActivityId sets field value
func (o *CancelAccessRequest) SetAccountActivityId(v string) {
	o.AccountActivityId = v
}

// GetComment returns the Comment field value
func (o *CancelAccessRequest) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *CancelAccessRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *CancelAccessRequest) SetComment(v string) {
	o.Comment = v
}

func (o CancelAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountActivityId"] = o.AccountActivityId
	toSerialize["comment"] = o.Comment

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelAccessRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountActivityId",
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

	varCancelAccessRequest := _CancelAccessRequest{}

	if err = json.Unmarshal(bytes, &varCancelAccessRequest); err == nil {
	*o = CancelAccessRequest(varCancelAccessRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountActivityId")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAccessRequest struct {
	value *CancelAccessRequest
	isSet bool
}

func (v NullableCancelAccessRequest) Get() *CancelAccessRequest {
	return v.value
}

func (v *NullableCancelAccessRequest) Set(val *CancelAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAccessRequest(val *CancelAccessRequest) *NullableCancelAccessRequest {
	return &NullableCancelAccessRequest{value: val, isSet: true}
}

func (v NullableCancelAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


