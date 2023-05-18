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

// checks if the PublicKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicKey{}

// PublicKey struct for PublicKey
type PublicKey struct {
	// ARM Public Key used to encrypt username and password credentials sent to ARM
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicKey PublicKey

// NewPublicKey instantiates a new PublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicKey() *PublicKey {
	this := PublicKey{}
	return &this
}

// NewPublicKeyWithDefaults instantiates a new PublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicKeyWithDefaults() *PublicKey {
	this := PublicKey{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PublicKey) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKey) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PublicKey) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PublicKey) SetValue(v string) {
	o.Value = &v
}

func (o PublicKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicKey) UnmarshalJSON(bytes []byte) (err error) {
	varPublicKey := _PublicKey{}

	if err = json.Unmarshal(bytes, &varPublicKey); err == nil {
		*o = PublicKey(varPublicKey)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicKey struct {
	value *PublicKey
	isSet bool
}

func (v NullablePublicKey) Get() *PublicKey {
	return v.value
}

func (v *NullablePublicKey) Set(val *PublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicKey(val *PublicKey) *NullablePublicKey {
	return &NullablePublicKey{value: val, isSet: true}
}

func (v NullablePublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


