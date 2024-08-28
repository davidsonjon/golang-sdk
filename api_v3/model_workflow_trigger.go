/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the WorkflowTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTrigger{}

// WorkflowTrigger The trigger that starts the workflow
type WorkflowTrigger struct {
	// The trigger type
	Type string `json:"type"`
	DisplayName NullableString `json:"displayName,omitempty"`
	// Workflow Trigger Attributes.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTrigger WorkflowTrigger

// NewWorkflowTrigger instantiates a new WorkflowTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTrigger(type_ string, attributes map[string]interface{}) *WorkflowTrigger {
	this := WorkflowTrigger{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTriggerWithDefaults() *WorkflowTrigger {
	this := WorkflowTrigger{}
	return &this
}

// GetType returns the Type field value
func (o *WorkflowTrigger) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WorkflowTrigger) SetType(v string) {
	o.Type = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTrigger) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTrigger) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *WorkflowTrigger) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *WorkflowTrigger) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *WorkflowTrigger) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetAttributes returns the Attributes field value
func (o *WorkflowTrigger) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *WorkflowTrigger) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o WorkflowTrigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowTrigger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varWorkflowTrigger := _WorkflowTrigger{}

	err = json.Unmarshal(data, &varWorkflowTrigger)

	if err != nil {
		return err
	}

	*o = WorkflowTrigger(varWorkflowTrigger)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTrigger struct {
	value *WorkflowTrigger
	isSet bool
}

func (v NullableWorkflowTrigger) Get() *WorkflowTrigger {
	return v.value
}

func (v *NullableWorkflowTrigger) Set(val *WorkflowTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTrigger(val *WorkflowTrigger) *NullableWorkflowTrigger {
	return &NullableWorkflowTrigger{value: val, isSet: true}
}

func (v NullableWorkflowTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


