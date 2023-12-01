/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateWorkflowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkflowRequest{}

// CreateWorkflowRequest struct for CreateWorkflowRequest
type CreateWorkflowRequest struct {
	// The name of the workflow
	Name string `json:"name"`
	Owner WorkflowBodyOwner `json:"owner"`
	// Description of what the workflow accomplishes
	Description *string `json:"description,omitempty"`
	Definition *WorkflowDefinition `json:"definition,omitempty"`
	// Enable or disable the workflow.  Workflows cannot be created in an enabled state.
	Enabled *bool `json:"enabled,omitempty"`
	Trigger *WorkflowTrigger `json:"trigger,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateWorkflowRequest CreateWorkflowRequest

// NewCreateWorkflowRequest instantiates a new CreateWorkflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkflowRequest(name string, owner WorkflowBodyOwner) *CreateWorkflowRequest {
	this := CreateWorkflowRequest{}
	this.Name = name
	this.Owner = owner
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewCreateWorkflowRequestWithDefaults instantiates a new CreateWorkflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkflowRequestWithDefaults() *CreateWorkflowRequest {
	this := CreateWorkflowRequest{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value
func (o *CreateWorkflowRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkflowRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkflowRequest) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value
func (o *CreateWorkflowRequest) GetOwner() WorkflowBodyOwner {
	if o == nil {
		var ret WorkflowBodyOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *CreateWorkflowRequest) GetOwnerOk() (*WorkflowBodyOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *CreateWorkflowRequest) SetOwner(v WorkflowBodyOwner) {
	o.Owner = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateWorkflowRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkflowRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateWorkflowRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateWorkflowRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *CreateWorkflowRequest) GetDefinition() WorkflowDefinition {
	if o == nil || isNil(o.Definition) {
		var ret WorkflowDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkflowRequest) GetDefinitionOk() (*WorkflowDefinition, bool) {
	if o == nil || isNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *CreateWorkflowRequest) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given WorkflowDefinition and assigns it to the Definition field.
func (o *CreateWorkflowRequest) SetDefinition(v WorkflowDefinition) {
	o.Definition = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateWorkflowRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkflowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateWorkflowRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateWorkflowRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *CreateWorkflowRequest) GetTrigger() WorkflowTrigger {
	if o == nil || isNil(o.Trigger) {
		var ret WorkflowTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkflowRequest) GetTriggerOk() (*WorkflowTrigger, bool) {
	if o == nil || isNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *CreateWorkflowRequest) HasTrigger() bool {
	if o != nil && !isNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given WorkflowTrigger and assigns it to the Trigger field.
func (o *CreateWorkflowRequest) SetTrigger(v WorkflowTrigger) {
	o.Trigger = &v
}

func (o CreateWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkflowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["owner"] = o.Owner
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateWorkflowRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner",
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

	varCreateWorkflowRequest := _CreateWorkflowRequest{}

	if err = json.Unmarshal(bytes, &varCreateWorkflowRequest); err == nil {
	*o = CreateWorkflowRequest(varCreateWorkflowRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "description")
		delete(additionalProperties, "definition")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "trigger")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateWorkflowRequest struct {
	value *CreateWorkflowRequest
	isSet bool
}

func (v NullableCreateWorkflowRequest) Get() *CreateWorkflowRequest {
	return v.value
}

func (v *NullableCreateWorkflowRequest) Set(val *CreateWorkflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkflowRequest(val *CreateWorkflowRequest) *NullableCreateWorkflowRequest {
	return &NullableCreateWorkflowRequest{value: val, isSet: true}
}

func (v NullableCreateWorkflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


