/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the MetricAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricAggregation{}

// MetricAggregation The calculation done on the results of the query
type MetricAggregation struct {
	// The name of the metric aggregate to be included in the result. If the metric aggregation is omitted, the resulting aggregation will be a count of the documents in the search results.
	Name string `json:"name"`
	Type *MetricType `json:"type,omitempty"`
	// The field the calculation is performed on.  Prefix the field name with '@' to reference a nested object. 
	Field string `json:"field"`
	AdditionalProperties map[string]interface{}
}

type _MetricAggregation MetricAggregation

// NewMetricAggregation instantiates a new MetricAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricAggregation(name string, field string) *MetricAggregation {
	this := MetricAggregation{}
	this.Name = name
	var type_ MetricType = METRICTYPE_UNIQUE_COUNT
	this.Type = &type_
	this.Field = field
	return &this
}

// NewMetricAggregationWithDefaults instantiates a new MetricAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricAggregationWithDefaults() *MetricAggregation {
	this := MetricAggregation{}
	var type_ MetricType = METRICTYPE_UNIQUE_COUNT
	this.Type = &type_
	return &this
}

// GetName returns the Name field value
func (o *MetricAggregation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricAggregation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricAggregation) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricAggregation) GetType() MetricType {
	if o == nil || isNil(o.Type) {
		var ret MetricType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAggregation) GetTypeOk() (*MetricType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricAggregation) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MetricType and assigns it to the Type field.
func (o *MetricAggregation) SetType(v MetricType) {
	o.Type = &v
}

// GetField returns the Field field value
func (o *MetricAggregation) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *MetricAggregation) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *MetricAggregation) SetField(v string) {
	o.Field = v
}

func (o MetricAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["field"] = o.Field

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricAggregation) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"field",
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

	varMetricAggregation := _MetricAggregation{}

	if err = json.Unmarshal(bytes, &varMetricAggregation); err == nil {
	*o = MetricAggregation(varMetricAggregation)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "field")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricAggregation struct {
	value *MetricAggregation
	isSet bool
}

func (v NullableMetricAggregation) Get() *MetricAggregation {
	return v.value
}

func (v *NullableMetricAggregation) Set(val *MetricAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricAggregation(val *MetricAggregation) *NullableMetricAggregation {
	return &NullableMetricAggregation{value: val, isSet: true}
}

func (v NullableMetricAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


