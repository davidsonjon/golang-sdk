/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Schedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule{}

// Schedule struct for Schedule
type Schedule struct {
	// Determines the overall schedule cadence. In general, all time period fields smaller than the chosen type can be configured. For example, a DAILY schedule can have 'hours' set, but not 'days'; a WEEKLY schedule can have both 'hours' and 'days' set.
	Type string `json:"type"`
	Months *ScheduleMonths `json:"months,omitempty"`
	Days *ScheduleDays `json:"days,omitempty"`
	Hours ScheduleHours `json:"hours"`
	// Specifies the time after which this schedule will no longer occur.
	Expiration *time.Time `json:"expiration,omitempty"`
	// The time zone to use when running the schedule. For instance, if the schedule is scheduled to run at 1AM, and this field is set to \"CST\", the schedule will run at 1AM CST.
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule Schedule

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule(type_ string, hours ScheduleHours) *Schedule {
	this := Schedule{}
	this.Type = type_
	this.Hours = hours
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule) SetType(v string) {
	o.Type = v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *Schedule) GetMonths() ScheduleMonths {
	if o == nil || IsNil(o.Months) {
		var ret ScheduleMonths
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetMonthsOk() (*ScheduleMonths, bool) {
	if o == nil || IsNil(o.Months) {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *Schedule) HasMonths() bool {
	if o != nil && !IsNil(o.Months) {
		return true
	}

	return false
}

// SetMonths gets a reference to the given ScheduleMonths and assigns it to the Months field.
func (o *Schedule) SetMonths(v ScheduleMonths) {
	o.Months = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *Schedule) GetDays() ScheduleDays {
	if o == nil || IsNil(o.Days) {
		var ret ScheduleDays
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDaysOk() (*ScheduleDays, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *Schedule) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given ScheduleDays and assigns it to the Days field.
func (o *Schedule) SetDays(v ScheduleDays) {
	o.Days = &v
}

// GetHours returns the Hours field value
func (o *Schedule) GetHours() ScheduleHours {
	if o == nil {
		var ret ScheduleHours
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetHoursOk() (*ScheduleHours, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value
func (o *Schedule) SetHours(v ScheduleHours) {
	o.Hours = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Schedule) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Schedule) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *Schedule) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetTimeZoneId returns the TimeZoneId field value if set, zero value otherwise.
func (o *Schedule) GetTimeZoneId() string {
	if o == nil || IsNil(o.TimeZoneId) {
		var ret string
		return ret
	}
	return *o.TimeZoneId
}

// GetTimeZoneIdOk returns a tuple with the TimeZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimeZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZoneId) {
		return nil, false
	}
	return o.TimeZoneId, true
}

// HasTimeZoneId returns a boolean if a field has been set.
func (o *Schedule) HasTimeZoneId() bool {
	if o != nil && !IsNil(o.TimeZoneId) {
		return true
	}

	return false
}

// SetTimeZoneId gets a reference to the given string and assigns it to the TimeZoneId field.
func (o *Schedule) SetTimeZoneId(v string) {
	o.TimeZoneId = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Months) {
		toSerialize["months"] = o.Months
	}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	toSerialize["hours"] = o.Hours
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.TimeZoneId) {
		toSerialize["timeZoneId"] = o.TimeZoneId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schedule) UnmarshalJSON(data []byte) (err error) {
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

	varSchedule := _Schedule{}

	err = json.Unmarshal(data, &varSchedule)

	if err != nil {
		return err
	}

	*o = Schedule(varSchedule)

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

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


