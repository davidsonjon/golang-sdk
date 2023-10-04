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

// checks if the AppAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAllOf{}

// AppAllOf struct for AppAllOf
type AppAllOf struct {
	Source *Reference `json:"source,omitempty"`
	Account *AppAllOfAccount `json:"account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAllOf AppAllOf

// NewAppAllOf instantiates a new AppAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAllOf() *AppAllOf {
	this := AppAllOf{}
	return &this
}

// NewAppAllOfWithDefaults instantiates a new AppAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAllOfWithDefaults() *AppAllOf {
	this := AppAllOf{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AppAllOf) GetSource() Reference {
	if o == nil || isNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAllOf) GetSourceOk() (*Reference, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AppAllOf) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *AppAllOf) SetSource(v Reference) {
	o.Source = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AppAllOf) GetAccount() AppAllOfAccount {
	if o == nil || isNil(o.Account) {
		var ret AppAllOfAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAllOf) GetAccountOk() (*AppAllOfAccount, bool) {
	if o == nil || isNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AppAllOf) HasAccount() bool {
	if o != nil && !isNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AppAllOfAccount and assigns it to the Account field.
func (o *AppAllOf) SetAccount(v AppAllOfAccount) {
	o.Account = &v
}

func (o AppAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Account) {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAppAllOf := _AppAllOf{}

	if err = json.Unmarshal(bytes, &varAppAllOf); err == nil {
		*o = AppAllOf(varAppAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAllOf struct {
	value *AppAllOf
	isSet bool
}

func (v NullableAppAllOf) Get() *AppAllOf {
	return v.value
}

func (v *NullableAppAllOf) Set(val *AppAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAllOf(val *AppAllOf) *NullableAppAllOf {
	return &NullableAppAllOf{value: val, isSet: true}
}

func (v NullableAppAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


