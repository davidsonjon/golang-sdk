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

// checks if the SodViolationContextCheckCompleted1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationContextCheckCompleted1{}

// SodViolationContextCheckCompleted1 An object referencing a completed SOD violation check
type SodViolationContextCheckCompleted1 struct {
	// The status of SOD violation check
	State *string `json:"state,omitempty"`
	// The id of the Violation check event
	Uuid *string `json:"uuid,omitempty"`
	ViolationCheckResult *SodViolationCheckResult1 `json:"violationCheckResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContextCheckCompleted1 SodViolationContextCheckCompleted1

// NewSodViolationContextCheckCompleted1 instantiates a new SodViolationContextCheckCompleted1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContextCheckCompleted1() *SodViolationContextCheckCompleted1 {
	this := SodViolationContextCheckCompleted1{}
	return &this
}

// NewSodViolationContextCheckCompleted1WithDefaults instantiates a new SodViolationContextCheckCompleted1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContextCheckCompleted1WithDefaults() *SodViolationContextCheckCompleted1 {
	this := SodViolationContextCheckCompleted1{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted1) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted1) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted1) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SodViolationContextCheckCompleted1) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted1) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted1) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted1) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SodViolationContextCheckCompleted1) SetUuid(v string) {
	o.Uuid = &v
}

// GetViolationCheckResult returns the ViolationCheckResult field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted1) GetViolationCheckResult() SodViolationCheckResult1 {
	if o == nil || isNil(o.ViolationCheckResult) {
		var ret SodViolationCheckResult1
		return ret
	}
	return *o.ViolationCheckResult
}

// GetViolationCheckResultOk returns a tuple with the ViolationCheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted1) GetViolationCheckResultOk() (*SodViolationCheckResult1, bool) {
	if o == nil || isNil(o.ViolationCheckResult) {
		return nil, false
	}
	return o.ViolationCheckResult, true
}

// HasViolationCheckResult returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted1) HasViolationCheckResult() bool {
	if o != nil && !isNil(o.ViolationCheckResult) {
		return true
	}

	return false
}

// SetViolationCheckResult gets a reference to the given SodViolationCheckResult1 and assigns it to the ViolationCheckResult field.
func (o *SodViolationContextCheckCompleted1) SetViolationCheckResult(v SodViolationCheckResult1) {
	o.ViolationCheckResult = &v
}

func (o SodViolationContextCheckCompleted1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationContextCheckCompleted1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.ViolationCheckResult) {
		toSerialize["violationCheckResult"] = o.ViolationCheckResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationContextCheckCompleted1) UnmarshalJSON(bytes []byte) (err error) {
	varSodViolationContextCheckCompleted1 := _SodViolationContextCheckCompleted1{}

	if err = json.Unmarshal(bytes, &varSodViolationContextCheckCompleted1); err == nil {
	*o = SodViolationContextCheckCompleted1(varSodViolationContextCheckCompleted1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "violationCheckResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContextCheckCompleted1 struct {
	value *SodViolationContextCheckCompleted1
	isSet bool
}

func (v NullableSodViolationContextCheckCompleted1) Get() *SodViolationContextCheckCompleted1 {
	return v.value
}

func (v *NullableSodViolationContextCheckCompleted1) Set(val *SodViolationContextCheckCompleted1) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContextCheckCompleted1) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContextCheckCompleted1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContextCheckCompleted1(val *SodViolationContextCheckCompleted1) *NullableSodViolationContextCheckCompleted1 {
	return &NullableSodViolationContextCheckCompleted1{value: val, isSet: true}
}

func (v NullableSodViolationContextCheckCompleted1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContextCheckCompleted1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


