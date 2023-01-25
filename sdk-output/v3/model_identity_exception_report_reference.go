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

// IdentityExceptionReportReference struct for IdentityExceptionReportReference
type IdentityExceptionReportReference struct {
	// The id of the task result.
	TaskResultId *string `json:"taskResultId,omitempty"`
	// The name of the report.
	ReportName *string `json:"reportName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityExceptionReportReference IdentityExceptionReportReference

// NewIdentityExceptionReportReference instantiates a new IdentityExceptionReportReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityExceptionReportReference() *IdentityExceptionReportReference {
	this := IdentityExceptionReportReference{}
	return &this
}

// NewIdentityExceptionReportReferenceWithDefaults instantiates a new IdentityExceptionReportReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityExceptionReportReferenceWithDefaults() *IdentityExceptionReportReference {
	this := IdentityExceptionReportReference{}
	return &this
}

// GetTaskResultId returns the TaskResultId field value if set, zero value otherwise.
func (o *IdentityExceptionReportReference) GetTaskResultId() string {
	if o == nil || isNil(o.TaskResultId) {
		var ret string
		return ret
	}
	return *o.TaskResultId
}

// GetTaskResultIdOk returns a tuple with the TaskResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityExceptionReportReference) GetTaskResultIdOk() (*string, bool) {
	if o == nil || isNil(o.TaskResultId) {
		return nil, false
	}
	return o.TaskResultId, true
}

// HasTaskResultId returns a boolean if a field has been set.
func (o *IdentityExceptionReportReference) HasTaskResultId() bool {
	if o != nil && !isNil(o.TaskResultId) {
		return true
	}

	return false
}

// SetTaskResultId gets a reference to the given string and assigns it to the TaskResultId field.
func (o *IdentityExceptionReportReference) SetTaskResultId(v string) {
	o.TaskResultId = &v
}

// GetReportName returns the ReportName field value if set, zero value otherwise.
func (o *IdentityExceptionReportReference) GetReportName() string {
	if o == nil || isNil(o.ReportName) {
		var ret string
		return ret
	}
	return *o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityExceptionReportReference) GetReportNameOk() (*string, bool) {
	if o == nil || isNil(o.ReportName) {
		return nil, false
	}
	return o.ReportName, true
}

// HasReportName returns a boolean if a field has been set.
func (o *IdentityExceptionReportReference) HasReportName() bool {
	if o != nil && !isNil(o.ReportName) {
		return true
	}

	return false
}

// SetReportName gets a reference to the given string and assigns it to the ReportName field.
func (o *IdentityExceptionReportReference) SetReportName(v string) {
	o.ReportName = &v
}

func (o IdentityExceptionReportReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TaskResultId) {
		toSerialize["taskResultId"] = o.TaskResultId
	}
	if !isNil(o.ReportName) {
		toSerialize["reportName"] = o.ReportName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityExceptionReportReference) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityExceptionReportReference := _IdentityExceptionReportReference{}

	if err = json.Unmarshal(bytes, &varIdentityExceptionReportReference); err == nil {
		*o = IdentityExceptionReportReference(varIdentityExceptionReportReference)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "taskResultId")
		delete(additionalProperties, "reportName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityExceptionReportReference struct {
	value *IdentityExceptionReportReference
	isSet bool
}

func (v NullableIdentityExceptionReportReference) Get() *IdentityExceptionReportReference {
	return v.value
}

func (v *NullableIdentityExceptionReportReference) Set(val *IdentityExceptionReportReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityExceptionReportReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityExceptionReportReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityExceptionReportReference(val *IdentityExceptionReportReference) *NullableIdentityExceptionReportReference {
	return &NullableIdentityExceptionReportReference{value: val, isSet: true}
}

func (v NullableIdentityExceptionReportReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityExceptionReportReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


