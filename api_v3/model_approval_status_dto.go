/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
)

// checks if the ApprovalStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalStatusDto{}

// ApprovalStatusDto struct for ApprovalStatusDto
type ApprovalStatusDto struct {
	// True if the request for this item was forwarded from one owner to another.
	Forwarded *bool `json:"forwarded,omitempty"`
	OriginalOwner *ApprovalStatusDtoOriginalOwner `json:"originalOwner,omitempty"`
	CurrentOwner *ApprovalStatusDtoCurrentOwner `json:"currentOwner,omitempty"`
	// Time at which item was modified.
	Modified NullableTime `json:"modified,omitempty"`
	Status *ManualWorkItemState `json:"status,omitempty"`
	Scheme *ApprovalScheme `json:"scheme,omitempty"`
	// If the request failed, includes any error messages that were generated.
	ErrorMessages []ErrorMessageDto `json:"errorMessages,omitempty"`
	// Comment, if any, provided by the approver.
	Comment NullableString `json:"comment,omitempty"`
	// The date the role or access profile or entitlement is no longer assigned to the specified identity.
	RemoveDate NullableTime `json:"removeDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalStatusDto ApprovalStatusDto

// NewApprovalStatusDto instantiates a new ApprovalStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalStatusDto() *ApprovalStatusDto {
	this := ApprovalStatusDto{}
	var forwarded bool = false
	this.Forwarded = &forwarded
	return &this
}

// NewApprovalStatusDtoWithDefaults instantiates a new ApprovalStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalStatusDtoWithDefaults() *ApprovalStatusDto {
	this := ApprovalStatusDto{}
	var forwarded bool = false
	this.Forwarded = &forwarded
	return &this
}

// GetForwarded returns the Forwarded field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetForwarded() bool {
	if o == nil || IsNil(o.Forwarded) {
		var ret bool
		return ret
	}
	return *o.Forwarded
}

// GetForwardedOk returns a tuple with the Forwarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetForwardedOk() (*bool, bool) {
	if o == nil || IsNil(o.Forwarded) {
		return nil, false
	}
	return o.Forwarded, true
}

// HasForwarded returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasForwarded() bool {
	if o != nil && !IsNil(o.Forwarded) {
		return true
	}

	return false
}

// SetForwarded gets a reference to the given bool and assigns it to the Forwarded field.
func (o *ApprovalStatusDto) SetForwarded(v bool) {
	o.Forwarded = &v
}

// GetOriginalOwner returns the OriginalOwner field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetOriginalOwner() ApprovalStatusDtoOriginalOwner {
	if o == nil || IsNil(o.OriginalOwner) {
		var ret ApprovalStatusDtoOriginalOwner
		return ret
	}
	return *o.OriginalOwner
}

// GetOriginalOwnerOk returns a tuple with the OriginalOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetOriginalOwnerOk() (*ApprovalStatusDtoOriginalOwner, bool) {
	if o == nil || IsNil(o.OriginalOwner) {
		return nil, false
	}
	return o.OriginalOwner, true
}

// HasOriginalOwner returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasOriginalOwner() bool {
	if o != nil && !IsNil(o.OriginalOwner) {
		return true
	}

	return false
}

// SetOriginalOwner gets a reference to the given ApprovalStatusDtoOriginalOwner and assigns it to the OriginalOwner field.
func (o *ApprovalStatusDto) SetOriginalOwner(v ApprovalStatusDtoOriginalOwner) {
	o.OriginalOwner = &v
}

// GetCurrentOwner returns the CurrentOwner field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetCurrentOwner() ApprovalStatusDtoCurrentOwner {
	if o == nil || IsNil(o.CurrentOwner) {
		var ret ApprovalStatusDtoCurrentOwner
		return ret
	}
	return *o.CurrentOwner
}

// GetCurrentOwnerOk returns a tuple with the CurrentOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetCurrentOwnerOk() (*ApprovalStatusDtoCurrentOwner, bool) {
	if o == nil || IsNil(o.CurrentOwner) {
		return nil, false
	}
	return o.CurrentOwner, true
}

// HasCurrentOwner returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasCurrentOwner() bool {
	if o != nil && !IsNil(o.CurrentOwner) {
		return true
	}

	return false
}

// SetCurrentOwner gets a reference to the given ApprovalStatusDtoCurrentOwner and assigns it to the CurrentOwner field.
func (o *ApprovalStatusDto) SetCurrentOwner(v ApprovalStatusDtoCurrentOwner) {
	o.CurrentOwner = &v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalStatusDto) GetModified() time.Time {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalStatusDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *ApprovalStatusDto) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *ApprovalStatusDto) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *ApprovalStatusDto) UnsetModified() {
	o.Modified.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetStatus() ManualWorkItemState {
	if o == nil || IsNil(o.Status) {
		var ret ManualWorkItemState
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetStatusOk() (*ManualWorkItemState, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ManualWorkItemState and assigns it to the Status field.
func (o *ApprovalStatusDto) SetStatus(v ManualWorkItemState) {
	o.Status = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetScheme() ApprovalScheme {
	if o == nil || IsNil(o.Scheme) {
		var ret ApprovalScheme
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetSchemeOk() (*ApprovalScheme, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given ApprovalScheme and assigns it to the Scheme field.
func (o *ApprovalStatusDto) SetScheme(v ApprovalScheme) {
	o.Scheme = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalStatusDto) GetErrorMessages() []ErrorMessageDto {
	if o == nil {
		var ret []ErrorMessageDto
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalStatusDto) GetErrorMessagesOk() ([]ErrorMessageDto, bool) {
	if o == nil || IsNil(o.ErrorMessages) {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasErrorMessages() bool {
	if o != nil && !IsNil(o.ErrorMessages) {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []ErrorMessageDto and assigns it to the ErrorMessages field.
func (o *ApprovalStatusDto) SetErrorMessages(v []ErrorMessageDto) {
	o.ErrorMessages = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalStatusDto) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalStatusDto) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ApprovalStatusDto) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ApprovalStatusDto) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ApprovalStatusDto) UnsetComment() {
	o.Comment.Unset()
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalStatusDto) GetRemoveDate() time.Time {
	if o == nil || IsNil(o.RemoveDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate.Get()
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalStatusDto) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveDate.Get(), o.RemoveDate.IsSet()
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasRemoveDate() bool {
	if o != nil && o.RemoveDate.IsSet() {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given NullableTime and assigns it to the RemoveDate field.
func (o *ApprovalStatusDto) SetRemoveDate(v time.Time) {
	o.RemoveDate.Set(&v)
}
// SetRemoveDateNil sets the value for RemoveDate to be an explicit nil
func (o *ApprovalStatusDto) SetRemoveDateNil() {
	o.RemoveDate.Set(nil)
}

// UnsetRemoveDate ensures that no value is present for RemoveDate, not even an explicit nil
func (o *ApprovalStatusDto) UnsetRemoveDate() {
	o.RemoveDate.Unset()
}

func (o ApprovalStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Forwarded) {
		toSerialize["forwarded"] = o.Forwarded
	}
	if !IsNil(o.OriginalOwner) {
		toSerialize["originalOwner"] = o.OriginalOwner
	}
	if !IsNil(o.CurrentOwner) {
		toSerialize["currentOwner"] = o.CurrentOwner
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if o.ErrorMessages != nil {
		toSerialize["errorMessages"] = o.ErrorMessages
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.RemoveDate.IsSet() {
		toSerialize["removeDate"] = o.RemoveDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalStatusDto) UnmarshalJSON(data []byte) (err error) {
	varApprovalStatusDto := _ApprovalStatusDto{}

	err = json.Unmarshal(data, &varApprovalStatusDto)

	if err != nil {
		return err
	}

	*o = ApprovalStatusDto(varApprovalStatusDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "forwarded")
		delete(additionalProperties, "originalOwner")
		delete(additionalProperties, "currentOwner")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "status")
		delete(additionalProperties, "scheme")
		delete(additionalProperties, "errorMessages")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "removeDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalStatusDto struct {
	value *ApprovalStatusDto
	isSet bool
}

func (v NullableApprovalStatusDto) Get() *ApprovalStatusDto {
	return v.value
}

func (v *NullableApprovalStatusDto) Set(val *ApprovalStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalStatusDto(val *ApprovalStatusDto) *NullableApprovalStatusDto {
	return &NullableApprovalStatusDto{value: val, isSet: true}
}

func (v NullableApprovalStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


