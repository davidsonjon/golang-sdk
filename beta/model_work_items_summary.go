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

// checks if the WorkItemsSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemsSummary{}

// WorkItemsSummary struct for WorkItemsSummary
type WorkItemsSummary struct {
	// The count of open work items
	Open *int32 `json:"open,omitempty"`
	// The count of completed work items
	Completed *int32 `json:"completed,omitempty"`
	// The count of total work items
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkItemsSummary WorkItemsSummary

// NewWorkItemsSummary instantiates a new WorkItemsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemsSummary() *WorkItemsSummary {
	this := WorkItemsSummary{}
	return &this
}

// NewWorkItemsSummaryWithDefaults instantiates a new WorkItemsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemsSummaryWithDefaults() *WorkItemsSummary {
	this := WorkItemsSummary{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *WorkItemsSummary) GetOpen() int32 {
	if o == nil || isNil(o.Open) {
		var ret int32
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemsSummary) GetOpenOk() (*int32, bool) {
	if o == nil || isNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *WorkItemsSummary) HasOpen() bool {
	if o != nil && !isNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given int32 and assigns it to the Open field.
func (o *WorkItemsSummary) SetOpen(v int32) {
	o.Open = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *WorkItemsSummary) GetCompleted() int32 {
	if o == nil || isNil(o.Completed) {
		var ret int32
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemsSummary) GetCompletedOk() (*int32, bool) {
	if o == nil || isNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *WorkItemsSummary) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given int32 and assigns it to the Completed field.
func (o *WorkItemsSummary) SetCompleted(v int32) {
	o.Completed = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *WorkItemsSummary) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemsSummary) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *WorkItemsSummary) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *WorkItemsSummary) SetTotal(v int32) {
	o.Total = &v
}

func (o WorkItemsSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemsSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkItemsSummary) UnmarshalJSON(bytes []byte) (err error) {
	varWorkItemsSummary := _WorkItemsSummary{}

	if err = json.Unmarshal(bytes, &varWorkItemsSummary); err == nil {
		*o = WorkItemsSummary(varWorkItemsSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "open")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkItemsSummary struct {
	value *WorkItemsSummary
	isSet bool
}

func (v NullableWorkItemsSummary) Get() *WorkItemsSummary {
	return v.value
}

func (v *NullableWorkItemsSummary) Set(val *WorkItemsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemsSummary(val *WorkItemsSummary) *NullableWorkItemsSummary {
	return &NullableWorkItemsSummary{value: val, isSet: true}
}

func (v NullableWorkItemsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


