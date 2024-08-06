/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the RequestedItemStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestedItemStatus{}

// RequestedItemStatus struct for RequestedItemStatus
type RequestedItemStatus struct {
	// Human-readable display name of the item being requested.
	Name NullableString `json:"name,omitempty"`
	// Type of requested object.
	Type NullableString `json:"type,omitempty"`
	CancelledRequestDetails *RequestedItemStatusCancelledRequestDetails `json:"cancelledRequestDetails,omitempty"`
	// List of list of localized error messages, if any, encountered during the approval/provisioning process.
	ErrorMessages [][]ErrorMessageDto `json:"errorMessages,omitempty"`
	State *RequestedItemStatusRequestState `json:"state,omitempty"`
	// Approval details for each item.
	ApprovalDetails []ApprovalStatusDto `json:"approvalDetails,omitempty"`
	// Manual work items created for provisioning the item.
	ManualWorkItemDetails []ManualWorkItemDetails `json:"manualWorkItemDetails,omitempty"`
	// Id of associated account activity item.
	AccountActivityItemId *string `json:"accountActivityItemId,omitempty"`
	RequestType NullableAccessRequestType `json:"requestType,omitempty"`
	// When the request was last modified.
	Modified NullableTime `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	Requester *AccessItemRequester `json:"requester,omitempty"`
	RequestedFor *RequestedItemStatusRequestedFor `json:"requestedFor,omitempty"`
	RequesterComment *RequestedItemStatusRequesterComment `json:"requesterComment,omitempty"`
	SodViolationContext *RequestedItemStatusSodViolationContext `json:"sodViolationContext,omitempty"`
	ProvisioningDetails *RequestedItemStatusProvisioningDetails `json:"provisioningDetails,omitempty"`
	PreApprovalTriggerDetails *RequestedItemStatusPreApprovalTriggerDetails `json:"preApprovalTriggerDetails,omitempty"`
	// A list of Phases that the Access Request has gone through in order, to help determine the status of the request.
	AccessRequestPhases []AccessRequestPhases `json:"accessRequestPhases,omitempty"`
	// Description associated to the requested object.
	Description NullableString `json:"description,omitempty"`
	// When the role access is scheduled for removal.
	RemoveDate NullableTime `json:"removeDate,omitempty"`
	// True if the request can be canceled.
	Cancelable *bool `json:"cancelable,omitempty"`
	// This is the account activity id.
	AccessRequestId *string `json:"accessRequestId,omitempty"`
	// Arbitrary key-value pairs, if any were included in the corresponding access request
	ClientMetadata map[string]string `json:"clientMetadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestedItemStatus RequestedItemStatus

// NewRequestedItemStatus instantiates a new RequestedItemStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedItemStatus() *RequestedItemStatus {
	this := RequestedItemStatus{}
	var cancelable bool = false
	this.Cancelable = &cancelable
	return &this
}

// NewRequestedItemStatusWithDefaults instantiates a new RequestedItemStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedItemStatusWithDefaults() *RequestedItemStatus {
	this := RequestedItemStatus{}
	var cancelable bool = false
	this.Cancelable = &cancelable
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RequestedItemStatus) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RequestedItemStatus) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RequestedItemStatus) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetType() string {
	if o == nil || isNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *RequestedItemStatus) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *RequestedItemStatus) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *RequestedItemStatus) UnsetType() {
	o.Type.Unset()
}

// GetCancelledRequestDetails returns the CancelledRequestDetails field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetCancelledRequestDetails() RequestedItemStatusCancelledRequestDetails {
	if o == nil || isNil(o.CancelledRequestDetails) {
		var ret RequestedItemStatusCancelledRequestDetails
		return ret
	}
	return *o.CancelledRequestDetails
}

// GetCancelledRequestDetailsOk returns a tuple with the CancelledRequestDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetCancelledRequestDetailsOk() (*RequestedItemStatusCancelledRequestDetails, bool) {
	if o == nil || isNil(o.CancelledRequestDetails) {
		return nil, false
	}
	return o.CancelledRequestDetails, true
}

// HasCancelledRequestDetails returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasCancelledRequestDetails() bool {
	if o != nil && !isNil(o.CancelledRequestDetails) {
		return true
	}

	return false
}

