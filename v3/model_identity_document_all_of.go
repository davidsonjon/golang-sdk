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

// checks if the IdentityDocumentAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityDocumentAllOf{}

// IdentityDocumentAllOf struct for IdentityDocumentAllOf
type IdentityDocumentAllOf struct {
	// The first name of the identity
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the identity
	LastName *string `json:"lastName,omitempty"`
	// The display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// The identity's primary email address
	Email *string `json:"email,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	// The phone number of the identity
	Phone *string `json:"phone,omitempty"`
	// Indicates if the identity is inactive
	Inactive *bool `json:"inactive,omitempty"`
	Protected *bool `json:"protected,omitempty"`
	// The identity's status in SailPoint
	Status *string `json:"status,omitempty"`
	EmployeeNumber *string `json:"employeeNumber,omitempty"`
	Manager *DisplayReference `json:"manager,omitempty"`
	// Indicates if this identity is a manager of other identities
	IsManager *bool `json:"isManager,omitempty"`
	IdentityProfile *Reference1 `json:"identityProfile,omitempty"`
	Source *Reference1 `json:"source,omitempty"`
	// a map or dictionary of key/value pairs
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	ProcessingState NullableString `json:"processingState,omitempty"`
	ProcessingDetails *ProcessingDetails `json:"processingDetails,omitempty"`
	// List of accounts associated with the identity
	Accounts []BaseAccount `json:"accounts,omitempty"`
	// Number of accounts associated with the identity
	AccountCount *int32 `json:"accountCount,omitempty"`
	// The list of applications the identity has access to
	Apps []App `json:"apps,omitempty"`
	// The number of applications the identity has access to
	AppCount *int32 `json:"appCount,omitempty"`
	// The list of access items assigned to the identity
	Access []IdentityAccess `json:"access,omitempty"`
	// The number of access items assigned to the identity
	AccessCount *int32 `json:"accessCount,omitempty"`
	// The number of access profiles assigned to the identity
	AccessProfileCount *int32 `json:"accessProfileCount,omitempty"`
	// The number of entitlements assigned to the identity
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// The number of roles assigned to the identity
	RoleCount *int32 `json:"roleCount,omitempty"`
	Owns *Owns `json:"owns,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityDocumentAllOf IdentityDocumentAllOf

// NewIdentityDocumentAllOf instantiates a new IdentityDocumentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityDocumentAllOf() *IdentityDocumentAllOf {
	this := IdentityDocumentAllOf{}
	return &this
}

