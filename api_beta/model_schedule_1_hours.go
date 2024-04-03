/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the Schedule1Hours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule1Hours{}

// Schedule1Hours struct for Schedule1Hours
type Schedule1Hours struct {
	Type SelectorType `json:"type"`
	// The selected values. 
	Values []string `json:"values"`
	// The selected interval for RANGE selectors. 
	Interval NullableInt32 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule1Hours Schedule1Hours

// NewSchedule1Hours instantiates a new Schedule1Hours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule1Hours(type_ SelectorType, values []string) *Schedule1Hours {
	this := Schedule1Hours{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewSchedule1HoursWithDefaults instantiates a new Schedule1Hours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule1HoursWithDefaults() *Schedule1Hours {
	this := Schedule1Hours{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule1Hours) GetType() SelectorType {
	if o == nil {
		var ret SelectorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule1Hours) GetTypeOk() (*SelectorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule1Hours) SetType(v SelectorType) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *Schedule1Hours) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Schedule1Hours) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *Schedule1Hours) SetValues(v []string) {
	o.Values = v
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1Hours) GetInterval() int32 {
	if o == nil || isNil(o.Interval.Get()) {
		var ret int32
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1Hours) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *Schedule1Hours) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.
func (o *Schedule1Hours) SetInterval(v int32) {
	o.Interval.Set(&v)
}
// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *Schedule1Hours) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *Schedule1Hours) UnsetInterval() {
	o.Interval.Unset()
}

func (o Schedule1Hours) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule1Hours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schedule1Hours) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"values",
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

	varSchedule1Hours := _Schedule1Hours{}

	if err = json.Unmarshal(bytes, &varSchedule1Hours); err == nil {
	*o = Schedule1Hours(varSchedule1Hours)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		delete(additionalProperties, "interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule1Hours struct {
	value *Schedule1Hours
	isSet bool
}

func (v NullableSchedule1Hours) Get() *Schedule1Hours {
	return v.value
}

func (v *NullableSchedule1Hours) Set(val *Schedule1Hours) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule1Hours) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule1Hours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule1Hours(val *Schedule1Hours) *NullableSchedule1Hours {
	return &NullableSchedule1Hours{value: val, isSet: true}
}

func (v NullableSchedule1Hours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule1Hours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


