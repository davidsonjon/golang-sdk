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

// checks if the BaseSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseSegment{}

// BaseSegment struct for BaseSegment
type BaseSegment struct {
	// Segment's unique ID.
	Id *string `json:"id,omitempty"`
	// Segment's display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseSegment BaseSegment

// NewBaseSegment instantiates a new BaseSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSegment() *BaseSegment {
	this := BaseSegment{}
	return &this
}

// NewBaseSegmentWithDefaults instantiates a new BaseSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSegmentWithDefaults() *BaseSegment {
	this := BaseSegment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseSegment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSegment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseSegment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseSegment) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseSegment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSegment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseSegment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseSegment) SetName(v string) {
	o.Name = &v
}

func (o BaseSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *BaseSegment) UnmarshalJSON(data []byte) (err error) {
	varBaseSegment := _BaseSegment{}

	err = json.Unmarshal(data, &varBaseSegment)

	if err != nil {
		return err
	}

	*o = BaseSegment(varBaseSegment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseSegment struct {
	value *BaseSegment
	isSet bool
}

func (v NullableBaseSegment) Get() *BaseSegment {
	return v.value
}

func (v *NullableBaseSegment) Set(val *BaseSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSegment(val *BaseSegment) *NullableBaseSegment {
	return &NullableBaseSegment{value: val, isSet: true}
}

func (v NullableBaseSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


