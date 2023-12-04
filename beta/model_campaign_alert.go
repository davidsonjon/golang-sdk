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

// checks if the CampaignAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAlert{}

// CampaignAlert struct for CampaignAlert
type CampaignAlert struct {
	// Denotes the level of the message
	Level *string `json:"level,omitempty"`
	Localizations []ErrorMessageDto `json:"localizations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAlert CampaignAlert

// NewCampaignAlert instantiates a new CampaignAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAlert() *CampaignAlert {
	this := CampaignAlert{}
	return &this
}

// NewCampaignAlertWithDefaults instantiates a new CampaignAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAlertWithDefaults() *CampaignAlert {
	this := CampaignAlert{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *CampaignAlert) GetLevel() string {
	if o == nil || isNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAlert) GetLevelOk() (*string, bool) {
	if o == nil || isNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *CampaignAlert) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *CampaignAlert) SetLevel(v string) {
	o.Level = &v
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *CampaignAlert) GetLocalizations() []ErrorMessageDto {
	if o == nil || isNil(o.Localizations) {
		var ret []ErrorMessageDto
		return ret
	}
	return o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAlert) GetLocalizationsOk() ([]ErrorMessageDto, bool) {
	if o == nil || isNil(o.Localizations) {
		return nil, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *CampaignAlert) HasLocalizations() bool {
	if o != nil && !isNil(o.Localizations) {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given []ErrorMessageDto and assigns it to the Localizations field.
func (o *CampaignAlert) SetLocalizations(v []ErrorMessageDto) {
	o.Localizations = v
}

func (o CampaignAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !isNil(o.Localizations) {
		toSerialize["localizations"] = o.Localizations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAlert) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignAlert := _CampaignAlert{}

	if err = json.Unmarshal(bytes, &varCampaignAlert); err == nil {
	*o = CampaignAlert(varCampaignAlert)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "level")
		delete(additionalProperties, "localizations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAlert struct {
	value *CampaignAlert
	isSet bool
}

func (v NullableCampaignAlert) Get() *CampaignAlert {
	return v.value
}

func (v *NullableCampaignAlert) Set(val *CampaignAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAlert(val *CampaignAlert) *NullableCampaignAlert {
	return &NullableCampaignAlert{value: val, isSet: true}
}

func (v NullableCampaignAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


