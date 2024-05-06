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

// checks if the SedApprovalStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SedApprovalStatus{}

// SedApprovalStatus SED Approval Status
type SedApprovalStatus struct {
	// failed reason will be display if status is failed
	FailedReason *string `json:"failedReason,omitempty"`
	// Sed id
	Id *string `json:"id,omitempty"`
	// SUCCESS | FAILED
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SedApprovalStatus SedApprovalStatus

// NewSedApprovalStatus instantiates a new SedApprovalStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSedApprovalStatus() *SedApprovalStatus {
	this := SedApprovalStatus{}
	return &this
}

// NewSedApprovalStatusWithDefaults instantiates a new SedApprovalStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSedApprovalStatusWithDefaults() *SedApprovalStatus {
	this := SedApprovalStatus{}
	return &this
}

// GetFailedReason returns the FailedReason field value if set, zero value otherwise.
func (o *SedApprovalStatus) GetFailedReason() string {
	if o == nil || isNil(o.FailedReason) {
		var ret string
		return ret
	}
	return *o.FailedReason
}

// GetFailedReasonOk returns a tuple with the FailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SedApprovalStatus) GetFailedReasonOk() (*string, bool) {
	if o == nil || isNil(o.FailedReason) {
		return nil, false
	}
	return o.FailedReason, true
}

// HasFailedReason returns a boolean if a field has been set.
func (o *SedApprovalStatus) HasFailedReason() bool {
	if o != nil && !isNil(o.FailedReason) {
		return true
	}

	return false
}

// SetFailedReason gets a reference to the given string and assigns it to the FailedReason field.
func (o *SedApprovalStatus) SetFailedReason(v string) {
	o.FailedReason = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SedApprovalStatus) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SedApprovalStatus) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SedApprovalStatus) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SedApprovalStatus) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SedApprovalStatus) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SedApprovalStatus) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SedApprovalStatus) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SedApprovalStatus) SetStatus(v string) {
	o.Status = &v
}

func (o SedApprovalStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SedApprovalStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FailedReason) {
		toSerialize["failedReason"] = o.FailedReason
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SedApprovalStatus) UnmarshalJSON(bytes []byte) (err error) {
	varSedApprovalStatus := _SedApprovalStatus{}

	if err = json.Unmarshal(bytes, &varSedApprovalStatus); err == nil {
	*o = SedApprovalStatus(varSedApprovalStatus)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "failedReason")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSedApprovalStatus struct {
	value *SedApprovalStatus
	isSet bool
}

func (v NullableSedApprovalStatus) Get() *SedApprovalStatus {
	return v.value
}

func (v *NullableSedApprovalStatus) Set(val *SedApprovalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSedApprovalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSedApprovalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSedApprovalStatus(val *SedApprovalStatus) *NullableSedApprovalStatus {
	return &NullableSedApprovalStatus{value: val, isSet: true}
}

func (v NullableSedApprovalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSedApprovalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


