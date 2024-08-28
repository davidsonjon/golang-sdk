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

// checks if the TaskReturnDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskReturnDetails{}

// TaskReturnDetails Task return details
type TaskReturnDetails struct {
	// Display name of the TaskReturnDetails
	Name string `json:"name"`
	// Attribute the TaskReturnDetails is for
	AttributeName string `json:"attributeName"`
	AdditionalProperties map[string]interface{}
}

type _TaskReturnDetails TaskReturnDetails

// NewTaskReturnDetails instantiates a new TaskReturnDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskReturnDetails(name string, attributeName string) *TaskReturnDetails {
	this := TaskReturnDetails{}
	this.Name = name
	this.AttributeName = attributeName
	return &this
}

// NewTaskReturnDetailsWithDefaults instantiates a new TaskReturnDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskReturnDetailsWithDefaults() *TaskReturnDetails {
	this := TaskReturnDetails{}
	return &this
}

// GetName returns the Name field value
func (o *TaskReturnDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskReturnDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaskReturnDetails) SetName(v string) {
	o.Name = v
}

// GetAttributeName returns the AttributeName field value
func (o *TaskReturnDetails) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *TaskReturnDetails) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *TaskReturnDetails) SetAttributeName(v string) {
	o.AttributeName = v
}

func (o TaskReturnDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskReturnDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["attributeName"] = o.AttributeName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskReturnDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"attributeName",
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

	varTaskReturnDetails := _TaskReturnDetails{}

	err = json.Unmarshal(data, &varTaskReturnDetails)

	if err != nil {
		return err
	}

	*o = TaskReturnDetails(varTaskReturnDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "attributeName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskReturnDetails struct {
	value *TaskReturnDetails
	isSet bool
}

func (v NullableTaskReturnDetails) Get() *TaskReturnDetails {
	return v.value
}

func (v *NullableTaskReturnDetails) Set(val *TaskReturnDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskReturnDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskReturnDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskReturnDetails(val *TaskReturnDetails) *NullableTaskReturnDetails {
	return &NullableTaskReturnDetails{value: val, isSet: true}
}

func (v NullableTaskReturnDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskReturnDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


