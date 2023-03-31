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

// checks if the TriggerInputIdentityDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputIdentityDeleted{}

// TriggerInputIdentityDeleted struct for TriggerInputIdentityDeleted
type TriggerInputIdentityDeleted struct {
	Identity TriggerInputIdentityDeletedIdentity `json:"identity"`
	// The attributes assigned to the identity.  Attributes are determined by the identity profile.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputIdentityDeleted TriggerInputIdentityDeleted

// NewTriggerInputIdentityDeleted instantiates a new TriggerInputIdentityDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputIdentityDeleted(identity TriggerInputIdentityDeletedIdentity, attributes map[string]interface{}) *TriggerInputIdentityDeleted {
	this := TriggerInputIdentityDeleted{}
	this.Identity = identity
	this.Attributes = attributes
	return &this
}

// NewTriggerInputIdentityDeletedWithDefaults instantiates a new TriggerInputIdentityDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputIdentityDeletedWithDefaults() *TriggerInputIdentityDeleted {
	this := TriggerInputIdentityDeleted{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *TriggerInputIdentityDeleted) GetIdentity() TriggerInputIdentityDeletedIdentity {
	if o == nil {
		var ret TriggerInputIdentityDeletedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputIdentityDeleted) GetIdentityOk() (*TriggerInputIdentityDeletedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *TriggerInputIdentityDeleted) SetIdentity(v TriggerInputIdentityDeletedIdentity) {
	o.Identity = v
}

// GetAttributes returns the Attributes field value
func (o *TriggerInputIdentityDeleted) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TriggerInputIdentityDeleted) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TriggerInputIdentityDeleted) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TriggerInputIdentityDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputIdentityDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identity"] = o.Identity
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputIdentityDeleted) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputIdentityDeleted := _TriggerInputIdentityDeleted{}

	if err = json.Unmarshal(bytes, &varTriggerInputIdentityDeleted); err == nil {
		*o = TriggerInputIdentityDeleted(varTriggerInputIdentityDeleted)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputIdentityDeleted struct {
	value *TriggerInputIdentityDeleted
	isSet bool
}

func (v NullableTriggerInputIdentityDeleted) Get() *TriggerInputIdentityDeleted {
	return v.value
}

func (v *NullableTriggerInputIdentityDeleted) Set(val *TriggerInputIdentityDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputIdentityDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputIdentityDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputIdentityDeleted(val *TriggerInputIdentityDeleted) *NullableTriggerInputIdentityDeleted {
	return &NullableTriggerInputIdentityDeleted{value: val, isSet: true}
}

func (v NullableTriggerInputIdentityDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputIdentityDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


