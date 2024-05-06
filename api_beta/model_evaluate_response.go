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

// checks if the EvaluateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvaluateResponse{}

// EvaluateResponse The response body for Evaluate Reassignment Configuration
type EvaluateResponse struct {
	// The Identity ID which should be the recipient of any work items sent to a specific identity & work type
	ReassignToId *string `json:"reassignToId,omitempty"`
	// List of Reassignments found by looking up the next `TargetIdentity` in a ReassignmentConfiguration
	LookupTrail []LookupStep `json:"lookupTrail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EvaluateResponse EvaluateResponse

// NewEvaluateResponse instantiates a new EvaluateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvaluateResponse() *EvaluateResponse {
	this := EvaluateResponse{}
	return &this
}

// NewEvaluateResponseWithDefaults instantiates a new EvaluateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvaluateResponseWithDefaults() *EvaluateResponse {
	this := EvaluateResponse{}
	return &this
}

// GetReassignToId returns the ReassignToId field value if set, zero value otherwise.
func (o *EvaluateResponse) GetReassignToId() string {
	if o == nil || isNil(o.ReassignToId) {
		var ret string
		return ret
	}
	return *o.ReassignToId
}

// GetReassignToIdOk returns a tuple with the ReassignToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluateResponse) GetReassignToIdOk() (*string, bool) {
	if o == nil || isNil(o.ReassignToId) {
		return nil, false
	}
	return o.ReassignToId, true
}

// HasReassignToId returns a boolean if a field has been set.
func (o *EvaluateResponse) HasReassignToId() bool {
	if o != nil && !isNil(o.ReassignToId) {
		return true
	}

	return false
}

// SetReassignToId gets a reference to the given string and assigns it to the ReassignToId field.
func (o *EvaluateResponse) SetReassignToId(v string) {
	o.ReassignToId = &v
}

// GetLookupTrail returns the LookupTrail field value if set, zero value otherwise.
func (o *EvaluateResponse) GetLookupTrail() []LookupStep {
	if o == nil || isNil(o.LookupTrail) {
		var ret []LookupStep
		return ret
	}
	return o.LookupTrail
}

// GetLookupTrailOk returns a tuple with the LookupTrail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluateResponse) GetLookupTrailOk() ([]LookupStep, bool) {
	if o == nil || isNil(o.LookupTrail) {
		return nil, false
	}
	return o.LookupTrail, true
}

// HasLookupTrail returns a boolean if a field has been set.
func (o *EvaluateResponse) HasLookupTrail() bool {
	if o != nil && !isNil(o.LookupTrail) {
		return true
	}

	return false
}

// SetLookupTrail gets a reference to the given []LookupStep and assigns it to the LookupTrail field.
func (o *EvaluateResponse) SetLookupTrail(v []LookupStep) {
	o.LookupTrail = v
}

func (o EvaluateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvaluateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReassignToId) {
		toSerialize["reassignToId"] = o.ReassignToId
	}
	if !isNil(o.LookupTrail) {
		toSerialize["lookupTrail"] = o.LookupTrail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EvaluateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEvaluateResponse := _EvaluateResponse{}

	if err = json.Unmarshal(bytes, &varEvaluateResponse); err == nil {
	*o = EvaluateResponse(varEvaluateResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reassignToId")
		delete(additionalProperties, "lookupTrail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvaluateResponse struct {
	value *EvaluateResponse
	isSet bool
}

func (v NullableEvaluateResponse) Get() *EvaluateResponse {
	return v.value
}

func (v *NullableEvaluateResponse) Set(val *EvaluateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluateResponse(val *EvaluateResponse) *NullableEvaluateResponse {
	return &NullableEvaluateResponse{value: val, isSet: true}
}

func (v NullableEvaluateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


