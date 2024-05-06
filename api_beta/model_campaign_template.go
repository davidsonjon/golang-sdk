/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CampaignTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignTemplate{}

// CampaignTemplate Campaign Template
type CampaignTemplate struct {
	// Id of the campaign template
	Id *string `json:"id,omitempty"`
	// This template's name. Has no bearing on generated campaigns' names.
	Name string `json:"name"`
	// This template's description. Has no bearing on generated campaigns' descriptions.
	Description string `json:"description"`
	// Creation date of Campaign Template
	Created time.Time `json:"created"`
	// Modification date of Campaign Template
	Modified time.Time `json:"modified"`
	// Indicates if this campaign template has been scheduled.
	Scheduled *bool `json:"scheduled,omitempty"`
	OwnerRef *CampaignTemplateOwnerRef `json:"ownerRef,omitempty"`
	// The time period during which the campaign should be completed, formatted as an ISO-8601 Duration. When this template generates a campaign, the campaign's deadline will be the current date plus this duration. For example, if generation occurred on 2020-01-01 and this field was \"P2W\" (two weeks), the resulting campaign's deadline would be 2020-01-15 (the current date plus 14 days).
	DeadlineDuration *string `json:"deadlineDuration,omitempty"`
	// This will hold campaign related information like name, description etc.
	Campaign Campaign `json:"campaign"`
	AdditionalProperties map[string]interface{}
}

type _CampaignTemplate CampaignTemplate

// NewCampaignTemplate instantiates a new CampaignTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignTemplate(name string, description string, created time.Time, modified time.Time, campaign Campaign) *CampaignTemplate {
	this := CampaignTemplate{}
	this.Name = name
	this.Description = description
	this.Created = created
	this.Modified = modified
	this.Campaign = campaign
	return &this
}

// NewCampaignTemplateWithDefaults instantiates a new CampaignTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignTemplateWithDefaults() *CampaignTemplate {
	this := CampaignTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CampaignTemplate) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CampaignTemplate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CampaignTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CampaignTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CampaignTemplate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CampaignTemplate) SetDescription(v string) {
	o.Description = v
}

// GetCreated returns the Created field value
func (o *CampaignTemplate) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignTemplate) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *CampaignTemplate) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *CampaignTemplate) SetModified(v time.Time) {
	o.Modified = v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *CampaignTemplate) GetScheduled() bool {
	if o == nil || isNil(o.Scheduled) {
		var ret bool
		return ret
	}
	return *o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetScheduledOk() (*bool, bool) {
	if o == nil || isNil(o.Scheduled) {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *CampaignTemplate) HasScheduled() bool {
	if o != nil && !isNil(o.Scheduled) {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given bool and assigns it to the Scheduled field.
func (o *CampaignTemplate) SetScheduled(v bool) {
	o.Scheduled = &v
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise.
func (o *CampaignTemplate) GetOwnerRef() CampaignTemplateOwnerRef {
	if o == nil || isNil(o.OwnerRef) {
		var ret CampaignTemplateOwnerRef
		return ret
	}
	return *o.OwnerRef
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetOwnerRefOk() (*CampaignTemplateOwnerRef, bool) {
	if o == nil || isNil(o.OwnerRef) {
		return nil, false
	}
	return o.OwnerRef, true
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *CampaignTemplate) HasOwnerRef() bool {
	if o != nil && !isNil(o.OwnerRef) {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given CampaignTemplateOwnerRef and assigns it to the OwnerRef field.
func (o *CampaignTemplate) SetOwnerRef(v CampaignTemplateOwnerRef) {
	o.OwnerRef = &v
}

// GetDeadlineDuration returns the DeadlineDuration field value if set, zero value otherwise.
func (o *CampaignTemplate) GetDeadlineDuration() string {
	if o == nil || isNil(o.DeadlineDuration) {
		var ret string
		return ret
	}
	return *o.DeadlineDuration
}

// GetDeadlineDurationOk returns a tuple with the DeadlineDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetDeadlineDurationOk() (*string, bool) {
	if o == nil || isNil(o.DeadlineDuration) {
		return nil, false
	}
	return o.DeadlineDuration, true
}

// HasDeadlineDuration returns a boolean if a field has been set.
func (o *CampaignTemplate) HasDeadlineDuration() bool {
	if o != nil && !isNil(o.DeadlineDuration) {
		return true
	}

	return false
}

// SetDeadlineDuration gets a reference to the given string and assigns it to the DeadlineDuration field.
func (o *CampaignTemplate) SetDeadlineDuration(v string) {
	o.DeadlineDuration = &v
}

// GetCampaign returns the Campaign field value
func (o *CampaignTemplate) GetCampaign() Campaign {
	if o == nil {
		var ret Campaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetCampaignOk() (Campaign, bool) {
	if o == nil {
		return Campaign{}, false
	}
	return o.Campaign, true
}

// SetCampaign sets field value
func (o *CampaignTemplate) SetCampaign(v Campaign) {
	o.Campaign = v
}

func (o CampaignTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	// skip: created is readOnly
	// skip: modified is readOnly
	// skip: scheduled is readOnly
	if !isNil(o.OwnerRef) {
		toSerialize["ownerRef"] = o.OwnerRef
	}
	if !isNil(o.DeadlineDuration) {
		toSerialize["deadlineDuration"] = o.DeadlineDuration
	}
	toSerialize["campaign"] = o.Campaign

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignTemplate) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"created",
		"modified",
		"campaign",
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

	varCampaignTemplate := _CampaignTemplate{}

	if err = json.Unmarshal(bytes, &varCampaignTemplate); err == nil {
	*o = CampaignTemplate(varCampaignTemplate)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "scheduled")
		delete(additionalProperties, "ownerRef")
		delete(additionalProperties, "deadlineDuration")
		delete(additionalProperties, "campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignTemplate struct {
	value *CampaignTemplate
	isSet bool
}

func (v NullableCampaignTemplate) Get() *CampaignTemplate {
	return v.value
}

func (v *NullableCampaignTemplate) Set(val *CampaignTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignTemplate(val *CampaignTemplate) *NullableCampaignTemplate {
	return &NullableCampaignTemplate{value: val, isSet: true}
}

func (v NullableCampaignTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


