/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the UpdateDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDetail{}

// UpdateDetail struct for UpdateDetail
type UpdateDetail struct {
	// The detailed message for an update. Typically the relevent error message when status is error.
	Message *string `json:"message,omitempty"`
	// The connector script name
	ScriptName *string `json:"scriptName,omitempty"`
	// The list of updated files supported by the connector
	UpdatedFiles []string `json:"updatedFiles,omitempty"`
	// The connector update status
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDetail UpdateDetail

// NewUpdateDetail instantiates a new UpdateDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDetail() *UpdateDetail {
	this := UpdateDetail{}
	return &this
}

// NewUpdateDetailWithDefaults instantiates a new UpdateDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDetailWithDefaults() *UpdateDetail {
	this := UpdateDetail{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateDetail) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDetail) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateDetail) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateDetail) SetMessage(v string) {
	o.Message = &v
}

// GetScriptName returns the ScriptName field value if set, zero value otherwise.
func (o *UpdateDetail) GetScriptName() string {
	if o == nil || isNil(o.ScriptName) {
		var ret string
		return ret
	}
	return *o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDetail) GetScriptNameOk() (*string, bool) {
	if o == nil || isNil(o.ScriptName) {
		return nil, false
	}
	return o.ScriptName, true
}

// HasScriptName returns a boolean if a field has been set.
func (o *UpdateDetail) HasScriptName() bool {
	if o != nil && !isNil(o.ScriptName) {
		return true
	}

	return false
}

// SetScriptName gets a reference to the given string and assigns it to the ScriptName field.
func (o *UpdateDetail) SetScriptName(v string) {
	o.ScriptName = &v
}

// GetUpdatedFiles returns the UpdatedFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDetail) GetUpdatedFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UpdatedFiles
}

// GetUpdatedFilesOk returns a tuple with the UpdatedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDetail) GetUpdatedFilesOk() ([]string, bool) {
	if o == nil || isNil(o.UpdatedFiles) {
		return nil, false
	}
	return o.UpdatedFiles, true
}

// HasUpdatedFiles returns a boolean if a field has been set.
func (o *UpdateDetail) HasUpdatedFiles() bool {
	if o != nil && isNil(o.UpdatedFiles) {
		return true
	}

	return false
}

// SetUpdatedFiles gets a reference to the given []string and assigns it to the UpdatedFiles field.
func (o *UpdateDetail) SetUpdatedFiles(v []string) {
	o.UpdatedFiles = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateDetail) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDetail) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateDetail) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateDetail) SetStatus(v string) {
	o.Status = &v
}

func (o UpdateDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ScriptName) {
		toSerialize["scriptName"] = o.ScriptName
	}
	if o.UpdatedFiles != nil {
		toSerialize["updatedFiles"] = o.UpdatedFiles
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDetail) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateDetail := _UpdateDetail{}

	if err = json.Unmarshal(bytes, &varUpdateDetail); err == nil {
	*o = UpdateDetail(varUpdateDetail)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "scriptName")
		delete(additionalProperties, "updatedFiles")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDetail struct {
	value *UpdateDetail
	isSet bool
}

func (v NullableUpdateDetail) Get() *UpdateDetail {
	return v.value
}

func (v *NullableUpdateDetail) Set(val *UpdateDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDetail(val *UpdateDetail) *NullableUpdateDetail {
	return &NullableUpdateDetail{value: val, isSet: true}
}

func (v NullableUpdateDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


