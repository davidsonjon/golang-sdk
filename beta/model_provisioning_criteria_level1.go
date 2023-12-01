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

// checks if the ProvisioningCriteriaLevel1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningCriteriaLevel1{}

// ProvisioningCriteriaLevel1 Defines matching criteria for an Account to be provisioned with a specific Access Profile
type ProvisioningCriteriaLevel1 struct {
	Operation *ProvisioningCriteriaOperation `json:"operation,omitempty"`
	// Name of the Account attribute to be tested. If **operation** is one of EQUALS, NOT_EQUALS, CONTAINS, or HAS, this field is required. Otherwise, specifying it is an error.
	Attribute NullableString `json:"attribute,omitempty"`
	// String value to test the Account attribute w/r/t the specified operation. If the operation is one of EQUALS, NOT_EQUALS, or CONTAINS, this field is required. Otherwise, specifying it is an error. If the Attribute is not String-typed, it will be converted to the appropriate type.
	Value NullableString `json:"value,omitempty"`
	// Array of child criteria. Required if the operation is AND or OR, otherwise it must be left null. A maximum of three levels of criteria are supported, including leaf nodes.
	Children []ProvisioningCriteriaLevel2 `json:"children,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningCriteriaLevel1 ProvisioningCriteriaLevel1

// NewProvisioningCriteriaLevel1 instantiates a new ProvisioningCriteriaLevel1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningCriteriaLevel1() *ProvisioningCriteriaLevel1 {
	this := ProvisioningCriteriaLevel1{}
	return &this
}

// NewProvisioningCriteriaLevel1WithDefaults instantiates a new ProvisioningCriteriaLevel1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningCriteriaLevel1WithDefaults() *ProvisioningCriteriaLevel1 {
	this := ProvisioningCriteriaLevel1{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ProvisioningCriteriaLevel1) GetOperation() ProvisioningCriteriaOperation {
	if o == nil || isNil(o.Operation) {
		var ret ProvisioningCriteriaOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningCriteriaLevel1) GetOperationOk() (*ProvisioningCriteriaOperation, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ProvisioningCriteriaLevel1) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given ProvisioningCriteriaOperation and assigns it to the Operation field.
func (o *ProvisioningCriteriaLevel1) SetOperation(v ProvisioningCriteriaOperation) {
	o.Operation = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningCriteriaLevel1) GetAttribute() string {
	if o == nil || isNil(o.Attribute.Get()) {
		var ret string
		return ret
	}
	return *o.Attribute.Get()
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningCriteriaLevel1) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute.Get(), o.Attribute.IsSet()
}

// HasAttribute returns a boolean if a field has been set.
func (o *ProvisioningCriteriaLevel1) HasAttribute() bool {
	if o != nil && o.Attribute.IsSet() {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given NullableString and assigns it to the Attribute field.
func (o *ProvisioningCriteriaLevel1) SetAttribute(v string) {
	o.Attribute.Set(&v)
}
// SetAttributeNil sets the value for Attribute to be an explicit nil
func (o *ProvisioningCriteriaLevel1) SetAttributeNil() {
	o.Attribute.Set(nil)
}

// UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
func (o *ProvisioningCriteriaLevel1) UnsetAttribute() {
	o.Attribute.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningCriteriaLevel1) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningCriteriaLevel1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProvisioningCriteriaLevel1) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProvisioningCriteriaLevel1) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ProvisioningCriteriaLevel1) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProvisioningCriteriaLevel1) UnsetValue() {
	o.Value.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningCriteriaLevel1) GetChildren() []ProvisioningCriteriaLevel2 {
	if o == nil {
		var ret []ProvisioningCriteriaLevel2
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningCriteriaLevel1) GetChildrenOk() ([]ProvisioningCriteriaLevel2, bool) {
	if o == nil || isNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ProvisioningCriteriaLevel1) HasChildren() bool {
	if o != nil && isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ProvisioningCriteriaLevel2 and assigns it to the Children field.
func (o *ProvisioningCriteriaLevel1) SetChildren(v []ProvisioningCriteriaLevel2) {
	o.Children = v
}

func (o ProvisioningCriteriaLevel1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningCriteriaLevel1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
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

func (o *ProvisioningCriteriaLevel1) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningCriteriaLevel1 := _ProvisioningCriteriaLevel1{}

	if err = json.Unmarshal(bytes, &varProvisioningCriteriaLevel1); err == nil {
	*o = ProvisioningCriteriaLevel1(varProvisioningCriteriaLevel1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "children")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningCriteriaLevel1 struct {
	value *ProvisioningCriteriaLevel1
	isSet bool
}

func (v NullableProvisioningCriteriaLevel1) Get() *ProvisioningCriteriaLevel1 {
	return v.value
}

func (v *NullableProvisioningCriteriaLevel1) Set(val *ProvisioningCriteriaLevel1) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningCriteriaLevel1) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningCriteriaLevel1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningCriteriaLevel1(val *ProvisioningCriteriaLevel1) *NullableProvisioningCriteriaLevel1 {
	return &NullableProvisioningCriteriaLevel1{value: val, isSet: true}
}

func (v NullableProvisioningCriteriaLevel1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningCriteriaLevel1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


