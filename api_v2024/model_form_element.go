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

// checks if the FormElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormElement{}

// FormElement struct for FormElement
type FormElement struct {
	// Form element identifier.
	Id *string `json:"id,omitempty"`
	// FormElementType value.  TEXT FormElementTypeText TOGGLE FormElementTypeToggle TEXTAREA FormElementTypeTextArea HIDDEN FormElementTypeHidden PHONE FormElementTypePhone EMAIL FormElementTypeEmail SELECT FormElementTypeSelect DATE FormElementTypeDate SECTION FormElementTypeSection COLUMN_SET FormElementTypeColumns IMAGE FormElementTypeImage DESCRIPTION FormElementTypeDescription
	ElementType *string `json:"elementType,omitempty"`
	// Config object.
	Config map[string]map[string]interface{} `json:"config,omitempty"`
	// Technical key.
	Key *string `json:"key,omitempty"`
	Validations []FormElementValidationsSet `json:"validations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormElement FormElement

// NewFormElement instantiates a new FormElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormElement() *FormElement {
	this := FormElement{}
	return &this
}

// NewFormElementWithDefaults instantiates a new FormElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormElementWithDefaults() *FormElement {
	this := FormElement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormElement) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormElement) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormElement) SetId(v string) {
	o.Id = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *FormElement) GetElementType() string {
	if o == nil || isNil(o.ElementType) {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetElementTypeOk() (*string, bool) {
	if o == nil || isNil(o.ElementType) {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *FormElement) HasElementType() bool {
	if o != nil && !isNil(o.ElementType) {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *FormElement) SetElementType(v string) {
	o.ElementType = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *FormElement) GetConfig() map[string]map[string]interface{} {
	if o == nil || isNil(o.Config) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetConfigOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Config) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *FormElement) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]map[string]interface{} and assigns it to the Config field.
func (o *FormElement) SetConfig(v map[string]map[string]interface{}) {
	o.Config = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FormElement) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FormElement) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FormElement) SetKey(v string) {
	o.Key = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormElement) GetValidations() []FormElementValidationsSet {
	if o == nil {
		var ret []FormElementValidationsSet
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormElement) GetValidationsOk() ([]FormElementValidationsSet, bool) {
	if o == nil || isNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *FormElement) HasValidations() bool {
	if o != nil && isNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []FormElementValidationsSet and assigns it to the Validations field.
func (o *FormElement) SetValidations(v []FormElementValidationsSet) {
	o.Validations = v
}

func (o FormElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ElementType) {
		toSerialize["elementType"] = o.ElementType
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.Validations != nil {
		toSerialize["validations"] = o.Validations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormElement) UnmarshalJSON(bytes []byte) (err error) {
	varFormElement := _FormElement{}

	if err = json.Unmarshal(bytes, &varFormElement); err == nil {
	*o = FormElement(varFormElement)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "elementType")
		delete(additionalProperties, "config")
		delete(additionalProperties, "key")
		delete(additionalProperties, "validations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormElement struct {
	value *FormElement
	isSet bool
}

func (v NullableFormElement) Get() *FormElement {
	return v.value
}

func (v *NullableFormElement) Set(val *FormElement) {
	v.value = val
	v.isSet = true
}

func (v NullableFormElement) IsSet() bool {
	return v.isSet
}

func (v *NullableFormElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormElement(val *FormElement) *NullableFormElement {
	return &NullableFormElement{value: val, isSet: true}
}

func (v NullableFormElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


