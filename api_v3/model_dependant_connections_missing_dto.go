/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the DependantConnectionsMissingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DependantConnectionsMissingDto{}

// DependantConnectionsMissingDto struct for DependantConnectionsMissingDto
type DependantConnectionsMissingDto struct {
	// The type of dependency type that is missing in the SourceConnections
	DependencyType *string `json:"dependencyType,omitempty"`
	// The reason why this dependency is missing
	Reason *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DependantConnectionsMissingDto DependantConnectionsMissingDto

// NewDependantConnectionsMissingDto instantiates a new DependantConnectionsMissingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependantConnectionsMissingDto() *DependantConnectionsMissingDto {
	this := DependantConnectionsMissingDto{}
	return &this
}

// NewDependantConnectionsMissingDtoWithDefaults instantiates a new DependantConnectionsMissingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependantConnectionsMissingDtoWithDefaults() *DependantConnectionsMissingDto {
	this := DependantConnectionsMissingDto{}
	return &this
}

// GetDependencyType returns the DependencyType field value if set, zero value otherwise.
func (o *DependantConnectionsMissingDto) GetDependencyType() string {
	if o == nil || IsNil(o.DependencyType) {
		var ret string
		return ret
	}
	return *o.DependencyType
}

// GetDependencyTypeOk returns a tuple with the DependencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DependantConnectionsMissingDto) GetDependencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DependencyType) {
		return nil, false
	}
	return o.DependencyType, true
}

// HasDependencyType returns a boolean if a field has been set.
func (o *DependantConnectionsMissingDto) HasDependencyType() bool {
	if o != nil && !IsNil(o.DependencyType) {
		return true
	}

	return false
}

// SetDependencyType gets a reference to the given string and assigns it to the DependencyType field.
func (o *DependantConnectionsMissingDto) SetDependencyType(v string) {
	o.DependencyType = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DependantConnectionsMissingDto) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DependantConnectionsMissingDto) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DependantConnectionsMissingDto) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DependantConnectionsMissingDto) SetReason(v string) {
	o.Reason = &v
}

func (o DependantConnectionsMissingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DependantConnectionsMissingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DependencyType) {
		toSerialize["dependencyType"] = o.DependencyType
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DependantConnectionsMissingDto) UnmarshalJSON(data []byte) (err error) {
	varDependantConnectionsMissingDto := _DependantConnectionsMissingDto{}

	err = json.Unmarshal(data, &varDependantConnectionsMissingDto)

	if err != nil {
		return err
	}

	*o = DependantConnectionsMissingDto(varDependantConnectionsMissingDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dependencyType")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDependantConnectionsMissingDto struct {
	value *DependantConnectionsMissingDto
	isSet bool
}

func (v NullableDependantConnectionsMissingDto) Get() *DependantConnectionsMissingDto {
	return v.value
}

func (v *NullableDependantConnectionsMissingDto) Set(val *DependantConnectionsMissingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDependantConnectionsMissingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDependantConnectionsMissingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependantConnectionsMissingDto(val *DependantConnectionsMissingDto) *NullableDependantConnectionsMissingDto {
	return &NullableDependantConnectionsMissingDto{value: val, isSet: true}
}

func (v NullableDependantConnectionsMissingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependantConnectionsMissingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


