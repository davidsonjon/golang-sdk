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

// checks if the UploadNonEmployeeRecordsInBulkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadNonEmployeeRecordsInBulkRequest{}

// UploadNonEmployeeRecordsInBulkRequest struct for UploadNonEmployeeRecordsInBulkRequest
type UploadNonEmployeeRecordsInBulkRequest struct {
	Data string `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _UploadNonEmployeeRecordsInBulkRequest UploadNonEmployeeRecordsInBulkRequest

// NewUploadNonEmployeeRecordsInBulkRequest instantiates a new UploadNonEmployeeRecordsInBulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadNonEmployeeRecordsInBulkRequest(data string) *UploadNonEmployeeRecordsInBulkRequest {
	this := UploadNonEmployeeRecordsInBulkRequest{}
	this.Data = data
	return &this
}

// NewUploadNonEmployeeRecordsInBulkRequestWithDefaults instantiates a new UploadNonEmployeeRecordsInBulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadNonEmployeeRecordsInBulkRequestWithDefaults() *UploadNonEmployeeRecordsInBulkRequest {
	this := UploadNonEmployeeRecordsInBulkRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UploadNonEmployeeRecordsInBulkRequest) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UploadNonEmployeeRecordsInBulkRequest) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UploadNonEmployeeRecordsInBulkRequest) SetData(v string) {
	o.Data = v
}

func (o UploadNonEmployeeRecordsInBulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadNonEmployeeRecordsInBulkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UploadNonEmployeeRecordsInBulkRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUploadNonEmployeeRecordsInBulkRequest := _UploadNonEmployeeRecordsInBulkRequest{}

	if err = json.Unmarshal(bytes, &varUploadNonEmployeeRecordsInBulkRequest); err == nil {
		*o = UploadNonEmployeeRecordsInBulkRequest(varUploadNonEmployeeRecordsInBulkRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadNonEmployeeRecordsInBulkRequest struct {
	value *UploadNonEmployeeRecordsInBulkRequest
	isSet bool
}

func (v NullableUploadNonEmployeeRecordsInBulkRequest) Get() *UploadNonEmployeeRecordsInBulkRequest {
	return v.value
}

func (v *NullableUploadNonEmployeeRecordsInBulkRequest) Set(val *UploadNonEmployeeRecordsInBulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadNonEmployeeRecordsInBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadNonEmployeeRecordsInBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadNonEmployeeRecordsInBulkRequest(val *UploadNonEmployeeRecordsInBulkRequest) *NullableUploadNonEmployeeRecordsInBulkRequest {
	return &NullableUploadNonEmployeeRecordsInBulkRequest{value: val, isSet: true}
}

func (v NullableUploadNonEmployeeRecordsInBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadNonEmployeeRecordsInBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


