/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the CampaignAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAllOf{}

// CampaignAllOf struct for CampaignAllOf
type CampaignAllOf struct {
	// Created time of the campaign
	Created *time.Time `json:"created,omitempty"`
	// Modified time of the campaign
	Modified *time.Time `json:"modified,omitempty"`
	// The correlatedStatus of the campaign. Only SOURCE_OWNER campaigns can be Uncorrelated. An Uncorrelated certification campaign only includes Uncorrelated identities (An identity is uncorrelated if it has no accounts on an authoritative source).
	CorrelatedStatus map[string]interface{} `json:"correlatedStatus,omitempty"`
	Filter *CampaignAllOfFilter `json:"filter,omitempty"`
	// Determines if comments on sunset date changes are required.
	SunsetCommentsRequired *bool `json:"sunsetCommentsRequired,omitempty"`
	SourceOwnerCampaignInfo *CampaignAllOfSourceOwnerCampaignInfo `json:"sourceOwnerCampaignInfo,omitempty"`
	SearchCampaignInfo *CampaignAllOfSearchCampaignInfo `json:"searchCampaignInfo,omitempty"`
	RoleCompositionCampaignInfo *CampaignAllOfRoleCompositionCampaignInfo `json:"roleCompositionCampaignInfo,omitempty"`
	// A list of errors and warnings that have accumulated.
	Alerts []CampaignAlert `json:"alerts,omitempty"`
	// The total number of certifications in this campaign.
	TotalCertifications *int32 `json:"totalCertifications,omitempty"`
	// The number of completed certifications in this campaign.
	CompletedCertifications *int32 `json:"completedCertifications,omitempty"`
	// A list of sources in the campaign that contain \\\"orphan entitlements\\\" (entitlements without a corresponding Managed Attribute). An empty list indicates the campaign has no orphan entitlements. Null indicates there may be unknown orphan entitlements in the campaign (the campaign was created before this feature was implemented).
	SourcesWithOrphanEntitlements []CampaignAllOfSourcesWithOrphanEntitlements `json:"sourcesWithOrphanEntitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAllOf CampaignAllOf

// NewCampaignAllOf instantiates a new CampaignAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOf() *CampaignAllOf {
	this := CampaignAllOf{}
	var sunsetCommentsRequired bool = true
	this.SunsetCommentsRequired = &sunsetCommentsRequired
	return &this
}

// NewCampaignAllOfWithDefaults instantiates a new CampaignAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfWithDefaults() *CampaignAllOf {
	this := CampaignAllOf{}
	var sunsetCommentsRequired bool = true
	this.SunsetCommentsRequired = &sunsetCommentsRequired
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CampaignAllOf) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CampaignAllOf) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CampaignAllOf) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *CampaignAllOf) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *CampaignAllOf) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *CampaignAllOf) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCorrelatedStatus returns the CorrelatedStatus field value if set, zero value otherwise.
func (o *CampaignAllOf) GetCorrelatedStatus() map[string]interface{} {
	if o == nil || isNil(o.CorrelatedStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.CorrelatedStatus
}

// GetCorrelatedStatusOk returns a tuple with the CorrelatedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetCorrelatedStatusOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.CorrelatedStatus) {
		return map[string]interface{}{}, false
	}
	return o.CorrelatedStatus, true
}

// HasCorrelatedStatus returns a boolean if a field has been set.
func (o *CampaignAllOf) HasCorrelatedStatus() bool {
	if o != nil && !isNil(o.CorrelatedStatus) {
		return true
	}

	return false
}

// SetCorrelatedStatus gets a reference to the given map[string]interface{} and assigns it to the CorrelatedStatus field.
func (o *CampaignAllOf) SetCorrelatedStatus(v map[string]interface{}) {
	o.CorrelatedStatus = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CampaignAllOf) GetFilter() CampaignAllOfFilter {
	if o == nil || isNil(o.Filter) {
		var ret CampaignAllOfFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetFilterOk() (*CampaignAllOfFilter, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CampaignAllOf) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given CampaignAllOfFilter and assigns it to the Filter field.
func (o *CampaignAllOf) SetFilter(v CampaignAllOfFilter) {
	o.Filter = &v
}

// GetSunsetCommentsRequired returns the SunsetCommentsRequired field value if set, zero value otherwise.
func (o *CampaignAllOf) GetSunsetCommentsRequired() bool {
	if o == nil || isNil(o.SunsetCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.SunsetCommentsRequired
}

// GetSunsetCommentsRequiredOk returns a tuple with the SunsetCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetSunsetCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.SunsetCommentsRequired) {
		return nil, false
	}
	return o.SunsetCommentsRequired, true
}

// HasSunsetCommentsRequired returns a boolean if a field has been set.
func (o *CampaignAllOf) HasSunsetCommentsRequired() bool {
	if o != nil && !isNil(o.SunsetCommentsRequired) {
		return true
	}

	return false
}

// SetSunsetCommentsRequired gets a reference to the given bool and assigns it to the SunsetCommentsRequired field.
func (o *CampaignAllOf) SetSunsetCommentsRequired(v bool) {
	o.SunsetCommentsRequired = &v
}

// GetSourceOwnerCampaignInfo returns the SourceOwnerCampaignInfo field value if set, zero value otherwise.
func (o *CampaignAllOf) GetSourceOwnerCampaignInfo() CampaignAllOfSourceOwnerCampaignInfo {
	if o == nil || isNil(o.SourceOwnerCampaignInfo) {
		var ret CampaignAllOfSourceOwnerCampaignInfo
		return ret
	}
	return *o.SourceOwnerCampaignInfo
}

// GetSourceOwnerCampaignInfoOk returns a tuple with the SourceOwnerCampaignInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetSourceOwnerCampaignInfoOk() (*CampaignAllOfSourceOwnerCampaignInfo, bool) {
	if o == nil || isNil(o.SourceOwnerCampaignInfo) {
		return nil, false
	}
	return o.SourceOwnerCampaignInfo, true
}

// HasSourceOwnerCampaignInfo returns a boolean if a field has been set.
func (o *CampaignAllOf) HasSourceOwnerCampaignInfo() bool {
	if o != nil && !isNil(o.SourceOwnerCampaignInfo) {
		return true
	}

	return false
}

// SetSourceOwnerCampaignInfo gets a reference to the given CampaignAllOfSourceOwnerCampaignInfo and assigns it to the SourceOwnerCampaignInfo field.
func (o *CampaignAllOf) SetSourceOwnerCampaignInfo(v CampaignAllOfSourceOwnerCampaignInfo) {
	o.SourceOwnerCampaignInfo = &v
}

// GetSearchCampaignInfo returns the SearchCampaignInfo field value if set, zero value otherwise.
func (o *CampaignAllOf) GetSearchCampaignInfo() CampaignAllOfSearchCampaignInfo {
	if o == nil || isNil(o.SearchCampaignInfo) {
		var ret CampaignAllOfSearchCampaignInfo
		return ret
	}
	return *o.SearchCampaignInfo
}

// GetSearchCampaignInfoOk returns a tuple with the SearchCampaignInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetSearchCampaignInfoOk() (*CampaignAllOfSearchCampaignInfo, bool) {
	if o == nil || isNil(o.SearchCampaignInfo) {
		return nil, false
	}
	return o.SearchCampaignInfo, true
}

// HasSearchCampaignInfo returns a boolean if a field has been set.
func (o *CampaignAllOf) HasSearchCampaignInfo() bool {
	if o != nil && !isNil(o.SearchCampaignInfo) {
		return true
	}

	return false
}

// SetSearchCampaignInfo gets a reference to the given CampaignAllOfSearchCampaignInfo and assigns it to the SearchCampaignInfo field.
func (o *CampaignAllOf) SetSearchCampaignInfo(v CampaignAllOfSearchCampaignInfo) {
	o.SearchCampaignInfo = &v
}

// GetRoleCompositionCampaignInfo returns the RoleCompositionCampaignInfo field value if set, zero value otherwise.
func (o *CampaignAllOf) GetRoleCompositionCampaignInfo() CampaignAllOfRoleCompositionCampaignInfo {
	if o == nil || isNil(o.RoleCompositionCampaignInfo) {
		var ret CampaignAllOfRoleCompositionCampaignInfo
		return ret
	}
	return *o.RoleCompositionCampaignInfo
}

// GetRoleCompositionCampaignInfoOk returns a tuple with the RoleCompositionCampaignInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetRoleCompositionCampaignInfoOk() (*CampaignAllOfRoleCompositionCampaignInfo, bool) {
	if o == nil || isNil(o.RoleCompositionCampaignInfo) {
		return nil, false
	}
	return o.RoleCompositionCampaignInfo, true
}

// HasRoleCompositionCampaignInfo returns a boolean if a field has been set.
func (o *CampaignAllOf) HasRoleCompositionCampaignInfo() bool {
	if o != nil && !isNil(o.RoleCompositionCampaignInfo) {
		return true
	}

	return false
}

// SetRoleCompositionCampaignInfo gets a reference to the given CampaignAllOfRoleCompositionCampaignInfo and assigns it to the RoleCompositionCampaignInfo field.
func (o *CampaignAllOf) SetRoleCompositionCampaignInfo(v CampaignAllOfRoleCompositionCampaignInfo) {
	o.RoleCompositionCampaignInfo = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *CampaignAllOf) GetAlerts() []CampaignAlert {
	if o == nil || isNil(o.Alerts) {
		var ret []CampaignAlert
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetAlertsOk() ([]CampaignAlert, bool) {
	if o == nil || isNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *CampaignAllOf) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []CampaignAlert and assigns it to the Alerts field.
func (o *CampaignAllOf) SetAlerts(v []CampaignAlert) {
	o.Alerts = v
}

// GetTotalCertifications returns the TotalCertifications field value if set, zero value otherwise.
func (o *CampaignAllOf) GetTotalCertifications() int32 {
	if o == nil || isNil(o.TotalCertifications) {
		var ret int32
		return ret
	}
	return *o.TotalCertifications
}

// GetTotalCertificationsOk returns a tuple with the TotalCertifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetTotalCertificationsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCertifications) {
		return nil, false
	}
	return o.TotalCertifications, true
}

// HasTotalCertifications returns a boolean if a field has been set.
func (o *CampaignAllOf) HasTotalCertifications() bool {
	if o != nil && !isNil(o.TotalCertifications) {
		return true
	}

	return false
}

// SetTotalCertifications gets a reference to the given int32 and assigns it to the TotalCertifications field.
func (o *CampaignAllOf) SetTotalCertifications(v int32) {
	o.TotalCertifications = &v
}

// GetCompletedCertifications returns the CompletedCertifications field value if set, zero value otherwise.
func (o *CampaignAllOf) GetCompletedCertifications() int32 {
	if o == nil || isNil(o.CompletedCertifications) {
		var ret int32
		return ret
	}
	return *o.CompletedCertifications
}

// GetCompletedCertificationsOk returns a tuple with the CompletedCertifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetCompletedCertificationsOk() (*int32, bool) {
	if o == nil || isNil(o.CompletedCertifications) {
		return nil, false
	}
	return o.CompletedCertifications, true
}

// HasCompletedCertifications returns a boolean if a field has been set.
func (o *CampaignAllOf) HasCompletedCertifications() bool {
	if o != nil && !isNil(o.CompletedCertifications) {
		return true
	}

	return false
}

// SetCompletedCertifications gets a reference to the given int32 and assigns it to the CompletedCertifications field.
func (o *CampaignAllOf) SetCompletedCertifications(v int32) {
	o.CompletedCertifications = &v
}

// GetSourcesWithOrphanEntitlements returns the SourcesWithOrphanEntitlements field value if set, zero value otherwise.
func (o *CampaignAllOf) GetSourcesWithOrphanEntitlements() []CampaignAllOfSourcesWithOrphanEntitlements {
	if o == nil || isNil(o.SourcesWithOrphanEntitlements) {
		var ret []CampaignAllOfSourcesWithOrphanEntitlements
		return ret
	}
	return o.SourcesWithOrphanEntitlements
}

// GetSourcesWithOrphanEntitlementsOk returns a tuple with the SourcesWithOrphanEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetSourcesWithOrphanEntitlementsOk() ([]CampaignAllOfSourcesWithOrphanEntitlements, bool) {
	if o == nil || isNil(o.SourcesWithOrphanEntitlements) {
		return nil, false
	}
	return o.SourcesWithOrphanEntitlements, true
}

// HasSourcesWithOrphanEntitlements returns a boolean if a field has been set.
func (o *CampaignAllOf) HasSourcesWithOrphanEntitlements() bool {
	if o != nil && !isNil(o.SourcesWithOrphanEntitlements) {
		return true
	}

	return false
}

// SetSourcesWithOrphanEntitlements gets a reference to the given []CampaignAllOfSourcesWithOrphanEntitlements and assigns it to the SourcesWithOrphanEntitlements field.
func (o *CampaignAllOf) SetSourcesWithOrphanEntitlements(v []CampaignAllOfSourcesWithOrphanEntitlements) {
	o.SourcesWithOrphanEntitlements = v
}

func (o CampaignAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created is readOnly
	// skip: modified is readOnly
	if !isNil(o.CorrelatedStatus) {
		toSerialize["correlatedStatus"] = o.CorrelatedStatus
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.SunsetCommentsRequired) {
		toSerialize["sunsetCommentsRequired"] = o.SunsetCommentsRequired
	}
	if !isNil(o.SourceOwnerCampaignInfo) {
		toSerialize["sourceOwnerCampaignInfo"] = o.SourceOwnerCampaignInfo
	}
	if !isNil(o.SearchCampaignInfo) {
		toSerialize["searchCampaignInfo"] = o.SearchCampaignInfo
	}
	if !isNil(o.RoleCompositionCampaignInfo) {
		toSerialize["roleCompositionCampaignInfo"] = o.RoleCompositionCampaignInfo
	}
	// skip: alerts is readOnly
	// skip: totalCertifications is readOnly
	// skip: completedCertifications is readOnly
	// skip: sourcesWithOrphanEntitlements is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignAllOf := _CampaignAllOf{}

	if err = json.Unmarshal(bytes, &varCampaignAllOf); err == nil {
		*o = CampaignAllOf(varCampaignAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "correlatedStatus")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "sunsetCommentsRequired")
		delete(additionalProperties, "sourceOwnerCampaignInfo")
		delete(additionalProperties, "searchCampaignInfo")
		delete(additionalProperties, "roleCompositionCampaignInfo")
		delete(additionalProperties, "alerts")
		delete(additionalProperties, "totalCertifications")
		delete(additionalProperties, "completedCertifications")
		delete(additionalProperties, "sourcesWithOrphanEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAllOf struct {
	value *CampaignAllOf
	isSet bool
}

func (v NullableCampaignAllOf) Get() *CampaignAllOf {
	return v.value
}

func (v *NullableCampaignAllOf) Set(val *CampaignAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOf(val *CampaignAllOf) *NullableCampaignAllOf {
	return &NullableCampaignAllOf{value: val, isSet: true}
}

func (v NullableCampaignAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


