/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the RoleMiningPotentialRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRole{}

// RoleMiningPotentialRole struct for RoleMiningPotentialRole
type RoleMiningPotentialRole struct {
	CreatedBy *RoleMiningSessionResponseCreatedBy `json:"createdBy,omitempty"`
	// The density of a potential role.
	Density *int32 `json:"density,omitempty"`
	// The description of a potential role.
	Description NullableString `json:"description,omitempty"`
	// The number of entitlements in a potential role.
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// The list of entitlement ids to be excluded.
	ExcludedEntitlements []string `json:"excludedEntitlements,omitempty"`
	// The freshness of a potential role.
	Freshness *int32 `json:"freshness,omitempty"`
	// The number of identities in a potential role.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	// Identity attribute distribution.
	IdentityDistribution []RoleMiningIdentityDistribution `json:"identityDistribution,omitempty"`
	// The list of ids in a potential role.
	IdentityIds []string `json:"identityIds,omitempty"`
	// Name of the potential role.
	Name *string `json:"name,omitempty"`
	ProvisionState *RoleMiningPotentialRoleProvisionState `json:"provisionState,omitempty"`
	// The quality of a potential role.
	Quality *int32 `json:"quality,omitempty"`
	// The roleId of a potential role.
	RoleId NullableString `json:"roleId,omitempty"`
	// The potential role's saved status.
	Saved *bool `json:"saved,omitempty"`
	Session *RoleMiningSessionParametersDto `json:"session,omitempty"`
	Type *RoleMiningRoleType `json:"type,omitempty"`
	// Id of the potential role
	Id *string `json:"id,omitempty"`
	// The date-time when this potential role was created.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// The date-time when this potential role was modified.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRole RoleMiningPotentialRole

// NewRoleMiningPotentialRole instantiates a new RoleMiningPotentialRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRole() *RoleMiningPotentialRole {
	this := RoleMiningPotentialRole{}
	return &this
}

// NewRoleMiningPotentialRoleWithDefaults instantiates a new RoleMiningPotentialRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleWithDefaults() *RoleMiningPotentialRole {
	this := RoleMiningPotentialRole{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetCreatedBy() RoleMiningSessionResponseCreatedBy {
	if o == nil || isNil(o.CreatedBy) {
		var ret RoleMiningSessionResponseCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetCreatedByOk() (*RoleMiningSessionResponseCreatedBy, bool) {
	if o == nil || isNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given RoleMiningSessionResponseCreatedBy and assigns it to the CreatedBy field.
func (o *RoleMiningPotentialRole) SetCreatedBy(v RoleMiningSessionResponseCreatedBy) {
	o.CreatedBy = &v
}

// GetDensity returns the Density field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetDensity() int32 {
	if o == nil || isNil(o.Density) {
		var ret int32
		return ret
	}
	return *o.Density
}

// GetDensityOk returns a tuple with the Density field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetDensityOk() (*int32, bool) {
	if o == nil || isNil(o.Density) {
		return nil, false
	}
	return o.Density, true
}

// HasDensity returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasDensity() bool {
	if o != nil && !isNil(o.Density) {
		return true
	}

	return false
}

// SetDensity gets a reference to the given int32 and assigns it to the Density field.
func (o *RoleMiningPotentialRole) SetDensity(v int32) {
	o.Density = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningPotentialRole) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningPotentialRole) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RoleMiningPotentialRole) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RoleMiningPotentialRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RoleMiningPotentialRole) UnsetDescription() {
	o.Description.Unset()
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *RoleMiningPotentialRole) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetExcludedEntitlements returns the ExcludedEntitlements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningPotentialRole) GetExcludedEntitlements() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludedEntitlements
}

// GetExcludedEntitlementsOk returns a tuple with the ExcludedEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningPotentialRole) GetExcludedEntitlementsOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedEntitlements) {
		return nil, false
	}
	return o.ExcludedEntitlements, true
}

// HasExcludedEntitlements returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasExcludedEntitlements() bool {
	if o != nil && isNil(o.ExcludedEntitlements) {
		return true
	}

	return false
}

// SetExcludedEntitlements gets a reference to the given []string and assigns it to the ExcludedEntitlements field.
func (o *RoleMiningPotentialRole) SetExcludedEntitlements(v []string) {
	o.ExcludedEntitlements = v
}

