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

// checks if the AccessRequestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestConfig{}

// AccessRequestConfig struct for AccessRequestConfig
type AccessRequestConfig struct {
	// If this is true, approvals must be processed by an external system. Also, if this is true, it blocks Request Center access requests and returns an error for any user who isn't an org admin.
	ApprovalsMustBeExternal *bool `json:"approvalsMustBeExternal,omitempty"`
	// If this is true and the requester and reviewer are the same, the request is automatically approved.
	AutoApprovalEnabled *bool `json:"autoApprovalEnabled,omitempty"`
	RequestOnBehalfOfConfig *RequestOnBehalfOfConfig `json:"requestOnBehalfOfConfig,omitempty"`
	ApprovalReminderAndEscalationConfig *ApprovalReminderAndEscalationConfig `json:"approvalReminderAndEscalationConfig,omitempty"`
	EntitlementRequestConfig *EntitlementRequestConfig1 `json:"entitlementRequestConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestConfig AccessRequestConfig

// NewAccessRequestConfig instantiates a new AccessRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestConfig() *AccessRequestConfig {
	this := AccessRequestConfig{}
	var approvalsMustBeExternal bool = false
	this.ApprovalsMustBeExternal = &approvalsMustBeExternal
	var autoApprovalEnabled bool = false
	this.AutoApprovalEnabled = &autoApprovalEnabled
	return &this
}

// NewAccessRequestConfigWithDefaults instantiates a new AccessRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestConfigWithDefaults() *AccessRequestConfig {
	this := AccessRequestConfig{}
	var approvalsMustBeExternal bool = false
	this.ApprovalsMustBeExternal = &approvalsMustBeExternal
	var autoApprovalEnabled bool = false
	this.AutoApprovalEnabled = &autoApprovalEnabled
	return &this
}

// GetApprovalsMustBeExternal returns the ApprovalsMustBeExternal field value if set, zero value otherwise.
func (o *AccessRequestConfig) GetApprovalsMustBeExternal() bool {
	if o == nil || isNil(o.ApprovalsMustBeExternal) {
		var ret bool
		return ret
	}
	return *o.ApprovalsMustBeExternal
}

// GetApprovalsMustBeExternalOk returns a tuple with the ApprovalsMustBeExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestConfig) GetApprovalsMustBeExternalOk() (*bool, bool) {
	if o == nil || isNil(o.ApprovalsMustBeExternal) {
		return nil, false
	}
	return o.ApprovalsMustBeExternal, true
}

// HasApprovalsMustBeExternal returns a boolean if a field has been set.
func (o *AccessRequestConfig) HasApprovalsMustBeExternal() bool {
	if o != nil && !isNil(o.ApprovalsMustBeExternal) {
		return true
	}

	return false
}

// SetApprovalsMustBeExternal gets a reference to the given bool and assigns it to the ApprovalsMustBeExternal field.
func (o *AccessRequestConfig) SetApprovalsMustBeExternal(v bool) {
	o.ApprovalsMustBeExternal = &v
}

// GetAutoApprovalEnabled returns the AutoApprovalEnabled field value if set, zero value otherwise.
func (o *AccessRequestConfig) GetAutoApprovalEnabled() bool {
	if o == nil || isNil(o.AutoApprovalEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoApprovalEnabled
}

// GetAutoApprovalEnabledOk returns a tuple with the AutoApprovalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestConfig) GetAutoApprovalEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AutoApprovalEnabled) {
		return nil, false
	}
	return o.AutoApprovalEnabled, true
}

// HasAutoApprovalEnabled returns a boolean if a field has been set.
func (o *AccessRequestConfig) HasAutoApprovalEnabled() bool {
	if o != nil && !isNil(o.AutoApprovalEnabled) {
		return true
	}

	return false
}

// SetAutoApprovalEnabled gets a reference to the given bool and assigns it to the AutoApprovalEnabled field.
func (o *AccessRequestConfig) SetAutoApprovalEnabled(v bool) {
	o.AutoApprovalEnabled = &v
}

// GetRequestOnBehalfOfConfig returns the RequestOnBehalfOfConfig field value if set, zero value otherwise.
func (o *AccessRequestConfig) GetRequestOnBehalfOfConfig() RequestOnBehalfOfConfig {
	if o == nil || isNil(o.RequestOnBehalfOfConfig) {
		var ret RequestOnBehalfOfConfig
		return ret
	}
	return *o.RequestOnBehalfOfConfig
}

// GetRequestOnBehalfOfConfigOk returns a tuple with the RequestOnBehalfOfConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestConfig) GetRequestOnBehalfOfConfigOk() (*RequestOnBehalfOfConfig, bool) {
	if o == nil || isNil(o.RequestOnBehalfOfConfig) {
		return nil, false
	}
	return o.RequestOnBehalfOfConfig, true
}

// HasRequestOnBehalfOfConfig returns a boolean if a field has been set.
func (o *AccessRequestConfig) HasRequestOnBehalfOfConfig() bool {
	if o != nil && !isNil(o.RequestOnBehalfOfConfig) {
		return true
	}

	return false
}

// SetRequestOnBehalfOfConfig gets a reference to the given RequestOnBehalfOfConfig and assigns it to the RequestOnBehalfOfConfig field.
func (o *AccessRequestConfig) SetRequestOnBehalfOfConfig(v RequestOnBehalfOfConfig) {
	o.RequestOnBehalfOfConfig = &v
}

// GetApprovalReminderAndEscalationConfig returns the ApprovalReminderAndEscalationConfig field value if set, zero value otherwise.
func (o *AccessRequestConfig) GetApprovalReminderAndEscalationConfig() ApprovalReminderAndEscalationConfig {
	if o == nil || isNil(o.ApprovalReminderAndEscalationConfig) {
		var ret ApprovalReminderAndEscalationConfig
		return ret
	}
	return *o.ApprovalReminderAndEscalationConfig
}

// GetApprovalReminderAndEscalationConfigOk returns a tuple with the ApprovalReminderAndEscalationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestConfig) GetApprovalReminderAndEscalationConfigOk() (*ApprovalReminderAndEscalationConfig, bool) {
	if o == nil || isNil(o.ApprovalReminderAndEscalationConfig) {
		return nil, false
	}
	return o.ApprovalReminderAndEscalationConfig, true
}

// HasApprovalReminderAndEscalationConfig returns a boolean if a field has been set.
func (o *AccessRequestConfig) HasApprovalReminderAndEscalationConfig() bool {
	if o != nil && !isNil(o.ApprovalReminderAndEscalationConfig) {
		return true
	}

	return false
}

// SetApprovalReminderAndEscalationConfig gets a reference to the given ApprovalReminderAndEscalationConfig and assigns it to the ApprovalReminderAndEscalationConfig field.
func (o *AccessRequestConfig) SetApprovalReminderAndEscalationConfig(v ApprovalReminderAndEscalationConfig) {
	o.ApprovalReminderAndEscalationConfig = &v
}

// GetEntitlementRequestConfig returns the EntitlementRequestConfig field value if set, zero value otherwise.
func (o *AccessRequestConfig) GetEntitlementRequestConfig() EntitlementRequestConfig1 {
	if o == nil || isNil(o.EntitlementRequestConfig) {
		var ret EntitlementRequestConfig1
		return ret
	}
	return *o.EntitlementRequestConfig
}

// GetEntitlementRequestConfigOk returns a tuple with the EntitlementRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestConfig) GetEntitlementRequestConfigOk() (*EntitlementRequestConfig1, bool) {
	if o == nil || isNil(o.EntitlementRequestConfig) {
		return nil, false
	}
	return o.EntitlementRequestConfig, true
}

// HasEntitlementRequestConfig returns a boolean if a field has been set.
func (o *AccessRequestConfig) HasEntitlementRequestConfig() bool {
	if o != nil && !isNil(o.EntitlementRequestConfig) {
		return true
	}

	return false
}

// SetEntitlementRequestConfig gets a reference to the given EntitlementRequestConfig1 and assigns it to the EntitlementRequestConfig field.
func (o *AccessRequestConfig) SetEntitlementRequestConfig(v EntitlementRequestConfig1) {
	o.EntitlementRequestConfig = &v
}

func (o AccessRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApprovalsMustBeExternal) {
		toSerialize["approvalsMustBeExternal"] = o.ApprovalsMustBeExternal
	}
	if !isNil(o.AutoApprovalEnabled) {
		toSerialize["autoApprovalEnabled"] = o.AutoApprovalEnabled
	}
	if !isNil(o.RequestOnBehalfOfConfig) {
		toSerialize["requestOnBehalfOfConfig"] = o.RequestOnBehalfOfConfig
	}
	if !isNil(o.ApprovalReminderAndEscalationConfig) {
		toSerialize["approvalReminderAndEscalationConfig"] = o.ApprovalReminderAndEscalationConfig
	}
	if !isNil(o.EntitlementRequestConfig) {
		toSerialize["entitlementRequestConfig"] = o.EntitlementRequestConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestConfig) UnmarshalJSON(bytes []byte) (err error) {
	varAccessRequestConfig := _AccessRequestConfig{}

	if err = json.Unmarshal(bytes, &varAccessRequestConfig); err == nil {
	*o = AccessRequestConfig(varAccessRequestConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approvalsMustBeExternal")
		delete(additionalProperties, "autoApprovalEnabled")
		delete(additionalProperties, "requestOnBehalfOfConfig")
		delete(additionalProperties, "approvalReminderAndEscalationConfig")
		delete(additionalProperties, "entitlementRequestConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestConfig struct {
	value *AccessRequestConfig
	isSet bool
}

func (v NullableAccessRequestConfig) Get() *AccessRequestConfig {
	return v.value
}

func (v *NullableAccessRequestConfig) Set(val *AccessRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestConfig(val *AccessRequestConfig) *NullableAccessRequestConfig {
	return &NullableAccessRequestConfig{value: val, isSet: true}
}

func (v NullableAccessRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


