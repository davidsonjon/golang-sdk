/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the FormError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormError{}

// FormError struct for FormError
type FormError struct {
	// Key is the technical key
	Key *string `json:"key,omitempty"`
	// Messages is a list of web.ErrorMessage items
	Messages []ErrorMessage `json:"messages,omitempty"`
	// Value is the value associated with a Key
	Value map[string]interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormError FormError

// NewFormError instantiates a new FormError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormError() *FormError {
	this := FormError{}
	return &this
}

// NewFormErrorWithDefaults instantiates a new FormError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormErrorWithDefaults() *FormError {
	this := FormError{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FormError) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormError) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FormError) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FormError) SetKey(v string) {
	o.Key = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *FormError) GetMessages() []ErrorMessage {
	if o == nil || IsNil(o.Messages) {
		var ret []ErrorMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormError) GetMessagesOk() ([]ErrorMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *FormError) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ErrorMessage and assigns it to the Messages field.
func (o *FormError) SetMessages(v []ErrorMessage) {
	o.Messages = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FormError) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormError) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FormError) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *FormError) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o FormError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormError) UnmarshalJSON(data []byte) (err error) {
	varFormError := _FormError{}

	err = json.Unmarshal(data, &varFormError)

	if err != nil {
		return err
	}

	*o = FormError(varFormError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormError struct {
	value *FormError
	isSet bool
}

func (v NullableFormError) Get() *FormError {
	return v.value
}

func (v *NullableFormError) Set(val *FormError) {
	v.value = val
	v.isSet = true
}

func (v NullableFormError) IsSet() bool {
	return v.isSet
}

func (v *NullableFormError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormError(val *FormError) *NullableFormError {
	return &NullableFormError{value: val, isSet: true}
}

func (v NullableFormError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


