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

// checks if the PatchPotentialRoleRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPotentialRoleRequestInner{}

// PatchPotentialRoleRequestInner struct for PatchPotentialRoleRequestInner
type PatchPotentialRoleRequestInner struct {
	// The operation to be performed
	Op *string `json:"op,omitempty"`
	// A string JSON Pointer representing the target path to an element to be affected by the operation
	Path string `json:"path"`
	Value *JsonPatchOperationValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchPotentialRoleRequestInner PatchPotentialRoleRequestInner

// NewPatchPotentialRoleRequestInner instantiates a new PatchPotentialRoleRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPotentialRoleRequestInner(op string, path string) *PatchPotentialRoleRequestInner {
	this := PatchPotentialRoleRequestInner{}
	this.Op = &op
	this.Path = path
	return &this
}

// NewPatchPotentialRoleRequestInnerWithDefaults instantiates a new PatchPotentialRoleRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPotentialRoleRequestInnerWithDefaults() *PatchPotentialRoleRequestInner {
	this := PatchPotentialRoleRequestInner{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *PatchPotentialRoleRequestInner) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPotentialRoleRequestInner) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *PatchPotentialRoleRequestInner) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *PatchPotentialRoleRequestInner) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value
func (o *PatchPotentialRoleRequestInner) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchPotentialRoleRequestInner) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchPotentialRoleRequestInner) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchPotentialRoleRequestInner) GetValue() JsonPatchOperationValue {
	if o == nil || IsNil(o.Value) {
		var ret JsonPatchOperationValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPotentialRoleRequestInner) GetValueOk() (*JsonPatchOperationValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchPotentialRoleRequestInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given JsonPatchOperationValue and assigns it to the Value field.
func (o *PatchPotentialRoleRequestInner) SetValue(v JsonPatchOperationValue) {
	o.Value = &v
}

func (o PatchPotentialRoleRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPotentialRoleRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchPotentialRoleRequestInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
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

	varPatchPotentialRoleRequestInner := _PatchPotentialRoleRequestInner{}

	err = json.Unmarshal(data, &varPatchPotentialRoleRequestInner)

	if err != nil {
		return err
	}

	*o = PatchPotentialRoleRequestInner(varPatchPotentialRoleRequestInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchPotentialRoleRequestInner struct {
	value *PatchPotentialRoleRequestInner
	isSet bool
}

func (v NullablePatchPotentialRoleRequestInner) Get() *PatchPotentialRoleRequestInner {
	return v.value
}

func (v *NullablePatchPotentialRoleRequestInner) Set(val *PatchPotentialRoleRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPotentialRoleRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPotentialRoleRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPotentialRoleRequestInner(val *PatchPotentialRoleRequestInner) *NullablePatchPotentialRoleRequestInner {
	return &NullablePatchPotentialRoleRequestInner{value: val, isSet: true}
}

func (v NullablePatchPotentialRoleRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPotentialRoleRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


