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

// checks if the PendingApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingApproval{}

// PendingApproval struct for PendingApproval
type PendingApproval struct {
	// The approval id.
	Id *string `json:"id,omitempty"`
	// The name of the approval.
	Name *string `json:"name,omitempty"`
	// When the approval was created.
	Created *time.Time `json:"created,omitempty"`
	// When the approval was modified last time.
	Modified *time.Time `json:"modified,omitempty"`
	// When the access-request was created.
	RequestCreated *time.Time `json:"requestCreated,omitempty"`
	RequestType NullableAccessRequestType `json:"requestType,omitempty"`
	Requester *AccessItemRequester `json:"requester,omitempty"`
	// Identities access was requested for.
	RequestedFor []AccessItemRequestedFor `json:"requestedFor,omitempty"`
	Owner *PendingApprovalOwner `json:"owner,omitempty"`
	RequestedObject *RequestableObjectReference `json:"requestedObject,omitempty"`
	RequesterComment *CommentDto `json:"requesterComment,omitempty"`
	// The history of the previous reviewers comments.
	PreviousReviewersComments []CommentDto `json:"previousReviewersComments,omitempty"`
	// The history of approval forward action.
	ForwardHistory []ApprovalForwardHistory `json:"forwardHistory,omitempty"`
	// When true the rejector has to provide comments when rejecting
	CommentRequiredWhenRejected *bool `json:"commentRequiredWhenRejected,omitempty"`
	ActionInProcess *PendingApprovalAction `json:"actionInProcess,omitempty"`
	// The date the role or access profile or entitlement is no longer assigned to the specified identity.
	RemoveDate *time.Time `json:"removeDate,omitempty"`
	// If true, then the request is to change the remove date or sunset date.
	RemoveDateUpdateRequested *bool `json:"removeDateUpdateRequested,omitempty"`
	// The remove date or sunset date that was assigned at the time of the request.
	CurrentRemoveDate *time.Time `json:"currentRemoveDate,omitempty"`
	SodViolationContext *SodViolationContextCheckCompleted `json:"sodViolationContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PendingApproval PendingApproval

// NewPendingApproval instantiates a new PendingApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingApproval() *PendingApproval {
	this := PendingApproval{}
	var commentRequiredWhenRejected bool = false
	this.CommentRequiredWhenRejected = &commentRequiredWhenRejected
	var removeDateUpdateRequested bool = false
	this.RemoveDateUpdateRequested = &removeDateUpdateRequested
	return &this
}

// NewPendingApprovalWithDefaults instantiates a new PendingApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingApprovalWithDefaults() *PendingApproval {
	this := PendingApproval{}
	var commentRequiredWhenRejected bool = false
	this.CommentRequiredWhenRejected = &commentRequiredWhenRejected
	var removeDateUpdateRequested bool = false
	this.RemoveDateUpdateRequested = &removeDateUpdateRequested
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PendingApproval) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PendingApproval) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PendingApproval) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PendingApproval) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PendingApproval) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PendingApproval) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PendingApproval) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PendingApproval) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PendingApproval) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *PendingApproval) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *PendingApproval) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *PendingApproval) SetModified(v time.Time) {
	o.Modified = &v
}

// GetRequestCreated returns the RequestCreated field value if set, zero value otherwise.
func (o *PendingApproval) GetRequestCreated() time.Time {
	if o == nil || isNil(o.RequestCreated) {
		var ret time.Time
		return ret
	}
	return *o.RequestCreated
}

// GetRequestCreatedOk returns a tuple with the RequestCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRequestCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestCreated) {
		return nil, false
	}
	return o.RequestCreated, true
}

// HasRequestCreated returns a boolean if a field has been set.
func (o *PendingApproval) HasRequestCreated() bool {
	if o != nil && !isNil(o.RequestCreated) {
		return true
	}

	return false
}

// SetRequestCreated gets a reference to the given time.Time and assigns it to the RequestCreated field.
func (o *PendingApproval) SetRequestCreated(v time.Time) {
	o.RequestCreated = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingApproval) GetRequestType() AccessRequestType {
	if o == nil || isNil(o.RequestType.Get()) {
		var ret AccessRequestType
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingApproval) GetRequestTypeOk() (*AccessRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *PendingApproval) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableAccessRequestType and assigns it to the RequestType field.
func (o *PendingApproval) SetRequestType(v AccessRequestType) {
	o.RequestType.Set(&v)
}
// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *PendingApproval) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *PendingApproval) UnsetRequestType() {
	o.RequestType.Unset()
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *PendingApproval) GetRequester() AccessItemRequester {
	if o == nil || isNil(o.Requester) {
		var ret AccessItemRequester
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRequesterOk() (*AccessItemRequester, bool) {
	if o == nil || isNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *PendingApproval) HasRequester() bool {
	if o != nil && !isNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given AccessItemRequester and assigns it to the Requester field.
func (o *PendingApproval) SetRequester(v AccessItemRequester) {
	o.Requester = &v
}

// GetRequestedFor returns the RequestedFor field value if set, zero value otherwise.
func (o *PendingApproval) GetRequestedFor() []AccessItemRequestedFor {
	if o == nil || isNil(o.RequestedFor) {
		var ret []AccessItemRequestedFor
		return ret
	}
	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRequestedForOk() ([]AccessItemRequestedFor, bool) {
	if o == nil || isNil(o.RequestedFor) {
		return nil, false
	}
	return o.RequestedFor, true
}

// HasRequestedFor returns a boolean if a field has been set.
func (o *PendingApproval) HasRequestedFor() bool {
	if o != nil && !isNil(o.RequestedFor) {
		return true
	}

	return false
}

// SetRequestedFor gets a reference to the given []AccessItemRequestedFor and assigns it to the RequestedFor field.
func (o *PendingApproval) SetRequestedFor(v []AccessItemRequestedFor) {
	o.RequestedFor = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PendingApproval) GetOwner() PendingApprovalOwner {
	if o == nil || isNil(o.Owner) {
		var ret PendingApprovalOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetOwnerOk() (*PendingApprovalOwner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PendingApproval) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given PendingApprovalOwner and assigns it to the Owner field.
func (o *PendingApproval) SetOwner(v PendingApprovalOwner) {
	o.Owner = &v
}

// GetRequestedObject returns the RequestedObject field value if set, zero value otherwise.
func (o *PendingApproval) GetRequestedObject() RequestableObjectReference {
	if o == nil || isNil(o.RequestedObject) {
		var ret RequestableObjectReference
		return ret
	}
	return *o.RequestedObject
}

// GetRequestedObjectOk returns a tuple with the RequestedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRequestedObjectOk() (*RequestableObjectReference, bool) {
	if o == nil || isNil(o.RequestedObject) {
		return nil, false
	}
	return o.RequestedObject, true
}

// HasRequestedObject returns a boolean if a field has been set.
func (o *PendingApproval) HasRequestedObject() bool {
	if o != nil && !isNil(o.RequestedObject) {
		return true
	}

	return false
}

// SetRequestedObject gets a reference to the given RequestableObjectReference and assigns it to the RequestedObject field.
func (o *PendingApproval) SetRequestedObject(v RequestableObjectReference) {
	o.RequestedObject = &v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise.
func (o *PendingApproval) GetRequesterComment() CommentDto {
	if o == nil || isNil(o.RequesterComment) {
		var ret CommentDto
		return ret
	}
	return *o.RequesterComment
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRequesterCommentOk() (*CommentDto, bool) {
	if o == nil || isNil(o.RequesterComment) {
		return nil, false
	}
	return o.RequesterComment, true
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *PendingApproval) HasRequesterComment() bool {
	if o != nil && !isNil(o.RequesterComment) {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given CommentDto and assigns it to the RequesterComment field.
func (o *PendingApproval) SetRequesterComment(v CommentDto) {
	o.RequesterComment = &v
}

// GetPreviousReviewersComments returns the PreviousReviewersComments field value if set, zero value otherwise.
func (o *PendingApproval) GetPreviousReviewersComments() []CommentDto {
	if o == nil || isNil(o.PreviousReviewersComments) {
		var ret []CommentDto
		return ret
	}
	return o.PreviousReviewersComments
}

// GetPreviousReviewersCommentsOk returns a tuple with the PreviousReviewersComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetPreviousReviewersCommentsOk() ([]CommentDto, bool) {
	if o == nil || isNil(o.PreviousReviewersComments) {
		return nil, false
	}
	return o.PreviousReviewersComments, true
}

// HasPreviousReviewersComments returns a boolean if a field has been set.
func (o *PendingApproval) HasPreviousReviewersComments() bool {
	if o != nil && !isNil(o.PreviousReviewersComments) {
		return true
	}

	return false
}

// SetPreviousReviewersComments gets a reference to the given []CommentDto and assigns it to the PreviousReviewersComments field.
func (o *PendingApproval) SetPreviousReviewersComments(v []CommentDto) {
	o.PreviousReviewersComments = v
}

// GetForwardHistory returns the ForwardHistory field value if set, zero value otherwise.
func (o *PendingApproval) GetForwardHistory() []ApprovalForwardHistory {
	if o == nil || isNil(o.ForwardHistory) {
		var ret []ApprovalForwardHistory
		return ret
	}
	return o.ForwardHistory
}

// GetForwardHistoryOk returns a tuple with the ForwardHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetForwardHistoryOk() ([]ApprovalForwardHistory, bool) {
	if o == nil || isNil(o.ForwardHistory) {
		return nil, false
	}
	return o.ForwardHistory, true
}

// HasForwardHistory returns a boolean if a field has been set.
func (o *PendingApproval) HasForwardHistory() bool {
	if o != nil && !isNil(o.ForwardHistory) {
		return true
	}

	return false
}

// SetForwardHistory gets a reference to the given []ApprovalForwardHistory and assigns it to the ForwardHistory field.
func (o *PendingApproval) SetForwardHistory(v []ApprovalForwardHistory) {
	o.ForwardHistory = v
}

// GetCommentRequiredWhenRejected returns the CommentRequiredWhenRejected field value if set, zero value otherwise.
func (o *PendingApproval) GetCommentRequiredWhenRejected() bool {
	if o == nil || isNil(o.CommentRequiredWhenRejected) {
		var ret bool
		return ret
	}
	return *o.CommentRequiredWhenRejected
}

// GetCommentRequiredWhenRejectedOk returns a tuple with the CommentRequiredWhenRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetCommentRequiredWhenRejectedOk() (*bool, bool) {
	if o == nil || isNil(o.CommentRequiredWhenRejected) {
		return nil, false
	}
	return o.CommentRequiredWhenRejected, true
}

// HasCommentRequiredWhenRejected returns a boolean if a field has been set.
func (o *PendingApproval) HasCommentRequiredWhenRejected() bool {
	if o != nil && !isNil(o.CommentRequiredWhenRejected) {
		return true
	}

	return false
}

// SetCommentRequiredWhenRejected gets a reference to the given bool and assigns it to the CommentRequiredWhenRejected field.
func (o *PendingApproval) SetCommentRequiredWhenRejected(v bool) {
	o.CommentRequiredWhenRejected = &v
}

// GetActionInProcess returns the ActionInProcess field value if set, zero value otherwise.
func (o *PendingApproval) GetActionInProcess() PendingApprovalAction {
	if o == nil || isNil(o.ActionInProcess) {
		var ret PendingApprovalAction
		return ret
	}
	return *o.ActionInProcess
}

// GetActionInProcessOk returns a tuple with the ActionInProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetActionInProcessOk() (*PendingApprovalAction, bool) {
	if o == nil || isNil(o.ActionInProcess) {
		return nil, false
	}
	return o.ActionInProcess, true
}

// HasActionInProcess returns a boolean if a field has been set.
func (o *PendingApproval) HasActionInProcess() bool {
	if o != nil && !isNil(o.ActionInProcess) {
		return true
	}

	return false
}

// SetActionInProcess gets a reference to the given PendingApprovalAction and assigns it to the ActionInProcess field.
func (o *PendingApproval) SetActionInProcess(v PendingApprovalAction) {
	o.ActionInProcess = &v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise.
func (o *PendingApproval) GetRemoveDate() time.Time {
	if o == nil || isNil(o.RemoveDate) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.RemoveDate) {
		return nil, false
	}
	return o.RemoveDate, true
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *PendingApproval) HasRemoveDate() bool {
	if o != nil && !isNil(o.RemoveDate) {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given time.Time and assigns it to the RemoveDate field.
func (o *PendingApproval) SetRemoveDate(v time.Time) {
	o.RemoveDate = &v
}

// GetRemoveDateUpdateRequested returns the RemoveDateUpdateRequested field value if set, zero value otherwise.
func (o *PendingApproval) GetRemoveDateUpdateRequested() bool {
	if o == nil || isNil(o.RemoveDateUpdateRequested) {
		var ret bool
		return ret
	}
	return *o.RemoveDateUpdateRequested
}

// GetRemoveDateUpdateRequestedOk returns a tuple with the RemoveDateUpdateRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetRemoveDateUpdateRequestedOk() (*bool, bool) {
	if o == nil || isNil(o.RemoveDateUpdateRequested) {
		return nil, false
	}
	return o.RemoveDateUpdateRequested, true
}

// HasRemoveDateUpdateRequested returns a boolean if a field has been set.
func (o *PendingApproval) HasRemoveDateUpdateRequested() bool {
	if o != nil && !isNil(o.RemoveDateUpdateRequested) {
		return true
	}

	return false
}

// SetRemoveDateUpdateRequested gets a reference to the given bool and assigns it to the RemoveDateUpdateRequested field.
func (o *PendingApproval) SetRemoveDateUpdateRequested(v bool) {
	o.RemoveDateUpdateRequested = &v
}

// GetCurrentRemoveDate returns the CurrentRemoveDate field value if set, zero value otherwise.
func (o *PendingApproval) GetCurrentRemoveDate() time.Time {
	if o == nil || isNil(o.CurrentRemoveDate) {
		var ret time.Time
		return ret
	}
	return *o.CurrentRemoveDate
}

// GetCurrentRemoveDateOk returns a tuple with the CurrentRemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetCurrentRemoveDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CurrentRemoveDate) {
		return nil, false
	}
	return o.CurrentRemoveDate, true
}

// HasCurrentRemoveDate returns a boolean if a field has been set.
func (o *PendingApproval) HasCurrentRemoveDate() bool {
	if o != nil && !isNil(o.CurrentRemoveDate) {
		return true
	}

	return false
}

// SetCurrentRemoveDate gets a reference to the given time.Time and assigns it to the CurrentRemoveDate field.
func (o *PendingApproval) SetCurrentRemoveDate(v time.Time) {
	o.CurrentRemoveDate = &v
}

// GetSodViolationContext returns the SodViolationContext field value if set, zero value otherwise.
func (o *PendingApproval) GetSodViolationContext() SodViolationContextCheckCompleted {
	if o == nil || isNil(o.SodViolationContext) {
		var ret SodViolationContextCheckCompleted
		return ret
	}
	return *o.SodViolationContext
}

// GetSodViolationContextOk returns a tuple with the SodViolationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingApproval) GetSodViolationContextOk() (*SodViolationContextCheckCompleted, bool) {
	if o == nil || isNil(o.SodViolationContext) {
		return nil, false
	}
	return o.SodViolationContext, true
}

// HasSodViolationContext returns a boolean if a field has been set.
func (o *PendingApproval) HasSodViolationContext() bool {
	if o != nil && !isNil(o.SodViolationContext) {
		return true
	}

	return false
}

// SetSodViolationContext gets a reference to the given SodViolationContextCheckCompleted and assigns it to the SodViolationContext field.
func (o *PendingApproval) SetSodViolationContext(v SodViolationContextCheckCompleted) {
	o.SodViolationContext = &v
}

func (o PendingApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.RequestCreated) {
		toSerialize["requestCreated"] = o.RequestCreated
	}
	if o.RequestType.IsSet() {
		toSerialize["requestType"] = o.RequestType.Get()
	}
	if !isNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !isNil(o.RequestedFor) {
		toSerialize["requestedFor"] = o.RequestedFor
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.RequestedObject) {
		toSerialize["requestedObject"] = o.RequestedObject
	}
	if !isNil(o.RequesterComment) {
		toSerialize["requesterComment"] = o.RequesterComment
	}
	if !isNil(o.PreviousReviewersComments) {
		toSerialize["previousReviewersComments"] = o.PreviousReviewersComments
	}
	if !isNil(o.ForwardHistory) {
		toSerialize["forwardHistory"] = o.ForwardHistory
	}
	if !isNil(o.CommentRequiredWhenRejected) {
		toSerialize["commentRequiredWhenRejected"] = o.CommentRequiredWhenRejected
	}
	if !isNil(o.ActionInProcess) {
		toSerialize["actionInProcess"] = o.ActionInProcess
	}
	if !isNil(o.RemoveDate) {
		toSerialize["removeDate"] = o.RemoveDate
	}
	if !isNil(o.RemoveDateUpdateRequested) {
		toSerialize["removeDateUpdateRequested"] = o.RemoveDateUpdateRequested
	}
	if !isNil(o.CurrentRemoveDate) {
		toSerialize["currentRemoveDate"] = o.CurrentRemoveDate
	}
	if !isNil(o.SodViolationContext) {
		toSerialize["sodViolationContext"] = o.SodViolationContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PendingApproval) UnmarshalJSON(bytes []byte) (err error) {
	varPendingApproval := _PendingApproval{}

	if err = json.Unmarshal(bytes, &varPendingApproval); err == nil {
	*o = PendingApproval(varPendingApproval)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "requestCreated")
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "requestedObject")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "previousReviewersComments")
		delete(additionalProperties, "forwardHistory")
		delete(additionalProperties, "commentRequiredWhenRejected")
		delete(additionalProperties, "actionInProcess")
		delete(additionalProperties, "removeDate")
		delete(additionalProperties, "removeDateUpdateRequested")
		delete(additionalProperties, "currentRemoveDate")
		delete(additionalProperties, "sodViolationContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingApproval struct {
	value *PendingApproval
	isSet bool
}

func (v NullablePendingApproval) Get() *PendingApproval {
	return v.value
}

func (v *NullablePendingApproval) Set(val *PendingApproval) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingApproval) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingApproval(val *PendingApproval) *NullablePendingApproval {
	return &NullablePendingApproval{value: val, isSet: true}
}

func (v NullablePendingApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


