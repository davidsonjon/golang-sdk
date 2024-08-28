/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the VerificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationResponse{}

// VerificationResponse struct for VerificationResponse
type VerificationResponse struct {
	// The verificationPollRequest request ID
	RequestId NullableString `json:"requestId,omitempty"`
	// MFA Authentication status
	Status *string `json:"status,omitempty"`
	// Error messages from MFA verification request
	Error NullableString `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerificationResponse VerificationResponse

// NewVerificationResponse instantiates a new VerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationResponse() *VerificationResponse {
	this := VerificationResponse{}
	return &this
}

// NewVerificationResponseWithDefaults instantiates a new VerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationResponseWithDefaults() *VerificationResponse {
	this := VerificationResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerificationResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerificationResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *VerificationResponse) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *VerificationResponse) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *VerificationResponse) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *VerificationResponse) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VerificationResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VerificationResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VerificationResponse) SetStatus(v string) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerificationResponse) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerificationResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *VerificationResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *VerificationResponse) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *VerificationResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *VerificationResponse) UnsetError() {
	o.Error.Unset()
}

func (o VerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId.IsSet() {
		toSerialize["requestId"] = o.RequestId.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerificationResponse) UnmarshalJSON(data []byte) (err error) {
	varVerificationResponse := _VerificationResponse{}

	err = json.Unmarshal(data, &varVerificationResponse)

	if err != nil {
		return err
	}

	*o = VerificationResponse(varVerificationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationResponse struct {
	value *VerificationResponse
	isSet bool
}

func (v NullableVerificationResponse) Get() *VerificationResponse {
	return v.value
}

func (v *NullableVerificationResponse) Set(val *VerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationResponse(val *VerificationResponse) *NullableVerificationResponse {
	return &NullableVerificationResponse{value: val, isSet: true}
}

func (v NullableVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


