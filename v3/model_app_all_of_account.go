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

// checks if the AppAllOfAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAllOfAccount{}

// AppAllOfAccount struct for AppAllOfAccount
type AppAllOfAccount struct {
	// The SailPoint generated unique ID
	Id *string `json:"id,omitempty"`
	// The account ID generated by the source
	AccountId *string `json:"accountId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAllOfAccount AppAllOfAccount

// NewAppAllOfAccount instantiates a new AppAllOfAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAllOfAccount() *AppAllOfAccount {
	this := AppAllOfAccount{}
	return &this
}

// NewAppAllOfAccountWithDefaults instantiates a new AppAllOfAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAllOfAccountWithDefaults() *AppAllOfAccount {
	this := AppAllOfAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppAllOfAccount) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAllOfAccount) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppAllOfAccount) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppAllOfAccount) SetId(v string) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AppAllOfAccount) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAllOfAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AppAllOfAccount) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AppAllOfAccount) SetAccountId(v string) {
	o.AccountId = &v
}

func (o AppAllOfAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAllOfAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAllOfAccount) UnmarshalJSON(bytes []byte) (err error) {
	varAppAllOfAccount := _AppAllOfAccount{}

	if err = json.Unmarshal(bytes, &varAppAllOfAccount); err == nil {
	*o = AppAllOfAccount(varAppAllOfAccount)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "accountId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAllOfAccount struct {
	value *AppAllOfAccount
	isSet bool
}

func (v NullableAppAllOfAccount) Get() *AppAllOfAccount {
	return v.value
}

func (v *NullableAppAllOfAccount) Set(val *AppAllOfAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAllOfAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAllOfAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAllOfAccount(val *AppAllOfAccount) *NullableAppAllOfAccount {
	return &NullableAppAllOfAccount{value: val, isSet: true}
}

func (v NullableAppAllOfAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAllOfAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


