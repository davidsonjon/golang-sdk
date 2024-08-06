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

// checks if the AggregationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationResult{}

// AggregationResult struct for AggregationResult
type AggregationResult struct {
	// The document containing the results of the aggregation. This document is controlled by Elasticsearch and depends on the type of aggregation query that is run.  See Elasticsearch [Aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/search-aggregations.html) documentation for information. 
	Aggregations map[string]interface{} `json:"aggregations,omitempty"`
	// The results of the aggregation search query. 
	Hits []map[string]interface{} `json:"hits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregationResult AggregationResult

// NewAggregationResult instantiates a new AggregationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationResult() *AggregationResult {
	this := AggregationResult{}
	return &this
}

// NewAggregationResultWithDefaults instantiates a new AggregationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationResultWithDefaults() *AggregationResult {
	this := AggregationResult{}
	return &this
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *AggregationResult) GetAggregations() map[string]interface{} {
	if o == nil || isNil(o.Aggregations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationResult) GetAggregationsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Aggregations) {
		return map[string]interface{}{}, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *AggregationResult) HasAggregations() bool {
	if o != nil && !isNil(o.Aggregations) {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given map[string]interface{} and assigns it to the Aggregations field.
func (o *AggregationResult) SetAggregations(v map[string]interface{}) {
	o.Aggregations = v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *AggregationResult) GetHits() []map[string]interface{} {
	if o == nil || isNil(o.Hits) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationResult) GetHitsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Hits) {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *AggregationResult) HasHits() bool {
	if o != nil && !isNil(o.Hits) {
		return true
	}

	return false
}

// SetHits gets a reference to the given []map[string]interface{} and assigns it to the Hits field.
func (o *AggregationResult) SetHits(v []map[string]interface{}) {
	o.Hits = v
}

func (o AggregationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aggregations) {
		toSerialize["aggregations"] = o.Aggregations
	}
	if !isNil(o.Hits) {
		toSerialize["hits"] = o.Hits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregationResult) UnmarshalJSON(bytes []byte) (err error) {
	varAggregationResult := _AggregationResult{}

	if err = json.Unmarshal(bytes, &varAggregationResult); err == nil {
	*o = AggregationResult(varAggregationResult)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "hits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregationResult struct {
	value *AggregationResult
	isSet bool
}

func (v NullableAggregationResult) Get() *AggregationResult {
	return v.value
}

func (v *NullableAggregationResult) Set(val *AggregationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationResult(val *AggregationResult) *NullableAggregationResult {
	return &NullableAggregationResult{value: val, isSet: true}
}

func (v NullableAggregationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