// GetFreshness returns the Freshness field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetFreshness() int32 {
	if o == nil || isNil(o.Freshness) {
		var ret int32
		return ret
	}
	return *o.Freshness
}

// GetFreshnessOk returns a tuple with the Freshness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetFreshnessOk() (*int32, bool) {
	if o == nil || isNil(o.Freshness) {
		return nil, false
	}
	return o.Freshness, true
}

// HasFreshness returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasFreshness() bool {
	if o != nil && !isNil(o.Freshness) {
		return true
	}

	return false
}

// SetFreshness gets a reference to the given int32 and assigns it to the Freshness field.
func (o *RoleMiningPotentialRole) SetFreshness(v int32) {
	o.Freshness = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *RoleMiningPotentialRole) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetIdentityDistribution returns the IdentityDistribution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningPotentialRole) GetIdentityDistribution() []RoleMiningIdentityDistribution {
	if o == nil {
		var ret []RoleMiningIdentityDistribution
		return ret
	}
	return o.IdentityDistribution
}

// GetIdentityDistributionOk returns a tuple with the IdentityDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningPotentialRole) GetIdentityDistributionOk() ([]RoleMiningIdentityDistribution, bool) {
	if o == nil || isNil(o.IdentityDistribution) {
		return nil, false
	}
	return o.IdentityDistribution, true
}

// HasIdentityDistribution returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasIdentityDistribution() bool {
	if o != nil && isNil(o.IdentityDistribution) {
		return true
	}

	return false
}

// SetIdentityDistribution gets a reference to the given []RoleMiningIdentityDistribution and assigns it to the IdentityDistribution field.
func (o *RoleMiningPotentialRole) SetIdentityDistribution(v []RoleMiningIdentityDistribution) {
	o.IdentityDistribution = v
}

// GetIdentityIds returns the IdentityIds field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetIdentityIds() []string {
	if o == nil || isNil(o.IdentityIds) {
		var ret []string
		return ret
	}
	return o.IdentityIds
}

// GetIdentityIdsOk returns a tuple with the IdentityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdentityIdsOk() ([]string, bool) {
	if o == nil || isNil(o.IdentityIds) {
		return nil, false
	}
	return o.IdentityIds, true
}

// HasIdentityIds returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasIdentityIds() bool {
	if o != nil && !isNil(o.IdentityIds) {
		return true
	}

	return false
}

// SetIdentityIds gets a reference to the given []string and assigns it to the IdentityIds field.
func (o *RoleMiningPotentialRole) SetIdentityIds(v []string) {
	o.IdentityIds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleMiningPotentialRole) SetName(v string) {
	o.Name = &v
}

// GetProvisionState returns the ProvisionState field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetProvisionState() RoleMiningPotentialRoleProvisionState {
	if o == nil || isNil(o.ProvisionState) {
		var ret RoleMiningPotentialRoleProvisionState
		return ret
	}
	return *o.ProvisionState
}

// GetProvisionStateOk returns a tuple with the ProvisionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetProvisionStateOk() (*RoleMiningPotentialRoleProvisionState, bool) {
	if o == nil || isNil(o.ProvisionState) {
		return nil, false
	}
	return o.ProvisionState, true
}

// HasProvisionState returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasProvisionState() bool {
	if o != nil && !isNil(o.ProvisionState) {
		return true
	}

	return false
}

