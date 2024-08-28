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

// checks if the SearchAttributeConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAttributeConfig{}

// SearchAttributeConfig struct for SearchAttributeConfig
type SearchAttributeConfig struct {
	// Name of the new attribute
	Name *string `json:"name,omitempty"`
	// The display name of the new attribute
	DisplayName *string `json:"displayName,omitempty"`
	// Map of application id and their associated attribute.
	ApplicationAttributes map[string]interface{} `json:"applicationAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchAttributeConfig SearchAttributeConfig

// NewSearchAttributeConfig instantiates a new SearchAttributeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAttributeConfig() *SearchAttributeConfig {
	this := SearchAttributeConfig{}
	return &this
}

// NewSearchAttributeConfigWithDefaults instantiates a new SearchAttributeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAttributeConfigWithDefaults() *SearchAttributeConfig {
	this := SearchAttributeConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchAttributeConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAttributeConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchAttributeConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchAttributeConfig) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SearchAttributeConfig) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAttributeConfig) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SearchAttributeConfig) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SearchAttributeConfig) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetApplicationAttributes returns the ApplicationAttributes field value if set, zero value otherwise.
func (o *SearchAttributeConfig) GetApplicationAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ApplicationAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ApplicationAttributes
}

// GetApplicationAttributesOk returns a tuple with the ApplicationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAttributeConfig) GetApplicationAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ApplicationAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ApplicationAttributes, true
}

// HasApplicationAttributes returns a boolean if a field has been set.
func (o *SearchAttributeConfig) HasApplicationAttributes() bool {
	if o != nil && !IsNil(o.ApplicationAttributes) {
		return true
	}

	return false
}

// SetApplicationAttributes gets a reference to the given map[string]interface{} and assigns it to the ApplicationAttributes field.
func (o *SearchAttributeConfig) SetApplicationAttributes(v map[string]interface{}) {
	o.ApplicationAttributes = v
}

func (o SearchAttributeConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAttributeConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ApplicationAttributes) {
		toSerialize["applicationAttributes"] = o.ApplicationAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchAttributeConfig) UnmarshalJSON(data []byte) (err error) {
	varSearchAttributeConfig := _SearchAttributeConfig{}

	err = json.Unmarshal(data, &varSearchAttributeConfig)

	if err != nil {
		return err
	}

	*o = SearchAttributeConfig(varSearchAttributeConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "applicationAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchAttributeConfig struct {
	value *SearchAttributeConfig
	isSet bool
}

func (v NullableSearchAttributeConfig) Get() *SearchAttributeConfig {
	return v.value
}

func (v *NullableSearchAttributeConfig) Set(val *SearchAttributeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAttributeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAttributeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAttributeConfig(val *SearchAttributeConfig) *NullableSearchAttributeConfig {
	return &NullableSearchAttributeConfig{value: val, isSet: true}
}

func (v NullableSearchAttributeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAttributeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


