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

// checks if the QueuedCheckConfigDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueuedCheckConfigDetails{}

// QueuedCheckConfigDetails Configuration of maximum number days and interval for checking Service Desk integration queue status
type QueuedCheckConfigDetails struct {
	// interval in minutes between status checks
	ProvisioningStatusCheckIntervalMinutes string `json:"provisioningStatusCheckIntervalMinutes"`
	// maximum number of days to check
	ProvisioningMaxStatusCheckDays string `json:"provisioningMaxStatusCheckDays"`
	AdditionalProperties map[string]interface{}
}

type _QueuedCheckConfigDetails QueuedCheckConfigDetails

// NewQueuedCheckConfigDetails instantiates a new QueuedCheckConfigDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueuedCheckConfigDetails(provisioningStatusCheckIntervalMinutes string, provisioningMaxStatusCheckDays string) *QueuedCheckConfigDetails {
	this := QueuedCheckConfigDetails{}
	this.ProvisioningStatusCheckIntervalMinutes = provisioningStatusCheckIntervalMinutes
	this.ProvisioningMaxStatusCheckDays = provisioningMaxStatusCheckDays
	return &this
}

// NewQueuedCheckConfigDetailsWithDefaults instantiates a new QueuedCheckConfigDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueuedCheckConfigDetailsWithDefaults() *QueuedCheckConfigDetails {
	this := QueuedCheckConfigDetails{}
	return &this
}

// GetProvisioningStatusCheckIntervalMinutes returns the ProvisioningStatusCheckIntervalMinutes field value
func (o *QueuedCheckConfigDetails) GetProvisioningStatusCheckIntervalMinutes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningStatusCheckIntervalMinutes
}

// GetProvisioningStatusCheckIntervalMinutesOk returns a tuple with the ProvisioningStatusCheckIntervalMinutes field value
// and a boolean to check if the value has been set.
func (o *QueuedCheckConfigDetails) GetProvisioningStatusCheckIntervalMinutesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningStatusCheckIntervalMinutes, true
}

// SetProvisioningStatusCheckIntervalMinutes sets field value
func (o *QueuedCheckConfigDetails) SetProvisioningStatusCheckIntervalMinutes(v string) {
	o.ProvisioningStatusCheckIntervalMinutes = v
}

// GetProvisioningMaxStatusCheckDays returns the ProvisioningMaxStatusCheckDays field value
func (o *QueuedCheckConfigDetails) GetProvisioningMaxStatusCheckDays() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningMaxStatusCheckDays
}

// GetProvisioningMaxStatusCheckDaysOk returns a tuple with the ProvisioningMaxStatusCheckDays field value
// and a boolean to check if the value has been set.
func (o *QueuedCheckConfigDetails) GetProvisioningMaxStatusCheckDaysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningMaxStatusCheckDays, true
}

// SetProvisioningMaxStatusCheckDays sets field value
func (o *QueuedCheckConfigDetails) SetProvisioningMaxStatusCheckDays(v string) {
	o.ProvisioningMaxStatusCheckDays = v
}

func (o QueuedCheckConfigDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueuedCheckConfigDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provisioningStatusCheckIntervalMinutes"] = o.ProvisioningStatusCheckIntervalMinutes
	toSerialize["provisioningMaxStatusCheckDays"] = o.ProvisioningMaxStatusCheckDays

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueuedCheckConfigDetails) UnmarshalJSON(bytes []byte) (err error) {
	varQueuedCheckConfigDetails := _QueuedCheckConfigDetails{}

	if err = json.Unmarshal(bytes, &varQueuedCheckConfigDetails); err == nil {
		*o = QueuedCheckConfigDetails(varQueuedCheckConfigDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "provisioningStatusCheckIntervalMinutes")
		delete(additionalProperties, "provisioningMaxStatusCheckDays")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueuedCheckConfigDetails struct {
	value *QueuedCheckConfigDetails
	isSet bool
}

func (v NullableQueuedCheckConfigDetails) Get() *QueuedCheckConfigDetails {
	return v.value
}

func (v *NullableQueuedCheckConfigDetails) Set(val *QueuedCheckConfigDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableQueuedCheckConfigDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableQueuedCheckConfigDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueuedCheckConfigDetails(val *QueuedCheckConfigDetails) *NullableQueuedCheckConfigDetails {
	return &NullableQueuedCheckConfigDetails{value: val, isSet: true}
}

func (v NullableQueuedCheckConfigDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueuedCheckConfigDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


