/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the ObjectImportResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectImportResult{}

// ObjectImportResult Response model for import of a single object.
type ObjectImportResult struct {
	// Informational messages returned from the target service on import.
	Infos []SpConfigMessage `json:"infos"`
	// Warning messages returned from the target service on import.
	Warnings []SpConfigMessage `json:"warnings"`
	// Error messages returned from the target service on import.
	Errors []SpConfigMessage `json:"errors"`
	// References to objects that were created or updated by the import.
	ImportedObjects []ImportObject `json:"importedObjects"`
	AdditionalProperties map[string]interface{}
}

type _ObjectImportResult ObjectImportResult

// NewObjectImportResult instantiates a new ObjectImportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectImportResult(infos []SpConfigMessage, warnings []SpConfigMessage, errors []SpConfigMessage, importedObjects []ImportObject) *ObjectImportResult {
	this := ObjectImportResult{}
	this.Infos = infos
	this.Warnings = warnings
	this.Errors = errors
	this.ImportedObjects = importedObjects
	return &this
}

// NewObjectImportResultWithDefaults instantiates a new ObjectImportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectImportResultWithDefaults() *ObjectImportResult {
	this := ObjectImportResult{}
	return &this
}

// GetInfos returns the Infos field value
func (o *ObjectImportResult) GetInfos() []SpConfigMessage {
	if o == nil {
		var ret []SpConfigMessage
		return ret
	}

	return o.Infos
}

// GetInfosOk returns a tuple with the Infos field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult) GetInfosOk() ([]SpConfigMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Infos, true
}

// SetInfos sets field value
func (o *ObjectImportResult) SetInfos(v []SpConfigMessage) {
	o.Infos = v
}

// GetWarnings returns the Warnings field value
func (o *ObjectImportResult) GetWarnings() []SpConfigMessage {
	if o == nil {
		var ret []SpConfigMessage
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult) GetWarningsOk() ([]SpConfigMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *ObjectImportResult) SetWarnings(v []SpConfigMessage) {
	o.Warnings = v
}

// GetErrors returns the Errors field value
func (o *ObjectImportResult) GetErrors() []SpConfigMessage {
	if o == nil {
		var ret []SpConfigMessage
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult) GetErrorsOk() ([]SpConfigMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ObjectImportResult) SetErrors(v []SpConfigMessage) {
	o.Errors = v
}

// GetImportedObjects returns the ImportedObjects field value
func (o *ObjectImportResult) GetImportedObjects() []ImportObject {
	if o == nil {
		var ret []ImportObject
		return ret
	}

	return o.ImportedObjects
}

// GetImportedObjectsOk returns a tuple with the ImportedObjects field value
// and a boolean to check if the value has been set.
func (o *ObjectImportResult) GetImportedObjectsOk() ([]ImportObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImportedObjects, true
}

// SetImportedObjects sets field value
func (o *ObjectImportResult) SetImportedObjects(v []ImportObject) {
	o.ImportedObjects = v
}

func (o ObjectImportResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectImportResult) ToMap() (map[string]interface{}, error) {
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

func (o *ObjectImportResult) UnmarshalJSON(bytes []byte) (err error) {
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

	varObjectImportResult := _ObjectImportResult{}

	if err = json.Unmarshal(bytes, &varObjectImportResult); err == nil {
	*o = ObjectImportResult(varObjectImportResult)
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

type NullableObjectImportResult struct {
	value *ObjectImportResult
	isSet bool
}

func (v NullableObjectImportResult) Get() *ObjectImportResult {
	return v.value
}

func (v *NullableObjectImportResult) Set(val *ObjectImportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectImportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectImportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectImportResult(val *ObjectImportResult) *NullableObjectImportResult {
	return &NullableObjectImportResult{value: val, isSet: true}
}

func (v NullableObjectImportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectImportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


