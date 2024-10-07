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

// checks if the DimensionBulkDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DimensionBulkDeleteRequest{}

// DimensionBulkDeleteRequest struct for DimensionBulkDeleteRequest
type DimensionBulkDeleteRequest struct {
	// List of IDs of Dimensions to be deleted.
	DimensionIds []string `json:"dimensionIds"`
	AdditionalProperties map[string]interface{}
}

type _DimensionBulkDeleteRequest DimensionBulkDeleteRequest

// NewDimensionBulkDeleteRequest instantiates a new DimensionBulkDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionBulkDeleteRequest(dimensionIds []string) *DimensionBulkDeleteRequest {
	this := DimensionBulkDeleteRequest{}
	this.DimensionIds = dimensionIds
	return &this
}

// NewDimensionBulkDeleteRequestWithDefaults instantiates a new DimensionBulkDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionBulkDeleteRequestWithDefaults() *DimensionBulkDeleteRequest {
	this := DimensionBulkDeleteRequest{}
	return &this
}

// GetDimensionIds returns the DimensionIds field value
func (o *DimensionBulkDeleteRequest) GetDimensionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DimensionIds
}

// GetDimensionIdsOk returns a tuple with the DimensionIds field value
// and a boolean to check if the value has been set.
func (o *DimensionBulkDeleteRequest) GetDimensionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DimensionIds, true
}

// SetDimensionIds sets field value
func (o *DimensionBulkDeleteRequest) SetDimensionIds(v []string) {
	o.DimensionIds = v
}

func (o DimensionBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DimensionBulkDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dimensionIds"] = o.DimensionIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DimensionBulkDeleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dimensionIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDimensionBulkDeleteRequest := _DimensionBulkDeleteRequest{}

	err = json.Unmarshal(data, &varDimensionBulkDeleteRequest)

	if err != nil {
		return err
	}

	*o = DimensionBulkDeleteRequest(varDimensionBulkDeleteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dimensionIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDimensionBulkDeleteRequest struct {
	value *DimensionBulkDeleteRequest
	isSet bool
}

func (v NullableDimensionBulkDeleteRequest) Get() *DimensionBulkDeleteRequest {
	return v.value
}

func (v *NullableDimensionBulkDeleteRequest) Set(val *DimensionBulkDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionBulkDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionBulkDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionBulkDeleteRequest(val *DimensionBulkDeleteRequest) *NullableDimensionBulkDeleteRequest {
	return &NullableDimensionBulkDeleteRequest{value: val, isSet: true}
}

func (v NullableDimensionBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionBulkDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


