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

// TriggerInputAccountCorrelatedAccount The account that was correlated.
type TriggerInputAccountCorrelatedAccount struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type"`
	// Unique ID of the account on the source.
	NativeIdentity string `json:"nativeIdentity"`
	// The source's unique identifier for the account. UUID is generated by the source system.
	Uuid NullableString `json:"uuid,omitempty"`
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountCorrelatedAccount TriggerInputAccountCorrelatedAccount

// NewTriggerInputAccountCorrelatedAccount instantiates a new TriggerInputAccountCorrelatedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountCorrelatedAccount(type_ map[string]interface{}, nativeIdentity string, id string, name string) *TriggerInputAccountCorrelatedAccount {
	this := TriggerInputAccountCorrelatedAccount{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewTriggerInputAccountCorrelatedAccountWithDefaults instantiates a new TriggerInputAccountCorrelatedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountCorrelatedAccountWithDefaults() *TriggerInputAccountCorrelatedAccount {
	this := TriggerInputAccountCorrelatedAccount{}
	return &this
}

// GetType returns the Type field value
func (o *TriggerInputAccountCorrelatedAccount) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedAccount) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccountCorrelatedAccount) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetNativeIdentity returns the NativeIdentity field value
func (o *TriggerInputAccountCorrelatedAccount) GetNativeIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentity, true
}

// SetNativeIdentity sets field value
func (o *TriggerInputAccountCorrelatedAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputAccountCorrelatedAccount) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccountCorrelatedAccount) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *TriggerInputAccountCorrelatedAccount) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *TriggerInputAccountCorrelatedAccount) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *TriggerInputAccountCorrelatedAccount) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *TriggerInputAccountCorrelatedAccount) UnsetUuid() {
	o.Uuid.Unset()
}

// GetId returns the Id field value
func (o *TriggerInputAccountCorrelatedAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccountCorrelatedAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputAccountCorrelatedAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccountCorrelatedAccount) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccountCorrelatedAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccountCorrelatedAccount) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountCorrelatedAccount := _TriggerInputAccountCorrelatedAccount{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountCorrelatedAccount); err == nil {
		*o = TriggerInputAccountCorrelatedAccount(varTriggerInputAccountCorrelatedAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountCorrelatedAccount struct {
	value *TriggerInputAccountCorrelatedAccount
	isSet bool
}

func (v NullableTriggerInputAccountCorrelatedAccount) Get() *TriggerInputAccountCorrelatedAccount {
	return v.value
}

func (v *NullableTriggerInputAccountCorrelatedAccount) Set(val *TriggerInputAccountCorrelatedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountCorrelatedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountCorrelatedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountCorrelatedAccount(val *TriggerInputAccountCorrelatedAccount) *NullableTriggerInputAccountCorrelatedAccount {
	return &NullableTriggerInputAccountCorrelatedAccount{value: val, isSet: true}
}

func (v NullableTriggerInputAccountCorrelatedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountCorrelatedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


