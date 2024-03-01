/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CertificationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificationDto{}

// CertificationDto struct for CertificationDto
type CertificationDto struct {
	CampaignRef CampaignReference `json:"campaignRef"`
	Phase CertificationPhase `json:"phase"`
	// The due date of the certification.
	Due time.Time `json:"due"`
	// The date the reviewer signed off on the certification.
	Signed time.Time `json:"signed"`
	Reviewer Reviewer `json:"reviewer"`
	Reassignment *Reassignment `json:"reassignment,omitempty"`
	// Indicates it the certification has any errors.
	HasErrors bool `json:"hasErrors"`
	// A message indicating what the error is.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Indicates if all certification decisions have been made.
	Completed bool `json:"completed"`
	// The number of approve/revoke/acknowledge decisions that have been made by the reviewer.
	DecisionsMade int32 `json:"decisionsMade"`
	// The total number of approve/revoke/acknowledge decisions for the certification.
	DecisionsTotal int32 `json:"decisionsTotal"`
	// The number of entities (identities, access profiles, roles, etc.) for which all decisions have been made and are complete.
	EntitiesCompleted int32 `json:"entitiesCompleted"`
	// The total number of entities (identities, access profiles, roles, etc.) in the certification, both complete and incomplete.
	EntitiesTotal int32 `json:"entitiesTotal"`
	AdditionalProperties map[string]interface{}
}

type _CertificationDto CertificationDto

// NewCertificationDto instantiates a new CertificationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationDto(campaignRef CampaignReference, phase CertificationPhase, due time.Time, signed time.Time, reviewer Reviewer, hasErrors bool, completed bool, decisionsMade int32, decisionsTotal int32, entitiesCompleted int32, entitiesTotal int32) *CertificationDto {
	this := CertificationDto{}
	this.CampaignRef = campaignRef
	this.Phase = phase
	this.Due = due
	this.Signed = signed
	this.Reviewer = reviewer
	this.HasErrors = hasErrors
	this.Completed = completed
	this.DecisionsMade = decisionsMade
	this.DecisionsTotal = decisionsTotal
	this.EntitiesCompleted = entitiesCompleted
	this.EntitiesTotal = entitiesTotal
	return &this
}

// NewCertificationDtoWithDefaults instantiates a new CertificationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationDtoWithDefaults() *CertificationDto {
	this := CertificationDto{}
	return &this
}

// GetCampaignRef returns the CampaignRef field value
func (o *CertificationDto) GetCampaignRef() CampaignReference {
	if o == nil {
		var ret CampaignReference
		return ret
	}

	return o.CampaignRef
}

// GetCampaignRefOk returns a tuple with the CampaignRef field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetCampaignRefOk() (*CampaignReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignRef, true
}

// SetCampaignRef sets field value
func (o *CertificationDto) SetCampaignRef(v CampaignReference) {
	o.CampaignRef = v
}

// GetPhase returns the Phase field value
func (o *CertificationDto) GetPhase() CertificationPhase {
	if o == nil {
		var ret CertificationPhase
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetPhaseOk() (*CertificationPhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *CertificationDto) SetPhase(v CertificationPhase) {
	o.Phase = v
}

// GetDue returns the Due field value
func (o *CertificationDto) GetDue() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Due
}

// GetDueOk returns a tuple with the Due field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetDueOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Due, true
}

// SetDue sets field value
func (o *CertificationDto) SetDue(v time.Time) {
	o.Due = v
}

// GetSigned returns the Signed field value
func (o *CertificationDto) GetSigned() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Signed
}

// GetSignedOk returns a tuple with the Signed field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetSignedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signed, true
}

// SetSigned sets field value
func (o *CertificationDto) SetSigned(v time.Time) {
	o.Signed = v
}

// GetReviewer returns the Reviewer field value
func (o *CertificationDto) GetReviewer() Reviewer {
	if o == nil {
		var ret Reviewer
		return ret
	}

	return o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetReviewerOk() (*Reviewer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reviewer, true
}

// SetReviewer sets field value
func (o *CertificationDto) SetReviewer(v Reviewer) {
	o.Reviewer = v
}

// GetReassignment returns the Reassignment field value if set, zero value otherwise.
func (o *CertificationDto) GetReassignment() Reassignment {
	if o == nil || isNil(o.Reassignment) {
		var ret Reassignment
		return ret
	}
	return *o.Reassignment
}

// GetReassignmentOk returns a tuple with the Reassignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetReassignmentOk() (*Reassignment, bool) {
	if o == nil || isNil(o.Reassignment) {
		return nil, false
	}
	return o.Reassignment, true
}

