/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the TaskResultSimplified type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResultSimplified{}

// TaskResultSimplified struct for TaskResultSimplified
type TaskResultSimplified struct {
	// Task identifier
	Id *string `json:"id,omitempty"`
	// Task name
	Name *string `json:"name,omitempty"`
	// Task description
	Description *string `json:"description,omitempty"`
	// User or process who launched the task
	Launcher *string `json:"launcher,omitempty"`
	// Date time of completion
	Completed *time.Time `json:"completed,omitempty"`
	// Date time when the task was launched
	Launched *time.Time `json:"launched,omitempty"`
	// Task result status
	CompletionStatus *string `json:"completionStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskResultSimplified TaskResultSimplified

// NewTaskResultSimplified instantiates a new TaskResultSimplified object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResultSimplified() *TaskResultSimplified {
	this := TaskResultSimplified{}
	return &this
}

// NewTaskResultSimplifiedWithDefaults instantiates a new TaskResultSimplified object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResultSimplifiedWithDefaults() *TaskResultSimplified {
	this := TaskResultSimplified{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskResultSimplified) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskResultSimplified) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TaskResultSimplified) SetDescription(v string) {
	o.Description = &v
}

// GetLauncher returns the Launcher field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetLauncher() string {
	if o == nil || IsNil(o.Launcher) {
		var ret string
		return ret
	}
	return *o.Launcher
}

// GetLauncherOk returns a tuple with the Launcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetLauncherOk() (*string, bool) {
	if o == nil || IsNil(o.Launcher) {
		return nil, false
	}
	return o.Launcher, true
}

// HasLauncher returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasLauncher() bool {
	if o != nil && !IsNil(o.Launcher) {
		return true
	}

	return false
}

// SetLauncher gets a reference to the given string and assigns it to the Launcher field.
func (o *TaskResultSimplified) SetLauncher(v string) {
	o.Launcher = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *TaskResultSimplified) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetLaunched returns the Launched field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetLaunched() time.Time {
	if o == nil || IsNil(o.Launched) {
		var ret time.Time
		return ret
	}
	return *o.Launched
}

// GetLaunchedOk returns a tuple with the Launched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetLaunchedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Launched) {
		return nil, false
	}
	return o.Launched, true
}

// HasLaunched returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasLaunched() bool {
	if o != nil && !IsNil(o.Launched) {
		return true
	}

	return false
}

// SetLaunched gets a reference to the given time.Time and assigns it to the Launched field.
func (o *TaskResultSimplified) SetLaunched(v time.Time) {
	o.Launched = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *TaskResultSimplified) GetCompletionStatus() string {
	if o == nil || IsNil(o.CompletionStatus) {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultSimplified) GetCompletionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *TaskResultSimplified) HasCompletionStatus() bool {
	if o != nil && !IsNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *TaskResultSimplified) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

func (o TaskResultSimplified) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskResultSimplified) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Launcher) {
		toSerialize["launcher"] = o.Launcher
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Launched) {
		toSerialize["launched"] = o.Launched
	}
	if !IsNil(o.CompletionStatus) {
		toSerialize["completionStatus"] = o.CompletionStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskResultSimplified) UnmarshalJSON(data []byte) (err error) {
	varTaskResultSimplified := _TaskResultSimplified{}

	err = json.Unmarshal(data, &varTaskResultSimplified)

	if err != nil {
		return err
	}

	*o = TaskResultSimplified(varTaskResultSimplified)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "launcher")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "launched")
		delete(additionalProperties, "completionStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskResultSimplified struct {
	value *TaskResultSimplified
	isSet bool
}

func (v NullableTaskResultSimplified) Get() *TaskResultSimplified {
	return v.value
}

func (v *NullableTaskResultSimplified) Set(val *TaskResultSimplified) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResultSimplified) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResultSimplified) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResultSimplified(val *TaskResultSimplified) *NullableTaskResultSimplified {
	return &NullableTaskResultSimplified{value: val, isSet: true}
}

func (v NullableTaskResultSimplified) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResultSimplified) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


