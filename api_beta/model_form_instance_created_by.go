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

// checks if the FormInstanceCreatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormInstanceCreatedBy{}

// FormInstanceCreatedBy struct for FormInstanceCreatedBy
type FormInstanceCreatedBy struct {
	// ID is a unique identifier
	Id *string `json:"id,omitempty"`
	// Type is a form instance created by type enum value WORKFLOW_EXECUTION FormInstanceCreatedByTypeWorkflowExecution SOURCE FormInstanceCreatedByTypeSource
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormInstanceCreatedBy FormInstanceCreatedBy

// NewFormInstanceCreatedBy instantiates a new FormInstanceCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormInstanceCreatedBy() *FormInstanceCreatedBy {
	this := FormInstanceCreatedBy{}
	return &this
}

// NewFormInstanceCreatedByWithDefaults instantiates a new FormInstanceCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormInstanceCreatedByWithDefaults() *FormInstanceCreatedBy {
	this := FormInstanceCreatedBy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormInstanceCreatedBy) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceCreatedBy) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormInstanceCreatedBy) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormInstanceCreatedBy) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FormInstanceCreatedBy) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceCreatedBy) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FormInstanceCreatedBy) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FormInstanceCreatedBy) SetType(v string) {
	o.Type = &v
}

func (o FormInstanceCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormInstanceCreatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormInstanceCreatedBy) UnmarshalJSON(bytes []byte) (err error) {
	varFormInstanceCreatedBy := _FormInstanceCreatedBy{}

	if err = json.Unmarshal(bytes, &varFormInstanceCreatedBy); err == nil {
	*o = FormInstanceCreatedBy(varFormInstanceCreatedBy)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormInstanceCreatedBy struct {
	value *FormInstanceCreatedBy
	isSet bool
}

func (v NullableFormInstanceCreatedBy) Get() *FormInstanceCreatedBy {
	return v.value
}

func (v *NullableFormInstanceCreatedBy) Set(val *FormInstanceCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableFormInstanceCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableFormInstanceCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormInstanceCreatedBy(val *FormInstanceCreatedBy) *NullableFormInstanceCreatedBy {
	return &NullableFormInstanceCreatedBy{value: val, isSet: true}
}

func (v NullableFormInstanceCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormInstanceCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


