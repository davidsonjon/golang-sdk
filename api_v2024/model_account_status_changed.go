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

// checks if the AccountStatusChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountStatusChanged{}

// AccountStatusChanged struct for AccountStatusChanged
type AccountStatusChanged struct {
	// the event type
	EventType *string `json:"eventType,omitempty"`
	// the identity id
	IdentityId *string `json:"identityId,omitempty"`
	// the date of event
	Dt *string `json:"dt,omitempty"`
	Account *AccountStatusChangedAccount `json:"account,omitempty"`
	StatusChange *AccountStatusChangedStatusChange `json:"statusChange,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountStatusChanged AccountStatusChanged

// NewAccountStatusChanged instantiates a new AccountStatusChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatusChanged() *AccountStatusChanged {
	this := AccountStatusChanged{}
	return &this
}

// NewAccountStatusChangedWithDefaults instantiates a new AccountStatusChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusChangedWithDefaults() *AccountStatusChanged {
	this := AccountStatusChanged{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *AccountStatusChanged) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChanged) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *AccountStatusChanged) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *AccountStatusChanged) SetEventType(v string) {
	o.EventType = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *AccountStatusChanged) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChanged) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *AccountStatusChanged) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *AccountStatusChanged) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *AccountStatusChanged) GetDt() string {
	if o == nil || IsNil(o.Dt) {
		var ret string
		return ret
	}
	return *o.Dt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChanged) GetDtOk() (*string, bool) {
	if o == nil || IsNil(o.Dt) {
		return nil, false
	}
	return o.Dt, true
}

// HasDt returns a boolean if a field has been set.
func (o *AccountStatusChanged) HasDt() bool {
	if o != nil && !IsNil(o.Dt) {
		return true
	}

	return false
}

// SetDt gets a reference to the given string and assigns it to the Dt field.
func (o *AccountStatusChanged) SetDt(v string) {
	o.Dt = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AccountStatusChanged) GetAccount() AccountStatusChangedAccount {
	if o == nil || IsNil(o.Account) {
		var ret AccountStatusChangedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChanged) GetAccountOk() (*AccountStatusChangedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AccountStatusChanged) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountStatusChangedAccount and assigns it to the Account field.
func (o *AccountStatusChanged) SetAccount(v AccountStatusChangedAccount) {
	o.Account = &v
}

// GetStatusChange returns the StatusChange field value if set, zero value otherwise.
func (o *AccountStatusChanged) GetStatusChange() AccountStatusChangedStatusChange {
	if o == nil || IsNil(o.StatusChange) {
		var ret AccountStatusChangedStatusChange
		return ret
	}
	return *o.StatusChange
}

// GetStatusChangeOk returns a tuple with the StatusChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChanged) GetStatusChangeOk() (*AccountStatusChangedStatusChange, bool) {
	if o == nil || IsNil(o.StatusChange) {
		return nil, false
	}
	return o.StatusChange, true
}

// HasStatusChange returns a boolean if a field has been set.
func (o *AccountStatusChanged) HasStatusChange() bool {
	if o != nil && !IsNil(o.StatusChange) {
		return true
	}

	return false
}

// SetStatusChange gets a reference to the given AccountStatusChangedStatusChange and assigns it to the StatusChange field.
func (o *AccountStatusChanged) SetStatusChange(v AccountStatusChangedStatusChange) {
	o.StatusChange = &v
}

func (o AccountStatusChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountStatusChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !IsNil(o.Dt) {
		toSerialize["dt"] = o.Dt
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.StatusChange) {
		toSerialize["statusChange"] = o.StatusChange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountStatusChanged) UnmarshalJSON(data []byte) (err error) {
	varAccountStatusChanged := _AccountStatusChanged{}

	err = json.Unmarshal(data, &varAccountStatusChanged)

	if err != nil {
		return err
	}

	*o = AccountStatusChanged(varAccountStatusChanged)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "dt")
		delete(additionalProperties, "account")
		delete(additionalProperties, "statusChange")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountStatusChanged struct {
	value *AccountStatusChanged
	isSet bool
}

func (v NullableAccountStatusChanged) Get() *AccountStatusChanged {
	return v.value
}

func (v *NullableAccountStatusChanged) Set(val *AccountStatusChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusChanged(val *AccountStatusChanged) *NullableAccountStatusChanged {
	return &NullableAccountStatusChanged{value: val, isSet: true}
}

func (v NullableAccountStatusChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


