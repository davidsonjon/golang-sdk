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

// checks if the EntitlementAccessRequestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementAccessRequestConfig{}

// EntitlementAccessRequestConfig struct for EntitlementAccessRequestConfig
type EntitlementAccessRequestConfig struct {
	// Ordered list of approval steps for the access request. Empty when no approval is required.
	ApprovalSchemes []EntitlementApprovalScheme `json:"approvalSchemes,omitempty"`
	// If the requester must provide a comment during access request.
	RequestCommentRequired *bool `json:"requestCommentRequired,omitempty"`
	// If the reviewer must provide a comment when denying the access request.
	DenialCommentRequired *bool `json:"denialCommentRequired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementAccessRequestConfig EntitlementAccessRequestConfig

// NewEntitlementAccessRequestConfig instantiates a new EntitlementAccessRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementAccessRequestConfig() *EntitlementAccessRequestConfig {
	this := EntitlementAccessRequestConfig{}
	var requestCommentRequired bool = false
	this.RequestCommentRequired = &requestCommentRequired
	var denialCommentRequired bool = false
	this.DenialCommentRequired = &denialCommentRequired
	return &this
}

// NewEntitlementAccessRequestConfigWithDefaults instantiates a new EntitlementAccessRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementAccessRequestConfigWithDefaults() *EntitlementAccessRequestConfig {
	this := EntitlementAccessRequestConfig{}
	var requestCommentRequired bool = false
	this.RequestCommentRequired = &requestCommentRequired
	var denialCommentRequired bool = false
	this.DenialCommentRequired = &denialCommentRequired
	return &this
}

// GetApprovalSchemes returns the ApprovalSchemes field value if set, zero value otherwise.
func (o *EntitlementAccessRequestConfig) GetApprovalSchemes() []EntitlementApprovalScheme {
	if o == nil || isNil(o.ApprovalSchemes) {
		var ret []EntitlementApprovalScheme
		return ret
	}
	return o.ApprovalSchemes
}

// GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAccessRequestConfig) GetApprovalSchemesOk() ([]EntitlementApprovalScheme, bool) {
	if o == nil || isNil(o.ApprovalSchemes) {
		return nil, false
	}
	return o.ApprovalSchemes, true
}

// HasApprovalSchemes returns a boolean if a field has been set.
func (o *EntitlementAccessRequestConfig) HasApprovalSchemes() bool {
	if o != nil && !isNil(o.ApprovalSchemes) {
		return true
	}

	return false
}

// SetApprovalSchemes gets a reference to the given []EntitlementApprovalScheme and assigns it to the ApprovalSchemes field.
func (o *EntitlementAccessRequestConfig) SetApprovalSchemes(v []EntitlementApprovalScheme) {
	o.ApprovalSchemes = v
}

// GetRequestCommentRequired returns the RequestCommentRequired field value if set, zero value otherwise.
func (o *EntitlementAccessRequestConfig) GetRequestCommentRequired() bool {
	if o == nil || isNil(o.RequestCommentRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentRequired
}

// GetRequestCommentRequiredOk returns a tuple with the RequestCommentRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAccessRequestConfig) GetRequestCommentRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.RequestCommentRequired) {
		return nil, false
	}
	return o.RequestCommentRequired, true
}

// HasRequestCommentRequired returns a boolean if a field has been set.
func (o *EntitlementAccessRequestConfig) HasRequestCommentRequired() bool {
	if o != nil && !isNil(o.RequestCommentRequired) {
		return true
	}

	return false
}

// SetRequestCommentRequired gets a reference to the given bool and assigns it to the RequestCommentRequired field.
func (o *EntitlementAccessRequestConfig) SetRequestCommentRequired(v bool) {
	o.RequestCommentRequired = &v
}

// GetDenialCommentRequired returns the DenialCommentRequired field value if set, zero value otherwise.
func (o *EntitlementAccessRequestConfig) GetDenialCommentRequired() bool {
	if o == nil || isNil(o.DenialCommentRequired) {
		var ret bool
		return ret
	}
	return *o.DenialCommentRequired
}

// GetDenialCommentRequiredOk returns a tuple with the DenialCommentRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAccessRequestConfig) GetDenialCommentRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.DenialCommentRequired) {
		return nil, false
	}
	return o.DenialCommentRequired, true
}

// HasDenialCommentRequired returns a boolean if a field has been set.
func (o *EntitlementAccessRequestConfig) HasDenialCommentRequired() bool {
	if o != nil && !isNil(o.DenialCommentRequired) {
		return true
	}

	return false
}

// SetDenialCommentRequired gets a reference to the given bool and assigns it to the DenialCommentRequired field.
func (o *EntitlementAccessRequestConfig) SetDenialCommentRequired(v bool) {
	o.DenialCommentRequired = &v
}

func (o EntitlementAccessRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementAccessRequestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApprovalSchemes) {
		toSerialize["approvalSchemes"] = o.ApprovalSchemes
	}
	if !isNil(o.RequestCommentRequired) {
		toSerialize["requestCommentRequired"] = o.RequestCommentRequired
	}
	if !isNil(o.DenialCommentRequired) {
		toSerialize["denialCommentRequired"] = o.DenialCommentRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementAccessRequestConfig) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementAccessRequestConfig := _EntitlementAccessRequestConfig{}

	if err = json.Unmarshal(bytes, &varEntitlementAccessRequestConfig); err == nil {
	*o = EntitlementAccessRequestConfig(varEntitlementAccessRequestConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approvalSchemes")
		delete(additionalProperties, "requestCommentRequired")
		delete(additionalProperties, "denialCommentRequired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementAccessRequestConfig struct {
	value *EntitlementAccessRequestConfig
	isSet bool
}

func (v NullableEntitlementAccessRequestConfig) Get() *EntitlementAccessRequestConfig {
	return v.value
}

func (v *NullableEntitlementAccessRequestConfig) Set(val *EntitlementAccessRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementAccessRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementAccessRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementAccessRequestConfig(val *EntitlementAccessRequestConfig) *NullableEntitlementAccessRequestConfig {
	return &NullableEntitlementAccessRequestConfig{value: val, isSet: true}
}

func (v NullableEntitlementAccessRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementAccessRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


