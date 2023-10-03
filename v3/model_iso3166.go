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

// checks if the ISO3166 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISO3166{}

// ISO3166 struct for ISO3166
type ISO3166 struct {
	// An optional value to denote which ISO 3166 format to return. Valid values are:   `alpha2` - Two-character country code (e.g., \"US\"); this is the default value if no format is supplied   `alpha3` - Three-character country code (e.g., \"USA\")   `numeric` - The numeric country code (e.g., \"840\") 
	Format *string `json:"format,omitempty"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ISO3166 ISO3166

// NewISO3166 instantiates a new ISO3166 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISO3166() *ISO3166 {
	this := ISO3166{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewISO3166WithDefaults instantiates a new ISO3166 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISO3166WithDefaults() *ISO3166 {
	this := ISO3166{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ISO3166) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISO3166) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ISO3166) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ISO3166) SetFormat(v string) {
	o.Format = &v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *ISO3166) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISO3166) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *ISO3166) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *ISO3166) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ISO3166) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISO3166) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ISO3166) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *ISO3166) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o ISO3166) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISO3166) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ISO3166) UnmarshalJSON(bytes []byte) (err error) {
	varISO3166 := _ISO3166{}

	if err = json.Unmarshal(bytes, &varISO3166); err == nil {
		*o = ISO3166(varISO3166)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableISO3166 struct {
	value *ISO3166
	isSet bool
}

func (v NullableISO3166) Get() *ISO3166 {
	return v.value
}

func (v *NullableISO3166) Set(val *ISO3166) {
	v.value = val
	v.isSet = true
}

func (v NullableISO3166) IsSet() bool {
	return v.isSet
}

func (v *NullableISO3166) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISO3166(val *ISO3166) *NullableISO3166 {
	return &NullableISO3166{value: val, isSet: true}
}

func (v NullableISO3166) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISO3166) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


