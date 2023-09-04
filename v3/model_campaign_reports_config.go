/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the CampaignReportsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignReportsConfig{}

// CampaignReportsConfig struct for CampaignReportsConfig
type CampaignReportsConfig struct {
	// list of identity attribute columns
	IdentityAttributeColumns []string `json:"identityAttributeColumns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignReportsConfig CampaignReportsConfig

// NewCampaignReportsConfig instantiates a new CampaignReportsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReportsConfig() *CampaignReportsConfig {
	this := CampaignReportsConfig{}
	return &this
}

// NewCampaignReportsConfigWithDefaults instantiates a new CampaignReportsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReportsConfigWithDefaults() *CampaignReportsConfig {
	this := CampaignReportsConfig{}
	return &this
}

// GetIdentityAttributeColumns returns the IdentityAttributeColumns field value if set, zero value otherwise.
func (o *CampaignReportsConfig) GetIdentityAttributeColumns() []string {
	if o == nil || isNil(o.IdentityAttributeColumns) {
		var ret []string
		return ret
	}
	return o.IdentityAttributeColumns
}

// GetIdentityAttributeColumnsOk returns a tuple with the IdentityAttributeColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportsConfig) GetIdentityAttributeColumnsOk() ([]string, bool) {
	if o == nil || isNil(o.IdentityAttributeColumns) {
		return nil, false
	}
	return o.IdentityAttributeColumns, true
}

// HasIdentityAttributeColumns returns a boolean if a field has been set.
func (o *CampaignReportsConfig) HasIdentityAttributeColumns() bool {
	if o != nil && !isNil(o.IdentityAttributeColumns) {
		return true
	}

	return false
}

// SetIdentityAttributeColumns gets a reference to the given []string and assigns it to the IdentityAttributeColumns field.
func (o *CampaignReportsConfig) SetIdentityAttributeColumns(v []string) {
	o.IdentityAttributeColumns = v
}

func (o CampaignReportsConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignReportsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityAttributeColumns) {
		toSerialize["identityAttributeColumns"] = o.IdentityAttributeColumns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignReportsConfig) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignReportsConfig := _CampaignReportsConfig{}

	if err = json.Unmarshal(bytes, &varCampaignReportsConfig); err == nil {
		*o = CampaignReportsConfig(varCampaignReportsConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityAttributeColumns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignReportsConfig struct {
	value *CampaignReportsConfig
	isSet bool
}

func (v NullableCampaignReportsConfig) Get() *CampaignReportsConfig {
	return v.value
}

func (v *NullableCampaignReportsConfig) Set(val *CampaignReportsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReportsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReportsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReportsConfig(val *CampaignReportsConfig) *NullableCampaignReportsConfig {
	return &NullableCampaignReportsConfig{value: val, isSet: true}
}

func (v NullableCampaignReportsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReportsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


