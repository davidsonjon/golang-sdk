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

// checks if the ReportDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportDetails{}

// ReportDetails Details about report to be processed.
type ReportDetails struct {
	// Use this property to define what report should be processed in the RDE service.
	ReportType *string `json:"reportType,omitempty"`
	Arguments *ReportDetailsArguments `json:"arguments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReportDetails ReportDetails

// NewReportDetails instantiates a new ReportDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportDetails() *ReportDetails {
	this := ReportDetails{}
	return &this
}

// NewReportDetailsWithDefaults instantiates a new ReportDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportDetailsWithDefaults() *ReportDetails {
	this := ReportDetails{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *ReportDetails) GetReportType() string {
	if o == nil || isNil(o.ReportType) {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDetails) GetReportTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReportType) {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportDetails) HasReportType() bool {
	if o != nil && !isNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportDetails) SetReportType(v string) {
	o.ReportType = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ReportDetails) GetArguments() ReportDetailsArguments {
	if o == nil || isNil(o.Arguments) {
		var ret ReportDetailsArguments
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDetails) GetArgumentsOk() (*ReportDetailsArguments, bool) {
	if o == nil || isNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ReportDetails) HasArguments() bool {
	if o != nil && !isNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given ReportDetailsArguments and assigns it to the Arguments field.
func (o *ReportDetails) SetArguments(v ReportDetailsArguments) {
	o.Arguments = &v
}

func (o ReportDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !isNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportDetails) UnmarshalJSON(bytes []byte) (err error) {
	varReportDetails := _ReportDetails{}

	if err = json.Unmarshal(bytes, &varReportDetails); err == nil {
		*o = ReportDetails(varReportDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reportType")
		delete(additionalProperties, "arguments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportDetails struct {
	value *ReportDetails
	isSet bool
}

func (v NullableReportDetails) Get() *ReportDetails {
	return v.value
}

func (v *NullableReportDetails) Set(val *ReportDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableReportDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableReportDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportDetails(val *ReportDetails) *NullableReportDetails {
	return &NullableReportDetails{value: val, isSet: true}
}

func (v NullableReportDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


