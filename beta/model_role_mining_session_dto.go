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

// RoleMiningSessionDto struct for RoleMiningSessionDto
type RoleMiningSessionDto struct {
	Scope *RoleMiningSessionScope `json:"scope,omitempty"`
	// The prune threshold to be used or null to calculate prescribedPruneThreshold
	PruneThreshold *int32 `json:"pruneThreshold,omitempty"`
	// The calculated prescribedPruneThreshold
	PrescribedPruneThreshold *int32 `json:"prescribedPruneThreshold,omitempty"`
	// Minimum number of identities in a potential role
	MinNumIdentitiesInPotentialRole *int32 `json:"minNumIdentitiesInPotentialRole,omitempty"`
	// Number of potential roles
	PotentialRoleCount *int32 `json:"potentialRoleCount,omitempty"`
	// Number of potential roles ready
	PotentialRolesReadyCount *int32 `json:"potentialRolesReadyCount,omitempty"`
	Status *RoleMiningSessionStatus `json:"status,omitempty"`
	Type *RoleMiningRoleType `json:"type,omitempty"`
	// The id of the user who will receive an email about the role mining session
	EmailRecipientId *string `json:"emailRecipientId,omitempty"`
	CreatedBy *EntityCreatedByDTO `json:"createdBy,omitempty"`
	// Number of identities in the population which meet the search criteria or identity list provided
	IdentityCount *int32 `json:"identityCount,omitempty"`
	// The session's saved status
	Saved *bool `json:"saved,omitempty"`
	// The session's saved name
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningSessionDto RoleMiningSessionDto

// NewRoleMiningSessionDto instantiates a new RoleMiningSessionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningSessionDto() *RoleMiningSessionDto {
	this := RoleMiningSessionDto{}
	return &this
}

// NewRoleMiningSessionDtoWithDefaults instantiates a new RoleMiningSessionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningSessionDtoWithDefaults() *RoleMiningSessionDto {
	this := RoleMiningSessionDto{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetScope() RoleMiningSessionScope {
	if o == nil || isNil(o.Scope) {
		var ret RoleMiningSessionScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetScopeOk() (*RoleMiningSessionScope, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given RoleMiningSessionScope and assigns it to the Scope field.
func (o *RoleMiningSessionDto) SetScope(v RoleMiningSessionScope) {
	o.Scope = &v
}

// GetPruneThreshold returns the PruneThreshold field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetPruneThreshold() int32 {
	if o == nil || isNil(o.PruneThreshold) {
		var ret int32
		return ret
	}
	return *o.PruneThreshold
}

// GetPruneThresholdOk returns a tuple with the PruneThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetPruneThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.PruneThreshold) {
		return nil, false
	}
	return o.PruneThreshold, true
}

// HasPruneThreshold returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasPruneThreshold() bool {
	if o != nil && !isNil(o.PruneThreshold) {
		return true
	}

	return false
}

// SetPruneThreshold gets a reference to the given int32 and assigns it to the PruneThreshold field.
func (o *RoleMiningSessionDto) SetPruneThreshold(v int32) {
	o.PruneThreshold = &v
}

// GetPrescribedPruneThreshold returns the PrescribedPruneThreshold field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetPrescribedPruneThreshold() int32 {
	if o == nil || isNil(o.PrescribedPruneThreshold) {
		var ret int32
		return ret
	}
	return *o.PrescribedPruneThreshold
}

// GetPrescribedPruneThresholdOk returns a tuple with the PrescribedPruneThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetPrescribedPruneThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.PrescribedPruneThreshold) {
		return nil, false
	}
	return o.PrescribedPruneThreshold, true
}

// HasPrescribedPruneThreshold returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasPrescribedPruneThreshold() bool {
	if o != nil && !isNil(o.PrescribedPruneThreshold) {
		return true
	}

	return false
}

// SetPrescribedPruneThreshold gets a reference to the given int32 and assigns it to the PrescribedPruneThreshold field.
func (o *RoleMiningSessionDto) SetPrescribedPruneThreshold(v int32) {
	o.PrescribedPruneThreshold = &v
}