// SetProvisionState gets a reference to the given RoleMiningPotentialRoleProvisionState and assigns it to the ProvisionState field.
func (o *RoleMiningPotentialRole) SetProvisionState(v RoleMiningPotentialRoleProvisionState) {
	o.ProvisionState = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetQuality() int32 {
	if o == nil || isNil(o.Quality) {
		var ret int32
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetQualityOk() (*int32, bool) {
	if o == nil || isNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given int32 and assigns it to the Quality field.
func (o *RoleMiningPotentialRole) SetQuality(v int32) {
	o.Quality = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningPotentialRole) GetRoleId() string {
	if o == nil || isNil(o.RoleId.Get()) {
		var ret string
		return ret
	}
	return *o.RoleId.Get()
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningPotentialRole) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleId.Get(), o.RoleId.IsSet()
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasRoleId() bool {
	if o != nil && o.RoleId.IsSet() {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given NullableString and assigns it to the RoleId field.
func (o *RoleMiningPotentialRole) SetRoleId(v string) {
	o.RoleId.Set(&v)
}
// SetRoleIdNil sets the value for RoleId to be an explicit nil
func (o *RoleMiningPotentialRole) SetRoleIdNil() {
	o.RoleId.Set(nil)
}

// UnsetRoleId ensures that no value is present for RoleId, not even an explicit nil
func (o *RoleMiningPotentialRole) UnsetRoleId() {
	o.RoleId.Unset()
}

// GetSaved returns the Saved field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetSaved() bool {
	if o == nil || isNil(o.Saved) {
		var ret bool
		return ret
	}
	return *o.Saved
}

// GetSavedOk returns a tuple with the Saved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetSavedOk() (*bool, bool) {
	if o == nil || isNil(o.Saved) {
		return nil, false
	}
	return o.Saved, true
}

// HasSaved returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasSaved() bool {
	if o != nil && !isNil(o.Saved) {
		return true
	}

	return false
}

// SetSaved gets a reference to the given bool and assigns it to the Saved field.
func (o *RoleMiningPotentialRole) SetSaved(v bool) {
	o.Saved = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetSession() RoleMiningSessionParametersDto {
	if o == nil || isNil(o.Session) {
		var ret RoleMiningSessionParametersDto
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetSessionOk() (*RoleMiningSessionParametersDto, bool) {
	if o == nil || isNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasSession() bool {
	if o != nil && !isNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given RoleMiningSessionParametersDto and assigns it to the Session field.
func (o *RoleMiningPotentialRole) SetSession(v RoleMiningSessionParametersDto) {
	o.Session = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetType() RoleMiningRoleType {
	if o == nil || isNil(o.Type) {
		var ret RoleMiningRoleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetTypeOk() (*RoleMiningRoleType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RoleMiningRoleType and assigns it to the Type field.
func (o *RoleMiningPotentialRole) SetType(v RoleMiningRoleType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMiningPotentialRole) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetCreatedDate() time.Time {
	if o == nil || isNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *RoleMiningPotentialRole) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetModifiedDate() time.Time {
	if o == nil || isNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasModifiedDate() bool {
	if o != nil && !isNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *RoleMiningPotentialRole) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

func (o RoleMiningPotentialRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.Density) {
		toSerialize["density"] = o.Density
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if o.ExcludedEntitlements != nil {
		toSerialize["excludedEntitlements"] = o.ExcludedEntitlements
	}
	if !isNil(o.Freshness) {
		toSerialize["freshness"] = o.Freshness
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if o.IdentityDistribution != nil {
		toSerialize["identityDistribution"] = o.IdentityDistribution
	}
	if !isNil(o.IdentityIds) {
		toSerialize["identityIds"] = o.IdentityIds
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ProvisionState) {
		toSerialize["provisionState"] = o.ProvisionState
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.RoleId.IsSet() {
		toSerialize["roleId"] = o.RoleId.Get()
	}
	if !isNil(o.Saved) {
		toSerialize["saved"] = o.Saved
	}
	if !isNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !isNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRole) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRole := _RoleMiningPotentialRole{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRole); err == nil {
	*o = RoleMiningPotentialRole(varRoleMiningPotentialRole)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "density")
		delete(additionalProperties, "description")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "excludedEntitlements")
		delete(additionalProperties, "freshness")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "identityDistribution")
		delete(additionalProperties, "identityIds")
		delete(additionalProperties, "name")
		delete(additionalProperties, "provisionState")
		delete(additionalProperties, "quality")
		delete(additionalProperties, "roleId")
		delete(additionalProperties, "saved")
		delete(additionalProperties, "session")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "modifiedDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRole struct {
	value *RoleMiningPotentialRole
	isSet bool
}

func (v NullableRoleMiningPotentialRole) Get() *RoleMiningPotentialRole {
	return v.value
}

func (v *NullableRoleMiningPotentialRole) Set(val *RoleMiningPotentialRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRole(val *RoleMiningPotentialRole) *NullableRoleMiningPotentialRole {
	return &NullableRoleMiningPotentialRole{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


