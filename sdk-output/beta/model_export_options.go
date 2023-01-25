/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// ExportOptions struct for ExportOptions
type ExportOptions struct {
	// Object type names to be excluded from an sp-config export command.
	ExcludeTypes []string `json:"excludeTypes,omitempty"`
	// Object type names to be included in an sp-config export command. IncludeTypes takes precedence over excludeTypes.
	IncludeTypes []string `json:"includeTypes,omitempty"`
	// Additional options targeting specific objects related to each item in the includeTypes field
	ObjectOptions *map[string]ObjectExportImportOptions `json:"objectOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportOptions ExportOptions

// NewExportOptions instantiates a new ExportOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportOptions() *ExportOptions {
	this := ExportOptions{}
	return &this
}

// NewExportOptionsWithDefaults instantiates a new ExportOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportOptionsWithDefaults() *ExportOptions {
	this := ExportOptions{}
	return &this
}

// GetExcludeTypes returns the ExcludeTypes field value if set, zero value otherwise.
func (o *ExportOptions) GetExcludeTypes() []string {
	if o == nil || isNil(o.ExcludeTypes) {
		var ret []string
		return ret
	}
	return o.ExcludeTypes
}

// GetExcludeTypesOk returns a tuple with the ExcludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetExcludeTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeTypes) {
		return nil, false
	}
	return o.ExcludeTypes, true
}

// HasExcludeTypes returns a boolean if a field has been set.
func (o *ExportOptions) HasExcludeTypes() bool {
	if o != nil && !isNil(o.ExcludeTypes) {
		return true
	}

	return false
}

// SetExcludeTypes gets a reference to the given []string and assigns it to the ExcludeTypes field.
func (o *ExportOptions) SetExcludeTypes(v []string) {
	o.ExcludeTypes = v
}

// GetIncludeTypes returns the IncludeTypes field value if set, zero value otherwise.
func (o *ExportOptions) GetIncludeTypes() []string {
	if o == nil || isNil(o.IncludeTypes) {
		var ret []string
		return ret
	}
	return o.IncludeTypes
}

// GetIncludeTypesOk returns a tuple with the IncludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetIncludeTypesOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeTypes) {
		return nil, false
	}
	return o.IncludeTypes, true
}

// HasIncludeTypes returns a boolean if a field has been set.
func (o *ExportOptions) HasIncludeTypes() bool {
	if o != nil && !isNil(o.IncludeTypes) {
		return true
	}

	return false
}

// SetIncludeTypes gets a reference to the given []string and assigns it to the IncludeTypes field.
func (o *ExportOptions) SetIncludeTypes(v []string) {
	o.IncludeTypes = v
}

// GetObjectOptions returns the ObjectOptions field value if set, zero value otherwise.
func (o *ExportOptions) GetObjectOptions() map[string]ObjectExportImportOptions {
	if o == nil || isNil(o.ObjectOptions) {
		var ret map[string]ObjectExportImportOptions
		return ret
	}
	return *o.ObjectOptions
}

// GetObjectOptionsOk returns a tuple with the ObjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetObjectOptionsOk() (*map[string]ObjectExportImportOptions, bool) {
	if o == nil || isNil(o.ObjectOptions) {
		return nil, false
	}
	return o.ObjectOptions, true
}

// HasObjectOptions returns a boolean if a field has been set.
func (o *ExportOptions) HasObjectOptions() bool {
	if o != nil && !isNil(o.ObjectOptions) {
		return true
	}

	return false
}

// SetObjectOptions gets a reference to the given map[string]ObjectExportImportOptions and assigns it to the ObjectOptions field.
func (o *ExportOptions) SetObjectOptions(v map[string]ObjectExportImportOptions) {
	o.ObjectOptions = &v
}

func (o ExportOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExcludeTypes) {
		toSerialize["excludeTypes"] = o.ExcludeTypes
	}
	if !isNil(o.IncludeTypes) {
		toSerialize["includeTypes"] = o.IncludeTypes
	}
	if !isNil(o.ObjectOptions) {
		toSerialize["objectOptions"] = o.ObjectOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportOptions) UnmarshalJSON(bytes []byte) (err error) {
	varExportOptions := _ExportOptions{}

	if err = json.Unmarshal(bytes, &varExportOptions); err == nil {
		*o = ExportOptions(varExportOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "excludeTypes")
		delete(additionalProperties, "includeTypes")
		delete(additionalProperties, "objectOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportOptions struct {
	value *ExportOptions
	isSet bool
}

func (v NullableExportOptions) Get() *ExportOptions {
	return v.value
}

func (v *NullableExportOptions) Set(val *ExportOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExportOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExportOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportOptions(val *ExportOptions) *NullableExportOptions {
	return &NullableExportOptions{value: val, isSet: true}
}

func (v NullableExportOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


