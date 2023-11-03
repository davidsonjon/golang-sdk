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

// checks if the SodViolationContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationContext{}

// SodViolationContext The contextual information of the violated criteria
type SodViolationContext struct {
	Policy *SodPolicyDto `json:"policy,omitempty"`
	ConflictingAccessCriteria *SodViolationContextConflictingAccessCriteria `json:"conflictingAccessCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContext SodViolationContext

// NewSodViolationContext instantiates a new SodViolationContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContext() *SodViolationContext {
	this := SodViolationContext{}
	return &this
}

// NewSodViolationContextWithDefaults instantiates a new SodViolationContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContextWithDefaults() *SodViolationContext {
	this := SodViolationContext{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *SodViolationContext) GetPolicy() SodPolicyDto {
	if o == nil || isNil(o.Policy) {
		var ret SodPolicyDto
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContext) GetPolicyOk() (*SodPolicyDto, bool) {
	if o == nil || isNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *SodViolationContext) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given SodPolicyDto and assigns it to the Policy field.
func (o *SodViolationContext) SetPolicy(v SodPolicyDto) {
	o.Policy = &v
}

// GetConflictingAccessCriteria returns the ConflictingAccessCriteria field value if set, zero value otherwise.
func (o *SodViolationContext) GetConflictingAccessCriteria() SodViolationContextConflictingAccessCriteria {
	if o == nil || isNil(o.ConflictingAccessCriteria) {
		var ret SodViolationContextConflictingAccessCriteria
		return ret
	}
	return *o.ConflictingAccessCriteria
}

// GetConflictingAccessCriteriaOk returns a tuple with the ConflictingAccessCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContext) GetConflictingAccessCriteriaOk() (*SodViolationContextConflictingAccessCriteria, bool) {
	if o == nil || isNil(o.ConflictingAccessCriteria) {
		return nil, false
	}
	return o.ConflictingAccessCriteria, true
}

// HasConflictingAccessCriteria returns a boolean if a field has been set.
func (o *SodViolationContext) HasConflictingAccessCriteria() bool {
	if o != nil && !isNil(o.ConflictingAccessCriteria) {
		return true
	}

	return false
}

// SetConflictingAccessCriteria gets a reference to the given SodViolationContextConflictingAccessCriteria and assigns it to the ConflictingAccessCriteria field.
func (o *SodViolationContext) SetConflictingAccessCriteria(v SodViolationContextConflictingAccessCriteria) {
	o.ConflictingAccessCriteria = &v
}

func (o SodViolationContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.ConflictingAccessCriteria) {
		toSerialize["conflictingAccessCriteria"] = o.ConflictingAccessCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationContext) UnmarshalJSON(bytes []byte) (err error) {
	varSodViolationContext := _SodViolationContext{}

	if err = json.Unmarshal(bytes, &varSodViolationContext); err == nil {
		*o = SodViolationContext(varSodViolationContext)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "policy")
		delete(additionalProperties, "conflictingAccessCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContext struct {
	value *SodViolationContext
	isSet bool
}

func (v NullableSodViolationContext) Get() *SodViolationContext {
	return v.value
}

func (v *NullableSodViolationContext) Set(val *SodViolationContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContext(val *SodViolationContext) *NullableSodViolationContext {
	return &NullableSodViolationContext{value: val, isSet: true}
}

func (v NullableSodViolationContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


