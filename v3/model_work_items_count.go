/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the WorkItemsCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemsCount{}

// WorkItemsCount struct for WorkItemsCount
type WorkItemsCount struct {
	// The count of work items
	Count *int32 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkItemsCount WorkItemsCount

// NewWorkItemsCount instantiates a new WorkItemsCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemsCount() *WorkItemsCount {
	this := WorkItemsCount{}
	return &this
}

// NewWorkItemsCountWithDefaults instantiates a new WorkItemsCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemsCountWithDefaults() *WorkItemsCount {
	this := WorkItemsCount{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WorkItemsCount) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemsCount) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WorkItemsCount) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *WorkItemsCount) SetCount(v int32) {
	o.Count = &v
}

func (o WorkItemsCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemsCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkItemsCount) UnmarshalJSON(bytes []byte) (err error) {
	varWorkItemsCount := _WorkItemsCount{}

	if err = json.Unmarshal(bytes, &varWorkItemsCount); err == nil {
		*o = WorkItemsCount(varWorkItemsCount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkItemsCount struct {
	value *WorkItemsCount
	isSet bool
}

func (v NullableWorkItemsCount) Get() *WorkItemsCount {
	return v.value
}

func (v *NullableWorkItemsCount) Set(val *WorkItemsCount) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemsCount) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemsCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemsCount(val *WorkItemsCount) *NullableWorkItemsCount {
	return &NullableWorkItemsCount{value: val, isSet: true}
}

func (v NullableWorkItemsCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemsCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


