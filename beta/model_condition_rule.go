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

// checks if the ConditionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionRule{}

// ConditionRule struct for ConditionRule
type ConditionRule struct {
	// Defines the type of object being selected. It will be either a reference to a form input (by input name) or a form element (by technical key). INPUT ConditionRuleSourceTypeInput ELEMENT ConditionRuleSourceTypeElement
	SourceType *string `json:"sourceType,omitempty"`
	// Source - if the sourceType is ConditionRuleSourceTypeInput, the source type is the name of the form input to accept. However, if the sourceType is ConditionRuleSourceTypeElement, the source is the name of a technical key of an element to retrieve its value.
	Source *string `json:"source,omitempty"`
	// ConditionRuleComparisonOperatorType value. EQ ConditionRuleComparisonOperatorTypeEquals  This comparison operator compares the source and target for equality. NE ConditionRuleComparisonOperatorTypeNotEquals  This comparison operator compares the source and target for inequality. CO ConditionRuleComparisonOperatorTypeContains  This comparison operator searches the source to see whether it contains the value. NOT_CO ConditionRuleComparisonOperatorTypeNotContains IN ConditionRuleComparisonOperatorTypeIncludes  This comparison operator searches the source if it equals any of the values. NOT_IN ConditionRuleComparisonOperatorTypeNotIncludes EM ConditionRuleComparisonOperatorTypeEmpty NOT_EM ConditionRuleComparisonOperatorTypeNotEmpty SW ConditionRuleComparisonOperatorTypeStartsWith  Checks whether a string starts with another substring of the same string. This operator is case-sensitive. NOT_SW ConditionRuleComparisonOperatorTypeNotStartsWith EW ConditionRuleComparisonOperatorTypeEndsWith  Checks whether a string ends with another substring of the same string. This operator is case-sensitive. NOT_EW ConditionRuleComparisonOperatorTypeNotEndsWith
	Operator *string `json:"operator,omitempty"`
	// ConditionRuleValueType type. STRING ConditionRuleValueTypeString  This value is a static string. STRING_LIST ConditionRuleValueTypeStringList  This value is an array of string values. INPUT ConditionRuleValueTypeInput  This value is a reference to a form input. ELEMENT ConditionRuleValueTypeElement  This value is a reference to a form element (by technical key). LIST ConditionRuleValueTypeList BOOLEAN ConditionRuleValueTypeBoolean
	ValueType *string `json:"valueType,omitempty"`
	// Based on the ValueType.
	Value map[string]interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionRule ConditionRule

// NewConditionRule instantiates a new ConditionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionRule() *ConditionRule {
	this := ConditionRule{}
	return &this
}

// NewConditionRuleWithDefaults instantiates a new ConditionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionRuleWithDefaults() *ConditionRule {
	this := ConditionRule{}
	return &this
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *ConditionRule) GetSourceType() string {
	if o == nil || isNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRule) GetSourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *ConditionRule) HasSourceType() bool {
	if o != nil && !isNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *ConditionRule) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConditionRule) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRule) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConditionRule) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConditionRule) SetSource(v string) {
	o.Source = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ConditionRule) GetOperator() string {
	if o == nil || isNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRule) GetOperatorOk() (*string, bool) {
	if o == nil || isNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ConditionRule) HasOperator() bool {
	if o != nil && !isNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ConditionRule) SetOperator(v string) {
	o.Operator = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ConditionRule) GetValueType() string {
	if o == nil || isNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRule) GetValueTypeOk() (*string, bool) {
	if o == nil || isNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ConditionRule) HasValueType() bool {
	if o != nil && !isNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ConditionRule) SetValueType(v string) {
	o.ValueType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConditionRule) GetValue() map[string]interface{} {
	if o == nil || isNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRule) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConditionRule) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *ConditionRule) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o ConditionRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !isNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionRule) UnmarshalJSON(bytes []byte) (err error) {
	varConditionRule := _ConditionRule{}

	if err = json.Unmarshal(bytes, &varConditionRule); err == nil {
		*o = ConditionRule(varConditionRule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sourceType")
		delete(additionalProperties, "source")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "valueType")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionRule struct {
	value *ConditionRule
	isSet bool
}

func (v NullableConditionRule) Get() *ConditionRule {
	return v.value
}

func (v *NullableConditionRule) Set(val *ConditionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionRule(val *ConditionRule) *NullableConditionRule {
	return &NullableConditionRule{value: val, isSet: true}
}

func (v NullableConditionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


