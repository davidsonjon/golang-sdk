/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the NetworkConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkConfiguration{}

// NetworkConfiguration struct for NetworkConfiguration
type NetworkConfiguration struct {
	// The collection of ip ranges.
	Range []string `json:"range,omitempty"`
	// The collection of country codes.
	Geolocation []string `json:"geolocation,omitempty"`
	// Denotes whether the provided lists are whitelisted or blacklisted for geo location.
	Whitelisted *bool `json:"whitelisted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkConfiguration NetworkConfiguration

// NewNetworkConfiguration instantiates a new NetworkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfiguration() *NetworkConfiguration {
	this := NetworkConfiguration{}
	var whitelisted bool = false
	this.Whitelisted = &whitelisted
	return &this
}

// NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigurationWithDefaults() *NetworkConfiguration {
	this := NetworkConfiguration{}
	var whitelisted bool = false
	this.Whitelisted = &whitelisted
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConfiguration) GetRange() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConfiguration) GetRangeOk() ([]string, bool) {
	if o == nil || isNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasRange() bool {
	if o != nil && isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given []string and assigns it to the Range field.
func (o *NetworkConfiguration) SetRange(v []string) {
	o.Range = v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConfiguration) GetGeolocation() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConfiguration) GetGeolocationOk() ([]string, bool) {
	if o == nil || isNil(o.Geolocation) {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasGeolocation() bool {
	if o != nil && isNil(o.Geolocation) {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given []string and assigns it to the Geolocation field.
func (o *NetworkConfiguration) SetGeolocation(v []string) {
	o.Geolocation = v
}

// GetWhitelisted returns the Whitelisted field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetWhitelisted() bool {
	if o == nil || isNil(o.Whitelisted) {
		var ret bool
		return ret
	}
	return *o.Whitelisted
}

// GetWhitelistedOk returns a tuple with the Whitelisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetWhitelistedOk() (*bool, bool) {
	if o == nil || isNil(o.Whitelisted) {
		return nil, false
	}
	return o.Whitelisted, true
}

// HasWhitelisted returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasWhitelisted() bool {
	if o != nil && !isNil(o.Whitelisted) {
		return true
	}

	return false
}

// SetWhitelisted gets a reference to the given bool and assigns it to the Whitelisted field.
func (o *NetworkConfiguration) SetWhitelisted(v bool) {
	o.Whitelisted = &v
}

func (o NetworkConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	if !isNil(o.Whitelisted) {
		toSerialize["whitelisted"] = o.Whitelisted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkConfiguration := _NetworkConfiguration{}

	if err = json.Unmarshal(bytes, &varNetworkConfiguration); err == nil {
	*o = NetworkConfiguration(varNetworkConfiguration)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "range")
		delete(additionalProperties, "geolocation")
		delete(additionalProperties, "whitelisted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkConfiguration struct {
	value *NetworkConfiguration
	isSet bool
}

func (v NullableNetworkConfiguration) Get() *NetworkConfiguration {
	return v.value
}

func (v *NullableNetworkConfiguration) Set(val *NetworkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfiguration(val *NetworkConfiguration) *NullableNetworkConfiguration {
	return &NullableNetworkConfiguration{value: val, isSet: true}
}

func (v NullableNetworkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


