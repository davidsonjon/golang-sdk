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

// checks if the ListFormElementDataByElementIDResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFormElementDataByElementIDResponse{}

// ListFormElementDataByElementIDResponse struct for ListFormElementDataByElementIDResponse
type ListFormElementDataByElementIDResponse struct {
	// Results holds a list of FormElementDataSourceConfigOptions items
	Results []FormElementDataSourceConfigOptions `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListFormElementDataByElementIDResponse ListFormElementDataByElementIDResponse

// NewListFormElementDataByElementIDResponse instantiates a new ListFormElementDataByElementIDResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFormElementDataByElementIDResponse() *ListFormElementDataByElementIDResponse {
	this := ListFormElementDataByElementIDResponse{}
	return &this
}

// NewListFormElementDataByElementIDResponseWithDefaults instantiates a new ListFormElementDataByElementIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFormElementDataByElementIDResponseWithDefaults() *ListFormElementDataByElementIDResponse {
	this := ListFormElementDataByElementIDResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListFormElementDataByElementIDResponse) GetResults() []FormElementDataSourceConfigOptions {
	if o == nil || isNil(o.Results) {
		var ret []FormElementDataSourceConfigOptions
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFormElementDataByElementIDResponse) GetResultsOk() ([]FormElementDataSourceConfigOptions, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListFormElementDataByElementIDResponse) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []FormElementDataSourceConfigOptions and assigns it to the Results field.
func (o *ListFormElementDataByElementIDResponse) SetResults(v []FormElementDataSourceConfigOptions) {
	o.Results = v
}

func (o ListFormElementDataByElementIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFormElementDataByElementIDResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListFormElementDataByElementIDResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListFormElementDataByElementIDResponse := _ListFormElementDataByElementIDResponse{}

	if err = json.Unmarshal(bytes, &varListFormElementDataByElementIDResponse); err == nil {
	*o = ListFormElementDataByElementIDResponse(varListFormElementDataByElementIDResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListFormElementDataByElementIDResponse struct {
	value *ListFormElementDataByElementIDResponse
	isSet bool
}

func (v NullableListFormElementDataByElementIDResponse) Get() *ListFormElementDataByElementIDResponse {
	return v.value
}

func (v *NullableListFormElementDataByElementIDResponse) Set(val *ListFormElementDataByElementIDResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFormElementDataByElementIDResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFormElementDataByElementIDResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFormElementDataByElementIDResponse(val *ListFormElementDataByElementIDResponse) *NullableListFormElementDataByElementIDResponse {
	return &NullableListFormElementDataByElementIDResponse{value: val, isSet: true}
}

func (v NullableListFormElementDataByElementIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFormElementDataByElementIDResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


