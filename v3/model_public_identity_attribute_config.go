/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the PublicIdentityAttributeConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIdentityAttributeConfig{}

// PublicIdentityAttributeConfig Used to map an attribute key for an Identity to its display name.
type PublicIdentityAttributeConfig struct {
	// The attribute key
	Key *string `json:"key,omitempty"`
	// The attribute display name
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicIdentityAttributeConfig PublicIdentityAttributeConfig

// NewPublicIdentityAttributeConfig instantiates a new PublicIdentityAttributeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIdentityAttributeConfig() *PublicIdentityAttributeConfig {
	this := PublicIdentityAttributeConfig{}
	return &this
}

// NewPublicIdentityAttributeConfigWithDefaults instantiates a new PublicIdentityAttributeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIdentityAttributeConfigWithDefaults() *PublicIdentityAttributeConfig {
	this := PublicIdentityAttributeConfig{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PublicIdentityAttributeConfig) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentityAttributeConfig) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PublicIdentityAttributeConfig) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PublicIdentityAttributeConfig) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicIdentityAttributeConfig) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentityAttributeConfig) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicIdentityAttributeConfig) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicIdentityAttributeConfig) SetName(v string) {
	o.Name = &v
}

func (o PublicIdentityAttributeConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicIdentityAttributeConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicIdentityAttributeConfig) UnmarshalJSON(bytes []byte) (err error) {
	varPublicIdentityAttributeConfig := _PublicIdentityAttributeConfig{}

	if err = json.Unmarshal(bytes, &varPublicIdentityAttributeConfig); err == nil {
	*o = PublicIdentityAttributeConfig(varPublicIdentityAttributeConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicIdentityAttributeConfig struct {
	value *PublicIdentityAttributeConfig
	isSet bool
}

func (v NullablePublicIdentityAttributeConfig) Get() *PublicIdentityAttributeConfig {
	return v.value
}

func (v *NullablePublicIdentityAttributeConfig) Set(val *PublicIdentityAttributeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIdentityAttributeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIdentityAttributeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIdentityAttributeConfig(val *PublicIdentityAttributeConfig) *NullablePublicIdentityAttributeConfig {
	return &NullablePublicIdentityAttributeConfig{value: val, isSet: true}
}

func (v NullablePublicIdentityAttributeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIdentityAttributeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


