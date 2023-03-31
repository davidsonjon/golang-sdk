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

// checks if the PasswordOrgConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordOrgConfig{}

// PasswordOrgConfig struct for PasswordOrgConfig
type PasswordOrgConfig struct {
	// Indicator whether custom password instructions feature is enabled. The default value is false.
	CustomInstructionsEnabled *bool `json:"customInstructionsEnabled,omitempty"`
	// Indicator whether \"digit token\" feature is enabled. The default value is false.
	DigitTokenEnabled *bool `json:"digitTokenEnabled,omitempty"`
	// The duration of \"digit token\" in minutes. The default value is 5.
	DigitTokenDurationMinutes *int32 `json:"digitTokenDurationMinutes,omitempty"`
	// The length of \"digit token\". The default value is 6.
	DigitTokenLength *int32 `json:"digitTokenLength,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordOrgConfig PasswordOrgConfig

// NewPasswordOrgConfig instantiates a new PasswordOrgConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordOrgConfig() *PasswordOrgConfig {
	this := PasswordOrgConfig{}
	return &this
}

// NewPasswordOrgConfigWithDefaults instantiates a new PasswordOrgConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordOrgConfigWithDefaults() *PasswordOrgConfig {
	this := PasswordOrgConfig{}
	return &this
}

// GetCustomInstructionsEnabled returns the CustomInstructionsEnabled field value if set, zero value otherwise.
func (o *PasswordOrgConfig) GetCustomInstructionsEnabled() bool {
	if o == nil || isNil(o.CustomInstructionsEnabled) {
		var ret bool
		return ret
	}
	return *o.CustomInstructionsEnabled
}

// GetCustomInstructionsEnabledOk returns a tuple with the CustomInstructionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordOrgConfig) GetCustomInstructionsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.CustomInstructionsEnabled) {
		return nil, false
	}
	return o.CustomInstructionsEnabled, true
}

// HasCustomInstructionsEnabled returns a boolean if a field has been set.
func (o *PasswordOrgConfig) HasCustomInstructionsEnabled() bool {
	if o != nil && !isNil(o.CustomInstructionsEnabled) {
		return true
	}

	return false
}

// SetCustomInstructionsEnabled gets a reference to the given bool and assigns it to the CustomInstructionsEnabled field.
func (o *PasswordOrgConfig) SetCustomInstructionsEnabled(v bool) {
	o.CustomInstructionsEnabled = &v
}

// GetDigitTokenEnabled returns the DigitTokenEnabled field value if set, zero value otherwise.
func (o *PasswordOrgConfig) GetDigitTokenEnabled() bool {
	if o == nil || isNil(o.DigitTokenEnabled) {
		var ret bool
		return ret
	}
	return *o.DigitTokenEnabled
}

// GetDigitTokenEnabledOk returns a tuple with the DigitTokenEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordOrgConfig) GetDigitTokenEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DigitTokenEnabled) {
		return nil, false
	}
	return o.DigitTokenEnabled, true
}

// HasDigitTokenEnabled returns a boolean if a field has been set.
func (o *PasswordOrgConfig) HasDigitTokenEnabled() bool {
	if o != nil && !isNil(o.DigitTokenEnabled) {
		return true
	}

	return false
}

// SetDigitTokenEnabled gets a reference to the given bool and assigns it to the DigitTokenEnabled field.
func (o *PasswordOrgConfig) SetDigitTokenEnabled(v bool) {
	o.DigitTokenEnabled = &v
}

// GetDigitTokenDurationMinutes returns the DigitTokenDurationMinutes field value if set, zero value otherwise.
func (o *PasswordOrgConfig) GetDigitTokenDurationMinutes() int32 {
	if o == nil || isNil(o.DigitTokenDurationMinutes) {
		var ret int32
		return ret
	}
	return *o.DigitTokenDurationMinutes
}

// GetDigitTokenDurationMinutesOk returns a tuple with the DigitTokenDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordOrgConfig) GetDigitTokenDurationMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.DigitTokenDurationMinutes) {
		return nil, false
	}
	return o.DigitTokenDurationMinutes, true
}

// HasDigitTokenDurationMinutes returns a boolean if a field has been set.
func (o *PasswordOrgConfig) HasDigitTokenDurationMinutes() bool {
	if o != nil && !isNil(o.DigitTokenDurationMinutes) {
		return true
	}

	return false
}

// SetDigitTokenDurationMinutes gets a reference to the given int32 and assigns it to the DigitTokenDurationMinutes field.
func (o *PasswordOrgConfig) SetDigitTokenDurationMinutes(v int32) {
	o.DigitTokenDurationMinutes = &v
}

// GetDigitTokenLength returns the DigitTokenLength field value if set, zero value otherwise.
func (o *PasswordOrgConfig) GetDigitTokenLength() int32 {
	if o == nil || isNil(o.DigitTokenLength) {
		var ret int32
		return ret
	}
	return *o.DigitTokenLength
}

// GetDigitTokenLengthOk returns a tuple with the DigitTokenLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordOrgConfig) GetDigitTokenLengthOk() (*int32, bool) {
	if o == nil || isNil(o.DigitTokenLength) {
		return nil, false
	}
	return o.DigitTokenLength, true
}

// HasDigitTokenLength returns a boolean if a field has been set.
func (o *PasswordOrgConfig) HasDigitTokenLength() bool {
	if o != nil && !isNil(o.DigitTokenLength) {
		return true
	}

	return false
}

// SetDigitTokenLength gets a reference to the given int32 and assigns it to the DigitTokenLength field.
func (o *PasswordOrgConfig) SetDigitTokenLength(v int32) {
	o.DigitTokenLength = &v
}

func (o PasswordOrgConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordOrgConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CustomInstructionsEnabled) {
		toSerialize["customInstructionsEnabled"] = o.CustomInstructionsEnabled
	}
	if !isNil(o.DigitTokenEnabled) {
		toSerialize["digitTokenEnabled"] = o.DigitTokenEnabled
	}
	if !isNil(o.DigitTokenDurationMinutes) {
		toSerialize["digitTokenDurationMinutes"] = o.DigitTokenDurationMinutes
	}
	if !isNil(o.DigitTokenLength) {
		toSerialize["digitTokenLength"] = o.DigitTokenLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordOrgConfig) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordOrgConfig := _PasswordOrgConfig{}

	if err = json.Unmarshal(bytes, &varPasswordOrgConfig); err == nil {
		*o = PasswordOrgConfig(varPasswordOrgConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customInstructionsEnabled")
		delete(additionalProperties, "digitTokenEnabled")
		delete(additionalProperties, "digitTokenDurationMinutes")
		delete(additionalProperties, "digitTokenLength")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordOrgConfig struct {
	value *PasswordOrgConfig
	isSet bool
}

func (v NullablePasswordOrgConfig) Get() *PasswordOrgConfig {
	return v.value
}

func (v *NullablePasswordOrgConfig) Set(val *PasswordOrgConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordOrgConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordOrgConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordOrgConfig(val *PasswordOrgConfig) *NullablePasswordOrgConfig {
	return &NullablePasswordOrgConfig{value: val, isSet: true}
}

func (v NullablePasswordOrgConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordOrgConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


