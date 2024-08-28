/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the ManualDiscoverApplications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualDiscoverApplications{}

// ManualDiscoverApplications struct for ManualDiscoverApplications
type ManualDiscoverApplications struct {
	// The CSV file to upload containing `application_name` and `description` columns. Each row represents an application to be discovered.
	File *os.File `json:"file"`
	AdditionalProperties map[string]interface{}
}

type _ManualDiscoverApplications ManualDiscoverApplications

// NewManualDiscoverApplications instantiates a new ManualDiscoverApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualDiscoverApplications(file *os.File) *ManualDiscoverApplications {
	this := ManualDiscoverApplications{}
	this.File = file
	return &this
}

// NewManualDiscoverApplicationsWithDefaults instantiates a new ManualDiscoverApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualDiscoverApplicationsWithDefaults() *ManualDiscoverApplications {
	this := ManualDiscoverApplications{}
	return &this
}

// GetFile returns the File field value
func (o *ManualDiscoverApplications) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ManualDiscoverApplications) GetFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ManualDiscoverApplications) SetFile(v *os.File) {
	o.File = v
}

func (o ManualDiscoverApplications) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualDiscoverApplications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManualDiscoverApplications) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
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

	varManualDiscoverApplications := _ManualDiscoverApplications{}

	err = json.Unmarshal(data, &varManualDiscoverApplications)

	if err != nil {
		return err
	}

	*o = ManualDiscoverApplications(varManualDiscoverApplications)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManualDiscoverApplications struct {
	value *ManualDiscoverApplications
	isSet bool
}

func (v NullableManualDiscoverApplications) Get() *ManualDiscoverApplications {
	return v.value
}

func (v *NullableManualDiscoverApplications) Set(val *ManualDiscoverApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableManualDiscoverApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableManualDiscoverApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualDiscoverApplications(val *ManualDiscoverApplications) *NullableManualDiscoverApplications {
	return &NullableManualDiscoverApplications{value: val, isSet: true}
}

func (v NullableManualDiscoverApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualDiscoverApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


