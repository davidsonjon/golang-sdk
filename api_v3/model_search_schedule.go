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

// checks if the SearchSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSchedule{}

// SearchSchedule struct for SearchSchedule
type SearchSchedule struct {
	// The ID of the saved search that will be executed.
	SavedSearchId string `json:"savedSearchId"`
	// The date the scheduled search was initially created.
	Created NullableTime `json:"created,omitempty"`
	// The last date the scheduled search was modified.
	Modified NullableTime `json:"modified,omitempty"`
	Schedule Schedule1 `json:"schedule"`
	// A list of identities that should receive the scheduled search report via email.
	Recipients []SearchScheduleRecipientsInner `json:"recipients"`
	// Indicates if the scheduled search is enabled. 
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates if email generation should occur when search returns no results. 
	EmailEmptyResults *bool `json:"emailEmptyResults,omitempty"`
	// Indicates if the generated email should include the query and search results preview (which could include PII). 
	DisplayQueryDetails *bool `json:"displayQueryDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchSchedule SearchSchedule

// NewSearchSchedule instantiates a new SearchSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSchedule(savedSearchId string, schedule Schedule1, recipients []SearchScheduleRecipientsInner) *SearchSchedule {
	this := SearchSchedule{}
	this.SavedSearchId = savedSearchId
	this.Schedule = schedule
	this.Recipients = recipients
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// NewSearchScheduleWithDefaults instantiates a new SearchSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchScheduleWithDefaults() *SearchSchedule {
	this := SearchSchedule{}
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// GetSavedSearchId returns the SavedSearchId field value
func (o *SearchSchedule) GetSavedSearchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SavedSearchId
}

// GetSavedSearchIdOk returns a tuple with the SavedSearchId field value
// and a boolean to check if the value has been set.
func (o *SearchSchedule) GetSavedSearchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavedSearchId, true
}

// SetSavedSearchId sets field value
func (o *SearchSchedule) SetSavedSearchId(v string) {
	o.SavedSearchId = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchSchedule) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchSchedule) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *SearchSchedule) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *SearchSchedule) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *SearchSchedule) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *SearchSchedule) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchSchedule) GetModified() time.Time {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchSchedule) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *SearchSchedule) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *SearchSchedule) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *SearchSchedule) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *SearchSchedule) UnsetModified() {
	o.Modified.Unset()
}

// GetSchedule returns the Schedule field value
func (o *SearchSchedule) GetSchedule() Schedule1 {
	if o == nil {
		var ret Schedule1
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *SearchSchedule) GetScheduleOk() (*Schedule1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *SearchSchedule) SetSchedule(v Schedule1) {
	o.Schedule = v
}

// GetRecipients returns the Recipients field value
func (o *SearchSchedule) GetRecipients() []SearchScheduleRecipientsInner {
	if o == nil {
		var ret []SearchScheduleRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *SearchSchedule) GetRecipientsOk() ([]SearchScheduleRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *SearchSchedule) SetRecipients(v []SearchScheduleRecipientsInner) {
	o.Recipients = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SearchSchedule) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSchedule) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SearchSchedule) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SearchSchedule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEmailEmptyResults returns the EmailEmptyResults field value if set, zero value otherwise.
func (o *SearchSchedule) GetEmailEmptyResults() bool {
	if o == nil || IsNil(o.EmailEmptyResults) {
		var ret bool
		return ret
	}
	return *o.EmailEmptyResults
}

// GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSchedule) GetEmailEmptyResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEmptyResults) {
		return nil, false
	}
	return o.EmailEmptyResults, true
}

// HasEmailEmptyResults returns a boolean if a field has been set.
func (o *SearchSchedule) HasEmailEmptyResults() bool {
	if o != nil && !IsNil(o.EmailEmptyResults) {
		return true
	}

	return false
}

// SetEmailEmptyResults gets a reference to the given bool and assigns it to the EmailEmptyResults field.
func (o *SearchSchedule) SetEmailEmptyResults(v bool) {
	o.EmailEmptyResults = &v
}

// GetDisplayQueryDetails returns the DisplayQueryDetails field value if set, zero value otherwise.
func (o *SearchSchedule) GetDisplayQueryDetails() bool {
	if o == nil || IsNil(o.DisplayQueryDetails) {
		var ret bool
		return ret
	}
	return *o.DisplayQueryDetails
}

// GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSchedule) GetDisplayQueryDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayQueryDetails) {
		return nil, false
	}
	return o.DisplayQueryDetails, true
}

// HasDisplayQueryDetails returns a boolean if a field has been set.
func (o *SearchSchedule) HasDisplayQueryDetails() bool {
	if o != nil && !IsNil(o.DisplayQueryDetails) {
		return true
	}

	return false
}

// SetDisplayQueryDetails gets a reference to the given bool and assigns it to the DisplayQueryDetails field.
func (o *SearchSchedule) SetDisplayQueryDetails(v bool) {
	o.DisplayQueryDetails = &v
}

func (o SearchSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["savedSearchId"] = o.SavedSearchId
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	toSerialize["schedule"] = o.Schedule
	toSerialize["recipients"] = o.Recipients
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EmailEmptyResults) {
		toSerialize["emailEmptyResults"] = o.EmailEmptyResults
	}
	if !IsNil(o.DisplayQueryDetails) {
		toSerialize["displayQueryDetails"] = o.DisplayQueryDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"savedSearchId",
		"schedule",
		"recipients",
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

	varSearchSchedule := _SearchSchedule{}

	err = json.Unmarshal(data, &varSearchSchedule)

	if err != nil {
		return err
	}

	*o = SearchSchedule(varSearchSchedule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "savedSearchId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "emailEmptyResults")
		delete(additionalProperties, "displayQueryDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchSchedule struct {
	value *SearchSchedule
	isSet bool
}

func (v NullableSearchSchedule) Get() *SearchSchedule {
	return v.value
}

func (v *NullableSearchSchedule) Set(val *SearchSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSchedule(val *SearchSchedule) *NullableSearchSchedule {
	return &NullableSearchSchedule{value: val, isSet: true}
}

func (v NullableSearchSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


