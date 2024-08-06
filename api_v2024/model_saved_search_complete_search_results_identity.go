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

// checks if the SavedSearchCompleteSearchResultsIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedSearchCompleteSearchResultsIdentity{}

// SavedSearchCompleteSearchResultsIdentity A table of identities that match the search criteria.
type SavedSearchCompleteSearchResultsIdentity struct {
	// The number of rows in the table.
	Count string `json:"count"`
	// The type of object represented in the table.
	Noun string `json:"noun"`
	// A sample of the data in the table.
	Preview [][]string `json:"preview"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearchCompleteSearchResultsIdentity SavedSearchCompleteSearchResultsIdentity

// NewSavedSearchCompleteSearchResultsIdentity instantiates a new SavedSearchCompleteSearchResultsIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchCompleteSearchResultsIdentity(count string, noun string, preview [][]string) *SavedSearchCompleteSearchResultsIdentity {
	this := SavedSearchCompleteSearchResultsIdentity{}
	this.Count = count
	this.Noun = noun
	this.Preview = preview
	return &this
}

// NewSavedSearchCompleteSearchResultsIdentityWithDefaults instantiates a new SavedSearchCompleteSearchResultsIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchCompleteSearchResultsIdentityWithDefaults() *SavedSearchCompleteSearchResultsIdentity {
	this := SavedSearchCompleteSearchResultsIdentity{}
	return &this
}

// GetCount returns the Count field value
func (o *SavedSearchCompleteSearchResultsIdentity) GetCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SavedSearchCompleteSearchResultsIdentity) GetCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SavedSearchCompleteSearchResultsIdentity) SetCount(v string) {
	o.Count = v
}

// GetNoun returns the Noun field value
func (o *SavedSearchCompleteSearchResultsIdentity) GetNoun() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Noun
}

// GetNounOk returns a tuple with the Noun field value
// and a boolean to check if the value has been set.
func (o *SavedSearchCompleteSearchResultsIdentity) GetNounOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Noun, true
}

// SetNoun sets field value
func (o *SavedSearchCompleteSearchResultsIdentity) SetNoun(v string) {
	o.Noun = v
}

// GetPreview returns the Preview field value
func (o *SavedSearchCompleteSearchResultsIdentity) GetPreview() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value
// and a boolean to check if the value has been set.
func (o *SavedSearchCompleteSearchResultsIdentity) GetPreviewOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preview, true
}

// SetPreview sets field value
func (o *SavedSearchCompleteSearchResultsIdentity) SetPreview(v [][]string) {
	o.Preview = v
}

func (o SavedSearchCompleteSearchResultsIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedSearchCompleteSearchResultsIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["noun"] = o.Noun
	toSerialize["preview"] = o.Preview

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavedSearchCompleteSearchResultsIdentity) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"noun",
		"preview",
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

	varSavedSearchCompleteSearchResultsIdentity := _SavedSearchCompleteSearchResultsIdentity{}

	if err = json.Unmarshal(bytes, &varSavedSearchCompleteSearchResultsIdentity); err == nil {
	*o = SavedSearchCompleteSearchResultsIdentity(varSavedSearchCompleteSearchResultsIdentity)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "noun")
		delete(additionalProperties, "preview")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearchCompleteSearchResultsIdentity struct {
	value *SavedSearchCompleteSearchResultsIdentity
	isSet bool
}

func (v NullableSavedSearchCompleteSearchResultsIdentity) Get() *SavedSearchCompleteSearchResultsIdentity {
	return v.value
}

func (v *NullableSavedSearchCompleteSearchResultsIdentity) Set(val *SavedSearchCompleteSearchResultsIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchCompleteSearchResultsIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchCompleteSearchResultsIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchCompleteSearchResultsIdentity(val *SavedSearchCompleteSearchResultsIdentity) *NullableSavedSearchCompleteSearchResultsIdentity {
	return &NullableSavedSearchCompleteSearchResultsIdentity{value: val, isSet: true}
}

func (v NullableSavedSearchCompleteSearchResultsIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchCompleteSearchResultsIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


