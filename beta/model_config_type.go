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

// checks if the ConfigType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigType{}

// ConfigType Type of Reassignment Configuration.
type ConfigType struct {
	InternalName *ConfigTypeEnum `json:"internalName,omitempty"`
	// Human readable display name of the type to be shown on UI
	DisplayName *string `json:"displayName,omitempty"`
	// Description of the type of work to be reassigned, displayed by the UI.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigType ConfigType

// NewConfigType instantiates a new ConfigType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigType() *ConfigType {
	this := ConfigType{}
	return &this
}

// NewConfigTypeWithDefaults instantiates a new ConfigType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigTypeWithDefaults() *ConfigType {
	this := ConfigType{}
	return &this
}

// GetInternalName returns the InternalName field value if set, zero value otherwise.
func (o *ConfigType) GetInternalName() ConfigTypeEnum {
	if o == nil || isNil(o.InternalName) {
		var ret ConfigTypeEnum
		return ret
	}
	return *o.InternalName
}

// GetInternalNameOk returns a tuple with the InternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigType) GetInternalNameOk() (*ConfigTypeEnum, bool) {
	if o == nil || isNil(o.InternalName) {
		return nil, false
	}
	return o.InternalName, true
}

// HasInternalName returns a boolean if a field has been set.
func (o *ConfigType) HasInternalName() bool {
	if o != nil && !isNil(o.InternalName) {
		return true
	}

	return false
}

// SetInternalName gets a reference to the given ConfigTypeEnum and assigns it to the InternalName field.
func (o *ConfigType) SetInternalName(v ConfigTypeEnum) {
	o.InternalName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConfigType) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigType) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConfigType) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConfigType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigType) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigType) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigType) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigType) SetDescription(v string) {
	o.Description = &v
}

func (o ConfigType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InternalName) {
		toSerialize["internalName"] = o.InternalName
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigType) UnmarshalJSON(bytes []byte) (err error) {
	varConfigType := _ConfigType{}

	if err = json.Unmarshal(bytes, &varConfigType); err == nil {
		*o = ConfigType(varConfigType)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "internalName")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigType struct {
	value *ConfigType
	isSet bool
}

func (v NullableConfigType) Get() *ConfigType {
	return v.value
}

func (v *NullableConfigType) Set(val *ConfigType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigType(val *ConfigType) *NullableConfigType {
	return &NullableConfigType{value: val, isSet: true}
}

func (v NullableConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