// NewIdentityDocumentAllOfWithDefaults instantiates a new IdentityDocumentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityDocumentAllOfWithDefaults() *IdentityDocumentAllOf {
	this := IdentityDocumentAllOf{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IdentityDocumentAllOf) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *IdentityDocumentAllOf) SetLastName(v string) {
	o.LastName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityDocumentAllOf) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IdentityDocumentAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityDocumentAllOf) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityDocumentAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *IdentityDocumentAllOf) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *IdentityDocumentAllOf) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *IdentityDocumentAllOf) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityDocumentAllOf) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityDocumentAllOf) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *IdentityDocumentAllOf) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *IdentityDocumentAllOf) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *IdentityDocumentAllOf) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityDocumentAllOf) GetSynced() time.Time {
	if o == nil || isNil(o.Synced.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityDocumentAllOf) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *IdentityDocumentAllOf) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *IdentityDocumentAllOf) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *IdentityDocumentAllOf) UnsetSynced() {
	o.Synced.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *IdentityDocumentAllOf) SetPhone(v string) {
	o.Phone = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetInactive() bool {
	if o == nil || isNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetInactiveOk() (*bool, bool) {
	if o == nil || isNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasInactive() bool {
	if o != nil && !isNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *IdentityDocumentAllOf) SetInactive(v bool) {
	o.Inactive = &v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetProtected() bool {
	if o == nil || isNil(o.Protected) {
		var ret bool
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetProtectedOk() (*bool, bool) {
	if o == nil || isNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasProtected() bool {
	if o != nil && !isNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *IdentityDocumentAllOf) SetProtected(v bool) {
	o.Protected = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentityDocumentAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetEmployeeNumber() string {
	if o == nil || isNil(o.EmployeeNumber) {
		var ret string
		return ret
	}
	return *o.EmployeeNumber
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetEmployeeNumberOk() (*string, bool) {
	if o == nil || isNil(o.EmployeeNumber) {
		return nil, false
	}
	return o.EmployeeNumber, true
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasEmployeeNumber() bool {
	if o != nil && !isNil(o.EmployeeNumber) {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given string and assigns it to the EmployeeNumber field.
func (o *IdentityDocumentAllOf) SetEmployeeNumber(v string) {
	o.EmployeeNumber = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetManager() DisplayReference {
	if o == nil || isNil(o.Manager) {
		var ret DisplayReference
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetManagerOk() (*DisplayReference, bool) {
	if o == nil || isNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasManager() bool {
	if o != nil && !isNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given DisplayReference and assigns it to the Manager field.
func (o *IdentityDocumentAllOf) SetManager(v DisplayReference) {
	o.Manager = &v
}

// GetIsManager returns the IsManager field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetIsManager() bool {
	if o == nil || isNil(o.IsManager) {
		var ret bool
		return ret
	}
	return *o.IsManager
}

// GetIsManagerOk returns a tuple with the IsManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetIsManagerOk() (*bool, bool) {
	if o == nil || isNil(o.IsManager) {
		return nil, false
	}
	return o.IsManager, true
}

// HasIsManager returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasIsManager() bool {
	if o != nil && !isNil(o.IsManager) {
		return true
	}

	return false
}

// SetIsManager gets a reference to the given bool and assigns it to the IsManager field.
func (o *IdentityDocumentAllOf) SetIsManager(v bool) {
	o.IsManager = &v
}

// GetIdentityProfile returns the IdentityProfile field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetIdentityProfile() Reference1 {
	if o == nil || isNil(o.IdentityProfile) {
		var ret Reference1
		return ret
	}
	return *o.IdentityProfile
}

// GetIdentityProfileOk returns a tuple with the IdentityProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetIdentityProfileOk() (*Reference1, bool) {
	if o == nil || isNil(o.IdentityProfile) {
		return nil, false
	}
	return o.IdentityProfile, true
}

// HasIdentityProfile returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasIdentityProfile() bool {
	if o != nil && !isNil(o.IdentityProfile) {
		return true
	}

	return false
}

// SetIdentityProfile gets a reference to the given Reference1 and assigns it to the IdentityProfile field.
func (o *IdentityDocumentAllOf) SetIdentityProfile(v Reference1) {
	o.IdentityProfile = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetSource() Reference1 {
	if o == nil || isNil(o.Source) {
		var ret Reference1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetSourceOk() (*Reference1, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference1 and assigns it to the Source field.
func (o *IdentityDocumentAllOf) SetSource(v Reference1) {
	o.Source = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *IdentityDocumentAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetProcessingState returns the ProcessingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityDocumentAllOf) GetProcessingState() string {
	if o == nil || isNil(o.ProcessingState.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessingState.Get()
}

// GetProcessingStateOk returns a tuple with the ProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityDocumentAllOf) GetProcessingStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessingState.Get(), o.ProcessingState.IsSet()
}

// HasProcessingState returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasProcessingState() bool {
	if o != nil && o.ProcessingState.IsSet() {
		return true
	}

	return false
}

// SetProcessingState gets a reference to the given NullableString and assigns it to the ProcessingState field.
func (o *IdentityDocumentAllOf) SetProcessingState(v string) {
	o.ProcessingState.Set(&v)
}
// SetProcessingStateNil sets the value for ProcessingState to be an explicit nil
func (o *IdentityDocumentAllOf) SetProcessingStateNil() {
	o.ProcessingState.Set(nil)
}

// UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
func (o *IdentityDocumentAllOf) UnsetProcessingState() {
	o.ProcessingState.Unset()
}

// GetProcessingDetails returns the ProcessingDetails field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetProcessingDetails() ProcessingDetails {
	if o == nil || isNil(o.ProcessingDetails) {
		var ret ProcessingDetails
		return ret
	}
	return *o.ProcessingDetails
}

// GetProcessingDetailsOk returns a tuple with the ProcessingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetProcessingDetailsOk() (*ProcessingDetails, bool) {
	if o == nil || isNil(o.ProcessingDetails) {
		return nil, false
	}
	return o.ProcessingDetails, true
}

// HasProcessingDetails returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasProcessingDetails() bool {
	if o != nil && !isNil(o.ProcessingDetails) {
		return true
	}

	return false
}

// SetProcessingDetails gets a reference to the given ProcessingDetails and assigns it to the ProcessingDetails field.
func (o *IdentityDocumentAllOf) SetProcessingDetails(v ProcessingDetails) {
	o.ProcessingDetails = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAccounts() []BaseAccount {
	if o == nil || isNil(o.Accounts) {
		var ret []BaseAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAccountsOk() ([]BaseAccount, bool) {
	if o == nil || isNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAccounts() bool {
	if o != nil && !isNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []BaseAccount and assigns it to the Accounts field.
func (o *IdentityDocumentAllOf) SetAccounts(v []BaseAccount) {
	o.Accounts = v
}

// GetAccountCount returns the AccountCount field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAccountCount() int32 {
	if o == nil || isNil(o.AccountCount) {
		var ret int32
		return ret
	}
	return *o.AccountCount
}

// GetAccountCountOk returns a tuple with the AccountCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAccountCountOk() (*int32, bool) {
	if o == nil || isNil(o.AccountCount) {
		return nil, false
	}
	return o.AccountCount, true
}

// HasAccountCount returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAccountCount() bool {
	if o != nil && !isNil(o.AccountCount) {
		return true
	}

	return false
}

// SetAccountCount gets a reference to the given int32 and assigns it to the AccountCount field.
func (o *IdentityDocumentAllOf) SetAccountCount(v int32) {
	o.AccountCount = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetApps() []App {
	if o == nil || isNil(o.Apps) {
		var ret []App
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAppsOk() ([]App, bool) {
	if o == nil || isNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasApps() bool {
	if o != nil && !isNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []App and assigns it to the Apps field.
func (o *IdentityDocumentAllOf) SetApps(v []App) {
	o.Apps = v
}

// GetAppCount returns the AppCount field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAppCount() int32 {
	if o == nil || isNil(o.AppCount) {
		var ret int32
		return ret
	}
	return *o.AppCount
}

// GetAppCountOk returns a tuple with the AppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAppCountOk() (*int32, bool) {
	if o == nil || isNil(o.AppCount) {
		return nil, false
	}
	return o.AppCount, true
}

// HasAppCount returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAppCount() bool {
	if o != nil && !isNil(o.AppCount) {
		return true
	}

	return false
}

// SetAppCount gets a reference to the given int32 and assigns it to the AppCount field.
func (o *IdentityDocumentAllOf) SetAppCount(v int32) {
	o.AppCount = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAccess() []IdentityAccess {
	if o == nil || isNil(o.Access) {
		var ret []IdentityAccess
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAccessOk() ([]IdentityAccess, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []IdentityAccess and assigns it to the Access field.
func (o *IdentityDocumentAllOf) SetAccess(v []IdentityAccess) {
	o.Access = v
}

// GetAccessCount returns the AccessCount field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAccessCount() int32 {
	if o == nil || isNil(o.AccessCount) {
		var ret int32
		return ret
	}
	return *o.AccessCount
}

// GetAccessCountOk returns a tuple with the AccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAccessCountOk() (*int32, bool) {
	if o == nil || isNil(o.AccessCount) {
		return nil, false
	}
	return o.AccessCount, true
}

// HasAccessCount returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAccessCount() bool {
	if o != nil && !isNil(o.AccessCount) {
		return true
	}

	return false
}

// SetAccessCount gets a reference to the given int32 and assigns it to the AccessCount field.
func (o *IdentityDocumentAllOf) SetAccessCount(v int32) {
	o.AccessCount = &v
}

// GetAccessProfileCount returns the AccessProfileCount field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetAccessProfileCount() int32 {
	if o == nil || isNil(o.AccessProfileCount) {
		var ret int32
		return ret
	}
	return *o.AccessProfileCount
}

// GetAccessProfileCountOk returns a tuple with the AccessProfileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetAccessProfileCountOk() (*int32, bool) {
	if o == nil || isNil(o.AccessProfileCount) {
		return nil, false
	}
	return o.AccessProfileCount, true
}

// HasAccessProfileCount returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasAccessProfileCount() bool {
	if o != nil && !isNil(o.AccessProfileCount) {
		return true
	}

	return false
}

// SetAccessProfileCount gets a reference to the given int32 and assigns it to the AccessProfileCount field.
func (o *IdentityDocumentAllOf) SetAccessProfileCount(v int32) {
	o.AccessProfileCount = &v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *IdentityDocumentAllOf) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetRoleCount returns the RoleCount field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetRoleCount() int32 {
	if o == nil || isNil(o.RoleCount) {
		var ret int32
		return ret
	}
	return *o.RoleCount
}

// GetRoleCountOk returns a tuple with the RoleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetRoleCountOk() (*int32, bool) {
	if o == nil || isNil(o.RoleCount) {
		return nil, false
	}
	return o.RoleCount, true
}

// HasRoleCount returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasRoleCount() bool {
	if o != nil && !isNil(o.RoleCount) {
		return true
	}

	return false
}

// SetRoleCount gets a reference to the given int32 and assigns it to the RoleCount field.
func (o *IdentityDocumentAllOf) SetRoleCount(v int32) {
	o.RoleCount = &v
}

// GetOwns returns the Owns field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetOwns() Owns {
	if o == nil || isNil(o.Owns) {
		var ret Owns
		return ret
	}
	return *o.Owns
}

// GetOwnsOk returns a tuple with the Owns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetOwnsOk() (*Owns, bool) {
	if o == nil || isNil(o.Owns) {
		return nil, false
	}
	return o.Owns, true
}

// HasOwns returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasOwns() bool {
	if o != nil && !isNil(o.Owns) {
		return true
	}

	return false
}

// SetOwns gets a reference to the given Owns and assigns it to the Owns field.
func (o *IdentityDocumentAllOf) SetOwns(v Owns) {
	o.Owns = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IdentityDocumentAllOf) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityDocumentAllOf) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IdentityDocumentAllOf) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IdentityDocumentAllOf) SetTags(v []string) {
	o.Tags = v
}

func (o IdentityDocumentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityDocumentAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !isNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if !isNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !isNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.EmployeeNumber) {
		toSerialize["employeeNumber"] = o.EmployeeNumber
	}
	if !isNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !isNil(o.IsManager) {
		toSerialize["isManager"] = o.IsManager
	}
	if !isNil(o.IdentityProfile) {
		toSerialize["identityProfile"] = o.IdentityProfile
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ProcessingState.IsSet() {
		toSerialize["processingState"] = o.ProcessingState.Get()
	}
	if !isNil(o.ProcessingDetails) {
		toSerialize["processingDetails"] = o.ProcessingDetails
	}
	if !isNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !isNil(o.AccountCount) {
		toSerialize["accountCount"] = o.AccountCount
	}
	if !isNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !isNil(o.AppCount) {
		toSerialize["appCount"] = o.AppCount
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.AccessCount) {
		toSerialize["accessCount"] = o.AccessCount
	}
	if !isNil(o.AccessProfileCount) {
		toSerialize["accessProfileCount"] = o.AccessProfileCount
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.RoleCount) {
		toSerialize["roleCount"] = o.RoleCount
	}
	if !isNil(o.Owns) {
		toSerialize["owns"] = o.Owns
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityDocumentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityDocumentAllOf := _IdentityDocumentAllOf{}

	if err = json.Unmarshal(bytes, &varIdentityDocumentAllOf); err == nil {
		*o = IdentityDocumentAllOf(varIdentityDocumentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "inactive")
		delete(additionalProperties, "protected")
		delete(additionalProperties, "status")
		delete(additionalProperties, "employeeNumber")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "isManager")
		delete(additionalProperties, "identityProfile")
		delete(additionalProperties, "source")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "processingState")
		delete(additionalProperties, "processingDetails")
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "accountCount")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "appCount")
		delete(additionalProperties, "access")
		delete(additionalProperties, "accessCount")
		delete(additionalProperties, "accessProfileCount")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "roleCount")
		delete(additionalProperties, "owns")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityDocumentAllOf struct {
	value *IdentityDocumentAllOf
	isSet bool
}

func (v NullableIdentityDocumentAllOf) Get() *IdentityDocumentAllOf {
	return v.value
}

func (v *NullableIdentityDocumentAllOf) Set(val *IdentityDocumentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityDocumentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityDocumentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityDocumentAllOf(val *IdentityDocumentAllOf) *NullableIdentityDocumentAllOf {
	return &NullableIdentityDocumentAllOf{value: val, isSet: true}
}

func (v NullableIdentityDocumentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityDocumentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


