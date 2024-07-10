/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
)

// checks if the VendorConnectorMappingUpdatedAt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorConnectorMappingUpdatedAt{}

// VendorConnectorMappingUpdatedAt An object representing the nullable timestamp of the last update.
type VendorConnectorMappingUpdatedAt struct {
	// The timestamp when the mapping was last updated, represented in ISO 8601 format.
	Time *time.Time `json:"Time,omitempty"`
	// A flag indicating if the 'Time' field is set and valid.
	Valid *bool `json:"Valid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VendorConnectorMappingUpdatedAt VendorConnectorMappingUpdatedAt

// NewVendorConnectorMappingUpdatedAt instantiates a new VendorConnectorMappingUpdatedAt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorConnectorMappingUpdatedAt() *VendorConnectorMappingUpdatedAt {
	this := VendorConnectorMappingUpdatedAt{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// NewVendorConnectorMappingUpdatedAtWithDefaults instantiates a new VendorConnectorMappingUpdatedAt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorConnectorMappingUpdatedAtWithDefaults() *VendorConnectorMappingUpdatedAt {
	this := VendorConnectorMappingUpdatedAt{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *VendorConnectorMappingUpdatedAt) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingUpdatedAt) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *VendorConnectorMappingUpdatedAt) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *VendorConnectorMappingUpdatedAt) SetTime(v time.Time) {
	o.Time = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *VendorConnectorMappingUpdatedAt) GetValid() bool {
	if o == nil || isNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingUpdatedAt) GetValidOk() (*bool, bool) {
	if o == nil || isNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *VendorConnectorMappingUpdatedAt) HasValid() bool {
	if o != nil && !isNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *VendorConnectorMappingUpdatedAt) SetValid(v bool) {
	o.Valid = &v
}

func (o VendorConnectorMappingUpdatedAt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorConnectorMappingUpdatedAt) ToMap() (map[string]interface{}, error) {
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

func (o *VendorConnectorMappingUpdatedAt) UnmarshalJSON(bytes []byte) (err error) {
	varVendorConnectorMappingUpdatedAt := _VendorConnectorMappingUpdatedAt{}

	if err = json.Unmarshal(bytes, &varVendorConnectorMappingUpdatedAt); err == nil {
	*o = VendorConnectorMappingUpdatedAt(varVendorConnectorMappingUpdatedAt)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Time")
		delete(additionalProperties, "Valid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVendorConnectorMappingUpdatedAt struct {
	value *VendorConnectorMappingUpdatedAt
	isSet bool
}

func (v NullableVendorConnectorMappingUpdatedAt) Get() *VendorConnectorMappingUpdatedAt {
	return v.value
}

func (v *NullableVendorConnectorMappingUpdatedAt) Set(val *VendorConnectorMappingUpdatedAt) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorConnectorMappingUpdatedAt) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorConnectorMappingUpdatedAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorConnectorMappingUpdatedAt(val *VendorConnectorMappingUpdatedAt) *NullableVendorConnectorMappingUpdatedAt {
	return &NullableVendorConnectorMappingUpdatedAt{value: val, isSet: true}
}

func (v NullableVendorConnectorMappingUpdatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorConnectorMappingUpdatedAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


