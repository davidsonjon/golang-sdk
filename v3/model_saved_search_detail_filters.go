/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the SavedSearchDetailFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedSearchDetailFilters{}

// SavedSearchDetailFilters struct for SavedSearchDetailFilters
type SavedSearchDetailFilters struct {
	Type *FilterType `json:"type,omitempty"`
	Range *Range `json:"range,omitempty"`
	// The terms to be filtered.
	Terms []string `json:"terms,omitempty"`
	// Indicates if the filter excludes results.
	Exclude *bool `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearchDetailFilters SavedSearchDetailFilters

// NewSavedSearchDetailFilters instantiates a new SavedSearchDetailFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchDetailFilters() *SavedSearchDetailFilters {
	this := SavedSearchDetailFilters{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// NewSavedSearchDetailFiltersWithDefaults instantiates a new SavedSearchDetailFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchDetailFiltersWithDefaults() *SavedSearchDetailFilters {
	this := SavedSearchDetailFilters{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SavedSearchDetailFilters) GetType() FilterType {
	if o == nil || isNil(o.Type) {
		var ret FilterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchDetailFilters) GetTypeOk() (*FilterType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SavedSearchDetailFilters) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FilterType and assigns it to the Type field.
func (o *SavedSearchDetailFilters) SetType(v FilterType) {
	o.Type = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *SavedSearchDetailFilters) GetRange() Range {
	if o == nil || isNil(o.Range) {
		var ret Range
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchDetailFilters) GetRangeOk() (*Range, bool) {
	if o == nil || isNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *SavedSearchDetailFilters) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given Range and assigns it to the Range field.
func (o *SavedSearchDetailFilters) SetRange(v Range) {
	o.Range = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *SavedSearchDetailFilters) GetTerms() []string {
	if o == nil || isNil(o.Terms) {
		var ret []string
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchDetailFilters) GetTermsOk() ([]string, bool) {
	if o == nil || isNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *SavedSearchDetailFilters) HasTerms() bool {
	if o != nil && !isNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given []string and assigns it to the Terms field.
func (o *SavedSearchDetailFilters) SetTerms(v []string) {
	o.Terms = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *SavedSearchDetailFilters) GetExclude() bool {
	if o == nil || isNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchDetailFilters) GetExcludeOk() (*bool, bool) {
	if o == nil || isNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *SavedSearchDetailFilters) HasExclude() bool {
	if o != nil && !isNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *SavedSearchDetailFilters) SetExclude(v bool) {
	o.Exclude = &v
}

func (o SavedSearchDetailFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedSearchDetailFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !isNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !isNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavedSearchDetailFilters) UnmarshalJSON(bytes []byte) (err error) {
	varSavedSearchDetailFilters := _SavedSearchDetailFilters{}

	if err = json.Unmarshal(bytes, &varSavedSearchDetailFilters); err == nil {
		*o = SavedSearchDetailFilters(varSavedSearchDetailFilters)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "range")
		delete(additionalProperties, "terms")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearchDetailFilters struct {
	value *SavedSearchDetailFilters
	isSet bool
}

func (v NullableSavedSearchDetailFilters) Get() *SavedSearchDetailFilters {
	return v.value
}

func (v *NullableSavedSearchDetailFilters) Set(val *SavedSearchDetailFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchDetailFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchDetailFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchDetailFilters(val *SavedSearchDetailFilters) *NullableSavedSearchDetailFilters {
	return &NullableSavedSearchDetailFilters{value: val, isSet: true}
}

func (v NullableSavedSearchDetailFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchDetailFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


