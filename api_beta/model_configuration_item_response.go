/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the ConfigurationItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationItemResponse{}

// ConfigurationItemResponse The response body of a Reassignment Configuration for a single identity
type ConfigurationItemResponse struct {
	Identity *Identity1 `json:"identity,omitempty"`
	// Details of how work should be reassigned for an Identity
	ConfigDetails []ConfigurationDetailsResponse `json:"configDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationItemResponse ConfigurationItemResponse

// NewConfigurationItemResponse instantiates a new ConfigurationItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationItemResponse() *ConfigurationItemResponse {
	this := ConfigurationItemResponse{}
	return &this
}

// NewConfigurationItemResponseWithDefaults instantiates a new ConfigurationItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationItemResponseWithDefaults() *ConfigurationItemResponse {
	this := ConfigurationItemResponse{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ConfigurationItemResponse) GetIdentity() Identity1 {
	if o == nil || IsNil(o.Identity) {
		var ret Identity1
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationItemResponse) GetIdentityOk() (*Identity1, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ConfigurationItemResponse) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity1 and assigns it to the Identity field.
func (o *ConfigurationItemResponse) SetIdentity(v Identity1) {
	o.Identity = &v
}

// GetConfigDetails returns the ConfigDetails field value if set, zero value otherwise.
func (o *ConfigurationItemResponse) GetConfigDetails() []ConfigurationDetailsResponse {
	if o == nil || IsNil(o.ConfigDetails) {
		var ret []ConfigurationDetailsResponse
		return ret
	}
	return o.ConfigDetails
}

// GetConfigDetailsOk returns a tuple with the ConfigDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationItemResponse) GetConfigDetailsOk() ([]ConfigurationDetailsResponse, bool) {
	if o == nil || IsNil(o.ConfigDetails) {
		return nil, false
	}
	return o.ConfigDetails, true
}

// HasConfigDetails returns a boolean if a field has been set.
func (o *ConfigurationItemResponse) HasConfigDetails() bool {
	if o != nil && !IsNil(o.ConfigDetails) {
		return true
	}

	return false
}

// SetConfigDetails gets a reference to the given []ConfigurationDetailsResponse and assigns it to the ConfigDetails field.
func (o *ConfigurationItemResponse) SetConfigDetails(v []ConfigurationDetailsResponse) {
	o.ConfigDetails = v
}

func (o ConfigurationItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.ConfigDetails) {
		toSerialize["configDetails"] = o.ConfigDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationItemResponse) UnmarshalJSON(data []byte) (err error) {
	varConfigurationItemResponse := _ConfigurationItemResponse{}

	err = json.Unmarshal(data, &varConfigurationItemResponse)

	if err != nil {
		return err
	}

	*o = ConfigurationItemResponse(varConfigurationItemResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "configDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationItemResponse struct {
	value *ConfigurationItemResponse
	isSet bool
}

func (v NullableConfigurationItemResponse) Get() *ConfigurationItemResponse {
	return v.value
}

func (v *NullableConfigurationItemResponse) Set(val *ConfigurationItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationItemResponse(val *ConfigurationItemResponse) *NullableConfigurationItemResponse {
	return &NullableConfigurationItemResponse{value: val, isSet: true}
}

func (v NullableConfigurationItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


