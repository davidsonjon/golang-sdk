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

// RoleMembershipSelector When present, specifies that the Role is to be granted to Identities which either satisfy specific criteria or which are members of a given list of Identities.
type RoleMembershipSelector struct {
	Type *RoleMembershipSelectorType `json:"type,omitempty"`
	Criteria NullableRoleCriteriaLevel1 `json:"criteria,omitempty"`
	// Defines role membership as being exclusive to the specified Identities, when type is IDENTITY_LIST.
	Identities []RoleMembershipIdentity `json:"identities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMembershipSelector RoleMembershipSelector

// NewRoleMembershipSelector instantiates a new RoleMembershipSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMembershipSelector() *RoleMembershipSelector {
	this := RoleMembershipSelector{}
	return &this
}

// NewRoleMembershipSelectorWithDefaults instantiates a new RoleMembershipSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMembershipSelectorWithDefaults() *RoleMembershipSelector {
	this := RoleMembershipSelector{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleMembershipSelector) GetType() RoleMembershipSelectorType {
	if o == nil || isNil(o.Type) {
		var ret RoleMembershipSelectorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMembershipSelector) GetTypeOk() (*RoleMembershipSelectorType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleMembershipSelector) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RoleMembershipSelectorType and assigns it to the Type field.
func (o *RoleMembershipSelector) SetType(v RoleMembershipSelectorType) {
	o.Type = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMembershipSelector) GetCriteria() RoleCriteriaLevel1 {
	if o == nil || isNil(o.Criteria.Get()) {
		var ret RoleCriteriaLevel1
		return ret
	}
	return *o.Criteria.Get()
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMembershipSelector) GetCriteriaOk() (*RoleCriteriaLevel1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria.Get(), o.Criteria.IsSet()
}

// HasCriteria returns a boolean if a field has been set.
func (o *RoleMembershipSelector) HasCriteria() bool {
	if o != nil && o.Criteria.IsSet() {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given NullableRoleCriteriaLevel1 and assigns it to the Criteria field.
func (o *RoleMembershipSelector) SetCriteria(v RoleCriteriaLevel1) {
	o.Criteria.Set(&v)
}
// SetCriteriaNil sets the value for Criteria to be an explicit nil
func (o *RoleMembershipSelector) SetCriteriaNil() {
	o.Criteria.Set(nil)
}

// UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil
func (o *RoleMembershipSelector) UnsetCriteria() {
	o.Criteria.Unset()
}

// GetIdentities returns the Identities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMembershipSelector) GetIdentities() []RoleMembershipIdentity {
	if o == nil {
		var ret []RoleMembershipIdentity
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMembershipSelector) GetIdentitiesOk() ([]RoleMembershipIdentity, bool) {
	if o == nil || isNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *RoleMembershipSelector) HasIdentities() bool {
	if o != nil && isNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []RoleMembershipIdentity and assigns it to the Identities field.
func (o *RoleMembershipSelector) SetIdentities(v []RoleMembershipIdentity) {
	o.Identities = v
}

func (o RoleMembershipSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Criteria.IsSet() {
		toSerialize["criteria"] = o.Criteria.Get()
	}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleMembershipSelector) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMembershipSelector := _RoleMembershipSelector{}

	if err = json.Unmarshal(bytes, &varRoleMembershipSelector); err == nil {
		*o = RoleMembershipSelector(varRoleMembershipSelector)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "criteria")
		delete(additionalProperties, "identities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMembershipSelector struct {
	value *RoleMembershipSelector
	isSet bool
}

func (v NullableRoleMembershipSelector) Get() *RoleMembershipSelector {
	return v.value
}

func (v *NullableRoleMembershipSelector) Set(val *RoleMembershipSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMembershipSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMembershipSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMembershipSelector(val *RoleMembershipSelector) *NullableRoleMembershipSelector {
	return &NullableRoleMembershipSelector{value: val, isSet: true}
}

func (v NullableRoleMembershipSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMembershipSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


