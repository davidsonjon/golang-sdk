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

// checks if the AccountStatusChangedStatusChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountStatusChangedStatusChange{}

// AccountStatusChangedStatusChange struct for AccountStatusChangedStatusChange
type AccountStatusChangedStatusChange struct {
	// the previous status of the account
	PreviousStatus *string `json:"previousStatus,omitempty"`
	// the new status of the account
	NewStatus *string `json:"newStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountStatusChangedStatusChange AccountStatusChangedStatusChange

// NewAccountStatusChangedStatusChange instantiates a new AccountStatusChangedStatusChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatusChangedStatusChange() *AccountStatusChangedStatusChange {
	this := AccountStatusChangedStatusChange{}
	return &this
}

// NewAccountStatusChangedStatusChangeWithDefaults instantiates a new AccountStatusChangedStatusChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusChangedStatusChangeWithDefaults() *AccountStatusChangedStatusChange {
	this := AccountStatusChangedStatusChange{}
	return &this
}

// GetPreviousStatus returns the PreviousStatus field value if set, zero value otherwise.
func (o *AccountStatusChangedStatusChange) GetPreviousStatus() string {
	if o == nil || isNil(o.PreviousStatus) {
		var ret string
		return ret
	}
	return *o.PreviousStatus
}

// GetPreviousStatusOk returns a tuple with the PreviousStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChangedStatusChange) GetPreviousStatusOk() (*string, bool) {
	if o == nil || isNil(o.PreviousStatus) {
		return nil, false
	}
	return o.PreviousStatus, true
}

// HasPreviousStatus returns a boolean if a field has been set.
func (o *AccountStatusChangedStatusChange) HasPreviousStatus() bool {
	if o != nil && !isNil(o.PreviousStatus) {
		return true
	}

	return false
}

// SetPreviousStatus gets a reference to the given string and assigns it to the PreviousStatus field.
func (o *AccountStatusChangedStatusChange) SetPreviousStatus(v string) {
	o.PreviousStatus = &v
}

// GetNewStatus returns the NewStatus field value if set, zero value otherwise.
func (o *AccountStatusChangedStatusChange) GetNewStatus() string {
	if o == nil || isNil(o.NewStatus) {
		var ret string
		return ret
	}
	return *o.NewStatus
}

// GetNewStatusOk returns a tuple with the NewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusChangedStatusChange) GetNewStatusOk() (*string, bool) {
	if o == nil || isNil(o.NewStatus) {
		return nil, false
	}
	return o.NewStatus, true
}

// HasNewStatus returns a boolean if a field has been set.
func (o *AccountStatusChangedStatusChange) HasNewStatus() bool {
	if o != nil && !isNil(o.NewStatus) {
		return true
	}

	return false
}

// SetNewStatus gets a reference to the given string and assigns it to the NewStatus field.
func (o *AccountStatusChangedStatusChange) SetNewStatus(v string) {
	o.NewStatus = &v
}

func (o AccountStatusChangedStatusChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountStatusChangedStatusChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PreviousStatus) {
		toSerialize["previousStatus"] = o.PreviousStatus
	}
	if !isNil(o.NewStatus) {
		toSerialize["newStatus"] = o.NewStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountStatusChangedStatusChange) UnmarshalJSON(bytes []byte) (err error) {
	varAccountStatusChangedStatusChange := _AccountStatusChangedStatusChange{}

	if err = json.Unmarshal(bytes, &varAccountStatusChangedStatusChange); err == nil {
	*o = AccountStatusChangedStatusChange(varAccountStatusChangedStatusChange)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "previousStatus")
		delete(additionalProperties, "newStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountStatusChangedStatusChange struct {
	value *AccountStatusChangedStatusChange
	isSet bool
}

func (v NullableAccountStatusChangedStatusChange) Get() *AccountStatusChangedStatusChange {
	return v.value
}

func (v *NullableAccountStatusChangedStatusChange) Set(val *AccountStatusChangedStatusChange) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusChangedStatusChange) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusChangedStatusChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusChangedStatusChange(val *AccountStatusChangedStatusChange) *NullableAccountStatusChangedStatusChange {
	return &NullableAccountStatusChangedStatusChange{value: val, isSet: true}
}

func (v NullableAccountStatusChangedStatusChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusChangedStatusChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


