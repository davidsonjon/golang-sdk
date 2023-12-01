/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"time"
)

// checks if the Workflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Workflow{}

// Workflow struct for Workflow
type Workflow struct {
	// The name of the workflow
	Name *string `json:"name,omitempty"`
	Owner *WorkflowBodyOwner `json:"owner,omitempty"`
	// Description of what the workflow accomplishes
	Description *string `json:"description,omitempty"`
	Definition *WorkflowDefinition `json:"definition,omitempty"`
	// Enable or disable the workflow.  Workflows cannot be created in an enabled state.
	Enabled *bool `json:"enabled,omitempty"`
	Trigger *WorkflowTrigger `json:"trigger,omitempty"`
	// Workflow ID. This is a UUID generated upon creation.
	Id *string `json:"id,omitempty"`
	// The number of times this workflow has been executed.
	ExecutionCount *int32 `json:"executionCount,omitempty"`
	// The number of times this workflow has failed during execution.
	FailureCount *int32 `json:"failureCount,omitempty"`
	// The date and time the workflow was created.
	Created *time.Time `json:"created,omitempty"`
	Creator *WorkflowAllOfCreator `json:"creator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Workflow Workflow

// NewWorkflow instantiates a new Workflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflow() *Workflow {
	this := Workflow{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewWorkflowWithDefaults instantiates a new Workflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWithDefaults() *Workflow {
	this := Workflow{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Workflow) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Workflow) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Workflow) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Workflow) GetOwner() WorkflowBodyOwner {
	if o == nil || isNil(o.Owner) {
		var ret WorkflowBodyOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetOwnerOk() (*WorkflowBodyOwner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Workflow) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given WorkflowBodyOwner and assigns it to the Owner field.
func (o *Workflow) SetOwner(v WorkflowBodyOwner) {
	o.Owner = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Workflow) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Workflow) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Workflow) SetDescription(v string) {
	o.Description = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *Workflow) GetDefinition() WorkflowDefinition {
	if o == nil || isNil(o.Definition) {
		var ret WorkflowDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetDefinitionOk() (*WorkflowDefinition, bool) {
	if o == nil || isNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *Workflow) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given WorkflowDefinition and assigns it to the Definition field.
func (o *Workflow) SetDefinition(v WorkflowDefinition) {
	o.Definition = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Workflow) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Workflow) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Workflow) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *Workflow) GetTrigger() WorkflowTrigger {
	if o == nil || isNil(o.Trigger) {
		var ret WorkflowTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetTriggerOk() (*WorkflowTrigger, bool) {
	if o == nil || isNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *Workflow) HasTrigger() bool {
	if o != nil && !isNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given WorkflowTrigger and assigns it to the Trigger field.
func (o *Workflow) SetTrigger(v WorkflowTrigger) {
	o.Trigger = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Workflow) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Workflow) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Workflow) SetId(v string) {
	o.Id = &v
}

// GetExecutionCount returns the ExecutionCount field value if set, zero value otherwise.
func (o *Workflow) GetExecutionCount() int32 {
	if o == nil || isNil(o.ExecutionCount) {
		var ret int32
		return ret
	}
	return *o.ExecutionCount
}

// GetExecutionCountOk returns a tuple with the ExecutionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetExecutionCountOk() (*int32, bool) {
	if o == nil || isNil(o.ExecutionCount) {
		return nil, false
	}
	return o.ExecutionCount, true
}

// HasExecutionCount returns a boolean if a field has been set.
func (o *Workflow) HasExecutionCount() bool {
	if o != nil && !isNil(o.ExecutionCount) {
		return true
	}

	return false
}

// SetExecutionCount gets a reference to the given int32 and assigns it to the ExecutionCount field.
func (o *Workflow) SetExecutionCount(v int32) {
	o.ExecutionCount = &v
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise.
func (o *Workflow) GetFailureCount() int32 {
	if o == nil || isNil(o.FailureCount) {
		var ret int32
		return ret
	}
	return *o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetFailureCountOk() (*int32, bool) {
	if o == nil || isNil(o.FailureCount) {
		return nil, false
	}
	return o.FailureCount, true
}

// HasFailureCount returns a boolean if a field has been set.
func (o *Workflow) HasFailureCount() bool {
	if o != nil && !isNil(o.FailureCount) {
		return true
	}

	return false
}

// SetFailureCount gets a reference to the given int32 and assigns it to the FailureCount field.
func (o *Workflow) SetFailureCount(v int32) {
	o.FailureCount = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Workflow) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Workflow) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Workflow) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Workflow) GetCreator() WorkflowAllOfCreator {
	if o == nil || isNil(o.Creator) {
		var ret WorkflowAllOfCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetCreatorOk() (*WorkflowAllOfCreator, bool) {
	if o == nil || isNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Workflow) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given WorkflowAllOfCreator and assigns it to the Creator field.
func (o *Workflow) SetCreator(v WorkflowAllOfCreator) {
	o.Creator = &v
}

func (o Workflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Workflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
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
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ExecutionCount) {
		toSerialize["executionCount"] = o.ExecutionCount
	}
	if !isNil(o.FailureCount) {
		toSerialize["failureCount"] = o.FailureCount
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Workflow) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflow := _Workflow{}

	if err = json.Unmarshal(bytes, &varWorkflow); err == nil {
	*o = Workflow(varWorkflow)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "description")
		delete(additionalProperties, "definition")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "trigger")
		delete(additionalProperties, "id")
		delete(additionalProperties, "executionCount")
		delete(additionalProperties, "failureCount")
		delete(additionalProperties, "created")
		delete(additionalProperties, "creator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflow struct {
	value *Workflow
	isSet bool
}

func (v NullableWorkflow) Get() *Workflow {
	return v.value
}

func (v *NullableWorkflow) Set(val *Workflow) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflow(val *Workflow) *NullableWorkflow {
	return &NullableWorkflow{value: val, isSet: true}
}

func (v NullableWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


