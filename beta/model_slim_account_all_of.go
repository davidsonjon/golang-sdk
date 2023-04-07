/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the SlimAccountAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimAccountAllOf{}

// SlimAccountAllOf struct for SlimAccountAllOf
type SlimAccountAllOf struct {
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
	AdditionalProperties map[string]interface{}
}

type _SlimAccountAllOf SlimAccountAllOf

// NewSlimAccountAllOf instantiates a new SlimAccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimAccountAllOf() *SlimAccountAllOf {
	this := SlimAccountAllOf{}
	return &this
}

// NewSlimAccountAllOfWithDefaults instantiates a new SlimAccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimAccountAllOfWithDefaults() *SlimAccountAllOf {
	this := SlimAccountAllOf{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SlimAccountAllOf) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SlimAccountAllOf) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *SlimAccountAllOf) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *SlimAccountAllOf) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *SlimAccountAllOf) UnsetUuid() {
	o.Uuid.Unset()
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetNativeIdentity() string {
	if o == nil || isNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetNativeIdentityOk() (*string, bool) {
	if o == nil || isNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasNativeIdentity() bool {
	if o != nil && !isNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *SlimAccountAllOf) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SlimAccountAllOf) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SlimAccountAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SlimAccountAllOf) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SlimAccountAllOf) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SlimAccountAllOf) UnsetDescription() {
	o.Description.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *SlimAccountAllOf) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *SlimAccountAllOf) SetLocked(v bool) {
	o.Locked = &v
}

// GetManuallyCorrelated returns the ManuallyCorrelated field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetManuallyCorrelated() bool {
	if o == nil || isNil(o.ManuallyCorrelated) {
		var ret bool
		return ret
	}
	return *o.ManuallyCorrelated
}

// GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetManuallyCorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.ManuallyCorrelated) {
		return nil, false
	}
	return o.ManuallyCorrelated, true
}

// HasManuallyCorrelated returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasManuallyCorrelated() bool {
	if o != nil && !isNil(o.ManuallyCorrelated) {
		return true
	}

	return false
}

// SetManuallyCorrelated gets a reference to the given bool and assigns it to the ManuallyCorrelated field.
func (o *SlimAccountAllOf) SetManuallyCorrelated(v bool) {
	o.ManuallyCorrelated = &v
}

// GetHasEntitlements returns the HasEntitlements field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetHasEntitlements() bool {
	if o == nil || isNil(o.HasEntitlements) {
		var ret bool
		return ret
	}
	return *o.HasEntitlements
}

// GetHasEntitlementsOk returns a tuple with the HasEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetHasEntitlementsOk() (*bool, bool) {
	if o == nil || isNil(o.HasEntitlements) {
		return nil, false
	}
	return o.HasEntitlements, true
}

// HasHasEntitlements returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasHasEntitlements() bool {
	if o != nil && !isNil(o.HasEntitlements) {
		return true
	}

	return false
}

// SetHasEntitlements gets a reference to the given bool and assigns it to the HasEntitlements field.
func (o *SlimAccountAllOf) SetHasEntitlements(v bool) {
	o.HasEntitlements = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *SlimAccountAllOf) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *SlimAccountAllOf) SetSourceName(v string) {
	o.SourceName = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *SlimAccountAllOf) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SlimAccountAllOf) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimAccountAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SlimAccountAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *SlimAccountAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o SlimAccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimAccountAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SlimAccountAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSlimAccountAllOf := _SlimAccountAllOf{}

	if err = json.Unmarshal(bytes, &varSlimAccountAllOf); err == nil {
		*o = SlimAccountAllOf(varSlimAccountAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSlimAccountAllOf struct {
	value *SlimAccountAllOf
	isSet bool
}

func (v NullableSlimAccountAllOf) Get() *SlimAccountAllOf {
	return v.value
}

func (v *NullableSlimAccountAllOf) Set(val *SlimAccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimAccountAllOf(val *SlimAccountAllOf) *NullableSlimAccountAllOf {
	return &NullableSlimAccountAllOf{value: val, isSet: true}
}

func (v NullableSlimAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

