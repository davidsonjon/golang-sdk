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

// checks if the ProvisioningConfigPlanInitializerScript type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConfigPlanInitializerScript{}

// ProvisioningConfigPlanInitializerScript This is a reference to a plan initializer script.
type ProvisioningConfigPlanInitializerScript struct {
	// This is a Rule that allows provisioning instruction changes.
	Source *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConfigPlanInitializerScript ProvisioningConfigPlanInitializerScript

// NewProvisioningConfigPlanInitializerScript instantiates a new ProvisioningConfigPlanInitializerScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConfigPlanInitializerScript() *ProvisioningConfigPlanInitializerScript {
	this := ProvisioningConfigPlanInitializerScript{}
	return &this
}

// NewProvisioningConfigPlanInitializerScriptWithDefaults instantiates a new ProvisioningConfigPlanInitializerScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConfigPlanInitializerScriptWithDefaults() *ProvisioningConfigPlanInitializerScript {
	this := ProvisioningConfigPlanInitializerScript{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProvisioningConfigPlanInitializerScript) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfigPlanInitializerScript) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProvisioningConfigPlanInitializerScript) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProvisioningConfigPlanInitializerScript) SetSource(v string) {
	o.Source = &v
}

func (o ProvisioningConfigPlanInitializerScript) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConfigPlanInitializerScript) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConfigPlanInitializerScript) UnmarshalJSON(data []byte) (err error) {
	varProvisioningConfigPlanInitializerScript := _ProvisioningConfigPlanInitializerScript{}

	err = json.Unmarshal(data, &varProvisioningConfigPlanInitializerScript)

	if err != nil {
		return err
	}

	*o = ProvisioningConfigPlanInitializerScript(varProvisioningConfigPlanInitializerScript)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConfigPlanInitializerScript struct {
	value *ProvisioningConfigPlanInitializerScript
	isSet bool
}

func (v NullableProvisioningConfigPlanInitializerScript) Get() *ProvisioningConfigPlanInitializerScript {
	return v.value
}

func (v *NullableProvisioningConfigPlanInitializerScript) Set(val *ProvisioningConfigPlanInitializerScript) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConfigPlanInitializerScript) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConfigPlanInitializerScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConfigPlanInitializerScript(val *ProvisioningConfigPlanInitializerScript) *NullableProvisioningConfigPlanInitializerScript {
	return &NullableProvisioningConfigPlanInitializerScript{value: val, isSet: true}
}

func (v NullableProvisioningConfigPlanInitializerScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConfigPlanInitializerScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


