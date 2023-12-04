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

// checks if the CampaignsDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignsDeleteRequest{}

// CampaignsDeleteRequest struct for CampaignsDeleteRequest
type CampaignsDeleteRequest struct {
	// The ids of the campaigns to delete
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignsDeleteRequest CampaignsDeleteRequest

// NewCampaignsDeleteRequest instantiates a new CampaignsDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignsDeleteRequest() *CampaignsDeleteRequest {
	this := CampaignsDeleteRequest{}
	return &this
}

// NewCampaignsDeleteRequestWithDefaults instantiates a new CampaignsDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignsDeleteRequestWithDefaults() *CampaignsDeleteRequest {
	this := CampaignsDeleteRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CampaignsDeleteRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignsDeleteRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CampaignsDeleteRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CampaignsDeleteRequest) SetIds(v []string) {
	o.Ids = v
}

func (o CampaignsDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignsDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignsDeleteRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignsDeleteRequest := _CampaignsDeleteRequest{}

	if err = json.Unmarshal(bytes, &varCampaignsDeleteRequest); err == nil {
	*o = CampaignsDeleteRequest(varCampaignsDeleteRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignsDeleteRequest struct {
	value *CampaignsDeleteRequest
	isSet bool
}

func (v NullableCampaignsDeleteRequest) Get() *CampaignsDeleteRequest {
	return v.value
}

func (v *NullableCampaignsDeleteRequest) Set(val *CampaignsDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignsDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignsDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignsDeleteRequest(val *CampaignsDeleteRequest) *NullableCampaignsDeleteRequest {
	return &NullableCampaignsDeleteRequest{value: val, isSet: true}
}

func (v NullableCampaignsDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignsDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


