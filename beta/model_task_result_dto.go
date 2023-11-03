/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the TaskResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResultDto{}

// TaskResultDto Task result.
type TaskResultDto struct {
	// Task result DTO type.
	Type *string `json:"type,omitempty"`
	// Task result ID.
	Id *string `json:"id,omitempty"`
	// Task result display name.
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskResultDto TaskResultDto

// NewTaskResultDto instantiates a new TaskResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResultDto() *TaskResultDto {
	this := TaskResultDto{}
	return &this
}

// NewTaskResultDtoWithDefaults instantiates a new TaskResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResultDtoWithDefaults() *TaskResultDto {
	this := TaskResultDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskResultDto) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultDto) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskResultDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaskResultDto) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskResultDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResultDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskResultDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskResultDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResultDto) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResultDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TaskResultDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TaskResultDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TaskResultDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TaskResultDto) UnsetName() {
	o.Name.Unset()
}

func (o TaskResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskResultDto) UnmarshalJSON(bytes []byte) (err error) {
	varTaskResultDto := _TaskResultDto{}

	if err = json.Unmarshal(bytes, &varTaskResultDto); err == nil {
		*o = TaskResultDto(varTaskResultDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskResultDto struct {
	value *TaskResultDto
	isSet bool
}

func (v NullableTaskResultDto) Get() *TaskResultDto {
	return v.value
}

func (v *NullableTaskResultDto) Set(val *TaskResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResultDto(val *TaskResultDto) *NullableTaskResultDto {
	return &NullableTaskResultDto{value: val, isSet: true}
}

func (v NullableTaskResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

