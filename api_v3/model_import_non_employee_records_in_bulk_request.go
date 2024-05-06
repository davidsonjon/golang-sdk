/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the ImportNonEmployeeRecordsInBulkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportNonEmployeeRecordsInBulkRequest{}

// ImportNonEmployeeRecordsInBulkRequest struct for ImportNonEmployeeRecordsInBulkRequest
type ImportNonEmployeeRecordsInBulkRequest struct {
	Data *os.File `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _ImportNonEmployeeRecordsInBulkRequest ImportNonEmployeeRecordsInBulkRequest

// NewImportNonEmployeeRecordsInBulkRequest instantiates a new ImportNonEmployeeRecordsInBulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportNonEmployeeRecordsInBulkRequest(data *os.File) *ImportNonEmployeeRecordsInBulkRequest {
	this := ImportNonEmployeeRecordsInBulkRequest{}
	this.Data = data
	return &this
}

// NewImportNonEmployeeRecordsInBulkRequestWithDefaults instantiates a new ImportNonEmployeeRecordsInBulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportNonEmployeeRecordsInBulkRequestWithDefaults() *ImportNonEmployeeRecordsInBulkRequest {
	this := ImportNonEmployeeRecordsInBulkRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ImportNonEmployeeRecordsInBulkRequest) GetData() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImportNonEmployeeRecordsInBulkRequest) GetDataOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ImportNonEmployeeRecordsInBulkRequest) SetData(v *os.File) {
	o.Data = v
}

func (o ImportNonEmployeeRecordsInBulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportNonEmployeeRecordsInBulkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportNonEmployeeRecordsInBulkRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varImportNonEmployeeRecordsInBulkRequest := _ImportNonEmployeeRecordsInBulkRequest{}

	if err = json.Unmarshal(bytes, &varImportNonEmployeeRecordsInBulkRequest); err == nil {
	*o = ImportNonEmployeeRecordsInBulkRequest(varImportNonEmployeeRecordsInBulkRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportNonEmployeeRecordsInBulkRequest struct {
	value *ImportNonEmployeeRecordsInBulkRequest
	isSet bool
}

func (v NullableImportNonEmployeeRecordsInBulkRequest) Get() *ImportNonEmployeeRecordsInBulkRequest {
	return v.value
}

func (v *NullableImportNonEmployeeRecordsInBulkRequest) Set(val *ImportNonEmployeeRecordsInBulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportNonEmployeeRecordsInBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportNonEmployeeRecordsInBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportNonEmployeeRecordsInBulkRequest(val *ImportNonEmployeeRecordsInBulkRequest) *NullableImportNonEmployeeRecordsInBulkRequest {
	return &NullableImportNonEmployeeRecordsInBulkRequest{value: val, isSet: true}
}

func (v NullableImportNonEmployeeRecordsInBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportNonEmployeeRecordsInBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


