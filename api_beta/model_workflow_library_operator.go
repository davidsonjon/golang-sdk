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
)

// checks if the WorkflowLibraryOperator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowLibraryOperator{}

// WorkflowLibraryOperator struct for WorkflowLibraryOperator
type WorkflowLibraryOperator struct {
	// Operator ID.
	Id *string `json:"id,omitempty"`
	// Operator friendly name
	Name *string `json:"name,omitempty"`
	// Operator type
	Type *string `json:"type,omitempty"`
	// Description of the operator
	Description *string `json:"description,omitempty"`
	// Determines whether the dynamic output schema is returned in place of the action's output schema. The dynamic schema lists non-static properties, like properties of a workflow form where each form has different fields. These will be provided dynamically based on available form fields.
	IsDynamicSchema *bool `json:"isDynamicSchema,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	DeprecatedBy *time.Time `json:"deprecatedBy,omitempty"`
	IsSimulationEnabled *bool `json:"isSimulationEnabled,omitempty"`
	// One or more inputs that the operator accepts
	FormFields []WorkflowLibraryFormFields `json:"formFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLibraryOperator WorkflowLibraryOperator

// NewWorkflowLibraryOperator instantiates a new WorkflowLibraryOperator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLibraryOperator() *WorkflowLibraryOperator {
	this := WorkflowLibraryOperator{}
	return &this
}

// NewWorkflowLibraryOperatorWithDefaults instantiates a new WorkflowLibraryOperator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLibraryOperatorWithDefaults() *WorkflowLibraryOperator {
	this := WorkflowLibraryOperator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowLibraryOperator) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowLibraryOperator) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowLibraryOperator) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowLibraryOperator) SetDescription(v string) {
	o.Description = &v
}

// GetIsDynamicSchema returns the IsDynamicSchema field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetIsDynamicSchema() bool {
	if o == nil || IsNil(o.IsDynamicSchema) {
		var ret bool
		return ret
	}
	return *o.IsDynamicSchema
}

// GetIsDynamicSchemaOk returns a tuple with the IsDynamicSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetIsDynamicSchemaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDynamicSchema) {
		return nil, false
	}
	return o.IsDynamicSchema, true
}

// HasIsDynamicSchema returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasIsDynamicSchema() bool {
	if o != nil && !IsNil(o.IsDynamicSchema) {
		return true
	}

	return false
}

// SetIsDynamicSchema gets a reference to the given bool and assigns it to the IsDynamicSchema field.
func (o *WorkflowLibraryOperator) SetIsDynamicSchema(v bool) {
	o.IsDynamicSchema = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *WorkflowLibraryOperator) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecatedBy returns the DeprecatedBy field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetDeprecatedBy() time.Time {
	if o == nil || IsNil(o.DeprecatedBy) {
		var ret time.Time
		return ret
	}
	return *o.DeprecatedBy
}

// GetDeprecatedByOk returns a tuple with the DeprecatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetDeprecatedByOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeprecatedBy) {
		return nil, false
	}
	return o.DeprecatedBy, true
}

// HasDeprecatedBy returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasDeprecatedBy() bool {
	if o != nil && !IsNil(o.DeprecatedBy) {
		return true
	}

	return false
}

// SetDeprecatedBy gets a reference to the given time.Time and assigns it to the DeprecatedBy field.
func (o *WorkflowLibraryOperator) SetDeprecatedBy(v time.Time) {
	o.DeprecatedBy = &v
}

// GetIsSimulationEnabled returns the IsSimulationEnabled field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetIsSimulationEnabled() bool {
	if o == nil || IsNil(o.IsSimulationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSimulationEnabled
}

// GetIsSimulationEnabledOk returns a tuple with the IsSimulationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetIsSimulationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSimulationEnabled) {
		return nil, false
	}
	return o.IsSimulationEnabled, true
}

// HasIsSimulationEnabled returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasIsSimulationEnabled() bool {
	if o != nil && !IsNil(o.IsSimulationEnabled) {
		return true
	}

	return false
}

// SetIsSimulationEnabled gets a reference to the given bool and assigns it to the IsSimulationEnabled field.
func (o *WorkflowLibraryOperator) SetIsSimulationEnabled(v bool) {
	o.IsSimulationEnabled = &v
}

// GetFormFields returns the FormFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowLibraryOperator) GetFormFields() []WorkflowLibraryFormFields {
	if o == nil {
		var ret []WorkflowLibraryFormFields
		return ret
	}
	return o.FormFields
}

// GetFormFieldsOk returns a tuple with the FormFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowLibraryOperator) GetFormFieldsOk() ([]WorkflowLibraryFormFields, bool) {
	if o == nil || IsNil(o.FormFields) {
		return nil, false
	}
	return o.FormFields, true
}

// HasFormFields returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasFormFields() bool {
	if o != nil && !IsNil(o.FormFields) {
		return true
	}

	return false
}

// SetFormFields gets a reference to the given []WorkflowLibraryFormFields and assigns it to the FormFields field.
func (o *WorkflowLibraryOperator) SetFormFields(v []WorkflowLibraryFormFields) {
	o.FormFields = v
}

func (o WorkflowLibraryOperator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowLibraryOperator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsDynamicSchema) {
		toSerialize["isDynamicSchema"] = o.IsDynamicSchema
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.DeprecatedBy) {
		toSerialize["deprecatedBy"] = o.DeprecatedBy
	}
	if !IsNil(o.IsSimulationEnabled) {
		toSerialize["isSimulationEnabled"] = o.IsSimulationEnabled
	}
	if o.FormFields != nil {
		toSerialize["formFields"] = o.FormFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowLibraryOperator) UnmarshalJSON(data []byte) (err error) {
	varWorkflowLibraryOperator := _WorkflowLibraryOperator{}

	err = json.Unmarshal(data, &varWorkflowLibraryOperator)

	if err != nil {
		return err
	}

	*o = WorkflowLibraryOperator(varWorkflowLibraryOperator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "isDynamicSchema")
		delete(additionalProperties, "deprecated")
		delete(additionalProperties, "deprecatedBy")
		delete(additionalProperties, "isSimulationEnabled")
		delete(additionalProperties, "formFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLibraryOperator struct {
	value *WorkflowLibraryOperator
	isSet bool
}

func (v NullableWorkflowLibraryOperator) Get() *WorkflowLibraryOperator {
	return v.value
}

func (v *NullableWorkflowLibraryOperator) Set(val *WorkflowLibraryOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryOperator(val *WorkflowLibraryOperator) *NullableWorkflowLibraryOperator {
	return &NullableWorkflowLibraryOperator{value: val, isSet: true}
}

func (v NullableWorkflowLibraryOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


