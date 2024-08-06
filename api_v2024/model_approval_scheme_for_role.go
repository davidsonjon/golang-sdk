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

// checks if the ApprovalSchemeForRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalSchemeForRole{}

// ApprovalSchemeForRole struct for ApprovalSchemeForRole
type ApprovalSchemeForRole struct {
	// Describes the individual or group that is responsible for an approval step. Values are as follows.  **OWNER**: Owner of the associated Role  **MANAGER**: Manager of the Identity making the request  **GOVERNANCE_GROUP**: A Governance Group, the ID of which is specified by the **approverId** field
	ApproverType *string `json:"approverType,omitempty"`
	// Id of the specific approver, used only when approverType is GOVERNANCE_GROUP
	ApproverId NullableString `json:"approverId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalSchemeForRole ApprovalSchemeForRole

// NewApprovalSchemeForRole instantiates a new ApprovalSchemeForRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalSchemeForRole() *ApprovalSchemeForRole {
	this := ApprovalSchemeForRole{}
	return &this
}

// NewApprovalSchemeForRoleWithDefaults instantiates a new ApprovalSchemeForRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalSchemeForRoleWithDefaults() *ApprovalSchemeForRole {
	this := ApprovalSchemeForRole{}
	return &this
}

// GetApproverType returns the ApproverType field value if set, zero value otherwise.
func (o *ApprovalSchemeForRole) GetApproverType() string {
	if o == nil || isNil(o.ApproverType) {
		var ret string
		return ret
	}
	return *o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalSchemeForRole) GetApproverTypeOk() (*string, bool) {
	if o == nil || isNil(o.ApproverType) {
		return nil, false
	}
	return o.ApproverType, true
}

// HasApproverType returns a boolean if a field has been set.
func (o *ApprovalSchemeForRole) HasApproverType() bool {
	if o != nil && !isNil(o.ApproverType) {
		return true
	}

	return false
}

// SetApproverType gets a reference to the given string and assigns it to the ApproverType field.
func (o *ApprovalSchemeForRole) SetApproverType(v string) {
	o.ApproverType = &v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalSchemeForRole) GetApproverId() string {
	if o == nil || isNil(o.ApproverId.Get()) {
		var ret string
		return ret
	}
	return *o.ApproverId.Get()
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalSchemeForRole) GetApproverIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverId.Get(), o.ApproverId.IsSet()
}

// HasApproverId returns a boolean if a field has been set.
func (o *ApprovalSchemeForRole) HasApproverId() bool {
	if o != nil && o.ApproverId.IsSet() {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given NullableString and assigns it to the ApproverId field.
func (o *ApprovalSchemeForRole) SetApproverId(v string) {
	o.ApproverId.Set(&v)
}
// SetApproverIdNil sets the value for ApproverId to be an explicit nil
func (o *ApprovalSchemeForRole) SetApproverIdNil() {
	o.ApproverId.Set(nil)
}

// UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil
func (o *ApprovalSchemeForRole) UnsetApproverId() {
	o.ApproverId.Unset()
}

func (o ApprovalSchemeForRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalSchemeForRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApproverType) {
		toSerialize["approverType"] = o.ApproverType
	}
	if o.ApproverId.IsSet() {
		toSerialize["approverId"] = o.ApproverId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalSchemeForRole) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalSchemeForRole := _ApprovalSchemeForRole{}

	if err = json.Unmarshal(bytes, &varApprovalSchemeForRole); err == nil {
	*o = ApprovalSchemeForRole(varApprovalSchemeForRole)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "approverId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalSchemeForRole struct {
	value *ApprovalSchemeForRole
	isSet bool
}

func (v NullableApprovalSchemeForRole) Get() *ApprovalSchemeForRole {
	return v.value
}

func (v *NullableApprovalSchemeForRole) Set(val *ApprovalSchemeForRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalSchemeForRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalSchemeForRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalSchemeForRole(val *ApprovalSchemeForRole) *NullableApprovalSchemeForRole {
	return &NullableApprovalSchemeForRole{value: val, isSet: true}
}

func (v NullableApprovalSchemeForRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalSchemeForRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


