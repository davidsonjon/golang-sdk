/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the VendorConnectorMappingDeletedAt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorConnectorMappingDeletedAt{}

// VendorConnectorMappingDeletedAt An object representing the nullable timestamp of when the mapping was deleted.
type VendorConnectorMappingDeletedAt struct {
	// The timestamp when the mapping was deleted, represented in ISO 8601 format, if applicable.
	Time *time.Time `json:"Time,omitempty"`
	// A flag indicating if the 'Time' field is set and valid, i.e., if the mapping has been deleted.
	Valid *bool `json:"Valid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VendorConnectorMappingDeletedAt VendorConnectorMappingDeletedAt

// NewVendorConnectorMappingDeletedAt instantiates a new VendorConnectorMappingDeletedAt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorConnectorMappingDeletedAt() *VendorConnectorMappingDeletedAt {
	this := VendorConnectorMappingDeletedAt{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// NewVendorConnectorMappingDeletedAtWithDefaults instantiates a new VendorConnectorMappingDeletedAt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorConnectorMappingDeletedAtWithDefaults() *VendorConnectorMappingDeletedAt {
	this := VendorConnectorMappingDeletedAt{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *VendorConnectorMappingDeletedAt) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingDeletedAt) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *VendorConnectorMappingDeletedAt) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *VendorConnectorMappingDeletedAt) SetTime(v time.Time) {
	o.Time = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *VendorConnectorMappingDeletedAt) GetValid() bool {
	if o == nil || isNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingDeletedAt) GetValidOk() (*bool, bool) {
	if o == nil || isNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *VendorConnectorMappingDeletedAt) HasValid() bool {
	if o != nil && !isNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *VendorConnectorMappingDeletedAt) SetValid(v bool) {
	o.Valid = &v
}

func (o VendorConnectorMappingDeletedAt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorConnectorMappingDeletedAt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["Time"] = o.Time
	}
	if !isNil(o.Valid) {
		toSerialize["Valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VendorConnectorMappingDeletedAt) UnmarshalJSON(bytes []byte) (err error) {
	varVendorConnectorMappingDeletedAt := _VendorConnectorMappingDeletedAt{}

	if err = json.Unmarshal(bytes, &varVendorConnectorMappingDeletedAt); err == nil {
	*o = VendorConnectorMappingDeletedAt(varVendorConnectorMappingDeletedAt)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Time")
		delete(additionalProperties, "Valid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVendorConnectorMappingDeletedAt struct {
	value *VendorConnectorMappingDeletedAt
	isSet bool
}

func (v NullableVendorConnectorMappingDeletedAt) Get() *VendorConnectorMappingDeletedAt {
	return v.value
}

func (v *NullableVendorConnectorMappingDeletedAt) Set(val *VendorConnectorMappingDeletedAt) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorConnectorMappingDeletedAt) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorConnectorMappingDeletedAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorConnectorMappingDeletedAt(val *VendorConnectorMappingDeletedAt) *NullableVendorConnectorMappingDeletedAt {
	return &NullableVendorConnectorMappingDeletedAt{value: val, isSet: true}
}

func (v NullableVendorConnectorMappingDeletedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorConnectorMappingDeletedAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


