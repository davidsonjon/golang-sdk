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

// checks if the BackupOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupOptions{}

// BackupOptions Backup options control what will be included in the backup.
type BackupOptions struct {
	// Object type names to be included in a Configuration Hub backup command.
	IncludeTypes []string `json:"includeTypes,omitempty"`
	// Additional options targeting specific objects related to each item in the includeTypes field.
	ObjectOptions *map[string]ObjectExportImportNames `json:"objectOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BackupOptions BackupOptions

// NewBackupOptions instantiates a new BackupOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupOptions() *BackupOptions {
	this := BackupOptions{}
	return &this
}

// NewBackupOptionsWithDefaults instantiates a new BackupOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupOptionsWithDefaults() *BackupOptions {
	this := BackupOptions{}
	return &this
}

// GetIncludeTypes returns the IncludeTypes field value if set, zero value otherwise.
func (o *BackupOptions) GetIncludeTypes() []string {
	if o == nil || IsNil(o.IncludeTypes) {
		var ret []string
		return ret
	}
	return o.IncludeTypes
}

// GetIncludeTypesOk returns a tuple with the IncludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptions) GetIncludeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeTypes) {
		return nil, false
	}
	return o.IncludeTypes, true
}

// HasIncludeTypes returns a boolean if a field has been set.
func (o *BackupOptions) HasIncludeTypes() bool {
	if o != nil && !IsNil(o.IncludeTypes) {
		return true
	}

	return false
}

// SetIncludeTypes gets a reference to the given []string and assigns it to the IncludeTypes field.
func (o *BackupOptions) SetIncludeTypes(v []string) {
	o.IncludeTypes = v
}

// GetObjectOptions returns the ObjectOptions field value if set, zero value otherwise.
func (o *BackupOptions) GetObjectOptions() map[string]ObjectExportImportNames {
	if o == nil || IsNil(o.ObjectOptions) {
		var ret map[string]ObjectExportImportNames
		return ret
	}
	return *o.ObjectOptions
}

// GetObjectOptionsOk returns a tuple with the ObjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptions) GetObjectOptionsOk() (*map[string]ObjectExportImportNames, bool) {
	if o == nil || IsNil(o.ObjectOptions) {
		return nil, false
	}
	return o.ObjectOptions, true
}

// HasObjectOptions returns a boolean if a field has been set.
func (o *BackupOptions) HasObjectOptions() bool {
	if o != nil && !IsNil(o.ObjectOptions) {
		return true
	}

	return false
}

// SetObjectOptions gets a reference to the given map[string]ObjectExportImportNames and assigns it to the ObjectOptions field.
func (o *BackupOptions) SetObjectOptions(v map[string]ObjectExportImportNames) {
	o.ObjectOptions = &v
}

func (o BackupOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *BackupOptions) UnmarshalJSON(data []byte) (err error) {
	varBackupOptions := _BackupOptions{}

	err = json.Unmarshal(data, &varBackupOptions)

	if err != nil {
		return err
	}

	*o = BackupOptions(varBackupOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "includeTypes")
		delete(additionalProperties, "objectOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBackupOptions struct {
	value *BackupOptions
	isSet bool
}

func (v NullableBackupOptions) Get() *BackupOptions {
	return v.value
}

func (v *NullableBackupOptions) Set(val *BackupOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupOptions(val *BackupOptions) *NullableBackupOptions {
	return &NullableBackupOptions{value: val, isSet: true}
}

func (v NullableBackupOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

