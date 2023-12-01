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
	"fmt"
)

// checks if the FullAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullAccount{}

// FullAccount struct for FullAccount
type FullAccount struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// Unique ID from the owning source
	Uuid NullableString `json:"uuid,omitempty"`
	// The native identifier of the account
	NativeIdentity *string `json:"nativeIdentity,omitempty"`
	// The description for the account
	Description NullableString `json:"description,omitempty"`
	// Whether the account is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Whether the account is locked
	Locked *bool `json:"locked,omitempty"`
	// Whether the account was manually correlated
	ManuallyCorrelated *bool `json:"manuallyCorrelated,omitempty"`
	// Whether the account has any entitlements associated with it
	HasEntitlements *bool `json:"hasEntitlements,omitempty"`
	// The ID of the source for which this account belongs
	SourceId *string `json:"sourceId,omitempty"`
	// The name of the source
	SourceName *string `json:"sourceName,omitempty"`
	// The ID of the identity for which this account is correlated to if not uncorrelated
	IdentityId *string `json:"identityId,omitempty"`
	// A map containing attributes associated with the account
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Whether this account belongs to an authoritative source
	Authoritative *bool `json:"authoritative,omitempty"`
	// Whether this account is for the IdentityNow source
	SystemAccount *bool `json:"systemAccount,omitempty"`
	// True if this account is not correlated to an identity
	Uncorrelated *bool `json:"uncorrelated,omitempty"`
	// A string list containing the owning source's features
	Features *string `json:"features,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullAccount FullAccount

// NewFullAccount instantiates a new FullAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullAccount(name string) *FullAccount {
	this := FullAccount{}
	this.Name = name
	return &this
}

// NewFullAccountWithDefaults instantiates a new FullAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullAccountWithDefaults() *FullAccount {
	this := FullAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullAccount) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullAccount) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullAccount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *FullAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FullAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FullAccount) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FullAccount) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FullAccount) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FullAccount) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *FullAccount) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *FullAccount) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *FullAccount) SetModified(v time.Time) {
	o.Modified = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FullAccount) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FullAccount) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *FullAccount) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *FullAccount) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *FullAccount) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *FullAccount) UnsetUuid() {
	o.Uuid.Unset()
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *FullAccount) GetNativeIdentity() string {
	if o == nil || isNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil || isNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *FullAccount) HasNativeIdentity() bool {
	if o != nil && !isNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *FullAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FullAccount) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FullAccount) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *FullAccount) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *FullAccount) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *FullAccount) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *FullAccount) UnsetDescription() {
	o.Description.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *FullAccount) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *FullAccount) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *FullAccount) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *FullAccount) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *FullAccount) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *FullAccount) SetLocked(v bool) {
	o.Locked = &v
}

// GetManuallyCorrelated returns the ManuallyCorrelated field value if set, zero value otherwise.
func (o *FullAccount) GetManuallyCorrelated() bool {
	if o == nil || isNil(o.ManuallyCorrelated) {
		var ret bool
		return ret
	}
	return *o.ManuallyCorrelated
}

// GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetManuallyCorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.ManuallyCorrelated) {
		return nil, false
	}
	return o.ManuallyCorrelated, true
}

// HasManuallyCorrelated returns a boolean if a field has been set.
func (o *FullAccount) HasManuallyCorrelated() bool {
	if o != nil && !isNil(o.ManuallyCorrelated) {
		return true
	}

	return false
}

// SetManuallyCorrelated gets a reference to the given bool and assigns it to the ManuallyCorrelated field.
func (o *FullAccount) SetManuallyCorrelated(v bool) {
	o.ManuallyCorrelated = &v
}

// GetHasEntitlements returns the HasEntitlements field value if set, zero value otherwise.
func (o *FullAccount) GetHasEntitlements() bool {
	if o == nil || isNil(o.HasEntitlements) {
		var ret bool
		return ret
	}
	return *o.HasEntitlements
}

// GetHasEntitlementsOk returns a tuple with the HasEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetHasEntitlementsOk() (*bool, bool) {
	if o == nil || isNil(o.HasEntitlements) {
		return nil, false
	}
	return o.HasEntitlements, true
}

// HasHasEntitlements returns a boolean if a field has been set.
func (o *FullAccount) HasHasEntitlements() bool {
	if o != nil && !isNil(o.HasEntitlements) {
		return true
	}

	return false
}

// SetHasEntitlements gets a reference to the given bool and assigns it to the HasEntitlements field.
func (o *FullAccount) SetHasEntitlements(v bool) {
	o.HasEntitlements = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *FullAccount) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *FullAccount) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *FullAccount) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *FullAccount) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *FullAccount) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *FullAccount) SetSourceName(v string) {
	o.SourceName = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *FullAccount) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *FullAccount) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *FullAccount) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FullAccount) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FullAccount) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *FullAccount) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetAuthoritative returns the Authoritative field value if set, zero value otherwise.
func (o *FullAccount) GetAuthoritative() bool {
	if o == nil || isNil(o.Authoritative) {
		var ret bool
		return ret
	}
	return *o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetAuthoritativeOk() (*bool, bool) {
	if o == nil || isNil(o.Authoritative) {
		return nil, false
	}
	return o.Authoritative, true
}

// HasAuthoritative returns a boolean if a field has been set.
func (o *FullAccount) HasAuthoritative() bool {
	if o != nil && !isNil(o.Authoritative) {
		return true
	}

	return false
}

// SetAuthoritative gets a reference to the given bool and assigns it to the Authoritative field.
func (o *FullAccount) SetAuthoritative(v bool) {
	o.Authoritative = &v
}

// GetSystemAccount returns the SystemAccount field value if set, zero value otherwise.
func (o *FullAccount) GetSystemAccount() bool {
	if o == nil || isNil(o.SystemAccount) {
		var ret bool
		return ret
	}
	return *o.SystemAccount
}

// GetSystemAccountOk returns a tuple with the SystemAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetSystemAccountOk() (*bool, bool) {
	if o == nil || isNil(o.SystemAccount) {
		return nil, false
	}
	return o.SystemAccount, true
}

// HasSystemAccount returns a boolean if a field has been set.
func (o *FullAccount) HasSystemAccount() bool {
	if o != nil && !isNil(o.SystemAccount) {
		return true
	}

	return false
}

// SetSystemAccount gets a reference to the given bool and assigns it to the SystemAccount field.
func (o *FullAccount) SetSystemAccount(v bool) {
	o.SystemAccount = &v
}

// GetUncorrelated returns the Uncorrelated field value if set, zero value otherwise.
func (o *FullAccount) GetUncorrelated() bool {
	if o == nil || isNil(o.Uncorrelated) {
		var ret bool
		return ret
	}
	return *o.Uncorrelated
}

// GetUncorrelatedOk returns a tuple with the Uncorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetUncorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.Uncorrelated) {
		return nil, false
	}
	return o.Uncorrelated, true
}

// HasUncorrelated returns a boolean if a field has been set.
func (o *FullAccount) HasUncorrelated() bool {
	if o != nil && !isNil(o.Uncorrelated) {
		return true
	}

	return false
}

// SetUncorrelated gets a reference to the given bool and assigns it to the Uncorrelated field.
func (o *FullAccount) SetUncorrelated(v bool) {
	o.Uncorrelated = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FullAccount) GetFeatures() string {
	if o == nil || isNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAccount) GetFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FullAccount) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *FullAccount) SetFeatures(v string) {
	o.Features = &v
}

func (o FullAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: created is readOnly
	// skip: modified is readOnly
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if !isNil(o.NativeIdentity) {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !isNil(o.ManuallyCorrelated) {
		toSerialize["manuallyCorrelated"] = o.ManuallyCorrelated
	}
	if !isNil(o.HasEntitlements) {
		toSerialize["hasEntitlements"] = o.HasEntitlements
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Authoritative) {
		toSerialize["authoritative"] = o.Authoritative
	}
	if !isNil(o.SystemAccount) {
		toSerialize["systemAccount"] = o.SystemAccount
	}
	if !isNil(o.Uncorrelated) {
		toSerialize["uncorrelated"] = o.Uncorrelated
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullAccount) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varFullAccount := _FullAccount{}

	if err = json.Unmarshal(bytes, &varFullAccount); err == nil {
	*o = FullAccount(varFullAccount)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "description")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "manuallyCorrelated")
		delete(additionalProperties, "hasEntitlements")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "authoritative")
		delete(additionalProperties, "systemAccount")
		delete(additionalProperties, "uncorrelated")
		delete(additionalProperties, "features")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullAccount struct {
	value *FullAccount
	isSet bool
}

func (v NullableFullAccount) Get() *FullAccount {
	return v.value
}

func (v *NullableFullAccount) Set(val *FullAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableFullAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableFullAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullAccount(val *FullAccount) *NullableFullAccount {
	return &NullableFullAccount{value: val, isSet: true}
}

func (v NullableFullAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


