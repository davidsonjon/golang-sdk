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

// checks if the SlimCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimCampaign{}

// SlimCampaign struct for SlimCampaign
type SlimCampaign struct {
	// Id of the campaign
	Id *string `json:"id,omitempty"`
	// The campaign name. If this object is part of a template, special formatting applies; see the `/campaign-templates/{id}/generate` endpoint documentation for details.
	Name string `json:"name"`
	// The campaign description. If this object is part of a template, special formatting applies; see the `/campaign-templates/{id}/generate` endpoint documentation for details.
	Description string `json:"description"`
	// The campaign's completion deadline.  This date must be in the future in order to activate the campaign.  If you try to activate a campaign with a deadline of today or in the past, you will receive a 400 error response.
	Deadline *time.Time `json:"deadline,omitempty"`
	// The type of campaign. Could be extended in the future.
	Type string `json:"type"`
	// Enables email notification for this campaign
	EmailNotificationEnabled *bool `json:"emailNotificationEnabled,omitempty"`
	// Allows auto revoke for this campaign
	AutoRevokeAllowed *bool `json:"autoRevokeAllowed,omitempty"`
	// Enables IAI for this campaign. Accepts true even if the IAI product feature is off. If IAI is turned off then campaigns generated from this template will indicate false. The real value will then be returned if IAI is ever enabled for the org in the future.
	RecommendationsEnabled *bool `json:"recommendationsEnabled,omitempty"`
	// The campaign's current status.
	Status *string `json:"status,omitempty"`
	// The correlatedStatus of the campaign. Only SOURCE_OWNER campaigns can be Uncorrelated. An Uncorrelated certification campaign only includes Uncorrelated identities (An identity is uncorrelated if it has no accounts on an authoritative source).
	CorrelatedStatus *string `json:"correlatedStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SlimCampaign SlimCampaign

// NewSlimCampaign instantiates a new SlimCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimCampaign(name string, description string, type_ string) *SlimCampaign {
	this := SlimCampaign{}
	this.Name = name
	this.Description = description
	this.Type = type_
	var emailNotificationEnabled bool = false
	this.EmailNotificationEnabled = &emailNotificationEnabled
	var autoRevokeAllowed bool = false
	this.AutoRevokeAllowed = &autoRevokeAllowed
	var recommendationsEnabled bool = false
	this.RecommendationsEnabled = &recommendationsEnabled
	return &this
}

// NewSlimCampaignWithDefaults instantiates a new SlimCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimCampaignWithDefaults() *SlimCampaign {
	this := SlimCampaign{}
	var emailNotificationEnabled bool = false
	this.EmailNotificationEnabled = &emailNotificationEnabled
	var autoRevokeAllowed bool = false
	this.AutoRevokeAllowed = &autoRevokeAllowed
	var recommendationsEnabled bool = false
	this.RecommendationsEnabled = &recommendationsEnabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SlimCampaign) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SlimCampaign) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SlimCampaign) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SlimCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlimCampaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *SlimCampaign) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SlimCampaign) SetDescription(v string) {
	o.Description = v
}

// GetDeadline returns the Deadline field value if set, zero value otherwise.
func (o *SlimCampaign) GetDeadline() time.Time {
	if o == nil || isNil(o.Deadline) {
		var ret time.Time
		return ret
	}
	return *o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetDeadlineOk() (*time.Time, bool) {
	if o == nil || isNil(o.Deadline) {
		return nil, false
	}
	return o.Deadline, true
}

// HasDeadline returns a boolean if a field has been set.
func (o *SlimCampaign) HasDeadline() bool {
	if o != nil && !isNil(o.Deadline) {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given time.Time and assigns it to the Deadline field.
func (o *SlimCampaign) SetDeadline(v time.Time) {
	o.Deadline = &v
}

// GetType returns the Type field value
func (o *SlimCampaign) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlimCampaign) SetType(v string) {
	o.Type = v
}

// GetEmailNotificationEnabled returns the EmailNotificationEnabled field value if set, zero value otherwise.
func (o *SlimCampaign) GetEmailNotificationEnabled() bool {
	if o == nil || isNil(o.EmailNotificationEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailNotificationEnabled
}

// GetEmailNotificationEnabledOk returns a tuple with the EmailNotificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetEmailNotificationEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EmailNotificationEnabled) {
		return nil, false
	}
	return o.EmailNotificationEnabled, true
}

// HasEmailNotificationEnabled returns a boolean if a field has been set.
func (o *SlimCampaign) HasEmailNotificationEnabled() bool {
	if o != nil && !isNil(o.EmailNotificationEnabled) {
		return true
	}

	return false
}

// SetEmailNotificationEnabled gets a reference to the given bool and assigns it to the EmailNotificationEnabled field.
func (o *SlimCampaign) SetEmailNotificationEnabled(v bool) {
	o.EmailNotificationEnabled = &v
}

// GetAutoRevokeAllowed returns the AutoRevokeAllowed field value if set, zero value otherwise.
func (o *SlimCampaign) GetAutoRevokeAllowed() bool {
	if o == nil || isNil(o.AutoRevokeAllowed) {
		var ret bool
		return ret
	}
	return *o.AutoRevokeAllowed
}

// GetAutoRevokeAllowedOk returns a tuple with the AutoRevokeAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetAutoRevokeAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.AutoRevokeAllowed) {
		return nil, false
	}
	return o.AutoRevokeAllowed, true
}

// HasAutoRevokeAllowed returns a boolean if a field has been set.
func (o *SlimCampaign) HasAutoRevokeAllowed() bool {
	if o != nil && !isNil(o.AutoRevokeAllowed) {
		return true
	}

	return false
}

// SetAutoRevokeAllowed gets a reference to the given bool and assigns it to the AutoRevokeAllowed field.
func (o *SlimCampaign) SetAutoRevokeAllowed(v bool) {
	o.AutoRevokeAllowed = &v
}

// GetRecommendationsEnabled returns the RecommendationsEnabled field value if set, zero value otherwise.
func (o *SlimCampaign) GetRecommendationsEnabled() bool {
	if o == nil || isNil(o.RecommendationsEnabled) {
		var ret bool
		return ret
	}
	return *o.RecommendationsEnabled
}

// GetRecommendationsEnabledOk returns a tuple with the RecommendationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetRecommendationsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RecommendationsEnabled) {
		return nil, false
	}
	return o.RecommendationsEnabled, true
}

// HasRecommendationsEnabled returns a boolean if a field has been set.
func (o *SlimCampaign) HasRecommendationsEnabled() bool {
	if o != nil && !isNil(o.RecommendationsEnabled) {
		return true
	}

	return false
}

// SetRecommendationsEnabled gets a reference to the given bool and assigns it to the RecommendationsEnabled field.
func (o *SlimCampaign) SetRecommendationsEnabled(v bool) {
	o.RecommendationsEnabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SlimCampaign) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SlimCampaign) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SlimCampaign) SetStatus(v string) {
	o.Status = &v
}

// GetCorrelatedStatus returns the CorrelatedStatus field value if set, zero value otherwise.
func (o *SlimCampaign) GetCorrelatedStatus() string {
	if o == nil || isNil(o.CorrelatedStatus) {
		var ret string
		return ret
	}
	return *o.CorrelatedStatus
}

// GetCorrelatedStatusOk returns a tuple with the CorrelatedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimCampaign) GetCorrelatedStatusOk() (*string, bool) {
	if o == nil || isNil(o.CorrelatedStatus) {
		return nil, false
	}
	return o.CorrelatedStatus, true
}

// HasCorrelatedStatus returns a boolean if a field has been set.
func (o *SlimCampaign) HasCorrelatedStatus() bool {
	if o != nil && !isNil(o.CorrelatedStatus) {
		return true
	}

	return false
}

// SetCorrelatedStatus gets a reference to the given string and assigns it to the CorrelatedStatus field.
func (o *SlimCampaign) SetCorrelatedStatus(v string) {
	o.CorrelatedStatus = &v
}

func (o SlimCampaign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !isNil(o.Deadline) {
		toSerialize["deadline"] = o.Deadline
	}
	toSerialize["type"] = o.Type
	if !isNil(o.EmailNotificationEnabled) {
		toSerialize["emailNotificationEnabled"] = o.EmailNotificationEnabled
	}
	if !isNil(o.AutoRevokeAllowed) {
		toSerialize["autoRevokeAllowed"] = o.AutoRevokeAllowed
	}
	if !isNil(o.RecommendationsEnabled) {
		toSerialize["recommendationsEnabled"] = o.RecommendationsEnabled
	}
	// skip: status is readOnly
	if !isNil(o.CorrelatedStatus) {
		toSerialize["correlatedStatus"] = o.CorrelatedStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SlimCampaign) UnmarshalJSON(bytes []byte) (err error) {
	varSlimCampaign := _SlimCampaign{}

	if err = json.Unmarshal(bytes, &varSlimCampaign); err == nil {
		*o = SlimCampaign(varSlimCampaign)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "deadline")
		delete(additionalProperties, "type")
		delete(additionalProperties, "emailNotificationEnabled")
		delete(additionalProperties, "autoRevokeAllowed")
		delete(additionalProperties, "recommendationsEnabled")
		delete(additionalProperties, "status")
		delete(additionalProperties, "correlatedStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSlimCampaign struct {
	value *SlimCampaign
	isSet bool
}

func (v NullableSlimCampaign) Get() *SlimCampaign {
	return v.value
}

func (v *NullableSlimCampaign) Set(val *SlimCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimCampaign(val *SlimCampaign) *NullableSlimCampaign {
	return &NullableSlimCampaign{value: val, isSet: true}
}

func (v NullableSlimCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


