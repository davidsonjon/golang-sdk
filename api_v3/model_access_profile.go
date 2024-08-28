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

// checks if the AccessProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfile{}

// AccessProfile Access Profile
type AccessProfile struct {
	// The ID of the Access Profile
	Id *string `json:"id,omitempty"`
	// Name of the Access Profile
	Name string `json:"name"`
	// Information about the Access Profile
	Description NullableString `json:"description,omitempty"`
	// Date the Access Profile was created
	Created *time.Time `json:"created,omitempty"`
	// Date the Access Profile was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Whether the Access Profile is enabled. If the Access Profile is enabled then you must include at least one Entitlement.
	Enabled *bool `json:"enabled,omitempty"`
	Owner OwnerReference `json:"owner"`
	Source AccessProfileSourceRef `json:"source"`
	// A list of entitlements associated with the Access Profile. If enabled is false this is allowed to be empty otherwise it needs to contain at least one Entitlement.
	Entitlements []EntitlementRef `json:"entitlements,omitempty"`
	// Whether the Access Profile is requestable via access request. Currently, making an Access Profile non-requestable is only supported  for customers enabled with the new Request Center. Otherwise, attempting to create an Access Profile with a value  **false** in this field results in a 400 error.
	Requestable *bool `json:"requestable,omitempty"`
	AccessRequestConfig NullableRequestability `json:"accessRequestConfig,omitempty"`
	RevocationRequestConfig NullableRevocability `json:"revocationRequestConfig,omitempty"`
	// List of IDs of segments, if any, to which this Access Profile is assigned.
	Segments []string `json:"segments,omitempty"`
	ProvisioningCriteria NullableProvisioningCriteriaLevel1 `json:"provisioningCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfile AccessProfile

// NewAccessProfile instantiates a new AccessProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfile(name string, owner OwnerReference, source AccessProfileSourceRef) *AccessProfile {
	this := AccessProfile{}
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	this.Owner = owner
	this.Source = source
	var requestable bool = true
	this.Requestable = &requestable
	return &this
}

// NewAccessProfileWithDefaults instantiates a new AccessProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileWithDefaults() *AccessProfile {
	this := AccessProfile{}
	var enabled bool = true
	this.Enabled = &enabled
	var requestable bool = true
	this.Requestable = &requestable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessProfile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessProfile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessProfile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *AccessProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessProfile) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfile) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfile) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessProfile) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessProfile) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessProfile) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessProfile) UnsetDescription() {
	o.Description.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AccessProfile) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AccessProfile) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AccessProfile) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AccessProfile) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AccessProfile) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *AccessProfile) SetModified(v time.Time) {
	o.Modified = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AccessProfile) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AccessProfile) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AccessProfile) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOwner returns the Owner field value
func (o *AccessProfile) GetOwner() OwnerReference {
	if o == nil {
		var ret OwnerReference
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetOwnerOk() (*OwnerReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *AccessProfile) SetOwner(v OwnerReference) {
	o.Owner = v
}

// GetSource returns the Source field value
func (o *AccessProfile) GetSource() AccessProfileSourceRef {
	if o == nil {
		var ret AccessProfileSourceRef
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetSourceOk() (*AccessProfileSourceRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AccessProfile) SetSource(v AccessProfileSourceRef) {
	o.Source = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfile) GetEntitlements() []EntitlementRef {
	if o == nil {
		var ret []EntitlementRef
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfile) GetEntitlementsOk() ([]EntitlementRef, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *AccessProfile) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementRef and assigns it to the Entitlements field.
func (o *AccessProfile) SetEntitlements(v []EntitlementRef) {
	o.Entitlements = v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *AccessProfile) GetRequestable() bool {
	if o == nil || IsNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfile) GetRequestableOk() (*bool, bool) {
	if o == nil || IsNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *AccessProfile) HasRequestable() bool {
	if o != nil && !IsNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *AccessProfile) SetRequestable(v bool) {
	o.Requestable = &v
}

// GetAccessRequestConfig returns the AccessRequestConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfile) GetAccessRequestConfig() Requestability {
	if o == nil || IsNil(o.AccessRequestConfig.Get()) {
		var ret Requestability
		return ret
	}
	return *o.AccessRequestConfig.Get()
}

// GetAccessRequestConfigOk returns a tuple with the AccessRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfile) GetAccessRequestConfigOk() (*Requestability, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessRequestConfig.Get(), o.AccessRequestConfig.IsSet()
}

// HasAccessRequestConfig returns a boolean if a field has been set.
func (o *AccessProfile) HasAccessRequestConfig() bool {
	if o != nil && o.AccessRequestConfig.IsSet() {
		return true
	}

	return false
}

// SetAccessRequestConfig gets a reference to the given NullableRequestability and assigns it to the AccessRequestConfig field.
func (o *AccessProfile) SetAccessRequestConfig(v Requestability) {
	o.AccessRequestConfig.Set(&v)
}
// SetAccessRequestConfigNil sets the value for AccessRequestConfig to be an explicit nil
func (o *AccessProfile) SetAccessRequestConfigNil() {
	o.AccessRequestConfig.Set(nil)
}

// UnsetAccessRequestConfig ensures that no value is present for AccessRequestConfig, not even an explicit nil
func (o *AccessProfile) UnsetAccessRequestConfig() {
	o.AccessRequestConfig.Unset()
}

// GetRevocationRequestConfig returns the RevocationRequestConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfile) GetRevocationRequestConfig() Revocability {
	if o == nil || IsNil(o.RevocationRequestConfig.Get()) {
		var ret Revocability
		return ret
	}
	return *o.RevocationRequestConfig.Get()
}

// GetRevocationRequestConfigOk returns a tuple with the RevocationRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfile) GetRevocationRequestConfigOk() (*Revocability, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationRequestConfig.Get(), o.RevocationRequestConfig.IsSet()
}

// HasRevocationRequestConfig returns a boolean if a field has been set.
func (o *AccessProfile) HasRevocationRequestConfig() bool {
	if o != nil && o.RevocationRequestConfig.IsSet() {
		return true
	}

	return false
}

// SetRevocationRequestConfig gets a reference to the given NullableRevocability and assigns it to the RevocationRequestConfig field.
func (o *AccessProfile) SetRevocationRequestConfig(v Revocability) {
	o.RevocationRequestConfig.Set(&v)
}
// SetRevocationRequestConfigNil sets the value for RevocationRequestConfig to be an explicit nil
func (o *AccessProfile) SetRevocationRequestConfigNil() {
	o.RevocationRequestConfig.Set(nil)
}

// UnsetRevocationRequestConfig ensures that no value is present for RevocationRequestConfig, not even an explicit nil
func (o *AccessProfile) UnsetRevocationRequestConfig() {
	o.RevocationRequestConfig.Unset()
}

// GetSegments returns the Segments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfile) GetSegments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfile) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *AccessProfile) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *AccessProfile) SetSegments(v []string) {
	o.Segments = v
}

// GetProvisioningCriteria returns the ProvisioningCriteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfile) GetProvisioningCriteria() ProvisioningCriteriaLevel1 {
	if o == nil || IsNil(o.ProvisioningCriteria.Get()) {
		var ret ProvisioningCriteriaLevel1
		return ret
	}
	return *o.ProvisioningCriteria.Get()
}

// GetProvisioningCriteriaOk returns a tuple with the ProvisioningCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfile) GetProvisioningCriteriaOk() (*ProvisioningCriteriaLevel1, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisioningCriteria.Get(), o.ProvisioningCriteria.IsSet()
}

// HasProvisioningCriteria returns a boolean if a field has been set.
func (o *AccessProfile) HasProvisioningCriteria() bool {
	if o != nil && o.ProvisioningCriteria.IsSet() {
		return true
	}

	return false
}

// SetProvisioningCriteria gets a reference to the given NullableProvisioningCriteriaLevel1 and assigns it to the ProvisioningCriteria field.
func (o *AccessProfile) SetProvisioningCriteria(v ProvisioningCriteriaLevel1) {
	o.ProvisioningCriteria.Set(&v)
}
// SetProvisioningCriteriaNil sets the value for ProvisioningCriteria to be an explicit nil
func (o *AccessProfile) SetProvisioningCriteriaNil() {
	o.ProvisioningCriteria.Set(nil)
}

// UnsetProvisioningCriteria ensures that no value is present for ProvisioningCriteria, not even an explicit nil
func (o *AccessProfile) UnsetProvisioningCriteria() {
	o.ProvisioningCriteria.Unset()
}

func (o AccessProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["owner"] = o.Owner
	toSerialize["source"] = o.Source
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if o.AccessRequestConfig.IsSet() {
		toSerialize["accessRequestConfig"] = o.AccessRequestConfig.Get()
	}
	if o.RevocationRequestConfig.IsSet() {
		toSerialize["revocationRequestConfig"] = o.RevocationRequestConfig.Get()
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	if o.ProvisioningCriteria.IsSet() {
		toSerialize["provisioningCriteria"] = o.ProvisioningCriteria.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner",
		"source",
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

	varAccessProfile := _AccessProfile{}

	err = json.Unmarshal(data, &varAccessProfile)

	if err != nil {
		return err
	}

	*o = AccessProfile(varAccessProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "source")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "accessRequestConfig")
		delete(additionalProperties, "revocationRequestConfig")
		delete(additionalProperties, "segments")
		delete(additionalProperties, "provisioningCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfile struct {
	value *AccessProfile
	isSet bool
}

func (v NullableAccessProfile) Get() *AccessProfile {
	return v.value
}

func (v *NullableAccessProfile) Set(val *AccessProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfile(val *AccessProfile) *NullableAccessProfile {
	return &NullableAccessProfile{value: val, isSet: true}
}

func (v NullableAccessProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