// HasReassignment returns a boolean if a field has been set.
func (o *CertificationDto) HasReassignment() bool {
	if o != nil && !isNil(o.Reassignment) {
		return true
	}

	return false
}

// SetReassignment gets a reference to the given Reassignment and assigns it to the Reassignment field.
func (o *CertificationDto) SetReassignment(v Reassignment) {
	o.Reassignment = &v
}

// GetHasErrors returns the HasErrors field value
func (o *CertificationDto) GetHasErrors() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasErrors
}

// GetHasErrorsOk returns a tuple with the HasErrors field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetHasErrorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasErrors, true
}

// SetHasErrors sets field value
func (o *CertificationDto) SetHasErrors(v bool) {
	o.HasErrors = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificationDto) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificationDto) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CertificationDto) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *CertificationDto) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *CertificationDto) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *CertificationDto) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetCompleted returns the Completed field value
func (o *CertificationDto) GetCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *CertificationDto) SetCompleted(v bool) {
	o.Completed = v
}

// GetDecisionsMade returns the DecisionsMade field value
func (o *CertificationDto) GetDecisionsMade() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DecisionsMade
}

// GetDecisionsMadeOk returns a tuple with the DecisionsMade field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetDecisionsMadeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecisionsMade, true
}

// SetDecisionsMade sets field value
func (o *CertificationDto) SetDecisionsMade(v int32) {
	o.DecisionsMade = v
}

// GetDecisionsTotal returns the DecisionsTotal field value
func (o *CertificationDto) GetDecisionsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DecisionsTotal
}

// GetDecisionsTotalOk returns a tuple with the DecisionsTotal field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetDecisionsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecisionsTotal, true
}

// SetDecisionsTotal sets field value
func (o *CertificationDto) SetDecisionsTotal(v int32) {
	o.DecisionsTotal = v
}

// GetEntitiesCompleted returns the EntitiesCompleted field value
func (o *CertificationDto) GetEntitiesCompleted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntitiesCompleted
}

// GetEntitiesCompletedOk returns a tuple with the EntitiesCompleted field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetEntitiesCompletedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitiesCompleted, true
}

// SetEntitiesCompleted sets field value
func (o *CertificationDto) SetEntitiesCompleted(v int32) {
	o.EntitiesCompleted = v
}

// GetEntitiesTotal returns the EntitiesTotal field value
func (o *CertificationDto) GetEntitiesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntitiesTotal
}

// GetEntitiesTotalOk returns a tuple with the EntitiesTotal field value
// and a boolean to check if the value has been set.
func (o *CertificationDto) GetEntitiesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitiesTotal, true
}

// SetEntitiesTotal sets field value
func (o *CertificationDto) SetEntitiesTotal(v int32) {
	o.EntitiesTotal = v
}

func (o CertificationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignRef"] = o.CampaignRef
	toSerialize["phase"] = o.Phase
	toSerialize["due"] = o.Due
	toSerialize["signed"] = o.Signed
	toSerialize["reviewer"] = o.Reviewer
	if !isNil(o.Reassignment) {
		toSerialize["reassignment"] = o.Reassignment
	}
	toSerialize["hasErrors"] = o.HasErrors
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	toSerialize["completed"] = o.Completed
	toSerialize["decisionsMade"] = o.DecisionsMade
	toSerialize["decisionsTotal"] = o.DecisionsTotal
	toSerialize["entitiesCompleted"] = o.EntitiesCompleted
	toSerialize["entitiesTotal"] = o.EntitiesTotal

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificationDto) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaignRef",
		"phase",
		"due",
		"signed",
		"reviewer",
		"hasErrors",
		"completed",
		"decisionsMade",
		"decisionsTotal",
		"entitiesCompleted",
		"entitiesTotal",
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

	varCertificationDto := _CertificationDto{}

	if err = json.Unmarshal(bytes, &varCertificationDto); err == nil {
	*o = CertificationDto(varCertificationDto)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "campaignRef")
		delete(additionalProperties, "phase")
		delete(additionalProperties, "due")
		delete(additionalProperties, "signed")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "reassignment")
		delete(additionalProperties, "hasErrors")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "decisionsMade")
		delete(additionalProperties, "decisionsTotal")
		delete(additionalProperties, "entitiesCompleted")
		delete(additionalProperties, "entitiesTotal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificationDto struct {
	value *CertificationDto
	isSet bool
}

func (v NullableCertificationDto) Get() *CertificationDto {
	return v.value
}

func (v *NullableCertificationDto) Set(val *CertificationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationDto(val *CertificationDto) *NullableCertificationDto {
	return &NullableCertificationDto{value: val, isSet: true}
}

func (v NullableCertificationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

