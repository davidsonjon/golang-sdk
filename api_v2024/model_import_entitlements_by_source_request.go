/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"os"
)

// checks if the ImportEntitlementsBySourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportEntitlementsBySourceRequest{}

// ImportEntitlementsBySourceRequest struct for ImportEntitlementsBySourceRequest
type ImportEntitlementsBySourceRequest struct {
	// The CSV file containing the source entitlements to aggregate.
	CsvFile **os.File `json:"csvFile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportEntitlementsBySourceRequest ImportEntitlementsBySourceRequest

// NewImportEntitlementsBySourceRequest instantiates a new ImportEntitlementsBySourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportEntitlementsBySourceRequest() *ImportEntitlementsBySourceRequest {
	this := ImportEntitlementsBySourceRequest{}
	return &this
}

// NewImportEntitlementsBySourceRequestWithDefaults instantiates a new ImportEntitlementsBySourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportEntitlementsBySourceRequestWithDefaults() *ImportEntitlementsBySourceRequest {
	this := ImportEntitlementsBySourceRequest{}
	return &this
}

// GetCsvFile returns the CsvFile field value if set, zero value otherwise.
func (o *ImportEntitlementsBySourceRequest) GetCsvFile() *os.File {
	if o == nil || isNil(o.CsvFile) {
		var ret *os.File
		return ret
	}
	return *o.CsvFile
}

// GetCsvFileOk returns a tuple with the CsvFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEntitlementsBySourceRequest) GetCsvFileOk() (**os.File, bool) {
	if o == nil || isNil(o.CsvFile) {
		return nil, false
	}
	return o.CsvFile, true
}

// HasCsvFile returns a boolean if a field has been set.
func (o *ImportEntitlementsBySourceRequest) HasCsvFile() bool {
	if o != nil && !isNil(o.CsvFile) {
		return true
	}

	return false
}

// SetCsvFile gets a reference to the given *os.File and assigns it to the CsvFile field.
func (o *ImportEntitlementsBySourceRequest) SetCsvFile(v *os.File) {
	o.CsvFile = &v
}

func (o ImportEntitlementsBySourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportEntitlementsBySourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CsvFile) {
		toSerialize["csvFile"] = o.CsvFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportEntitlementsBySourceRequest) UnmarshalJSON(bytes []byte) (err error) {
	varImportEntitlementsBySourceRequest := _ImportEntitlementsBySourceRequest{}

	if err = json.Unmarshal(bytes, &varImportEntitlementsBySourceRequest); err == nil {
	*o = ImportEntitlementsBySourceRequest(varImportEntitlementsBySourceRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csvFile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportEntitlementsBySourceRequest struct {
	value *ImportEntitlementsBySourceRequest
	isSet bool
}

func (v NullableImportEntitlementsBySourceRequest) Get() *ImportEntitlementsBySourceRequest {
	return v.value
}

func (v *NullableImportEntitlementsBySourceRequest) Set(val *ImportEntitlementsBySourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportEntitlementsBySourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportEntitlementsBySourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportEntitlementsBySourceRequest(val *ImportEntitlementsBySourceRequest) *NullableImportEntitlementsBySourceRequest {
	return &NullableImportEntitlementsBySourceRequest{value: val, isSet: true}
}

func (v NullableImportEntitlementsBySourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportEntitlementsBySourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


