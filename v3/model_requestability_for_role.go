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

// checks if the RequestabilityForRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestabilityForRole{}

// RequestabilityForRole struct for RequestabilityForRole
type RequestabilityForRole struct {
	// Whether the requester of the containing object must provide comments justifying the request
	CommentsRequired NullableBool `json:"commentsRequired,omitempty"`
	// Whether an approver must provide comments when denying the request
	DenialCommentsRequired NullableBool `json:"denialCommentsRequired,omitempty"`
	// List describing the steps in approving the request
	ApprovalSchemes []ApprovalSchemeForRole `json:"approvalSchemes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestabilityForRole RequestabilityForRole

// NewRequestabilityForRole instantiates a new RequestabilityForRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestabilityForRole() *RequestabilityForRole {
	this := RequestabilityForRole{}
	var commentsRequired bool = false
	this.CommentsRequired = *NewNullableBool(&commentsRequired)
	var denialCommentsRequired bool = false
	this.DenialCommentsRequired = *NewNullableBool(&denialCommentsRequired)
	return &this
}

// NewRequestabilityForRoleWithDefaults instantiates a new RequestabilityForRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestabilityForRoleWithDefaults() *RequestabilityForRole {
	this := RequestabilityForRole{}
	var commentsRequired bool = false
	this.CommentsRequired = *NewNullableBool(&commentsRequired)
	var denialCommentsRequired bool = false
	this.DenialCommentsRequired = *NewNullableBool(&denialCommentsRequired)
	return &this
}

// GetCommentsRequired returns the CommentsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestabilityForRole) GetCommentsRequired() bool {
	if o == nil || isNil(o.CommentsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.CommentsRequired.Get()
}

// GetCommentsRequiredOk returns a tuple with the CommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestabilityForRole) GetCommentsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommentsRequired.Get(), o.CommentsRequired.IsSet()
}

// HasCommentsRequired returns a boolean if a field has been set.
func (o *RequestabilityForRole) HasCommentsRequired() bool {
	if o != nil && o.CommentsRequired.IsSet() {
		return true
	}

	return false
}

// SetCommentsRequired gets a reference to the given NullableBool and assigns it to the CommentsRequired field.
func (o *RequestabilityForRole) SetCommentsRequired(v bool) {
	o.CommentsRequired.Set(&v)
}
// SetCommentsRequiredNil sets the value for CommentsRequired to be an explicit nil
func (o *RequestabilityForRole) SetCommentsRequiredNil() {
	o.CommentsRequired.Set(nil)
}

// UnsetCommentsRequired ensures that no value is present for CommentsRequired, not even an explicit nil
func (o *RequestabilityForRole) UnsetCommentsRequired() {
	o.CommentsRequired.Unset()
}

// GetDenialCommentsRequired returns the DenialCommentsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestabilityForRole) GetDenialCommentsRequired() bool {
	if o == nil || isNil(o.DenialCommentsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.DenialCommentsRequired.Get()
}

// GetDenialCommentsRequiredOk returns a tuple with the DenialCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestabilityForRole) GetDenialCommentsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenialCommentsRequired.Get(), o.DenialCommentsRequired.IsSet()
}

// HasDenialCommentsRequired returns a boolean if a field has been set.
func (o *RequestabilityForRole) HasDenialCommentsRequired() bool {
	if o != nil && o.DenialCommentsRequired.IsSet() {
		return true
	}

	return false
}

// SetDenialCommentsRequired gets a reference to the given NullableBool and assigns it to the DenialCommentsRequired field.
func (o *RequestabilityForRole) SetDenialCommentsRequired(v bool) {
	o.DenialCommentsRequired.Set(&v)
}
// SetDenialCommentsRequiredNil sets the value for DenialCommentsRequired to be an explicit nil
func (o *RequestabilityForRole) SetDenialCommentsRequiredNil() {
	o.DenialCommentsRequired.Set(nil)
}

// UnsetDenialCommentsRequired ensures that no value is present for DenialCommentsRequired, not even an explicit nil
func (o *RequestabilityForRole) UnsetDenialCommentsRequired() {
	o.DenialCommentsRequired.Unset()
}

// GetApprovalSchemes returns the ApprovalSchemes field value if set, zero value otherwise.
func (o *RequestabilityForRole) GetApprovalSchemes() []ApprovalSchemeForRole {
	if o == nil || isNil(o.ApprovalSchemes) {
		var ret []ApprovalSchemeForRole
		return ret
	}
	return o.ApprovalSchemes
}

// GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestabilityForRole) GetApprovalSchemesOk() ([]ApprovalSchemeForRole, bool) {
	if o == nil || isNil(o.ApprovalSchemes) {
		return nil, false
	}
	return o.ApprovalSchemes, true
}

// HasApprovalSchemes returns a boolean if a field has been set.
func (o *RequestabilityForRole) HasApprovalSchemes() bool {
	if o != nil && !isNil(o.ApprovalSchemes) {
		return true
	}

	return false
}

// SetApprovalSchemes gets a reference to the given []ApprovalSchemeForRole and assigns it to the ApprovalSchemes field.
func (o *RequestabilityForRole) SetApprovalSchemes(v []ApprovalSchemeForRole) {
	o.ApprovalSchemes = v
}

func (o RequestabilityForRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestabilityForRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CommentsRequired.IsSet() {
		toSerialize["commentsRequired"] = o.CommentsRequired.Get()
	}
	if o.DenialCommentsRequired.IsSet() {
		toSerialize["denialCommentsRequired"] = o.DenialCommentsRequired.Get()
	}
	if !isNil(o.ApprovalSchemes) {
		toSerialize["approvalSchemes"] = o.ApprovalSchemes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestabilityForRole) UnmarshalJSON(bytes []byte) (err error) {
	varRequestabilityForRole := _RequestabilityForRole{}

	if err = json.Unmarshal(bytes, &varRequestabilityForRole); err == nil {
	*o = RequestabilityForRole(varRequestabilityForRole)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "commentsRequired")
		delete(additionalProperties, "denialCommentsRequired")
		delete(additionalProperties, "approvalSchemes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestabilityForRole struct {
	value *RequestabilityForRole
	isSet bool
}

func (v NullableRequestabilityForRole) Get() *RequestabilityForRole {
	return v.value
}

func (v *NullableRequestabilityForRole) Set(val *RequestabilityForRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestabilityForRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestabilityForRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestabilityForRole(val *RequestabilityForRole) *NullableRequestabilityForRole {
	return &NullableRequestabilityForRole{value: val, isSet: true}
}

func (v NullableRequestabilityForRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestabilityForRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


