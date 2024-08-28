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

// checks if the Requestability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Requestability{}

// Requestability struct for Requestability
type Requestability struct {
	// Whether the requester of the containing object must provide comments justifying the request
	CommentsRequired NullableBool `json:"commentsRequired,omitempty"`
	// Whether an approver must provide comments when denying the request
	DenialCommentsRequired NullableBool `json:"denialCommentsRequired,omitempty"`
	// List describing the steps in approving the request
	ApprovalSchemes []AccessProfileApprovalScheme `json:"approvalSchemes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Requestability Requestability

// NewRequestability instantiates a new Requestability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestability() *Requestability {
	this := Requestability{}
	var commentsRequired bool = false
	this.CommentsRequired = *NewNullableBool(&commentsRequired)
	var denialCommentsRequired bool = false
	this.DenialCommentsRequired = *NewNullableBool(&denialCommentsRequired)
	return &this
}

// NewRequestabilityWithDefaults instantiates a new Requestability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestabilityWithDefaults() *Requestability {
	this := Requestability{}
	var commentsRequired bool = false
	this.CommentsRequired = *NewNullableBool(&commentsRequired)
	var denialCommentsRequired bool = false
	this.DenialCommentsRequired = *NewNullableBool(&denialCommentsRequired)
	return &this
}

// GetCommentsRequired returns the CommentsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requestability) GetCommentsRequired() bool {
	if o == nil || IsNil(o.CommentsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.CommentsRequired.Get()
}

// GetCommentsRequiredOk returns a tuple with the CommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requestability) GetCommentsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommentsRequired.Get(), o.CommentsRequired.IsSet()
}

// HasCommentsRequired returns a boolean if a field has been set.
func (o *Requestability) HasCommentsRequired() bool {
	if o != nil && o.CommentsRequired.IsSet() {
		return true
	}

	return false
}

// SetCommentsRequired gets a reference to the given NullableBool and assigns it to the CommentsRequired field.
func (o *Requestability) SetCommentsRequired(v bool) {
	o.CommentsRequired.Set(&v)
}
// SetCommentsRequiredNil sets the value for CommentsRequired to be an explicit nil
func (o *Requestability) SetCommentsRequiredNil() {
	o.CommentsRequired.Set(nil)
}

// UnsetCommentsRequired ensures that no value is present for CommentsRequired, not even an explicit nil
func (o *Requestability) UnsetCommentsRequired() {
	o.CommentsRequired.Unset()
}

// GetDenialCommentsRequired returns the DenialCommentsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requestability) GetDenialCommentsRequired() bool {
	if o == nil || IsNil(o.DenialCommentsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.DenialCommentsRequired.Get()
}

// GetDenialCommentsRequiredOk returns a tuple with the DenialCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requestability) GetDenialCommentsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenialCommentsRequired.Get(), o.DenialCommentsRequired.IsSet()
}

// HasDenialCommentsRequired returns a boolean if a field has been set.
func (o *Requestability) HasDenialCommentsRequired() bool {
	if o != nil && o.DenialCommentsRequired.IsSet() {
		return true
	}

	return false
}

// SetDenialCommentsRequired gets a reference to the given NullableBool and assigns it to the DenialCommentsRequired field.
func (o *Requestability) SetDenialCommentsRequired(v bool) {
	o.DenialCommentsRequired.Set(&v)
}
// SetDenialCommentsRequiredNil sets the value for DenialCommentsRequired to be an explicit nil
func (o *Requestability) SetDenialCommentsRequiredNil() {
	o.DenialCommentsRequired.Set(nil)
}

// UnsetDenialCommentsRequired ensures that no value is present for DenialCommentsRequired, not even an explicit nil
func (o *Requestability) UnsetDenialCommentsRequired() {
	o.DenialCommentsRequired.Unset()
}

// GetApprovalSchemes returns the ApprovalSchemes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requestability) GetApprovalSchemes() []AccessProfileApprovalScheme {
	if o == nil {
		var ret []AccessProfileApprovalScheme
		return ret
	}
	return o.ApprovalSchemes
}

// GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requestability) GetApprovalSchemesOk() ([]AccessProfileApprovalScheme, bool) {
	if o == nil || IsNil(o.ApprovalSchemes) {
		return nil, false
	}
	return o.ApprovalSchemes, true
}

// HasApprovalSchemes returns a boolean if a field has been set.
func (o *Requestability) HasApprovalSchemes() bool {
	if o != nil && !IsNil(o.ApprovalSchemes) {
		return true
	}

	return false
}

// SetApprovalSchemes gets a reference to the given []AccessProfileApprovalScheme and assigns it to the ApprovalSchemes field.
func (o *Requestability) SetApprovalSchemes(v []AccessProfileApprovalScheme) {
	o.ApprovalSchemes = v
}

func (o Requestability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requestability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CommentsRequired.IsSet() {
		toSerialize["commentsRequired"] = o.CommentsRequired.Get()
	}
	if o.DenialCommentsRequired.IsSet() {
		toSerialize["denialCommentsRequired"] = o.DenialCommentsRequired.Get()
	}
	if o.ApprovalSchemes != nil {
		toSerialize["approvalSchemes"] = o.ApprovalSchemes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Requestability) UnmarshalJSON(data []byte) (err error) {
	varRequestability := _Requestability{}

	err = json.Unmarshal(data, &varRequestability)

	if err != nil {
		return err
	}

	*o = Requestability(varRequestability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commentsRequired")
		delete(additionalProperties, "denialCommentsRequired")
		delete(additionalProperties, "approvalSchemes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestability struct {
	value *Requestability
	isSet bool
}

func (v NullableRequestability) Get() *Requestability {
	return v.value
}

func (v *NullableRequestability) Set(val *Requestability) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestability) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestability(val *Requestability) *NullableRequestability {
	return &NullableRequestability{value: val, isSet: true}
}

func (v NullableRequestability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


