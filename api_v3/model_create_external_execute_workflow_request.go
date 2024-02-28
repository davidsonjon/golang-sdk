/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the CreateExternalExecuteWorkflowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExternalExecuteWorkflowRequest{}

// CreateExternalExecuteWorkflowRequest struct for CreateExternalExecuteWorkflowRequest
type CreateExternalExecuteWorkflowRequest struct {
	// The input for the workflow
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateExternalExecuteWorkflowRequest CreateExternalExecuteWorkflowRequest

// NewCreateExternalExecuteWorkflowRequest instantiates a new CreateExternalExecuteWorkflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExternalExecuteWorkflowRequest() *CreateExternalExecuteWorkflowRequest {
	this := CreateExternalExecuteWorkflowRequest{}
	return &this
}

// NewCreateExternalExecuteWorkflowRequestWithDefaults instantiates a new CreateExternalExecuteWorkflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExternalExecuteWorkflowRequestWithDefaults() *CreateExternalExecuteWorkflowRequest {
	this := CreateExternalExecuteWorkflowRequest{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *CreateExternalExecuteWorkflowRequest) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExternalExecuteWorkflowRequest) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *CreateExternalExecuteWorkflowRequest) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *CreateExternalExecuteWorkflowRequest) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o CreateExternalExecuteWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateExternalExecuteWorkflowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateExternalExecuteWorkflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateExternalExecuteWorkflowRequest := _CreateExternalExecuteWorkflowRequest{}

	if err = json.Unmarshal(bytes, &varCreateExternalExecuteWorkflowRequest); err == nil {
	*o = CreateExternalExecuteWorkflowRequest(varCreateExternalExecuteWorkflowRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateExternalExecuteWorkflowRequest struct {
	value *CreateExternalExecuteWorkflowRequest
	isSet bool
}

func (v NullableCreateExternalExecuteWorkflowRequest) Get() *CreateExternalExecuteWorkflowRequest {
	return v.value
}

func (v *NullableCreateExternalExecuteWorkflowRequest) Set(val *CreateExternalExecuteWorkflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExternalExecuteWorkflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExternalExecuteWorkflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExternalExecuteWorkflowRequest(val *CreateExternalExecuteWorkflowRequest) *NullableCreateExternalExecuteWorkflowRequest {
	return &NullableCreateExternalExecuteWorkflowRequest{value: val, isSet: true}
}

func (v NullableCreateExternalExecuteWorkflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExternalExecuteWorkflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


