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

// checks if the RoleMiningPotentialRoleEditEntitlements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRoleEditEntitlements{}

// RoleMiningPotentialRoleEditEntitlements struct for RoleMiningPotentialRoleEditEntitlements
type RoleMiningPotentialRoleEditEntitlements struct {
	// The list of entitlement ids to be edited
	Ids []string `json:"ids,omitempty"`
	// If true, add ids to be exclusion list. If false, remove ids from the exclusion list.
	Exclude *bool `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleEditEntitlements RoleMiningPotentialRoleEditEntitlements

// NewRoleMiningPotentialRoleEditEntitlements instantiates a new RoleMiningPotentialRoleEditEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleEditEntitlements() *RoleMiningPotentialRoleEditEntitlements {
	this := RoleMiningPotentialRoleEditEntitlements{}
	return &this
}

// NewRoleMiningPotentialRoleEditEntitlementsWithDefaults instantiates a new RoleMiningPotentialRoleEditEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleEditEntitlementsWithDefaults() *RoleMiningPotentialRoleEditEntitlements {
	this := RoleMiningPotentialRoleEditEntitlements{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleEditEntitlements) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleEditEntitlements) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleEditEntitlements) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *RoleMiningPotentialRoleEditEntitlements) SetIds(v []string) {
	o.Ids = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleEditEntitlements) GetExclude() bool {
	if o == nil || isNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleEditEntitlements) GetExcludeOk() (*bool, bool) {
	if o == nil || isNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleEditEntitlements) HasExclude() bool {
	if o != nil && !isNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *RoleMiningPotentialRoleEditEntitlements) SetExclude(v bool) {
	o.Exclude = &v
}

func (o RoleMiningPotentialRoleEditEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRoleEditEntitlements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRoleEditEntitlements) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRoleEditEntitlements := _RoleMiningPotentialRoleEditEntitlements{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRoleEditEntitlements); err == nil {
	*o = RoleMiningPotentialRoleEditEntitlements(varRoleMiningPotentialRoleEditEntitlements)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleEditEntitlements struct {
	value *RoleMiningPotentialRoleEditEntitlements
	isSet bool
}

func (v NullableRoleMiningPotentialRoleEditEntitlements) Get() *RoleMiningPotentialRoleEditEntitlements {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleEditEntitlements) Set(val *RoleMiningPotentialRoleEditEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleEditEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleEditEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleEditEntitlements(val *RoleMiningPotentialRoleEditEntitlements) *NullableRoleMiningPotentialRoleEditEntitlements {
	return &NullableRoleMiningPotentialRoleEditEntitlements{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleEditEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleEditEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


