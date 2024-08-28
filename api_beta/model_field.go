/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the Field type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Field{}

// Field struct for Field
type Field struct {
	// Name of the FormItem
	Name *string `json:"name,omitempty"`
	// Display name of the field
	DisplayName *string `json:"displayName,omitempty"`
	// Type of the field to display
	DisplayType *string `json:"displayType,omitempty"`
	// True if the field is required
	Required *bool `json:"required,omitempty"`
	// List of allowed values for the field
	AllowedValuesList []map[string]interface{} `json:"allowedValuesList,omitempty"`
	// Value of the field
	Value map[string]interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Field Field

// NewField instantiates a new Field object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewField() *Field {
	this := Field{}
	return &this
}

// NewFieldWithDefaults instantiates a new Field object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldWithDefaults() *Field {
	this := Field{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Field) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Field) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Field) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Field) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Field) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Field) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *Field) GetDisplayType() string {
	if o == nil || IsNil(o.DisplayType) {
		var ret string
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetDisplayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayType) {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *Field) HasDisplayType() bool {
	if o != nil && !IsNil(o.DisplayType) {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *Field) SetDisplayType(v string) {
	o.DisplayType = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Field) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Field) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Field) SetRequired(v bool) {
	o.Required = &v
}

// GetAllowedValuesList returns the AllowedValuesList field value if set, zero value otherwise.
func (o *Field) GetAllowedValuesList() []map[string]interface{} {
	if o == nil || IsNil(o.AllowedValuesList) {
		var ret []map[string]interface{}
		return ret
	}
	return o.AllowedValuesList
}

// GetAllowedValuesListOk returns a tuple with the AllowedValuesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetAllowedValuesListOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AllowedValuesList) {
		return nil, false
	}
	return o.AllowedValuesList, true
}

// HasAllowedValuesList returns a boolean if a field has been set.
func (o *Field) HasAllowedValuesList() bool {
	if o != nil && !IsNil(o.AllowedValuesList) {
		return true
	}

	return false
}

// SetAllowedValuesList gets a reference to the given []map[string]interface{} and assigns it to the AllowedValuesList field.
func (o *Field) SetAllowedValuesList(v []map[string]interface{}) {
	o.AllowedValuesList = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Field) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Field) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *Field) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o Field) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Field) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DisplayType) {
		toSerialize["displayType"] = o.DisplayType
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.AllowedValuesList) {
		toSerialize["allowedValuesList"] = o.AllowedValuesList
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Field) UnmarshalJSON(data []byte) (err error) {
	varField := _Field{}

	err = json.Unmarshal(data, &varField)

	if err != nil {
		return err
	}

	*o = Field(varField)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "displayType")
		delete(additionalProperties, "required")
		delete(additionalProperties, "allowedValuesList")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableField struct {
	value *Field
	isSet bool
}

func (v NullableField) Get() *Field {
	return v.value
}

func (v *NullableField) Set(val *Field) {
	v.value = val
	v.isSet = true
}

func (v NullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableField(val *Field) *NullableField {
	return &NullableField{value: val, isSet: true}
}

func (v NullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


