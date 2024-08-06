/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the FormElementDynamicDataSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormElementDynamicDataSource{}

// FormElementDynamicDataSource struct for FormElementDynamicDataSource
type FormElementDynamicDataSource struct {
	Config *FormElementDynamicDataSourceConfig `json:"config,omitempty"`
	// DataSourceType is a FormElementDataSourceType value STATIC FormElementDataSourceTypeStatic INTERNAL FormElementDataSourceTypeInternal SEARCH FormElementDataSourceTypeSearch FORM_INPUT FormElementDataSourceTypeFormInput
	DataSourceType *string `json:"dataSourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormElementDynamicDataSource FormElementDynamicDataSource

// NewFormElementDynamicDataSource instantiates a new FormElementDynamicDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormElementDynamicDataSource() *FormElementDynamicDataSource {
	this := FormElementDynamicDataSource{}
	return &this
}

// NewFormElementDynamicDataSourceWithDefaults instantiates a new FormElementDynamicDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormElementDynamicDataSourceWithDefaults() *FormElementDynamicDataSource {
	this := FormElementDynamicDataSource{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *FormElementDynamicDataSource) GetConfig() FormElementDynamicDataSourceConfig {
	if o == nil || isNil(o.Config) {
		var ret FormElementDynamicDataSourceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElementDynamicDataSource) GetConfigOk() (*FormElementDynamicDataSourceConfig, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *FormElementDynamicDataSource) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given FormElementDynamicDataSourceConfig and assigns it to the Config field.
func (o *FormElementDynamicDataSource) SetConfig(v FormElementDynamicDataSourceConfig) {
	o.Config = &v
}

// GetDataSourceType returns the DataSourceType field value if set, zero value otherwise.
func (o *FormElementDynamicDataSource) GetDataSourceType() string {
	if o == nil || isNil(o.DataSourceType) {
		var ret string
		return ret
	}
	return *o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElementDynamicDataSource) GetDataSourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.DataSourceType) {
		return nil, false
	}
	return o.DataSourceType, true
}

// HasDataSourceType returns a boolean if a field has been set.
func (o *FormElementDynamicDataSource) HasDataSourceType() bool {
	if o != nil && !isNil(o.DataSourceType) {
		return true
	}

	return false
}

// SetDataSourceType gets a reference to the given string and assigns it to the DataSourceType field.
func (o *FormElementDynamicDataSource) SetDataSourceType(v string) {
	o.DataSourceType = &v
}

func (o FormElementDynamicDataSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormElementDynamicDataSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.DataSourceType) {
		toSerialize["dataSourceType"] = o.DataSourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormElementDynamicDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varFormElementDynamicDataSource := _FormElementDynamicDataSource{}

	if err = json.Unmarshal(bytes, &varFormElementDynamicDataSource); err == nil {
	*o = FormElementDynamicDataSource(varFormElementDynamicDataSource)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "dataSourceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormElementDynamicDataSource struct {
	value *FormElementDynamicDataSource
	isSet bool
}

func (v NullableFormElementDynamicDataSource) Get() *FormElementDynamicDataSource {
	return v.value
}

func (v *NullableFormElementDynamicDataSource) Set(val *FormElementDynamicDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableFormElementDynamicDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableFormElementDynamicDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormElementDynamicDataSource(val *FormElementDynamicDataSource) *NullableFormElementDynamicDataSource {
	return &NullableFormElementDynamicDataSource{value: val, isSet: true}
}

func (v NullableFormElementDynamicDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormElementDynamicDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


