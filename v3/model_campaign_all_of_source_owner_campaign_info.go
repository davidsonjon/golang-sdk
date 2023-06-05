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

// checks if the CampaignAllOfSourceOwnerCampaignInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAllOfSourceOwnerCampaignInfo{}

// CampaignAllOfSourceOwnerCampaignInfo Must be set only if the campaign type is SOURCE_OWNER.
type CampaignAllOfSourceOwnerCampaignInfo struct {
	// The list of sources to be included in the campaign.
	SourceIds []string `json:"sourceIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAllOfSourceOwnerCampaignInfo CampaignAllOfSourceOwnerCampaignInfo

// NewCampaignAllOfSourceOwnerCampaignInfo instantiates a new CampaignAllOfSourceOwnerCampaignInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOfSourceOwnerCampaignInfo() *CampaignAllOfSourceOwnerCampaignInfo {
	this := CampaignAllOfSourceOwnerCampaignInfo{}
	return &this
}

// NewCampaignAllOfSourceOwnerCampaignInfoWithDefaults instantiates a new CampaignAllOfSourceOwnerCampaignInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfSourceOwnerCampaignInfoWithDefaults() *CampaignAllOfSourceOwnerCampaignInfo {
	this := CampaignAllOfSourceOwnerCampaignInfo{}
	return &this
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise.
func (o *CampaignAllOfSourceOwnerCampaignInfo) GetSourceIds() []string {
	if o == nil || isNil(o.SourceIds) {
		var ret []string
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSourceOwnerCampaignInfo) GetSourceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SourceIds) {
		return nil, false
	}
	return o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *CampaignAllOfSourceOwnerCampaignInfo) HasSourceIds() bool {
	if o != nil && !isNil(o.SourceIds) {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []string and assigns it to the SourceIds field.
func (o *CampaignAllOfSourceOwnerCampaignInfo) SetSourceIds(v []string) {
	o.SourceIds = v
}

func (o CampaignAllOfSourceOwnerCampaignInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAllOfSourceOwnerCampaignInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceIds) {
		toSerialize["sourceIds"] = o.SourceIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAllOfSourceOwnerCampaignInfo) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignAllOfSourceOwnerCampaignInfo := _CampaignAllOfSourceOwnerCampaignInfo{}

	if err = json.Unmarshal(bytes, &varCampaignAllOfSourceOwnerCampaignInfo); err == nil {
		*o = CampaignAllOfSourceOwnerCampaignInfo(varCampaignAllOfSourceOwnerCampaignInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sourceIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAllOfSourceOwnerCampaignInfo struct {
	value *CampaignAllOfSourceOwnerCampaignInfo
	isSet bool
}

func (v NullableCampaignAllOfSourceOwnerCampaignInfo) Get() *CampaignAllOfSourceOwnerCampaignInfo {
	return v.value
}

func (v *NullableCampaignAllOfSourceOwnerCampaignInfo) Set(val *CampaignAllOfSourceOwnerCampaignInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOfSourceOwnerCampaignInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOfSourceOwnerCampaignInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOfSourceOwnerCampaignInfo(val *CampaignAllOfSourceOwnerCampaignInfo) *NullableCampaignAllOfSourceOwnerCampaignInfo {
	return &NullableCampaignAllOfSourceOwnerCampaignInfo{value: val, isSet: true}
}

func (v NullableCampaignAllOfSourceOwnerCampaignInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOfSourceOwnerCampaignInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


