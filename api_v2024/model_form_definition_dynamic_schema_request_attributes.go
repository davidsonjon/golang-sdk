/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the FormDefinitionDynamicSchemaRequestAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDefinitionDynamicSchemaRequestAttributes{}

// FormDefinitionDynamicSchemaRequestAttributes struct for FormDefinitionDynamicSchemaRequestAttributes
type FormDefinitionDynamicSchemaRequestAttributes struct {
	// FormDefinitionID is a unique guid identifying this form definition
	FormDefinitionId *string `json:"formDefinitionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormDefinitionDynamicSchemaRequestAttributes FormDefinitionDynamicSchemaRequestAttributes

// NewFormDefinitionDynamicSchemaRequestAttributes instantiates a new FormDefinitionDynamicSchemaRequestAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDefinitionDynamicSchemaRequestAttributes() *FormDefinitionDynamicSchemaRequestAttributes {
	this := FormDefinitionDynamicSchemaRequestAttributes{}
	return &this
}

// NewFormDefinitionDynamicSchemaRequestAttributesWithDefaults instantiates a new FormDefinitionDynamicSchemaRequestAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDefinitionDynamicSchemaRequestAttributesWithDefaults() *FormDefinitionDynamicSchemaRequestAttributes {
	this := FormDefinitionDynamicSchemaRequestAttributes{}
	return &this
}

// GetFormDefinitionId returns the FormDefinitionId field value if set, zero value otherwise.
func (o *FormDefinitionDynamicSchemaRequestAttributes) GetFormDefinitionId() string {
	if o == nil || isNil(o.FormDefinitionId) {
		var ret string
		return ret
	}
	return *o.FormDefinitionId
}

// GetFormDefinitionIdOk returns a tuple with the FormDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionDynamicSchemaRequestAttributes) GetFormDefinitionIdOk() (*string, bool) {
	if o == nil || isNil(o.FormDefinitionId) {
		return nil, false
	}
	return o.FormDefinitionId, true
}

// HasFormDefinitionId returns a boolean if a field has been set.
func (o *FormDefinitionDynamicSchemaRequestAttributes) HasFormDefinitionId() bool {
	if o != nil && !isNil(o.FormDefinitionId) {
		return true
	}

	return false
}

// SetFormDefinitionId gets a reference to the given string and assigns it to the FormDefinitionId field.
func (o *FormDefinitionDynamicSchemaRequestAttributes) SetFormDefinitionId(v string) {
	o.FormDefinitionId = &v
}

func (o FormDefinitionDynamicSchemaRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDefinitionDynamicSchemaRequestAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FormDefinitionId) {
		toSerialize["formDefinitionId"] = o.FormDefinitionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormDefinitionDynamicSchemaRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varFormDefinitionDynamicSchemaRequestAttributes := _FormDefinitionDynamicSchemaRequestAttributes{}

	if err = json.Unmarshal(bytes, &varFormDefinitionDynamicSchemaRequestAttributes); err == nil {
	*o = FormDefinitionDynamicSchemaRequestAttributes(varFormDefinitionDynamicSchemaRequestAttributes)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "formDefinitionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormDefinitionDynamicSchemaRequestAttributes struct {
	value *FormDefinitionDynamicSchemaRequestAttributes
	isSet bool
}

func (v NullableFormDefinitionDynamicSchemaRequestAttributes) Get() *FormDefinitionDynamicSchemaRequestAttributes {
	return v.value
}

func (v *NullableFormDefinitionDynamicSchemaRequestAttributes) Set(val *FormDefinitionDynamicSchemaRequestAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDefinitionDynamicSchemaRequestAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDefinitionDynamicSchemaRequestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDefinitionDynamicSchemaRequestAttributes(val *FormDefinitionDynamicSchemaRequestAttributes) *NullableFormDefinitionDynamicSchemaRequestAttributes {
	return &NullableFormDefinitionDynamicSchemaRequestAttributes{value: val, isSet: true}
}

func (v NullableFormDefinitionDynamicSchemaRequestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDefinitionDynamicSchemaRequestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


