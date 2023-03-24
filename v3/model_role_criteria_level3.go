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

// RoleCriteriaLevel3 Defines STANDARD type Role membership
type RoleCriteriaLevel3 struct {
	Operation *RoleCriteriaOperation `json:"operation,omitempty"`
	Key NullableRoleCriteriaKey `json:"key,omitempty"`
	// String value to test the Identity attribute, Account attribute, or Entitlement specified in the key w/r/t the specified operation. If this criteria is a leaf node, that is, if the operation is one of EQUALS, NOT_EQUALS, CONTAINS, STARTS_WITH, or ENDS_WITH, this field is required. Otherwise, specifying it is an error.
	StringValue *string `json:"stringValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleCriteriaLevel3 RoleCriteriaLevel3

// NewRoleCriteriaLevel3 instantiates a new RoleCriteriaLevel3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCriteriaLevel3() *RoleCriteriaLevel3 {
	this := RoleCriteriaLevel3{}
	return &this
}

// NewRoleCriteriaLevel3WithDefaults instantiates a new RoleCriteriaLevel3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCriteriaLevel3WithDefaults() *RoleCriteriaLevel3 {
	this := RoleCriteriaLevel3{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RoleCriteriaLevel3) GetOperation() RoleCriteriaOperation {
	if o == nil || isNil(o.Operation) {
		var ret RoleCriteriaOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCriteriaLevel3) GetOperationOk() (*RoleCriteriaOperation, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RoleCriteriaLevel3) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given RoleCriteriaOperation and assigns it to the Operation field.
func (o *RoleCriteriaLevel3) SetOperation(v RoleCriteriaOperation) {
	o.Operation = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleCriteriaLevel3) GetKey() RoleCriteriaKey {
	if o == nil || isNil(o.Key.Get()) {
		var ret RoleCriteriaKey
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleCriteriaLevel3) GetKeyOk() (*RoleCriteriaKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *RoleCriteriaLevel3) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableRoleCriteriaKey and assigns it to the Key field.
func (o *RoleCriteriaLevel3) SetKey(v RoleCriteriaKey) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *RoleCriteriaLevel3) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *RoleCriteriaLevel3) UnsetKey() {
	o.Key.Unset()
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *RoleCriteriaLevel3) GetStringValue() string {
	if o == nil || isNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCriteriaLevel3) GetStringValueOk() (*string, bool) {
	if o == nil || isNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *RoleCriteriaLevel3) HasStringValue() bool {
	if o != nil && !isNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *RoleCriteriaLevel3) SetStringValue(v string) {
	o.StringValue = &v
}

func (o RoleCriteriaLevel3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !isNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleCriteriaLevel3) UnmarshalJSON(bytes []byte) (err error) {
	varRoleCriteriaLevel3 := _RoleCriteriaLevel3{}

	if err = json.Unmarshal(bytes, &varRoleCriteriaLevel3); err == nil {
		*o = RoleCriteriaLevel3(varRoleCriteriaLevel3)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "key")
		delete(additionalProperties, "stringValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleCriteriaLevel3 struct {
	value *RoleCriteriaLevel3
	isSet bool
}

func (v NullableRoleCriteriaLevel3) Get() *RoleCriteriaLevel3 {
	return v.value
}

func (v *NullableRoleCriteriaLevel3) Set(val *RoleCriteriaLevel3) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCriteriaLevel3) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCriteriaLevel3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCriteriaLevel3(val *RoleCriteriaLevel3) *NullableRoleCriteriaLevel3 {
	return &NullableRoleCriteriaLevel3{value: val, isSet: true}
}

func (v NullableRoleCriteriaLevel3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCriteriaLevel3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

