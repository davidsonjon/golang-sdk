/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the ReportResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportResults{}

// ReportResults Details about report result or current state.
type ReportResults struct {
	// Use this property to define what report should be processed in the RDE service.
	ReportType *string `json:"reportType,omitempty"`
	// Name of the task definition which is started to process requesting report. Usually the same as report name
	TaskDefName *string `json:"taskDefName,omitempty"`
	// Unique task definition identifier.
	Id *string `json:"id,omitempty"`
	// Report processing start date
	Created *time.Time `json:"created,omitempty"`
	// Report current state or result status.
	Status *string `json:"status,omitempty"`
	// Report processing time in ms.
	Duration *int64 `json:"duration,omitempty"`
	// Report size in rows.
	Rows *int64 `json:"rows,omitempty"`
	// Output report file formats. This are formats for calling get endpoint as a query parameter 'fileFormat'.  In case report won't have this argument there will be ['CSV', 'PDF'] as default.
	AvailableFormats []string `json:"availableFormats,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReportResults ReportResults

// NewReportResults instantiates a new ReportResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResults() *ReportResults {
	this := ReportResults{}
	return &this
}

// NewReportResultsWithDefaults instantiates a new ReportResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResultsWithDefaults() *ReportResults {
	this := ReportResults{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *ReportResults) GetReportType() string {
	if o == nil || isNil(o.ReportType) {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetReportTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReportType) {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportResults) HasReportType() bool {
	if o != nil && !isNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportResults) SetReportType(v string) {
	o.ReportType = &v
}

// GetTaskDefName returns the TaskDefName field value if set, zero value otherwise.
func (o *ReportResults) GetTaskDefName() string {
	if o == nil || isNil(o.TaskDefName) {
		var ret string
		return ret
	}
	return *o.TaskDefName
}

// GetTaskDefNameOk returns a tuple with the TaskDefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetTaskDefNameOk() (*string, bool) {
	if o == nil || isNil(o.TaskDefName) {
		return nil, false
	}
	return o.TaskDefName, true
}

// HasTaskDefName returns a boolean if a field has been set.
func (o *ReportResults) HasTaskDefName() bool {
	if o != nil && !isNil(o.TaskDefName) {
		return true
	}

	return false
}

// SetTaskDefName gets a reference to the given string and assigns it to the TaskDefName field.
func (o *ReportResults) SetTaskDefName(v string) {
	o.TaskDefName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportResults) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportResults) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportResults) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ReportResults) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ReportResults) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ReportResults) SetCreated(v time.Time) {
	o.Created = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportResults) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportResults) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportResults) SetStatus(v string) {
	o.Status = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ReportResults) GetDuration() int64 {
	if o == nil || isNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetDurationOk() (*int64, bool) {
	if o == nil || isNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ReportResults) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *ReportResults) SetDuration(v int64) {
	o.Duration = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *ReportResults) GetRows() int64 {
	if o == nil || isNil(o.Rows) {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetRowsOk() (*int64, bool) {
	if o == nil || isNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ReportResults) HasRows() bool {
	if o != nil && !isNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *ReportResults) SetRows(v int64) {
	o.Rows = &v
}

// GetAvailableFormats returns the AvailableFormats field value if set, zero value otherwise.
func (o *ReportResults) GetAvailableFormats() []string {
	if o == nil || isNil(o.AvailableFormats) {
		var ret []string
		return ret
	}
	return o.AvailableFormats
}

// GetAvailableFormatsOk returns a tuple with the AvailableFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResults) GetAvailableFormatsOk() ([]string, bool) {
	if o == nil || isNil(o.AvailableFormats) {
		return nil, false
	}
	return o.AvailableFormats, true
}

// HasAvailableFormats returns a boolean if a field has been set.
func (o *ReportResults) HasAvailableFormats() bool {
	if o != nil && !isNil(o.AvailableFormats) {
		return true
	}

	return false
}

// SetAvailableFormats gets a reference to the given []string and assigns it to the AvailableFormats field.
func (o *ReportResults) SetAvailableFormats(v []string) {
	o.AvailableFormats = v
}

func (o ReportResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !isNil(o.TaskDefName) {
		toSerialize["taskDefName"] = o.TaskDefName
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !isNil(o.AvailableFormats) {
		toSerialize["availableFormats"] = o.AvailableFormats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportResults) UnmarshalJSON(bytes []byte) (err error) {
	varReportResults := _ReportResults{}

	if err = json.Unmarshal(bytes, &varReportResults); err == nil {
	*o = ReportResults(varReportResults)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reportType")
		delete(additionalProperties, "taskDefName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "status")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "rows")
		delete(additionalProperties, "availableFormats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportResults struct {
	value *ReportResults
	isSet bool
}

func (v NullableReportResults) Get() *ReportResults {
	return v.value
}

func (v *NullableReportResults) Set(val *ReportResults) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResults) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResults(val *ReportResults) *NullableReportResults {
	return &NullableReportResults{value: val, isSet: true}
}

func (v NullableReportResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


