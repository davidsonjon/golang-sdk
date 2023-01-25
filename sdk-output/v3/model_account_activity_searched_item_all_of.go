/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// AccountActivitySearchedItemAllOf struct for AccountActivitySearchedItemAllOf
type AccountActivitySearchedItemAllOf struct {
	// The type of action that this activity performed
	Action *string `json:"action,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// The current stage of the activity
	Stage *string `json:"stage,omitempty"`
	Origin NullableString `json:"origin,omitempty"`
	// the current status of the activity
	Status *string `json:"status,omitempty"`
	Requester *Source1 `json:"requester,omitempty"`
	Recipient *Source1 `json:"recipient,omitempty"`
	TrackingNumber *string `json:"trackingNumber,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Approvals []Approval `json:"approvals,omitempty"`
	OriginalRequests []OriginalRequest `json:"originalRequests,omitempty"`
	ExpansionItems []ExpansionItem `json:"expansionItems,omitempty"`
	AccountRequests []AccountRequest `json:"accountRequests,omitempty"`
	Sources *string `json:"sources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountActivitySearchedItemAllOf AccountActivitySearchedItemAllOf

// NewAccountActivitySearchedItemAllOf instantiates a new AccountActivitySearchedItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountActivitySearchedItemAllOf() *AccountActivitySearchedItemAllOf {
	this := AccountActivitySearchedItemAllOf{}
	return &this
}

// NewAccountActivitySearchedItemAllOfWithDefaults instantiates a new AccountActivitySearchedItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountActivitySearchedItemAllOfWithDefaults() *AccountActivitySearchedItemAllOf {
	this := AccountActivitySearchedItemAllOf{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AccountActivitySearchedItemAllOf) SetAction(v string) {
	o.Action = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItemAllOf) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItemAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *AccountActivitySearchedItemAllOf) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *AccountActivitySearchedItemAllOf) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *AccountActivitySearchedItemAllOf) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItemAllOf) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItemAllOf) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *AccountActivitySearchedItemAllOf) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *AccountActivitySearchedItemAllOf) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *AccountActivitySearchedItemAllOf) UnsetModified() {
	o.Modified.Unset()
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetStage() string {
	if o == nil || isNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetStageOk() (*string, bool) {
	if o == nil || isNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasStage() bool {
	if o != nil && !isNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *AccountActivitySearchedItemAllOf) SetStage(v string) {
	o.Stage = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItemAllOf) GetOrigin() string {
	if o == nil || isNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItemAllOf) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *AccountActivitySearchedItemAllOf) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *AccountActivitySearchedItemAllOf) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *AccountActivitySearchedItemAllOf) UnsetOrigin() {
	o.Origin.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountActivitySearchedItemAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetRequester() Source1 {
	if o == nil || isNil(o.Requester) {
		var ret Source1
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetRequesterOk() (*Source1, bool) {
	if o == nil || isNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasRequester() bool {
	if o != nil && !isNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given Source1 and assigns it to the Requester field.
func (o *AccountActivitySearchedItemAllOf) SetRequester(v Source1) {
	o.Requester = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetRecipient() Source1 {
	if o == nil || isNil(o.Recipient) {
		var ret Source1
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetRecipientOk() (*Source1, bool) {
	if o == nil || isNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasRecipient() bool {
	if o != nil && !isNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given Source1 and assigns it to the Recipient field.
func (o *AccountActivitySearchedItemAllOf) SetRecipient(v Source1) {
	o.Recipient = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetTrackingNumber() string {
	if o == nil || isNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetTrackingNumberOk() (*string, bool) {
	if o == nil || isNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasTrackingNumber() bool {
	if o != nil && !isNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *AccountActivitySearchedItemAllOf) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItemAllOf) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItemAllOf) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasErrors() bool {
	if o != nil && isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AccountActivitySearchedItemAllOf) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountActivitySearchedItemAllOf) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountActivitySearchedItemAllOf) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasWarnings() bool {
	if o != nil && isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *AccountActivitySearchedItemAllOf) SetWarnings(v []string) {
	o.Warnings = v
}

// GetApprovals returns the Approvals field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetApprovals() []Approval {
	if o == nil || isNil(o.Approvals) {
		var ret []Approval
		return ret
	}
	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetApprovalsOk() ([]Approval, bool) {
	if o == nil || isNil(o.Approvals) {
		return nil, false
	}
	return o.Approvals, true
}

// HasApprovals returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasApprovals() bool {
	if o != nil && !isNil(o.Approvals) {
		return true
	}

	return false
}

// SetApprovals gets a reference to the given []Approval and assigns it to the Approvals field.
func (o *AccountActivitySearchedItemAllOf) SetApprovals(v []Approval) {
	o.Approvals = v
}

// GetOriginalRequests returns the OriginalRequests field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetOriginalRequests() []OriginalRequest {
	if o == nil || isNil(o.OriginalRequests) {
		var ret []OriginalRequest
		return ret
	}
	return o.OriginalRequests
}

// GetOriginalRequestsOk returns a tuple with the OriginalRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetOriginalRequestsOk() ([]OriginalRequest, bool) {
	if o == nil || isNil(o.OriginalRequests) {
		return nil, false
	}
	return o.OriginalRequests, true
}

// HasOriginalRequests returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasOriginalRequests() bool {
	if o != nil && !isNil(o.OriginalRequests) {
		return true
	}

	return false
}

// SetOriginalRequests gets a reference to the given []OriginalRequest and assigns it to the OriginalRequests field.
func (o *AccountActivitySearchedItemAllOf) SetOriginalRequests(v []OriginalRequest) {
	o.OriginalRequests = v
}

// GetExpansionItems returns the ExpansionItems field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetExpansionItems() []ExpansionItem {
	if o == nil || isNil(o.ExpansionItems) {
		var ret []ExpansionItem
		return ret
	}
	return o.ExpansionItems
}

// GetExpansionItemsOk returns a tuple with the ExpansionItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetExpansionItemsOk() ([]ExpansionItem, bool) {
	if o == nil || isNil(o.ExpansionItems) {
		return nil, false
	}
	return o.ExpansionItems, true
}

// HasExpansionItems returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasExpansionItems() bool {
	if o != nil && !isNil(o.ExpansionItems) {
		return true
	}

	return false
}

// SetExpansionItems gets a reference to the given []ExpansionItem and assigns it to the ExpansionItems field.
func (o *AccountActivitySearchedItemAllOf) SetExpansionItems(v []ExpansionItem) {
	o.ExpansionItems = v
}

// GetAccountRequests returns the AccountRequests field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetAccountRequests() []AccountRequest {
	if o == nil || isNil(o.AccountRequests) {
		var ret []AccountRequest
		return ret
	}
	return o.AccountRequests
}

// GetAccountRequestsOk returns a tuple with the AccountRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetAccountRequestsOk() ([]AccountRequest, bool) {
	if o == nil || isNil(o.AccountRequests) {
		return nil, false
	}
	return o.AccountRequests, true
}

// HasAccountRequests returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasAccountRequests() bool {
	if o != nil && !isNil(o.AccountRequests) {
		return true
	}

	return false
}

// SetAccountRequests gets a reference to the given []AccountRequest and assigns it to the AccountRequests field.
func (o *AccountActivitySearchedItemAllOf) SetAccountRequests(v []AccountRequest) {
	o.AccountRequests = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *AccountActivitySearchedItemAllOf) GetSources() string {
	if o == nil || isNil(o.Sources) {
		var ret string
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountActivitySearchedItemAllOf) GetSourcesOk() (*string, bool) {
	if o == nil || isNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *AccountActivitySearchedItemAllOf) HasSources() bool {
	if o != nil && !isNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given string and assigns it to the Sources field.
func (o *AccountActivitySearchedItemAllOf) SetSources(v string) {
	o.Sources = &v
}

func (o AccountActivitySearchedItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !isNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !isNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !isNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if !isNil(o.Approvals) {
		toSerialize["approvals"] = o.Approvals
	}
	if !isNil(o.OriginalRequests) {
		toSerialize["originalRequests"] = o.OriginalRequests
	}
	if !isNil(o.ExpansionItems) {
		toSerialize["expansionItems"] = o.ExpansionItems
	}
	if !isNil(o.AccountRequests) {
		toSerialize["accountRequests"] = o.AccountRequests
	}
	if !isNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountActivitySearchedItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccountActivitySearchedItemAllOf := _AccountActivitySearchedItemAllOf{}

	if err = json.Unmarshal(bytes, &varAccountActivitySearchedItemAllOf); err == nil {
		*o = AccountActivitySearchedItemAllOf(varAccountActivitySearchedItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

type NullableAccountActivitySearchedItemAllOf struct {
	value *AccountActivitySearchedItemAllOf
	isSet bool
}

func (v NullableAccountActivitySearchedItemAllOf) Get() *AccountActivitySearchedItemAllOf {
	return v.value
}

func (v *NullableAccountActivitySearchedItemAllOf) Set(val *AccountActivitySearchedItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountActivitySearchedItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountActivitySearchedItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountActivitySearchedItemAllOf(val *AccountActivitySearchedItemAllOf) *NullableAccountActivitySearchedItemAllOf {
	return &NullableAccountActivitySearchedItemAllOf{value: val, isSet: true}
}

func (v NullableAccountActivitySearchedItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountActivitySearchedItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


