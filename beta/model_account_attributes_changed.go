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

// AccountAttributesChanged struct for AccountAttributesChanged
type AccountAttributesChanged struct {
	Identity TriggerInputAccountAttributesChangedIdentity `json:"identity"`
	Source TriggerInputAccountAttributesChangedSource `json:"source"`
	Account TriggerInputAccountAttributesChangedAccount `json:"account"`
	// A list of attributes that changed.
	Changes []TriggerInputAccountAttributesChangedChangesInner `json:"changes"`
	AdditionalProperties map[string]interface{}
}

type _AccountAttributesChanged AccountAttributesChanged

// NewAccountAttributesChanged instantiates a new AccountAttributesChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAttributesChanged(identity TriggerInputAccountAttributesChangedIdentity, source TriggerInputAccountAttributesChangedSource, account TriggerInputAccountAttributesChangedAccount, changes []TriggerInputAccountAttributesChangedChangesInner) *AccountAttributesChanged {
	this := AccountAttributesChanged{}
	this.Identity = identity
	this.Source = source
	this.Account = account
	this.Changes = changes
	return &this
}

// NewAccountAttributesChangedWithDefaults instantiates a new AccountAttributesChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAttributesChangedWithDefaults() *AccountAttributesChanged {
	this := AccountAttributesChanged{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *AccountAttributesChanged) GetIdentity() TriggerInputAccountAttributesChangedIdentity {
	if o == nil {
		var ret TriggerInputAccountAttributesChangedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *AccountAttributesChanged) GetIdentityOk() (*TriggerInputAccountAttributesChangedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *AccountAttributesChanged) SetIdentity(v TriggerInputAccountAttributesChangedIdentity) {
	o.Identity = v
}

// GetSource returns the Source field value
func (o *AccountAttributesChanged) GetSource() TriggerInputAccountAttributesChangedSource {
	if o == nil {
		var ret TriggerInputAccountAttributesChangedSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AccountAttributesChanged) GetSourceOk() (*TriggerInputAccountAttributesChangedSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AccountAttributesChanged) SetSource(v TriggerInputAccountAttributesChangedSource) {
	o.Source = v
}

// GetAccount returns the Account field value
func (o *AccountAttributesChanged) GetAccount() TriggerInputAccountAttributesChangedAccount {
	if o == nil {
		var ret TriggerInputAccountAttributesChangedAccount
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AccountAttributesChanged) GetAccountOk() (*TriggerInputAccountAttributesChangedAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AccountAttributesChanged) SetAccount(v TriggerInputAccountAttributesChangedAccount) {
	o.Account = v
}

// GetChanges returns the Changes field value
func (o *AccountAttributesChanged) GetChanges() []TriggerInputAccountAttributesChangedChangesInner {
	if o == nil {
		var ret []TriggerInputAccountAttributesChangedChangesInner
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *AccountAttributesChanged) GetChangesOk() ([]TriggerInputAccountAttributesChangedChangesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *AccountAttributesChanged) SetChanges(v []TriggerInputAccountAttributesChangedChangesInner) {
	o.Changes = v
}

func (o AccountAttributesChanged) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["changes"] = o.Changes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountAttributesChanged) UnmarshalJSON(bytes []byte) (err error) {
	varAccountAttributesChanged := _AccountAttributesChanged{}

	if err = json.Unmarshal(bytes, &varAccountAttributesChanged); err == nil {
		*o = AccountAttributesChanged(varAccountAttributesChanged)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "account")
		delete(additionalProperties, "changes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAttributesChanged struct {
	value *AccountAttributesChanged
	isSet bool
}

func (v NullableAccountAttributesChanged) Get() *AccountAttributesChanged {
	return v.value
}

func (v *NullableAccountAttributesChanged) Set(val *AccountAttributesChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAttributesChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAttributesChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAttributesChanged(val *AccountAttributesChanged) *NullableAccountAttributesChanged {
	return &NullableAccountAttributesChanged{value: val, isSet: true}
}

func (v NullableAccountAttributesChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAttributesChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


