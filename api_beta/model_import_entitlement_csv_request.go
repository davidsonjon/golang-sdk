/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the ImportEntitlementCsvRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportEntitlementCsvRequest{}

// ImportEntitlementCsvRequest struct for ImportEntitlementCsvRequest
type ImportEntitlementCsvRequest struct {
	CsvFile *os.File `json:"csvFile"`
	AdditionalProperties map[string]interface{}
}

type _ImportEntitlementCsvRequest ImportEntitlementCsvRequest

// NewImportEntitlementCsvRequest instantiates a new ImportEntitlementCsvRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportEntitlementCsvRequest(csvFile *os.File) *ImportEntitlementCsvRequest {
	this := ImportEntitlementCsvRequest{}
	this.CsvFile = csvFile
	return &this
}

// NewImportEntitlementCsvRequestWithDefaults instantiates a new ImportEntitlementCsvRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportEntitlementCsvRequestWithDefaults() *ImportEntitlementCsvRequest {
	this := ImportEntitlementCsvRequest{}
	return &this
}

// GetCsvFile returns the CsvFile field value
func (o *ImportEntitlementCsvRequest) GetCsvFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.CsvFile
}

// GetCsvFileOk returns a tuple with the CsvFile field value
// and a boolean to check if the value has been set.
func (o *ImportEntitlementCsvRequest) GetCsvFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsvFile, true
}

// SetCsvFile sets field value
func (o *ImportEntitlementCsvRequest) SetCsvFile(v *os.File) {
	o.CsvFile = v
}

func (o ImportEntitlementCsvRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportEntitlementCsvRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["csvFile"] = o.CsvFile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportEntitlementCsvRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"csvFile",
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

	varImportEntitlementCsvRequest := _ImportEntitlementCsvRequest{}

	if err = json.Unmarshal(bytes, &varImportEntitlementCsvRequest); err == nil {
	*o = ImportEntitlementCsvRequest(varImportEntitlementCsvRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csvFile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportEntitlementCsvRequest struct {
	value *ImportEntitlementCsvRequest
	isSet bool
}

func (v NullableImportEntitlementCsvRequest) Get() *ImportEntitlementCsvRequest {
	return v.value
}

func (v *NullableImportEntitlementCsvRequest) Set(val *ImportEntitlementCsvRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportEntitlementCsvRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportEntitlementCsvRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportEntitlementCsvRequest(val *ImportEntitlementCsvRequest) *NullableImportEntitlementCsvRequest {
	return &NullableImportEntitlementCsvRequest{value: val, isSet: true}
}

func (v NullableImportEntitlementCsvRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportEntitlementCsvRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


