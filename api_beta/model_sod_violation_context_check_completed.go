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

// checks if the SodViolationContextCheckCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationContextCheckCompleted{}

// SodViolationContextCheckCompleted An object referencing a completed SOD violation check
type SodViolationContextCheckCompleted struct {
	// The status of SOD violation check
	State NullableString `json:"state,omitempty"`
	// The id of the Violation check event
	Uuid NullableString `json:"uuid,omitempty"`
	ViolationCheckResult *SodViolationCheckResult `json:"violationCheckResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContextCheckCompleted SodViolationContextCheckCompleted

// NewSodViolationContextCheckCompleted instantiates a new SodViolationContextCheckCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContextCheckCompleted() *SodViolationContextCheckCompleted {
	this := SodViolationContextCheckCompleted{}
	return &this
}

// NewSodViolationContextCheckCompletedWithDefaults instantiates a new SodViolationContextCheckCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContextCheckCompletedWithDefaults() *SodViolationContextCheckCompleted {
	this := SodViolationContextCheckCompleted{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodViolationContextCheckCompleted) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodViolationContextCheckCompleted) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *SodViolationContextCheckCompleted) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *SodViolationContextCheckCompleted) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *SodViolationContextCheckCompleted) UnsetState() {
	o.State.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodViolationContextCheckCompleted) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodViolationContextCheckCompleted) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *SodViolationContextCheckCompleted) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *SodViolationContextCheckCompleted) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *SodViolationContextCheckCompleted) UnsetUuid() {
	o.Uuid.Unset()
}

// GetViolationCheckResult returns the ViolationCheckResult field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted) GetViolationCheckResult() SodViolationCheckResult {
	if o == nil || IsNil(o.ViolationCheckResult) {
		var ret SodViolationCheckResult
		return ret
	}
	return *o.ViolationCheckResult
}

// GetViolationCheckResultOk returns a tuple with the ViolationCheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted) GetViolationCheckResultOk() (*SodViolationCheckResult, bool) {
	if o == nil || IsNil(o.ViolationCheckResult) {
		return nil, false
	}
	return o.ViolationCheckResult, true
}

// HasViolationCheckResult returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted) HasViolationCheckResult() bool {
	if o != nil && !IsNil(o.ViolationCheckResult) {
		return true
	}

	return false
}

// SetViolationCheckResult gets a reference to the given SodViolationCheckResult and assigns it to the ViolationCheckResult field.
func (o *SodViolationContextCheckCompleted) SetViolationCheckResult(v SodViolationCheckResult) {
	o.ViolationCheckResult = &v
}

func (o SodViolationContextCheckCompleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationContextCheckCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if !IsNil(o.ViolationCheckResult) {
		toSerialize["violationCheckResult"] = o.ViolationCheckResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationContextCheckCompleted) UnmarshalJSON(data []byte) (err error) {
	varSodViolationContextCheckCompleted := _SodViolationContextCheckCompleted{}

	err = json.Unmarshal(data, &varSodViolationContextCheckCompleted)

	if err != nil {
		return err
	}

	*o = SodViolationContextCheckCompleted(varSodViolationContextCheckCompleted)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "violationCheckResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContextCheckCompleted struct {
	value *SodViolationContextCheckCompleted
	isSet bool
}

func (v NullableSodViolationContextCheckCompleted) Get() *SodViolationContextCheckCompleted {
	return v.value
}

func (v *NullableSodViolationContextCheckCompleted) Set(val *SodViolationContextCheckCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContextCheckCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContextCheckCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContextCheckCompleted(val *SodViolationContextCheckCompleted) *NullableSodViolationContextCheckCompleted {
	return &NullableSodViolationContextCheckCompleted{value: val, isSet: true}
}

func (v NullableSodViolationContextCheckCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContextCheckCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


