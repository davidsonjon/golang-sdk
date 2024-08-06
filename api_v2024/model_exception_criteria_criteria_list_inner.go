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

// checks if the ExceptionCriteriaCriteriaListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionCriteriaCriteriaListInner{}

// ExceptionCriteriaCriteriaListInner The types of objects supported for SOD violations
type ExceptionCriteriaCriteriaListInner struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies
	Name *string `json:"name,omitempty"`
	// Whether the subject identity already had that access or not
	Existing *bool `json:"existing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExceptionCriteriaCriteriaListInner ExceptionCriteriaCriteriaListInner

// NewExceptionCriteriaCriteriaListInner instantiates a new ExceptionCriteriaCriteriaListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionCriteriaCriteriaListInner() *ExceptionCriteriaCriteriaListInner {
	this := ExceptionCriteriaCriteriaListInner{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// NewExceptionCriteriaCriteriaListInnerWithDefaults instantiates a new ExceptionCriteriaCriteriaListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionCriteriaCriteriaListInnerWithDefaults() *ExceptionCriteriaCriteriaListInner {
	this := ExceptionCriteriaCriteriaListInner{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExceptionCriteriaCriteriaListInner) GetType() map[string]interface{} {
	if o == nil || isNil(o.Type) {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaCriteriaListInner) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExceptionCriteriaCriteriaListInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *ExceptionCriteriaCriteriaListInner) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExceptionCriteriaCriteriaListInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaCriteriaListInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExceptionCriteriaCriteriaListInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExceptionCriteriaCriteriaListInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExceptionCriteriaCriteriaListInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaCriteriaListInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExceptionCriteriaCriteriaListInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExceptionCriteriaCriteriaListInner) SetName(v string) {
	o.Name = &v
}

// GetExisting returns the Existing field value if set, zero value otherwise.
func (o *ExceptionCriteriaCriteriaListInner) GetExisting() bool {
	if o == nil || isNil(o.Existing) {
		var ret bool
		return ret
	}
	return *o.Existing
}

// GetExistingOk returns a tuple with the Existing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaCriteriaListInner) GetExistingOk() (*bool, bool) {
	if o == nil || isNil(o.Existing) {
		return nil, false
	}
	return o.Existing, true
}

// HasExisting returns a boolean if a field has been set.
func (o *ExceptionCriteriaCriteriaListInner) HasExisting() bool {
	if o != nil && !isNil(o.Existing) {
		return true
	}

	return false
}

// SetExisting gets a reference to the given bool and assigns it to the Existing field.
func (o *ExceptionCriteriaCriteriaListInner) SetExisting(v bool) {
	o.Existing = &v
}

func (o ExceptionCriteriaCriteriaListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionCriteriaCriteriaListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Existing) {
		toSerialize["existing"] = o.Existing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExceptionCriteriaCriteriaListInner) UnmarshalJSON(bytes []byte) (err error) {
	varExceptionCriteriaCriteriaListInner := _ExceptionCriteriaCriteriaListInner{}

	if err = json.Unmarshal(bytes, &varExceptionCriteriaCriteriaListInner); err == nil {
	*o = ExceptionCriteriaCriteriaListInner(varExceptionCriteriaCriteriaListInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "existing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExceptionCriteriaCriteriaListInner struct {
	value *ExceptionCriteriaCriteriaListInner
	isSet bool
}

func (v NullableExceptionCriteriaCriteriaListInner) Get() *ExceptionCriteriaCriteriaListInner {
	return v.value
}

func (v *NullableExceptionCriteriaCriteriaListInner) Set(val *ExceptionCriteriaCriteriaListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionCriteriaCriteriaListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionCriteriaCriteriaListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionCriteriaCriteriaListInner(val *ExceptionCriteriaCriteriaListInner) *NullableExceptionCriteriaCriteriaListInner {
	return &NullableExceptionCriteriaCriteriaListInner{value: val, isSet: true}
}

func (v NullableExceptionCriteriaCriteriaListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionCriteriaCriteriaListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


