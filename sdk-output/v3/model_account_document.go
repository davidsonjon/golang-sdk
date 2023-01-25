/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"time"
	"encoding/json"
)

// AccountDocument Account
type AccountDocument struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type DocumentType `json:"_type"`
	// The ID of the account
	AccountId *string `json:"accountId,omitempty"`
	Source *Source1 `json:"source,omitempty"`
	// Indicates if the account is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Indicates if the account is locked
	Locked *bool `json:"locked,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	// Indicates if the account has been manually correlated to an identity
	ManuallyCorrelated *bool `json:"manuallyCorrelated,omitempty"`
	// A date-time in ISO-8601 format
	PasswordLastSet NullableTime `json:"passwordLastSet,omitempty"`
	// a map or dictionary of key/value pairs
	EntitlementAttributes map[string]interface{} `json:"entitlementAttributes,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// a map or dictionary of key/value pairs
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Identity *DisplayReference `json:"identity,omitempty"`
	Access []Entitlement1 `json:"access,omitempty"`
	// The number of entitlements assigned to the account
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// Indicates if the account is not correlated to an identity
	Uncorrelated *bool `json:"uncorrelated,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountDocument AccountDocument

// NewAccountDocument instantiates a new AccountDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDocument(id string, name string, type_ DocumentType) *AccountDocument {
	this := AccountDocument{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewAccountDocumentWithDefaults instantiates a new AccountDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDocumentWithDefaults() *AccountDocument {
	this := AccountDocument{}
	return &this
}

// GetId returns the Id field value
func (o *AccountDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountDocument) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccountDocument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountDocument) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AccountDocument) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountDocument) SetType(v DocumentType) {
	o.Type = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccountDocument) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccountDocument) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccountDocument) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AccountDocument) GetSource() Source1 {
	if o == nil || isNil(o.Source) {
		var ret Source1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetSourceOk() (*Source1, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AccountDocument) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source1 and assigns it to the Source field.
func (o *AccountDocument) SetSource(v Source1) {
	o.Source = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AccountDocument) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AccountDocument) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AccountDocument) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *AccountDocument) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *AccountDocument) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *AccountDocument) SetLocked(v bool) {
	o.Locked = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *AccountDocument) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *AccountDocument) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *AccountDocument) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetManuallyCorrelated returns the ManuallyCorrelated field value if set, zero value otherwise.
func (o *AccountDocument) GetManuallyCorrelated() bool {
	if o == nil || isNil(o.ManuallyCorrelated) {
		var ret bool
		return ret
	}
	return *o.ManuallyCorrelated
}

// GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetManuallyCorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.ManuallyCorrelated) {
		return nil, false
	}
	return o.ManuallyCorrelated, true
}

// HasManuallyCorrelated returns a boolean if a field has been set.
func (o *AccountDocument) HasManuallyCorrelated() bool {
	if o != nil && !isNil(o.ManuallyCorrelated) {
		return true
	}

	return false
}

// SetManuallyCorrelated gets a reference to the given bool and assigns it to the ManuallyCorrelated field.
func (o *AccountDocument) SetManuallyCorrelated(v bool) {
	o.ManuallyCorrelated = &v
}

// GetPasswordLastSet returns the PasswordLastSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountDocument) GetPasswordLastSet() time.Time {
	if o == nil || isNil(o.PasswordLastSet.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PasswordLastSet.Get()
}

// GetPasswordLastSetOk returns a tuple with the PasswordLastSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountDocument) GetPasswordLastSetOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordLastSet.Get(), o.PasswordLastSet.IsSet()
}

// HasPasswordLastSet returns a boolean if a field has been set.
func (o *AccountDocument) HasPasswordLastSet() bool {
	if o != nil && o.PasswordLastSet.IsSet() {
		return true
	}

	return false
}

// SetPasswordLastSet gets a reference to the given NullableTime and assigns it to the PasswordLastSet field.
func (o *AccountDocument) SetPasswordLastSet(v time.Time) {
	o.PasswordLastSet.Set(&v)
}
// SetPasswordLastSetNil sets the value for PasswordLastSet to be an explicit nil
func (o *AccountDocument) SetPasswordLastSetNil() {
	o.PasswordLastSet.Set(nil)
}

// UnsetPasswordLastSet ensures that no value is present for PasswordLastSet, not even an explicit nil
func (o *AccountDocument) UnsetPasswordLastSet() {
	o.PasswordLastSet.Unset()
}

// GetEntitlementAttributes returns the EntitlementAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountDocument) GetEntitlementAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EntitlementAttributes
}

// GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountDocument) GetEntitlementAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.EntitlementAttributes) {
		return map[string]interface{}{}, false
	}
	return o.EntitlementAttributes, true
}

// HasEntitlementAttributes returns a boolean if a field has been set.
func (o *AccountDocument) HasEntitlementAttributes() bool {
	if o != nil && isNil(o.EntitlementAttributes) {
		return true
	}

	return false
}

// SetEntitlementAttributes gets a reference to the given map[string]interface{} and assigns it to the EntitlementAttributes field.
func (o *AccountDocument) SetEntitlementAttributes(v map[string]interface{}) {
	o.EntitlementAttributes = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountDocument) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountDocument) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *AccountDocument) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *AccountDocument) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *AccountDocument) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *AccountDocument) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountDocument) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountDocument) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *AccountDocument) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *AccountDocument) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *AccountDocument) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *AccountDocument) UnsetModified() {
	o.Modified.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AccountDocument) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AccountDocument) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *AccountDocument) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AccountDocument) GetIdentity() DisplayReference {
	if o == nil || isNil(o.Identity) {
		var ret DisplayReference
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetIdentityOk() (*DisplayReference, bool) {
	if o == nil || isNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AccountDocument) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given DisplayReference and assigns it to the Identity field.
func (o *AccountDocument) SetIdentity(v DisplayReference) {
	o.Identity = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AccountDocument) GetAccess() []Entitlement1 {
	if o == nil || isNil(o.Access) {
		var ret []Entitlement1
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetAccessOk() ([]Entitlement1, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AccountDocument) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []Entitlement1 and assigns it to the Access field.
func (o *AccountDocument) SetAccess(v []Entitlement1) {
	o.Access = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccountDocument) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccountDocument) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *AccountDocument) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetUncorrelated returns the Uncorrelated field value if set, zero value otherwise.
func (o *AccountDocument) GetUncorrelated() bool {
	if o == nil || isNil(o.Uncorrelated) {
		var ret bool
		return ret
	}
	return *o.Uncorrelated
}

// GetUncorrelatedOk returns a tuple with the Uncorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetUncorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.Uncorrelated) {
		return nil, false
	}
	return o.Uncorrelated, true
}

// HasUncorrelated returns a boolean if a field has been set.
func (o *AccountDocument) HasUncorrelated() bool {
	if o != nil && !isNil(o.Uncorrelated) {
		return true
	}

	return false
}

// SetUncorrelated gets a reference to the given bool and assigns it to the Uncorrelated field.
func (o *AccountDocument) SetUncorrelated(v bool) {
	o.Uncorrelated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AccountDocument) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AccountDocument) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AccountDocument) SetTags(v []string) {
	o.Tags = v
}

func (o AccountDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.ManuallyCorrelated) {
		toSerialize["manuallyCorrelated"] = o.ManuallyCorrelated
	}
	if o.PasswordLastSet.IsSet() {
		toSerialize["passwordLastSet"] = o.PasswordLastSet.Get()
	}
	if o.EntitlementAttributes != nil {
		toSerialize["entitlementAttributes"] = o.EntitlementAttributes
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.Uncorrelated) {
		toSerialize["uncorrelated"] = o.Uncorrelated
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountDocument) UnmarshalJSON(bytes []byte) (err error) {
	varAccountDocument := _AccountDocument{}

	if err = json.Unmarshal(bytes, &varAccountDocument); err == nil {
		*o = AccountDocument(varAccountDocument)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "source")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "manuallyCorrelated")
		delete(additionalProperties, "passwordLastSet")
		delete(additionalProperties, "entitlementAttributes")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "access")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "uncorrelated")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountDocument struct {
	value *AccountDocument
	isSet bool
}

func (v NullableAccountDocument) Get() *AccountDocument {
	return v.value
}

func (v *NullableAccountDocument) Set(val *AccountDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDocument(val *AccountDocument) *NullableAccountDocument {
	return &NullableAccountDocument{value: val, isSet: true}
}

func (v NullableAccountDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


