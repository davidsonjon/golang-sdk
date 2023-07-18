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

// checks if the ImportFormDefinitions202Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportFormDefinitions202Response{}

// ImportFormDefinitions202Response struct for ImportFormDefinitions202Response
type ImportFormDefinitions202Response struct {
	Errors []ImportFormDefinitions202ResponseErrorsInner `json:"errors,omitempty"`
	ImportedObjects []ExportFormDefinitionsByTenant200ResponseInner `json:"importedObjects,omitempty"`
	Infos []ImportFormDefinitions202ResponseErrorsInner `json:"infos,omitempty"`
	Warnings []ImportFormDefinitions202ResponseErrorsInner `json:"warnings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportFormDefinitions202Response ImportFormDefinitions202Response

// NewImportFormDefinitions202Response instantiates a new ImportFormDefinitions202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportFormDefinitions202Response() *ImportFormDefinitions202Response {
	this := ImportFormDefinitions202Response{}
	return &this
}

// NewImportFormDefinitions202ResponseWithDefaults instantiates a new ImportFormDefinitions202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportFormDefinitions202ResponseWithDefaults() *ImportFormDefinitions202Response {
	this := ImportFormDefinitions202Response{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ImportFormDefinitions202Response) GetErrors() []ImportFormDefinitions202ResponseErrorsInner {
	if o == nil || isNil(o.Errors) {
		var ret []ImportFormDefinitions202ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFormDefinitions202Response) GetErrorsOk() ([]ImportFormDefinitions202ResponseErrorsInner, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ImportFormDefinitions202Response) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ImportFormDefinitions202ResponseErrorsInner and assigns it to the Errors field.
func (o *ImportFormDefinitions202Response) SetErrors(v []ImportFormDefinitions202ResponseErrorsInner) {
	o.Errors = v
}

// GetImportedObjects returns the ImportedObjects field value if set, zero value otherwise.
func (o *ImportFormDefinitions202Response) GetImportedObjects() []ExportFormDefinitionsByTenant200ResponseInner {
	if o == nil || isNil(o.ImportedObjects) {
		var ret []ExportFormDefinitionsByTenant200ResponseInner
		return ret
	}
	return o.ImportedObjects
}

// GetImportedObjectsOk returns a tuple with the ImportedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFormDefinitions202Response) GetImportedObjectsOk() ([]ExportFormDefinitionsByTenant200ResponseInner, bool) {
	if o == nil || isNil(o.ImportedObjects) {
		return nil, false
	}
	return o.ImportedObjects, true
}

// HasImportedObjects returns a boolean if a field has been set.
func (o *ImportFormDefinitions202Response) HasImportedObjects() bool {
	if o != nil && !isNil(o.ImportedObjects) {
		return true
	}

	return false
}

// SetImportedObjects gets a reference to the given []ExportFormDefinitionsByTenant200ResponseInner and assigns it to the ImportedObjects field.
func (o *ImportFormDefinitions202Response) SetImportedObjects(v []ExportFormDefinitionsByTenant200ResponseInner) {
	o.ImportedObjects = v
}

// GetInfos returns the Infos field value if set, zero value otherwise.
func (o *ImportFormDefinitions202Response) GetInfos() []ImportFormDefinitions202ResponseErrorsInner {
	if o == nil || isNil(o.Infos) {
		var ret []ImportFormDefinitions202ResponseErrorsInner
		return ret
	}
	return o.Infos
}

// GetInfosOk returns a tuple with the Infos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFormDefinitions202Response) GetInfosOk() ([]ImportFormDefinitions202ResponseErrorsInner, bool) {
	if o == nil || isNil(o.Infos) {
		return nil, false
	}
	return o.Infos, true
}

// HasInfos returns a boolean if a field has been set.
func (o *ImportFormDefinitions202Response) HasInfos() bool {
	if o != nil && !isNil(o.Infos) {
		return true
	}

	return false
}

// SetInfos gets a reference to the given []ImportFormDefinitions202ResponseErrorsInner and assigns it to the Infos field.
func (o *ImportFormDefinitions202Response) SetInfos(v []ImportFormDefinitions202ResponseErrorsInner) {
	o.Infos = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ImportFormDefinitions202Response) GetWarnings() []ImportFormDefinitions202ResponseErrorsInner {
	if o == nil || isNil(o.Warnings) {
		var ret []ImportFormDefinitions202ResponseErrorsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFormDefinitions202Response) GetWarningsOk() ([]ImportFormDefinitions202ResponseErrorsInner, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ImportFormDefinitions202Response) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []ImportFormDefinitions202ResponseErrorsInner and assigns it to the Warnings field.
func (o *ImportFormDefinitions202Response) SetWarnings(v []ImportFormDefinitions202ResponseErrorsInner) {
	o.Warnings = v
}

func (o ImportFormDefinitions202Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportFormDefinitions202Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.ImportedObjects) {
		toSerialize["importedObjects"] = o.ImportedObjects
	}
	if !isNil(o.Infos) {
		toSerialize["infos"] = o.Infos
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportFormDefinitions202Response) UnmarshalJSON(bytes []byte) (err error) {
	varImportFormDefinitions202Response := _ImportFormDefinitions202Response{}

	if err = json.Unmarshal(bytes, &varImportFormDefinitions202Response); err == nil {
		*o = ImportFormDefinitions202Response(varImportFormDefinitions202Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		delete(additionalProperties, "importedObjects")
		delete(additionalProperties, "infos")
		delete(additionalProperties, "warnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportFormDefinitions202Response struct {
	value *ImportFormDefinitions202Response
	isSet bool
}

func (v NullableImportFormDefinitions202Response) Get() *ImportFormDefinitions202Response {
	return v.value
}

func (v *NullableImportFormDefinitions202Response) Set(val *ImportFormDefinitions202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableImportFormDefinitions202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableImportFormDefinitions202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportFormDefinitions202Response(val *ImportFormDefinitions202Response) *NullableImportFormDefinitions202Response {
	return &NullableImportFormDefinitions202Response{value: val, isSet: true}
}

func (v NullableImportFormDefinitions202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportFormDefinitions202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


