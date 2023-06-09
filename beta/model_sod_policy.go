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
)

// checks if the SodPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodPolicy{}

// SodPolicy struct for SodPolicy
type SodPolicy struct {
	// Policy id
	Id *string `json:"id,omitempty"`
	// Policy Business Name
	Name *string `json:"name,omitempty"`
	// The time when this SOD policy is created.
	Created *time.Time `json:"created,omitempty"`
	// The time when this SOD policy is modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Optional description of the SOD policy
	Description NullableString `json:"description,omitempty"`
	OwnerRef *BaseReferenceDto1 `json:"ownerRef,omitempty"`
	// Optional External Policy Reference
	ExternalPolicyReference NullableString `json:"externalPolicyReference,omitempty"`
	// Search query of the SOD policy
	PolicyQuery *string `json:"policyQuery,omitempty"`
	// Optional compensating controls(Mitigating Controls)
	CompensatingControls NullableString `json:"compensatingControls,omitempty"`
	// Optional correction advice
	CorrectionAdvice NullableString `json:"correctionAdvice,omitempty"`
	// whether the policy is enforced or not
	State *string `json:"state,omitempty"`
	// tags for this policy object
	Tags []string `json:"tags,omitempty"`
	// Policy's creator ID
	CreatorId *string `json:"creatorId,omitempty"`
	// Policy's modifier ID
	ModifierId NullableString `json:"modifierId,omitempty"`
	ViolationOwnerAssignmentConfig *ViolationOwnerAssignmentConfig `json:"violationOwnerAssignmentConfig,omitempty"`
	// defines whether a policy has been scheduled or not
	Scheduled *bool `json:"scheduled,omitempty"`
	// whether a policy is query based or conflicting access based
	Type *string `json:"type,omitempty"`
	ConflictingAccessCriteria *SodPolicyConflictingAccessCriteria `json:"conflictingAccessCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodPolicy SodPolicy

// NewSodPolicy instantiates a new SodPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodPolicy() *SodPolicy {
	this := SodPolicy{}
	var scheduled bool = false
	this.Scheduled = &scheduled
	var type_ string = "GENERAL"
	this.Type = &type_
	return &this
}

// NewSodPolicyWithDefaults instantiates a new SodPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodPolicyWithDefaults() *SodPolicy {
	this := SodPolicy{}
	var scheduled bool = false
	this.Scheduled = &scheduled
	var type_ string = "GENERAL"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SodPolicy) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SodPolicy) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SodPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SodPolicy) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SodPolicy) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SodPolicy) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SodPolicy) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SodPolicy) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SodPolicy) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SodPolicy) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SodPolicy) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SodPolicy) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodPolicy) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SodPolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SodPolicy) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SodPolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SodPolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise.
func (o *SodPolicy) GetOwnerRef() BaseReferenceDto1 {
	if o == nil || isNil(o.OwnerRef) {
		var ret BaseReferenceDto1
		return ret
	}
	return *o.OwnerRef
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetOwnerRefOk() (*BaseReferenceDto1, bool) {
	if o == nil || isNil(o.OwnerRef) {
		return nil, false
	}
	return o.OwnerRef, true
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *SodPolicy) HasOwnerRef() bool {
	if o != nil && !isNil(o.OwnerRef) {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given BaseReferenceDto1 and assigns it to the OwnerRef field.
func (o *SodPolicy) SetOwnerRef(v BaseReferenceDto1) {
	o.OwnerRef = &v
}

// GetExternalPolicyReference returns the ExternalPolicyReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodPolicy) GetExternalPolicyReference() string {
	if o == nil || isNil(o.ExternalPolicyReference.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalPolicyReference.Get()
}

// GetExternalPolicyReferenceOk returns a tuple with the ExternalPolicyReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodPolicy) GetExternalPolicyReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalPolicyReference.Get(), o.ExternalPolicyReference.IsSet()
}

// HasExternalPolicyReference returns a boolean if a field has been set.
func (o *SodPolicy) HasExternalPolicyReference() bool {
	if o != nil && o.ExternalPolicyReference.IsSet() {
		return true
	}

	return false
}

// SetExternalPolicyReference gets a reference to the given NullableString and assigns it to the ExternalPolicyReference field.
func (o *SodPolicy) SetExternalPolicyReference(v string) {
	o.ExternalPolicyReference.Set(&v)
}
// SetExternalPolicyReferenceNil sets the value for ExternalPolicyReference to be an explicit nil
func (o *SodPolicy) SetExternalPolicyReferenceNil() {
	o.ExternalPolicyReference.Set(nil)
}

// UnsetExternalPolicyReference ensures that no value is present for ExternalPolicyReference, not even an explicit nil
func (o *SodPolicy) UnsetExternalPolicyReference() {
	o.ExternalPolicyReference.Unset()
}

// GetPolicyQuery returns the PolicyQuery field value if set, zero value otherwise.
func (o *SodPolicy) GetPolicyQuery() string {
	if o == nil || isNil(o.PolicyQuery) {
		var ret string
		return ret
	}
	return *o.PolicyQuery
}

// GetPolicyQueryOk returns a tuple with the PolicyQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetPolicyQueryOk() (*string, bool) {
	if o == nil || isNil(o.PolicyQuery) {
		return nil, false
	}
	return o.PolicyQuery, true
}

// HasPolicyQuery returns a boolean if a field has been set.
func (o *SodPolicy) HasPolicyQuery() bool {
	if o != nil && !isNil(o.PolicyQuery) {
		return true
	}

	return false
}

// SetPolicyQuery gets a reference to the given string and assigns it to the PolicyQuery field.
func (o *SodPolicy) SetPolicyQuery(v string) {
	o.PolicyQuery = &v
}

// GetCompensatingControls returns the CompensatingControls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodPolicy) GetCompensatingControls() string {
	if o == nil || isNil(o.CompensatingControls.Get()) {
		var ret string
		return ret
	}
	return *o.CompensatingControls.Get()
}

// GetCompensatingControlsOk returns a tuple with the CompensatingControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodPolicy) GetCompensatingControlsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompensatingControls.Get(), o.CompensatingControls.IsSet()
}

// HasCompensatingControls returns a boolean if a field has been set.
func (o *SodPolicy) HasCompensatingControls() bool {
	if o != nil && o.CompensatingControls.IsSet() {
		return true
	}

	return false
}

// SetCompensatingControls gets a reference to the given NullableString and assigns it to the CompensatingControls field.
func (o *SodPolicy) SetCompensatingControls(v string) {
	o.CompensatingControls.Set(&v)
}
// SetCompensatingControlsNil sets the value for CompensatingControls to be an explicit nil
func (o *SodPolicy) SetCompensatingControlsNil() {
	o.CompensatingControls.Set(nil)
}

// UnsetCompensatingControls ensures that no value is present for CompensatingControls, not even an explicit nil
func (o *SodPolicy) UnsetCompensatingControls() {
	o.CompensatingControls.Unset()
}

// GetCorrectionAdvice returns the CorrectionAdvice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodPolicy) GetCorrectionAdvice() string {
	if o == nil || isNil(o.CorrectionAdvice.Get()) {
		var ret string
		return ret
	}
	return *o.CorrectionAdvice.Get()
}

// GetCorrectionAdviceOk returns a tuple with the CorrectionAdvice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodPolicy) GetCorrectionAdviceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrectionAdvice.Get(), o.CorrectionAdvice.IsSet()
}

// HasCorrectionAdvice returns a boolean if a field has been set.
func (o *SodPolicy) HasCorrectionAdvice() bool {
	if o != nil && o.CorrectionAdvice.IsSet() {
		return true
	}

	return false
}

// SetCorrectionAdvice gets a reference to the given NullableString and assigns it to the CorrectionAdvice field.
func (o *SodPolicy) SetCorrectionAdvice(v string) {
	o.CorrectionAdvice.Set(&v)
}
// SetCorrectionAdviceNil sets the value for CorrectionAdvice to be an explicit nil
func (o *SodPolicy) SetCorrectionAdviceNil() {
	o.CorrectionAdvice.Set(nil)
}

// UnsetCorrectionAdvice ensures that no value is present for CorrectionAdvice, not even an explicit nil
func (o *SodPolicy) UnsetCorrectionAdvice() {
	o.CorrectionAdvice.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SodPolicy) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SodPolicy) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SodPolicy) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SodPolicy) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SodPolicy) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SodPolicy) SetTags(v []string) {
	o.Tags = v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SodPolicy) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SodPolicy) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SodPolicy) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetModifierId returns the ModifierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SodPolicy) GetModifierId() string {
	if o == nil || isNil(o.ModifierId.Get()) {
		var ret string
		return ret
	}
	return *o.ModifierId.Get()
}

// GetModifierIdOk returns a tuple with the ModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SodPolicy) GetModifierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifierId.Get(), o.ModifierId.IsSet()
}

// HasModifierId returns a boolean if a field has been set.
func (o *SodPolicy) HasModifierId() bool {
	if o != nil && o.ModifierId.IsSet() {
		return true
	}

	return false
}

// SetModifierId gets a reference to the given NullableString and assigns it to the ModifierId field.
func (o *SodPolicy) SetModifierId(v string) {
	o.ModifierId.Set(&v)
}
// SetModifierIdNil sets the value for ModifierId to be an explicit nil
func (o *SodPolicy) SetModifierIdNil() {
	o.ModifierId.Set(nil)
}

// UnsetModifierId ensures that no value is present for ModifierId, not even an explicit nil
func (o *SodPolicy) UnsetModifierId() {
	o.ModifierId.Unset()
}

// GetViolationOwnerAssignmentConfig returns the ViolationOwnerAssignmentConfig field value if set, zero value otherwise.
func (o *SodPolicy) GetViolationOwnerAssignmentConfig() ViolationOwnerAssignmentConfig {
	if o == nil || isNil(o.ViolationOwnerAssignmentConfig) {
		var ret ViolationOwnerAssignmentConfig
		return ret
	}
	return *o.ViolationOwnerAssignmentConfig
}

// GetViolationOwnerAssignmentConfigOk returns a tuple with the ViolationOwnerAssignmentConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetViolationOwnerAssignmentConfigOk() (*ViolationOwnerAssignmentConfig, bool) {
	if o == nil || isNil(o.ViolationOwnerAssignmentConfig) {
		return nil, false
	}
	return o.ViolationOwnerAssignmentConfig, true
}

// HasViolationOwnerAssignmentConfig returns a boolean if a field has been set.
func (o *SodPolicy) HasViolationOwnerAssignmentConfig() bool {
	if o != nil && !isNil(o.ViolationOwnerAssignmentConfig) {
		return true
	}

	return false
}

// SetViolationOwnerAssignmentConfig gets a reference to the given ViolationOwnerAssignmentConfig and assigns it to the ViolationOwnerAssignmentConfig field.
func (o *SodPolicy) SetViolationOwnerAssignmentConfig(v ViolationOwnerAssignmentConfig) {
	o.ViolationOwnerAssignmentConfig = &v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *SodPolicy) GetScheduled() bool {
	if o == nil || isNil(o.Scheduled) {
		var ret bool
		return ret
	}
	return *o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetScheduledOk() (*bool, bool) {
	if o == nil || isNil(o.Scheduled) {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *SodPolicy) HasScheduled() bool {
	if o != nil && !isNil(o.Scheduled) {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given bool and assigns it to the Scheduled field.
func (o *SodPolicy) SetScheduled(v bool) {
	o.Scheduled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SodPolicy) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SodPolicy) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SodPolicy) SetType(v string) {
	o.Type = &v
}

// GetConflictingAccessCriteria returns the ConflictingAccessCriteria field value if set, zero value otherwise.
func (o *SodPolicy) GetConflictingAccessCriteria() SodPolicyConflictingAccessCriteria {
	if o == nil || isNil(o.ConflictingAccessCriteria) {
		var ret SodPolicyConflictingAccessCriteria
		return ret
	}
	return *o.ConflictingAccessCriteria
}

// GetConflictingAccessCriteriaOk returns a tuple with the ConflictingAccessCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicy) GetConflictingAccessCriteriaOk() (*SodPolicyConflictingAccessCriteria, bool) {
	if o == nil || isNil(o.ConflictingAccessCriteria) {
		return nil, false
	}
	return o.ConflictingAccessCriteria, true
}

// HasConflictingAccessCriteria returns a boolean if a field has been set.
func (o *SodPolicy) HasConflictingAccessCriteria() bool {
	if o != nil && !isNil(o.ConflictingAccessCriteria) {
		return true
	}

	return false
}

// SetConflictingAccessCriteria gets a reference to the given SodPolicyConflictingAccessCriteria and assigns it to the ConflictingAccessCriteria field.
func (o *SodPolicy) SetConflictingAccessCriteria(v SodPolicyConflictingAccessCriteria) {
	o.ConflictingAccessCriteria = &v
}

func (o SodPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	// skip: created is readOnly
	// skip: modified is readOnly
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.OwnerRef) {
		toSerialize["ownerRef"] = o.OwnerRef
	}
	if o.ExternalPolicyReference.IsSet() {
		toSerialize["externalPolicyReference"] = o.ExternalPolicyReference.Get()
	}
	if !isNil(o.PolicyQuery) {
		toSerialize["policyQuery"] = o.PolicyQuery
	}
	if o.CompensatingControls.IsSet() {
		toSerialize["compensatingControls"] = o.CompensatingControls.Get()
	}
	if o.CorrectionAdvice.IsSet() {
		toSerialize["correctionAdvice"] = o.CorrectionAdvice.Get()
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	// skip: creatorId is readOnly
	if o.ModifierId.IsSet() {
		toSerialize["modifierId"] = o.ModifierId.Get()
	}
	if !isNil(o.ViolationOwnerAssignmentConfig) {
		toSerialize["violationOwnerAssignmentConfig"] = o.ViolationOwnerAssignmentConfig
	}
	if !isNil(o.Scheduled) {
		toSerialize["scheduled"] = o.Scheduled
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ConflictingAccessCriteria) {
		toSerialize["conflictingAccessCriteria"] = o.ConflictingAccessCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varSodPolicy := _SodPolicy{}

	if err = json.Unmarshal(bytes, &varSodPolicy); err == nil {
		*o = SodPolicy(varSodPolicy)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ownerRef")
		delete(additionalProperties, "externalPolicyReference")
		delete(additionalProperties, "policyQuery")
		delete(additionalProperties, "compensatingControls")
		delete(additionalProperties, "correctionAdvice")
		delete(additionalProperties, "state")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "creatorId")
		delete(additionalProperties, "modifierId")
		delete(additionalProperties, "violationOwnerAssignmentConfig")
		delete(additionalProperties, "scheduled")
		delete(additionalProperties, "type")
		delete(additionalProperties, "conflictingAccessCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodPolicy struct {
	value *SodPolicy
	isSet bool
}

func (v NullableSodPolicy) Get() *SodPolicy {
	return v.value
}

func (v *NullableSodPolicy) Set(val *SodPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSodPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSodPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodPolicy(val *SodPolicy) *NullableSodPolicy {
	return &NullableSodPolicy{value: val, isSet: true}
}

func (v NullableSodPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


