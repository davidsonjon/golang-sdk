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

// checks if the MfaConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MfaConfig{}

// MfaConfig struct for MfaConfig
type MfaConfig struct {
	// If MFA method is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The server host name or IP address of the MFA provider.
	Host *string `json:"host,omitempty"`
	// The secret key for authenticating requests to the MFA provider.
	AccessKey *string `json:"accessKey,omitempty"`
	// Optional. The name of the attribute for mapping IdentityNow identity to the MFA provider.
	IdentityAttribute *string `json:"identityAttribute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MfaConfig MfaConfig

// NewMfaConfig instantiates a new MfaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaConfig() *MfaConfig {
	this := MfaConfig{}
	return &this
}

// NewMfaConfigWithDefaults instantiates a new MfaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaConfigWithDefaults() *MfaConfig {
	this := MfaConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MfaConfig) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MfaConfig) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MfaConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *MfaConfig) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfig) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *MfaConfig) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *MfaConfig) SetHost(v string) {
	o.Host = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *MfaConfig) GetAccessKey() string {
	if o == nil || isNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfig) GetAccessKeyOk() (*string, bool) {
	if o == nil || isNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *MfaConfig) HasAccessKey() bool {
	if o != nil && !isNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *MfaConfig) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetIdentityAttribute returns the IdentityAttribute field value if set, zero value otherwise.
func (o *MfaConfig) GetIdentityAttribute() string {
	if o == nil || isNil(o.IdentityAttribute) {
		var ret string
		return ret
	}
	return *o.IdentityAttribute
}

// GetIdentityAttributeOk returns a tuple with the IdentityAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfig) GetIdentityAttributeOk() (*string, bool) {
	if o == nil || isNil(o.IdentityAttribute) {
		return nil, false
	}
	return o.IdentityAttribute, true
}

// HasIdentityAttribute returns a boolean if a field has been set.
func (o *MfaConfig) HasIdentityAttribute() bool {
	if o != nil && !isNil(o.IdentityAttribute) {
		return true
	}

	return false
}

// SetIdentityAttribute gets a reference to the given string and assigns it to the IdentityAttribute field.
func (o *MfaConfig) SetIdentityAttribute(v string) {
	o.IdentityAttribute = &v
}

func (o MfaConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MfaConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.AccessKey) {
		toSerialize["accessKey"] = o.AccessKey
	}
	if !isNil(o.IdentityAttribute) {
		toSerialize["identityAttribute"] = o.IdentityAttribute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MfaConfig) UnmarshalJSON(bytes []byte) (err error) {
	varMfaConfig := _MfaConfig{}

	if err = json.Unmarshal(bytes, &varMfaConfig); err == nil {
		*o = MfaConfig(varMfaConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "host")
		delete(additionalProperties, "accessKey")
		delete(additionalProperties, "identityAttribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMfaConfig struct {
	value *MfaConfig
	isSet bool
}

func (v NullableMfaConfig) Get() *MfaConfig {
	return v.value
}

func (v *NullableMfaConfig) Set(val *MfaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMfaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMfaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMfaConfig(val *MfaConfig) *NullableMfaConfig {
	return &NullableMfaConfig{value: val, isSet: true}
}

func (v NullableMfaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMfaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


