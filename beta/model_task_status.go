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

// checks if the TaskStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskStatus{}

// TaskStatus Details and current status of a specific task
type TaskStatus struct {
	// System-generated unique ID of the task this TaskStatus represents
	Id string `json:"id"`
	// Type of task this TaskStatus represents
	Type string `json:"type"`
	// Name of the task this TaskStatus represents
	UniqueName string `json:"uniqueName"`
	// Description of the task this TaskStatus represents
	Description string `json:"description"`
	// Name of the parent of the task this TaskStatus represents
	ParentName string `json:"parentName"`
	// Service to execute the task this TaskStatus represents
	Launcher string `json:"launcher"`
	// Creation date of the task this TaskStatus represents
	Created time.Time `json:"created"`
	// Last modification date of the task this TaskStatus represents
	Modified time.Time `json:"modified"`
	// Launch date of the task this TaskStatus represents
	Launched time.Time `json:"launched"`
	// Completion date of the task this TaskStatus represents
	Completed time.Time `json:"completed"`
	// Completion status of the task this TaskStatus represents
	CompletionStatus string `json:"completionStatus"`
	// Messages associated with the task this TaskStatus represents
	Messages []TaskStatusMessage `json:"messages"`
	// Return values from the task this TaskStatus represents
	Returns []TaskReturnDetails `json:"returns"`
	// Attributes of the task this TaskStatus represents
	Attributes map[string]interface{} `json:"attributes"`
	// Current progress of the task this TaskStatus represents
	Progress string `json:"progress"`
	// Current percentage completion of the task this TaskStatus represents
	PercentComplete int32 `json:"percentComplete"`
	AdditionalProperties map[string]interface{}
}

type _TaskStatus TaskStatus

// NewTaskStatus instantiates a new TaskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskStatus(id string, type_ string, uniqueName string, description string, parentName string, launcher string, created time.Time, modified time.Time, launched time.Time, completed time.Time, completionStatus string, messages []TaskStatusMessage, returns []TaskReturnDetails, attributes map[string]interface{}, progress string, percentComplete int32) *TaskStatus {
	this := TaskStatus{}
	this.Id = id
	this.Type = type_
	this.UniqueName = uniqueName
	this.Description = description
	this.ParentName = parentName
	this.Launcher = launcher
	this.Created = created
	this.Modified = modified
	this.Launched = launched
	this.Completed = completed
	this.CompletionStatus = completionStatus
	this.Messages = messages
	this.Returns = returns
	this.Attributes = attributes
	this.Progress = progress
	this.PercentComplete = percentComplete
	return &this
}

// NewTaskStatusWithDefaults instantiates a new TaskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskStatusWithDefaults() *TaskStatus {
	this := TaskStatus{}
	return &this
}

// GetId returns the Id field value
func (o *TaskStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskStatus) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TaskStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaskStatus) SetType(v string) {
	o.Type = v
}

// GetUniqueName returns the UniqueName field value
func (o *TaskStatus) GetUniqueName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetUniqueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueName, true
}

// SetUniqueName sets field value
func (o *TaskStatus) SetUniqueName(v string) {
	o.UniqueName = v
}

// GetDescription returns the Description field value
func (o *TaskStatus) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TaskStatus) SetDescription(v string) {
	o.Description = v
}

// GetParentName returns the ParentName field value
func (o *TaskStatus) GetParentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentName, true
}

// SetParentName sets field value
func (o *TaskStatus) SetParentName(v string) {
	o.ParentName = v
}

// GetLauncher returns the Launcher field value
func (o *TaskStatus) GetLauncher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Launcher
}

// GetLauncherOk returns a tuple with the Launcher field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetLauncherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Launcher, true
}

// SetLauncher sets field value
func (o *TaskStatus) SetLauncher(v string) {
	o.Launcher = v
}

// GetCreated returns the Created field value
func (o *TaskStatus) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TaskStatus) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *TaskStatus) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *TaskStatus) SetModified(v time.Time) {
	o.Modified = v
}

// GetLaunched returns the Launched field value
func (o *TaskStatus) GetLaunched() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Launched
}

// GetLaunchedOk returns a tuple with the Launched field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetLaunchedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Launched, true
}

// SetLaunched sets field value
func (o *TaskStatus) SetLaunched(v time.Time) {
	o.Launched = v
}

// GetCompleted returns the Completed field value
func (o *TaskStatus) GetCompleted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *TaskStatus) SetCompleted(v time.Time) {
	o.Completed = v
}

// GetCompletionStatus returns the CompletionStatus field value
func (o *TaskStatus) GetCompletionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetCompletionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionStatus, true
}

// SetCompletionStatus sets field value
func (o *TaskStatus) SetCompletionStatus(v string) {
	o.CompletionStatus = v
}

// GetMessages returns the Messages field value
func (o *TaskStatus) GetMessages() []TaskStatusMessage {
	if o == nil {
		var ret []TaskStatusMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetMessagesOk() ([]TaskStatusMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *TaskStatus) SetMessages(v []TaskStatusMessage) {
	o.Messages = v
}

// GetReturns returns the Returns field value
func (o *TaskStatus) GetReturns() []TaskReturnDetails {
	if o == nil {
		var ret []TaskReturnDetails
		return ret
	}

	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetReturnsOk() ([]TaskReturnDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Returns, true
}

// SetReturns sets field value
func (o *TaskStatus) SetReturns(v []TaskReturnDetails) {
	o.Returns = v
}

// GetAttributes returns the Attributes field value
func (o *TaskStatus) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TaskStatus) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetProgress returns the Progress field value
func (o *TaskStatus) GetProgress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetProgressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *TaskStatus) SetProgress(v string) {
	o.Progress = v
}

// GetPercentComplete returns the PercentComplete field value
func (o *TaskStatus) GetPercentComplete() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetPercentCompleteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentComplete, true
}

// SetPercentComplete sets field value
func (o *TaskStatus) SetPercentComplete(v int32) {
	o.PercentComplete = v
}

func (o TaskStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["uniqueName"] = o.UniqueName
	toSerialize["description"] = o.Description
	toSerialize["parentName"] = o.ParentName
	toSerialize["launcher"] = o.Launcher
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	toSerialize["launched"] = o.Launched
	toSerialize["completed"] = o.Completed
	toSerialize["completionStatus"] = o.CompletionStatus
	toSerialize["messages"] = o.Messages
	toSerialize["returns"] = o.Returns
	toSerialize["attributes"] = o.Attributes
	toSerialize["progress"] = o.Progress
	toSerialize["percentComplete"] = o.PercentComplete

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskStatus) UnmarshalJSON(bytes []byte) (err error) {
	varTaskStatus := _TaskStatus{}

	if err = json.Unmarshal(bytes, &varTaskStatus); err == nil {
		*o = TaskStatus(varTaskStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uniqueName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "parentName")
		delete(additionalProperties, "launcher")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "launched")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "completionStatus")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "returns")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "percentComplete")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskStatus struct {
	value *TaskStatus
	isSet bool
}

func (v NullableTaskStatus) Get() *TaskStatus {
	return v.value
}

func (v *NullableTaskStatus) Set(val *TaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatus(val *TaskStatus) *NullableTaskStatus {
	return &NullableTaskStatus{value: val, isSet: true}
}

func (v NullableTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


