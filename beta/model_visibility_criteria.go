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

// checks if the VisibilityCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisibilityCriteria{}

// VisibilityCriteria struct for VisibilityCriteria
type VisibilityCriteria struct {
	Expression *Expression `json:"expression,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VisibilityCriteria VisibilityCriteria

// NewVisibilityCriteria instantiates a new VisibilityCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisibilityCriteria() *VisibilityCriteria {
	this := VisibilityCriteria{}
	return &this
}

// NewVisibilityCriteriaWithDefaults instantiates a new VisibilityCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisibilityCriteriaWithDefaults() *VisibilityCriteria {
	this := VisibilityCriteria{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *VisibilityCriteria) GetExpression() Expression {
	if o == nil || isNil(o.Expression) {
		var ret Expression
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisibilityCriteria) GetExpressionOk() (*Expression, bool) {
	if o == nil || isNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *VisibilityCriteria) HasExpression() bool {
	if o != nil && !isNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given Expression and assigns it to the Expression field.
func (o *VisibilityCriteria) SetExpression(v Expression) {
	o.Expression = &v
}

func (o VisibilityCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisibilityCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VisibilityCriteria) UnmarshalJSON(bytes []byte) (err error) {
	varVisibilityCriteria := _VisibilityCriteria{}

	if err = json.Unmarshal(bytes, &varVisibilityCriteria); err == nil {
	*o = VisibilityCriteria(varVisibilityCriteria)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVisibilityCriteria struct {
	value *VisibilityCriteria
	isSet bool
}

func (v NullableVisibilityCriteria) Get() *VisibilityCriteria {
	return v.value
}

func (v *NullableVisibilityCriteria) Set(val *VisibilityCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableVisibilityCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableVisibilityCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisibilityCriteria(val *VisibilityCriteria) *NullableVisibilityCriteria {
	return &NullableVisibilityCriteria{value: val, isSet: true}
}

func (v NullableVisibilityCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisibilityCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


