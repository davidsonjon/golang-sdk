/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}

// ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner struct for ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner
type ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner struct {
	// The name of the attribute being provisioned.
	AttributeName string `json:"attributeName"`
	// The value of the attribute being provisioned.
	AttributeValue NullableString `json:"attributeValue,omitempty"`
	// The operation to handle the attribute.
	Operation map[string]interface{} `json:"operation"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner

// NewProvisioningCompletedAccountRequestsInnerAttributeRequestsInner instantiates a new ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningCompletedAccountRequestsInnerAttributeRequestsInner(attributeName string, operation map[string]interface{}) *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	this := ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}
	this.AttributeName = attributeName
	this.Operation = operation
	return &this
}

// NewProvisioningCompletedAccountRequestsInnerAttributeRequestsInnerWithDefaults instantiates a new ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningCompletedAccountRequestsInnerAttributeRequestsInnerWithDefaults() *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	this := ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeValue() string {
	if o == nil || isNil(o.AttributeValue.Get()) {
		var ret string
		return ret
	}
	return *o.AttributeValue.Get()
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeValue.Get(), o.AttributeValue.IsSet()
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) HasAttributeValue() bool {
	if o != nil && o.AttributeValue.IsSet() {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given NullableString and assigns it to the AttributeValue field.
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetAttributeValue(v string) {
	o.AttributeValue.Set(&v)
}
// SetAttributeValueNil sets the value for AttributeValue to be an explicit nil
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetAttributeValueNil() {
	o.AttributeValue.Set(nil)
}

// UnsetAttributeValue ensures that no value is present for AttributeValue, not even an explicit nil
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) UnsetAttributeValue() {
	o.AttributeValue.Unset()
}

// GetOperation returns the Operation field value
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetOperation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetOperationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Operation, true
}

// SetOperation sets field value
func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetOperation(v map[string]interface{}) {
	o.Operation = v
}

func (o ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeName"] = o.AttributeName
	if o.AttributeValue.IsSet() {
		toSerialize["attributeValue"] = o.AttributeValue.Get()
	}
	toSerialize["operation"] = o.Operation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributeName",
		"operation",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProvisioningCompletedAccountRequestsInnerAttributeRequestsInner := _ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}

	if err = json.Unmarshal(bytes, &varProvisioningCompletedAccountRequestsInnerAttributeRequestsInner); err == nil {
	*o = ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner(varProvisioningCompletedAccountRequestsInnerAttributeRequestsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "attributeValue")
		delete(additionalProperties, "operation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner struct {
	value *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner
	isSet bool
}

func (v NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) Get() *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	return v.value
}

func (v *NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) Set(val *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner(val *ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) *NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	return &NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{value: val, isSet: true}
}

func (v NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

