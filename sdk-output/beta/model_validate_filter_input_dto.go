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

// ValidateFilterInputDto struct for ValidateFilterInputDto
type ValidateFilterInputDto struct {
	// Mock input to evaluate filter expression against.
	Input map[string]interface{} `json:"input"`
	// JSONPath filter to conditionally invoke trigger when expression evaluates to true.
	Filter string `json:"filter"`
	AdditionalProperties map[string]interface{}
}

type _ValidateFilterInputDto ValidateFilterInputDto

// NewValidateFilterInputDto instantiates a new ValidateFilterInputDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateFilterInputDto(input map[string]interface{}, filter string) *ValidateFilterInputDto {
	this := ValidateFilterInputDto{}
	this.Input = input
	this.Filter = filter
	return &this
}

// NewValidateFilterInputDtoWithDefaults instantiates a new ValidateFilterInputDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateFilterInputDtoWithDefaults() *ValidateFilterInputDto {
	this := ValidateFilterInputDto{}
	return &this
}

// GetInput returns the Input field value
func (o *ValidateFilterInputDto) GetInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *ValidateFilterInputDto) GetInputOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// SetInput sets field value
func (o *ValidateFilterInputDto) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetFilter returns the Filter field value
func (o *ValidateFilterInputDto) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *ValidateFilterInputDto) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *ValidateFilterInputDto) SetFilter(v string) {
	o.Filter = v
}

func (o ValidateFilterInputDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["input"] = o.Input
	}
	if true {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ValidateFilterInputDto) UnmarshalJSON(bytes []byte) (err error) {
	varValidateFilterInputDto := _ValidateFilterInputDto{}

	if err = json.Unmarshal(bytes, &varValidateFilterInputDto); err == nil {
		*o = ValidateFilterInputDto(varValidateFilterInputDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "input")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateFilterInputDto struct {
	value *ValidateFilterInputDto
	isSet bool
}

func (v NullableValidateFilterInputDto) Get() *ValidateFilterInputDto {
	return v.value
}

func (v *NullableValidateFilterInputDto) Set(val *ValidateFilterInputDto) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateFilterInputDto) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateFilterInputDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateFilterInputDto(val *ValidateFilterInputDto) *NullableValidateFilterInputDto {
	return &NullableValidateFilterInputDto{value: val, isSet: true}
}

func (v NullableValidateFilterInputDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateFilterInputDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