// GetMinNumIdentitiesInPotentialRole returns the MinNumIdentitiesInPotentialRole field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetMinNumIdentitiesInPotentialRole() int32 {
	if o == nil || isNil(o.MinNumIdentitiesInPotentialRole) {
		var ret int32
		return ret
	}
	return *o.MinNumIdentitiesInPotentialRole
}

// GetMinNumIdentitiesInPotentialRoleOk returns a tuple with the MinNumIdentitiesInPotentialRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetMinNumIdentitiesInPotentialRoleOk() (*int32, bool) {
	if o == nil || isNil(o.MinNumIdentitiesInPotentialRole) {
		return nil, false
	}
	return o.MinNumIdentitiesInPotentialRole, true
}

// HasMinNumIdentitiesInPotentialRole returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasMinNumIdentitiesInPotentialRole() bool {
	if o != nil && !isNil(o.MinNumIdentitiesInPotentialRole) {
		return true
	}

	return false
}

// SetMinNumIdentitiesInPotentialRole gets a reference to the given int32 and assigns it to the MinNumIdentitiesInPotentialRole field.
func (o *RoleMiningSessionDto) SetMinNumIdentitiesInPotentialRole(v int32) {
	o.MinNumIdentitiesInPotentialRole = &v
}

// GetPotentialRoleCount returns the PotentialRoleCount field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetPotentialRoleCount() int32 {
	if o == nil || isNil(o.PotentialRoleCount) {
		var ret int32
		return ret
	}
	return *o.PotentialRoleCount
}

// GetPotentialRoleCountOk returns a tuple with the PotentialRoleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetPotentialRoleCountOk() (*int32, bool) {
	if o == nil || isNil(o.PotentialRoleCount) {
		return nil, false
	}
	return o.PotentialRoleCount, true
}

// HasPotentialRoleCount returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasPotentialRoleCount() bool {
	if o != nil && !isNil(o.PotentialRoleCount) {
		return true
	}

	return false
}

// SetPotentialRoleCount gets a reference to the given int32 and assigns it to the PotentialRoleCount field.
func (o *RoleMiningSessionDto) SetPotentialRoleCount(v int32) {
	o.PotentialRoleCount = &v
}

// GetPotentialRolesReadyCount returns the PotentialRolesReadyCount field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetPotentialRolesReadyCount() int32 {
	if o == nil || isNil(o.PotentialRolesReadyCount) {
		var ret int32
		return ret
	}
	return *o.PotentialRolesReadyCount
}

// GetPotentialRolesReadyCountOk returns a tuple with the PotentialRolesReadyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetPotentialRolesReadyCountOk() (*int32, bool) {
	if o == nil || isNil(o.PotentialRolesReadyCount) {
		return nil, false
	}
	return o.PotentialRolesReadyCount, true
}

// HasPotentialRolesReadyCount returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasPotentialRolesReadyCount() bool {
	if o != nil && !isNil(o.PotentialRolesReadyCount) {
		return true
	}

	return false
}

