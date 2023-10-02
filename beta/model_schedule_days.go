/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the ScheduleDays type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleDays{}

// ScheduleDays Specifies which day(s) a schedule is active for. This is required for all schedule types. The \"values\" field holds different data depending on the type of schedule: * WEEKLY: days of the week (1-7) * MONTHLY: days of the month (1-31, L, L-1...) * ANNUALLY: if the \"months\" field is also set: days of the month (1-31, L, L-1...); otherwise: ISO-8601 dates without year (\"--12-31\") * CALENDAR: ISO-8601 dates (\"2020-12-31\")  Note that CALENDAR only supports the LIST type, and ANNUALLY does not support the RANGE type when provided with ISO-8601 dates without year.  Examples:  On Sundays: * type LIST * values \"1\"  The second to last day of the month: * type LIST * values \"L-1\"  From the 20th to the last day of the month: * type RANGE * values \"20\", \"L\"  Every March 2nd: * type LIST * values \"--03-02\"  On March 2nd, 2021: * type: LIST * values \"2021-03-02\" 
type ScheduleDays struct {
	// Enum type to specify days value
	Type string `json:"type"`
	// Values of the days based on the enum type mentioned above
	Values []string `json:"values"`
	// Interval between the cert generations
	Interval *int64 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleDays ScheduleDays

// NewScheduleDays instantiates a new ScheduleDays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleDays(type_ string, values []string) *ScheduleDays {
	this := ScheduleDays{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewScheduleDaysWithDefaults instantiates a new ScheduleDays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleDaysWithDefaults() *ScheduleDays {
	this := ScheduleDays{}
	return &this
}

// GetType returns the Type field value
func (o *ScheduleDays) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleDays) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleDays) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *ScheduleDays) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ScheduleDays) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ScheduleDays) SetValues(v []string) {
	o.Values = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ScheduleDays) GetInterval() int64 {
	if o == nil || isNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDays) GetIntervalOk() (*int64, bool) {
	if o == nil || isNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ScheduleDays) HasInterval() bool {
	if o != nil && !isNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ScheduleDays) SetInterval(v int64) {
	o.Interval = &v
}

func (o ScheduleDays) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleDays) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values
	if !isNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduleDays) UnmarshalJSON(bytes []byte) (err error) {
	varScheduleDays := _ScheduleDays{}

	if err = json.Unmarshal(bytes, &varScheduleDays); err == nil {
		*o = ScheduleDays(varScheduleDays)
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

type NullableScheduleDays struct {
	value *ScheduleDays
	isSet bool
}

func (v NullableScheduleDays) Get() *ScheduleDays {
	return v.value
}

func (v *NullableScheduleDays) Set(val *ScheduleDays) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleDays) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleDays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleDays(val *ScheduleDays) *NullableScheduleDays {
	return &NullableScheduleDays{value: val, isSet: true}
}

func (v NullableScheduleDays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleDays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


