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

// checks if the VendorConnectorMappingDeletedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorConnectorMappingDeletedBy{}

// VendorConnectorMappingDeletedBy An object representing the nullable identifier of the user who deleted the mapping.
type VendorConnectorMappingDeletedBy struct {
	// The identifier of the user who deleted the mapping, if applicable.
	String *string `json:"String,omitempty"`
	// A flag indicating if the 'String' field is set and valid, i.e., if the mapping has been deleted.
	Valid *bool `json:"Valid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VendorConnectorMappingDeletedBy VendorConnectorMappingDeletedBy

// NewVendorConnectorMappingDeletedBy instantiates a new VendorConnectorMappingDeletedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorConnectorMappingDeletedBy() *VendorConnectorMappingDeletedBy {
	this := VendorConnectorMappingDeletedBy{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// NewVendorConnectorMappingDeletedByWithDefaults instantiates a new VendorConnectorMappingDeletedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorConnectorMappingDeletedByWithDefaults() *VendorConnectorMappingDeletedBy {
	this := VendorConnectorMappingDeletedBy{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *VendorConnectorMappingDeletedBy) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingDeletedBy) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *VendorConnectorMappingDeletedBy) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *VendorConnectorMappingDeletedBy) SetString(v string) {
	o.String = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *VendorConnectorMappingDeletedBy) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingDeletedBy) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *VendorConnectorMappingDeletedBy) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *VendorConnectorMappingDeletedBy) SetValid(v bool) {
	o.Valid = &v
}

func (o VendorConnectorMappingDeletedBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorConnectorMappingDeletedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["String"] = o.String
	}
	if !IsNil(o.Valid) {
		toSerialize["Valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VendorConnectorMappingDeletedBy) UnmarshalJSON(data []byte) (err error) {
	varVendorConnectorMappingDeletedBy := _VendorConnectorMappingDeletedBy{}

	err = json.Unmarshal(data, &varVendorConnectorMappingDeletedBy)

	if err != nil {
		return err
	}

	*o = VendorConnectorMappingDeletedBy(varVendorConnectorMappingDeletedBy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "String")
		delete(additionalProperties, "Valid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVendorConnectorMappingDeletedBy struct {
	value *VendorConnectorMappingDeletedBy
	isSet bool
}

func (v NullableVendorConnectorMappingDeletedBy) Get() *VendorConnectorMappingDeletedBy {
	return v.value
}

func (v *NullableVendorConnectorMappingDeletedBy) Set(val *VendorConnectorMappingDeletedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorConnectorMappingDeletedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorConnectorMappingDeletedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorConnectorMappingDeletedBy(val *VendorConnectorMappingDeletedBy) *NullableVendorConnectorMappingDeletedBy {
	return &NullableVendorConnectorMappingDeletedBy{value: val, isSet: true}
}

func (v NullableVendorConnectorMappingDeletedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorConnectorMappingDeletedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


