/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Schedule1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule1{}

// Schedule1 The schedule information.
type Schedule1 struct {
	Type ScheduleType `json:"type"`
	Months *Schedule1Months `json:"months,omitempty"`
	Days *Schedule1Days `json:"days,omitempty"`
	Hours Schedule1Hours `json:"hours"`
	// A date-time in ISO-8601 format
	Expiration NullableTime `json:"expiration,omitempty"`
	// The canonical TZ identifier the schedule will run in (ex. America/New_York).  If no timezone is specified, the org's default timezone is used.
	TimeZoneId NullableString `json:"timeZoneId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule1 Schedule1

// NewSchedule1 instantiates a new Schedule1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule1(type_ ScheduleType, hours Schedule1Hours) *Schedule1 {
	this := Schedule1{}
	this.Type = type_
	this.Hours = hours
	return &this
}

// NewSchedule1WithDefaults instantiates a new Schedule1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule1WithDefaults() *Schedule1 {
	this := Schedule1{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule1) GetType() ScheduleType {
	if o == nil {
		var ret ScheduleType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetTypeOk() (*ScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule1) SetType(v ScheduleType) {
	o.Type = v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *Schedule1) GetMonths() Schedule1Months {
	if o == nil || IsNil(o.Months) {
		var ret Schedule1Months
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule1) GetMonthsOk() (*Schedule1Months, bool) {
	if o == nil || IsNil(o.Months) {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *Schedule1) HasMonths() bool {
	if o != nil && !IsNil(o.Months) {
		return true
	}

	return false
}

// SetMonths gets a reference to the given Schedule1Months and assigns it to the Months field.
func (o *Schedule1) SetMonths(v Schedule1Months) {
	o.Months = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *Schedule1) GetDays() Schedule1Days {
	if o == nil || IsNil(o.Days) {
		var ret Schedule1Days
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule1) GetDaysOk() (*Schedule1Days, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *Schedule1) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given Schedule1Days and assigns it to the Days field.
func (o *Schedule1) SetDays(v Schedule1Days) {
	o.Days = &v
}

// GetHours returns the Hours field value
func (o *Schedule1) GetHours() Schedule1Hours {
	if o == nil {
		var ret Schedule1Hours
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetHoursOk() (*Schedule1Hours, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value
func (o *Schedule1) SetHours(v Schedule1Hours) {
	o.Hours = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *Schedule1) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *Schedule1) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}
// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *Schedule1) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *Schedule1) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetTimeZoneId returns the TimeZoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1) GetTimeZoneId() string {
	if o == nil || IsNil(o.TimeZoneId.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZoneId.Get()
}

// GetTimeZoneIdOk returns a tuple with the TimeZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1) GetTimeZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZoneId.Get(), o.TimeZoneId.IsSet()
}

// HasTimeZoneId returns a boolean if a field has been set.
func (o *Schedule1) HasTimeZoneId() bool {
	if o != nil && o.TimeZoneId.IsSet() {
		return true
	}

	return false
}

// SetTimeZoneId gets a reference to the given NullableString and assigns it to the TimeZoneId field.
func (o *Schedule1) SetTimeZoneId(v string) {
	o.TimeZoneId.Set(&v)
}
// SetTimeZoneIdNil sets the value for TimeZoneId to be an explicit nil
func (o *Schedule1) SetTimeZoneIdNil() {
	o.TimeZoneId.Set(nil)
}

// UnsetTimeZoneId ensures that no value is present for TimeZoneId, not even an explicit nil
func (o *Schedule1) UnsetTimeZoneId() {
	o.TimeZoneId.Unset()
}

func (o Schedule1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Months) {
		toSerialize["months"] = o.Months
	}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	toSerialize["hours"] = o.Hours
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.TimeZoneId.IsSet() {
		toSerialize["timeZoneId"] = o.TimeZoneId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schedule1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"hours",
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

	varSchedule1 := _Schedule1{}

	err = json.Unmarshal(data, &varSchedule1)

	if err != nil {
		return err
	}

	*o = Schedule1(varSchedule1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "months")
		delete(additionalProperties, "days")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "timeZoneId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule1 struct {
	value *Schedule1
	isSet bool
}

func (v NullableSchedule1) Get() *Schedule1 {
	return v.value
}

func (v *NullableSchedule1) Set(val *Schedule1) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule1) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule1(val *Schedule1) *NullableSchedule1 {
	return &NullableSchedule1{value: val, isSet: true}
}

func (v NullableSchedule1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


