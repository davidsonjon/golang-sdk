/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the FieldDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldDetails{}

// FieldDetails struct for FieldDetails
type FieldDetails struct {
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

type _FieldDetails FieldDetails

// NewFieldDetails instantiates a new FieldDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldDetails() *FieldDetails {
	this := FieldDetails{}
	return &this
}

// NewFieldDetailsWithDefaults instantiates a new FieldDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldDetailsWithDefaults() *FieldDetails {
	this := FieldDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldDetails) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FieldDetails) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FieldDetails) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FieldDetails) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *FieldDetails) GetDisplayType() string {
	if o == nil || isNil(o.DisplayType) {
		var ret string
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetails) GetDisplayTypeOk() (*string, bool) {
	if o == nil || isNil(o.DisplayType) {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *FieldDetails) HasDisplayType() bool {
	if o != nil && !isNil(o.DisplayType) {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *FieldDetails) SetDisplayType(v string) {
	o.DisplayType = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FieldDetails) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetails) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FieldDetails) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FieldDetails) SetRequired(v bool) {
	o.Required = &v
}

// GetAllowedValuesList returns the AllowedValuesList field value if set, zero value otherwise.
func (o *FieldDetails) GetAllowedValuesList() []map[string]interface{} {
	if o == nil || isNil(o.AllowedValuesList) {
		var ret []map[string]interface{}
		return ret
	}
	return o.AllowedValuesList
}

// GetAllowedValuesListOk returns a tuple with the AllowedValuesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetails) GetAllowedValuesListOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.AllowedValuesList) {
		return nil, false
	}
	return o.AllowedValuesList, true
}

// HasAllowedValuesList returns a boolean if a field has been set.
func (o *FieldDetails) HasAllowedValuesList() bool {
	if o != nil && !isNil(o.AllowedValuesList) {
		return true
	}

	return false
}

// SetAllowedValuesList gets a reference to the given []map[string]interface{} and assigns it to the AllowedValuesList field.
func (o *FieldDetails) SetAllowedValuesList(v []map[string]interface{}) {
	o.AllowedValuesList = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FieldDetails) GetValue() map[string]interface{} {
	if o == nil || isNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetails) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FieldDetails) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *FieldDetails) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o FieldDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.DisplayType) {
		toSerialize["displayType"] = o.DisplayType
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !isNil(o.AllowedValuesList) {
		toSerialize["allowedValuesList"] = o.AllowedValuesList
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FieldDetails) UnmarshalJSON(bytes []byte) (err error) {
	varFieldDetails := _FieldDetails{}

	if err = json.Unmarshal(bytes, &varFieldDetails); err == nil {
	*o = FieldDetails(varFieldDetails)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

type NullableFieldDetails struct {
	value *FieldDetails
	isSet bool
}

func (v NullableFieldDetails) Get() *FieldDetails {
	return v.value
}

func (v *NullableFieldDetails) Set(val *FieldDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldDetails(val *FieldDetails) *NullableFieldDetails {
	return &NullableFieldDetails{value: val, isSet: true}
}

func (v NullableFieldDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

