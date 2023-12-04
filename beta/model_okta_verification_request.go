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

// checks if the OktaVerificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaVerificationRequest{}

// OktaVerificationRequest struct for OktaVerificationRequest
type OktaVerificationRequest struct {
	// User identifier for Verification request. The value of the user's attribute.
	UserId string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _OktaVerificationRequest OktaVerificationRequest

// NewOktaVerificationRequest instantiates a new OktaVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaVerificationRequest(userId string) *OktaVerificationRequest {
	this := OktaVerificationRequest{}
	this.UserId = userId
	return &this
}

// NewOktaVerificationRequestWithDefaults instantiates a new OktaVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaVerificationRequestWithDefaults() *OktaVerificationRequest {
	this := OktaVerificationRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *OktaVerificationRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OktaVerificationRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *OktaVerificationRequest) SetUserId(v string) {
	o.UserId = v
}

func (o OktaVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaVerificationRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
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

	varOktaVerificationRequest := _OktaVerificationRequest{}

	if err = json.Unmarshal(bytes, &varOktaVerificationRequest); err == nil {
	*o = OktaVerificationRequest(varOktaVerificationRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaVerificationRequest struct {
	value *OktaVerificationRequest
	isSet bool
}

func (v NullableOktaVerificationRequest) Get() *OktaVerificationRequest {
	return v.value
}

func (v *NullableOktaVerificationRequest) Set(val *OktaVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaVerificationRequest(val *OktaVerificationRequest) *NullableOktaVerificationRequest {
	return &NullableOktaVerificationRequest{value: val, isSet: true}
}

func (v NullableOktaVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


