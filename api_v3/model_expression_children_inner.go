/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the ExpressionChildrenInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpressionChildrenInner{}

// ExpressionChildrenInner struct for ExpressionChildrenInner
type ExpressionChildrenInner struct {
	// Operator for the expression
	Operator *string `json:"operator,omitempty"`
	// Name for the attribute
	Attribute NullableString `json:"attribute,omitempty"`
	Value NullableValue `json:"value,omitempty"`
	// There cannot be anymore nested children. This will always be null.
	Children NullableString `json:"children,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExpressionChildrenInner ExpressionChildrenInner

// NewExpressionChildrenInner instantiates a new ExpressionChildrenInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressionChildrenInner() *ExpressionChildrenInner {
	this := ExpressionChildrenInner{}
	return &this
}

// NewExpressionChildrenInnerWithDefaults instantiates a new ExpressionChildrenInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionChildrenInnerWithDefaults() *ExpressionChildrenInner {
	this := ExpressionChildrenInner{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ExpressionChildrenInner) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionChildrenInner) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ExpressionChildrenInner) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ExpressionChildrenInner) SetOperator(v string) {
	o.Operator = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpressionChildrenInner) GetAttribute() string {
	if o == nil || IsNil(o.Attribute.Get()) {
		var ret string
		return ret
	}
	return *o.Attribute.Get()
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpressionChildrenInner) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute.Get(), o.Attribute.IsSet()
}

// HasAttribute returns a boolean if a field has been set.
func (o *ExpressionChildrenInner) HasAttribute() bool {
	if o != nil && o.Attribute.IsSet() {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given NullableString and assigns it to the Attribute field.
func (o *ExpressionChildrenInner) SetAttribute(v string) {
	o.Attribute.Set(&v)
}
// SetAttributeNil sets the value for Attribute to be an explicit nil
func (o *ExpressionChildrenInner) SetAttributeNil() {
	o.Attribute.Set(nil)
}

// UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
func (o *ExpressionChildrenInner) UnsetAttribute() {
	o.Attribute.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpressionChildrenInner) GetValue() Value {
	if o == nil || IsNil(o.Value.Get()) {
		var ret Value
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpressionChildrenInner) GetValueOk() (*Value, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ExpressionChildrenInner) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableValue and assigns it to the Value field.
func (o *ExpressionChildrenInner) SetValue(v Value) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ExpressionChildrenInner) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ExpressionChildrenInner) UnsetValue() {
	o.Value.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpressionChildrenInner) GetChildren() string {
	if o == nil || IsNil(o.Children.Get()) {
		var ret string
		return ret
	}
	return *o.Children.Get()
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpressionChildrenInner) GetChildrenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children.Get(), o.Children.IsSet()
}

// HasChildren returns a boolean if a field has been set.
func (o *ExpressionChildrenInner) HasChildren() bool {
	if o != nil && o.Children.IsSet() {
		return true
	}

	return false
}

// SetChildren gets a reference to the given NullableString and assigns it to the Children field.
func (o *ExpressionChildrenInner) SetChildren(v string) {
	o.Children.Set(&v)
}
// SetChildrenNil sets the value for Children to be an explicit nil
func (o *ExpressionChildrenInner) SetChildrenNil() {
	o.Children.Set(nil)
}

// UnsetChildren ensures that no value is present for Children, not even an explicit nil
func (o *ExpressionChildrenInner) UnsetChildren() {
	o.Children.Unset()
}

func (o ExpressionChildrenInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpressionChildrenInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if o.Attribute.IsSet() {
		toSerialize["attribute"] = o.Attribute.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Children.IsSet() {
		toSerialize["children"] = o.Children.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExpressionChildrenInner) UnmarshalJSON(data []byte) (err error) {
	varExpressionChildrenInner := _ExpressionChildrenInner{}

	err = json.Unmarshal(data, &varExpressionChildrenInner)

	if err != nil {
		return err
	}

	*o = ExpressionChildrenInner(varExpressionChildrenInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "children")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpressionChildrenInner struct {
	value *ExpressionChildrenInner
	isSet bool
}

func (v NullableExpressionChildrenInner) Get() *ExpressionChildrenInner {
	return v.value
}

func (v *NullableExpressionChildrenInner) Set(val *ExpressionChildrenInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionChildrenInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionChildrenInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionChildrenInner(val *ExpressionChildrenInner) *NullableExpressionChildrenInner {
	return &NullableExpressionChildrenInner{value: val, isSet: true}
}

func (v NullableExpressionChildrenInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionChildrenInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


