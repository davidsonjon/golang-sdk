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

// checks if the NativeChangeDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NativeChangeDetectionConfig{}

// NativeChangeDetectionConfig Source configuration information for Native Change Detection that is read and used by account aggregation process.
type NativeChangeDetectionConfig struct {
	// A flag indicating if Native Change Detection is enabled for a source.
	Enabled *bool `json:"enabled,omitempty"`
	// Operation types for which Native Change Detection is enabled for a source.
	Operations []string `json:"operations,omitempty"`
	// A flag indicating that all entitlements participate in Native Change Detection.
	AllEntitlements *bool `json:"allEntitlements,omitempty"`
	// A flag indicating that all non-entitlement account attributes participate in Native Change Detection.
	AllNonEntitlementAttributes *bool `json:"allNonEntitlementAttributes,omitempty"`
	// If allEntitlements flag is off this field lists entitlements that participate in Native Change Detection.
	SelectedEntitlements []string `json:"selectedEntitlements,omitempty"`
	// If allNonEntitlementAttributes flag is off this field lists non-entitlement account attributes that participate in Native Change Detection.
	SelectedNonEntitlementAttributes []string `json:"selectedNonEntitlementAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NativeChangeDetectionConfig NativeChangeDetectionConfig

// NewNativeChangeDetectionConfig instantiates a new NativeChangeDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeChangeDetectionConfig() *NativeChangeDetectionConfig {
	this := NativeChangeDetectionConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	var allEntitlements bool = false
	this.AllEntitlements = &allEntitlements
	var allNonEntitlementAttributes bool = false
	this.AllNonEntitlementAttributes = &allNonEntitlementAttributes
	return &this
}

// NewNativeChangeDetectionConfigWithDefaults instantiates a new NativeChangeDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeChangeDetectionConfigWithDefaults() *NativeChangeDetectionConfig {
	this := NativeChangeDetectionConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	var allEntitlements bool = false
	this.AllEntitlements = &allEntitlements
	var allNonEntitlementAttributes bool = false
	this.AllNonEntitlementAttributes = &allNonEntitlementAttributes
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NativeChangeDetectionConfig) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeChangeDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NativeChangeDetectionConfig) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NativeChangeDetectionConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *NativeChangeDetectionConfig) GetOperations() []string {
	if o == nil || isNil(o.Operations) {
		var ret []string
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeChangeDetectionConfig) GetOperationsOk() ([]string, bool) {
	if o == nil || isNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *NativeChangeDetectionConfig) HasOperations() bool {
	if o != nil && !isNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []string and assigns it to the Operations field.
func (o *NativeChangeDetectionConfig) SetOperations(v []string) {
	o.Operations = v
}

// GetAllEntitlements returns the AllEntitlements field value if set, zero value otherwise.
func (o *NativeChangeDetectionConfig) GetAllEntitlements() bool {
	if o == nil || isNil(o.AllEntitlements) {
		var ret bool
		return ret
	}
	return *o.AllEntitlements
}

// GetAllEntitlementsOk returns a tuple with the AllEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeChangeDetectionConfig) GetAllEntitlementsOk() (*bool, bool) {
	if o == nil || isNil(o.AllEntitlements) {
		return nil, false
	}
	return o.AllEntitlements, true
}

// HasAllEntitlements returns a boolean if a field has been set.
func (o *NativeChangeDetectionConfig) HasAllEntitlements() bool {
	if o != nil && !isNil(o.AllEntitlements) {
		return true
	}

	return false
}

// SetAllEntitlements gets a reference to the given bool and assigns it to the AllEntitlements field.
func (o *NativeChangeDetectionConfig) SetAllEntitlements(v bool) {
	o.AllEntitlements = &v
}

// GetAllNonEntitlementAttributes returns the AllNonEntitlementAttributes field value if set, zero value otherwise.
func (o *NativeChangeDetectionConfig) GetAllNonEntitlementAttributes() bool {
	if o == nil || isNil(o.AllNonEntitlementAttributes) {
		var ret bool
		return ret
	}
	return *o.AllNonEntitlementAttributes
}

// GetAllNonEntitlementAttributesOk returns a tuple with the AllNonEntitlementAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeChangeDetectionConfig) GetAllNonEntitlementAttributesOk() (*bool, bool) {
	if o == nil || isNil(o.AllNonEntitlementAttributes) {
		return nil, false
	}
	return o.AllNonEntitlementAttributes, true
}

// HasAllNonEntitlementAttributes returns a boolean if a field has been set.
func (o *NativeChangeDetectionConfig) HasAllNonEntitlementAttributes() bool {
	if o != nil && !isNil(o.AllNonEntitlementAttributes) {
		return true
	}

	return false
}

// SetAllNonEntitlementAttributes gets a reference to the given bool and assigns it to the AllNonEntitlementAttributes field.
func (o *NativeChangeDetectionConfig) SetAllNonEntitlementAttributes(v bool) {
	o.AllNonEntitlementAttributes = &v
}

// GetSelectedEntitlements returns the SelectedEntitlements field value if set, zero value otherwise.
func (o *NativeChangeDetectionConfig) GetSelectedEntitlements() []string {
	if o == nil || isNil(o.SelectedEntitlements) {
		var ret []string
		return ret
	}
	return o.SelectedEntitlements
}

// GetSelectedEntitlementsOk returns a tuple with the SelectedEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeChangeDetectionConfig) GetSelectedEntitlementsOk() ([]string, bool) {
	if o == nil || isNil(o.SelectedEntitlements) {
		return nil, false
	}
	return o.SelectedEntitlements, true
}

// HasSelectedEntitlements returns a boolean if a field has been set.
func (o *NativeChangeDetectionConfig) HasSelectedEntitlements() bool {
	if o != nil && !isNil(o.SelectedEntitlements) {
		return true
	}

	return false
}

// SetSelectedEntitlements gets a reference to the given []string and assigns it to the SelectedEntitlements field.
func (o *NativeChangeDetectionConfig) SetSelectedEntitlements(v []string) {
	o.SelectedEntitlements = v
}

// GetSelectedNonEntitlementAttributes returns the SelectedNonEntitlementAttributes field value if set, zero value otherwise.
func (o *NativeChangeDetectionConfig) GetSelectedNonEntitlementAttributes() []string {
	if o == nil || isNil(o.SelectedNonEntitlementAttributes) {
		var ret []string
		return ret
	}
	return o.SelectedNonEntitlementAttributes
}

// GetSelectedNonEntitlementAttributesOk returns a tuple with the SelectedNonEntitlementAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeChangeDetectionConfig) GetSelectedNonEntitlementAttributesOk() ([]string, bool) {
	if o == nil || isNil(o.SelectedNonEntitlementAttributes) {
		return nil, false
	}
	return o.SelectedNonEntitlementAttributes, true
}

// HasSelectedNonEntitlementAttributes returns a boolean if a field has been set.
func (o *NativeChangeDetectionConfig) HasSelectedNonEntitlementAttributes() bool {
	if o != nil && !isNil(o.SelectedNonEntitlementAttributes) {
		return true
	}

	return false
}

// SetSelectedNonEntitlementAttributes gets a reference to the given []string and assigns it to the SelectedNonEntitlementAttributes field.
func (o *NativeChangeDetectionConfig) SetSelectedNonEntitlementAttributes(v []string) {
	o.SelectedNonEntitlementAttributes = v
}

func (o NativeChangeDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NativeChangeDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !isNil(o.AllEntitlements) {
		toSerialize["allEntitlements"] = o.AllEntitlements
	}
	if !isNil(o.AllNonEntitlementAttributes) {
		toSerialize["allNonEntitlementAttributes"] = o.AllNonEntitlementAttributes
	}
	if !isNil(o.SelectedEntitlements) {
		toSerialize["selectedEntitlements"] = o.SelectedEntitlements
	}
	if !isNil(o.SelectedNonEntitlementAttributes) {
		toSerialize["selectedNonEntitlementAttributes"] = o.SelectedNonEntitlementAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NativeChangeDetectionConfig) UnmarshalJSON(bytes []byte) (err error) {
	varNativeChangeDetectionConfig := _NativeChangeDetectionConfig{}

	if err = json.Unmarshal(bytes, &varNativeChangeDetectionConfig); err == nil {
	*o = NativeChangeDetectionConfig(varNativeChangeDetectionConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "operations")
		delete(additionalProperties, "allEntitlements")
		delete(additionalProperties, "allNonEntitlementAttributes")
		delete(additionalProperties, "selectedEntitlements")
		delete(additionalProperties, "selectedNonEntitlementAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNativeChangeDetectionConfig struct {
	value *NativeChangeDetectionConfig
	isSet bool
}

func (v NullableNativeChangeDetectionConfig) Get() *NativeChangeDetectionConfig {
	return v.value
}

func (v *NullableNativeChangeDetectionConfig) Set(val *NativeChangeDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeChangeDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeChangeDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeChangeDetectionConfig(val *NativeChangeDetectionConfig) *NullableNativeChangeDetectionConfig {
	return &NullableNativeChangeDetectionConfig{value: val, isSet: true}
}

func (v NullableNativeChangeDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeChangeDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

