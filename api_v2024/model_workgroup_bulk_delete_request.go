/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the WorkgroupBulkDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupBulkDeleteRequest{}

// WorkgroupBulkDeleteRequest struct for WorkgroupBulkDeleteRequest
type WorkgroupBulkDeleteRequest struct {
	// List of IDs of Governance Groups to be deleted.
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupBulkDeleteRequest WorkgroupBulkDeleteRequest

// NewWorkgroupBulkDeleteRequest instantiates a new WorkgroupBulkDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupBulkDeleteRequest() *WorkgroupBulkDeleteRequest {
	this := WorkgroupBulkDeleteRequest{}
	return &this
}

// NewWorkgroupBulkDeleteRequestWithDefaults instantiates a new WorkgroupBulkDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupBulkDeleteRequestWithDefaults() *WorkgroupBulkDeleteRequest {
	this := WorkgroupBulkDeleteRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *WorkgroupBulkDeleteRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupBulkDeleteRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *WorkgroupBulkDeleteRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *WorkgroupBulkDeleteRequest) SetIds(v []string) {
	o.Ids = v
}

func (o WorkgroupBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupBulkDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupBulkDeleteRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWorkgroupBulkDeleteRequest := _WorkgroupBulkDeleteRequest{}

	if err = json.Unmarshal(bytes, &varWorkgroupBulkDeleteRequest); err == nil {
	*o = WorkgroupBulkDeleteRequest(varWorkgroupBulkDeleteRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupBulkDeleteRequest struct {
	value *WorkgroupBulkDeleteRequest
	isSet bool
}

func (v NullableWorkgroupBulkDeleteRequest) Get() *WorkgroupBulkDeleteRequest {
	return v.value
}

func (v *NullableWorkgroupBulkDeleteRequest) Set(val *WorkgroupBulkDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupBulkDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupBulkDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupBulkDeleteRequest(val *WorkgroupBulkDeleteRequest) *NullableWorkgroupBulkDeleteRequest {
	return &NullableWorkgroupBulkDeleteRequest{value: val, isSet: true}
}

func (v NullableWorkgroupBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupBulkDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


