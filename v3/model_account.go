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

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// The unique ID of the source this account belongs to
	SourceId string `json:"sourceId"`
	// The display name of the source this account belongs to
	SourceName string `json:"sourceName"`
	// The unique ID of the identity this account is correlated to
	IdentityId *string `json:"identityId,omitempty"`
	// The account attributes that are aggregated
	Attributes map[string]interface{} `json:"attributes"`
	// Indicates if this account is from an authoritative source
	Authoritative bool `json:"authoritative"`
	// A description of the account
	Description NullableString `json:"description,omitempty"`
	// Indicates if the account is currently disabled
	Disabled bool `json:"disabled"`
	// Indicates if the account is currently locked
	Locked bool `json:"locked"`
	// The unique ID of the account generated by the source system
	NativeIdentity string `json:"nativeIdentity"`
	// If true, this is a user account within IdentityNow.  If false, this is an account from a source system.
	SystemAccount bool `json:"systemAccount"`
	// Indicates if this account is not correlated to an identity
	Uncorrelated bool `json:"uncorrelated"`
	// The unique ID of the account as determined by the account schema
	Uuid NullableString `json:"uuid,omitempty"`
	// Indicates if the account has been manually correlated to an identity
	ManuallyCorrelated bool `json:"manuallyCorrelated"`
	// Indicates if the account has entitlements
	HasEntitlements bool `json:"hasEntitlements"`
	AdditionalProperties map[string]interface{}
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(name string, sourceId string, sourceName string, attributes map[string]interface{}, authoritative bool, disabled bool, locked bool, nativeIdentity string, systemAccount bool, uncorrelated bool, manuallyCorrelated bool, hasEntitlements bool) *Account {
	this := Account{}
	this.Name = name
	this.SourceId = sourceId
	this.SourceName = sourceName
	this.Attributes = attributes
	this.Authoritative = authoritative
	this.Disabled = disabled
	this.Locked = locked
	this.NativeIdentity = nativeIdentity
	this.SystemAccount = systemAccount
	this.Uncorrelated = uncorrelated
	this.ManuallyCorrelated = manuallyCorrelated
	this.HasEntitlements = hasEntitlements
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Account) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Account) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Account) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Account) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Account) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Account) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Account) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Account) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Account) SetModified(v time.Time) {
	o.Modified = &v
}

// GetSourceId returns the SourceId field value
func (o *Account) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *Account) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *Account) SetSourceId(v string) {
	o.SourceId = v
}

// GetSourceName returns the SourceName field value
func (o *Account) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *Account) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *Account) SetSourceName(v string) {
	o.SourceName = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *Account) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *Account) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *Account) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAttributes returns the Attributes field value
func (o *Account) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Account) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *Account) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetAuthoritative returns the Authoritative field value
func (o *Account) GetAuthoritative() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value
// and a boolean to check if the value has been set.
func (o *Account) GetAuthoritativeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authoritative, true
}

// SetAuthoritative sets field value
func (o *Account) SetAuthoritative(v bool) {
	o.Authoritative = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Account) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Account) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Account) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Account) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Account) UnsetDescription() {
	o.Description.Unset()
}

// GetDisabled returns the Disabled field value
func (o *Account) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *Account) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *Account) SetDisabled(v bool) {
	o.Disabled = v
}

// GetLocked returns the Locked field value
func (o *Account) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *Account) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *Account) SetLocked(v bool) {
	o.Locked = v
}

// GetNativeIdentity returns the NativeIdentity field value
func (o *Account) GetNativeIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value
// and a boolean to check if the value has been set.
func (o *Account) GetNativeIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentity, true
}

// SetNativeIdentity sets field value
func (o *Account) SetNativeIdentity(v string) {
	o.NativeIdentity = v
}

// GetSystemAccount returns the SystemAccount field value
func (o *Account) GetSystemAccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SystemAccount
}

// GetSystemAccountOk returns a tuple with the SystemAccount field value
// and a boolean to check if the value has been set.
func (o *Account) GetSystemAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemAccount, true
}

// SetSystemAccount sets field value
func (o *Account) SetSystemAccount(v bool) {
	o.SystemAccount = v
}

// GetUncorrelated returns the Uncorrelated field value
func (o *Account) GetUncorrelated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Uncorrelated
}

// GetUncorrelatedOk returns a tuple with the Uncorrelated field value
// and a boolean to check if the value has been set.
func (o *Account) GetUncorrelatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uncorrelated, true
}

// SetUncorrelated sets field value
func (o *Account) SetUncorrelated(v bool) {
	o.Uncorrelated = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Account) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *Account) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *Account) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *Account) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *Account) UnsetUuid() {
	o.Uuid.Unset()
}

// GetManuallyCorrelated returns the ManuallyCorrelated field value
func (o *Account) GetManuallyCorrelated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ManuallyCorrelated
}

// GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field value
// and a boolean to check if the value has been set.
func (o *Account) GetManuallyCorrelatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManuallyCorrelated, true
}

// SetManuallyCorrelated sets field value
func (o *Account) SetManuallyCorrelated(v bool) {
	o.ManuallyCorrelated = v
}

// GetHasEntitlements returns the HasEntitlements field value
func (o *Account) GetHasEntitlements() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasEntitlements
}

// GetHasEntitlementsOk returns a tuple with the HasEntitlements field value
// and a boolean to check if the value has been set.
func (o *Account) GetHasEntitlementsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasEntitlements, true
}

// SetHasEntitlements sets field value
func (o *Account) SetHasEntitlements(v bool) {
	o.HasEntitlements = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: created is readOnly
	// skip: modified is readOnly
	toSerialize["sourceId"] = o.SourceId
	toSerialize["sourceName"] = o.SourceName
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["authoritative"] = o.Authoritative
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["locked"] = o.Locked
	toSerialize["nativeIdentity"] = o.NativeIdentity
	toSerialize["systemAccount"] = o.SystemAccount
	toSerialize["uncorrelated"] = o.Uncorrelated
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	toSerialize["manuallyCorrelated"] = o.ManuallyCorrelated
	toSerialize["hasEntitlements"] = o.HasEntitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Account) UnmarshalJSON(bytes []byte) (err error) {
	varAccount := _Account{}

	if err = json.Unmarshal(bytes, &varAccount); err == nil {
		*o = Account(varAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "authoritative")
		delete(additionalProperties, "description")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "systemAccount")
		delete(additionalProperties, "uncorrelated")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "manuallyCorrelated")
		delete(additionalProperties, "hasEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


