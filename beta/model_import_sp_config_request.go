/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the ImportSpConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportSpConfigRequest{}

// ImportSpConfigRequest struct for ImportSpConfigRequest
type ImportSpConfigRequest struct {
	// JSON file containing the objects to be imported.
	Data *os.File `json:"data"`
	Options *ImportOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportSpConfigRequest ImportSpConfigRequest

// NewImportSpConfigRequest instantiates a new ImportSpConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSpConfigRequest(data *os.File) *ImportSpConfigRequest {
	this := ImportSpConfigRequest{}
	this.Data = data
	return &this
}

// NewImportSpConfigRequestWithDefaults instantiates a new ImportSpConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSpConfigRequestWithDefaults() *ImportSpConfigRequest {
	this := ImportSpConfigRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ImportSpConfigRequest) GetData() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImportSpConfigRequest) GetDataOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ImportSpConfigRequest) SetData(v *os.File) {
	o.Data = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ImportSpConfigRequest) GetOptions() ImportOptions {
	if o == nil || isNil(o.Options) {
		var ret ImportOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSpConfigRequest) GetOptionsOk() (*ImportOptions, bool) {
	if o == nil || isNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ImportSpConfigRequest) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ImportOptions and assigns it to the Options field.
func (o *ImportSpConfigRequest) SetOptions(v ImportOptions) {
	o.Options = &v
}

func (o ImportSpConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportSpConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportSpConfigRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varImportSpConfigRequest := _ImportSpConfigRequest{}

	if err = json.Unmarshal(bytes, &varImportSpConfigRequest); err == nil {
	*o = ImportSpConfigRequest(varImportSpConfigRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportSpConfigRequest struct {
	value *ImportSpConfigRequest
	isSet bool
}

func (v NullableImportSpConfigRequest) Get() *ImportSpConfigRequest {
	return v.value
}

func (v *NullableImportSpConfigRequest) Set(val *ImportSpConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSpConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSpConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSpConfigRequest(val *ImportSpConfigRequest) *NullableImportSpConfigRequest {
	return &NullableImportSpConfigRequest{value: val, isSet: true}
}

func (v NullableImportSpConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSpConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


