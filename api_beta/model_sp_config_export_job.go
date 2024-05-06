/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SpConfigExportJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpConfigExportJob{}

// SpConfigExportJob struct for SpConfigExportJob
type SpConfigExportJob struct {
	// Unique id assigned to this job.
	JobId string `json:"jobId"`
	// Status of the job.
	Status string `json:"status"`
	// Type of the job, either export or import.
	Type string `json:"type"`
	// The time until which the artifacts will be available for download.
	Expiration time.Time `json:"expiration"`
	// The time the job was started.
	Created time.Time `json:"created"`
	// The time of the last update to the job.
	Modified time.Time `json:"modified"`
	// Optional user defined description/name for export job.
	Description string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _SpConfigExportJob SpConfigExportJob

// NewSpConfigExportJob instantiates a new SpConfigExportJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpConfigExportJob(jobId string, status string, type_ string, expiration time.Time, created time.Time, modified time.Time, description string) *SpConfigExportJob {
	this := SpConfigExportJob{}
	this.JobId = jobId
	this.Status = status
	this.Type = type_
	this.Expiration = expiration
	this.Created = created
	this.Modified = modified
	this.Description = description
	return &this
}

// NewSpConfigExportJobWithDefaults instantiates a new SpConfigExportJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpConfigExportJobWithDefaults() *SpConfigExportJob {
	this := SpConfigExportJob{}
	return &this
}

// GetJobId returns the JobId field value
func (o *SpConfigExportJob) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *SpConfigExportJob) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *SpConfigExportJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SpConfigExportJob) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *SpConfigExportJob) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpConfigExportJob) SetType(v string) {
	o.Type = v
}

// GetExpiration returns the Expiration field value
func (o *SpConfigExportJob) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *SpConfigExportJob) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetCreated returns the Created field value
func (o *SpConfigExportJob) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SpConfigExportJob) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *SpConfigExportJob) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *SpConfigExportJob) SetModified(v time.Time) {
	o.Modified = v
}

// GetDescription returns the Description field value
func (o *SpConfigExportJob) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SpConfigExportJob) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SpConfigExportJob) SetDescription(v string) {
	o.Description = v
}

func (o SpConfigExportJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpConfigExportJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["expiration"] = o.Expiration
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpConfigExportJob) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobId",
		"status",
		"type",
		"expiration",
		"created",
		"modified",
		"description",
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

	varSpConfigExportJob := _SpConfigExportJob{}

	if err = json.Unmarshal(bytes, &varSpConfigExportJob); err == nil {
	*o = SpConfigExportJob(varSpConfigExportJob)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpConfigExportJob struct {
	value *SpConfigExportJob
	isSet bool
}

func (v NullableSpConfigExportJob) Get() *SpConfigExportJob {
	return v.value
}

func (v *NullableSpConfigExportJob) Set(val *SpConfigExportJob) {
	v.value = val
	v.isSet = true
}

func (v NullableSpConfigExportJob) IsSet() bool {
	return v.isSet
}

func (v *NullableSpConfigExportJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpConfigExportJob(val *SpConfigExportJob) *NullableSpConfigExportJob {
	return &NullableSpConfigExportJob{value: val, isSet: true}
}

func (v NullableSpConfigExportJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpConfigExportJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


