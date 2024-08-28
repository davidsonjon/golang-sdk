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

// checks if the LoadEntitlementTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadEntitlementTask{}

// LoadEntitlementTask struct for LoadEntitlementTask
type LoadEntitlementTask struct {
	// System-generated unique ID of the task this taskStatus represents
	Id *string `json:"id,omitempty"`
	// Type of task this task represents
	Type *string `json:"type,omitempty"`
	// The name of the task
	UniqueName *string `json:"uniqueName,omitempty"`
	// The description of the task
	Description *string `json:"description,omitempty"`
	// The user who initiated the task
	Launcher *string `json:"launcher,omitempty"`
	// The creation date of the task
	Created *time.Time `json:"created,omitempty"`
	// Return values from the task
	Returns []LoadEntitlementTaskReturnsInner `json:"returns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadEntitlementTask LoadEntitlementTask

// NewLoadEntitlementTask instantiates a new LoadEntitlementTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadEntitlementTask() *LoadEntitlementTask {
	this := LoadEntitlementTask{}
	return &this
}

// NewLoadEntitlementTaskWithDefaults instantiates a new LoadEntitlementTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadEntitlementTaskWithDefaults() *LoadEntitlementTask {
	this := LoadEntitlementTask{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadEntitlementTask) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LoadEntitlementTask) SetType(v string) {
	o.Type = &v
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetUniqueName() string {
	if o == nil || IsNil(o.UniqueName) {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetUniqueNameOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueName) {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasUniqueName() bool {
	if o != nil && !IsNil(o.UniqueName) {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *LoadEntitlementTask) SetUniqueName(v string) {
	o.UniqueName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadEntitlementTask) SetDescription(v string) {
	o.Description = &v
}

// GetLauncher returns the Launcher field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetLauncher() string {
	if o == nil || IsNil(o.Launcher) {
		var ret string
		return ret
	}
	return *o.Launcher
}

// GetLauncherOk returns a tuple with the Launcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetLauncherOk() (*string, bool) {
	if o == nil || IsNil(o.Launcher) {
		return nil, false
	}
	return o.Launcher, true
}

// HasLauncher returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasLauncher() bool {
	if o != nil && !IsNil(o.Launcher) {
		return true
	}

	return false
}

// SetLauncher gets a reference to the given string and assigns it to the Launcher field.
func (o *LoadEntitlementTask) SetLauncher(v string) {
	o.Launcher = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *LoadEntitlementTask) SetCreated(v time.Time) {
	o.Created = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *LoadEntitlementTask) GetReturns() []LoadEntitlementTaskReturnsInner {
	if o == nil || IsNil(o.Returns) {
		var ret []LoadEntitlementTaskReturnsInner
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadEntitlementTask) GetReturnsOk() ([]LoadEntitlementTaskReturnsInner, bool) {
	if o == nil || IsNil(o.Returns) {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *LoadEntitlementTask) HasReturns() bool {
	if o != nil && !IsNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []LoadEntitlementTaskReturnsInner and assigns it to the Returns field.
func (o *LoadEntitlementTask) SetReturns(v []LoadEntitlementTaskReturnsInner) {
	o.Returns = v
}

func (o LoadEntitlementTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadEntitlementTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UniqueName) {
		toSerialize["uniqueName"] = o.UniqueName
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
	if !IsNil(o.Returns) {
		toSerialize["returns"] = o.Returns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadEntitlementTask) UnmarshalJSON(data []byte) (err error) {
	varLoadEntitlementTask := _LoadEntitlementTask{}

	err = json.Unmarshal(data, &varLoadEntitlementTask)

	if err != nil {
		return err
	}

	*o = LoadEntitlementTask(varLoadEntitlementTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uniqueName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "launcher")
		delete(additionalProperties, "created")
		delete(additionalProperties, "returns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadEntitlementTask struct {
	value *LoadEntitlementTask
	isSet bool
}

func (v NullableLoadEntitlementTask) Get() *LoadEntitlementTask {
	return v.value
}

func (v *NullableLoadEntitlementTask) Set(val *LoadEntitlementTask) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadEntitlementTask) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadEntitlementTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadEntitlementTask(val *LoadEntitlementTask) *NullableLoadEntitlementTask {
	return &NullableLoadEntitlementTask{value: val, isSet: true}
}

func (v NullableLoadEntitlementTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadEntitlementTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


