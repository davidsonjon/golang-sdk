/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the PublicIdentityConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIdentityConfig{}

// PublicIdentityConfig Details of up to 5 Identity attributes that will be publicly accessible for all Identities to anyone in the org
type PublicIdentityConfig struct {
	Attributes []PublicIdentityAttributeConfig `json:"attributes,omitempty"`
	ModifiedBy NullableIdentityReference `json:"modifiedBy,omitempty"`
	// the date/time of the modification
	Modified NullableTime `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicIdentityConfig PublicIdentityConfig

// NewPublicIdentityConfig instantiates a new PublicIdentityConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIdentityConfig() *PublicIdentityConfig {
	this := PublicIdentityConfig{}
	return &this
}

// NewPublicIdentityConfigWithDefaults instantiates a new PublicIdentityConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIdentityConfigWithDefaults() *PublicIdentityConfig {
	this := PublicIdentityConfig{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PublicIdentityConfig) GetAttributes() []PublicIdentityAttributeConfig {
	if o == nil || IsNil(o.Attributes) {
		var ret []PublicIdentityAttributeConfig
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentityConfig) GetAttributesOk() ([]PublicIdentityAttributeConfig, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PublicIdentityConfig) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []PublicIdentityAttributeConfig and assigns it to the Attributes field.
func (o *PublicIdentityConfig) SetAttributes(v []PublicIdentityAttributeConfig) {
	o.Attributes = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIdentityConfig) GetModifiedBy() IdentityReference {
	if o == nil || IsNil(o.ModifiedBy.Get()) {
		var ret IdentityReference
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIdentityConfig) GetModifiedByOk() (*IdentityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *PublicIdentityConfig) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableIdentityReference and assigns it to the ModifiedBy field.
func (o *PublicIdentityConfig) SetModifiedBy(v IdentityReference) {
	o.ModifiedBy.Set(&v)
}
// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *PublicIdentityConfig) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *PublicIdentityConfig) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIdentityConfig) GetModified() time.Time {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIdentityConfig) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *PublicIdentityConfig) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *PublicIdentityConfig) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *PublicIdentityConfig) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *PublicIdentityConfig) UnsetModified() {
	o.Modified.Unset()
}

func (o PublicIdentityConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicIdentityConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modifiedBy"] = o.ModifiedBy.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicIdentityConfig) UnmarshalJSON(data []byte) (err error) {
	varPublicIdentityConfig := _PublicIdentityConfig{}

	err = json.Unmarshal(data, &varPublicIdentityConfig)

	if err != nil {
		return err
	}

	*o = PublicIdentityConfig(varPublicIdentityConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "modifiedBy")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicIdentityConfig struct {
	value *PublicIdentityConfig
	isSet bool
}

func (v NullablePublicIdentityConfig) Get() *PublicIdentityConfig {
	return v.value
}

func (v *NullablePublicIdentityConfig) Set(val *PublicIdentityConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIdentityConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIdentityConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIdentityConfig(val *PublicIdentityConfig) *NullablePublicIdentityConfig {
	return &NullablePublicIdentityConfig{value: val, isSet: true}
}

func (v NullablePublicIdentityConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIdentityConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


