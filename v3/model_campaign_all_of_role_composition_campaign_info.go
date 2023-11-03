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

// checks if the CampaignAllOfRoleCompositionCampaignInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAllOfRoleCompositionCampaignInfo{}

// CampaignAllOfRoleCompositionCampaignInfo Optional configuration options for role composition campaigns.
type CampaignAllOfRoleCompositionCampaignInfo struct {
	Reviewer *CampaignAllOfSearchCampaignInfoReviewer `json:"reviewer,omitempty"`
	// Optional list of roles to include in this campaign. Only one of `roleIds` and `query` may be set; if neither are set, all roles are included.
	RoleIds []string `json:"roleIds,omitempty"`
	RemediatorRef CampaignAllOfRoleCompositionCampaignInfoRemediatorRef `json:"remediatorRef"`
	// Optional search query to scope this campaign to a set of roles. Only one of `roleIds` and `query` may be set; if neither are set, all roles are included.
	Query *string `json:"query,omitempty"`
	// Describes this role composition campaign. Intended for storing the query used, and possibly the number of roles selected/available.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAllOfRoleCompositionCampaignInfo CampaignAllOfRoleCompositionCampaignInfo

// NewCampaignAllOfRoleCompositionCampaignInfo instantiates a new CampaignAllOfRoleCompositionCampaignInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOfRoleCompositionCampaignInfo(remediatorRef CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) *CampaignAllOfRoleCompositionCampaignInfo {
	this := CampaignAllOfRoleCompositionCampaignInfo{}
	this.RemediatorRef = remediatorRef
	return &this
}

// NewCampaignAllOfRoleCompositionCampaignInfoWithDefaults instantiates a new CampaignAllOfRoleCompositionCampaignInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfRoleCompositionCampaignInfoWithDefaults() *CampaignAllOfRoleCompositionCampaignInfo {
	this := CampaignAllOfRoleCompositionCampaignInfo{}
	return &this
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetReviewer() CampaignAllOfSearchCampaignInfoReviewer {
	if o == nil || isNil(o.Reviewer) {
		var ret CampaignAllOfSearchCampaignInfoReviewer
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetReviewerOk() (*CampaignAllOfSearchCampaignInfoReviewer, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given CampaignAllOfSearchCampaignInfoReviewer and assigns it to the Reviewer field.
func (o *CampaignAllOfRoleCompositionCampaignInfo) SetReviewer(v CampaignAllOfSearchCampaignInfoReviewer) {
	o.Reviewer = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetRoleIds() []string {
	if o == nil || isNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetRoleIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *CampaignAllOfRoleCompositionCampaignInfo) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetRemediatorRef returns the RemediatorRef field value
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetRemediatorRef() CampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	if o == nil {
		var ret CampaignAllOfRoleCompositionCampaignInfoRemediatorRef
		return ret
	}

	return o.RemediatorRef
}

// GetRemediatorRefOk returns a tuple with the RemediatorRef field value
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetRemediatorRefOk() (*CampaignAllOfRoleCompositionCampaignInfoRemediatorRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediatorRef, true
}

// SetRemediatorRef sets field value
func (o *CampaignAllOfRoleCompositionCampaignInfo) SetRemediatorRef(v CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) {
	o.RemediatorRef = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CampaignAllOfRoleCompositionCampaignInfo) SetQuery(v string) {
	o.Query = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignAllOfRoleCompositionCampaignInfo) SetDescription(v string) {
	o.Description = &v
}

func (o CampaignAllOfRoleCompositionCampaignInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAllOfRoleCompositionCampaignInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !isNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	toSerialize["remediatorRef"] = o.RemediatorRef
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAllOfRoleCompositionCampaignInfo) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignAllOfRoleCompositionCampaignInfo := _CampaignAllOfRoleCompositionCampaignInfo{}

	if err = json.Unmarshal(bytes, &varCampaignAllOfRoleCompositionCampaignInfo); err == nil {
		*o = CampaignAllOfRoleCompositionCampaignInfo(varCampaignAllOfRoleCompositionCampaignInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "roleIds")
		delete(additionalProperties, "remediatorRef")
		delete(additionalProperties, "query")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAllOfRoleCompositionCampaignInfo struct {
	value *CampaignAllOfRoleCompositionCampaignInfo
	isSet bool
}

func (v NullableCampaignAllOfRoleCompositionCampaignInfo) Get() *CampaignAllOfRoleCompositionCampaignInfo {
	return v.value
}

func (v *NullableCampaignAllOfRoleCompositionCampaignInfo) Set(val *CampaignAllOfRoleCompositionCampaignInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOfRoleCompositionCampaignInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOfRoleCompositionCampaignInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOfRoleCompositionCampaignInfo(val *CampaignAllOfRoleCompositionCampaignInfo) *NullableCampaignAllOfRoleCompositionCampaignInfo {
	return &NullableCampaignAllOfRoleCompositionCampaignInfo{value: val, isSet: true}
}

func (v NullableCampaignAllOfRoleCompositionCampaignInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOfRoleCompositionCampaignInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


