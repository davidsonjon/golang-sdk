/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"os"
)

// checks if the ImportAccountsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportAccountsRequest{}

// ImportAccountsRequest struct for ImportAccountsRequest
type ImportAccountsRequest struct {
	// The CSV file containing the source accounts to aggregate.
	File **os.File `json:"file,omitempty"`
	// Use this flag to reprocess every account whether or not the data has changed.
	DisableOptimization *string `json:"disableOptimization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportAccountsRequest ImportAccountsRequest

// NewImportAccountsRequest instantiates a new ImportAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportAccountsRequest() *ImportAccountsRequest {
	this := ImportAccountsRequest{}
	return &this
}

// NewImportAccountsRequestWithDefaults instantiates a new ImportAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportAccountsRequestWithDefaults() *ImportAccountsRequest {
	this := ImportAccountsRequest{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ImportAccountsRequest) GetFile() *os.File {
	if o == nil || isNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportAccountsRequest) GetFileOk() (**os.File, bool) {
	if o == nil || isNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ImportAccountsRequest) HasFile() bool {
	if o != nil && !isNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *ImportAccountsRequest) SetFile(v *os.File) {
	o.File = &v
}

// GetDisableOptimization returns the DisableOptimization field value if set, zero value otherwise.
func (o *ImportAccountsRequest) GetDisableOptimization() string {
	if o == nil || isNil(o.DisableOptimization) {
		var ret string
		return ret
	}
	return *o.DisableOptimization
}

// GetDisableOptimizationOk returns a tuple with the DisableOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportAccountsRequest) GetDisableOptimizationOk() (*string, bool) {
	if o == nil || isNil(o.DisableOptimization) {
		return nil, false
	}
	return o.DisableOptimization, true
}

// HasDisableOptimization returns a boolean if a field has been set.
func (o *ImportAccountsRequest) HasDisableOptimization() bool {
	if o != nil && !isNil(o.DisableOptimization) {
		return true
	}

	return false
}

// SetDisableOptimization gets a reference to the given string and assigns it to the DisableOptimization field.
func (o *ImportAccountsRequest) SetDisableOptimization(v string) {
	o.DisableOptimization = &v
}

func (o ImportAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportAccountsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !isNil(o.DisableOptimization) {
		toSerialize["disableOptimization"] = o.DisableOptimization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportAccountsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varImportAccountsRequest := _ImportAccountsRequest{}

	if err = json.Unmarshal(bytes, &varImportAccountsRequest); err == nil {
	*o = ImportAccountsRequest(varImportAccountsRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "disableOptimization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportAccountsRequest struct {
	value *ImportAccountsRequest
	isSet bool
}

func (v NullableImportAccountsRequest) Get() *ImportAccountsRequest {
	return v.value
}

func (v *NullableImportAccountsRequest) Set(val *ImportAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportAccountsRequest(val *ImportAccountsRequest) *NullableImportAccountsRequest {
	return &NullableImportAccountsRequest{value: val, isSet: true}
}

func (v NullableImportAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


