/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the AccessReviewReassignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessReviewReassignment{}

// AccessReviewReassignment struct for AccessReviewReassignment
type AccessReviewReassignment struct {
	Reassign []ReassignReference `json:"reassign"`
	// The ID of the identity to which the certification is reassigned
	ReassignTo string `json:"reassignTo"`
	// The reason comment for why the reassign was made
	Reason string `json:"reason"`
	AdditionalProperties map[string]interface{}
}

type _AccessReviewReassignment AccessReviewReassignment

// NewAccessReviewReassignment instantiates a new AccessReviewReassignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessReviewReassignment(reassign []ReassignReference, reassignTo string, reason string) *AccessReviewReassignment {
	this := AccessReviewReassignment{}
	this.Reassign = reassign
	this.ReassignTo = reassignTo
	this.Reason = reason
	return &this
}

// NewAccessReviewReassignmentWithDefaults instantiates a new AccessReviewReassignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessReviewReassignmentWithDefaults() *AccessReviewReassignment {
	this := AccessReviewReassignment{}
	return &this
}

// GetReassign returns the Reassign field value
func (o *AccessReviewReassignment) GetReassign() []ReassignReference {
	if o == nil {
		var ret []ReassignReference
		return ret
	}

	return o.Reassign
}

// GetReassignOk returns a tuple with the Reassign field value
// and a boolean to check if the value has been set.
func (o *AccessReviewReassignment) GetReassignOk() ([]ReassignReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reassign, true
}

// SetReassign sets field value
func (o *AccessReviewReassignment) SetReassign(v []ReassignReference) {
	o.Reassign = v
}

// GetReassignTo returns the ReassignTo field value
func (o *AccessReviewReassignment) GetReassignTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReassignTo
}

// GetReassignToOk returns a tuple with the ReassignTo field value
// and a boolean to check if the value has been set.
func (o *AccessReviewReassignment) GetReassignToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReassignTo, true
}

// SetReassignTo sets field value
func (o *AccessReviewReassignment) SetReassignTo(v string) {
	o.ReassignTo = v
}

// GetReason returns the Reason field value
func (o *AccessReviewReassignment) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *AccessReviewReassignment) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *AccessReviewReassignment) SetReason(v string) {
	o.Reason = v
}

func (o AccessReviewReassignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessReviewReassignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reassign"] = o.Reassign
	toSerialize["reassignTo"] = o.ReassignTo
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessReviewReassignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reassign",
		"reassignTo",
		"reason",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessReviewReassignment := _AccessReviewReassignment{}

	err = json.Unmarshal(data, &varAccessReviewReassignment)

	if err != nil {
		return err
	}

	*o = AccessReviewReassignment(varAccessReviewReassignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reassign")
		delete(additionalProperties, "reassignTo")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessReviewReassignment struct {
	value *AccessReviewReassignment
	isSet bool
}

func (v NullableAccessReviewReassignment) Get() *AccessReviewReassignment {
	return v.value
}

func (v *NullableAccessReviewReassignment) Set(val *AccessReviewReassignment) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessReviewReassignment) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessReviewReassignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessReviewReassignment(val *AccessReviewReassignment) *NullableAccessReviewReassignment {
	return &NullableAccessReviewReassignment{value: val, isSet: true}
}

func (v NullableAccessReviewReassignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessReviewReassignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


