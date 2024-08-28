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

// checks if the WorkItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItems{}

// WorkItems struct for WorkItems
type WorkItems struct {
	// ID of the work item
	Id *string `json:"id,omitempty"`
	// ID of the requester
	RequesterId NullableString `json:"requesterId,omitempty"`
	// The displayname of the requester
	RequesterDisplayName NullableString `json:"requesterDisplayName,omitempty"`
	// The ID of the owner
	OwnerId NullableString `json:"ownerId,omitempty"`
	// The name of the owner
	OwnerName *string `json:"ownerName,omitempty"`
	// Time when the work item was created
	Created *time.Time `json:"created,omitempty"`
	// Time when the work item was last updated
	Modified NullableTime `json:"modified,omitempty"`
	// The description of the work item
	Description *string `json:"description,omitempty"`
	State *WorkItemStateManualWorkItems `json:"state,omitempty"`
	Type *WorkItemTypeManualWorkItems `json:"type,omitempty"`
	// A list of remediation items
	RemediationItems []RemediationItemDetails `json:"remediationItems,omitempty"`
	// A list of items that need to be approved
	ApprovalItems []ApprovalItemDetails `json:"approvalItems,omitempty"`
	// The work item name
	Name NullableString `json:"name,omitempty"`
	// The time at which the work item completed
	Completed NullableTime `json:"completed,omitempty"`
	// The number of items in the work item
	NumItems NullableInt32 `json:"numItems,omitempty"`
	Form *WorkItemsForm `json:"form,omitempty"`
	// An array of errors that ocurred during the work item
	Errors []string `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkItems WorkItems

// NewWorkItems instantiates a new WorkItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItems() *WorkItems {
	this := WorkItems{}
	return &this
}

// NewWorkItemsWithDefaults instantiates a new WorkItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemsWithDefaults() *WorkItems {
	this := WorkItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkItems) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkItems) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkItems) SetId(v string) {
	o.Id = &v
}

// GetRequesterId returns the RequesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetRequesterId() string {
	if o == nil || IsNil(o.RequesterId.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterId.Get()
}

// GetRequesterIdOk returns a tuple with the RequesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetRequesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterId.Get(), o.RequesterId.IsSet()
}

// HasRequesterId returns a boolean if a field has been set.
func (o *WorkItems) HasRequesterId() bool {
	if o != nil && o.RequesterId.IsSet() {
		return true
	}

	return false
}

// SetRequesterId gets a reference to the given NullableString and assigns it to the RequesterId field.
func (o *WorkItems) SetRequesterId(v string) {
	o.RequesterId.Set(&v)
}
// SetRequesterIdNil sets the value for RequesterId to be an explicit nil
func (o *WorkItems) SetRequesterIdNil() {
	o.RequesterId.Set(nil)
}

// UnsetRequesterId ensures that no value is present for RequesterId, not even an explicit nil
func (o *WorkItems) UnsetRequesterId() {
	o.RequesterId.Unset()
}

// GetRequesterDisplayName returns the RequesterDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetRequesterDisplayName() string {
	if o == nil || IsNil(o.RequesterDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterDisplayName.Get()
}

// GetRequesterDisplayNameOk returns a tuple with the RequesterDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetRequesterDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterDisplayName.Get(), o.RequesterDisplayName.IsSet()
}

// HasRequesterDisplayName returns a boolean if a field has been set.
func (o *WorkItems) HasRequesterDisplayName() bool {
	if o != nil && o.RequesterDisplayName.IsSet() {
		return true
	}

	return false
}

// SetRequesterDisplayName gets a reference to the given NullableString and assigns it to the RequesterDisplayName field.
func (o *WorkItems) SetRequesterDisplayName(v string) {
	o.RequesterDisplayName.Set(&v)
}
// SetRequesterDisplayNameNil sets the value for RequesterDisplayName to be an explicit nil
func (o *WorkItems) SetRequesterDisplayNameNil() {
	o.RequesterDisplayName.Set(nil)
}

// UnsetRequesterDisplayName ensures that no value is present for RequesterDisplayName, not even an explicit nil
func (o *WorkItems) UnsetRequesterDisplayName() {
	o.RequesterDisplayName.Unset()
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *WorkItems) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *WorkItems) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *WorkItems) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *WorkItems) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *WorkItems) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *WorkItems) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *WorkItems) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WorkItems) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WorkItems) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *WorkItems) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetModified() time.Time {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *WorkItems) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *WorkItems) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *WorkItems) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *WorkItems) UnsetModified() {
	o.Modified.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkItems) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkItems) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkItems) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WorkItems) GetState() WorkItemStateManualWorkItems {
	if o == nil || IsNil(o.State) {
		var ret WorkItemStateManualWorkItems
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetStateOk() (*WorkItemStateManualWorkItems, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WorkItems) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given WorkItemStateManualWorkItems and assigns it to the State field.
func (o *WorkItems) SetState(v WorkItemStateManualWorkItems) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkItems) GetType() WorkItemTypeManualWorkItems {
	if o == nil || IsNil(o.Type) {
		var ret WorkItemTypeManualWorkItems
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetTypeOk() (*WorkItemTypeManualWorkItems, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkItems) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WorkItemTypeManualWorkItems and assigns it to the Type field.
func (o *WorkItems) SetType(v WorkItemTypeManualWorkItems) {
	o.Type = &v
}

// GetRemediationItems returns the RemediationItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetRemediationItems() []RemediationItemDetails {
	if o == nil {
		var ret []RemediationItemDetails
		return ret
	}
	return o.RemediationItems
}

// GetRemediationItemsOk returns a tuple with the RemediationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetRemediationItemsOk() ([]RemediationItemDetails, bool) {
	if o == nil || IsNil(o.RemediationItems) {
		return nil, false
	}
	return o.RemediationItems, true
}

// HasRemediationItems returns a boolean if a field has been set.
func (o *WorkItems) HasRemediationItems() bool {
	if o != nil && !IsNil(o.RemediationItems) {
		return true
	}

	return false
}

// SetRemediationItems gets a reference to the given []RemediationItemDetails and assigns it to the RemediationItems field.
func (o *WorkItems) SetRemediationItems(v []RemediationItemDetails) {
	o.RemediationItems = v
}

// GetApprovalItems returns the ApprovalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetApprovalItems() []ApprovalItemDetails {
	if o == nil {
		var ret []ApprovalItemDetails
		return ret
	}
	return o.ApprovalItems
}

// GetApprovalItemsOk returns a tuple with the ApprovalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetApprovalItemsOk() ([]ApprovalItemDetails, bool) {
	if o == nil || IsNil(o.ApprovalItems) {
		return nil, false
	}
	return o.ApprovalItems, true
}

// HasApprovalItems returns a boolean if a field has been set.
func (o *WorkItems) HasApprovalItems() bool {
	if o != nil && !IsNil(o.ApprovalItems) {
		return true
	}

	return false
}

// SetApprovalItems gets a reference to the given []ApprovalItemDetails and assigns it to the ApprovalItems field.
func (o *WorkItems) SetApprovalItems(v []ApprovalItemDetails) {
	o.ApprovalItems = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkItems) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkItems) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkItems) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkItems) UnsetName() {
	o.Name.Unset()
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *WorkItems) HasCompleted() bool {
	if o != nil && o.Completed.IsSet() {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given NullableTime and assigns it to the Completed field.
func (o *WorkItems) SetCompleted(v time.Time) {
	o.Completed.Set(&v)
}
// SetCompletedNil sets the value for Completed to be an explicit nil
func (o *WorkItems) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
func (o *WorkItems) UnsetCompleted() {
	o.Completed.Unset()
}

// GetNumItems returns the NumItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItems) GetNumItems() int32 {
	if o == nil || IsNil(o.NumItems.Get()) {
		var ret int32
		return ret
	}
	return *o.NumItems.Get()
}

// GetNumItemsOk returns a tuple with the NumItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItems) GetNumItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumItems.Get(), o.NumItems.IsSet()
}

// HasNumItems returns a boolean if a field has been set.
func (o *WorkItems) HasNumItems() bool {
	if o != nil && o.NumItems.IsSet() {
		return true
	}

	return false
}

// SetNumItems gets a reference to the given NullableInt32 and assigns it to the NumItems field.
func (o *WorkItems) SetNumItems(v int32) {
	o.NumItems.Set(&v)
}
// SetNumItemsNil sets the value for NumItems to be an explicit nil
func (o *WorkItems) SetNumItemsNil() {
	o.NumItems.Set(nil)
}

// UnsetNumItems ensures that no value is present for NumItems, not even an explicit nil
func (o *WorkItems) UnsetNumItems() {
	o.NumItems.Unset()
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *WorkItems) GetForm() WorkItemsForm {
	if o == nil || IsNil(o.Form) {
		var ret WorkItemsForm
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetFormOk() (*WorkItemsForm, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *WorkItems) HasForm() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given WorkItemsForm and assigns it to the Form field.
func (o *WorkItems) SetForm(v WorkItemsForm) {
	o.Form = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *WorkItems) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *WorkItems) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *WorkItems) SetErrors(v []string) {
	o.Errors = v
}

func (o WorkItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.RequesterId.IsSet() {
		toSerialize["requesterId"] = o.RequesterId.Get()
	}
	if o.RequesterDisplayName.IsSet() {
		toSerialize["requesterDisplayName"] = o.RequesterDisplayName.Get()
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.RemediationItems != nil {
		toSerialize["remediationItems"] = o.RemediationItems
	}
	if o.ApprovalItems != nil {
		toSerialize["approvalItems"] = o.ApprovalItems
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	if o.NumItems.IsSet() {
		toSerialize["numItems"] = o.NumItems.Get()
	}
	if !IsNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkItems) UnmarshalJSON(data []byte) (err error) {
	varWorkItems := _WorkItems{}

	err = json.Unmarshal(data, &varWorkItems)

	if err != nil {
		return err
	}

	*o = WorkItems(varWorkItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "requesterId")
		delete(additionalProperties, "requesterDisplayName")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "ownerName")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "type")
		delete(additionalProperties, "remediationItems")
		delete(additionalProperties, "approvalItems")
		delete(additionalProperties, "name")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "numItems")
		delete(additionalProperties, "form")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkItems struct {
	value *WorkItems
	isSet bool
}

func (v NullableWorkItems) Get() *WorkItems {
	return v.value
}

func (v *NullableWorkItems) Set(val *WorkItems) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItems) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItems(val *WorkItems) *NullableWorkItems {
	return &NullableWorkItems{value: val, isSet: true}
}

func (v NullableWorkItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


