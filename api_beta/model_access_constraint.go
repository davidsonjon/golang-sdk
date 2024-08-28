/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the AccessConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessConstraint{}

// AccessConstraint struct for AccessConstraint
type AccessConstraint struct {
	// Type of Access
	Type string `json:"type"`
	// Must be set only if operator is SELECTED.
	Ids []string `json:"ids,omitempty"`
	// Used to determine whether the scope of the campaign should be reduced for selected ids or all.
	Operator string `json:"operator"`
	AdditionalProperties map[string]interface{}
}

type _AccessConstraint AccessConstraint

// NewAccessConstraint instantiates a new AccessConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessConstraint(type_ string, operator string) *AccessConstraint {
	this := AccessConstraint{}
	this.Type = type_
	this.Operator = operator
	return &this
}

// NewAccessConstraintWithDefaults instantiates a new AccessConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessConstraintWithDefaults() *AccessConstraint {
	this := AccessConstraint{}
	return &this
}

// GetType returns the Type field value
func (o *AccessConstraint) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessConstraint) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessConstraint) SetType(v string) {
	o.Type = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *AccessConstraint) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessConstraint) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *AccessConstraint) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *AccessConstraint) SetIds(v []string) {
	o.Ids = v
}

// GetOperator returns the Operator field value
func (o *AccessConstraint) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *AccessConstraint) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *AccessConstraint) SetOperator(v string) {
	o.Operator = v
}

func (o AccessConstraint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	toSerialize["operator"] = o.Operator

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessConstraint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"operator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessConstraint := _AccessConstraint{}

	err = json.Unmarshal(data, &varAccessConstraint)

	if err != nil {
		return err
	}

	*o = AccessConstraint(varAccessConstraint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "operator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessConstraint struct {
	value *AccessConstraint
	isSet bool
}

func (v NullableAccessConstraint) Get() *AccessConstraint {
	return v.value
}

func (v *NullableAccessConstraint) Set(val *AccessConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessConstraint(val *AccessConstraint) *NullableAccessConstraint {
	return &NullableAccessConstraint{value: val, isSet: true}
}

func (v NullableAccessConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


