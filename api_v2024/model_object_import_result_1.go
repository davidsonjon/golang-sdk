/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the ObjectImportResult1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectImportResult1{}

// ObjectImportResult1 Response model for import of a single object.
type ObjectImportResult1 struct {
	// Informational messages returned from the target service on import.
	Infos []SpConfigMessage1 `json:"infos"`
	// Warning messages returned from the target service on import.
	Warnings []SpConfigMessage1 `json:"warnings"`
	// Error messages returned from the target service on import.
	Errors []SpConfigMessage1 `json:"errors"`
	// References to objects that were created or updated by the import.
	ImportedObjects []ImportObject `json:"importedObjects"`
	AdditionalProperties map[string]interface{}
}

type _ObjectImportResult1 ObjectImportResult1

// NewObjectImportResult1 instantiates a new ObjectImportResult1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectImportResult1(infos []SpConfigMessage1, warnings []SpConfigMessage1, errors []SpConfigMessage1, importedObjects []ImportObject) *ObjectImportResult1 {
	this := ObjectImportResult1{}
	this.Infos = infos
	this.Warnings = warnings
	this.Errors = errors
	this.ImportedObjects = importedObjects
	return &this
}

// NewObjectImportResult1WithDefaults instantiates a new ObjectImportResult1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectImportResult1WithDefaults() *ObjectImportResult1 {
	this := ObjectImportResult1{}
	return &this
}

// GetInfos returns the Infos field value
func (o *ObjectImportResult1) GetInfos() []SpConfigMessage1 {
	if o == nil {
		var ret []SpConfigMessage1
		return ret
	}

	return o.Infos
}

// GetInfosOk returns a tuple with the Infos field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult1) GetInfosOk() ([]SpConfigMessage1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Infos, true
}

// SetInfos sets field value
func (o *ObjectImportResult1) SetInfos(v []SpConfigMessage1) {
	o.Infos = v
}

// GetWarnings returns the Warnings field value
func (o *ObjectImportResult1) GetWarnings() []SpConfigMessage1 {
	if o == nil {
		var ret []SpConfigMessage1
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult1) GetWarningsOk() ([]SpConfigMessage1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *ObjectImportResult1) SetWarnings(v []SpConfigMessage1) {
	o.Warnings = v
}

// GetErrors returns the Errors field value
func (o *ObjectImportResult1) GetErrors() []SpConfigMessage1 {
	if o == nil {
		var ret []SpConfigMessage1
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult1) GetErrorsOk() ([]SpConfigMessage1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ObjectImportResult1) SetErrors(v []SpConfigMessage1) {
	o.Errors = v
}

// GetImportedObjects returns the ImportedObjects field value
func (o *ObjectImportResult1) GetImportedObjects() []ImportObject {
	if o == nil {
		var ret []ImportObject
		return ret
	}

	return o.ImportedObjects
}

// GetImportedObjectsOk returns a tuple with the ImportedObjects field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult1) GetImportedObjectsOk() ([]ImportObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImportedObjects, true
}

// SetImportedObjects sets field value
func (o *ObjectImportResult1) SetImportedObjects(v []ImportObject) {
	o.ImportedObjects = v
}

func (o ObjectImportResult1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectImportResult1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["infos"] = o.Infos
	toSerialize["warnings"] = o.Warnings
	toSerialize["errors"] = o.Errors
	toSerialize["importedObjects"] = o.ImportedObjects

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObjectImportResult1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"infos",
		"warnings",
		"errors",
		"importedObjects",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varObjectImportResult1 := _ObjectImportResult1{}

	if err = json.Unmarshal(bytes, &varObjectImportResult1); err == nil {
	*o = ObjectImportResult1(varObjectImportResult1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "infos")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "importedObjects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectImportResult1 struct {
	value *ObjectImportResult1
	isSet bool
}

func (v NullableObjectImportResult1) Get() *ObjectImportResult1 {
	return v.value
}

func (v *NullableObjectImportResult1) Set(val *ObjectImportResult1) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectImportResult1) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectImportResult1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectImportResult1(val *ObjectImportResult1) *NullableObjectImportResult1 {
	return &NullableObjectImportResult1{value: val, isSet: true}
}

func (v NullableObjectImportResult1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectImportResult1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

