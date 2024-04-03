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

// checks if the CorrelatedGovernanceEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorrelatedGovernanceEvent{}

// CorrelatedGovernanceEvent struct for CorrelatedGovernanceEvent
type CorrelatedGovernanceEvent struct {
	// The name of the governance event, such as the certification name or access request ID.
	Name *string `json:"name,omitempty"`
	// The date that the certification or access request was completed.
	Dt *string `json:"dt,omitempty"`
	// The type of governance event.
	Type *string `json:"type,omitempty"`
	// The ID of the instance that caused the event - either the certification ID or access request ID.
	GovernanceId *string `json:"governanceId,omitempty"`
	// The owners of the governance event (the certifiers or approvers)
	Owners []CertifierResponse `json:"owners,omitempty"`
	// The owners of the governance event (the certifiers or approvers), this field should be preferred over owners
	Reviewers []CertifierResponse `json:"reviewers,omitempty"`
	DecisionMaker *CertifierResponse `json:"decisionMaker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CorrelatedGovernanceEvent CorrelatedGovernanceEvent

// NewCorrelatedGovernanceEvent instantiates a new CorrelatedGovernanceEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrelatedGovernanceEvent() *CorrelatedGovernanceEvent {
	this := CorrelatedGovernanceEvent{}
	return &this
}

// NewCorrelatedGovernanceEventWithDefaults instantiates a new CorrelatedGovernanceEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrelatedGovernanceEventWithDefaults() *CorrelatedGovernanceEvent {
	this := CorrelatedGovernanceEvent{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CorrelatedGovernanceEvent) SetName(v string) {
	o.Name = &v
}

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetDt() string {
	if o == nil || isNil(o.Dt) {
		var ret string
		return ret
	}
	return *o.Dt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetDtOk() (*string, bool) {
	if o == nil || isNil(o.Dt) {
		return nil, false
	}
	return o.Dt, true
}

// HasDt returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasDt() bool {
	if o != nil && !isNil(o.Dt) {
		return true
	}

	return false
}

// SetDt gets a reference to the given string and assigns it to the Dt field.
func (o *CorrelatedGovernanceEvent) SetDt(v string) {
	o.Dt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CorrelatedGovernanceEvent) SetType(v string) {
	o.Type = &v
}

// GetGovernanceId returns the GovernanceId field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetGovernanceId() string {
	if o == nil || isNil(o.GovernanceId) {
		var ret string
		return ret
	}
	return *o.GovernanceId
}

// GetGovernanceIdOk returns a tuple with the GovernanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetGovernanceIdOk() (*string, bool) {
	if o == nil || isNil(o.GovernanceId) {
		return nil, false
	}
	return o.GovernanceId, true
}

// HasGovernanceId returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasGovernanceId() bool {
	if o != nil && !isNil(o.GovernanceId) {
		return true
	}

	return false
}

// SetGovernanceId gets a reference to the given string and assigns it to the GovernanceId field.
func (o *CorrelatedGovernanceEvent) SetGovernanceId(v string) {
	o.GovernanceId = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetOwners() []CertifierResponse {
	if o == nil || isNil(o.Owners) {
		var ret []CertifierResponse
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetOwnersOk() ([]CertifierResponse, bool) {
	if o == nil || isNil(o.Owners) {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasOwners() bool {
	if o != nil && !isNil(o.Owners) {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []CertifierResponse and assigns it to the Owners field.
func (o *CorrelatedGovernanceEvent) SetOwners(v []CertifierResponse) {
	o.Owners = v
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetReviewers() []CertifierResponse {
	if o == nil || isNil(o.Reviewers) {
		var ret []CertifierResponse
		return ret
	}
	return o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetReviewersOk() ([]CertifierResponse, bool) {
	if o == nil || isNil(o.Reviewers) {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasReviewers() bool {
	if o != nil && !isNil(o.Reviewers) {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []CertifierResponse and assigns it to the Reviewers field.
func (o *CorrelatedGovernanceEvent) SetReviewers(v []CertifierResponse) {
	o.Reviewers = v
}

// GetDecisionMaker returns the DecisionMaker field value if set, zero value otherwise.
func (o *CorrelatedGovernanceEvent) GetDecisionMaker() CertifierResponse {
	if o == nil || isNil(o.DecisionMaker) {
		var ret CertifierResponse
		return ret
	}
	return *o.DecisionMaker
}

// GetDecisionMakerOk returns a tuple with the DecisionMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrelatedGovernanceEvent) GetDecisionMakerOk() (*CertifierResponse, bool) {
	if o == nil || isNil(o.DecisionMaker) {
		return nil, false
	}
	return o.DecisionMaker, true
}

// HasDecisionMaker returns a boolean if a field has been set.
func (o *CorrelatedGovernanceEvent) HasDecisionMaker() bool {
	if o != nil && !isNil(o.DecisionMaker) {
		return true
	}

	return false
}

// SetDecisionMaker gets a reference to the given CertifierResponse and assigns it to the DecisionMaker field.
func (o *CorrelatedGovernanceEvent) SetDecisionMaker(v CertifierResponse) {
	o.DecisionMaker = &v
}

func (o CorrelatedGovernanceEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorrelatedGovernanceEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Dt) {
		toSerialize["dt"] = o.Dt
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.GovernanceId) {
		toSerialize["governanceId"] = o.GovernanceId
	}
	if !isNil(o.Owners) {
		toSerialize["owners"] = o.Owners
	}
	if !isNil(o.Reviewers) {
		toSerialize["reviewers"] = o.Reviewers
	}
	if !isNil(o.DecisionMaker) {
		toSerialize["decisionMaker"] = o.DecisionMaker
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CorrelatedGovernanceEvent) UnmarshalJSON(bytes []byte) (err error) {
	varCorrelatedGovernanceEvent := _CorrelatedGovernanceEvent{}

	if err = json.Unmarshal(bytes, &varCorrelatedGovernanceEvent); err == nil {
	*o = CorrelatedGovernanceEvent(varCorrelatedGovernanceEvent)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "dt")
		delete(additionalProperties, "type")
		delete(additionalProperties, "governanceId")
		delete(additionalProperties, "owners")
		delete(additionalProperties, "reviewers")
		delete(additionalProperties, "decisionMaker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCorrelatedGovernanceEvent struct {
	value *CorrelatedGovernanceEvent
	isSet bool
}

func (v NullableCorrelatedGovernanceEvent) Get() *CorrelatedGovernanceEvent {
	return v.value
}

func (v *NullableCorrelatedGovernanceEvent) Set(val *CorrelatedGovernanceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrelatedGovernanceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrelatedGovernanceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrelatedGovernanceEvent(val *CorrelatedGovernanceEvent) *NullableCorrelatedGovernanceEvent {
	return &NullableCorrelatedGovernanceEvent{value: val, isSet: true}
}

func (v NullableCorrelatedGovernanceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrelatedGovernanceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


