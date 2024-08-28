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

// checks if the CampaignCompleteOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignCompleteOptions{}

// CampaignCompleteOptions struct for CampaignCompleteOptions
type CampaignCompleteOptions struct {
	// Determines whether to auto-approve(APPROVE) or auto-revoke(REVOKE) upon campaign completion.
	AutoCompleteAction *string `json:"autoCompleteAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignCompleteOptions CampaignCompleteOptions

// NewCampaignCompleteOptions instantiates a new CampaignCompleteOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignCompleteOptions() *CampaignCompleteOptions {
	this := CampaignCompleteOptions{}
	var autoCompleteAction string = "APPROVE"
	this.AutoCompleteAction = &autoCompleteAction
	return &this
}

// NewCampaignCompleteOptionsWithDefaults instantiates a new CampaignCompleteOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignCompleteOptionsWithDefaults() *CampaignCompleteOptions {
	this := CampaignCompleteOptions{}
	var autoCompleteAction string = "APPROVE"
	this.AutoCompleteAction = &autoCompleteAction
	return &this
}

// GetAutoCompleteAction returns the AutoCompleteAction field value if set, zero value otherwise.
func (o *CampaignCompleteOptions) GetAutoCompleteAction() string {
	if o == nil || IsNil(o.AutoCompleteAction) {
		var ret string
		return ret
	}
	return *o.AutoCompleteAction
}

// GetAutoCompleteActionOk returns a tuple with the AutoCompleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCompleteOptions) GetAutoCompleteActionOk() (*string, bool) {
	if o == nil || IsNil(o.AutoCompleteAction) {
		return nil, false
	}
	return o.AutoCompleteAction, true
}

// HasAutoCompleteAction returns a boolean if a field has been set.
func (o *CampaignCompleteOptions) HasAutoCompleteAction() bool {
	if o != nil && !IsNil(o.AutoCompleteAction) {
		return true
	}

	return false
}

// SetAutoCompleteAction gets a reference to the given string and assigns it to the AutoCompleteAction field.
func (o *CampaignCompleteOptions) SetAutoCompleteAction(v string) {
	o.AutoCompleteAction = &v
}

func (o CampaignCompleteOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignCompleteOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoCompleteAction) {
		toSerialize["autoCompleteAction"] = o.AutoCompleteAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignCompleteOptions) UnmarshalJSON(data []byte) (err error) {
	varCampaignCompleteOptions := _CampaignCompleteOptions{}

	err = json.Unmarshal(data, &varCampaignCompleteOptions)

	if err != nil {
		return err
	}

	*o = CampaignCompleteOptions(varCampaignCompleteOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoCompleteAction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignCompleteOptions struct {
	value *CampaignCompleteOptions
	isSet bool
}

func (v NullableCampaignCompleteOptions) Get() *CampaignCompleteOptions {
	return v.value
}

func (v *NullableCampaignCompleteOptions) Set(val *CampaignCompleteOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignCompleteOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignCompleteOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignCompleteOptions(val *CampaignCompleteOptions) *NullableCampaignCompleteOptions {
	return &NullableCampaignCompleteOptions{value: val, isSet: true}
}

func (v NullableCampaignCompleteOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignCompleteOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


