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

// checks if the OwnerReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerReference{}

// OwnerReference The owner of this object.
type OwnerReference struct {
	// Owner type. This field must be either left null or set to 'IDENTITY' on input, otherwise a 400 Bad Request error will result.
	Type *string `json:"type,omitempty"`
	// Identity id
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the owner. It may be left null or omitted in a POST or PATCH. If set, it must match the current value of the owner's display name, otherwise a 400 Bad Request error will result.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OwnerReference OwnerReference

// NewOwnerReference instantiates a new OwnerReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerReference() *OwnerReference {
	this := OwnerReference{}
	return &this
}

// NewOwnerReferenceWithDefaults instantiates a new OwnerReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerReferenceWithDefaults() *OwnerReference {
	this := OwnerReference{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OwnerReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OwnerReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OwnerReference) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OwnerReference) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReference) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OwnerReference) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OwnerReference) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OwnerReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OwnerReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OwnerReference) SetName(v string) {
	o.Name = &v
}

func (o OwnerReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerReference) ToMap() (map[string]interface{}, error) {
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

func (o *OwnerReference) UnmarshalJSON(data []byte) (err error) {
	varOwnerReference := _OwnerReference{}

	err = json.Unmarshal(data, &varOwnerReference)

	if err != nil {
		return err
	}

	*o = OwnerReference(varOwnerReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOwnerReference struct {
	value *OwnerReference
	isSet bool
}

func (v NullableOwnerReference) Get() *OwnerReference {
	return v.value
}

func (v *NullableOwnerReference) Set(val *OwnerReference) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerReference) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerReference(val *OwnerReference) *NullableOwnerReference {
	return &NullableOwnerReference{value: val, isSet: true}
}

func (v NullableOwnerReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


