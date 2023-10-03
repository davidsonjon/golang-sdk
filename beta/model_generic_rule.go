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

// checks if the GenericRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericRule{}

// GenericRule struct for GenericRule
type GenericRule struct {
	// This is the name of the Generic rule that needs to be invoked by the transform
	Name string `json:"name"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericRule GenericRule

// NewGenericRule instantiates a new GenericRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericRule(name string) *GenericRule {
	this := GenericRule{}
	this.Name = name
	return &this
}

// NewGenericRuleWithDefaults instantiates a new GenericRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericRuleWithDefaults() *GenericRule {
	this := GenericRule{}
	return &this
}

// GetName returns the Name field value
func (o *GenericRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenericRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenericRule) SetName(v string) {
	o.Name = v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *GenericRule) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericRule) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *GenericRule) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *GenericRule) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

func (o GenericRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenericRule) UnmarshalJSON(bytes []byte) (err error) {
	varGenericRule := _GenericRule{}

	if err = json.Unmarshal(bytes, &varGenericRule); err == nil {
		*o = GenericRule(varGenericRule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "requiresPeriodicRefresh")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericRule struct {
	value *GenericRule
	isSet bool
}

func (v NullableGenericRule) Get() *GenericRule {
	return v.value
}

func (v *NullableGenericRule) Set(val *GenericRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericRule(val *GenericRule) *NullableGenericRule {
	return &NullableGenericRule{value: val, isSet: true}
}

func (v NullableGenericRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


