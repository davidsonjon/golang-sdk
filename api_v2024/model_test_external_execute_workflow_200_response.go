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

// checks if the TestExternalExecuteWorkflow200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestExternalExecuteWorkflow200Response{}

// TestExternalExecuteWorkflow200Response struct for TestExternalExecuteWorkflow200Response
type TestExternalExecuteWorkflow200Response struct {
	// The input that was received
	Payload map[string]interface{} `json:"payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestExternalExecuteWorkflow200Response TestExternalExecuteWorkflow200Response

// NewTestExternalExecuteWorkflow200Response instantiates a new TestExternalExecuteWorkflow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestExternalExecuteWorkflow200Response() *TestExternalExecuteWorkflow200Response {
	this := TestExternalExecuteWorkflow200Response{}
	return &this
}

// NewTestExternalExecuteWorkflow200ResponseWithDefaults instantiates a new TestExternalExecuteWorkflow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestExternalExecuteWorkflow200ResponseWithDefaults() *TestExternalExecuteWorkflow200Response {
	this := TestExternalExecuteWorkflow200Response{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *TestExternalExecuteWorkflow200Response) GetPayload() map[string]interface{} {
	if o == nil || isNil(o.Payload) {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestExternalExecuteWorkflow200Response) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Payload) {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *TestExternalExecuteWorkflow200Response) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *TestExternalExecuteWorkflow200Response) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

func (o TestExternalExecuteWorkflow200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestExternalExecuteWorkflow200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestExternalExecuteWorkflow200Response) UnmarshalJSON(bytes []byte) (err error) {
	varTestExternalExecuteWorkflow200Response := _TestExternalExecuteWorkflow200Response{}

	if err = json.Unmarshal(bytes, &varTestExternalExecuteWorkflow200Response); err == nil {
	*o = TestExternalExecuteWorkflow200Response(varTestExternalExecuteWorkflow200Response)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestExternalExecuteWorkflow200Response struct {
	value *TestExternalExecuteWorkflow200Response
	isSet bool
}

func (v NullableTestExternalExecuteWorkflow200Response) Get() *TestExternalExecuteWorkflow200Response {
	return v.value
}

func (v *NullableTestExternalExecuteWorkflow200Response) Set(val *TestExternalExecuteWorkflow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTestExternalExecuteWorkflow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTestExternalExecuteWorkflow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestExternalExecuteWorkflow200Response(val *TestExternalExecuteWorkflow200Response) *NullableTestExternalExecuteWorkflow200Response {
	return &NullableTestExternalExecuteWorkflow200Response{value: val, isSet: true}
}

func (v NullableTestExternalExecuteWorkflow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestExternalExecuteWorkflow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


