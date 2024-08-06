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

// checks if the ManagedClusterRedis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedClusterRedis{}

// ManagedClusterRedis Managed Cluster Redis Configuration
type ManagedClusterRedis struct {
	// ManagedCluster redisHost
	RedisHost *string `json:"redisHost,omitempty"`
	// ManagedCluster redisPort
	RedisPort *int32 `json:"redisPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClusterRedis ManagedClusterRedis

// NewManagedClusterRedis instantiates a new ManagedClusterRedis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClusterRedis() *ManagedClusterRedis {
	this := ManagedClusterRedis{}
	return &this
}

// NewManagedClusterRedisWithDefaults instantiates a new ManagedClusterRedis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClusterRedisWithDefaults() *ManagedClusterRedis {
	this := ManagedClusterRedis{}
	return &this
}

// GetRedisHost returns the RedisHost field value if set, zero value otherwise.
func (o *ManagedClusterRedis) GetRedisHost() string {
	if o == nil || isNil(o.RedisHost) {
		var ret string
		return ret
	}
	return *o.RedisHost
}

// GetRedisHostOk returns a tuple with the RedisHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedClusterRedis) GetRedisHostOk() (*string, bool) {
	if o == nil || isNil(o.RedisHost) {
		return nil, false
	}
	return o.RedisHost, true
}

// HasRedisHost returns a boolean if a field has been set.
func (o *ManagedClusterRedis) HasRedisHost() bool {
	if o != nil && !isNil(o.RedisHost) {
		return true
	}

	return false
}

// SetRedisHost gets a reference to the given string and assigns it to the RedisHost field.
func (o *ManagedClusterRedis) SetRedisHost(v string) {
	o.RedisHost = &v
}

// GetRedisPort returns the RedisPort field value if set, zero value otherwise.
func (o *ManagedClusterRedis) GetRedisPort() int32 {
	if o == nil || isNil(o.RedisPort) {
		var ret int32
		return ret
	}
	return *o.RedisPort
}

// GetRedisPortOk returns a tuple with the RedisPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedClusterRedis) GetRedisPortOk() (*int32, bool) {
	if o == nil || isNil(o.RedisPort) {
		return nil, false
	}
	return o.RedisPort, true
}

// HasRedisPort returns a boolean if a field has been set.
func (o *ManagedClusterRedis) HasRedisPort() bool {
	if o != nil && !isNil(o.RedisPort) {
		return true
	}

	return false
}

// SetRedisPort gets a reference to the given int32 and assigns it to the RedisPort field.
func (o *ManagedClusterRedis) SetRedisPort(v int32) {
	o.RedisPort = &v
}

func (o ManagedClusterRedis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedClusterRedis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RedisHost) {
		toSerialize["redisHost"] = o.RedisHost
	}
	if !isNil(o.RedisPort) {
		toSerialize["redisPort"] = o.RedisPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedClusterRedis) UnmarshalJSON(bytes []byte) (err error) {
	varManagedClusterRedis := _ManagedClusterRedis{}

	if err = json.Unmarshal(bytes, &varManagedClusterRedis); err == nil {
	*o = ManagedClusterRedis(varManagedClusterRedis)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "redisHost")
		delete(additionalProperties, "redisPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedClusterRedis struct {
	value *ManagedClusterRedis
	isSet bool
}

func (v NullableManagedClusterRedis) Get() *ManagedClusterRedis {
	return v.value
}

func (v *NullableManagedClusterRedis) Set(val *ManagedClusterRedis) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClusterRedis) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClusterRedis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClusterRedis(val *ManagedClusterRedis) *NullableManagedClusterRedis {
	return &NullableManagedClusterRedis{value: val, isSet: true}
}

func (v NullableManagedClusterRedis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClusterRedis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


