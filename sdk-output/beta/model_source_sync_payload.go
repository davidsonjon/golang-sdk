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

// SourceSyncPayload struct for SourceSyncPayload
type SourceSyncPayload struct {
	// Payload type.
	Type string `json:"type"`
	// Payload type.
	DataJson string `json:"dataJson"`
	AdditionalProperties map[string]interface{}
}

type _SourceSyncPayload SourceSyncPayload

// NewSourceSyncPayload instantiates a new SourceSyncPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSyncPayload(type_ string, dataJson string) *SourceSyncPayload {
	this := SourceSyncPayload{}
	this.Type = type_
	this.DataJson = dataJson
	return &this
}

// NewSourceSyncPayloadWithDefaults instantiates a new SourceSyncPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSyncPayloadWithDefaults() *SourceSyncPayload {
	this := SourceSyncPayload{}
	return &this
}

// GetType returns the Type field value
func (o *SourceSyncPayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceSyncPayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceSyncPayload) SetType(v string) {
	o.Type = v
}

// GetDataJson returns the DataJson field value
func (o *SourceSyncPayload) GetDataJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataJson
}

// GetDataJsonOk returns a tuple with the DataJson field value
// and a boolean to check if the value has been set.
func (o *SourceSyncPayload) GetDataJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataJson, true
}

// SetDataJson sets field value
func (o *SourceSyncPayload) SetDataJson(v string) {
	o.DataJson = v
}

func (o SourceSyncPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["dataJson"] = o.DataJson
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SourceSyncPayload) UnmarshalJSON(bytes []byte) (err error) {
	varSourceSyncPayload := _SourceSyncPayload{}

	if err = json.Unmarshal(bytes, &varSourceSyncPayload); err == nil {
		*o = SourceSyncPayload(varSourceSyncPayload)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "dataJson")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceSyncPayload struct {
	value *SourceSyncPayload
	isSet bool
}

func (v NullableSourceSyncPayload) Get() *SourceSyncPayload {
	return v.value
}

func (v *NullableSourceSyncPayload) Set(val *SourceSyncPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSyncPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSyncPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSyncPayload(val *SourceSyncPayload) *NullableSourceSyncPayload {
	return &NullableSourceSyncPayload{value: val, isSet: true}
}

func (v NullableSourceSyncPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSyncPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


