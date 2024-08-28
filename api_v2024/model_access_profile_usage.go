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

// checks if the AccessProfileUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfileUsage{}

// AccessProfileUsage struct for AccessProfileUsage
type AccessProfileUsage struct {
	// ID of the Access Profile that is in use
	AccessProfileId *string `json:"accessProfileId,omitempty"`
	// List of references to objects which are using the indicated Access Profile
	UsedBy []AccessProfileUsageUsedByInner `json:"usedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileUsage AccessProfileUsage

// NewAccessProfileUsage instantiates a new AccessProfileUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileUsage() *AccessProfileUsage {
	this := AccessProfileUsage{}
	return &this
}

// NewAccessProfileUsageWithDefaults instantiates a new AccessProfileUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileUsageWithDefaults() *AccessProfileUsage {
	this := AccessProfileUsage{}
	return &this
}

// GetAccessProfileId returns the AccessProfileId field value if set, zero value otherwise.
func (o *AccessProfileUsage) GetAccessProfileId() string {
	if o == nil || IsNil(o.AccessProfileId) {
		var ret string
		return ret
	}
	return *o.AccessProfileId
}

// GetAccessProfileIdOk returns a tuple with the AccessProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileUsage) GetAccessProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessProfileId) {
		return nil, false
	}
	return o.AccessProfileId, true
}

// HasAccessProfileId returns a boolean if a field has been set.
func (o *AccessProfileUsage) HasAccessProfileId() bool {
	if o != nil && !IsNil(o.AccessProfileId) {
		return true
	}

	return false
}

// SetAccessProfileId gets a reference to the given string and assigns it to the AccessProfileId field.
func (o *AccessProfileUsage) SetAccessProfileId(v string) {
	o.AccessProfileId = &v
}

// GetUsedBy returns the UsedBy field value if set, zero value otherwise.
func (o *AccessProfileUsage) GetUsedBy() []AccessProfileUsageUsedByInner {
	if o == nil || IsNil(o.UsedBy) {
		var ret []AccessProfileUsageUsedByInner
		return ret
	}
	return o.UsedBy
}

// GetUsedByOk returns a tuple with the UsedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileUsage) GetUsedByOk() ([]AccessProfileUsageUsedByInner, bool) {
	if o == nil || IsNil(o.UsedBy) {
		return nil, false
	}
	return o.UsedBy, true
}

// HasUsedBy returns a boolean if a field has been set.
func (o *AccessProfileUsage) HasUsedBy() bool {
	if o != nil && !IsNil(o.UsedBy) {
		return true
	}

	return false
}

// SetUsedBy gets a reference to the given []AccessProfileUsageUsedByInner and assigns it to the UsedBy field.
func (o *AccessProfileUsage) SetUsedBy(v []AccessProfileUsageUsedByInner) {
	o.UsedBy = v
}

func (o AccessProfileUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfileUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessProfileId) {
		toSerialize["accessProfileId"] = o.AccessProfileId
	}
	if !IsNil(o.UsedBy) {
		toSerialize["usedBy"] = o.UsedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessProfileUsage) UnmarshalJSON(data []byte) (err error) {
	varAccessProfileUsage := _AccessProfileUsage{}

	err = json.Unmarshal(data, &varAccessProfileUsage)

	if err != nil {
		return err
	}

	*o = AccessProfileUsage(varAccessProfileUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessProfileId")
		delete(additionalProperties, "usedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileUsage struct {
	value *AccessProfileUsage
	isSet bool
}

func (v NullableAccessProfileUsage) Get() *AccessProfileUsage {
	return v.value
}

func (v *NullableAccessProfileUsage) Set(val *AccessProfileUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileUsage(val *AccessProfileUsage) *NullableAccessProfileUsage {
	return &NullableAccessProfileUsage{value: val, isSet: true}
}

func (v NullableAccessProfileUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


