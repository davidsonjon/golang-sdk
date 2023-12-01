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

// checks if the OutlierFeatureSummaryOutlierFeatureDisplayValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutlierFeatureSummaryOutlierFeatureDisplayValuesInner{}

// OutlierFeatureSummaryOutlierFeatureDisplayValuesInner struct for OutlierFeatureSummaryOutlierFeatureDisplayValuesInner
type OutlierFeatureSummaryOutlierFeatureDisplayValuesInner struct {
	// display name
	DisplayName *string `json:"displayName,omitempty"`
	// value
	Value *string `json:"value,omitempty"`
	// The data type of the value field
	ValueType *string `json:"valueType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutlierFeatureSummaryOutlierFeatureDisplayValuesInner OutlierFeatureSummaryOutlierFeatureDisplayValuesInner

// NewOutlierFeatureSummaryOutlierFeatureDisplayValuesInner instantiates a new OutlierFeatureSummaryOutlierFeatureDisplayValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlierFeatureSummaryOutlierFeatureDisplayValuesInner() *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner {
	this := OutlierFeatureSummaryOutlierFeatureDisplayValuesInner{}
	return &this
}

// NewOutlierFeatureSummaryOutlierFeatureDisplayValuesInnerWithDefaults instantiates a new OutlierFeatureSummaryOutlierFeatureDisplayValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlierFeatureSummaryOutlierFeatureDisplayValuesInnerWithDefaults() *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner {
	this := OutlierFeatureSummaryOutlierFeatureDisplayValuesInner{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) SetValue(v string) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) GetValueType() string {
	if o == nil || isNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) GetValueTypeOk() (*string, bool) {
	if o == nil || isNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) HasValueType() bool {
	if o != nil && !isNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) SetValueType(v string) {
	o.ValueType = &v
}

func (o OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) UnmarshalJSON(bytes []byte) (err error) {
	varOutlierFeatureSummaryOutlierFeatureDisplayValuesInner := _OutlierFeatureSummaryOutlierFeatureDisplayValuesInner{}

	if err = json.Unmarshal(bytes, &varOutlierFeatureSummaryOutlierFeatureDisplayValuesInner); err == nil {
	*o = OutlierFeatureSummaryOutlierFeatureDisplayValuesInner(varOutlierFeatureSummaryOutlierFeatureDisplayValuesInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "value")
		delete(additionalProperties, "valueType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner struct {
	value *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner
	isSet bool
}

func (v NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner) Get() *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner {
	return v.value
}

func (v *NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner) Set(val *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner(val *OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) *NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner {
	return &NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner{value: val, isSet: true}
}

func (v NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlierFeatureSummaryOutlierFeatureDisplayValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


