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

// checks if the FilterAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterAggregation{}

// FilterAggregation An additional filter to constrain the results of the search query.
type FilterAggregation struct {
	// The name of the filter aggregate to be included in the result.
	Name string `json:"name"`
	Type *SearchFilterType `json:"type,omitempty"`
	// The search field to apply the filter to.  Prefix the field name with '@' to reference a nested object. 
	Field string `json:"field"`
	// The value to filter on.
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FilterAggregation FilterAggregation

// NewFilterAggregation instantiates a new FilterAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterAggregation(name string, field string, value string) *FilterAggregation {
	this := FilterAggregation{}
	this.Name = name
	var type_ SearchFilterType = SEARCHFILTERTYPE_TERM
	this.Type = &type_
	this.Field = field
	this.Value = value
	return &this
}

// NewFilterAggregationWithDefaults instantiates a new FilterAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterAggregationWithDefaults() *FilterAggregation {
	this := FilterAggregation{}
	var type_ SearchFilterType = SEARCHFILTERTYPE_TERM
	this.Type = &type_
	return &this
}

// GetName returns the Name field value
func (o *FilterAggregation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilterAggregation) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FilterAggregation) GetType() SearchFilterType {
	if o == nil || isNil(o.Type) {
		var ret SearchFilterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetTypeOk() (*SearchFilterType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FilterAggregation) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SearchFilterType and assigns it to the Type field.
func (o *FilterAggregation) SetType(v SearchFilterType) {
	o.Type = &v
}

// GetField returns the Field field value
func (o *FilterAggregation) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *FilterAggregation) SetField(v string) {
	o.Field = v
}

// GetValue returns the Value field value
func (o *FilterAggregation) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilterAggregation) SetValue(v string) {
	o.Value = v
}

func (o FilterAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["field"] = o.Field
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilterAggregation) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"field",
		"value",
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

	varFilterAggregation := _FilterAggregation{}

	if err = json.Unmarshal(bytes, &varFilterAggregation); err == nil {
	*o = FilterAggregation(varFilterAggregation)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "field")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilterAggregation struct {
	value *FilterAggregation
	isSet bool
}

func (v NullableFilterAggregation) Get() *FilterAggregation {
	return v.value
}

func (v *NullableFilterAggregation) Set(val *FilterAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterAggregation(val *FilterAggregation) *NullableFilterAggregation {
	return &NullableFilterAggregation{value: val, isSet: true}
}

func (v NullableFilterAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