// SetCancelledRequestDetails gets a reference to the given RequestedItemStatusCancelledRequestDetails and assigns it to the CancelledRequestDetails field.
func (o *RequestedItemStatus) SetCancelledRequestDetails(v RequestedItemStatusCancelledRequestDetails) {
	o.CancelledRequestDetails = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetErrorMessages() [][]ErrorMessageDto {
	if o == nil {
		var ret [][]ErrorMessageDto
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetErrorMessagesOk() ([][]ErrorMessageDto, bool) {
	if o == nil || isNil(o.ErrorMessages) {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasErrorMessages() bool {
	if o != nil && isNil(o.ErrorMessages) {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given [][]ErrorMessageDto and assigns it to the ErrorMessages field.
func (o *RequestedItemStatus) SetErrorMessages(v [][]ErrorMessageDto) {
	o.ErrorMessages = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetState() RequestedItemStatusRequestState {
	if o == nil || isNil(o.State) {
		var ret RequestedItemStatusRequestState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetStateOk() (*RequestedItemStatusRequestState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given RequestedItemStatusRequestState and assigns it to the State field.
func (o *RequestedItemStatus) SetState(v RequestedItemStatusRequestState) {
	o.State = &v
}

// GetApprovalDetails returns the ApprovalDetails field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetApprovalDetails() []ApprovalStatusDto {
	if o == nil || isNil(o.ApprovalDetails) {
		var ret []ApprovalStatusDto
		return ret
	}
	return o.ApprovalDetails
}

// GetApprovalDetailsOk returns a tuple with the ApprovalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetApprovalDetailsOk() ([]ApprovalStatusDto, bool) {
	if o == nil || isNil(o.ApprovalDetails) {
		return nil, false
	}
	return o.ApprovalDetails, true
}

// HasApprovalDetails returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasApprovalDetails() bool {
	if o != nil && !isNil(o.ApprovalDetails) {
		return true
	}

	return false
}

// SetApprovalDetails gets a reference to the given []ApprovalStatusDto and assigns it to the ApprovalDetails field.
func (o *RequestedItemStatus) SetApprovalDetails(v []ApprovalStatusDto) {
	o.ApprovalDetails = v
}

// GetManualWorkItemDetails returns the ManualWorkItemDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetManualWorkItemDetails() []ManualWorkItemDetails {
	if o == nil {
		var ret []ManualWorkItemDetails
		return ret
	}
	return o.ManualWorkItemDetails
}

// GetManualWorkItemDetailsOk returns a tuple with the ManualWorkItemDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetManualWorkItemDetailsOk() ([]ManualWorkItemDetails, bool) {
	if o == nil || isNil(o.ManualWorkItemDetails) {
		return nil, false
	}
	return o.ManualWorkItemDetails, true
}

// HasManualWorkItemDetails returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasManualWorkItemDetails() bool {
	if o != nil && isNil(o.ManualWorkItemDetails) {
		return true
	}

	return false
}

// SetManualWorkItemDetails gets a reference to the given []ManualWorkItemDetails and assigns it to the ManualWorkItemDetails field.
func (o *RequestedItemStatus) SetManualWorkItemDetails(v []ManualWorkItemDetails) {
	o.ManualWorkItemDetails = v
}

// GetAccountActivityItemId returns the AccountActivityItemId field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetAccountActivityItemId() string {
	if o == nil || isNil(o.AccountActivityItemId) {
		var ret string
		return ret
	}
	return *o.AccountActivityItemId
}

// GetAccountActivityItemIdOk returns a tuple with the AccountActivityItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetAccountActivityItemIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountActivityItemId) {
		return nil, false
	}
	return o.AccountActivityItemId, true
}

// HasAccountActivityItemId returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasAccountActivityItemId() bool {
	if o != nil && !isNil(o.AccountActivityItemId) {
		return true
	}

	return false
}

// SetAccountActivityItemId gets a reference to the given string and assigns it to the AccountActivityItemId field.
func (o *RequestedItemStatus) SetAccountActivityItemId(v string) {
	o.AccountActivityItemId = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetRequestType() AccessRequestType {
	if o == nil || isNil(o.RequestType.Get()) {
		var ret AccessRequestType
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetRequestTypeOk() (*AccessRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableAccessRequestType and assigns it to the RequestType field.
func (o *RequestedItemStatus) SetRequestType(v AccessRequestType) {
	o.RequestType.Set(&v)
}
// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *RequestedItemStatus) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *RequestedItemStatus) UnsetRequestType() {
	o.RequestType.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *RequestedItemStatus) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *RequestedItemStatus) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *RequestedItemStatus) UnsetModified() {
	o.Modified.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RequestedItemStatus) SetCreated(v time.Time) {
	o.Created = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetRequester() AccessItemRequester {
	if o == nil || isNil(o.Requester) {
		var ret AccessItemRequester
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetRequesterOk() (*AccessItemRequester, bool) {
	if o == nil || isNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasRequester() bool {
	if o != nil && !isNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given AccessItemRequester and assigns it to the Requester field.
func (o *RequestedItemStatus) SetRequester(v AccessItemRequester) {
	o.Requester = &v
}

// GetRequestedFor returns the RequestedFor field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetRequestedFor() RequestedItemStatusRequestedFor {
	if o == nil || isNil(o.RequestedFor) {
		var ret RequestedItemStatusRequestedFor
		return ret
	}
	return *o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetRequestedForOk() (*RequestedItemStatusRequestedFor, bool) {
	if o == nil || isNil(o.RequestedFor) {
		return nil, false
	}
	return o.RequestedFor, true
}

// HasRequestedFor returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasRequestedFor() bool {
	if o != nil && !isNil(o.RequestedFor) {
		return true
	}

	return false
}

// SetRequestedFor gets a reference to the given RequestedItemStatusRequestedFor and assigns it to the RequestedFor field.
func (o *RequestedItemStatus) SetRequestedFor(v RequestedItemStatusRequestedFor) {
	o.RequestedFor = &v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetRequesterComment() RequestedItemStatusRequesterComment {
	if o == nil || isNil(o.RequesterComment) {
		var ret RequestedItemStatusRequesterComment
		return ret
	}
	return *o.RequesterComment
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetRequesterCommentOk() (*RequestedItemStatusRequesterComment, bool) {
	if o == nil || isNil(o.RequesterComment) {
		return nil, false
	}
	return o.RequesterComment, true
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasRequesterComment() bool {
	if o != nil && !isNil(o.RequesterComment) {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given RequestedItemStatusRequesterComment and assigns it to the RequesterComment field.
func (o *RequestedItemStatus) SetRequesterComment(v RequestedItemStatusRequesterComment) {
	o.RequesterComment = &v
}

// GetSodViolationContext returns the SodViolationContext field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetSodViolationContext() RequestedItemStatusSodViolationContext {
	if o == nil || isNil(o.SodViolationContext) {
		var ret RequestedItemStatusSodViolationContext
		return ret
	}
	return *o.SodViolationContext
}

// GetSodViolationContextOk returns a tuple with the SodViolationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetSodViolationContextOk() (*RequestedItemStatusSodViolationContext, bool) {
	if o == nil || isNil(o.SodViolationContext) {
		return nil, false
	}
	return o.SodViolationContext, true
}

// HasSodViolationContext returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasSodViolationContext() bool {
	if o != nil && !isNil(o.SodViolationContext) {
		return true
	}

	return false
}

// SetSodViolationContext gets a reference to the given RequestedItemStatusSodViolationContext and assigns it to the SodViolationContext field.
func (o *RequestedItemStatus) SetSodViolationContext(v RequestedItemStatusSodViolationContext) {
	o.SodViolationContext = &v
}

// GetProvisioningDetails returns the ProvisioningDetails field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetProvisioningDetails() RequestedItemStatusProvisioningDetails {
	if o == nil || isNil(o.ProvisioningDetails) {
		var ret RequestedItemStatusProvisioningDetails
		return ret
	}
	return *o.ProvisioningDetails
}

// GetProvisioningDetailsOk returns a tuple with the ProvisioningDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetProvisioningDetailsOk() (*RequestedItemStatusProvisioningDetails, bool) {
	if o == nil || isNil(o.ProvisioningDetails) {
		return nil, false
	}
	return o.ProvisioningDetails, true
}

// HasProvisioningDetails returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasProvisioningDetails() bool {
	if o != nil && !isNil(o.ProvisioningDetails) {
		return true
	}

	return false
}

// SetProvisioningDetails gets a reference to the given RequestedItemStatusProvisioningDetails and assigns it to the ProvisioningDetails field.
func (o *RequestedItemStatus) SetProvisioningDetails(v RequestedItemStatusProvisioningDetails) {
	o.ProvisioningDetails = &v
}

// GetPreApprovalTriggerDetails returns the PreApprovalTriggerDetails field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetPreApprovalTriggerDetails() RequestedItemStatusPreApprovalTriggerDetails {
	if o == nil || isNil(o.PreApprovalTriggerDetails) {
		var ret RequestedItemStatusPreApprovalTriggerDetails
		return ret
	}
	return *o.PreApprovalTriggerDetails
}

// GetPreApprovalTriggerDetailsOk returns a tuple with the PreApprovalTriggerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetPreApprovalTriggerDetailsOk() (*RequestedItemStatusPreApprovalTriggerDetails, bool) {
	if o == nil || isNil(o.PreApprovalTriggerDetails) {
		return nil, false
	}
	return o.PreApprovalTriggerDetails, true
}

// HasPreApprovalTriggerDetails returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasPreApprovalTriggerDetails() bool {
	if o != nil && !isNil(o.PreApprovalTriggerDetails) {
		return true
	}

	return false
}

// SetPreApprovalTriggerDetails gets a reference to the given RequestedItemStatusPreApprovalTriggerDetails and assigns it to the PreApprovalTriggerDetails field.
func (o *RequestedItemStatus) SetPreApprovalTriggerDetails(v RequestedItemStatusPreApprovalTriggerDetails) {
	o.PreApprovalTriggerDetails = &v
}

// GetAccessRequestPhases returns the AccessRequestPhases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetAccessRequestPhases() []AccessRequestPhases {
	if o == nil {
		var ret []AccessRequestPhases
		return ret
	}
	return o.AccessRequestPhases
}

// GetAccessRequestPhasesOk returns a tuple with the AccessRequestPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetAccessRequestPhasesOk() ([]AccessRequestPhases, bool) {
	if o == nil || isNil(o.AccessRequestPhases) {
		return nil, false
	}
	return o.AccessRequestPhases, true
}

