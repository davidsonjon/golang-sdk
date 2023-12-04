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

// checks if the PreApprovalTriggerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreApprovalTriggerDetails{}

// PreApprovalTriggerDetails Provides additional details about the pre-approval trigger for this request.
type PreApprovalTriggerDetails struct {
	// Comment left for the pre-approval decision
	Comment *string `json:"comment,omitempty"`
	// The reviewer of the pre-approval decision
	Reviewer *string `json:"reviewer,omitempty"`
	// The decision of the pre-approval trigger
	Decision *string `json:"decision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreApprovalTriggerDetails PreApprovalTriggerDetails

// NewPreApprovalTriggerDetails instantiates a new PreApprovalTriggerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreApprovalTriggerDetails() *PreApprovalTriggerDetails {
	this := PreApprovalTriggerDetails{}
	return &this
}

// NewPreApprovalTriggerDetailsWithDefaults instantiates a new PreApprovalTriggerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreApprovalTriggerDetailsWithDefaults() *PreApprovalTriggerDetails {
	this := PreApprovalTriggerDetails{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PreApprovalTriggerDetails) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreApprovalTriggerDetails) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PreApprovalTriggerDetails) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *PreApprovalTriggerDetails) SetComment(v string) {
	o.Comment = &v
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *PreApprovalTriggerDetails) GetReviewer() string {
	if o == nil || isNil(o.Reviewer) {
		var ret string
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreApprovalTriggerDetails) GetReviewerOk() (*string, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *PreApprovalTriggerDetails) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given string and assigns it to the Reviewer field.
func (o *PreApprovalTriggerDetails) SetReviewer(v string) {
	o.Reviewer = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *PreApprovalTriggerDetails) GetDecision() string {
	if o == nil || isNil(o.Decision) {
		var ret string
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreApprovalTriggerDetails) GetDecisionOk() (*string, bool) {
	if o == nil || isNil(o.Decision) {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *PreApprovalTriggerDetails) HasDecision() bool {
	if o != nil && !isNil(o.Decision) {
		return true
	}

	return false
}

// SetDecision gets a reference to the given string and assigns it to the Decision field.
func (o *PreApprovalTriggerDetails) SetDecision(v string) {
	o.Decision = &v
}

func (o PreApprovalTriggerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreApprovalTriggerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !isNil(o.Decision) {
		toSerialize["decision"] = o.Decision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreApprovalTriggerDetails) UnmarshalJSON(bytes []byte) (err error) {
	varPreApprovalTriggerDetails := _PreApprovalTriggerDetails{}

	if err = json.Unmarshal(bytes, &varPreApprovalTriggerDetails); err == nil {
	*o = PreApprovalTriggerDetails(varPreApprovalTriggerDetails)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "decision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreApprovalTriggerDetails struct {
	value *PreApprovalTriggerDetails
	isSet bool
}

func (v NullablePreApprovalTriggerDetails) Get() *PreApprovalTriggerDetails {
	return v.value
}

func (v *NullablePreApprovalTriggerDetails) Set(val *PreApprovalTriggerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePreApprovalTriggerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePreApprovalTriggerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreApprovalTriggerDetails(val *PreApprovalTriggerDetails) *NullablePreApprovalTriggerDetails {
	return &NullablePreApprovalTriggerDetails{value: val, isSet: true}
}

func (v NullablePreApprovalTriggerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreApprovalTriggerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


