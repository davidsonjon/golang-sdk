/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// CreateAccessRequest429Response struct for CreateAccessRequest429Response
type CreateAccessRequest429Response struct {
	// A message describing the error
	Message map[string]interface{} `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccessRequest429Response CreateAccessRequest429Response

// NewCreateAccessRequest429Response instantiates a new CreateAccessRequest429Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessRequest429Response() *CreateAccessRequest429Response {
	this := CreateAccessRequest429Response{}
	return &this
}

// NewCreateAccessRequest429ResponseWithDefaults instantiates a new CreateAccessRequest429Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessRequest429ResponseWithDefaults() *CreateAccessRequest429Response {
	this := CreateAccessRequest429Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateAccessRequest429Response) GetMessage() map[string]interface{} {
	if o == nil || isNil(o.Message) {
		var ret map[string]interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessRequest429Response) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Message) {
		return map[string]interface{}{}, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateAccessRequest429Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given map[string]interface{} and assigns it to the Message field.
func (o *CreateAccessRequest429Response) SetMessage(v map[string]interface{}) {
	o.Message = v
}

func (o CreateAccessRequest429Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateAccessRequest429Response) UnmarshalJSON(bytes []byte) (err error) {
	varCreateAccessRequest429Response := _CreateAccessRequest429Response{}

	if err = json.Unmarshal(bytes, &varCreateAccessRequest429Response); err == nil {
		*o = CreateAccessRequest429Response(varCreateAccessRequest429Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccessRequest429Response struct {
	value *CreateAccessRequest429Response
	isSet bool
}

func (v NullableCreateAccessRequest429Response) Get() *CreateAccessRequest429Response {
	return v.value
}

func (v *NullableCreateAccessRequest429Response) Set(val *CreateAccessRequest429Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessRequest429Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessRequest429Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessRequest429Response(val *CreateAccessRequest429Response) *NullableCreateAccessRequest429Response {
	return &NullableCreateAccessRequest429Response{value: val, isSet: true}
}

func (v NullableCreateAccessRequest429Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessRequest429Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


