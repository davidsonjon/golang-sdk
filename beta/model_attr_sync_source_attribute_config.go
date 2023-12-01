/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// checks if the AttrSyncSourceAttributeConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttrSyncSourceAttributeConfig{}

// AttrSyncSourceAttributeConfig Specification of source attribute sync mapping configuration for an identity attribute
type AttrSyncSourceAttributeConfig struct {
	// Name of the identity attribute
	Name string `json:"name"`
	// Display name of the identity attribute
	DisplayName string `json:"displayName"`
	// Determines whether or not the attribute is enabled for synchronization
	Enabled bool `json:"enabled"`
	// Name of the source account attribute to which the identity attribute value will be synchronized if enabled
	Target string `json:"target"`
	AdditionalProperties map[string]interface{}
}

type _AttrSyncSourceAttributeConfig AttrSyncSourceAttributeConfig

// NewAttrSyncSourceAttributeConfig instantiates a new AttrSyncSourceAttributeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttrSyncSourceAttributeConfig(name string, displayName string, enabled bool, target string) *AttrSyncSourceAttributeConfig {
	this := AttrSyncSourceAttributeConfig{}
	this.Name = name
	this.DisplayName = displayName
	this.Enabled = enabled
	this.Target = target
	return &this
}

// NewAttrSyncSourceAttributeConfigWithDefaults instantiates a new AttrSyncSourceAttributeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttrSyncSourceAttributeConfigWithDefaults() *AttrSyncSourceAttributeConfig {
	this := AttrSyncSourceAttributeConfig{}
	return &this
}

// GetName returns the Name field value
func (o *AttrSyncSourceAttributeConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttrSyncSourceAttributeConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttrSyncSourceAttributeConfig) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *AttrSyncSourceAttributeConfig) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AttrSyncSourceAttributeConfig) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AttrSyncSourceAttributeConfig) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value
func (o *AttrSyncSourceAttributeConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AttrSyncSourceAttributeConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AttrSyncSourceAttributeConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTarget returns the Target field value
func (o *AttrSyncSourceAttributeConfig) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *AttrSyncSourceAttributeConfig) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *AttrSyncSourceAttributeConfig) SetTarget(v string) {
	o.Target = v
}

func (o AttrSyncSourceAttributeConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttrSyncSourceAttributeConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["enabled"] = o.Enabled
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttrSyncSourceAttributeConfig) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"displayName",
		"enabled",
		"target",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAttrSyncSourceAttributeConfig := _AttrSyncSourceAttributeConfig{}

	if err = json.Unmarshal(bytes, &varAttrSyncSourceAttributeConfig); err == nil {
	*o = AttrSyncSourceAttributeConfig(varAttrSyncSourceAttributeConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttrSyncSourceAttributeConfig struct {
	value *AttrSyncSourceAttributeConfig
	isSet bool
}

func (v NullableAttrSyncSourceAttributeConfig) Get() *AttrSyncSourceAttributeConfig {
	return v.value
}

func (v *NullableAttrSyncSourceAttributeConfig) Set(val *AttrSyncSourceAttributeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAttrSyncSourceAttributeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAttrSyncSourceAttributeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttrSyncSourceAttributeConfig(val *AttrSyncSourceAttributeConfig) *NullableAttrSyncSourceAttributeConfig {
	return &NullableAttrSyncSourceAttributeConfig{value: val, isSet: true}
}

func (v NullableAttrSyncSourceAttributeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttrSyncSourceAttributeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


