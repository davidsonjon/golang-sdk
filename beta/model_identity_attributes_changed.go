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

// checks if the IdentityAttributesChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAttributesChanged{}

// IdentityAttributesChanged struct for IdentityAttributesChanged
type IdentityAttributesChanged struct {
	Identity IdentityAttributesChangedIdentity `json:"identity"`
	// A list of one or more identity attributes that changed on the identity.
	Changes []IdentityAttributesChangedChangesInner `json:"changes"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttributesChanged IdentityAttributesChanged

// NewIdentityAttributesChanged instantiates a new IdentityAttributesChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttributesChanged(identity IdentityAttributesChangedIdentity, changes []IdentityAttributesChangedChangesInner) *IdentityAttributesChanged {
	this := IdentityAttributesChanged{}
	this.Identity = identity
	this.Changes = changes
	return &this
}

// NewIdentityAttributesChangedWithDefaults instantiates a new IdentityAttributesChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttributesChangedWithDefaults() *IdentityAttributesChanged {
	this := IdentityAttributesChanged{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *IdentityAttributesChanged) GetIdentity() IdentityAttributesChangedIdentity {
	if o == nil {
		var ret IdentityAttributesChangedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *IdentityAttributesChanged) GetIdentityOk() (*IdentityAttributesChangedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *IdentityAttributesChanged) SetIdentity(v IdentityAttributesChangedIdentity) {
	o.Identity = v
}

// GetChanges returns the Changes field value
func (o *IdentityAttributesChanged) GetChanges() []IdentityAttributesChangedChangesInner {
	if o == nil {
		var ret []IdentityAttributesChangedChangesInner
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *IdentityAttributesChanged) GetChangesOk() ([]IdentityAttributesChangedChangesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *IdentityAttributesChanged) SetChanges(v []IdentityAttributesChangedChangesInner) {
	o.Changes = v
}

func (o IdentityAttributesChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAttributesChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identity"] = o.Identity
	toSerialize["changes"] = o.Changes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityAttributesChanged) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityAttributesChanged := _IdentityAttributesChanged{}

	if err = json.Unmarshal(bytes, &varIdentityAttributesChanged); err == nil {
		*o = IdentityAttributesChanged(varIdentityAttributesChanged)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "changes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAttributesChanged struct {
	value *IdentityAttributesChanged
	isSet bool
}

func (v NullableIdentityAttributesChanged) Get() *IdentityAttributesChanged {
	return v.value
}

func (v *NullableIdentityAttributesChanged) Set(val *IdentityAttributesChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributesChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributesChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributesChanged(val *IdentityAttributesChanged) *NullableIdentityAttributesChanged {
	return &NullableIdentityAttributesChanged{value: val, isSet: true}
}

func (v NullableIdentityAttributesChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributesChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


