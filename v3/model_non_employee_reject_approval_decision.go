/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the NonEmployeeRejectApprovalDecision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeRejectApprovalDecision{}

// NonEmployeeRejectApprovalDecision struct for NonEmployeeRejectApprovalDecision
type NonEmployeeRejectApprovalDecision struct {
	// Comment on the approval item.
	Comment string `json:"comment"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeRejectApprovalDecision NonEmployeeRejectApprovalDecision

// NewNonEmployeeRejectApprovalDecision instantiates a new NonEmployeeRejectApprovalDecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeRejectApprovalDecision(comment string) *NonEmployeeRejectApprovalDecision {
	this := NonEmployeeRejectApprovalDecision{}
	this.Comment = comment
	return &this
}

// NewNonEmployeeRejectApprovalDecisionWithDefaults instantiates a new NonEmployeeRejectApprovalDecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeRejectApprovalDecisionWithDefaults() *NonEmployeeRejectApprovalDecision {
	this := NonEmployeeRejectApprovalDecision{}
	return &this
}

// GetComment returns the Comment field value
func (o *NonEmployeeRejectApprovalDecision) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRejectApprovalDecision) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *NonEmployeeRejectApprovalDecision) SetComment(v string) {
	o.Comment = v
}

func (o NonEmployeeRejectApprovalDecision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeRejectApprovalDecision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comment"] = o.Comment

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeRejectApprovalDecision) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"comment",
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

	varNonEmployeeRejectApprovalDecision := _NonEmployeeRejectApprovalDecision{}

	if err = json.Unmarshal(bytes, &varNonEmployeeRejectApprovalDecision); err == nil {
	*o = NonEmployeeRejectApprovalDecision(varNonEmployeeRejectApprovalDecision)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeRejectApprovalDecision struct {
	value *NonEmployeeRejectApprovalDecision
	isSet bool
}

func (v NullableNonEmployeeRejectApprovalDecision) Get() *NonEmployeeRejectApprovalDecision {
	return v.value
}

func (v *NullableNonEmployeeRejectApprovalDecision) Set(val *NonEmployeeRejectApprovalDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeRejectApprovalDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeRejectApprovalDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeRejectApprovalDecision(val *NonEmployeeRejectApprovalDecision) *NullableNonEmployeeRejectApprovalDecision {
	return &NullableNonEmployeeRejectApprovalDecision{value: val, isSet: true}
}

func (v NullableNonEmployeeRejectApprovalDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeRejectApprovalDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


