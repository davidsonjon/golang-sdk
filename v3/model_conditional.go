/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the Conditional type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conditional{}

// Conditional struct for Conditional
type Conditional struct {
	// A comparison statement that follows the structure of `ValueA eq ValueB` where `ValueA` and `ValueB` are static strings or outputs of other transforms.   The `eq` operator is the only valid comparison
	Expression string `json:"expression"`
	// The output of the transform if the expression evalutes to true
	PositiveCondition string `json:"positiveCondition"`
	// The output of the transform if the expression evalutes to false
	NegativeCondition string `json:"negativeCondition"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Conditional Conditional

// NewConditional instantiates a new Conditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditional(expression string, positiveCondition string, negativeCondition string) *Conditional {
	this := Conditional{}
	this.Expression = expression
	this.PositiveCondition = positiveCondition
	this.NegativeCondition = negativeCondition
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewConditionalWithDefaults instantiates a new Conditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalWithDefaults() *Conditional {
	this := Conditional{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetExpression returns the Expression field value
func (o *Conditional) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *Conditional) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *Conditional) SetExpression(v string) {
	o.Expression = v
}

// GetPositiveCondition returns the PositiveCondition field value
func (o *Conditional) GetPositiveCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PositiveCondition
}

// GetPositiveConditionOk returns a tuple with the PositiveCondition field value
// and a boolean to check if the value has been set.
func (o *Conditional) GetPositiveConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositiveCondition, true
}

// SetPositiveCondition sets field value
func (o *Conditional) SetPositiveCondition(v string) {
	o.PositiveCondition = v
}

// GetNegativeCondition returns the NegativeCondition field value
func (o *Conditional) GetNegativeCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NegativeCondition
}

// GetNegativeConditionOk returns a tuple with the NegativeCondition field value
// and a boolean to check if the value has been set.
func (o *Conditional) GetNegativeConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCondition, true
}

// SetNegativeCondition sets field value
func (o *Conditional) SetNegativeCondition(v string) {
	o.NegativeCondition = v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *Conditional) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conditional) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *Conditional) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *Conditional) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *Conditional) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conditional) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *Conditional) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *Conditional) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o Conditional) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conditional) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	toSerialize["positiveCondition"] = o.PositiveCondition
	toSerialize["negativeCondition"] = o.NegativeCondition
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Conditional) UnmarshalJSON(bytes []byte) (err error) {
	varConditional := _Conditional{}

	if err = json.Unmarshal(bytes, &varConditional); err == nil {
		*o = Conditional(varConditional)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "positiveCondition")
		delete(additionalProperties, "negativeCondition")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditional struct {
	value *Conditional
	isSet bool
}

func (v NullableConditional) Get() *Conditional {
	return v.value
}

func (v *NullableConditional) Set(val *Conditional) {
	v.value = val
	v.isSet = true
}

func (v NullableConditional) IsSet() bool {
	return v.isSet
}

func (v *NullableConditional) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditional(val *Conditional) *NullableConditional {
	return &NullableConditional{value: val, isSet: true}
}

func (v NullableConditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


