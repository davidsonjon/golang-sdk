/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the SpConfigImportResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpConfigImportResults{}

// SpConfigImportResults Response Body for Config Import command.
type SpConfigImportResults struct {
	// The results of an object configuration import job.
	Results map[string]ObjectImportResult1 `json:"results"`
	// If a backup was performed before the import, this will contain the jobId of the backup job. This id can be used to retrieve the json file of the backup export.
	ExportJobId *string `json:"exportJobId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpConfigImportResults SpConfigImportResults

// NewSpConfigImportResults instantiates a new SpConfigImportResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpConfigImportResults(results map[string]ObjectImportResult1) *SpConfigImportResults {
	this := SpConfigImportResults{}
	this.Results = results
	return &this
}

// NewSpConfigImportResultsWithDefaults instantiates a new SpConfigImportResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpConfigImportResultsWithDefaults() *SpConfigImportResults {
	this := SpConfigImportResults{}
	return &this
}

// GetResults returns the Results field value
func (o *SpConfigImportResults) GetResults() map[string]ObjectImportResult1 {
	if o == nil {
		var ret map[string]ObjectImportResult1
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SpConfigImportResults) GetResultsOk() (*map[string]ObjectImportResult1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *SpConfigImportResults) SetResults(v map[string]ObjectImportResult1) {
	o.Results = v
}

// GetExportJobId returns the ExportJobId field value if set, zero value otherwise.
func (o *SpConfigImportResults) GetExportJobId() string {
	if o == nil || isNil(o.ExportJobId) {
		var ret string
		return ret
	}
	return *o.ExportJobId
}

// GetExportJobIdOk returns a tuple with the ExportJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigImportResults) GetExportJobIdOk() (*string, bool) {
	if o == nil || isNil(o.ExportJobId) {
		return nil, false
	}
	return o.ExportJobId, true
}

// HasExportJobId returns a boolean if a field has been set.
func (o *SpConfigImportResults) HasExportJobId() bool {
	if o != nil && !isNil(o.ExportJobId) {
		return true
	}

	return false
}

// SetExportJobId gets a reference to the given string and assigns it to the ExportJobId field.
func (o *SpConfigImportResults) SetExportJobId(v string) {
	o.ExportJobId = &v
}

func (o SpConfigImportResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpConfigImportResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	if !isNil(o.ExportJobId) {
		toSerialize["exportJobId"] = o.ExportJobId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpConfigImportResults) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varSpConfigImportResults := _SpConfigImportResults{}

	if err = json.Unmarshal(bytes, &varSpConfigImportResults); err == nil {
	*o = SpConfigImportResults(varSpConfigImportResults)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		delete(additionalProperties, "exportJobId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpConfigImportResults struct {
	value *SpConfigImportResults
	isSet bool
}

func (v NullableSpConfigImportResults) Get() *SpConfigImportResults {
	return v.value
}

func (v *NullableSpConfigImportResults) Set(val *SpConfigImportResults) {
	v.value = val
	v.isSet = true
}

func (v NullableSpConfigImportResults) IsSet() bool {
	return v.isSet
}

func (v *NullableSpConfigImportResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpConfigImportResults(val *SpConfigImportResults) *NullableSpConfigImportResults {
	return &NullableSpConfigImportResults{value: val, isSet: true}
}

func (v NullableSpConfigImportResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpConfigImportResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