// HasAccessRequestPhases returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasAccessRequestPhases() bool {
	if o != nil && isNil(o.AccessRequestPhases) {
		return true
	}

	return false
}

// SetAccessRequestPhases gets a reference to the given []AccessRequestPhases and assigns it to the AccessRequestPhases field.
func (o *RequestedItemStatus) SetAccessRequestPhases(v []AccessRequestPhases) {
	o.AccessRequestPhases = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RequestedItemStatus) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RequestedItemStatus) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RequestedItemStatus) UnsetDescription() {
	o.Description.Unset()
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetRemoveDate() time.Time {
	if o == nil || isNil(o.RemoveDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate.Get()
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveDate.Get(), o.RemoveDate.IsSet()
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasRemoveDate() bool {
	if o != nil && o.RemoveDate.IsSet() {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given NullableTime and assigns it to the RemoveDate field.
func (o *RequestedItemStatus) SetRemoveDate(v time.Time) {
	o.RemoveDate.Set(&v)
}
// SetRemoveDateNil sets the value for RemoveDate to be an explicit nil
func (o *RequestedItemStatus) SetRemoveDateNil() {
	o.RemoveDate.Set(nil)
}

// UnsetRemoveDate ensures that no value is present for RemoveDate, not even an explicit nil
func (o *RequestedItemStatus) UnsetRemoveDate() {
	o.RemoveDate.Unset()
}

// GetCancelable returns the Cancelable field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetCancelable() bool {
	if o == nil || isNil(o.Cancelable) {
		var ret bool
		return ret
	}
	return *o.Cancelable
}

// GetCancelableOk returns a tuple with the Cancelable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetCancelableOk() (*bool, bool) {
	if o == nil || isNil(o.Cancelable) {
		return nil, false
	}
	return o.Cancelable, true
}

// HasCancelable returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasCancelable() bool {
	if o != nil && !isNil(o.Cancelable) {
		return true
	}

	return false
}

// SetCancelable gets a reference to the given bool and assigns it to the Cancelable field.
func (o *RequestedItemStatus) SetCancelable(v bool) {
	o.Cancelable = &v
}

// GetAccessRequestId returns the AccessRequestId field value if set, zero value otherwise.
func (o *RequestedItemStatus) GetAccessRequestId() string {
	if o == nil || isNil(o.AccessRequestId) {
		var ret string
		return ret
	}
	return *o.AccessRequestId
}

// GetAccessRequestIdOk returns a tuple with the AccessRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatus) GetAccessRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.AccessRequestId) {
		return nil, false
	}
	return o.AccessRequestId, true
}

