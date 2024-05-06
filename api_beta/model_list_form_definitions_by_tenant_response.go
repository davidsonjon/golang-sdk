/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the ListFormDefinitionsByTenantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFormDefinitionsByTenantResponse{}

// ListFormDefinitionsByTenantResponse struct for ListFormDefinitionsByTenantResponse
type ListFormDefinitionsByTenantResponse struct {
	// Count number of results.
	Count *int64 `json:"count,omitempty"`
	// List of FormDefinitionResponse items.
	Results []FormDefinitionResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListFormDefinitionsByTenantResponse ListFormDefinitionsByTenantResponse

// NewListFormDefinitionsByTenantResponse instantiates a new ListFormDefinitionsByTenantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFormDefinitionsByTenantResponse() *ListFormDefinitionsByTenantResponse {
	this := ListFormDefinitionsByTenantResponse{}
	return &this
}

// NewListFormDefinitionsByTenantResponseWithDefaults instantiates a new ListFormDefinitionsByTenantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFormDefinitionsByTenantResponseWithDefaults() *ListFormDefinitionsByTenantResponse {
	this := ListFormDefinitionsByTenantResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListFormDefinitionsByTenantResponse) GetCount() int64 {
	if o == nil || isNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFormDefinitionsByTenantResponse) GetCountOk() (*int64, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListFormDefinitionsByTenantResponse) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListFormDefinitionsByTenantResponse) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListFormDefinitionsByTenantResponse) GetResults() []FormDefinitionResponse {
	if o == nil || isNil(o.Results) {
		var ret []FormDefinitionResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFormDefinitionsByTenantResponse) GetResultsOk() ([]FormDefinitionResponse, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListFormDefinitionsByTenantResponse) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []FormDefinitionResponse and assigns it to the Results field.
func (o *ListFormDefinitionsByTenantResponse) SetResults(v []FormDefinitionResponse) {
	o.Results = v
}

func (o ListFormDefinitionsByTenantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFormDefinitionsByTenantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListFormDefinitionsByTenantResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListFormDefinitionsByTenantResponse := _ListFormDefinitionsByTenantResponse{}

	if err = json.Unmarshal(bytes, &varListFormDefinitionsByTenantResponse); err == nil {
	*o = ListFormDefinitionsByTenantResponse(varListFormDefinitionsByTenantResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListFormDefinitionsByTenantResponse struct {
	value *ListFormDefinitionsByTenantResponse
	isSet bool
}

func (v NullableListFormDefinitionsByTenantResponse) Get() *ListFormDefinitionsByTenantResponse {
	return v.value
}

func (v *NullableListFormDefinitionsByTenantResponse) Set(val *ListFormDefinitionsByTenantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFormDefinitionsByTenantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFormDefinitionsByTenantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFormDefinitionsByTenantResponse(val *ListFormDefinitionsByTenantResponse) *NullableListFormDefinitionsByTenantResponse {
	return &NullableListFormDefinitionsByTenantResponse{value: val, isSet: true}
}

func (v NullableListFormDefinitionsByTenantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFormDefinitionsByTenantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


