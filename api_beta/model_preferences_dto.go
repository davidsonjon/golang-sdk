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
)

// checks if the PreferencesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreferencesDto{}

// PreferencesDto Maps an Identity's attribute key to a list of preferred notification mediums.
type PreferencesDto struct {
	// The template notification key.
	Key *string `json:"key,omitempty"`
	// List of preferred notification mediums, i.e., the mediums (or method) for which notifications are enabled. More mediums may be added in the future.
	Mediums []Medium `json:"mediums,omitempty"`
	// Modified date of preference
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreferencesDto PreferencesDto

// NewPreferencesDto instantiates a new PreferencesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferencesDto() *PreferencesDto {
	this := PreferencesDto{}
	return &this
}

// NewPreferencesDtoWithDefaults instantiates a new PreferencesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferencesDtoWithDefaults() *PreferencesDto {
	this := PreferencesDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PreferencesDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PreferencesDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PreferencesDto) SetKey(v string) {
	o.Key = &v
}

// GetMediums returns the Mediums field value if set, zero value otherwise.
func (o *PreferencesDto) GetMediums() []Medium {
	if o == nil || IsNil(o.Mediums) {
		var ret []Medium
		return ret
	}
	return o.Mediums
}

// GetMediumsOk returns a tuple with the Mediums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesDto) GetMediumsOk() ([]Medium, bool) {
	if o == nil || IsNil(o.Mediums) {
		return nil, false
	}
	return o.Mediums, true
}

// HasMediums returns a boolean if a field has been set.
func (o *PreferencesDto) HasMediums() bool {
	if o != nil && !IsNil(o.Mediums) {
		return true
	}

	return false
}

// SetMediums gets a reference to the given []Medium and assigns it to the Mediums field.
func (o *PreferencesDto) SetMediums(v []Medium) {
	o.Mediums = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *PreferencesDto) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *PreferencesDto) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *PreferencesDto) SetModified(v time.Time) {
	o.Modified = &v
}

func (o PreferencesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreferencesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Mediums) {
		toSerialize["mediums"] = o.Mediums
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreferencesDto) UnmarshalJSON(data []byte) (err error) {
	varPreferencesDto := _PreferencesDto{}

	err = json.Unmarshal(data, &varPreferencesDto)

	if err != nil {
		return err
	}

	*o = PreferencesDto(varPreferencesDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "mediums")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreferencesDto struct {
	value *PreferencesDto
	isSet bool
}

func (v NullablePreferencesDto) Get() *PreferencesDto {
	return v.value
}

func (v *NullablePreferencesDto) Set(val *PreferencesDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferencesDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferencesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferencesDto(val *PreferencesDto) *NullablePreferencesDto {
	return &NullablePreferencesDto{value: val, isSet: true}
}

func (v NullablePreferencesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferencesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


