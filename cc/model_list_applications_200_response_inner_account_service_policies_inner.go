/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the ListApplications200ResponseInnerAccountServicePoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplications200ResponseInnerAccountServicePoliciesInner{}

// ListApplications200ResponseInnerAccountServicePoliciesInner struct for ListApplications200ResponseInnerAccountServicePoliciesInner
type ListApplications200ResponseInnerAccountServicePoliciesInner struct {
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	Selectors map[string]interface{} `json:"selectors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListApplications200ResponseInnerAccountServicePoliciesInner ListApplications200ResponseInnerAccountServicePoliciesInner

// NewListApplications200ResponseInnerAccountServicePoliciesInner instantiates a new ListApplications200ResponseInnerAccountServicePoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplications200ResponseInnerAccountServicePoliciesInner() *ListApplications200ResponseInnerAccountServicePoliciesInner {
	this := ListApplications200ResponseInnerAccountServicePoliciesInner{}
	return &this
}

// NewListApplications200ResponseInnerAccountServicePoliciesInnerWithDefaults instantiates a new ListApplications200ResponseInnerAccountServicePoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplications200ResponseInnerAccountServicePoliciesInnerWithDefaults() *ListApplications200ResponseInnerAccountServicePoliciesInner {
	this := ListApplications200ResponseInnerAccountServicePoliciesInner{}
	return &this
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) GetPolicyId() string {
	if o == nil || isNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) GetPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) HasPolicyId() bool {
	if o != nil && !isNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) GetPolicyName() string {
	if o == nil || isNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) GetPolicyNameOk() (*string, bool) {
	if o == nil || isNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) HasPolicyName() bool {
	if o != nil && !isNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) GetSelectors() map[string]interface{} {
	if o == nil || isNil(o.Selectors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) GetSelectorsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Selectors) {
		return map[string]interface{}{}, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) HasSelectors() bool {
	if o != nil && !isNil(o.Selectors) {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given map[string]interface{} and assigns it to the Selectors field.
func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) SetSelectors(v map[string]interface{}) {
	o.Selectors = v
}

func (o ListApplications200ResponseInnerAccountServicePoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplications200ResponseInnerAccountServicePoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !isNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	if !isNil(o.Selectors) {
		toSerialize["selectors"] = o.Selectors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListApplications200ResponseInnerAccountServicePoliciesInner) UnmarshalJSON(bytes []byte) (err error) {
	varListApplications200ResponseInnerAccountServicePoliciesInner := _ListApplications200ResponseInnerAccountServicePoliciesInner{}

	if err = json.Unmarshal(bytes, &varListApplications200ResponseInnerAccountServicePoliciesInner); err == nil {
	*o = ListApplications200ResponseInnerAccountServicePoliciesInner(varListApplications200ResponseInnerAccountServicePoliciesInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "policyId")
		delete(additionalProperties, "policyName")
		delete(additionalProperties, "selectors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListApplications200ResponseInnerAccountServicePoliciesInner struct {
	value *ListApplications200ResponseInnerAccountServicePoliciesInner
	isSet bool
}

func (v NullableListApplications200ResponseInnerAccountServicePoliciesInner) Get() *ListApplications200ResponseInnerAccountServicePoliciesInner {
	return v.value
}

func (v *NullableListApplications200ResponseInnerAccountServicePoliciesInner) Set(val *ListApplications200ResponseInnerAccountServicePoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInnerAccountServicePoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInnerAccountServicePoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInnerAccountServicePoliciesInner(val *ListApplications200ResponseInnerAccountServicePoliciesInner) *NullableListApplications200ResponseInnerAccountServicePoliciesInner {
	return &NullableListApplications200ResponseInnerAccountServicePoliciesInner{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInnerAccountServicePoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInnerAccountServicePoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


