/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the ExportPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportPayload{}

// ExportPayload struct for ExportPayload
type ExportPayload struct {
	// Optional user defined description/name for export job.
	Description *string `json:"description,omitempty"`
	// Object type names to be excluded from an sp-config export command.
	ExcludeTypes []string `json:"excludeTypes,omitempty"`
	// Object type names to be included in an sp-config export command. IncludeTypes takes precedence over excludeTypes.
	IncludeTypes []string `json:"includeTypes,omitempty"`
	// Additional options targeting specific objects related to each item in the includeTypes field
	ObjectOptions *map[string]ObjectExportImportOptions `json:"objectOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportPayload ExportPayload

// NewExportPayload instantiates a new ExportPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportPayload() *ExportPayload {
	this := ExportPayload{}
	return &this
}

// NewExportPayloadWithDefaults instantiates a new ExportPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportPayloadWithDefaults() *ExportPayload {
	this := ExportPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExportPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExportPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExportPayload) SetDescription(v string) {
	o.Description = &v
}

// GetExcludeTypes returns the ExcludeTypes field value if set, zero value otherwise.
func (o *ExportPayload) GetExcludeTypes() []string {
	if o == nil || IsNil(o.ExcludeTypes) {
		var ret []string
		return ret
	}
	return o.ExcludeTypes
}

// GetExcludeTypesOk returns a tuple with the ExcludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPayload) GetExcludeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeTypes) {
		return nil, false
	}
	return o.ExcludeTypes, true
}

// HasExcludeTypes returns a boolean if a field has been set.
func (o *ExportPayload) HasExcludeTypes() bool {
	if o != nil && !IsNil(o.ExcludeTypes) {
		return true
	}

	return false
}

// SetExcludeTypes gets a reference to the given []string and assigns it to the ExcludeTypes field.
func (o *ExportPayload) SetExcludeTypes(v []string) {
	o.ExcludeTypes = v
}

// GetIncludeTypes returns the IncludeTypes field value if set, zero value otherwise.
func (o *ExportPayload) GetIncludeTypes() []string {
	if o == nil || IsNil(o.IncludeTypes) {
		var ret []string
		return ret
	}
	return o.IncludeTypes
}

// GetIncludeTypesOk returns a tuple with the IncludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPayload) GetIncludeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeTypes) {
		return nil, false
	}
	return o.IncludeTypes, true
}

// HasIncludeTypes returns a boolean if a field has been set.
func (o *ExportPayload) HasIncludeTypes() bool {
	if o != nil && !IsNil(o.IncludeTypes) {
		return true
	}

	return false
}

// SetIncludeTypes gets a reference to the given []string and assigns it to the IncludeTypes field.
func (o *ExportPayload) SetIncludeTypes(v []string) {
	o.IncludeTypes = v
}

// GetObjectOptions returns the ObjectOptions field value if set, zero value otherwise.
func (o *ExportPayload) GetObjectOptions() map[string]ObjectExportImportOptions {
	if o == nil || IsNil(o.ObjectOptions) {
		var ret map[string]ObjectExportImportOptions
		return ret
	}
	return *o.ObjectOptions
}

// GetObjectOptionsOk returns a tuple with the ObjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPayload) GetObjectOptionsOk() (*map[string]ObjectExportImportOptions, bool) {
	if o == nil || IsNil(o.ObjectOptions) {
		return nil, false
	}
	return o.ObjectOptions, true
}

// HasObjectOptions returns a boolean if a field has been set.
func (o *ExportPayload) HasObjectOptions() bool {
	if o != nil && !IsNil(o.ObjectOptions) {
		return true
	}

	return false
}

// SetObjectOptions gets a reference to the given map[string]ObjectExportImportOptions and assigns it to the ObjectOptions field.
func (o *ExportPayload) SetObjectOptions(v map[string]ObjectExportImportOptions) {
	o.ObjectOptions = &v
}

func (o ExportPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExcludeTypes) {
		toSerialize["excludeTypes"] = o.ExcludeTypes
	}
	if !IsNil(o.IncludeTypes) {
		toSerialize["includeTypes"] = o.IncludeTypes
	}
	if !IsNil(o.ObjectOptions) {
		toSerialize["objectOptions"] = o.ObjectOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExportPayload) UnmarshalJSON(data []byte) (err error) {
	varExportPayload := _ExportPayload{}

	err = json.Unmarshal(data, &varExportPayload)

	if err != nil {
		return err
	}

	*o = ExportPayload(varExportPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "excludeTypes")
		delete(additionalProperties, "includeTypes")
		delete(additionalProperties, "objectOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportPayload struct {
	value *ExportPayload
	isSet bool
}

func (v NullableExportPayload) Get() *ExportPayload {
	return v.value
}

func (v *NullableExportPayload) Set(val *ExportPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableExportPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableExportPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportPayload(val *ExportPayload) *NullableExportPayload {
	return &NullableExportPayload{value: val, isSet: true}
}

func (v NullableExportPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


