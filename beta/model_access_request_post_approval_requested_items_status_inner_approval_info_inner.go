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

// checks if the AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner{}

// AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner struct for AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner
type AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner struct {
	// A comment left by the approver.
	ApprovalComment NullableString `json:"approvalComment,omitempty"`
	// The final decision of the approver.
	ApprovalDecision map[string]interface{} `json:"approvalDecision"`
	// The name of the approver
	ApproverName string `json:"approverName"`
	Approver AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover `json:"approver"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner

// NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner instantiates a new AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner(approvalDecision map[string]interface{}, approverName string, approver AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner {
	this := AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner{}
	this.ApprovalDecision = approvalDecision
	this.ApproverName = approverName
	this.Approver = approver
	return &this
}

// NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerWithDefaults instantiates a new AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerWithDefaults() *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner {
	this := AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner{}
	return &this
}

// GetApprovalComment returns the ApprovalComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalComment() string {
	if o == nil || isNil(o.ApprovalComment.Get()) {
		var ret string
		return ret
	}
	return *o.ApprovalComment.Get()
}

// GetApprovalCommentOk returns a tuple with the ApprovalComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalComment.Get(), o.ApprovalComment.IsSet()
}

// HasApprovalComment returns a boolean if a field has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) HasApprovalComment() bool {
	if o != nil && o.ApprovalComment.IsSet() {
		return true
	}

	return false
}

// SetApprovalComment gets a reference to the given NullableString and assigns it to the ApprovalComment field.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprovalComment(v string) {
	o.ApprovalComment.Set(&v)
}
// SetApprovalCommentNil sets the value for ApprovalComment to be an explicit nil
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprovalCommentNil() {
	o.ApprovalComment.Set(nil)
}

// UnsetApprovalComment ensures that no value is present for ApprovalComment, not even an explicit nil
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) UnsetApprovalComment() {
	o.ApprovalComment.Unset()
}

// GetApprovalDecision returns the ApprovalDecision field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalDecision() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ApprovalDecision
}

// GetApprovalDecisionOk returns a tuple with the ApprovalDecision field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalDecisionOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ApprovalDecision, true
}

// SetApprovalDecision sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprovalDecision(v map[string]interface{}) {
	o.ApprovalDecision = v
}

// GetApproverName returns the ApproverName field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApproverName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverName
}

// GetApproverNameOk returns a tuple with the ApproverName field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApproverNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverName, true
}

// SetApproverName sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApproverName(v string) {
	o.ApproverName = v
}

// GetApprover returns the Approver field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprover() AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	if o == nil {
		var ret AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover
		return ret
	}

	return o.Approver
}

// GetApproverOk returns a tuple with the Approver field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApproverOk() (*AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approver, true
}

// SetApprover sets field value
func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprover(v AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) {
	o.Approver = v
}

func (o AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApprovalComment.IsSet() {
		toSerialize["approvalComment"] = o.ApprovalComment.Get()
	}
	toSerialize["approvalDecision"] = o.ApprovalDecision
	toSerialize["approverName"] = o.ApproverName
	toSerialize["approver"] = o.Approver

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approvalDecision",
		"approverName",
		"approver",
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

	varAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner := _AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner{}

	if err = json.Unmarshal(bytes, &varAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner); err == nil {
	*o = AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner(varAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approvalComment")
		delete(additionalProperties, "approvalDecision")
		delete(additionalProperties, "approverName")
		delete(additionalProperties, "approver")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner struct {
	value *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner
	isSet bool
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) Get() *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner {
	return v.value
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) Set(val *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner(val *AccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner {
	return &NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner{value: val, isSet: true}
}

func (v NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


