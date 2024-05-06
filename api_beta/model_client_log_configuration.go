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

// checks if the ClientLogConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientLogConfiguration{}

// ClientLogConfiguration Client Runtime Logging Configuration
type ClientLogConfiguration struct {
	// Log configuration's client ID
	ClientId *string `json:"clientId,omitempty"`
	// Duration in minutes for log configuration to remain in effect before resetting to defaults
	DurationMinutes int32 `json:"durationMinutes"`
	// Expiration date-time of the log configuration request
	Expiration *time.Time `json:"expiration,omitempty"`
	RootLevel StandardLevel `json:"rootLevel"`
	// Mapping of identifiers to Standard Log Level values
	LogLevels *map[string]StandardLevel `json:"logLevels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientLogConfiguration ClientLogConfiguration

// NewClientLogConfiguration instantiates a new ClientLogConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLogConfiguration(durationMinutes int32, rootLevel StandardLevel) *ClientLogConfiguration {
	this := ClientLogConfiguration{}
	this.DurationMinutes = durationMinutes
	this.RootLevel = rootLevel
	return &this
}

// NewClientLogConfigurationWithDefaults instantiates a new ClientLogConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLogConfigurationWithDefaults() *ClientLogConfiguration {
	this := ClientLogConfiguration{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientLogConfiguration) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfiguration) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientLogConfiguration) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientLogConfiguration) SetClientId(v string) {
	o.ClientId = &v
}

// GetDurationMinutes returns the DurationMinutes field value
func (o *ClientLogConfiguration) GetDurationMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMinutes
}

// GetDurationMinutesOk returns a tuple with the DurationMinutes field value
// and a boolean to check if the value has been set.
func (o *ClientLogConfiguration) GetDurationMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMinutes, true
}

// SetDurationMinutes sets field value
func (o *ClientLogConfiguration) SetDurationMinutes(v int32) {
	o.DurationMinutes = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ClientLogConfiguration) GetExpiration() time.Time {
	if o == nil || isNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfiguration) GetExpirationOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ClientLogConfiguration) HasExpiration() bool {
	if o != nil && !isNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *ClientLogConfiguration) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetRootLevel returns the RootLevel field value
func (o *ClientLogConfiguration) GetRootLevel() StandardLevel {
	if o == nil {
		var ret StandardLevel
		return ret
	}

	return o.RootLevel
}

// GetRootLevelOk returns a tuple with the RootLevel field value
// and a boolean to check if the value has been set.
func (o *ClientLogConfiguration) GetRootLevelOk() (*StandardLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootLevel, true
}

// SetRootLevel sets field value
func (o *ClientLogConfiguration) SetRootLevel(v StandardLevel) {
	o.RootLevel = v
}

// GetLogLevels returns the LogLevels field value if set, zero value otherwise.
func (o *ClientLogConfiguration) GetLogLevels() map[string]StandardLevel {
	if o == nil || isNil(o.LogLevels) {
		var ret map[string]StandardLevel
		return ret
	}
	return *o.LogLevels
}

// GetLogLevelsOk returns a tuple with the LogLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfiguration) GetLogLevelsOk() (*map[string]StandardLevel, bool) {
	if o == nil || isNil(o.LogLevels) {
		return nil, false
	}
	return o.LogLevels, true
}

// HasLogLevels returns a boolean if a field has been set.
func (o *ClientLogConfiguration) HasLogLevels() bool {
	if o != nil && !isNil(o.LogLevels) {
		return true
	}

	return false
}

// SetLogLevels gets a reference to the given map[string]StandardLevel and assigns it to the LogLevels field.
func (o *ClientLogConfiguration) SetLogLevels(v map[string]StandardLevel) {
	o.LogLevels = &v
}

func (o ClientLogConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientLogConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	toSerialize["durationMinutes"] = o.DurationMinutes
	if !isNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	toSerialize["rootLevel"] = o.RootLevel
	if !isNil(o.LogLevels) {
		toSerialize["logLevels"] = o.LogLevels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientLogConfiguration) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"durationMinutes",
		"rootLevel",
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

	varClientLogConfiguration := _ClientLogConfiguration{}

	if err = json.Unmarshal(bytes, &varClientLogConfiguration); err == nil {
	*o = ClientLogConfiguration(varClientLogConfiguration)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "durationMinutes")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "rootLevel")
		delete(additionalProperties, "logLevels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientLogConfiguration struct {
	value *ClientLogConfiguration
	isSet bool
}

func (v NullableClientLogConfiguration) Get() *ClientLogConfiguration {
	return v.value
}

func (v *NullableClientLogConfiguration) Set(val *ClientLogConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLogConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLogConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLogConfiguration(val *ClientLogConfiguration) *NullableClientLogConfiguration {
	return &NullableClientLogConfiguration{value: val, isSet: true}
}

func (v NullableClientLogConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLogConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


