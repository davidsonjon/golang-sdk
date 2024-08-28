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

// checks if the LoadAccountsTaskTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadAccountsTaskTask{}

// LoadAccountsTaskTask struct for LoadAccountsTaskTask
type LoadAccountsTaskTask struct {
	// System-generated unique ID of the task this taskStatus represents
	Id *string `json:"id,omitempty"`
	// Type of task this task represents
	Type *string `json:"type,omitempty"`
	// The name of the aggregation process
	Name *string `json:"name,omitempty"`
	// The description of the task
	Description *string `json:"description,omitempty"`
	// The user who initiated the task
	Launcher *string `json:"launcher,omitempty"`
	// The Task creation date
	Created *time.Time `json:"created,omitempty"`
	// The task start date
	Launched NullableTime `json:"launched,omitempty"`
	// The task completion date
	Completed NullableTime `json:"completed,omitempty"`
	// Task completion status.
	CompletionStatus NullableString `json:"completionStatus,omitempty"`
	// Name of the parent task if exists.
	ParentName NullableString `json:"parentName,omitempty"`
	// List of the messages dedicated to the report.  From task definition perspective here usually should be warnings or errors.
	Messages []LoadAccountsTaskTaskMessagesInner `json:"messages,omitempty"`
	// Current task state.
	Progress NullableString `json:"progress,omitempty"`
	Attributes *LoadAccountsTaskTaskAttributes `json:"attributes,omitempty"`
	// Return values from the task
	Returns []LoadAccountsTaskTaskReturnsInner `json:"returns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadAccountsTaskTask LoadAccountsTaskTask

// NewLoadAccountsTaskTask instantiates a new LoadAccountsTaskTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadAccountsTaskTask() *LoadAccountsTaskTask {
	this := LoadAccountsTaskTask{}
	return &this
}

// NewLoadAccountsTaskTaskWithDefaults instantiates a new LoadAccountsTaskTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadAccountsTaskTaskWithDefaults() *LoadAccountsTaskTask {
	this := LoadAccountsTaskTask{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadAccountsTaskTask) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LoadAccountsTaskTask) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadAccountsTaskTask) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadAccountsTaskTask) SetDescription(v string) {
	o.Description = &v
}

// GetLauncher returns the Launcher field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetLauncher() string {
	if o == nil || IsNil(o.Launcher) {
		var ret string
		return ret
	}
	return *o.Launcher
}

// GetLauncherOk returns a tuple with the Launcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetLauncherOk() (*string, bool) {
	if o == nil || IsNil(o.Launcher) {
		return nil, false
	}
	return o.Launcher, true
}

// HasLauncher returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasLauncher() bool {
	if o != nil && !IsNil(o.Launcher) {
		return true
	}

	return false
}

// SetLauncher gets a reference to the given string and assigns it to the Launcher field.
func (o *LoadAccountsTaskTask) SetLauncher(v string) {
	o.Launcher = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *LoadAccountsTaskTask) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLaunched returns the Launched field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadAccountsTaskTask) GetLaunched() time.Time {
	if o == nil || IsNil(o.Launched.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Launched.Get()
}

// GetLaunchedOk returns a tuple with the Launched field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadAccountsTaskTask) GetLaunchedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Launched.Get(), o.Launched.IsSet()
}

// HasLaunched returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasLaunched() bool {
	if o != nil && o.Launched.IsSet() {
		return true
	}

	return false
}

// SetLaunched gets a reference to the given NullableTime and assigns it to the Launched field.
func (o *LoadAccountsTaskTask) SetLaunched(v time.Time) {
	o.Launched.Set(&v)
}
// SetLaunchedNil sets the value for Launched to be an explicit nil
func (o *LoadAccountsTaskTask) SetLaunchedNil() {
	o.Launched.Set(nil)
}

// UnsetLaunched ensures that no value is present for Launched, not even an explicit nil
func (o *LoadAccountsTaskTask) UnsetLaunched() {
	o.Launched.Unset()
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadAccountsTaskTask) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadAccountsTaskTask) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasCompleted() bool {
	if o != nil && o.Completed.IsSet() {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given NullableTime and assigns it to the Completed field.
func (o *LoadAccountsTaskTask) SetCompleted(v time.Time) {
	o.Completed.Set(&v)
}
// SetCompletedNil sets the value for Completed to be an explicit nil
func (o *LoadAccountsTaskTask) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
func (o *LoadAccountsTaskTask) UnsetCompleted() {
	o.Completed.Unset()
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadAccountsTaskTask) GetCompletionStatus() string {
	if o == nil || IsNil(o.CompletionStatus.Get()) {
		var ret string
		return ret
	}
	return *o.CompletionStatus.Get()
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadAccountsTaskTask) GetCompletionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletionStatus.Get(), o.CompletionStatus.IsSet()
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasCompletionStatus() bool {
	if o != nil && o.CompletionStatus.IsSet() {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given NullableString and assigns it to the CompletionStatus field.
func (o *LoadAccountsTaskTask) SetCompletionStatus(v string) {
	o.CompletionStatus.Set(&v)
}
// SetCompletionStatusNil sets the value for CompletionStatus to be an explicit nil
func (o *LoadAccountsTaskTask) SetCompletionStatusNil() {
	o.CompletionStatus.Set(nil)
}

// UnsetCompletionStatus ensures that no value is present for CompletionStatus, not even an explicit nil
func (o *LoadAccountsTaskTask) UnsetCompletionStatus() {
	o.CompletionStatus.Unset()
}

// GetParentName returns the ParentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadAccountsTaskTask) GetParentName() string {
	if o == nil || IsNil(o.ParentName.Get()) {
		var ret string
		return ret
	}
	return *o.ParentName.Get()
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadAccountsTaskTask) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentName.Get(), o.ParentName.IsSet()
}

// HasParentName returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasParentName() bool {
	if o != nil && o.ParentName.IsSet() {
		return true
	}

	return false
}

// SetParentName gets a reference to the given NullableString and assigns it to the ParentName field.
func (o *LoadAccountsTaskTask) SetParentName(v string) {
	o.ParentName.Set(&v)
}
// SetParentNameNil sets the value for ParentName to be an explicit nil
func (o *LoadAccountsTaskTask) SetParentNameNil() {
	o.ParentName.Set(nil)
}

// UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
func (o *LoadAccountsTaskTask) UnsetParentName() {
	o.ParentName.Unset()
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetMessages() []LoadAccountsTaskTaskMessagesInner {
	if o == nil || IsNil(o.Messages) {
		var ret []LoadAccountsTaskTaskMessagesInner
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetMessagesOk() ([]LoadAccountsTaskTaskMessagesInner, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []LoadAccountsTaskTaskMessagesInner and assigns it to the Messages field.
func (o *LoadAccountsTaskTask) SetMessages(v []LoadAccountsTaskTaskMessagesInner) {
	o.Messages = v
}

// GetProgress returns the Progress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadAccountsTaskTask) GetProgress() string {
	if o == nil || IsNil(o.Progress.Get()) {
		var ret string
		return ret
	}
	return *o.Progress.Get()
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadAccountsTaskTask) GetProgressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress.Get(), o.Progress.IsSet()
}

// HasProgress returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasProgress() bool {
	if o != nil && o.Progress.IsSet() {
		return true
	}

	return false
}

// SetProgress gets a reference to the given NullableString and assigns it to the Progress field.
func (o *LoadAccountsTaskTask) SetProgress(v string) {
	o.Progress.Set(&v)
}
// SetProgressNil sets the value for Progress to be an explicit nil
func (o *LoadAccountsTaskTask) SetProgressNil() {
	o.Progress.Set(nil)
}

// UnsetProgress ensures that no value is present for Progress, not even an explicit nil
func (o *LoadAccountsTaskTask) UnsetProgress() {
	o.Progress.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetAttributes() LoadAccountsTaskTaskAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret LoadAccountsTaskTaskAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetAttributesOk() (*LoadAccountsTaskTaskAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given LoadAccountsTaskTaskAttributes and assigns it to the Attributes field.
func (o *LoadAccountsTaskTask) SetAttributes(v LoadAccountsTaskTaskAttributes) {
	o.Attributes = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *LoadAccountsTaskTask) GetReturns() []LoadAccountsTaskTaskReturnsInner {
	if o == nil || IsNil(o.Returns) {
		var ret []LoadAccountsTaskTaskReturnsInner
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTask) GetReturnsOk() ([]LoadAccountsTaskTaskReturnsInner, bool) {
	if o == nil || IsNil(o.Returns) {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *LoadAccountsTaskTask) HasReturns() bool {
	if o != nil && !IsNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []LoadAccountsTaskTaskReturnsInner and assigns it to the Returns field.
func (o *LoadAccountsTaskTask) SetReturns(v []LoadAccountsTaskTaskReturnsInner) {
	o.Returns = v
}

func (o LoadAccountsTaskTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadAccountsTaskTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.Launched.IsSet() {
		toSerialize["launched"] = o.Launched.Get()
	}
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	if o.CompletionStatus.IsSet() {
		toSerialize["completionStatus"] = o.CompletionStatus.Get()
	}
	if o.ParentName.IsSet() {
		toSerialize["parentName"] = o.ParentName.Get()
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if o.Progress.IsSet() {
		toSerialize["progress"] = o.Progress.Get()
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Returns) {
		toSerialize["returns"] = o.Returns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadAccountsTaskTask) UnmarshalJSON(data []byte) (err error) {
	varLoadAccountsTaskTask := _LoadAccountsTaskTask{}

	err = json.Unmarshal(data, &varLoadAccountsTaskTask)

	if err != nil {
		return err
	}

	*o = LoadAccountsTaskTask(varLoadAccountsTaskTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "launcher")
		delete(additionalProperties, "created")
		delete(additionalProperties, "launched")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "completionStatus")
		delete(additionalProperties, "parentName")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "returns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadAccountsTaskTask struct {
	value *LoadAccountsTaskTask
	isSet bool
}

func (v NullableLoadAccountsTaskTask) Get() *LoadAccountsTaskTask {
	return v.value
}

func (v *NullableLoadAccountsTaskTask) Set(val *LoadAccountsTaskTask) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadAccountsTaskTask) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadAccountsTaskTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadAccountsTaskTask(val *LoadAccountsTaskTask) *NullableLoadAccountsTaskTask {
	return &NullableLoadAccountsTaskTask{value: val, isSet: true}
}

func (v NullableLoadAccountsTaskTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadAccountsTaskTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


