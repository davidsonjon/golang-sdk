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

// AccountActivityItem struct for AccountActivityItem
type AccountActivityItem struct {
	// Item id
	Id *string `json:"id,omitempty"`
	// Human-readable display name of item
	Name *string `json:"name,omitempty"`
	// Date and time item was requested
	Requested *time.Time `json:"requested,omitempty"`
	ApprovalStatus *WorkItemState `json:"approvalStatus,omitempty"`
	ProvisioningStatus *ProvisioningState `json:"provisioningStatus,omitempty"`
	RequesterComment NullableComment `json:"requesterComment,omitempty"`
	ReviewerIdentitySummary NullableIdentitySummary `json:"reviewerIdentitySummary,omitempty"`
	ReviewerComment NullableComment `json:"reviewerComment,omitempty"`
	Operation *AccountActivityItemOperation `json:"operation,omitempty"`
	// Attribute to which account activity applies
	Attribute NullableString `json:"attribute,omitempty"`
	// Value of attribute
	Value NullableString `json:"value,omitempty"`
	// Native identity in the target system to which the account activity applies
	NativeIdentity NullableString `json:"nativeIdentity,omitempty"`
	// Id of Source to which account activity applies
	SourceId *string `json:"sourceId,omitempty"`
	AccountRequestInfo NullableAccountRequestInfo `json:"accountRequestInfo,omitempty"`
	// Arbitrary key-value pairs, if any were included in the corresponding access request item
	ClientMetadata map[string]string `json:"clientMetadata,omitempty"`
	// The date the role or access profile is no longer assigned to the specified identity.
	RemoveDate NullableTime `json:"removeDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountActivityItem AccountActivityItem

// NewAccountActivityItem instantiates a new AccountActivityItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountActivityItem() *AccountActivityItem {
	this := AccountActivityItem{}
	return &this
}

// NewAccountActivityItemWithDefaults instantiates a new AccountActivityItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountActivityItemWithDefaults() *AccountActivityItem {
	this := AccountActivityItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountActivityItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountActivityItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountActivityItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountActivityItem) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountActivityItem) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountActivityItem) SetName(v string) {
	o.Name = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *AccountActivityItem) GetRequested() time.Time {
	if o == nil || isNil(o.Requested) {
		var ret time.Time
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetRequestedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *AccountActivityItem) HasRequested() bool {
	if o != nil && !isNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given time.Time and assigns it to the Requested field.
func (o *AccountActivityItem) SetRequested(v time.Time) {
	o.Requested = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *AccountActivityItem) GetApprovalStatus() WorkItemState {
	if o == nil || isNil(o.ApprovalStatus) {
		var ret WorkItemState
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetApprovalStatusOk() (*WorkItemState, bool) {
	if o == nil || isNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *AccountActivityItem) HasApprovalStatus() bool {
	if o != nil && !isNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given WorkItemState and assigns it to the ApprovalStatus field.
func (o *AccountActivityItem) SetApprovalStatus(v WorkItemState) {
	o.ApprovalStatus = &v
}

// GetProvisioningStatus returns the ProvisioningStatus field value if set, zero value otherwise.
func (o *AccountActivityItem) GetProvisioningStatus() ProvisioningState {
	if o == nil || isNil(o.ProvisioningStatus) {
		var ret ProvisioningState
		return ret
	}
	return *o.ProvisioningStatus
}

// GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetProvisioningStatusOk() (*ProvisioningState, bool) {
	if o == nil || isNil(o.ProvisioningStatus) {
		return nil, false
	}
	return o.ProvisioningStatus, true
}

// HasProvisioningStatus returns a boolean if a field has been set.
func (o *AccountActivityItem) HasProvisioningStatus() bool {
	if o != nil && !isNil(o.ProvisioningStatus) {
		return true
	}

	return false
}

// SetProvisioningStatus gets a reference to the given ProvisioningState and assigns it to the ProvisioningStatus field.
func (o *AccountActivityItem) SetProvisioningStatus(v ProvisioningState) {
	o.ProvisioningStatus = &v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetRequesterComment() Comment {
	if o == nil || isNil(o.RequesterComment.Get()) {
		var ret Comment
		return ret
	}
	return *o.RequesterComment.Get()
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetRequesterCommentOk() (*Comment, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterComment.Get(), o.RequesterComment.IsSet()
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *AccountActivityItem) HasRequesterComment() bool {
	if o != nil && o.RequesterComment.IsSet() {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given NullableComment and assigns it to the RequesterComment field.
func (o *AccountActivityItem) SetRequesterComment(v Comment) {
	o.RequesterComment.Set(&v)
}
// SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil
func (o *AccountActivityItem) SetRequesterCommentNil() {
	o.RequesterComment.Set(nil)
}

// UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
func (o *AccountActivityItem) UnsetRequesterComment() {
	o.RequesterComment.Unset()
}

// GetReviewerIdentitySummary returns the ReviewerIdentitySummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetReviewerIdentitySummary() IdentitySummary {
	if o == nil || isNil(o.ReviewerIdentitySummary.Get()) {
		var ret IdentitySummary
		return ret
	}
	return *o.ReviewerIdentitySummary.Get()
}

// GetReviewerIdentitySummaryOk returns a tuple with the ReviewerIdentitySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetReviewerIdentitySummaryOk() (*IdentitySummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewerIdentitySummary.Get(), o.ReviewerIdentitySummary.IsSet()
}

// HasReviewerIdentitySummary returns a boolean if a field has been set.
func (o *AccountActivityItem) HasReviewerIdentitySummary() bool {
	if o != nil && o.ReviewerIdentitySummary.IsSet() {
		return true
	}

	return false
}

// SetReviewerIdentitySummary gets a reference to the given NullableIdentitySummary and assigns it to the ReviewerIdentitySummary field.
func (o *AccountActivityItem) SetReviewerIdentitySummary(v IdentitySummary) {
	o.ReviewerIdentitySummary.Set(&v)
}
// SetReviewerIdentitySummaryNil sets the value for ReviewerIdentitySummary to be an explicit nil
func (o *AccountActivityItem) SetReviewerIdentitySummaryNil() {
	o.ReviewerIdentitySummary.Set(nil)
}

// UnsetReviewerIdentitySummary ensures that no value is present for ReviewerIdentitySummary, not even an explicit nil
func (o *AccountActivityItem) UnsetReviewerIdentitySummary() {
	o.ReviewerIdentitySummary.Unset()
}

// GetReviewerComment returns the ReviewerComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetReviewerComment() Comment {
	if o == nil || isNil(o.ReviewerComment.Get()) {
		var ret Comment
		return ret
	}
	return *o.ReviewerComment.Get()
}

// GetReviewerCommentOk returns a tuple with the ReviewerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetReviewerCommentOk() (*Comment, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewerComment.Get(), o.ReviewerComment.IsSet()
}

// HasReviewerComment returns a boolean if a field has been set.
func (o *AccountActivityItem) HasReviewerComment() bool {
	if o != nil && o.ReviewerComment.IsSet() {
		return true
	}

	return false
}

// SetReviewerComment gets a reference to the given NullableComment and assigns it to the ReviewerComment field.
func (o *AccountActivityItem) SetReviewerComment(v Comment) {
	o.ReviewerComment.Set(&v)
}
// SetReviewerCommentNil sets the value for ReviewerComment to be an explicit nil
func (o *AccountActivityItem) SetReviewerCommentNil() {
	o.ReviewerComment.Set(nil)
}

// UnsetReviewerComment ensures that no value is present for ReviewerComment, not even an explicit nil
func (o *AccountActivityItem) UnsetReviewerComment() {
	o.ReviewerComment.Unset()
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *AccountActivityItem) GetOperation() AccountActivityItemOperation {
	if o == nil || isNil(o.Operation) {
		var ret AccountActivityItemOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetOperationOk() (*AccountActivityItemOperation, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *AccountActivityItem) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given AccountActivityItemOperation and assigns it to the Operation field.
func (o *AccountActivityItem) SetOperation(v AccountActivityItemOperation) {
	o.Operation = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetAttribute() string {
	if o == nil || isNil(o.Attribute.Get()) {
		var ret string
		return ret
	}
	return *o.Attribute.Get()
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute.Get(), o.Attribute.IsSet()
}

// HasAttribute returns a boolean if a field has been set.
func (o *AccountActivityItem) HasAttribute() bool {
	if o != nil && o.Attribute.IsSet() {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given NullableString and assigns it to the Attribute field.
func (o *AccountActivityItem) SetAttribute(v string) {
	o.Attribute.Set(&v)
}
// SetAttributeNil sets the value for Attribute to be an explicit nil
func (o *AccountActivityItem) SetAttributeNil() {
	o.Attribute.Set(nil)
}

// UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
func (o *AccountActivityItem) UnsetAttribute() {
	o.Attribute.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *AccountActivityItem) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *AccountActivityItem) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *AccountActivityItem) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *AccountActivityItem) UnsetValue() {
	o.Value.Unset()
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetNativeIdentity() string {
	if o == nil || isNil(o.NativeIdentity.Get()) {
		var ret string
		return ret
	}
	return *o.NativeIdentity.Get()
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetNativeIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NativeIdentity.Get(), o.NativeIdentity.IsSet()
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *AccountActivityItem) HasNativeIdentity() bool {
	if o != nil && o.NativeIdentity.IsSet() {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given NullableString and assigns it to the NativeIdentity field.
func (o *AccountActivityItem) SetNativeIdentity(v string) {
	o.NativeIdentity.Set(&v)
}
// SetNativeIdentityNil sets the value for NativeIdentity to be an explicit nil
func (o *AccountActivityItem) SetNativeIdentityNil() {
	o.NativeIdentity.Set(nil)
}

// UnsetNativeIdentity ensures that no value is present for NativeIdentity, not even an explicit nil
func (o *AccountActivityItem) UnsetNativeIdentity() {
	o.NativeIdentity.Unset()
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AccountActivityItem) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivityItem) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AccountActivityItem) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AccountActivityItem) SetSourceId(v string) {
	o.SourceId = &v
}

// GetAccountRequestInfo returns the AccountRequestInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetAccountRequestInfo() AccountRequestInfo {
	if o == nil || isNil(o.AccountRequestInfo.Get()) {
		var ret AccountRequestInfo
		return ret
	}
	return *o.AccountRequestInfo.Get()
}

// GetAccountRequestInfoOk returns a tuple with the AccountRequestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetAccountRequestInfoOk() (*AccountRequestInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountRequestInfo.Get(), o.AccountRequestInfo.IsSet()
}

// HasAccountRequestInfo returns a boolean if a field has been set.
func (o *AccountActivityItem) HasAccountRequestInfo() bool {
	if o != nil && o.AccountRequestInfo.IsSet() {
		return true
	}

	return false
}

// SetAccountRequestInfo gets a reference to the given NullableAccountRequestInfo and assigns it to the AccountRequestInfo field.
func (o *AccountActivityItem) SetAccountRequestInfo(v AccountRequestInfo) {
	o.AccountRequestInfo.Set(&v)
}
// SetAccountRequestInfoNil sets the value for AccountRequestInfo to be an explicit nil
func (o *AccountActivityItem) SetAccountRequestInfoNil() {
	o.AccountRequestInfo.Set(nil)
}

// UnsetAccountRequestInfo ensures that no value is present for AccountRequestInfo, not even an explicit nil
func (o *AccountActivityItem) UnsetAccountRequestInfo() {
	o.AccountRequestInfo.Unset()
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetClientMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return &o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *AccountActivityItem) HasClientMetadata() bool {
	if o != nil && isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *AccountActivityItem) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivityItem) GetRemoveDate() time.Time {
	if o == nil || isNil(o.RemoveDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate.Get()
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivityItem) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveDate.Get(), o.RemoveDate.IsSet()
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *AccountActivityItem) HasRemoveDate() bool {
	if o != nil && o.RemoveDate.IsSet() {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given NullableTime and assigns it to the RemoveDate field.
func (o *AccountActivityItem) SetRemoveDate(v time.Time) {
	o.RemoveDate.Set(&v)
}
// SetRemoveDateNil sets the value for RemoveDate to be an explicit nil
func (o *AccountActivityItem) SetRemoveDateNil() {
	o.RemoveDate.Set(nil)
}

// UnsetRemoveDate ensures that no value is present for RemoveDate, not even an explicit nil
func (o *AccountActivityItem) UnsetRemoveDate() {
	o.RemoveDate.Unset()
}

func (o AccountActivityItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !isNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if !isNil(o.ProvisioningStatus) {
		toSerialize["provisioningStatus"] = o.ProvisioningStatus
	}
	if o.RequesterComment.IsSet() {
		toSerialize["requesterComment"] = o.RequesterComment.Get()
	}
	if o.ReviewerIdentitySummary.IsSet() {
		toSerialize["reviewerIdentitySummary"] = o.ReviewerIdentitySummary.Get()
	}
	if o.ReviewerComment.IsSet() {
		toSerialize["reviewerComment"] = o.ReviewerComment.Get()
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if o.Attribute.IsSet() {
		toSerialize["attribute"] = o.Attribute.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.NativeIdentity.IsSet() {
		toSerialize["nativeIdentity"] = o.NativeIdentity.Get()
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.AccountRequestInfo.IsSet() {
		toSerialize["accountRequestInfo"] = o.AccountRequestInfo.Get()
	}
	if o.ClientMetadata != nil {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if o.RemoveDate.IsSet() {
		toSerialize["removeDate"] = o.RemoveDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountActivityItem) UnmarshalJSON(bytes []byte) (err error) {
	varAccountActivityItem := _AccountActivityItem{}

	if err = json.Unmarshal(bytes, &varAccountActivityItem); err == nil {
		*o = AccountActivityItem(varAccountActivityItem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "requested")
		delete(additionalProperties, "approvalStatus")
		delete(additionalProperties, "provisioningStatus")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "reviewerIdentitySummary")
		delete(additionalProperties, "reviewerComment")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "accountRequestInfo")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "removeDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountActivityItem struct {
	value *AccountActivityItem
	isSet bool
}

func (v NullableAccountActivityItem) Get() *AccountActivityItem {
	return v.value
}

func (v *NullableAccountActivityItem) Set(val *AccountActivityItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountActivityItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountActivityItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountActivityItem(val *AccountActivityItem) *NullableAccountActivityItem {
	return &NullableAccountActivityItem{value: val, isSet: true}
}

func (v NullableAccountActivityItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountActivityItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


