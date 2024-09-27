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

// checks if the RoleMetadataBulkUpdateByFilterRequestValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMetadataBulkUpdateByFilterRequestValuesInner{}

// RoleMetadataBulkUpdateByFilterRequestValuesInner struct for RoleMetadataBulkUpdateByFilterRequestValuesInner
type RoleMetadataBulkUpdateByFilterRequestValuesInner struct {
	// the key of metadata attribute
	AttributeKey *string `json:"attributeKey,omitempty"`
	// the values of attribute to be updated
	Values []string `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _RoleMetadataBulkUpdateByFilterRequestValuesInner RoleMetadataBulkUpdateByFilterRequestValuesInner

// NewRoleMetadataBulkUpdateByFilterRequestValuesInner instantiates a new RoleMetadataBulkUpdateByFilterRequestValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMetadataBulkUpdateByFilterRequestValuesInner(values []string) *RoleMetadataBulkUpdateByFilterRequestValuesInner {
	this := RoleMetadataBulkUpdateByFilterRequestValuesInner{}
	this.Values = values
	return &this
}

// NewRoleMetadataBulkUpdateByFilterRequestValuesInnerWithDefaults instantiates a new RoleMetadataBulkUpdateByFilterRequestValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMetadataBulkUpdateByFilterRequestValuesInnerWithDefaults() *RoleMetadataBulkUpdateByFilterRequestValuesInner {
	this := RoleMetadataBulkUpdateByFilterRequestValuesInner{}
	return &this
}

// GetAttributeKey returns the AttributeKey field value if set, zero value otherwise.
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) GetAttributeKey() string {
	if o == nil || IsNil(o.AttributeKey) {
		var ret string
		return ret
	}
	return *o.AttributeKey
}

// GetAttributeKeyOk returns a tuple with the AttributeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) GetAttributeKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeKey) {
		return nil, false
	}
	return o.AttributeKey, true
}

// HasAttributeKey returns a boolean if a field has been set.
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) HasAttributeKey() bool {
	if o != nil && !IsNil(o.AttributeKey) {
		return true
	}

	return false
}

// SetAttributeKey gets a reference to the given string and assigns it to the AttributeKey field.
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) SetAttributeKey(v string) {
	o.AttributeKey = &v
}

// GetValues returns the Values field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) SetValues(v []string) {
	o.Values = v
}

func (o RoleMetadataBulkUpdateByFilterRequestValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMetadataBulkUpdateByFilterRequestValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeKey) {
		toSerialize["attributeKey"] = o.AttributeKey
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMetadataBulkUpdateByFilterRequestValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
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

	varRoleMetadataBulkUpdateByFilterRequestValuesInner := _RoleMetadataBulkUpdateByFilterRequestValuesInner{}

	err = json.Unmarshal(data, &varRoleMetadataBulkUpdateByFilterRequestValuesInner)

	if err != nil {
		return err
	}

	*o = RoleMetadataBulkUpdateByFilterRequestValuesInner(varRoleMetadataBulkUpdateByFilterRequestValuesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeKey")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMetadataBulkUpdateByFilterRequestValuesInner struct {
	value *RoleMetadataBulkUpdateByFilterRequestValuesInner
	isSet bool
}

func (v NullableRoleMetadataBulkUpdateByFilterRequestValuesInner) Get() *RoleMetadataBulkUpdateByFilterRequestValuesInner {
	return v.value
}

func (v *NullableRoleMetadataBulkUpdateByFilterRequestValuesInner) Set(val *RoleMetadataBulkUpdateByFilterRequestValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMetadataBulkUpdateByFilterRequestValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMetadataBulkUpdateByFilterRequestValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMetadataBulkUpdateByFilterRequestValuesInner(val *RoleMetadataBulkUpdateByFilterRequestValuesInner) *NullableRoleMetadataBulkUpdateByFilterRequestValuesInner {
	return &NullableRoleMetadataBulkUpdateByFilterRequestValuesInner{value: val, isSet: true}
}

func (v NullableRoleMetadataBulkUpdateByFilterRequestValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMetadataBulkUpdateByFilterRequestValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


