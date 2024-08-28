/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the SourceAccountUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAccountUpdated{}

// SourceAccountUpdated struct for SourceAccountUpdated
type SourceAccountUpdated struct {
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

type _SourceAccountUpdated SourceAccountUpdated

// NewSourceAccountUpdated instantiates a new SourceAccountUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccountUpdated(id string, nativeIdentifier string, sourceId string, sourceName string, identityId string, identityName string, attributes map[string]interface{}) *SourceAccountUpdated {
	this := SourceAccountUpdated{}
	this.Id = id
	this.NativeIdentifier = nativeIdentifier
	this.SourceId = sourceId
	this.SourceName = sourceName
	this.IdentityId = identityId
	this.IdentityName = identityName
	this.Attributes = attributes
	return &this
}

// NewSourceAccountUpdatedWithDefaults instantiates a new SourceAccountUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountUpdatedWithDefaults() *SourceAccountUpdated {
	this := SourceAccountUpdated{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SourceAccountUpdated) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SourceAccountUpdated) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SourceAccountUpdated) SetUuid(v string) {
	o.Uuid = &v
}

// GetId returns the Id field value
func (o *SourceAccountUpdated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceAccountUpdated) SetId(v string) {
	o.Id = v
}

// GetNativeIdentifier returns the NativeIdentifier field value
func (o *SourceAccountUpdated) GetNativeIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentifier
}

// GetNativeIdentifierOk returns a tuple with the NativeIdentifier field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetNativeIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentifier, true
}

// SetNativeIdentifier sets field value
func (o *SourceAccountUpdated) SetNativeIdentifier(v string) {
	o.NativeIdentifier = v
}

// GetSourceId returns the SourceId field value
func (o *SourceAccountUpdated) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *SourceAccountUpdated) SetSourceId(v string) {
	o.SourceId = v
}

// GetSourceName returns the SourceName field value
func (o *SourceAccountUpdated) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *SourceAccountUpdated) SetSourceName(v string) {
	o.SourceName = v
}

// GetIdentityId returns the IdentityId field value
func (o *SourceAccountUpdated) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *SourceAccountUpdated) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetIdentityName returns the IdentityName field value
func (o *SourceAccountUpdated) GetIdentityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityName
}

// GetIdentityNameOk returns a tuple with the IdentityName field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetIdentityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityName, true
}

// SetIdentityName sets field value
func (o *SourceAccountUpdated) SetIdentityName(v string) {
	o.IdentityName = v
}

// GetAttributes returns the Attributes field value
func (o *SourceAccountUpdated) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SourceAccountUpdated) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *SourceAccountUpdated) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o SourceAccountUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAccountUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
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

func (o *SourceAccountUpdated) UnmarshalJSON(data []byte) (err error) {
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

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSourceAccountUpdated := _SourceAccountUpdated{}

	err = json.Unmarshal(data, &varSourceAccountUpdated)

	if err != nil {
		return err
	}

	*o = SourceAccountUpdated(varSourceAccountUpdated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

type NullableSourceAccountUpdated struct {
	value *SourceAccountUpdated
	isSet bool
}

func (v NullableSourceAccountUpdated) Get() *SourceAccountUpdated {
	return v.value
}

func (v *NullableSourceAccountUpdated) Set(val *SourceAccountUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccountUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccountUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccountUpdated(val *SourceAccountUpdated) *NullableSourceAccountUpdated {
	return &NullableSourceAccountUpdated{value: val, isSet: true}
}

func (v NullableSourceAccountUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccountUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


