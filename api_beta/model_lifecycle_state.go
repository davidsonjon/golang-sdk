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
)

// checks if the LifecycleState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleState{}

// LifecycleState struct for LifecycleState
type LifecycleState struct {
	// Lifecycle state ID.
	Id *string `json:"id,omitempty"`
	// Lifecycle state name.
	Name *string `json:"name,omitempty"`
	// Lifecycle state technical name. This is for internal use.
	TechnicalName *string `json:"technicalName,omitempty"`
	// Lifecycle state description.
	Description *string `json:"description,omitempty"`
	// Lifecycle state created date.
	Created *time.Time `json:"created,omitempty"`
	// Lifecycle state modified date.
	Modified *time.Time `json:"modified,omitempty"`
	// Indicates whether the lifecycle state is enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Number of identities that have the lifecycle state.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	EmailNotificationOption *EmailNotificationOption `json:"emailNotificationOption,omitempty"`
	AccountActions []AccountAction `json:"accountActions,omitempty"`
	// List of access-profile IDs that are associated with the lifecycle state.
	AccessProfileIds []string `json:"accessProfileIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleState LifecycleState

// NewLifecycleState instantiates a new LifecycleState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleState() *LifecycleState {
	this := LifecycleState{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewLifecycleStateWithDefaults instantiates a new LifecycleState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleStateWithDefaults() *LifecycleState {
	this := LifecycleState{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LifecycleState) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LifecycleState) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LifecycleState) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LifecycleState) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LifecycleState) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LifecycleState) SetName(v string) {
	o.Name = &v
}

// GetTechnicalName returns the TechnicalName field value if set, zero value otherwise.
func (o *LifecycleState) GetTechnicalName() string {
	if o == nil || IsNil(o.TechnicalName) {
		var ret string
		return ret
	}
	return *o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetTechnicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicalName) {
		return nil, false
	}
	return o.TechnicalName, true
}

// HasTechnicalName returns a boolean if a field has been set.
func (o *LifecycleState) HasTechnicalName() bool {
	if o != nil && !IsNil(o.TechnicalName) {
		return true
	}

	return false
}

// SetTechnicalName gets a reference to the given string and assigns it to the TechnicalName field.
func (o *LifecycleState) SetTechnicalName(v string) {
	o.TechnicalName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LifecycleState) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LifecycleState) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LifecycleState) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LifecycleState) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LifecycleState) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
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
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *LifecycleState) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
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
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LifecycleState) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LifecycleState) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *LifecycleState) GetIdentityCount() int32 {
	if o == nil || IsNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetIdentityCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *LifecycleState) HasIdentityCount() bool {
	if o != nil && !IsNil(o.IdentityCount) {
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
	if o == nil || IsNil(o.EmailNotificationOption) {
		var ret EmailNotificationOption
		return ret
	}
	return *o.EmailNotificationOption
}

// GetEmailNotificationOptionOk returns a tuple with the EmailNotificationOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetEmailNotificationOptionOk() (*EmailNotificationOption, bool) {
	if o == nil || IsNil(o.EmailNotificationOption) {
		return nil, false
	}
	return o.EmailNotificationOption, true
}

// HasEmailNotificationOption returns a boolean if a field has been set.
func (o *LifecycleState) HasEmailNotificationOption() bool {
	if o != nil && !IsNil(o.EmailNotificationOption) {
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
	if o == nil || IsNil(o.AccountActions) {
		var ret []AccountAction
		return ret
	}
	return o.AccountActions
}

// GetAccountActionsOk returns a tuple with the AccountActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetAccountActionsOk() ([]AccountAction, bool) {
	if o == nil || IsNil(o.AccountActions) {
		return nil, false
	}
	return o.AccountActions, true
}

// HasAccountActions returns a boolean if a field has been set.
func (o *LifecycleState) HasAccountActions() bool {
	if o != nil && !IsNil(o.AccountActions) {
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
	if o == nil || IsNil(o.AccessProfileIds) {
		var ret []string
		return ret
	}
	return o.AccessProfileIds
}

// GetAccessProfileIdsOk returns a tuple with the AccessProfileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetAccessProfileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessProfileIds) {
		return nil, false
	}
	return o.AccessProfileIds, true
}

// HasAccessProfileIds returns a boolean if a field has been set.
func (o *LifecycleState) HasAccessProfileIds() bool {
	if o != nil && !IsNil(o.AccessProfileIds) {
		return true
	}

	return false
}

// SetAccessProfileIds gets a reference to the given []string and assigns it to the AccessProfileIds field.
func (o *LifecycleState) SetAccessProfileIds(v []string) {
	o.AccessProfileIds = v
}

func (o LifecycleState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TechnicalName) {
		toSerialize["technicalName"] = o.TechnicalName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !IsNil(o.EmailNotificationOption) {
		toSerialize["emailNotificationOption"] = o.EmailNotificationOption
	}
	if !IsNil(o.AccountActions) {
		toSerialize["accountActions"] = o.AccountActions
	}
	if !IsNil(o.AccessProfileIds) {
		toSerialize["accessProfileIds"] = o.AccessProfileIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleState) UnmarshalJSON(data []byte) (err error) {
	varLifecycleState := _LifecycleState{}

	err = json.Unmarshal(data, &varLifecycleState)

	if err != nil {
		return err
	}

	*o = LifecycleState(varLifecycleState)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "technicalName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "enabled")
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


