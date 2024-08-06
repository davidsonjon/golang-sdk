/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the ConfigurationItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationItemRequest{}

// ConfigurationItemRequest The request body for creation or update of a Reassignment Configuration for a single identity and work type
type ConfigurationItemRequest struct {
	// The identity id to reassign an item from
	ReassignedFromId *string `json:"reassignedFromId,omitempty"`
	// The identity id to reassign an item to
	ReassignedToId *string `json:"reassignedToId,omitempty"`
	ConfigType *ConfigTypeEnum `json:"configType,omitempty"`
	// The date from which to start reassigning work items
	StartDate *time.Time `json:"startDate,omitempty"`
	// The date from which to stop reassigning work items.  If this is an null string it indicates a permanent reassignment.
	EndDate NullableTime `json:"endDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationItemRequest ConfigurationItemRequest

// NewConfigurationItemRequest instantiates a new ConfigurationItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationItemRequest() *ConfigurationItemRequest {
	this := ConfigurationItemRequest{}
	return &this
}

// NewConfigurationItemRequestWithDefaults instantiates a new ConfigurationItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationItemRequestWithDefaults() *ConfigurationItemRequest {
	this := ConfigurationItemRequest{}
	return &this
}

// GetReassignedFromId returns the ReassignedFromId field value if set, zero value otherwise.
func (o *ConfigurationItemRequest) GetReassignedFromId() string {
	if o == nil || isNil(o.ReassignedFromId) {
		var ret string
		return ret
	}
	return *o.ReassignedFromId
}

// GetReassignedFromIdOk returns a tuple with the ReassignedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationItemRequest) GetReassignedFromIdOk() (*string, bool) {
	if o == nil || isNil(o.ReassignedFromId) {
		return nil, false
	}
	return o.ReassignedFromId, true
}

// HasReassignedFromId returns a boolean if a field has been set.
func (o *ConfigurationItemRequest) HasReassignedFromId() bool {
	if o != nil && !isNil(o.ReassignedFromId) {
		return true
	}

	return false
}

// SetReassignedFromId gets a reference to the given string and assigns it to the ReassignedFromId field.
func (o *ConfigurationItemRequest) SetReassignedFromId(v string) {
	o.ReassignedFromId = &v
}

// GetReassignedToId returns the ReassignedToId field value if set, zero value otherwise.
func (o *ConfigurationItemRequest) GetReassignedToId() string {
	if o == nil || isNil(o.ReassignedToId) {
		var ret string
		return ret
	}
	return *o.ReassignedToId
}

// GetReassignedToIdOk returns a tuple with the ReassignedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationItemRequest) GetReassignedToIdOk() (*string, bool) {
	if o == nil || isNil(o.ReassignedToId) {
		return nil, false
	}
	return o.ReassignedToId, true
}

// HasReassignedToId returns a boolean if a field has been set.
func (o *ConfigurationItemRequest) HasReassignedToId() bool {
	if o != nil && !isNil(o.ReassignedToId) {
		return true
	}

	return false
}

// SetReassignedToId gets a reference to the given string and assigns it to the ReassignedToId field.
func (o *ConfigurationItemRequest) SetReassignedToId(v string) {
	o.ReassignedToId = &v
}

// GetConfigType returns the ConfigType field value if set, zero value otherwise.
func (o *ConfigurationItemRequest) GetConfigType() ConfigTypeEnum {
	if o == nil || isNil(o.ConfigType) {
		var ret ConfigTypeEnum
		return ret
	}
	return *o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationItemRequest) GetConfigTypeOk() (*ConfigTypeEnum, bool) {
	if o == nil || isNil(o.ConfigType) {
		return nil, false
	}
	return o.ConfigType, true
}

// HasConfigType returns a boolean if a field has been set.
func (o *ConfigurationItemRequest) HasConfigType() bool {
	if o != nil && !isNil(o.ConfigType) {
		return true
	}

	return false
}

// SetConfigType gets a reference to the given ConfigTypeEnum and assigns it to the ConfigType field.
func (o *ConfigurationItemRequest) SetConfigType(v ConfigTypeEnum) {
	o.ConfigType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ConfigurationItemRequest) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationItemRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ConfigurationItemRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ConfigurationItemRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationItemRequest) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationItemRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ConfigurationItemRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *ConfigurationItemRequest) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ConfigurationItemRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ConfigurationItemRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

func (o ConfigurationItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReassignedFromId) {
		toSerialize["reassignedFromId"] = o.ReassignedFromId
	}
	if !isNil(o.ReassignedToId) {
		toSerialize["reassignedToId"] = o.ReassignedToId
	}
	if !isNil(o.ConfigType) {
		toSerialize["configType"] = o.ConfigType
	}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationItemRequest) UnmarshalJSON(bytes []byte) (err error) {
	varConfigurationItemRequest := _ConfigurationItemRequest{}

	if err = json.Unmarshal(bytes, &varConfigurationItemRequest); err == nil {
	*o = ConfigurationItemRequest(varConfigurationItemRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reassignedFromId")
		delete(additionalProperties, "reassignedToId")
		delete(additionalProperties, "configType")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationItemRequest struct {
	value *ConfigurationItemRequest
	isSet bool
}

func (v NullableConfigurationItemRequest) Get() *ConfigurationItemRequest {
	return v.value
}

func (v *NullableConfigurationItemRequest) Set(val *ConfigurationItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationItemRequest(val *ConfigurationItemRequest) *NullableConfigurationItemRequest {
	return &NullableConfigurationItemRequest{value: val, isSet: true}
}

func (v NullableConfigurationItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


