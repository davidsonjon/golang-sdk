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
)

// checks if the ReviewRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewRecommendation{}

// ReviewRecommendation struct for ReviewRecommendation
type ReviewRecommendation struct {
	// The recommendation from IAI at the time of the decision. This field will be null if no recommendation was made.
	Recommendation NullableString `json:"recommendation,omitempty"`
	// A list of reasons for the recommendation.
	Reasons []string `json:"reasons,omitempty"`
	// The time at which the recommendation was recorded.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewRecommendation ReviewRecommendation

// NewReviewRecommendation instantiates a new ReviewRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewRecommendation() *ReviewRecommendation {
	this := ReviewRecommendation{}
	return &this
}

// NewReviewRecommendationWithDefaults instantiates a new ReviewRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewRecommendationWithDefaults() *ReviewRecommendation {
	this := ReviewRecommendation{}
	return &this
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewRecommendation) GetRecommendation() string {
	if o == nil || isNil(o.Recommendation.Get()) {
		var ret string
		return ret
	}
	return *o.Recommendation.Get()
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewRecommendation) GetRecommendationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recommendation.Get(), o.Recommendation.IsSet()
}

// HasRecommendation returns a boolean if a field has been set.
func (o *ReviewRecommendation) HasRecommendation() bool {
	if o != nil && o.Recommendation.IsSet() {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given NullableString and assigns it to the Recommendation field.
func (o *ReviewRecommendation) SetRecommendation(v string) {
	o.Recommendation.Set(&v)
}
// SetRecommendationNil sets the value for Recommendation to be an explicit nil
func (o *ReviewRecommendation) SetRecommendationNil() {
	o.Recommendation.Set(nil)
}

// UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
func (o *ReviewRecommendation) UnsetRecommendation() {
	o.Recommendation.Unset()
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *ReviewRecommendation) GetReasons() []string {
	if o == nil || isNil(o.Reasons) {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewRecommendation) GetReasonsOk() ([]string, bool) {
	if o == nil || isNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *ReviewRecommendation) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *ReviewRecommendation) SetReasons(v []string) {
	o.Reasons = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReviewRecommendation) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewRecommendation) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReviewRecommendation) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ReviewRecommendation) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o ReviewRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Recommendation.IsSet() {
		toSerialize["recommendation"] = o.Recommendation.Get()
	}
	if !isNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewRecommendation) UnmarshalJSON(bytes []byte) (err error) {
	varReviewRecommendation := _ReviewRecommendation{}

	if err = json.Unmarshal(bytes, &varReviewRecommendation); err == nil {
	*o = ReviewRecommendation(varReviewRecommendation)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recommendation")
		delete(additionalProperties, "reasons")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewRecommendation struct {
	value *ReviewRecommendation
	isSet bool
}

func (v NullableReviewRecommendation) Get() *ReviewRecommendation {
	return v.value
}

func (v *NullableReviewRecommendation) Set(val *ReviewRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewRecommendation(val *ReviewRecommendation) *NullableReviewRecommendation {
	return &NullableReviewRecommendation{value: val, isSet: true}
}

func (v NullableReviewRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


