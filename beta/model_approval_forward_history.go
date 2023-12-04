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

// checks if the ApprovalForwardHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalForwardHistory{}

// ApprovalForwardHistory struct for ApprovalForwardHistory
type ApprovalForwardHistory struct {
	// Display name of approver from whom the approval was forwarded.
	OldApproverName *string `json:"oldApproverName,omitempty"`
	// Display name of approver to whom the approval was forwarded.
	NewApproverName *string `json:"newApproverName,omitempty"`
	// Comment made while forwarding.
	Comment NullableString `json:"comment,omitempty"`
	// Time at which approval was forwarded.
	Modified *time.Time `json:"modified,omitempty"`
	// Display name of forwarder who forwarded the approval.
	ForwarderName NullableString `json:"forwarderName,omitempty"`
	ReassignmentType *ReassignmentType `json:"reassignmentType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalForwardHistory ApprovalForwardHistory

// NewApprovalForwardHistory instantiates a new ApprovalForwardHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalForwardHistory() *ApprovalForwardHistory {
	this := ApprovalForwardHistory{}
	return &this
}

// NewApprovalForwardHistoryWithDefaults instantiates a new ApprovalForwardHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalForwardHistoryWithDefaults() *ApprovalForwardHistory {
	this := ApprovalForwardHistory{}
	return &this
}

// GetOldApproverName returns the OldApproverName field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetOldApproverName() string {
	if o == nil || isNil(o.OldApproverName) {
		var ret string
		return ret
	}
	return *o.OldApproverName
}

// GetOldApproverNameOk returns a tuple with the OldApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetOldApproverNameOk() (*string, bool) {
	if o == nil || isNil(o.OldApproverName) {
		return nil, false
	}
	return o.OldApproverName, true
}

// HasOldApproverName returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasOldApproverName() bool {
	if o != nil && !isNil(o.OldApproverName) {
		return true
	}

	return false
}

// SetOldApproverName gets a reference to the given string and assigns it to the OldApproverName field.
func (o *ApprovalForwardHistory) SetOldApproverName(v string) {
	o.OldApproverName = &v
}

// GetNewApproverName returns the NewApproverName field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetNewApproverName() string {
	if o == nil || isNil(o.NewApproverName) {
		var ret string
		return ret
	}
	return *o.NewApproverName
}

// GetNewApproverNameOk returns a tuple with the NewApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetNewApproverNameOk() (*string, bool) {
	if o == nil || isNil(o.NewApproverName) {
		return nil, false
	}
	return o.NewApproverName, true
}

// HasNewApproverName returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasNewApproverName() bool {
	if o != nil && !isNil(o.NewApproverName) {
		return true
	}

	return false
}

// SetNewApproverName gets a reference to the given string and assigns it to the NewApproverName field.
func (o *ApprovalForwardHistory) SetNewApproverName(v string) {
	o.NewApproverName = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalForwardHistory) GetComment() string {
	if o == nil || isNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalForwardHistory) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ApprovalForwardHistory) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ApprovalForwardHistory) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ApprovalForwardHistory) UnsetComment() {
	o.Comment.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ApprovalForwardHistory) SetModified(v time.Time) {
	o.Modified = &v
}

// GetForwarderName returns the ForwarderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalForwardHistory) GetForwarderName() string {
	if o == nil || isNil(o.ForwarderName.Get()) {
		var ret string
		return ret
	}
	return *o.ForwarderName.Get()
}

// GetForwarderNameOk returns a tuple with the ForwarderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalForwardHistory) GetForwarderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwarderName.Get(), o.ForwarderName.IsSet()
}

// HasForwarderName returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasForwarderName() bool {
	if o != nil && o.ForwarderName.IsSet() {
		return true
	}

	return false
}

// SetForwarderName gets a reference to the given NullableString and assigns it to the ForwarderName field.
func (o *ApprovalForwardHistory) SetForwarderName(v string) {
	o.ForwarderName.Set(&v)
}
// SetForwarderNameNil sets the value for ForwarderName to be an explicit nil
func (o *ApprovalForwardHistory) SetForwarderNameNil() {
	o.ForwarderName.Set(nil)
}

// UnsetForwarderName ensures that no value is present for ForwarderName, not even an explicit nil
func (o *ApprovalForwardHistory) UnsetForwarderName() {
	o.ForwarderName.Unset()
}

// GetReassignmentType returns the ReassignmentType field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetReassignmentType() ReassignmentType {
	if o == nil || isNil(o.ReassignmentType) {
		var ret ReassignmentType
		return ret
	}
	return *o.ReassignmentType
}

// GetReassignmentTypeOk returns a tuple with the ReassignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetReassignmentTypeOk() (*ReassignmentType, bool) {
	if o == nil || isNil(o.ReassignmentType) {
		return nil, false
	}
	return o.ReassignmentType, true
}

// HasReassignmentType returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasReassignmentType() bool {
	if o != nil && !isNil(o.ReassignmentType) {
		return true
	}

	return false
}

// SetReassignmentType gets a reference to the given ReassignmentType and assigns it to the ReassignmentType field.
func (o *ApprovalForwardHistory) SetReassignmentType(v ReassignmentType) {
	o.ReassignmentType = &v
}

func (o ApprovalForwardHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalForwardHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OldApproverName) {
		toSerialize["oldApproverName"] = o.OldApproverName
	}
	if !isNil(o.NewApproverName) {
		toSerialize["newApproverName"] = o.NewApproverName
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if o.ForwarderName.IsSet() {
		toSerialize["forwarderName"] = o.ForwarderName.Get()
	}
	if !isNil(o.ReassignmentType) {
		toSerialize["reassignmentType"] = o.ReassignmentType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalForwardHistory) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalForwardHistory := _ApprovalForwardHistory{}

	if err = json.Unmarshal(bytes, &varApprovalForwardHistory); err == nil {
	*o = ApprovalForwardHistory(varApprovalForwardHistory)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "oldApproverName")
		delete(additionalProperties, "newApproverName")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "forwarderName")
		delete(additionalProperties, "reassignmentType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalForwardHistory struct {
	value *ApprovalForwardHistory
	isSet bool
}

func (v NullableApprovalForwardHistory) Get() *ApprovalForwardHistory {
	return v.value
}

func (v *NullableApprovalForwardHistory) Set(val *ApprovalForwardHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalForwardHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalForwardHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalForwardHistory(val *ApprovalForwardHistory) *NullableApprovalForwardHistory {
	return &NullableApprovalForwardHistory{value: val, isSet: true}
}

func (v NullableApprovalForwardHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalForwardHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


