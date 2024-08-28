/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the CompleteInvocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteInvocation{}

// CompleteInvocation struct for CompleteInvocation
type CompleteInvocation struct {
	// Unique invocation secret that was generated when the invocation was created. Required to authenticate to the endpoint.
	Secret string `json:"secret"`
	// The error message to indicate a failed invocation or error if any.
	Error *string `json:"error,omitempty"`
	// Trigger output to complete the invocation. Its schema is defined in the trigger definition.
	Output map[string]interface{} `json:"output"`
	AdditionalProperties map[string]interface{}
}

type _CompleteInvocation CompleteInvocation

// NewCompleteInvocation instantiates a new CompleteInvocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteInvocation(secret string, output map[string]interface{}) *CompleteInvocation {
	this := CompleteInvocation{}
	this.Secret = secret
	this.Output = output
	return &this
}

// NewCompleteInvocationWithDefaults instantiates a new CompleteInvocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteInvocationWithDefaults() *CompleteInvocation {
	this := CompleteInvocation{}
	return &this
}

// GetSecret returns the Secret field value
func (o *CompleteInvocation) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CompleteInvocation) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CompleteInvocation) SetSecret(v string) {
	o.Secret = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CompleteInvocation) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteInvocation) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CompleteInvocation) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *CompleteInvocation) SetError(v string) {
	o.Error = &v
}

// GetOutput returns the Output field value
func (o *CompleteInvocation) GetOutput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
func (o *CompleteInvocation) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// SetOutput sets field value
func (o *CompleteInvocation) SetOutput(v map[string]interface{}) {
	o.Output = v
}

func (o CompleteInvocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteInvocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secret"] = o.Secret
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["output"] = o.Output

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompleteInvocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secret",
		"output",
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

	varCompleteInvocation := _CompleteInvocation{}

	err = json.Unmarshal(data, &varCompleteInvocation)

	if err != nil {
		return err
	}

	*o = CompleteInvocation(varCompleteInvocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "secret")
		delete(additionalProperties, "error")
		delete(additionalProperties, "output")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompleteInvocation struct {
	value *CompleteInvocation
	isSet bool
}

func (v NullableCompleteInvocation) Get() *CompleteInvocation {
	return v.value
}

func (v *NullableCompleteInvocation) Set(val *CompleteInvocation) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteInvocation) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteInvocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteInvocation(val *CompleteInvocation) *NullableCompleteInvocation {
	return &NullableCompleteInvocation{value: val, isSet: true}
}

func (v NullableCompleteInvocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteInvocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


