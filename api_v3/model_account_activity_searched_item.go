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
	"fmt"
)

// checks if the AccountActivitySearchedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountActivitySearchedItem{}

// AccountActivitySearchedItem AccountActivity
type AccountActivitySearchedItem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type DocumentType `json:"_type"`
	// Type of action performed in the activity.
	Action *string `json:"action,omitempty"`
	// ISO-8601 date-time referring to the time when the object was created.
	Created NullableTime `json:"created,omitempty"`
	// ISO-8601 date-time referring to the time when the object was last modified.
	Modified NullableTime `json:"modified,omitempty"`
	// Activity's current stage.
	Stage *string `json:"stage,omitempty"`
	// Activity's origin.
	Origin NullableString `json:"origin,omitempty"`
	// Activity's current status.
	Status *string `json:"status,omitempty"`
	Requester *AccountSource `json:"requester,omitempty"`
	Recipient *AccountSource `json:"recipient,omitempty"`
	// Account activity's tracking number.
	TrackingNumber *string `json:"trackingNumber,omitempty"`
	// Errors provided by the source while completing account actions.
	Errors []string `json:"errors,omitempty"`
	// Warnings provided by the source while completing account actions.
	Warnings []string `json:"warnings,omitempty"`
	// Approvals performed on an item during activity.
	Approvals []Approval `json:"approvals,omitempty"`
	// Original actions that triggered all individual source actions related to the account action.
	OriginalRequests []OriginalRequest `json:"originalRequests,omitempty"`
	// Controls that translated the attribute requests into actual provisioning actions on the source.
	ExpansionItems []ExpansionItem `json:"expansionItems,omitempty"`
	// Account data for each individual source action triggered by the original requests.
	AccountRequests []AccountRequest `json:"accountRequests,omitempty"`
	// Sources involved in the account activity.
	Sources *string `json:"sources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountActivitySearchedItem AccountActivitySearchedItem

// NewAccountActivitySearchedItem instantiates a new AccountActivitySearchedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountActivitySearchedItem(id string, name string, type_ DocumentType) *AccountActivitySearchedItem {
	this := AccountActivitySearchedItem{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewAccountActivitySearchedItemWithDefaults instantiates a new AccountActivitySearchedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountActivitySearchedItemWithDefaults() *AccountActivitySearchedItem {
	this := AccountActivitySearchedItem{}
	return &this
}

// GetId returns the Id field value
func (o *AccountActivitySearchedItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountActivitySearchedItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccountActivitySearchedItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountActivitySearchedItem) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AccountActivitySearchedItem) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountActivitySearchedItem) SetType(v DocumentType) {
	o.Type = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AccountActivitySearchedItem) SetAction(v string) {
	o.Action = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItem) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItem) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *AccountActivitySearchedItem) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *AccountActivitySearchedItem) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *AccountActivitySearchedItem) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItem) GetModified() time.Time {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItem) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *AccountActivitySearchedItem) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *AccountActivitySearchedItem) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *AccountActivitySearchedItem) UnsetModified() {
	o.Modified.Unset()
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *AccountActivitySearchedItem) SetStage(v string) {
	o.Stage = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItem) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItem) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *AccountActivitySearchedItem) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *AccountActivitySearchedItem) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *AccountActivitySearchedItem) UnsetOrigin() {
	o.Origin.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountActivitySearchedItem) SetStatus(v string) {
	o.Status = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetRequester() AccountSource {
	if o == nil || IsNil(o.Requester) {
		var ret AccountSource
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetRequesterOk() (*AccountSource, bool) {
	if o == nil || IsNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasRequester() bool {
	if o != nil && !IsNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given AccountSource and assigns it to the Requester field.
func (o *AccountActivitySearchedItem) SetRequester(v AccountSource) {
	o.Requester = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetRecipient() AccountSource {
	if o == nil || IsNil(o.Recipient) {
		var ret AccountSource
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetRecipientOk() (*AccountSource, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given AccountSource and assigns it to the Recipient field.
func (o *AccountActivitySearchedItem) SetRecipient(v AccountSource) {
	o.Recipient = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *AccountActivitySearchedItem) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItem) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItem) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AccountActivitySearchedItem) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItem) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItem) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *AccountActivitySearchedItem) SetWarnings(v []string) {
	o.Warnings = v
}

// GetApprovals returns the Approvals field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetApprovals() []Approval {
	if o == nil || IsNil(o.Approvals) {
		var ret []Approval
		return ret
	}
	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetApprovalsOk() ([]Approval, bool) {
	if o == nil || IsNil(o.Approvals) {
		return nil, false
	}
	return o.Approvals, true
}

// HasApprovals returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasApprovals() bool {
	if o != nil && !IsNil(o.Approvals) {
		return true
	}

	return false
}

// SetApprovals gets a reference to the given []Approval and assigns it to the Approvals field.
func (o *AccountActivitySearchedItem) SetApprovals(v []Approval) {
	o.Approvals = v
}

// GetOriginalRequests returns the OriginalRequests field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetOriginalRequests() []OriginalRequest {
	if o == nil || IsNil(o.OriginalRequests) {
		var ret []OriginalRequest
		return ret
	}
	return o.OriginalRequests
}

// GetOriginalRequestsOk returns a tuple with the OriginalRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetOriginalRequestsOk() ([]OriginalRequest, bool) {
	if o == nil || IsNil(o.OriginalRequests) {
		return nil, false
	}
	return o.OriginalRequests, true
}

// HasOriginalRequests returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasOriginalRequests() bool {
	if o != nil && !IsNil(o.OriginalRequests) {
		return true
	}

	return false
}

// SetOriginalRequests gets a reference to the given []OriginalRequest and assigns it to the OriginalRequests field.
func (o *AccountActivitySearchedItem) SetOriginalRequests(v []OriginalRequest) {
	o.OriginalRequests = v
}

// GetExpansionItems returns the ExpansionItems field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetExpansionItems() []ExpansionItem {
	if o == nil || IsNil(o.ExpansionItems) {
		var ret []ExpansionItem
		return ret
	}
	return o.ExpansionItems
}

// GetExpansionItemsOk returns a tuple with the ExpansionItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetExpansionItemsOk() ([]ExpansionItem, bool) {
	if o == nil || IsNil(o.ExpansionItems) {
		return nil, false
	}
	return o.ExpansionItems, true
}

// HasExpansionItems returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasExpansionItems() bool {
	if o != nil && !IsNil(o.ExpansionItems) {
		return true
	}

	return false
}

// SetExpansionItems gets a reference to the given []ExpansionItem and assigns it to the ExpansionItems field.
func (o *AccountActivitySearchedItem) SetExpansionItems(v []ExpansionItem) {
	o.ExpansionItems = v
}

// GetAccountRequests returns the AccountRequests field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetAccountRequests() []AccountRequest {
	if o == nil || IsNil(o.AccountRequests) {
		var ret []AccountRequest
		return ret
	}
	return o.AccountRequests
}

// GetAccountRequestsOk returns a tuple with the AccountRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetAccountRequestsOk() ([]AccountRequest, bool) {
	if o == nil || IsNil(o.AccountRequests) {
		return nil, false
	}
	return o.AccountRequests, true
}

// HasAccountRequests returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasAccountRequests() bool {
	if o != nil && !IsNil(o.AccountRequests) {
		return true
	}

	return false
}

// SetAccountRequests gets a reference to the given []AccountRequest and assigns it to the AccountRequests field.
func (o *AccountActivitySearchedItem) SetAccountRequests(v []AccountRequest) {
	o.AccountRequests = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *AccountActivitySearchedItem) GetSources() string {
	if o == nil || IsNil(o.Sources) {
		var ret string
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItem) GetSourcesOk() (*string, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *AccountActivitySearchedItem) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given string and assigns it to the Sources field.
func (o *AccountActivitySearchedItem) SetSources(v string) {
	o.Sources = &v
}

func (o AccountActivitySearchedItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountActivitySearchedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["_type"] = o.Type
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Approvals) {
		toSerialize["approvals"] = o.Approvals
	}
	if !IsNil(o.OriginalRequests) {
		toSerialize["originalRequests"] = o.OriginalRequests
	}
	if !IsNil(o.ExpansionItems) {
		toSerialize["expansionItems"] = o.ExpansionItems
	}
	if !IsNil(o.AccountRequests) {
		toSerialize["accountRequests"] = o.AccountRequests
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountActivitySearchedItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountActivitySearchedItem := _AccountActivitySearchedItem{}

	err = json.Unmarshal(data, &varAccountActivitySearchedItem)

	if err != nil {
		return err
	}

	*o = AccountActivitySearchedItem(varAccountActivitySearchedItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "action")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "stage")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "status")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "trackingNumber")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "approvals")
		delete(additionalProperties, "originalRequests")
		delete(additionalProperties, "expansionItems")
		delete(additionalProperties, "accountRequests")
		delete(additionalProperties, "sources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountActivitySearchedItem struct {
	value *AccountActivitySearchedItem
	isSet bool
}

func (v NullableAccountActivitySearchedItem) Get() *AccountActivitySearchedItem {
	return v.value
}

func (v *NullableAccountActivitySearchedItem) Set(val *AccountActivitySearchedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountActivitySearchedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountActivitySearchedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountActivitySearchedItem(val *AccountActivitySearchedItem) *NullableAccountActivitySearchedItem {
	return &NullableAccountActivitySearchedItem{value: val, isSet: true}
}

func (v NullableAccountActivitySearchedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountActivitySearchedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


