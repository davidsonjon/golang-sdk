/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ApprovalSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalSummary{}

// ApprovalSummary struct for ApprovalSummary
type ApprovalSummary struct {
	// The number of pending access requests approvals.
	Pending *int32 `json:"pending,omitempty"`
	// The number of approved access requests approvals.
	Approved *int32 `json:"approved,omitempty"`
	// The number of rejected access requests approvals.
	Rejected *int32 `json:"rejected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalSummary ApprovalSummary

// NewApprovalSummary instantiates a new ApprovalSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalSummary() *ApprovalSummary {
	this := ApprovalSummary{}
	return &this
}

// NewApprovalSummaryWithDefaults instantiates a new ApprovalSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalSummaryWithDefaults() *ApprovalSummary {
	this := ApprovalSummary{}
	return &this
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *ApprovalSummary) GetPending() int32 {
	if o == nil || isNil(o.Pending) {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalSummary) GetPendingOk() (*int32, bool) {
	if o == nil || isNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *ApprovalSummary) HasPending() bool {
	if o != nil && !isNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *ApprovalSummary) SetPending(v int32) {
	o.Pending = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *ApprovalSummary) GetApproved() int32 {
	if o == nil || isNil(o.Approved) {
		var ret int32
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalSummary) GetApprovedOk() (*int32, bool) {
	if o == nil || isNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *ApprovalSummary) HasApproved() bool {
	if o != nil && !isNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given int32 and assigns it to the Approved field.
func (o *ApprovalSummary) SetApproved(v int32) {
	o.Approved = &v
}

// GetRejected returns the Rejected field value if set, zero value otherwise.
func (o *ApprovalSummary) GetRejected() int32 {
	if o == nil || isNil(o.Rejected) {
		var ret int32
		return ret
	}
	return *o.Rejected
}

// GetRejectedOk returns a tuple with the Rejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalSummary) GetRejectedOk() (*int32, bool) {
	if o == nil || isNil(o.Rejected) {
		return nil, false
	}
	return o.Rejected, true
}

// HasRejected returns a boolean if a field has been set.
func (o *ApprovalSummary) HasRejected() bool {
	if o != nil && !isNil(o.Rejected) {
		return true
	}

	return false
}

// SetRejected gets a reference to the given int32 and assigns it to the Rejected field.
func (o *ApprovalSummary) SetRejected(v int32) {
	o.Rejected = &v
}

func (o ApprovalSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !isNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !isNil(o.Rejected) {
		toSerialize["rejected"] = o.Rejected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalSummary) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalSummary := _ApprovalSummary{}

	if err = json.Unmarshal(bytes, &varApprovalSummary); err == nil {
	*o = ApprovalSummary(varApprovalSummary)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pending")
		delete(additionalProperties, "approved")
		delete(additionalProperties, "rejected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalSummary struct {
	value *ApprovalSummary
	isSet bool
}

func (v NullableApprovalSummary) Get() *ApprovalSummary {
	return v.value
}

func (v *NullableApprovalSummary) Set(val *ApprovalSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalSummary(val *ApprovalSummary) *NullableApprovalSummary {
	return &NullableApprovalSummary{value: val, isSet: true}
}

func (v NullableApprovalSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


