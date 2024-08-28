/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the TenantConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantConfigurationRequest{}

// TenantConfigurationRequest Tenant-wide Reassignment Configuration settings
type TenantConfigurationRequest struct {
	ConfigDetails *TenantConfigurationDetails `json:"configDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TenantConfigurationRequest TenantConfigurationRequest

// NewTenantConfigurationRequest instantiates a new TenantConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantConfigurationRequest() *TenantConfigurationRequest {
	this := TenantConfigurationRequest{}
	return &this
}

// NewTenantConfigurationRequestWithDefaults instantiates a new TenantConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantConfigurationRequestWithDefaults() *TenantConfigurationRequest {
	this := TenantConfigurationRequest{}
	return &this
}

// GetConfigDetails returns the ConfigDetails field value if set, zero value otherwise.
func (o *TenantConfigurationRequest) GetConfigDetails() TenantConfigurationDetails {
	if o == nil || IsNil(o.ConfigDetails) {
		var ret TenantConfigurationDetails
		return ret
	}
	return *o.ConfigDetails
}

// GetConfigDetailsOk returns a tuple with the ConfigDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfigurationRequest) GetConfigDetailsOk() (*TenantConfigurationDetails, bool) {
	if o == nil || IsNil(o.ConfigDetails) {
		return nil, false
	}
	return o.ConfigDetails, true
}

// HasConfigDetails returns a boolean if a field has been set.
func (o *TenantConfigurationRequest) HasConfigDetails() bool {
	if o != nil && !IsNil(o.ConfigDetails) {
		return true
	}

	return false
}

// SetConfigDetails gets a reference to the given TenantConfigurationDetails and assigns it to the ConfigDetails field.
func (o *TenantConfigurationRequest) SetConfigDetails(v TenantConfigurationDetails) {
	o.ConfigDetails = &v
}

func (o TenantConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigDetails) {
		toSerialize["configDetails"] = o.ConfigDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TenantConfigurationRequest) UnmarshalJSON(data []byte) (err error) {
	varTenantConfigurationRequest := _TenantConfigurationRequest{}

	err = json.Unmarshal(data, &varTenantConfigurationRequest)

	if err != nil {
		return err
	}

	*o = TenantConfigurationRequest(varTenantConfigurationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantConfigurationRequest struct {
	value *TenantConfigurationRequest
	isSet bool
}

func (v NullableTenantConfigurationRequest) Get() *TenantConfigurationRequest {
	return v.value
}

func (v *NullableTenantConfigurationRequest) Set(val *TenantConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantConfigurationRequest(val *TenantConfigurationRequest) *NullableTenantConfigurationRequest {
	return &NullableTenantConfigurationRequest{value: val, isSet: true}
}

func (v NullableTenantConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


