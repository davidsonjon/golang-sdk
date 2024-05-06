/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the PasswordInfoAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordInfoAccount{}

// PasswordInfoAccount struct for PasswordInfoAccount
type PasswordInfoAccount struct {
	// Account ID of the account. This is specified per account schema in the source configuration. It is used to distinguish accounts. More info can be found here https://community.sailpoint.com/t5/IdentityNow-Connectors/How-do-I-designate-an-account-attribute-as-the-Account-ID-for-a/ta-p/80350
	AccountId *string `json:"accountId,omitempty"`
	// Display name of the account. This is specified per account schema in the source configuration. It is used to display name of the account. More info can be found here https://community.sailpoint.com/t5/IdentityNow-Connectors/How-do-I-designate-an-account-attribute-as-the-Account-Name-for/ta-p/74008
	AccountName *string `json:"accountName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordInfoAccount PasswordInfoAccount

// NewPasswordInfoAccount instantiates a new PasswordInfoAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordInfoAccount() *PasswordInfoAccount {
	this := PasswordInfoAccount{}
	return &this
}

// NewPasswordInfoAccountWithDefaults instantiates a new PasswordInfoAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordInfoAccountWithDefaults() *PasswordInfoAccount {
	this := PasswordInfoAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PasswordInfoAccount) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfoAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PasswordInfoAccount) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PasswordInfoAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *PasswordInfoAccount) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfoAccount) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *PasswordInfoAccount) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *PasswordInfoAccount) SetAccountName(v string) {
	o.AccountName = &v
}

func (o PasswordInfoAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordInfoAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordInfoAccount) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordInfoAccount := _PasswordInfoAccount{}

	if err = json.Unmarshal(bytes, &varPasswordInfoAccount); err == nil {
	*o = PasswordInfoAccount(varPasswordInfoAccount)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "accountName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordInfoAccount struct {
	value *PasswordInfoAccount
	isSet bool
}

func (v NullablePasswordInfoAccount) Get() *PasswordInfoAccount {
	return v.value
}

func (v *NullablePasswordInfoAccount) Set(val *PasswordInfoAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordInfoAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordInfoAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordInfoAccount(val *PasswordInfoAccount) *NullablePasswordInfoAccount {
	return &NullablePasswordInfoAccount{value: val, isSet: true}
}

func (v NullablePasswordInfoAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordInfoAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


