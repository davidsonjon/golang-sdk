/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the RoleMiningEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningEntitlement{}

// RoleMiningEntitlement struct for RoleMiningEntitlement
type RoleMiningEntitlement struct {
	EntitlementRef *RoleMiningEntitlementRef `json:"entitlementRef,omitempty"`
	// Name of the entitlement
	Name *string `json:"name,omitempty"`
	// Application name of the entitlement
	ApplicationName *string `json:"applicationName,omitempty"`
	// The number of identities with this entitlement in a role.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	// The % popularity of this entitlement in a role.
	Popularity *float32 `json:"popularity,omitempty"`
	// The % popularity of this entitlement in the org.
	PopularityInOrg *float32 `json:"popularityInOrg,omitempty"`
	// The ID of the source/application.
	SourceId *string `json:"sourceId,omitempty"`
	// The status of activity data for the source.   Value is complete or notComplete.
	ActivitySourceState NullableString `json:"activitySourceState,omitempty"`
	// The percentage of identities in the potential role that have usage of the source/application of this entitlement.
	SourceUsagePercent NullableFloat32 `json:"sourceUsagePercent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningEntitlement RoleMiningEntitlement

// NewRoleMiningEntitlement instantiates a new RoleMiningEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningEntitlement() *RoleMiningEntitlement {
	this := RoleMiningEntitlement{}
	return &this
}

// NewRoleMiningEntitlementWithDefaults instantiates a new RoleMiningEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningEntitlementWithDefaults() *RoleMiningEntitlement {
	this := RoleMiningEntitlement{}
	return &this
}

// GetEntitlementRef returns the EntitlementRef field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetEntitlementRef() RoleMiningEntitlementRef {
	if o == nil || isNil(o.EntitlementRef) {
		var ret RoleMiningEntitlementRef
		return ret
	}
	return *o.EntitlementRef
}

// GetEntitlementRefOk returns a tuple with the EntitlementRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetEntitlementRefOk() (*RoleMiningEntitlementRef, bool) {
	if o == nil || isNil(o.EntitlementRef) {
		return nil, false
	}
	return o.EntitlementRef, true
}

// HasEntitlementRef returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasEntitlementRef() bool {
	if o != nil && !isNil(o.EntitlementRef) {
		return true
	}

	return false
}

// SetEntitlementRef gets a reference to the given RoleMiningEntitlementRef and assigns it to the EntitlementRef field.
func (o *RoleMiningEntitlement) SetEntitlementRef(v RoleMiningEntitlementRef) {
	o.EntitlementRef = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleMiningEntitlement) SetName(v string) {
	o.Name = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetApplicationName() string {
	if o == nil || isNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetApplicationNameOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasApplicationName() bool {
	if o != nil && !isNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *RoleMiningEntitlement) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *RoleMiningEntitlement) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetPopularity() float32 {
	if o == nil || isNil(o.Popularity) {
		var ret float32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetPopularityOk() (*float32, bool) {
	if o == nil || isNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasPopularity() bool {
	if o != nil && !isNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given float32 and assigns it to the Popularity field.
func (o *RoleMiningEntitlement) SetPopularity(v float32) {
	o.Popularity = &v
}

// GetPopularityInOrg returns the PopularityInOrg field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetPopularityInOrg() float32 {
	if o == nil || isNil(o.PopularityInOrg) {
		var ret float32
		return ret
	}
	return *o.PopularityInOrg
}

// GetPopularityInOrgOk returns a tuple with the PopularityInOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetPopularityInOrgOk() (*float32, bool) {
	if o == nil || isNil(o.PopularityInOrg) {
		return nil, false
	}
	return o.PopularityInOrg, true
}

// HasPopularityInOrg returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasPopularityInOrg() bool {
	if o != nil && !isNil(o.PopularityInOrg) {
		return true
	}

	return false
}

// SetPopularityInOrg gets a reference to the given float32 and assigns it to the PopularityInOrg field.
func (o *RoleMiningEntitlement) SetPopularityInOrg(v float32) {
	o.PopularityInOrg = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *RoleMiningEntitlement) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlement) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *RoleMiningEntitlement) SetSourceId(v string) {
	o.SourceId = &v
}

// GetActivitySourceState returns the ActivitySourceState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningEntitlement) GetActivitySourceState() string {
	if o == nil || isNil(o.ActivitySourceState.Get()) {
		var ret string
		return ret
	}
	return *o.ActivitySourceState.Get()
}

// GetActivitySourceStateOk returns a tuple with the ActivitySourceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningEntitlement) GetActivitySourceStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivitySourceState.Get(), o.ActivitySourceState.IsSet()
}

// HasActivitySourceState returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasActivitySourceState() bool {
	if o != nil && o.ActivitySourceState.IsSet() {
		return true
	}

	return false
}

// SetActivitySourceState gets a reference to the given NullableString and assigns it to the ActivitySourceState field.
func (o *RoleMiningEntitlement) SetActivitySourceState(v string) {
	o.ActivitySourceState.Set(&v)
}
// SetActivitySourceStateNil sets the value for ActivitySourceState to be an explicit nil
func (o *RoleMiningEntitlement) SetActivitySourceStateNil() {
	o.ActivitySourceState.Set(nil)
}

// UnsetActivitySourceState ensures that no value is present for ActivitySourceState, not even an explicit nil
func (o *RoleMiningEntitlement) UnsetActivitySourceState() {
	o.ActivitySourceState.Unset()
}

// GetSourceUsagePercent returns the SourceUsagePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningEntitlement) GetSourceUsagePercent() float32 {
	if o == nil || isNil(o.SourceUsagePercent.Get()) {
		var ret float32
		return ret
	}
	return *o.SourceUsagePercent.Get()
}

// GetSourceUsagePercentOk returns a tuple with the SourceUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningEntitlement) GetSourceUsagePercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceUsagePercent.Get(), o.SourceUsagePercent.IsSet()
}

// HasSourceUsagePercent returns a boolean if a field has been set.
func (o *RoleMiningEntitlement) HasSourceUsagePercent() bool {
	if o != nil && o.SourceUsagePercent.IsSet() {
		return true
	}

	return false
}

// SetSourceUsagePercent gets a reference to the given NullableFloat32 and assigns it to the SourceUsagePercent field.
func (o *RoleMiningEntitlement) SetSourceUsagePercent(v float32) {
	o.SourceUsagePercent.Set(&v)
}
// SetSourceUsagePercentNil sets the value for SourceUsagePercent to be an explicit nil
func (o *RoleMiningEntitlement) SetSourceUsagePercentNil() {
	o.SourceUsagePercent.Set(nil)
}

// UnsetSourceUsagePercent ensures that no value is present for SourceUsagePercent, not even an explicit nil
func (o *RoleMiningEntitlement) UnsetSourceUsagePercent() {
	o.SourceUsagePercent.Unset()
}

func (o RoleMiningEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EntitlementRef) {
		toSerialize["entitlementRef"] = o.EntitlementRef
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !isNil(o.PopularityInOrg) {
		toSerialize["popularityInOrg"] = o.PopularityInOrg
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.ActivitySourceState.IsSet() {
		toSerialize["activitySourceState"] = o.ActivitySourceState.Get()
	}
	if o.SourceUsagePercent.IsSet() {
		toSerialize["sourceUsagePercent"] = o.SourceUsagePercent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningEntitlement := _RoleMiningEntitlement{}

	if err = json.Unmarshal(bytes, &varRoleMiningEntitlement); err == nil {
	*o = RoleMiningEntitlement(varRoleMiningEntitlement)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementRef")
		delete(additionalProperties, "name")
		delete(additionalProperties, "applicationName")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "popularity")
		delete(additionalProperties, "popularityInOrg")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "activitySourceState")
		delete(additionalProperties, "sourceUsagePercent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningEntitlement struct {
	value *RoleMiningEntitlement
	isSet bool
}

func (v NullableRoleMiningEntitlement) Get() *RoleMiningEntitlement {
	return v.value
}

func (v *NullableRoleMiningEntitlement) Set(val *RoleMiningEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningEntitlement(val *RoleMiningEntitlement) *NullableRoleMiningEntitlement {
	return &NullableRoleMiningEntitlement{value: val, isSet: true}
}

func (v NullableRoleMiningEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


