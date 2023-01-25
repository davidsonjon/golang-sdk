/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"time"
	"encoding/json"
)

// LifecycleState struct for LifecycleState
type LifecycleState struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// Whether the lifecycle state is enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The technical name for lifecycle state. This is for internal use.
	TechnicalName string `json:"technicalName"`
	// Lifecycle state description.
	Description *string `json:"description,omitempty"`
	// Number of identities that have the lifecycle state.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	EmailNotificationOption *EmailNotificationOption `json:"emailNotificationOption,omitempty"`
	AccountActions []AccountAction `json:"accountActions,omitempty"`
	// List of unique access-profile IDs that are associated with the lifecycle state.
	AccessProfileIds []string `json:"accessProfileIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleState LifecycleState

// NewLifecycleState instantiates a new LifecycleState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleState(name string, technicalName string) *LifecycleState {
	this := LifecycleState{}
	this.Name = name
	this.TechnicalName = technicalName
	return &this
}

// NewLifecycleStateWithDefaults instantiates a new LifecycleState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleStateWithDefaults() *LifecycleState {
	this := LifecycleState{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LifecycleState) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LifecycleState) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LifecycleState) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *LifecycleState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LifecycleState) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LifecycleState) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LifecycleState) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *LifecycleState) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *LifecycleState) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *LifecycleState) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *LifecycleState) SetModified(v time.Time) {
	o.Modified = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LifecycleState) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LifecycleState) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LifecycleState) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTechnicalName returns the TechnicalName field value
func (o *LifecycleState) GetTechnicalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetTechnicalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TechnicalName, true
}

// SetTechnicalName sets field value
func (o *LifecycleState) SetTechnicalName(v string) {
	o.TechnicalName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LifecycleState) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LifecycleState) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LifecycleState) SetDescription(v string) {
	o.Description = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *LifecycleState) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *LifecycleState) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *LifecycleState) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetEmailNotificationOption returns the EmailNotificationOption field value if set, zero value otherwise.
func (o *LifecycleState) GetEmailNotificationOption() EmailNotificationOption {
	if o == nil || isNil(o.EmailNotificationOption) {
		var ret EmailNotificationOption
		return ret
	}
	return *o.EmailNotificationOption
}

// GetEmailNotificationOptionOk returns a tuple with the EmailNotificationOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetEmailNotificationOptionOk() (*EmailNotificationOption, bool) {
	if o == nil || isNil(o.EmailNotificationOption) {
		return nil, false
	}
	return o.EmailNotificationOption, true
}

// HasEmailNotificationOption returns a boolean if a field has been set.
func (o *LifecycleState) HasEmailNotificationOption() bool {
	if o != nil && !isNil(o.EmailNotificationOption) {
		return true
	}

	return false
}

// SetEmailNotificationOption gets a reference to the given EmailNotificationOption and assigns it to the EmailNotificationOption field.
func (o *LifecycleState) SetEmailNotificationOption(v EmailNotificationOption) {
	o.EmailNotificationOption = &v
}

// GetAccountActions returns the AccountActions field value if set, zero value otherwise.
func (o *LifecycleState) GetAccountActions() []AccountAction {
	if o == nil || isNil(o.AccountActions) {
		var ret []AccountAction
		return ret
	}
	return o.AccountActions
}

// GetAccountActionsOk returns a tuple with the AccountActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetAccountActionsOk() ([]AccountAction, bool) {
	if o == nil || isNil(o.AccountActions) {
		return nil, false
	}
	return o.AccountActions, true
}

// HasAccountActions returns a boolean if a field has been set.
func (o *LifecycleState) HasAccountActions() bool {
	if o != nil && !isNil(o.AccountActions) {
		return true
	}

	return false
}

// SetAccountActions gets a reference to the given []AccountAction and assigns it to the AccountActions field.
func (o *LifecycleState) SetAccountActions(v []AccountAction) {
	o.AccountActions = v
}

// GetAccessProfileIds returns the AccessProfileIds field value if set, zero value otherwise.
func (o *LifecycleState) GetAccessProfileIds() []string {
	if o == nil || isNil(o.AccessProfileIds) {
		var ret []string
		return ret
	}
	return o.AccessProfileIds
}

// GetAccessProfileIdsOk returns a tuple with the AccessProfileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetAccessProfileIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AccessProfileIds) {
		return nil, false
	}
	return o.AccessProfileIds, true
}

// HasAccessProfileIds returns a boolean if a field has been set.
func (o *LifecycleState) HasAccessProfileIds() bool {
	if o != nil && !isNil(o.AccessProfileIds) {
		return true
	}

	return false
}

// SetAccessProfileIds gets a reference to the given []string and assigns it to the AccessProfileIds field.
func (o *LifecycleState) SetAccessProfileIds(v []string) {
	o.AccessProfileIds = v
}

func (o LifecycleState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["technicalName"] = o.TechnicalName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.EmailNotificationOption) {
		toSerialize["emailNotificationOption"] = o.EmailNotificationOption
	}
	if !isNil(o.AccountActions) {
		toSerialize["accountActions"] = o.AccountActions
	}
	if !isNil(o.AccessProfileIds) {
		toSerialize["accessProfileIds"] = o.AccessProfileIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LifecycleState) UnmarshalJSON(bytes []byte) (err error) {
	varLifecycleState := _LifecycleState{}

	if err = json.Unmarshal(bytes, &varLifecycleState); err == nil {
		*o = LifecycleState(varLifecycleState)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "technicalName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "emailNotificationOption")
		delete(additionalProperties, "accountActions")
		delete(additionalProperties, "accessProfileIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleState struct {
	value *LifecycleState
	isSet bool
}

func (v NullableLifecycleState) Get() *LifecycleState {
	return v.value
}

func (v *NullableLifecycleState) Set(val *LifecycleState) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleState) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleState(val *LifecycleState) *NullableLifecycleState {
	return &NullableLifecycleState{value: val, isSet: true}
}

func (v NullableLifecycleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


