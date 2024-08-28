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

// checks if the NonEmployeeApprovalSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeApprovalSummary{}

// NonEmployeeApprovalSummary struct for NonEmployeeApprovalSummary
type NonEmployeeApprovalSummary struct {
	// The number of approved non-employee approval requests.
	Approved *int32 `json:"approved,omitempty"`
	// The number of pending non-employee approval requests.
	Pending *int32 `json:"pending,omitempty"`
	// The number of rejected non-employee approval requests.
	Rejected *int32 `json:"rejected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeApprovalSummary NonEmployeeApprovalSummary

// NewNonEmployeeApprovalSummary instantiates a new NonEmployeeApprovalSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeApprovalSummary() *NonEmployeeApprovalSummary {
	this := NonEmployeeApprovalSummary{}
	return &this
}

// NewNonEmployeeApprovalSummaryWithDefaults instantiates a new NonEmployeeApprovalSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeApprovalSummaryWithDefaults() *NonEmployeeApprovalSummary {
	this := NonEmployeeApprovalSummary{}
	return &this
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *NonEmployeeApprovalSummary) GetApproved() int32 {
	if o == nil || IsNil(o.Approved) {
		var ret int32
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalSummary) GetApprovedOk() (*int32, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *NonEmployeeApprovalSummary) HasApproved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given int32 and assigns it to the Approved field.
func (o *NonEmployeeApprovalSummary) SetApproved(v int32) {
	o.Approved = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *NonEmployeeApprovalSummary) GetPending() int32 {
	if o == nil || IsNil(o.Pending) {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalSummary) GetPendingOk() (*int32, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *NonEmployeeApprovalSummary) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *NonEmployeeApprovalSummary) SetPending(v int32) {
	o.Pending = &v
}

// GetRejected returns the Rejected field value if set, zero value otherwise.
func (o *NonEmployeeApprovalSummary) GetRejected() int32 {
	if o == nil || IsNil(o.Rejected) {
		var ret int32
		return ret
	}
	return *o.Rejected
}

// GetRejectedOk returns a tuple with the Rejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalSummary) GetRejectedOk() (*int32, bool) {
	if o == nil || IsNil(o.Rejected) {
		return nil, false
	}
	return o.Rejected, true
}

// HasRejected returns a boolean if a field has been set.
func (o *NonEmployeeApprovalSummary) HasRejected() bool {
	if o != nil && !IsNil(o.Rejected) {
		return true
	}

	return false
}

// SetRejected gets a reference to the given int32 and assigns it to the Rejected field.
func (o *NonEmployeeApprovalSummary) SetRejected(v int32) {
	o.Rejected = &v
}

func (o NonEmployeeApprovalSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeApprovalSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !IsNil(o.Rejected) {
		toSerialize["rejected"] = o.Rejected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeApprovalSummary) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeApprovalSummary := _NonEmployeeApprovalSummary{}

	err = json.Unmarshal(data, &varNonEmployeeApprovalSummary)

	if err != nil {
		return err
	}

	*o = NonEmployeeApprovalSummary(varNonEmployeeApprovalSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approved")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "rejected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeApprovalSummary struct {
	value *NonEmployeeApprovalSummary
	isSet bool
}

func (v NullableNonEmployeeApprovalSummary) Get() *NonEmployeeApprovalSummary {
	return v.value
}

func (v *NullableNonEmployeeApprovalSummary) Set(val *NonEmployeeApprovalSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeApprovalSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeApprovalSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeApprovalSummary(val *NonEmployeeApprovalSummary) *NullableNonEmployeeApprovalSummary {
	return &NullableNonEmployeeApprovalSummary{value: val, isSet: true}
}

func (v NullableNonEmployeeApprovalSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeApprovalSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


