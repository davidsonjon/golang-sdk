/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the WorkflowLibraryFormFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowLibraryFormFields{}

// WorkflowLibraryFormFields struct for WorkflowLibraryFormFields
type WorkflowLibraryFormFields struct {
	// Description of the form field
	Description *string `json:"description,omitempty"`
	// Describes the form field in the UI
	HelpText *string `json:"helpText,omitempty"`
	// A human readable name for this form field in the UI
	Label *string `json:"label,omitempty"`
	// The name of the input attribute
	Name *string `json:"name,omitempty"`
	// Denotes if this field is a required attribute
	Required *bool `json:"required,omitempty"`
	// The type of the form field
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLibraryFormFields WorkflowLibraryFormFields

// NewWorkflowLibraryFormFields instantiates a new WorkflowLibraryFormFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLibraryFormFields() *WorkflowLibraryFormFields {
	this := WorkflowLibraryFormFields{}
	var required bool = false
	this.Required = &required
	return &this
}

// NewWorkflowLibraryFormFieldsWithDefaults instantiates a new WorkflowLibraryFormFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLibraryFormFieldsWithDefaults() *WorkflowLibraryFormFields {
	this := WorkflowLibraryFormFields{}
	var required bool = false
	this.Required = &required
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowLibraryFormFields) SetDescription(v string) {
	o.Description = &v
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetHelpText() string {
	if o == nil || IsNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetHelpTextOk() (*string, bool) {
	if o == nil || IsNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasHelpText() bool {
	if o != nil && !IsNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *WorkflowLibraryFormFields) SetHelpText(v string) {
	o.HelpText = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowLibraryFormFields) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowLibraryFormFields) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *WorkflowLibraryFormFields) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowLibraryFormFields) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowLibraryFormFields) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *WorkflowLibraryFormFields) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *WorkflowLibraryFormFields) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *WorkflowLibraryFormFields) UnsetType() {
	o.Type.Unset()
}

func (o WorkflowLibraryFormFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowLibraryFormFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HelpText) {
		toSerialize["helpText"] = o.HelpText
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowLibraryFormFields) UnmarshalJSON(data []byte) (err error) {
	varWorkflowLibraryFormFields := _WorkflowLibraryFormFields{}

	err = json.Unmarshal(data, &varWorkflowLibraryFormFields)

	if err != nil {
		return err
	}

	*o = WorkflowLibraryFormFields(varWorkflowLibraryFormFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "helpText")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "required")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLibraryFormFields struct {
	value *WorkflowLibraryFormFields
	isSet bool
}

func (v NullableWorkflowLibraryFormFields) Get() *WorkflowLibraryFormFields {
	return v.value
}

func (v *NullableWorkflowLibraryFormFields) Set(val *WorkflowLibraryFormFields) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryFormFields) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryFormFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryFormFields(val *WorkflowLibraryFormFields) *NullableWorkflowLibraryFormFields {
	return &NullableWorkflowLibraryFormFields{value: val, isSet: true}
}

func (v NullableWorkflowLibraryFormFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryFormFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


