/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the CampaignReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignReference{}

// CampaignReference struct for CampaignReference
type CampaignReference struct {
	// The unique ID of the campaign.
	Id string `json:"id"`
	// The name of the campaign.
	Name string `json:"name"`
	// The type of object that is being referenced.
	Type string `json:"type"`
	// The type of the campaign.
	CampaignType string `json:"campaignType"`
	// The description of the campaign set by the admin who created it.
	Description NullableString `json:"description"`
	// The correlatedStatus of the campaign. Only SOURCE_OWNER campaigns can be Uncorrelated. An Uncorrelated certification campaign only includes Uncorrelated identities (An identity is uncorrelated if it has no accounts on an authoritative source).
	CorrelatedStatus map[string]interface{} `json:"correlatedStatus"`
	// Determines whether comments are required for decisions during certification reviews. You can require comments for all decisions, revoke-only decisions, or no decisions. By default, comments are not required for decisions.
	MandatoryCommentRequirement string `json:"mandatoryCommentRequirement"`
	AdditionalProperties map[string]interface{}
}

type _CampaignReference CampaignReference

// NewCampaignReference instantiates a new CampaignReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReference(id string, name string, type_ string, campaignType string, description NullableString, correlatedStatus map[string]interface{}, mandatoryCommentRequirement string) *CampaignReference {
	this := CampaignReference{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.CampaignType = campaignType
	this.Description = description
	this.CorrelatedStatus = correlatedStatus
	this.MandatoryCommentRequirement = mandatoryCommentRequirement
	return &this
}

// NewCampaignReferenceWithDefaults instantiates a new CampaignReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReferenceWithDefaults() *CampaignReference {
	this := CampaignReference{}
	return &this
}

// GetId returns the Id field value
func (o *CampaignReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignReference) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CampaignReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignReference) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CampaignReference) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignReference) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignReference) SetType(v string) {
	o.Type = v
}

// GetCampaignType returns the CampaignType field value
func (o *CampaignReference) GetCampaignType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignType
}

// GetCampaignTypeOk returns a tuple with the CampaignType field value
// and a boolean to check if the value has been set.
func (o *CampaignReference) GetCampaignTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignType, true
}

// SetCampaignType sets field value
func (o *CampaignReference) SetCampaignType(v string) {
	o.CampaignType = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CampaignReference) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignReference) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *CampaignReference) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCorrelatedStatus returns the CorrelatedStatus field value
func (o *CampaignReference) GetCorrelatedStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CorrelatedStatus
}

// GetCorrelatedStatusOk returns a tuple with the CorrelatedStatus field value
// and a boolean to check if the value has been set.
func (o *CampaignReference) GetCorrelatedStatusOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CorrelatedStatus, true
}

// SetCorrelatedStatus sets field value
func (o *CampaignReference) SetCorrelatedStatus(v map[string]interface{}) {
	o.CorrelatedStatus = v
}

// GetMandatoryCommentRequirement returns the MandatoryCommentRequirement field value
func (o *CampaignReference) GetMandatoryCommentRequirement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandatoryCommentRequirement
}

// GetMandatoryCommentRequirementOk returns a tuple with the MandatoryCommentRequirement field value
// and a boolean to check if the value has been set.
func (o *CampaignReference) GetMandatoryCommentRequirementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandatoryCommentRequirement, true
}

// SetMandatoryCommentRequirement sets field value
func (o *CampaignReference) SetMandatoryCommentRequirement(v string) {
	o.MandatoryCommentRequirement = v
}

func (o CampaignReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["campaignType"] = o.CampaignType
	toSerialize["description"] = o.Description.Get()
	toSerialize["correlatedStatus"] = o.CorrelatedStatus
	toSerialize["mandatoryCommentRequirement"] = o.MandatoryCommentRequirement

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignReference) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"campaignType",
		"description",
		"correlatedStatus",
		"mandatoryCommentRequirement",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCampaignReference := _CampaignReference{}

	if err = json.Unmarshal(bytes, &varCampaignReference); err == nil {
	*o = CampaignReference(varCampaignReference)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "campaignType")
		delete(additionalProperties, "description")
		delete(additionalProperties, "correlatedStatus")
		delete(additionalProperties, "mandatoryCommentRequirement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignReference struct {
	value *CampaignReference
	isSet bool
}

func (v NullableCampaignReference) Get() *CampaignReference {
	return v.value
}

func (v *NullableCampaignReference) Set(val *CampaignReference) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReference) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReference(val *CampaignReference) *NullableCampaignReference {
	return &NullableCampaignReference{value: val, isSet: true}
}

func (v NullableCampaignReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


