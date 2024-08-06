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

// checks if the AccessCriteriaCriteriaListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessCriteriaCriteriaListInner{}

// AccessCriteriaCriteriaListInner struct for AccessCriteriaCriteriaListInner
type AccessCriteriaCriteriaListInner struct {
	// Type of the propery to which this reference applies to
	Type *string `json:"type,omitempty"`
	// ID of the object to which this reference applies to
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies to
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessCriteriaCriteriaListInner AccessCriteriaCriteriaListInner

// NewAccessCriteriaCriteriaListInner instantiates a new AccessCriteriaCriteriaListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessCriteriaCriteriaListInner() *AccessCriteriaCriteriaListInner {
	this := AccessCriteriaCriteriaListInner{}
	return &this
}

// NewAccessCriteriaCriteriaListInnerWithDefaults instantiates a new AccessCriteriaCriteriaListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessCriteriaCriteriaListInnerWithDefaults() *AccessCriteriaCriteriaListInner {
	this := AccessCriteriaCriteriaListInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessCriteriaCriteriaListInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCriteriaCriteriaListInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessCriteriaCriteriaListInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessCriteriaCriteriaListInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessCriteriaCriteriaListInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCriteriaCriteriaListInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessCriteriaCriteriaListInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessCriteriaCriteriaListInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessCriteriaCriteriaListInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCriteriaCriteriaListInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessCriteriaCriteriaListInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessCriteriaCriteriaListInner) SetName(v string) {
	o.Name = &v
}

func (o AccessCriteriaCriteriaListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessCriteriaCriteriaListInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessCriteriaCriteriaListInner) UnmarshalJSON(bytes []byte) (err error) {
	varAccessCriteriaCriteriaListInner := _AccessCriteriaCriteriaListInner{}

	if err = json.Unmarshal(bytes, &varAccessCriteriaCriteriaListInner); err == nil {
	*o = AccessCriteriaCriteriaListInner(varAccessCriteriaCriteriaListInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessCriteriaCriteriaListInner struct {
	value *AccessCriteriaCriteriaListInner
	isSet bool
}

func (v NullableAccessCriteriaCriteriaListInner) Get() *AccessCriteriaCriteriaListInner {
	return v.value
}

func (v *NullableAccessCriteriaCriteriaListInner) Set(val *AccessCriteriaCriteriaListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessCriteriaCriteriaListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessCriteriaCriteriaListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessCriteriaCriteriaListInner(val *AccessCriteriaCriteriaListInner) *NullableAccessCriteriaCriteriaListInner {
	return &NullableAccessCriteriaCriteriaListInner{value: val, isSet: true}
}

func (v NullableAccessCriteriaCriteriaListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessCriteriaCriteriaListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


