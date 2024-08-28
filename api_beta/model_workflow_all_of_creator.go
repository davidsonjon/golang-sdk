/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the WorkflowAllOfCreator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowAllOfCreator{}

// WorkflowAllOfCreator Workflow creator's identity.
type WorkflowAllOfCreator struct {
	// Workflow creator's DTO type.
	Type *string `json:"type,omitempty"`
	// Workflow creator's identity ID.
	Id *string `json:"id,omitempty"`
	// Workflow creator's display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAllOfCreator WorkflowAllOfCreator

// NewWorkflowAllOfCreator instantiates a new WorkflowAllOfCreator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAllOfCreator() *WorkflowAllOfCreator {
	this := WorkflowAllOfCreator{}
	return &this
}

// NewWorkflowAllOfCreatorWithDefaults instantiates a new WorkflowAllOfCreator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAllOfCreatorWithDefaults() *WorkflowAllOfCreator {
	this := WorkflowAllOfCreator{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowAllOfCreator) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOfCreator) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowAllOfCreator) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowAllOfCreator) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowAllOfCreator) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOfCreator) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowAllOfCreator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowAllOfCreator) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowAllOfCreator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOfCreator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowAllOfCreator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowAllOfCreator) SetName(v string) {
	o.Name = &v
}

func (o WorkflowAllOfCreator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowAllOfCreator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowAllOfCreator) UnmarshalJSON(data []byte) (err error) {
	varWorkflowAllOfCreator := _WorkflowAllOfCreator{}

	err = json.Unmarshal(data, &varWorkflowAllOfCreator)

	if err != nil {
		return err
	}

	*o = WorkflowAllOfCreator(varWorkflowAllOfCreator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowAllOfCreator struct {
	value *WorkflowAllOfCreator
	isSet bool
}

func (v NullableWorkflowAllOfCreator) Get() *WorkflowAllOfCreator {
	return v.value
}

func (v *NullableWorkflowAllOfCreator) Set(val *WorkflowAllOfCreator) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAllOfCreator) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAllOfCreator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAllOfCreator(val *WorkflowAllOfCreator) *NullableWorkflowAllOfCreator {
	return &NullableWorkflowAllOfCreator{value: val, isSet: true}
}

func (v NullableWorkflowAllOfCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAllOfCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


