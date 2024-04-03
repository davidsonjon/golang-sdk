/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the SourceAccountDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAccountDeleted{}

// SourceAccountDeleted struct for SourceAccountDeleted
type SourceAccountDeleted struct {
	// Source unique identifier for the identity. UUID is generated by the source system.
	Uuid *string `json:"uuid,omitempty"`
	// SailPoint generated unique identifier.
	Id string `json:"id"`
	// Unique ID of the account on the source.
	NativeIdentifier string `json:"nativeIdentifier"`
	// The ID of the source.
	SourceId string `json:"sourceId"`
	// The name of the source.
	SourceName string `json:"sourceName"`
	// The ID of the identity that is correlated with this account.
	IdentityId string `json:"identityId"`
	// The name of the identity that is correlated with this account.
	IdentityName string `json:"identityName"`
	// The attributes of the account. The contents of attributes depends on the account schema for the source.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _SourceAccountDeleted SourceAccountDeleted

// NewSourceAccountDeleted instantiates a new SourceAccountDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccountDeleted(id string, nativeIdentifier string, sourceId string, sourceName string, identityId string, identityName string, attributes map[string]interface{}) *SourceAccountDeleted {
	this := SourceAccountDeleted{}
	this.Id = id
	this.NativeIdentifier = nativeIdentifier
	this.SourceId = sourceId
	this.SourceName = sourceName
	this.IdentityId = identityId
	this.IdentityName = identityName
	this.Attributes = attributes
	return &this
}

// NewSourceAccountDeletedWithDefaults instantiates a new SourceAccountDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountDeletedWithDefaults() *SourceAccountDeleted {
	this := SourceAccountDeleted{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SourceAccountDeleted) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SourceAccountDeleted) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SourceAccountDeleted) SetUuid(v string) {
	o.Uuid = &v
}

// GetId returns the Id field value
func (o *SourceAccountDeleted) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceAccountDeleted) SetId(v string) {
	o.Id = v
}

// GetNativeIdentifier returns the NativeIdentifier field value
func (o *SourceAccountDeleted) GetNativeIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentifier
}

// GetNativeIdentifierOk returns a tuple with the NativeIdentifier field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetNativeIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentifier, true
}

// SetNativeIdentifier sets field value
func (o *SourceAccountDeleted) SetNativeIdentifier(v string) {
	o.NativeIdentifier = v
}

// GetSourceId returns the SourceId field value
func (o *SourceAccountDeleted) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *SourceAccountDeleted) SetSourceId(v string) {
	o.SourceId = v
}

// GetSourceName returns the SourceName field value
func (o *SourceAccountDeleted) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *SourceAccountDeleted) SetSourceName(v string) {
	o.SourceName = v
}

// GetIdentityId returns the IdentityId field value
func (o *SourceAccountDeleted) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *SourceAccountDeleted) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetIdentityName returns the IdentityName field value
func (o *SourceAccountDeleted) GetIdentityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityName
}

// GetIdentityNameOk returns a tuple with the IdentityName field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetIdentityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityName, true
}

// SetIdentityName sets field value
func (o *SourceAccountDeleted) SetIdentityName(v string) {
	o.IdentityName = v
}

// GetAttributes returns the Attributes field value
func (o *SourceAccountDeleted) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SourceAccountDeleted) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *SourceAccountDeleted) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o SourceAccountDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAccountDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["id"] = o.Id
	toSerialize["nativeIdentifier"] = o.NativeIdentifier
	toSerialize["sourceId"] = o.SourceId
	toSerialize["sourceName"] = o.SourceName
	toSerialize["identityId"] = o.IdentityId
	toSerialize["identityName"] = o.IdentityName
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceAccountDeleted) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"nativeIdentifier",
		"sourceId",
		"sourceName",
		"identityId",
		"identityName",
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSourceAccountDeleted := _SourceAccountDeleted{}

	if err = json.Unmarshal(bytes, &varSourceAccountDeleted); err == nil {
	*o = SourceAccountDeleted(varSourceAccountDeleted)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "nativeIdentifier")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "identityName")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceAccountDeleted struct {
	value *SourceAccountDeleted
	isSet bool
}

func (v NullableSourceAccountDeleted) Get() *SourceAccountDeleted {
	return v.value
}

func (v *NullableSourceAccountDeleted) Set(val *SourceAccountDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccountDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccountDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccountDeleted(val *SourceAccountDeleted) *NullableSourceAccountDeleted {
	return &NullableSourceAccountDeleted{value: val, isSet: true}
}

func (v NullableSourceAccountDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccountDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


