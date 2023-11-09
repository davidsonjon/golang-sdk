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

// checks if the TaskResultDetailsReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResultDetailsReturnsInner{}

// TaskResultDetailsReturnsInner struct for TaskResultDetailsReturnsInner
type TaskResultDetailsReturnsInner struct {
	// Attribute description.
	DisplayLabel *string `json:"displayLabel,omitempty"`
	// System or database attribute name.
	AttributeName *string `json:"attributeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskResultDetailsReturnsInner TaskResultDetailsReturnsInner

// NewTaskResultDetailsReturnsInner instantiates a new TaskResultDetailsReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResultDetailsReturnsInner() *TaskResultDetailsReturnsInner {
	this := TaskResultDetailsReturnsInner{}
	return &this
}

// NewTaskResultDetailsReturnsInnerWithDefaults instantiates a new TaskResultDetailsReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResultDetailsReturnsInnerWithDefaults() *TaskResultDetailsReturnsInner {
	this := TaskResultDetailsReturnsInner{}
	return &this
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *TaskResultDetailsReturnsInner) GetDisplayLabel() string {
	if o == nil || isNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultDetailsReturnsInner) GetDisplayLabelOk() (*string, bool) {
	if o == nil || isNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *TaskResultDetailsReturnsInner) HasDisplayLabel() bool {
	if o != nil && !isNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *TaskResultDetailsReturnsInner) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *TaskResultDetailsReturnsInner) GetAttributeName() string {
	if o == nil || isNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultDetailsReturnsInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *TaskResultDetailsReturnsInner) HasAttributeName() bool {
	if o != nil && !isNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *TaskResultDetailsReturnsInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

func (o TaskResultDetailsReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskResultDetailsReturnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if !isNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskResultDetailsReturnsInner) UnmarshalJSON(bytes []byte) (err error) {
	varTaskResultDetailsReturnsInner := _TaskResultDetailsReturnsInner{}

	if err = json.Unmarshal(bytes, &varTaskResultDetailsReturnsInner); err == nil {
		*o = TaskResultDetailsReturnsInner(varTaskResultDetailsReturnsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayLabel")
		delete(additionalProperties, "attributeName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskResultDetailsReturnsInner struct {
	value *TaskResultDetailsReturnsInner
	isSet bool
}

func (v NullableTaskResultDetailsReturnsInner) Get() *TaskResultDetailsReturnsInner {
	return v.value
}

func (v *NullableTaskResultDetailsReturnsInner) Set(val *TaskResultDetailsReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResultDetailsReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResultDetailsReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResultDetailsReturnsInner(val *TaskResultDetailsReturnsInner) *NullableTaskResultDetailsReturnsInner {
	return &NullableTaskResultDetailsReturnsInner{value: val, isSet: true}
}

func (v NullableTaskResultDetailsReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResultDetailsReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


