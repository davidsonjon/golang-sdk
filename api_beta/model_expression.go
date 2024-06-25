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

// checks if the Expression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Expression{}

// Expression struct for Expression
type Expression struct {
	// Operator for the expression
	Operator *string `json:"operator,omitempty"`
	// Name for the attribute
	Attribute NullableString `json:"attribute,omitempty"`
	Value NullableValue `json:"value,omitempty"`
	// List of expressions
	Children []Children `json:"children,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Expression Expression

// NewExpression instantiates a new Expression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpression() *Expression {
	this := Expression{}
	return &this
}

// NewExpressionWithDefaults instantiates a new Expression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionWithDefaults() *Expression {
	this := Expression{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Expression) GetOperator() string {
	if o == nil || isNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expression) GetOperatorOk() (*string, bool) {
	if o == nil || isNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Expression) HasOperator() bool {
	if o != nil && !isNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *Expression) SetOperator(v string) {
	o.Operator = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Expression) GetAttribute() string {
	if o == nil || isNil(o.Attribute.Get()) {
		var ret string
		return ret
	}
	return *o.Attribute.Get()
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Expression) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute.Get(), o.Attribute.IsSet()
}

// HasAttribute returns a boolean if a field has been set.
func (o *Expression) HasAttribute() bool {
	if o != nil && o.Attribute.IsSet() {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given NullableString and assigns it to the Attribute field.
func (o *Expression) SetAttribute(v string) {
	o.Attribute.Set(&v)
}
// SetAttributeNil sets the value for Attribute to be an explicit nil
func (o *Expression) SetAttributeNil() {
	o.Attribute.Set(nil)
}

// UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
func (o *Expression) UnsetAttribute() {
	o.Attribute.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Expression) GetValue() Value {
	if o == nil || isNil(o.Value.Get()) {
		var ret Value
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Expression) GetValueOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *Expression) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableValue and assigns it to the Value field.
func (o *Expression) SetValue(v Value) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *Expression) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *Expression) UnsetValue() {
	o.Value.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Expression) GetChildren() []Children {
	if o == nil {
		var ret []Children
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Expression) GetChildrenOk() ([]Children, bool) {
	if o == nil || isNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Expression) HasChildren() bool {
	if o != nil && isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Children and assigns it to the Children field.
func (o *Expression) SetChildren(v []Children) {
	o.Children = v
}

func (o Expression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Expression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if o.Attribute.IsSet() {
		toSerialize["attribute"] = o.Attribute.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Expression) UnmarshalJSON(bytes []byte) (err error) {
	varExpression := _Expression{}

	if err = json.Unmarshal(bytes, &varExpression); err == nil {
	*o = Expression(varExpression)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "children")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpression struct {
	value *Expression
	isSet bool
}

func (v NullableExpression) Get() *Expression {
	return v.value
}

func (v *NullableExpression) Set(val *Expression) {
	v.value = val
	v.isSet = true
}

func (v NullableExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpression(val *Expression) *NullableExpression {
	return &NullableExpression{value: val, isSet: true}
}

func (v NullableExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


