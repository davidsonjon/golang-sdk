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

// checks if the ReviewableEntitlementAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewableEntitlementAccount{}

// ReviewableEntitlementAccount Information about the status of the entitlement
type ReviewableEntitlementAccount struct {
	// The native identity for this account
	NativeIdentity *string `json:"nativeIdentity,omitempty"`
	// Indicates whether this account is currently disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Indicates whether this account is currently locked
	Locked *bool `json:"locked,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	// The id associated with the account
	Id NullableString `json:"id,omitempty"`
	// The account name
	Name NullableString `json:"name,omitempty"`
	// When the account was created
	Created NullableTime `json:"created,omitempty"`
	// When the account was last modified
	Modified NullableTime `json:"modified,omitempty"`
	ActivityInsights *ActivityInsights `json:"activityInsights,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewableEntitlementAccount ReviewableEntitlementAccount

// NewReviewableEntitlementAccount instantiates a new ReviewableEntitlementAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewableEntitlementAccount() *ReviewableEntitlementAccount {
	this := ReviewableEntitlementAccount{}
	var disabled bool = false
	this.Disabled = &disabled
	var locked bool = false
	this.Locked = &locked
	return &this
}

// NewReviewableEntitlementAccountWithDefaults instantiates a new ReviewableEntitlementAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewableEntitlementAccountWithDefaults() *ReviewableEntitlementAccount {
	this := ReviewableEntitlementAccount{}
	var disabled bool = false
	this.Disabled = &disabled
	var locked bool = false
	this.Locked = &locked
	return &this
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *ReviewableEntitlementAccount) GetNativeIdentity() string {
	if o == nil || isNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlementAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil || isNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasNativeIdentity() bool {
	if o != nil && !isNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *ReviewableEntitlementAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ReviewableEntitlementAccount) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlementAccount) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ReviewableEntitlementAccount) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ReviewableEntitlementAccount) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlementAccount) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *ReviewableEntitlementAccount) SetLocked(v bool) {
	o.Locked = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReviewableEntitlementAccount) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlementAccount) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *ReviewableEntitlementAccount) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlementAccount) GetId() string {
	if o == nil || isNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlementAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ReviewableEntitlementAccount) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ReviewableEntitlementAccount) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ReviewableEntitlementAccount) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlementAccount) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlementAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReviewableEntitlementAccount) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ReviewableEntitlementAccount) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReviewableEntitlementAccount) UnsetName() {
	o.Name.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlementAccount) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlementAccount) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *ReviewableEntitlementAccount) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *ReviewableEntitlementAccount) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *ReviewableEntitlementAccount) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlementAccount) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlementAccount) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *ReviewableEntitlementAccount) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *ReviewableEntitlementAccount) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *ReviewableEntitlementAccount) UnsetModified() {
	o.Modified.Unset()
}

// GetActivityInsights returns the ActivityInsights field value if set, zero value otherwise.
func (o *ReviewableEntitlementAccount) GetActivityInsights() ActivityInsights {
	if o == nil || isNil(o.ActivityInsights) {
		var ret ActivityInsights
		return ret
	}
	return *o.ActivityInsights
}

// GetActivityInsightsOk returns a tuple with the ActivityInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlementAccount) GetActivityInsightsOk() (*ActivityInsights, bool) {
	if o == nil || isNil(o.ActivityInsights) {
		return nil, false
	}
	return o.ActivityInsights, true
}

// HasActivityInsights returns a boolean if a field has been set.
func (o *ReviewableEntitlementAccount) HasActivityInsights() bool {
	if o != nil && !isNil(o.ActivityInsights) {
		return true
	}

	return false
}

// SetActivityInsights gets a reference to the given ActivityInsights and assigns it to the ActivityInsights field.
func (o *ReviewableEntitlementAccount) SetActivityInsights(v ActivityInsights) {
	o.ActivityInsights = &v
}

func (o ReviewableEntitlementAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewableEntitlementAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NativeIdentity) {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !isNil(o.ActivityInsights) {
		toSerialize["activityInsights"] = o.ActivityInsights
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewableEntitlementAccount) UnmarshalJSON(bytes []byte) (err error) {
	varReviewableEntitlementAccount := _ReviewableEntitlementAccount{}

	if err = json.Unmarshal(bytes, &varReviewableEntitlementAccount); err == nil {
	*o = ReviewableEntitlementAccount(varReviewableEntitlementAccount)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "activityInsights")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewableEntitlementAccount struct {
	value *ReviewableEntitlementAccount
	isSet bool
}

func (v NullableReviewableEntitlementAccount) Get() *ReviewableEntitlementAccount {
	return v.value
}

func (v *NullableReviewableEntitlementAccount) Set(val *ReviewableEntitlementAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewableEntitlementAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewableEntitlementAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewableEntitlementAccount(val *ReviewableEntitlementAccount) *NullableReviewableEntitlementAccount {
	return &NullableReviewableEntitlementAccount{value: val, isSet: true}
}

func (v NullableReviewableEntitlementAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewableEntitlementAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


