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

// checks if the VAClusterStatusChangeEventPreviousHealthCheckResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VAClusterStatusChangeEventPreviousHealthCheckResult{}

// VAClusterStatusChangeEventPreviousHealthCheckResult The results of the last health check.
type VAClusterStatusChangeEventPreviousHealthCheckResult struct {
	// Detailed message of the result of the health check.
	Message string `json:"message"`
	// The type of the health check result.
	ResultType string `json:"resultType"`
	// The status of the health check.
	Status map[string]interface{} `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _VAClusterStatusChangeEventPreviousHealthCheckResult VAClusterStatusChangeEventPreviousHealthCheckResult

// NewVAClusterStatusChangeEventPreviousHealthCheckResult instantiates a new VAClusterStatusChangeEventPreviousHealthCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVAClusterStatusChangeEventPreviousHealthCheckResult(message string, resultType string, status map[string]interface{}) *VAClusterStatusChangeEventPreviousHealthCheckResult {
	this := VAClusterStatusChangeEventPreviousHealthCheckResult{}
	this.Message = message
	this.ResultType = resultType
	this.Status = status
	return &this
}

// NewVAClusterStatusChangeEventPreviousHealthCheckResultWithDefaults instantiates a new VAClusterStatusChangeEventPreviousHealthCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVAClusterStatusChangeEventPreviousHealthCheckResultWithDefaults() *VAClusterStatusChangeEventPreviousHealthCheckResult {
	this := VAClusterStatusChangeEventPreviousHealthCheckResult{}
	return &this
}

// GetMessage returns the Message field value
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) SetMessage(v string) {
	o.Message = v
}

// GetResultType returns the ResultType field value
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) GetResultType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) GetResultTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) SetResultType(v string) {
	o.ResultType = v
}

// GetStatus returns the Status field value
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) SetStatus(v map[string]interface{}) {
	o.Status = v
}

func (o VAClusterStatusChangeEventPreviousHealthCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VAClusterStatusChangeEventPreviousHealthCheckResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["resultType"] = o.ResultType
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VAClusterStatusChangeEventPreviousHealthCheckResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"resultType",
		"status",
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

	varVAClusterStatusChangeEventPreviousHealthCheckResult := _VAClusterStatusChangeEventPreviousHealthCheckResult{}

	if err = json.Unmarshal(bytes, &varVAClusterStatusChangeEventPreviousHealthCheckResult); err == nil {
	*o = VAClusterStatusChangeEventPreviousHealthCheckResult(varVAClusterStatusChangeEventPreviousHealthCheckResult)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "resultType")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVAClusterStatusChangeEventPreviousHealthCheckResult struct {
	value *VAClusterStatusChangeEventPreviousHealthCheckResult
	isSet bool
}

func (v NullableVAClusterStatusChangeEventPreviousHealthCheckResult) Get() *VAClusterStatusChangeEventPreviousHealthCheckResult {
	return v.value
}

func (v *NullableVAClusterStatusChangeEventPreviousHealthCheckResult) Set(val *VAClusterStatusChangeEventPreviousHealthCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVAClusterStatusChangeEventPreviousHealthCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVAClusterStatusChangeEventPreviousHealthCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVAClusterStatusChangeEventPreviousHealthCheckResult(val *VAClusterStatusChangeEventPreviousHealthCheckResult) *NullableVAClusterStatusChangeEventPreviousHealthCheckResult {
	return &NullableVAClusterStatusChangeEventPreviousHealthCheckResult{value: val, isSet: true}
}

func (v NullableVAClusterStatusChangeEventPreviousHealthCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVAClusterStatusChangeEventPreviousHealthCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


