/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the CampaignGenerated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignGenerated{}

// CampaignGenerated struct for CampaignGenerated
type CampaignGenerated struct {
	Campaign CampaignGeneratedCampaign `json:"campaign"`
	AdditionalProperties map[string]interface{}
}

type _CampaignGenerated CampaignGenerated

// NewCampaignGenerated instantiates a new CampaignGenerated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignGenerated(campaign CampaignGeneratedCampaign) *CampaignGenerated {
	this := CampaignGenerated{}
	this.Campaign = campaign
	return &this
}

// NewCampaignGeneratedWithDefaults instantiates a new CampaignGenerated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignGeneratedWithDefaults() *CampaignGenerated {
	this := CampaignGenerated{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *CampaignGenerated) GetCampaign() CampaignGeneratedCampaign {
	if o == nil {
		var ret CampaignGeneratedCampaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *CampaignGenerated) GetCampaignOk() (*CampaignGeneratedCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *CampaignGenerated) SetCampaign(v CampaignGeneratedCampaign) {
	o.Campaign = v
}

func (o CampaignGenerated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignGenerated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignGenerated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaign",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCampaignGenerated := _CampaignGenerated{}

	err = json.Unmarshal(data, &varCampaignGenerated)

	if err != nil {
		return err
	}

	*o = CampaignGenerated(varCampaignGenerated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignGenerated struct {
	value *CampaignGenerated
	isSet bool
}

func (v NullableCampaignGenerated) Get() *CampaignGenerated {
	return v.value
}

func (v *NullableCampaignGenerated) Set(val *CampaignGenerated) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignGenerated) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignGenerated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignGenerated(val *CampaignGenerated) *NullableCampaignGenerated {
	return &NullableCampaignGenerated{value: val, isSet: true}
}

func (v NullableCampaignGenerated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignGenerated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


