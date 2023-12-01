/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the PasswordDigitToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordDigitToken{}

// PasswordDigitToken struct for PasswordDigitToken
type PasswordDigitToken struct {
	// The digit token for password management
	DigitToken *string `json:"digitToken,omitempty"`
	// The reference ID of the digit token generation request
	RequestId *string `json:"requestId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordDigitToken PasswordDigitToken

// NewPasswordDigitToken instantiates a new PasswordDigitToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordDigitToken() *PasswordDigitToken {
	this := PasswordDigitToken{}
	return &this
}

// NewPasswordDigitTokenWithDefaults instantiates a new PasswordDigitToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordDigitTokenWithDefaults() *PasswordDigitToken {
	this := PasswordDigitToken{}
	return &this
}

// GetDigitToken returns the DigitToken field value if set, zero value otherwise.
func (o *PasswordDigitToken) GetDigitToken() string {
	if o == nil || isNil(o.DigitToken) {
		var ret string
		return ret
	}
	return *o.DigitToken
}

// GetDigitTokenOk returns a tuple with the DigitToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDigitToken) GetDigitTokenOk() (*string, bool) {
	if o == nil || isNil(o.DigitToken) {
		return nil, false
	}
	return o.DigitToken, true
}

// HasDigitToken returns a boolean if a field has been set.
func (o *PasswordDigitToken) HasDigitToken() bool {
	if o != nil && !isNil(o.DigitToken) {
		return true
	}

	return false
}

// SetDigitToken gets a reference to the given string and assigns it to the DigitToken field.
func (o *PasswordDigitToken) SetDigitToken(v string) {
	o.DigitToken = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PasswordDigitToken) GetRequestId() string {
	if o == nil || isNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDigitToken) GetRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PasswordDigitToken) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PasswordDigitToken) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PasswordDigitToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordDigitToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DigitToken) {
		toSerialize["digitToken"] = o.DigitToken
	}
	if !isNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordDigitToken) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordDigitToken := _PasswordDigitToken{}

	if err = json.Unmarshal(bytes, &varPasswordDigitToken); err == nil {
	*o = PasswordDigitToken(varPasswordDigitToken)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "digitToken")
		delete(additionalProperties, "requestId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordDigitToken struct {
	value *PasswordDigitToken
	isSet bool
}

func (v NullablePasswordDigitToken) Get() *PasswordDigitToken {
	return v.value
}

func (v *NullablePasswordDigitToken) Set(val *PasswordDigitToken) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordDigitToken) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordDigitToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordDigitToken(val *PasswordDigitToken) *NullablePasswordDigitToken {
	return &NullablePasswordDigitToken{value: val, isSet: true}
}

func (v NullablePasswordDigitToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordDigitToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


