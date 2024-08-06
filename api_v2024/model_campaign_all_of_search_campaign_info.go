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

// checks if the CampaignAllOfSearchCampaignInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAllOfSearchCampaignInfo{}

// CampaignAllOfSearchCampaignInfo Must be set only if the campaign type is SEARCH.
type CampaignAllOfSearchCampaignInfo struct {
	// The type of search campaign represented.
	Type string `json:"type"`
	// Describes this search campaign. Intended for storing the query used, and possibly the number of identities selected/available.
	Description *string `json:"description,omitempty"`
	Reviewer *CampaignAllOfSearchCampaignInfoReviewer `json:"reviewer,omitempty"`
	// The scope for the campaign. The campaign will cover identities returned by the query and identities that have access items returned by the query. One of `query` or `identityIds` must be set.
	Query *string `json:"query,omitempty"`
	// A direct list of identities to include in this campaign. One of `identityIds` or `query` must be set.
	IdentityIds []string `json:"identityIds,omitempty"`
	// Further reduces the scope of the campaign by excluding identities (from `query` or `identityIds`) that do not have this access.
	AccessConstraints []AccessConstraint `json:"accessConstraints,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAllOfSearchCampaignInfo CampaignAllOfSearchCampaignInfo

// NewCampaignAllOfSearchCampaignInfo instantiates a new CampaignAllOfSearchCampaignInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOfSearchCampaignInfo(type_ string) *CampaignAllOfSearchCampaignInfo {
	this := CampaignAllOfSearchCampaignInfo{}
	this.Type = type_
	return &this
}

// NewCampaignAllOfSearchCampaignInfoWithDefaults instantiates a new CampaignAllOfSearchCampaignInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfSearchCampaignInfoWithDefaults() *CampaignAllOfSearchCampaignInfo {
	this := CampaignAllOfSearchCampaignInfo{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignAllOfSearchCampaignInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSearchCampaignInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignAllOfSearchCampaignInfo) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignAllOfSearchCampaignInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSearchCampaignInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignAllOfSearchCampaignInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignAllOfSearchCampaignInfo) SetDescription(v string) {
	o.Description = &v
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *CampaignAllOfSearchCampaignInfo) GetReviewer() CampaignAllOfSearchCampaignInfoReviewer {
	if o == nil || isNil(o.Reviewer) {
		var ret CampaignAllOfSearchCampaignInfoReviewer
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSearchCampaignInfo) GetReviewerOk() (*CampaignAllOfSearchCampaignInfoReviewer, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *CampaignAllOfSearchCampaignInfo) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given CampaignAllOfSearchCampaignInfoReviewer and assigns it to the Reviewer field.
func (o *CampaignAllOfSearchCampaignInfo) SetReviewer(v CampaignAllOfSearchCampaignInfoReviewer) {
	o.Reviewer = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CampaignAllOfSearchCampaignInfo) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSearchCampaignInfo) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CampaignAllOfSearchCampaignInfo) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CampaignAllOfSearchCampaignInfo) SetQuery(v string) {
	o.Query = &v
}

// GetIdentityIds returns the IdentityIds field value if set, zero value otherwise.
func (o *CampaignAllOfSearchCampaignInfo) GetIdentityIds() []string {
	if o == nil || isNil(o.IdentityIds) {
		var ret []string
		return ret
	}
	return o.IdentityIds
}

// GetIdentityIdsOk returns a tuple with the IdentityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSearchCampaignInfo) GetIdentityIdsOk() ([]string, bool) {
	if o == nil || isNil(o.IdentityIds) {
		return nil, false
	}
	return o.IdentityIds, true
}

// HasIdentityIds returns a boolean if a field has been set.
func (o *CampaignAllOfSearchCampaignInfo) HasIdentityIds() bool {
	if o != nil && !isNil(o.IdentityIds) {
		return true
	}

	return false
}

// SetIdentityIds gets a reference to the given []string and assigns it to the IdentityIds field.
func (o *CampaignAllOfSearchCampaignInfo) SetIdentityIds(v []string) {
	o.IdentityIds = v
}

// GetAccessConstraints returns the AccessConstraints field value if set, zero value otherwise.
func (o *CampaignAllOfSearchCampaignInfo) GetAccessConstraints() []AccessConstraint {
	if o == nil || isNil(o.AccessConstraints) {
		var ret []AccessConstraint
		return ret
	}
	return o.AccessConstraints
}

// GetAccessConstraintsOk returns a tuple with the AccessConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfSearchCampaignInfo) GetAccessConstraintsOk() ([]AccessConstraint, bool) {
	if o == nil || isNil(o.AccessConstraints) {
		return nil, false
	}
	return o.AccessConstraints, true
}

// HasAccessConstraints returns a boolean if a field has been set.
func (o *CampaignAllOfSearchCampaignInfo) HasAccessConstraints() bool {
	if o != nil && !isNil(o.AccessConstraints) {
		return true
	}

	return false
}

// SetAccessConstraints gets a reference to the given []AccessConstraint and assigns it to the AccessConstraints field.
func (o *CampaignAllOfSearchCampaignInfo) SetAccessConstraints(v []AccessConstraint) {
	o.AccessConstraints = v
}

func (o CampaignAllOfSearchCampaignInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAllOfSearchCampaignInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.IdentityIds) {
		toSerialize["identityIds"] = o.IdentityIds
	}
	if !isNil(o.AccessConstraints) {
		toSerialize["accessConstraints"] = o.AccessConstraints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAllOfSearchCampaignInfo) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCampaignAllOfSearchCampaignInfo := _CampaignAllOfSearchCampaignInfo{}

	if err = json.Unmarshal(bytes, &varCampaignAllOfSearchCampaignInfo); err == nil {
	*o = CampaignAllOfSearchCampaignInfo(varCampaignAllOfSearchCampaignInfo)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "query")
		delete(additionalProperties, "identityIds")
		delete(additionalProperties, "accessConstraints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAllOfSearchCampaignInfo struct {
	value *CampaignAllOfSearchCampaignInfo
	isSet bool
}

func (v NullableCampaignAllOfSearchCampaignInfo) Get() *CampaignAllOfSearchCampaignInfo {
	return v.value
}

func (v *NullableCampaignAllOfSearchCampaignInfo) Set(val *CampaignAllOfSearchCampaignInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOfSearchCampaignInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOfSearchCampaignInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOfSearchCampaignInfo(val *CampaignAllOfSearchCampaignInfo) *NullableCampaignAllOfSearchCampaignInfo {
	return &NullableCampaignAllOfSearchCampaignInfo{value: val, isSet: true}
}

func (v NullableCampaignAllOfSearchCampaignInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOfSearchCampaignInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


