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

// checks if the MfaOktaConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MfaOktaConfig{}

// MfaOktaConfig struct for MfaOktaConfig
type MfaOktaConfig struct {
	// Mfa method name
	MfaMethod NullableString `json:"mfaMethod,omitempty"`
	// If MFA method is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The server host name or IP address of the MFA provider.
	Host NullableString `json:"host,omitempty"`
	// The secret key for authenticating requests to the MFA provider.
	AccessKey NullableString `json:"accessKey,omitempty"`
	// Optional. The name of the attribute for mapping IdentityNow identity to the MFA provider.
	IdentityAttribute NullableString `json:"identityAttribute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MfaOktaConfig MfaOktaConfig

// NewMfaOktaConfig instantiates a new MfaOktaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaOktaConfig() *MfaOktaConfig {
	this := MfaOktaConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewMfaOktaConfigWithDefaults instantiates a new MfaOktaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaOktaConfigWithDefaults() *MfaOktaConfig {
	this := MfaOktaConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetMfaMethod returns the MfaMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MfaOktaConfig) GetMfaMethod() string {
	if o == nil || isNil(o.MfaMethod.Get()) {
		var ret string
		return ret
	}
	return *o.MfaMethod.Get()
}

// GetMfaMethodOk returns a tuple with the MfaMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MfaOktaConfig) GetMfaMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MfaMethod.Get(), o.MfaMethod.IsSet()
}

// HasMfaMethod returns a boolean if a field has been set.
func (o *MfaOktaConfig) HasMfaMethod() bool {
	if o != nil && o.MfaMethod.IsSet() {
		return true
	}

	return false
}

// SetMfaMethod gets a reference to the given NullableString and assigns it to the MfaMethod field.
func (o *MfaOktaConfig) SetMfaMethod(v string) {
	o.MfaMethod.Set(&v)
}
// SetMfaMethodNil sets the value for MfaMethod to be an explicit nil
func (o *MfaOktaConfig) SetMfaMethodNil() {
	o.MfaMethod.Set(nil)
}

// UnsetMfaMethod ensures that no value is present for MfaMethod, not even an explicit nil
func (o *MfaOktaConfig) UnsetMfaMethod() {
	o.MfaMethod.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MfaOktaConfig) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaOktaConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MfaOktaConfig) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MfaOktaConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MfaOktaConfig) GetHost() string {
	if o == nil || isNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MfaOktaConfig) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *MfaOktaConfig) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *MfaOktaConfig) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *MfaOktaConfig) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *MfaOktaConfig) UnsetHost() {
	o.Host.Unset()
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MfaOktaConfig) GetAccessKey() string {
	if o == nil || isNil(o.AccessKey.Get()) {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MfaOktaConfig) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *MfaOktaConfig) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *MfaOktaConfig) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}
// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *MfaOktaConfig) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *MfaOktaConfig) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetIdentityAttribute returns the IdentityAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MfaOktaConfig) GetIdentityAttribute() string {
	if o == nil || isNil(o.IdentityAttribute.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityAttribute.Get()
}

// GetIdentityAttributeOk returns a tuple with the IdentityAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MfaOktaConfig) GetIdentityAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityAttribute.Get(), o.IdentityAttribute.IsSet()
}

// HasIdentityAttribute returns a boolean if a field has been set.
func (o *MfaOktaConfig) HasIdentityAttribute() bool {
	if o != nil && o.IdentityAttribute.IsSet() {
		return true
	}

	return false
}

// SetIdentityAttribute gets a reference to the given NullableString and assigns it to the IdentityAttribute field.
func (o *MfaOktaConfig) SetIdentityAttribute(v string) {
	o.IdentityAttribute.Set(&v)
}
// SetIdentityAttributeNil sets the value for IdentityAttribute to be an explicit nil
func (o *MfaOktaConfig) SetIdentityAttributeNil() {
	o.IdentityAttribute.Set(nil)
}

// UnsetIdentityAttribute ensures that no value is present for IdentityAttribute, not even an explicit nil
func (o *MfaOktaConfig) UnsetIdentityAttribute() {
	o.IdentityAttribute.Unset()
}

func (o MfaOktaConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MfaOktaConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MfaMethod.IsSet() {
		toSerialize["mfaMethod"] = o.MfaMethod.Get()
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.AccessKey.IsSet() {
		toSerialize["accessKey"] = o.AccessKey.Get()
	}
	if o.IdentityAttribute.IsSet() {
		toSerialize["identityAttribute"] = o.IdentityAttribute.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MfaOktaConfig) UnmarshalJSON(bytes []byte) (err error) {
	varMfaOktaConfig := _MfaOktaConfig{}

	if err = json.Unmarshal(bytes, &varMfaOktaConfig); err == nil {
	*o = MfaOktaConfig(varMfaOktaConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "mfaMethod")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "host")
		delete(additionalProperties, "accessKey")
		delete(additionalProperties, "identityAttribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMfaOktaConfig struct {
	value *MfaOktaConfig
	isSet bool
}

func (v NullableMfaOktaConfig) Get() *MfaOktaConfig {
	return v.value
}

func (v *NullableMfaOktaConfig) Set(val *MfaOktaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMfaOktaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMfaOktaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMfaOktaConfig(val *MfaOktaConfig) *NullableMfaOktaConfig {
	return &NullableMfaOktaConfig{value: val, isSet: true}
}

func (v NullableMfaOktaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMfaOktaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


