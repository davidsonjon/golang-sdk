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

// checks if the Certification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certification{}

// Certification struct for Certification
type Certification struct {
	// id of the certification
	Id *string `json:"id,omitempty"`
	// name of the certification
	Name *string `json:"name,omitempty"`
	Campaign *CampaignReference `json:"campaign,omitempty"`
	// Have all decisions been made?
	Completed *bool `json:"completed,omitempty"`
	// The number of identities for whom all decisions have been made and are complete.
	IdentitiesCompleted *int32 `json:"identitiesCompleted,omitempty"`
	// The total number of identities in the Certification, both complete and incomplete.
	IdentitiesTotal *int32 `json:"identitiesTotal,omitempty"`
	// created date
	Created *time.Time `json:"created,omitempty"`
	// modified date
	Modified *time.Time `json:"modified,omitempty"`
	// The number of approve/revoke/acknowledge decisions that have been made.
	DecisionsMade *int32 `json:"decisionsMade,omitempty"`
	// The total number of approve/revoke/acknowledge decisions.
	DecisionsTotal *int32 `json:"decisionsTotal,omitempty"`
	// The due date of the certification.
	Due *time.Time `json:"due,omitempty"`
	// The date the reviewer signed off on the Certification.
	Signed NullableTime `json:"signed,omitempty"`
	Reviewer *Reviewer `json:"reviewer,omitempty"`
	Reassignment NullableReassignment `json:"reassignment,omitempty"`
	// Identifies if the certification has an error
	HasErrors *bool `json:"hasErrors,omitempty"`
	// Description of the certification error
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	Phase *CertificationPhase `json:"phase,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Certification Certification

// NewCertification instantiates a new Certification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertification() *Certification {
	this := Certification{}
	return &this
}

// NewCertificationWithDefaults instantiates a new Certification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationWithDefaults() *Certification {
	this := Certification{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Certification) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Certification) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Certification) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Certification) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Certification) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Certification) SetName(v string) {
	o.Name = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *Certification) GetCampaign() CampaignReference {
	if o == nil || isNil(o.Campaign) {
		var ret CampaignReference
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetCampaignOk() (*CampaignReference, bool) {
	if o == nil || isNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *Certification) HasCampaign() bool {
	if o != nil && !isNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given CampaignReference and assigns it to the Campaign field.
func (o *Certification) SetCampaign(v CampaignReference) {
	o.Campaign = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *Certification) GetCompleted() bool {
	if o == nil || isNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *Certification) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *Certification) SetCompleted(v bool) {
	o.Completed = &v
}

// GetIdentitiesCompleted returns the IdentitiesCompleted field value if set, zero value otherwise.
func (o *Certification) GetIdentitiesCompleted() int32 {
	if o == nil || isNil(o.IdentitiesCompleted) {
		var ret int32
		return ret
	}
	return *o.IdentitiesCompleted
}

// GetIdentitiesCompletedOk returns a tuple with the IdentitiesCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetIdentitiesCompletedOk() (*int32, bool) {
	if o == nil || isNil(o.IdentitiesCompleted) {
		return nil, false
	}
	return o.IdentitiesCompleted, true
}

// HasIdentitiesCompleted returns a boolean if a field has been set.
func (o *Certification) HasIdentitiesCompleted() bool {
	if o != nil && !isNil(o.IdentitiesCompleted) {
		return true
	}

	return false
}

// SetIdentitiesCompleted gets a reference to the given int32 and assigns it to the IdentitiesCompleted field.
func (o *Certification) SetIdentitiesCompleted(v int32) {
	o.IdentitiesCompleted = &v
}

// GetIdentitiesTotal returns the IdentitiesTotal field value if set, zero value otherwise.
func (o *Certification) GetIdentitiesTotal() int32 {
	if o == nil || isNil(o.IdentitiesTotal) {
		var ret int32
		return ret
	}
	return *o.IdentitiesTotal
}

// GetIdentitiesTotalOk returns a tuple with the IdentitiesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetIdentitiesTotalOk() (*int32, bool) {
	if o == nil || isNil(o.IdentitiesTotal) {
		return nil, false
	}
	return o.IdentitiesTotal, true
}

// HasIdentitiesTotal returns a boolean if a field has been set.
func (o *Certification) HasIdentitiesTotal() bool {
	if o != nil && !isNil(o.IdentitiesTotal) {
		return true
	}

	return false
}

// SetIdentitiesTotal gets a reference to the given int32 and assigns it to the IdentitiesTotal field.
func (o *Certification) SetIdentitiesTotal(v int32) {
	o.IdentitiesTotal = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Certification) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Certification) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Certification) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Certification) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Certification) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Certification) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDecisionsMade returns the DecisionsMade field value if set, zero value otherwise.
func (o *Certification) GetDecisionsMade() int32 {
	if o == nil || isNil(o.DecisionsMade) {
		var ret int32
		return ret
	}
	return *o.DecisionsMade
}

// GetDecisionsMadeOk returns a tuple with the DecisionsMade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetDecisionsMadeOk() (*int32, bool) {
	if o == nil || isNil(o.DecisionsMade) {
		return nil, false
	}
	return o.DecisionsMade, true
}

// HasDecisionsMade returns a boolean if a field has been set.
func (o *Certification) HasDecisionsMade() bool {
	if o != nil && !isNil(o.DecisionsMade) {
		return true
	}

	return false
}

// SetDecisionsMade gets a reference to the given int32 and assigns it to the DecisionsMade field.
func (o *Certification) SetDecisionsMade(v int32) {
	o.DecisionsMade = &v
}

// GetDecisionsTotal returns the DecisionsTotal field value if set, zero value otherwise.
func (o *Certification) GetDecisionsTotal() int32 {
	if o == nil || isNil(o.DecisionsTotal) {
		var ret int32
		return ret
	}
	return *o.DecisionsTotal
}

// GetDecisionsTotalOk returns a tuple with the DecisionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetDecisionsTotalOk() (*int32, bool) {
	if o == nil || isNil(o.DecisionsTotal) {
		return nil, false
	}
	return o.DecisionsTotal, true
}

// HasDecisionsTotal returns a boolean if a field has been set.
func (o *Certification) HasDecisionsTotal() bool {
	if o != nil && !isNil(o.DecisionsTotal) {
		return true
	}

	return false
}

// SetDecisionsTotal gets a reference to the given int32 and assigns it to the DecisionsTotal field.
func (o *Certification) SetDecisionsTotal(v int32) {
	o.DecisionsTotal = &v
}

// GetDue returns the Due field value if set, zero value otherwise.
func (o *Certification) GetDue() time.Time {
	if o == nil || isNil(o.Due) {
		var ret time.Time
		return ret
	}
	return *o.Due
}

// GetDueOk returns a tuple with the Due field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetDueOk() (*time.Time, bool) {
	if o == nil || isNil(o.Due) {
		return nil, false
	}
	return o.Due, true
}

// HasDue returns a boolean if a field has been set.
func (o *Certification) HasDue() bool {
	if o != nil && !isNil(o.Due) {
		return true
	}

	return false
}

// SetDue gets a reference to the given time.Time and assigns it to the Due field.
func (o *Certification) SetDue(v time.Time) {
	o.Due = &v
}

// GetSigned returns the Signed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Certification) GetSigned() time.Time {
	if o == nil || isNil(o.Signed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Signed.Get()
}

// GetSignedOk returns a tuple with the Signed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Certification) GetSignedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signed.Get(), o.Signed.IsSet()
}

// HasSigned returns a boolean if a field has been set.
func (o *Certification) HasSigned() bool {
	if o != nil && o.Signed.IsSet() {
		return true
	}

	return false
}

// SetSigned gets a reference to the given NullableTime and assigns it to the Signed field.
func (o *Certification) SetSigned(v time.Time) {
	o.Signed.Set(&v)
}
// SetSignedNil sets the value for Signed to be an explicit nil
func (o *Certification) SetSignedNil() {
	o.Signed.Set(nil)
}

// UnsetSigned ensures that no value is present for Signed, not even an explicit nil
func (o *Certification) UnsetSigned() {
	o.Signed.Unset()
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *Certification) GetReviewer() Reviewer {
	if o == nil || isNil(o.Reviewer) {
		var ret Reviewer
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetReviewerOk() (*Reviewer, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *Certification) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given Reviewer and assigns it to the Reviewer field.
func (o *Certification) SetReviewer(v Reviewer) {
	o.Reviewer = &v
}

// GetReassignment returns the Reassignment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Certification) GetReassignment() Reassignment {
	if o == nil || isNil(o.Reassignment.Get()) {
		var ret Reassignment
		return ret
	}
	return *o.Reassignment.Get()
}

// GetReassignmentOk returns a tuple with the Reassignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Certification) GetReassignmentOk() (*Reassignment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reassignment.Get(), o.Reassignment.IsSet()
}

// HasReassignment returns a boolean if a field has been set.
func (o *Certification) HasReassignment() bool {
	if o != nil && o.Reassignment.IsSet() {
		return true
	}

	return false
}

// SetReassignment gets a reference to the given NullableReassignment and assigns it to the Reassignment field.
func (o *Certification) SetReassignment(v Reassignment) {
	o.Reassignment.Set(&v)
}
// SetReassignmentNil sets the value for Reassignment to be an explicit nil
func (o *Certification) SetReassignmentNil() {
	o.Reassignment.Set(nil)
}

// UnsetReassignment ensures that no value is present for Reassignment, not even an explicit nil
func (o *Certification) UnsetReassignment() {
	o.Reassignment.Unset()
}

// GetHasErrors returns the HasErrors field value if set, zero value otherwise.
func (o *Certification) GetHasErrors() bool {
	if o == nil || isNil(o.HasErrors) {
		var ret bool
		return ret
	}
	return *o.HasErrors
}

// GetHasErrorsOk returns a tuple with the HasErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetHasErrorsOk() (*bool, bool) {
	if o == nil || isNil(o.HasErrors) {
		return nil, false
	}
	return o.HasErrors, true
}

// HasHasErrors returns a boolean if a field has been set.
func (o *Certification) HasHasErrors() bool {
	if o != nil && !isNil(o.HasErrors) {
		return true
	}

	return false
}

// SetHasErrors gets a reference to the given bool and assigns it to the HasErrors field.
func (o *Certification) SetHasErrors(v bool) {
	o.HasErrors = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Certification) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Certification) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Certification) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *Certification) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *Certification) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *Certification) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *Certification) GetPhase() CertificationPhase {
	if o == nil || isNil(o.Phase) {
		var ret CertificationPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetPhaseOk() (*CertificationPhase, bool) {
	if o == nil || isNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *Certification) HasPhase() bool {
	if o != nil && !isNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given CertificationPhase and assigns it to the Phase field.
func (o *Certification) SetPhase(v CertificationPhase) {
	o.Phase = &v
}

func (o Certification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.IdentitiesCompleted) {
		toSerialize["identitiesCompleted"] = o.IdentitiesCompleted
	}
	if !isNil(o.IdentitiesTotal) {
		toSerialize["identitiesTotal"] = o.IdentitiesTotal
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.DecisionsMade) {
		toSerialize["decisionsMade"] = o.DecisionsMade
	}
	if !isNil(o.DecisionsTotal) {
		toSerialize["decisionsTotal"] = o.DecisionsTotal
	}
	if !isNil(o.Due) {
		toSerialize["due"] = o.Due
	}
	if o.Signed.IsSet() {
		toSerialize["signed"] = o.Signed.Get()
	}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if o.Reassignment.IsSet() {
		toSerialize["reassignment"] = o.Reassignment.Get()
	}
	if !isNil(o.HasErrors) {
		toSerialize["hasErrors"] = o.HasErrors
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if !isNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Certification) UnmarshalJSON(bytes []byte) (err error) {
	varCertification := _Certification{}

	if err = json.Unmarshal(bytes, &varCertification); err == nil {
		*o = Certification(varCertification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "campaign")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "identitiesCompleted")
		delete(additionalProperties, "identitiesTotal")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "decisionsMade")
		delete(additionalProperties, "decisionsTotal")
		delete(additionalProperties, "due")
		delete(additionalProperties, "signed")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "reassignment")
		delete(additionalProperties, "hasErrors")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "phase")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertification struct {
	value *Certification
	isSet bool
}

func (v NullableCertification) Get() *Certification {
	return v.value
}

func (v *NullableCertification) Set(val *Certification) {
	v.value = val
	v.isSet = true
}

func (v NullableCertification) IsSet() bool {
	return v.isSet
}

func (v *NullableCertification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertification(val *Certification) *NullableCertification {
	return &NullableCertification{value: val, isSet: true}
}

func (v NullableCertification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


