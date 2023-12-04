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

// checks if the RoleCriteriaLevel1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleCriteriaLevel1{}

// RoleCriteriaLevel1 Defines STANDARD type Role membership
type RoleCriteriaLevel1 struct {
	Operation *RoleCriteriaOperation `json:"operation,omitempty"`
	Key NullableRoleCriteriaKey `json:"key,omitempty"`
	// String value to test the Identity attribute, Account attribute, or Entitlement specified in the key w/r/t the specified operation. If this criteria is a leaf node, that is, if the operation is one of EQUALS, NOT_EQUALS, CONTAINS, STARTS_WITH, or ENDS_WITH, this field is required. Otherwise, specifying it is an error.
	StringValue NullableString `json:"stringValue,omitempty"`
	// Array of child criteria. Required if the operation is AND or OR, otherwise it must be left null. A maximum of three levels of criteria are supported, including leaf nodes. Additionally, AND nodes can only be children or OR nodes and vice-versa.
	Children []RoleCriteriaLevel2 `json:"children,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleCriteriaLevel1 RoleCriteriaLevel1

// NewRoleCriteriaLevel1 instantiates a new RoleCriteriaLevel1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCriteriaLevel1() *RoleCriteriaLevel1 {
	this := RoleCriteriaLevel1{}
	return &this
}

// NewRoleCriteriaLevel1WithDefaults instantiates a new RoleCriteriaLevel1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCriteriaLevel1WithDefaults() *RoleCriteriaLevel1 {
	this := RoleCriteriaLevel1{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RoleCriteriaLevel1) GetOperation() RoleCriteriaOperation {
	if o == nil || isNil(o.Operation) {
		var ret RoleCriteriaOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCriteriaLevel1) GetOperationOk() (*RoleCriteriaOperation, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given RoleCriteriaOperation and assigns it to the Operation field.
func (o *RoleCriteriaLevel1) SetOperation(v RoleCriteriaOperation) {
	o.Operation = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleCriteriaLevel1) GetKey() RoleCriteriaKey {
	if o == nil || isNil(o.Key.Get()) {
		var ret RoleCriteriaKey
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleCriteriaLevel1) GetKeyOk() (*RoleCriteriaKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableRoleCriteriaKey and assigns it to the Key field.
func (o *RoleCriteriaLevel1) SetKey(v RoleCriteriaKey) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *RoleCriteriaLevel1) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *RoleCriteriaLevel1) UnsetKey() {
	o.Key.Unset()
}

// GetStringValue returns the StringValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleCriteriaLevel1) GetStringValue() string {
	if o == nil || isNil(o.StringValue.Get()) {
		var ret string
		return ret
	}
	return *o.StringValue.Get()
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleCriteriaLevel1) GetStringValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringValue.Get(), o.StringValue.IsSet()
}

// HasStringValue returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasStringValue() bool {
	if o != nil && o.StringValue.IsSet() {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given NullableString and assigns it to the StringValue field.
func (o *RoleCriteriaLevel1) SetStringValue(v string) {
	o.StringValue.Set(&v)
}
// SetStringValueNil sets the value for StringValue to be an explicit nil
func (o *RoleCriteriaLevel1) SetStringValueNil() {
	o.StringValue.Set(nil)
}

// UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
func (o *RoleCriteriaLevel1) UnsetStringValue() {
	o.StringValue.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleCriteriaLevel1) GetChildren() []RoleCriteriaLevel2 {
	if o == nil {
		var ret []RoleCriteriaLevel2
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleCriteriaLevel1) GetChildrenOk() ([]RoleCriteriaLevel2, bool) {
	if o == nil || isNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasChildren() bool {
	if o != nil && isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []RoleCriteriaLevel2 and assigns it to the Children field.
func (o *RoleCriteriaLevel1) SetChildren(v []RoleCriteriaLevel2) {
	o.Children = v
}

func (o RoleCriteriaLevel1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleCriteriaLevel1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.StringValue.IsSet() {
		toSerialize["stringValue"] = o.StringValue.Get()
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleCriteriaLevel1) UnmarshalJSON(bytes []byte) (err error) {
	varRoleCriteriaLevel1 := _RoleCriteriaLevel1{}

	if err = json.Unmarshal(bytes, &varRoleCriteriaLevel1); err == nil {
	*o = RoleCriteriaLevel1(varRoleCriteriaLevel1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "key")
		delete(additionalProperties, "stringValue")
		delete(additionalProperties, "children")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleCriteriaLevel1 struct {
	value *RoleCriteriaLevel1
	isSet bool
}

func (v NullableRoleCriteriaLevel1) Get() *RoleCriteriaLevel1 {
	return v.value
}

func (v *NullableRoleCriteriaLevel1) Set(val *RoleCriteriaLevel1) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCriteriaLevel1) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCriteriaLevel1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCriteriaLevel1(val *RoleCriteriaLevel1) *NullableRoleCriteriaLevel1 {
	return &NullableRoleCriteriaLevel1{value: val, isSet: true}
}

func (v NullableRoleCriteriaLevel1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCriteriaLevel1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