// HasAccessRequestId returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasAccessRequestId() bool {
	if o != nil && !isNil(o.AccessRequestId) {
		return true
	}

	return false
}

// SetAccessRequestId gets a reference to the given string and assigns it to the AccessRequestId field.
func (o *RequestedItemStatus) SetAccessRequestId(v string) {
	o.AccessRequestId = &v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatus) GetClientMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatus) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return &o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *RequestedItemStatus) HasClientMetadata() bool {
	if o != nil && isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *RequestedItemStatus) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = v
}

func (o RequestedItemStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestedItemStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !isNil(o.CancelledRequestDetails) {
		toSerialize["cancelledRequestDetails"] = o.CancelledRequestDetails
	}
	if o.ErrorMessages != nil {
		toSerialize["errorMessages"] = o.ErrorMessages
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.ApprovalDetails) {
		toSerialize["approvalDetails"] = o.ApprovalDetails
	}
	if o.ManualWorkItemDetails != nil {
		toSerialize["manualWorkItemDetails"] = o.ManualWorkItemDetails
	}
	if !isNil(o.AccountActivityItemId) {
		toSerialize["accountActivityItemId"] = o.AccountActivityItemId
	}
	if o.RequestType.IsSet() {
		toSerialize["requestType"] = o.RequestType.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !isNil(o.RequestedFor) {
		toSerialize["requestedFor"] = o.RequestedFor
	}
	if !isNil(o.RequesterComment) {
		toSerialize["requesterComment"] = o.RequesterComment
	}
	if !isNil(o.SodViolationContext) {
		toSerialize["sodViolationContext"] = o.SodViolationContext
	}
	if !isNil(o.ProvisioningDetails) {
		toSerialize["provisioningDetails"] = o.ProvisioningDetails
	}
	if !isNil(o.PreApprovalTriggerDetails) {
		toSerialize["preApprovalTriggerDetails"] = o.PreApprovalTriggerDetails
	}
	if o.AccessRequestPhases != nil {
		toSerialize["accessRequestPhases"] = o.AccessRequestPhases
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RemoveDate.IsSet() {
		toSerialize["removeDate"] = o.RemoveDate.Get()
	}
	if !isNil(o.Cancelable) {
		toSerialize["cancelable"] = o.Cancelable
	}
	if !isNil(o.AccessRequestId) {
		toSerialize["accessRequestId"] = o.AccessRequestId
	}
	if o.ClientMetadata != nil {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestedItemStatus) UnmarshalJSON(bytes []byte) (err error) {
	varRequestedItemStatus := _RequestedItemStatus{}

	if err = json.Unmarshal(bytes, &varRequestedItemStatus); err == nil {
	*o = RequestedItemStatus(varRequestedItemStatus)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "cancelledRequestDetails")
		delete(additionalProperties, "errorMessages")
		delete(additionalProperties, "state")
		delete(additionalProperties, "approvalDetails")
		delete(additionalProperties, "manualWorkItemDetails")
		delete(additionalProperties, "accountActivityItemId")
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "sodViolationContext")
		delete(additionalProperties, "provisioningDetails")
		delete(additionalProperties, "preApprovalTriggerDetails")
		delete(additionalProperties, "accessRequestPhases")
		delete(additionalProperties, "description")
		delete(additionalProperties, "removeDate")
		delete(additionalProperties, "cancelable")
		delete(additionalProperties, "accessRequestId")
		delete(additionalProperties, "clientMetadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestedItemStatus struct {
	value *RequestedItemStatus
	isSet bool
}

func (v NullableRequestedItemStatus) Get() *RequestedItemStatus {
	return v.value
}

func (v *NullableRequestedItemStatus) Set(val *RequestedItemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedItemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedItemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedItemStatus(val *RequestedItemStatus) *NullableRequestedItemStatus {
	return &NullableRequestedItemStatus{value: val, isSet: true}
}

func (v NullableRequestedItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedItemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


