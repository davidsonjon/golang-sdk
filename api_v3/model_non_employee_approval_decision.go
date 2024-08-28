/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the NonEmployeeApprovalDecision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeApprovalDecision{}

// NonEmployeeApprovalDecision struct for NonEmployeeApprovalDecision
type NonEmployeeApprovalDecision struct {
	// Comment on the approval item.
	Comment *string `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeApprovalDecision NonEmployeeApprovalDecision

// NewNonEmployeeApprovalDecision instantiates a new NonEmployeeApprovalDecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeApprovalDecision() *NonEmployeeApprovalDecision {
	this := NonEmployeeApprovalDecision{}
	return &this
}

// NewNonEmployeeApprovalDecisionWithDefaults instantiates a new NonEmployeeApprovalDecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeApprovalDecisionWithDefaults() *NonEmployeeApprovalDecision {
	this := NonEmployeeApprovalDecision{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NonEmployeeApprovalDecision) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalDecision) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NonEmployeeApprovalDecision) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NonEmployeeApprovalDecision) SetComment(v string) {
	o.Comment = &v
}

func (o NonEmployeeApprovalDecision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeApprovalDecision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeApprovalDecision) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeApprovalDecision := _NonEmployeeApprovalDecision{}

	err = json.Unmarshal(data, &varNonEmployeeApprovalDecision)

	if err != nil {
		return err
	}

	*o = NonEmployeeApprovalDecision(varNonEmployeeApprovalDecision)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeApprovalDecision struct {
	value *NonEmployeeApprovalDecision
	isSet bool
}

func (v NullableNonEmployeeApprovalDecision) Get() *NonEmployeeApprovalDecision {
	return v.value
}

func (v *NullableNonEmployeeApprovalDecision) Set(val *NonEmployeeApprovalDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeApprovalDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeApprovalDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeApprovalDecision(val *NonEmployeeApprovalDecision) *NullableNonEmployeeApprovalDecision {
	return &NullableNonEmployeeApprovalDecision{value: val, isSet: true}
}

func (v NullableNonEmployeeApprovalDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeApprovalDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


