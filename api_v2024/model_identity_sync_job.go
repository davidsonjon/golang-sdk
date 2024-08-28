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

// checks if the IdentitySyncJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySyncJob{}

// IdentitySyncJob struct for IdentitySyncJob
type IdentitySyncJob struct {
	// Job ID.
	Id string `json:"id"`
	// The job status.
	Status string `json:"status"`
	Payload IdentitySyncPayload `json:"payload"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySyncJob IdentitySyncJob

// NewIdentitySyncJob instantiates a new IdentitySyncJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySyncJob(id string, status string, payload IdentitySyncPayload) *IdentitySyncJob {
	this := IdentitySyncJob{}
	this.Id = id
	this.Status = status
	this.Payload = payload
	return &this
}

// NewIdentitySyncJobWithDefaults instantiates a new IdentitySyncJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySyncJobWithDefaults() *IdentitySyncJob {
	this := IdentitySyncJob{}
	return &this
}

// GetId returns the Id field value
func (o *IdentitySyncJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentitySyncJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentitySyncJob) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *IdentitySyncJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IdentitySyncJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IdentitySyncJob) SetStatus(v string) {
	o.Status = v
}

// GetPayload returns the Payload field value
func (o *IdentitySyncJob) GetPayload() IdentitySyncPayload {
	if o == nil {
		var ret IdentitySyncPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *IdentitySyncJob) GetPayloadOk() (*IdentitySyncPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *IdentitySyncJob) SetPayload(v IdentitySyncPayload) {
	o.Payload = v
}

func (o IdentitySyncJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySyncJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["payload"] = o.Payload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySyncJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"payload",
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

	varIdentitySyncJob := _IdentitySyncJob{}

	err = json.Unmarshal(data, &varIdentitySyncJob)

	if err != nil {
		return err
	}

	*o = IdentitySyncJob(varIdentitySyncJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySyncJob struct {
	value *IdentitySyncJob
	isSet bool
}

func (v NullableIdentitySyncJob) Get() *IdentitySyncJob {
	return v.value
}

func (v *NullableIdentitySyncJob) Set(val *IdentitySyncJob) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySyncJob) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySyncJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySyncJob(val *IdentitySyncJob) *NullableIdentitySyncJob {
	return &NullableIdentitySyncJob{value: val, isSet: true}
}

func (v NullableIdentitySyncJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySyncJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


