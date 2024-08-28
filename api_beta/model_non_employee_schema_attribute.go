/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the NonEmployeeSchemaAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeSchemaAttribute{}

// NonEmployeeSchemaAttribute struct for NonEmployeeSchemaAttribute
type NonEmployeeSchemaAttribute struct {
	// Schema Attribute Id
	Id *string `json:"id,omitempty"`
	// True if this schema attribute is mandatory on all non-employees sources.
	System *bool `json:"system,omitempty"`
	// When the schema attribute was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the schema attribute was created.
	Created *time.Time `json:"created,omitempty"`
	Type NonEmployeeSchemaAttributeType `json:"type"`
	// Label displayed on the UI for this schema attribute.
	Label string `json:"label"`
	// The technical name of the attribute. Must be unique per source.
	TechnicalName string `json:"technicalName"`
	// help text displayed by UI.
	HelpText *string `json:"helpText,omitempty"`
	// Hint text that fills UI box.
	Placeholder *string `json:"placeholder,omitempty"`
	// If true, the schema attribute is required for all non-employees in the source
	Required *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSchemaAttribute NonEmployeeSchemaAttribute

// NewNonEmployeeSchemaAttribute instantiates a new NonEmployeeSchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSchemaAttribute(type_ NonEmployeeSchemaAttributeType, label string, technicalName string) *NonEmployeeSchemaAttribute {
	this := NonEmployeeSchemaAttribute{}
	this.Type = type_
	this.Label = label
	this.TechnicalName = technicalName
	return &this
}

// NewNonEmployeeSchemaAttributeWithDefaults instantiates a new NonEmployeeSchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSchemaAttributeWithDefaults() *NonEmployeeSchemaAttribute {
	this := NonEmployeeSchemaAttribute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeSchemaAttribute) SetId(v string) {
	o.Id = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *NonEmployeeSchemaAttribute) SetSystem(v bool) {
	o.System = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeSchemaAttribute) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeSchemaAttribute) SetCreated(v time.Time) {
	o.Created = &v
}

// GetType returns the Type field value
func (o *NonEmployeeSchemaAttribute) GetType() NonEmployeeSchemaAttributeType {
	if o == nil {
		var ret NonEmployeeSchemaAttributeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetTypeOk() (*NonEmployeeSchemaAttributeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NonEmployeeSchemaAttribute) SetType(v NonEmployeeSchemaAttributeType) {
	o.Type = v
}

// GetLabel returns the Label field value
func (o *NonEmployeeSchemaAttribute) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *NonEmployeeSchemaAttribute) SetLabel(v string) {
	o.Label = v
}

// GetTechnicalName returns the TechnicalName field value
func (o *NonEmployeeSchemaAttribute) GetTechnicalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetTechnicalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TechnicalName, true
}

// SetTechnicalName sets field value
func (o *NonEmployeeSchemaAttribute) SetTechnicalName(v string) {
	o.TechnicalName = v
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetHelpText() string {
	if o == nil || IsNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetHelpTextOk() (*string, bool) {
	if o == nil || IsNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasHelpText() bool {
	if o != nil && !IsNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *NonEmployeeSchemaAttribute) SetHelpText(v string) {
	o.HelpText = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *NonEmployeeSchemaAttribute) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *NonEmployeeSchemaAttribute) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSchemaAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *NonEmployeeSchemaAttribute) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *NonEmployeeSchemaAttribute) SetRequired(v bool) {
	o.Required = &v
}

func (o NonEmployeeSchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeSchemaAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	toSerialize["type"] = o.Type
	toSerialize["label"] = o.Label
	toSerialize["technicalName"] = o.TechnicalName
	if !IsNil(o.HelpText) {
		toSerialize["helpText"] = o.HelpText
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeSchemaAttribute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"label",
		"technicalName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNonEmployeeSchemaAttribute := _NonEmployeeSchemaAttribute{}

	err = json.Unmarshal(data, &varNonEmployeeSchemaAttribute)

	if err != nil {
		return err
	}

	*o = NonEmployeeSchemaAttribute(varNonEmployeeSchemaAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "system")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		delete(additionalProperties, "type")
		delete(additionalProperties, "label")
		delete(additionalProperties, "technicalName")
		delete(additionalProperties, "helpText")
		delete(additionalProperties, "placeholder")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSchemaAttribute struct {
	value *NonEmployeeSchemaAttribute
	isSet bool
}

func (v NullableNonEmployeeSchemaAttribute) Get() *NonEmployeeSchemaAttribute {
	return v.value
}

func (v *NullableNonEmployeeSchemaAttribute) Set(val *NonEmployeeSchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSchemaAttribute(val *NonEmployeeSchemaAttribute) *NullableNonEmployeeSchemaAttribute {
	return &NullableNonEmployeeSchemaAttribute{value: val, isSet: true}
}

func (v NullableNonEmployeeSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


