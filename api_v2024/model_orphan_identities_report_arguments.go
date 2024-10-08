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

// checks if the OrphanIdentitiesReportArguments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrphanIdentitiesReportArguments{}

// OrphanIdentitiesReportArguments Arguments for Orphan Identities report (ORPHAN_IDENTITIES)
type OrphanIdentitiesReportArguments struct {
	// Output report file formats. These are formats for calling GET endpoint as query parameter 'fileFormat'.  In case report won't have this argument there will be ['CSV', 'PDF'] as default.
	SelectedFormats []string `json:"selectedFormats,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrphanIdentitiesReportArguments OrphanIdentitiesReportArguments

// NewOrphanIdentitiesReportArguments instantiates a new OrphanIdentitiesReportArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrphanIdentitiesReportArguments() *OrphanIdentitiesReportArguments {
	this := OrphanIdentitiesReportArguments{}
	return &this
}

// NewOrphanIdentitiesReportArgumentsWithDefaults instantiates a new OrphanIdentitiesReportArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrphanIdentitiesReportArgumentsWithDefaults() *OrphanIdentitiesReportArguments {
	this := OrphanIdentitiesReportArguments{}
	return &this
}

// GetSelectedFormats returns the SelectedFormats field value if set, zero value otherwise.
func (o *OrphanIdentitiesReportArguments) GetSelectedFormats() []string {
	if o == nil || IsNil(o.SelectedFormats) {
		var ret []string
		return ret
	}
	return o.SelectedFormats
}

// GetSelectedFormatsOk returns a tuple with the SelectedFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrphanIdentitiesReportArguments) GetSelectedFormatsOk() ([]string, bool) {
	if o == nil || IsNil(o.SelectedFormats) {
		return nil, false
	}
	return o.SelectedFormats, true
}

// HasSelectedFormats returns a boolean if a field has been set.
func (o *OrphanIdentitiesReportArguments) HasSelectedFormats() bool {
	if o != nil && !IsNil(o.SelectedFormats) {
		return true
	}

	return false
}

// SetSelectedFormats gets a reference to the given []string and assigns it to the SelectedFormats field.
func (o *OrphanIdentitiesReportArguments) SetSelectedFormats(v []string) {
	o.SelectedFormats = v
}

func (o OrphanIdentitiesReportArguments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrphanIdentitiesReportArguments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SelectedFormats) {
		toSerialize["selectedFormats"] = o.SelectedFormats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrphanIdentitiesReportArguments) UnmarshalJSON(data []byte) (err error) {
	varOrphanIdentitiesReportArguments := _OrphanIdentitiesReportArguments{}

	err = json.Unmarshal(data, &varOrphanIdentitiesReportArguments)

	if err != nil {
		return err
	}

	*o = OrphanIdentitiesReportArguments(varOrphanIdentitiesReportArguments)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selectedFormats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrphanIdentitiesReportArguments struct {
	value *OrphanIdentitiesReportArguments
	isSet bool
}

func (v NullableOrphanIdentitiesReportArguments) Get() *OrphanIdentitiesReportArguments {
	return v.value
}

func (v *NullableOrphanIdentitiesReportArguments) Set(val *OrphanIdentitiesReportArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableOrphanIdentitiesReportArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableOrphanIdentitiesReportArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrphanIdentitiesReportArguments(val *OrphanIdentitiesReportArguments) *NullableOrphanIdentitiesReportArguments {
	return &NullableOrphanIdentitiesReportArguments{value: val, isSet: true}
}

func (v NullableOrphanIdentitiesReportArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrphanIdentitiesReportArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


