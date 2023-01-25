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

// TriggerInputCampaignEnded struct for TriggerInputCampaignEnded
type TriggerInputCampaignEnded struct {
	Campaign TriggerInputCampaignEndedCampaign `json:"campaign"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputCampaignEnded TriggerInputCampaignEnded

// NewTriggerInputCampaignEnded instantiates a new TriggerInputCampaignEnded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputCampaignEnded(campaign TriggerInputCampaignEndedCampaign) *TriggerInputCampaignEnded {
	this := TriggerInputCampaignEnded{}
	this.Campaign = campaign
	return &this
}

// NewTriggerInputCampaignEndedWithDefaults instantiates a new TriggerInputCampaignEnded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputCampaignEndedWithDefaults() *TriggerInputCampaignEnded {
	this := TriggerInputCampaignEnded{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *TriggerInputCampaignEnded) GetCampaign() TriggerInputCampaignEndedCampaign {
	if o == nil {
		var ret TriggerInputCampaignEndedCampaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignEnded) GetCampaignOk() (*TriggerInputCampaignEndedCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *TriggerInputCampaignEnded) SetCampaign(v TriggerInputCampaignEndedCampaign) {
	o.Campaign = v
}

func (o TriggerInputCampaignEnded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaign"] = o.Campaign
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputCampaignEnded) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputCampaignEnded := _TriggerInputCampaignEnded{}

	if err = json.Unmarshal(bytes, &varTriggerInputCampaignEnded); err == nil {
		*o = TriggerInputCampaignEnded(varTriggerInputCampaignEnded)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputCampaignEnded struct {
	value *TriggerInputCampaignEnded
	isSet bool
}

func (v NullableTriggerInputCampaignEnded) Get() *TriggerInputCampaignEnded {
	return v.value
}

func (v *NullableTriggerInputCampaignEnded) Set(val *TriggerInputCampaignEnded) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputCampaignEnded) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputCampaignEnded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputCampaignEnded(val *TriggerInputCampaignEnded) *NullableTriggerInputCampaignEnded {
	return &NullableTriggerInputCampaignEnded{value: val, isSet: true}
}

func (v NullableTriggerInputCampaignEnded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputCampaignEnded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


