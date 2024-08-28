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

// checks if the SourceUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceUsage{}

// SourceUsage struct for SourceUsage
type SourceUsage struct {
	// The first day of the month for which activity is aggregated.
	Date *string `json:"date,omitempty"`
	// The average number of days that accounts were active within this source, for the month.
	Count *float32 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceUsage SourceUsage

// NewSourceUsage instantiates a new SourceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceUsage() *SourceUsage {
	this := SourceUsage{}
	return &this
}

// NewSourceUsageWithDefaults instantiates a new SourceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceUsageWithDefaults() *SourceUsage {
	this := SourceUsage{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SourceUsage) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUsage) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SourceUsage) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SourceUsage) SetDate(v string) {
	o.Date = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SourceUsage) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUsage) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SourceUsage) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *SourceUsage) SetCount(v float32) {
	o.Count = &v
}

func (o SourceUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceUsage) UnmarshalJSON(data []byte) (err error) {
	varSourceUsage := _SourceUsage{}

	err = json.Unmarshal(data, &varSourceUsage)

	if err != nil {
		return err
	}

	*o = SourceUsage(varSourceUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceUsage struct {
	value *SourceUsage
	isSet bool
}

func (v NullableSourceUsage) Get() *SourceUsage {
	return v.value
}

func (v *NullableSourceUsage) Set(val *SourceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceUsage(val *SourceUsage) *NullableSourceUsage {
	return &NullableSourceUsage{value: val, isSet: true}
}

func (v NullableSourceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


