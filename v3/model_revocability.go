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

// checks if the Revocability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Revocability{}

// Revocability struct for Revocability
type Revocability struct {
	// List describing the steps in approving the revocation request
	ApprovalSchemes []AccessProfileApprovalScheme `json:"approvalSchemes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Revocability Revocability

// NewRevocability instantiates a new Revocability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevocability() *Revocability {
	this := Revocability{}
	return &this
}

// NewRevocabilityWithDefaults instantiates a new Revocability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevocabilityWithDefaults() *Revocability {
	this := Revocability{}
	return &this
}

// GetApprovalSchemes returns the ApprovalSchemes field value if set, zero value otherwise.
func (o *Revocability) GetApprovalSchemes() []AccessProfileApprovalScheme {
	if o == nil || isNil(o.ApprovalSchemes) {
		var ret []AccessProfileApprovalScheme
		return ret
	}
	return o.ApprovalSchemes
}

// GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revocability) GetApprovalSchemesOk() ([]AccessProfileApprovalScheme, bool) {
	if o == nil || isNil(o.ApprovalSchemes) {
		return nil, false
	}
	return o.ApprovalSchemes, true
}

// HasApprovalSchemes returns a boolean if a field has been set.
func (o *Revocability) HasApprovalSchemes() bool {
	if o != nil && !isNil(o.ApprovalSchemes) {
		return true
	}

	return false
}

// SetApprovalSchemes gets a reference to the given []AccessProfileApprovalScheme and assigns it to the ApprovalSchemes field.
func (o *Revocability) SetApprovalSchemes(v []AccessProfileApprovalScheme) {
	o.ApprovalSchemes = v
}

func (o Revocability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Revocability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApprovalSchemes) {
		toSerialize["approvalSchemes"] = o.ApprovalSchemes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Revocability) UnmarshalJSON(bytes []byte) (err error) {
	varRevocability := _Revocability{}

	if err = json.Unmarshal(bytes, &varRevocability); err == nil {
		*o = Revocability(varRevocability)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approvalSchemes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevocability struct {
	value *Revocability
	isSet bool
}

func (v NullableRevocability) Get() *Revocability {
	return v.value
}

func (v *NullableRevocability) Set(val *Revocability) {
	v.value = val
	v.isSet = true
}

func (v NullableRevocability) IsSet() bool {
	return v.isSet
}

func (v *NullableRevocability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevocability(val *Revocability) *NullableRevocability {
	return &NullableRevocability{value: val, isSet: true}
}

func (v NullableRevocability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevocability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