// SetPotentialRolesReadyCount gets a reference to the given int32 and assigns it to the PotentialRolesReadyCount field.
func (o *RoleMiningSessionDto) SetPotentialRolesReadyCount(v int32) {
	o.PotentialRolesReadyCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetStatus() RoleMiningSessionStatus {
	if o == nil || isNil(o.Status) {
		var ret RoleMiningSessionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetStatusOk() (*RoleMiningSessionStatus, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RoleMiningSessionStatus and assigns it to the Status field.
func (o *RoleMiningSessionDto) SetStatus(v RoleMiningSessionStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetType() RoleMiningRoleType {
	if o == nil || isNil(o.Type) {
		var ret RoleMiningRoleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetTypeOk() (*RoleMiningRoleType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RoleMiningRoleType and assigns it to the Type field.
func (o *RoleMiningSessionDto) SetType(v RoleMiningRoleType) {
	o.Type = &v
}

// GetEmailRecipientId returns the EmailRecipientId field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetEmailRecipientId() string {
	if o == nil || isNil(o.EmailRecipientId) {
		var ret string
		return ret
	}
	return *o.EmailRecipientId
}

// GetEmailRecipientIdOk returns a tuple with the EmailRecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetEmailRecipientIdOk() (*string, bool) {
	if o == nil || isNil(o.EmailRecipientId) {
		return nil, false
	}
	return o.EmailRecipientId, true
}

// HasEmailRecipientId returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasEmailRecipientId() bool {
	if o != nil && !isNil(o.EmailRecipientId) {
		return true
	}

	return false
}

// SetEmailRecipientId gets a reference to the given string and assigns it to the EmailRecipientId field.
func (o *RoleMiningSessionDto) SetEmailRecipientId(v string) {
	o.EmailRecipientId = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetCreatedBy() EntityCreatedByDTO {
	if o == nil || isNil(o.CreatedBy) {
		var ret EntityCreatedByDTO
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetCreatedByOk() (*EntityCreatedByDTO, bool) {
	if o == nil || isNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EntityCreatedByDTO and assigns it to the CreatedBy field.
func (o *RoleMiningSessionDto) SetCreatedBy(v EntityCreatedByDTO) {
	o.CreatedBy = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *RoleMiningSessionDto) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetSaved returns the Saved field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetSaved() bool {
	if o == nil || isNil(o.Saved) {
		var ret bool
		return ret
	}
	return *o.Saved
}

// GetSavedOk returns a tuple with the Saved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetSavedOk() (*bool, bool) {
	if o == nil || isNil(o.Saved) {
		return nil, false
	}
	return o.Saved, true
}

// HasSaved returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasSaved() bool {
	if o != nil && !isNil(o.Saved) {
		return true
	}

	return false
}

// SetSaved gets a reference to the given bool and assigns it to the Saved field.
func (o *RoleMiningSessionDto) SetSaved(v bool) {
	o.Saved = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleMiningSessionDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningSessionDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleMiningSessionDto) SetName(v string) {
	o.Name = &v
}

func (o RoleMiningSessionDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.PruneThreshold) {
		toSerialize["pruneThreshold"] = o.PruneThreshold
	}
	if !isNil(o.PrescribedPruneThreshold) {
		toSerialize["prescribedPruneThreshold"] = o.PrescribedPruneThreshold
	}
	if !isNil(o.MinNumIdentitiesInPotentialRole) {
		toSerialize["minNumIdentitiesInPotentialRole"] = o.MinNumIdentitiesInPotentialRole
	}
	if !isNil(o.PotentialRoleCount) {
		toSerialize["potentialRoleCount"] = o.PotentialRoleCount
	}
	if !isNil(o.PotentialRolesReadyCount) {
		toSerialize["potentialRolesReadyCount"] = o.PotentialRolesReadyCount
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.EmailRecipientId) {
		toSerialize["emailRecipientId"] = o.EmailRecipientId
	}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.Saved) {
		toSerialize["saved"] = o.Saved
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleMiningSessionDto) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningSessionDto := _RoleMiningSessionDto{}

	if err = json.Unmarshal(bytes, &varRoleMiningSessionDto); err == nil {
		*o = RoleMiningSessionDto(varRoleMiningSessionDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "scope")
		delete(additionalProperties, "pruneThreshold")
		delete(additionalProperties, "prescribedPruneThreshold")
		delete(additionalProperties, "minNumIdentitiesInPotentialRole")
		delete(additionalProperties, "potentialRoleCount")
		delete(additionalProperties, "potentialRolesReadyCount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "emailRecipientId")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "saved")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningSessionDto struct {
	value *RoleMiningSessionDto
	isSet bool
}

func (v NullableRoleMiningSessionDto) Get() *RoleMiningSessionDto {
	return v.value
}

func (v *NullableRoleMiningSessionDto) Set(val *RoleMiningSessionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningSessionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningSessionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningSessionDto(val *RoleMiningSessionDto) *NullableRoleMiningSessionDto {
	return &NullableRoleMiningSessionDto{value: val, isSet: true}
}

func (v NullableRoleMiningSessionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningSessionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

