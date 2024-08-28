/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the AdminReviewReassign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminReviewReassign{}

// AdminReviewReassign struct for AdminReviewReassign
type AdminReviewReassign struct {
	// List of certification IDs to reassign
	CertificationIds []string `json:"certificationIds,omitempty"`
	ReassignTo *AdminReviewReassignReassignTo `json:"reassignTo,omitempty"`
	// Comment to explain why the certification was reassigned
	Reason *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdminReviewReassign AdminReviewReassign

// NewAdminReviewReassign instantiates a new AdminReviewReassign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminReviewReassign() *AdminReviewReassign {
	this := AdminReviewReassign{}
	return &this
}

// NewAdminReviewReassignWithDefaults instantiates a new AdminReviewReassign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminReviewReassignWithDefaults() *AdminReviewReassign {
	this := AdminReviewReassign{}
	return &this
}

// GetCertificationIds returns the CertificationIds field value if set, zero value otherwise.
func (o *AdminReviewReassign) GetCertificationIds() []string {
	if o == nil || IsNil(o.CertificationIds) {
		var ret []string
		return ret
	}
	return o.CertificationIds
}

// GetCertificationIdsOk returns a tuple with the CertificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReviewReassign) GetCertificationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CertificationIds) {
		return nil, false
	}
	return o.CertificationIds, true
}

// HasCertificationIds returns a boolean if a field has been set.
func (o *AdminReviewReassign) HasCertificationIds() bool {
	if o != nil && !IsNil(o.CertificationIds) {
		return true
	}

	return false
}

// SetCertificationIds gets a reference to the given []string and assigns it to the CertificationIds field.
func (o *AdminReviewReassign) SetCertificationIds(v []string) {
	o.CertificationIds = v
}

// GetReassignTo returns the ReassignTo field value if set, zero value otherwise.
func (o *AdminReviewReassign) GetReassignTo() AdminReviewReassignReassignTo {
	if o == nil || IsNil(o.ReassignTo) {
		var ret AdminReviewReassignReassignTo
		return ret
	}
	return *o.ReassignTo
}

// GetReassignToOk returns a tuple with the ReassignTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReviewReassign) GetReassignToOk() (*AdminReviewReassignReassignTo, bool) {
	if o == nil || IsNil(o.ReassignTo) {
		return nil, false
	}
	return o.ReassignTo, true
}

// HasReassignTo returns a boolean if a field has been set.
func (o *AdminReviewReassign) HasReassignTo() bool {
	if o != nil && !IsNil(o.ReassignTo) {
		return true
	}

	return false
}

// SetReassignTo gets a reference to the given AdminReviewReassignReassignTo and assigns it to the ReassignTo field.
func (o *AdminReviewReassign) SetReassignTo(v AdminReviewReassignReassignTo) {
	o.ReassignTo = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AdminReviewReassign) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReviewReassign) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AdminReviewReassign) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AdminReviewReassign) SetReason(v string) {
	o.Reason = &v
}

func (o AdminReviewReassign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminReviewReassign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificationIds) {
		toSerialize["certificationIds"] = o.CertificationIds
	}
	if !IsNil(o.ReassignTo) {
		toSerialize["reassignTo"] = o.ReassignTo
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdminReviewReassign) UnmarshalJSON(data []byte) (err error) {
	varAdminReviewReassign := _AdminReviewReassign{}

	err = json.Unmarshal(data, &varAdminReviewReassign)

	if err != nil {
		return err
	}

	*o = AdminReviewReassign(varAdminReviewReassign)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificationIds")
		delete(additionalProperties, "reassignTo")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdminReviewReassign struct {
	value *AdminReviewReassign
	isSet bool
}

func (v NullableAdminReviewReassign) Get() *AdminReviewReassign {
	return v.value
}

func (v *NullableAdminReviewReassign) Set(val *AdminReviewReassign) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminReviewReassign) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminReviewReassign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminReviewReassign(val *AdminReviewReassign) *NullableAdminReviewReassign {
	return &NullableAdminReviewReassign{value: val, isSet: true}
}

func (v NullableAdminReviewReassign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminReviewReassign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


