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

// CommonAccessIDStatus struct for CommonAccessIDStatus
type CommonAccessIDStatus struct {
	// List of confirmed common access ids.
	ConfirmedIds []string `json:"confirmedIds,omitempty"`
	// List of denied common access ids.
	DeniedIds []string `json:"deniedIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonAccessIDStatus CommonAccessIDStatus

// NewCommonAccessIDStatus instantiates a new CommonAccessIDStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAccessIDStatus() *CommonAccessIDStatus {
	this := CommonAccessIDStatus{}
	return &this
}

// NewCommonAccessIDStatusWithDefaults instantiates a new CommonAccessIDStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAccessIDStatusWithDefaults() *CommonAccessIDStatus {
	this := CommonAccessIDStatus{}
	return &this
}

// GetConfirmedIds returns the ConfirmedIds field value if set, zero value otherwise.
func (o *CommonAccessIDStatus) GetConfirmedIds() []string {
	if o == nil || isNil(o.ConfirmedIds) {
		var ret []string
		return ret
	}
	return o.ConfirmedIds
}

// GetConfirmedIdsOk returns a tuple with the ConfirmedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessIDStatus) GetConfirmedIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ConfirmedIds) {
		return nil, false
	}
	return o.ConfirmedIds, true
}

// HasConfirmedIds returns a boolean if a field has been set.
func (o *CommonAccessIDStatus) HasConfirmedIds() bool {
	if o != nil && !isNil(o.ConfirmedIds) {
		return true
	}

	return false
}

// SetConfirmedIds gets a reference to the given []string and assigns it to the ConfirmedIds field.
func (o *CommonAccessIDStatus) SetConfirmedIds(v []string) {
	o.ConfirmedIds = v
}

// GetDeniedIds returns the DeniedIds field value if set, zero value otherwise.
func (o *CommonAccessIDStatus) GetDeniedIds() []string {
	if o == nil || isNil(o.DeniedIds) {
		var ret []string
		return ret
	}
	return o.DeniedIds
}

// GetDeniedIdsOk returns a tuple with the DeniedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessIDStatus) GetDeniedIdsOk() ([]string, bool) {
	if o == nil || isNil(o.DeniedIds) {
		return nil, false
	}
	return o.DeniedIds, true
}

// HasDeniedIds returns a boolean if a field has been set.
func (o *CommonAccessIDStatus) HasDeniedIds() bool {
	if o != nil && !isNil(o.DeniedIds) {
		return true
	}

	return false
}

// SetDeniedIds gets a reference to the given []string and assigns it to the DeniedIds field.
func (o *CommonAccessIDStatus) SetDeniedIds(v []string) {
	o.DeniedIds = v
}

func (o CommonAccessIDStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConfirmedIds) {
		toSerialize["confirmedIds"] = o.ConfirmedIds
	}
	if !isNil(o.DeniedIds) {
		toSerialize["deniedIds"] = o.DeniedIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommonAccessIDStatus) UnmarshalJSON(bytes []byte) (err error) {
	varCommonAccessIDStatus := _CommonAccessIDStatus{}

	if err = json.Unmarshal(bytes, &varCommonAccessIDStatus); err == nil {
		*o = CommonAccessIDStatus(varCommonAccessIDStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "confirmedIds")
		delete(additionalProperties, "deniedIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAccessIDStatus struct {
	value *CommonAccessIDStatus
	isSet bool
}

func (v NullableCommonAccessIDStatus) Get() *CommonAccessIDStatus {
	return v.value
}

func (v *NullableCommonAccessIDStatus) Set(val *CommonAccessIDStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAccessIDStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAccessIDStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAccessIDStatus(val *CommonAccessIDStatus) *NullableCommonAccessIDStatus {
	return &NullableCommonAccessIDStatus{value: val, isSet: true}
}

func (v NullableCommonAccessIDStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAccessIDStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


