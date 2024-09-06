/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the UserApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserApp{}

// UserApp struct for UserApp
type UserApp struct {
	// The user app id
	Id *string `json:"id,omitempty"`
	// Time when the user app was created
	Created *time.Time `json:"created,omitempty"`
	// Time when the user app was last modified
	Modified *time.Time `json:"modified,omitempty"`
	// True if the owner has multiple accounts for the source
	HasMultipleAccounts *bool `json:"hasMultipleAccounts,omitempty"`
	// True if the source has password feature
	UseForPasswordManagement *bool `json:"useForPasswordManagement,omitempty"`
	// True if the source app related to the user app is provision request enabled
	ProvisionRequestEnabled *bool `json:"provisionRequestEnabled,omitempty"`
	// True if the source app related to the user app is shown in the app center
	AppCenterEnabled *bool `json:"appCenterEnabled,omitempty"`
	SourceApp *UserAppSourceApp `json:"sourceApp,omitempty"`
	Source *UserAppSource `json:"source,omitempty"`
	Account *UserAppAccount `json:"account,omitempty"`
	Owner *UserAppOwner `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserApp UserApp

// NewUserApp instantiates a new UserApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserApp() *UserApp {
	this := UserApp{}
	var hasMultipleAccounts bool = false
	this.HasMultipleAccounts = &hasMultipleAccounts
	var useForPasswordManagement bool = false
	this.UseForPasswordManagement = &useForPasswordManagement
	var provisionRequestEnabled bool = false
	this.ProvisionRequestEnabled = &provisionRequestEnabled
	var appCenterEnabled bool = true
	this.AppCenterEnabled = &appCenterEnabled
	return &this
}

// NewUserAppWithDefaults instantiates a new UserApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAppWithDefaults() *UserApp {
	this := UserApp{}
	var hasMultipleAccounts bool = false
	this.HasMultipleAccounts = &hasMultipleAccounts
	var useForPasswordManagement bool = false
	this.UseForPasswordManagement = &useForPasswordManagement
	var provisionRequestEnabled bool = false
	this.ProvisionRequestEnabled = &provisionRequestEnabled
	var appCenterEnabled bool = true
	this.AppCenterEnabled = &appCenterEnabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserApp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserApp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserApp) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserApp) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserApp) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserApp) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *UserApp) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *UserApp) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *UserApp) SetModified(v time.Time) {
	o.Modified = &v
}

// GetHasMultipleAccounts returns the HasMultipleAccounts field value if set, zero value otherwise.
func (o *UserApp) GetHasMultipleAccounts() bool {
	if o == nil || IsNil(o.HasMultipleAccounts) {
		var ret bool
		return ret
	}
	return *o.HasMultipleAccounts
}

// GetHasMultipleAccountsOk returns a tuple with the HasMultipleAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetHasMultipleAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMultipleAccounts) {
		return nil, false
	}
	return o.HasMultipleAccounts, true
}

// HasHasMultipleAccounts returns a boolean if a field has been set.
func (o *UserApp) HasHasMultipleAccounts() bool {
	if o != nil && !IsNil(o.HasMultipleAccounts) {
		return true
	}

	return false
}

// SetHasMultipleAccounts gets a reference to the given bool and assigns it to the HasMultipleAccounts field.
func (o *UserApp) SetHasMultipleAccounts(v bool) {
	o.HasMultipleAccounts = &v
}

// GetUseForPasswordManagement returns the UseForPasswordManagement field value if set, zero value otherwise.
func (o *UserApp) GetUseForPasswordManagement() bool {
	if o == nil || IsNil(o.UseForPasswordManagement) {
		var ret bool
		return ret
	}
	return *o.UseForPasswordManagement
}

// GetUseForPasswordManagementOk returns a tuple with the UseForPasswordManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetUseForPasswordManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.UseForPasswordManagement) {
		return nil, false
	}
	return o.UseForPasswordManagement, true
}

// HasUseForPasswordManagement returns a boolean if a field has been set.
func (o *UserApp) HasUseForPasswordManagement() bool {
	if o != nil && !IsNil(o.UseForPasswordManagement) {
		return true
	}

	return false
}

// SetUseForPasswordManagement gets a reference to the given bool and assigns it to the UseForPasswordManagement field.
func (o *UserApp) SetUseForPasswordManagement(v bool) {
	o.UseForPasswordManagement = &v
}

// GetProvisionRequestEnabled returns the ProvisionRequestEnabled field value if set, zero value otherwise.
func (o *UserApp) GetProvisionRequestEnabled() bool {
	if o == nil || IsNil(o.ProvisionRequestEnabled) {
		var ret bool
		return ret
	}
	return *o.ProvisionRequestEnabled
}

// GetProvisionRequestEnabledOk returns a tuple with the ProvisionRequestEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetProvisionRequestEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProvisionRequestEnabled) {
		return nil, false
	}
	return o.ProvisionRequestEnabled, true
}

// HasProvisionRequestEnabled returns a boolean if a field has been set.
func (o *UserApp) HasProvisionRequestEnabled() bool {
	if o != nil && !IsNil(o.ProvisionRequestEnabled) {
		return true
	}

	return false
}

// SetProvisionRequestEnabled gets a reference to the given bool and assigns it to the ProvisionRequestEnabled field.
func (o *UserApp) SetProvisionRequestEnabled(v bool) {
	o.ProvisionRequestEnabled = &v
}

// GetAppCenterEnabled returns the AppCenterEnabled field value if set, zero value otherwise.
func (o *UserApp) GetAppCenterEnabled() bool {
	if o == nil || IsNil(o.AppCenterEnabled) {
		var ret bool
		return ret
	}
	return *o.AppCenterEnabled
}

// GetAppCenterEnabledOk returns a tuple with the AppCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetAppCenterEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AppCenterEnabled) {
		return nil, false
	}
	return o.AppCenterEnabled, true
}

// HasAppCenterEnabled returns a boolean if a field has been set.
func (o *UserApp) HasAppCenterEnabled() bool {
	if o != nil && !IsNil(o.AppCenterEnabled) {
		return true
	}

	return false
}

// SetAppCenterEnabled gets a reference to the given bool and assigns it to the AppCenterEnabled field.
func (o *UserApp) SetAppCenterEnabled(v bool) {
	o.AppCenterEnabled = &v
}

// GetSourceApp returns the SourceApp field value if set, zero value otherwise.
func (o *UserApp) GetSourceApp() UserAppSourceApp {
	if o == nil || IsNil(o.SourceApp) {
		var ret UserAppSourceApp
		return ret
	}
	return *o.SourceApp
}

// GetSourceAppOk returns a tuple with the SourceApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetSourceAppOk() (*UserAppSourceApp, bool) {
	if o == nil || IsNil(o.SourceApp) {
		return nil, false
	}
	return o.SourceApp, true
}

// HasSourceApp returns a boolean if a field has been set.
func (o *UserApp) HasSourceApp() bool {
	if o != nil && !IsNil(o.SourceApp) {
		return true
	}

	return false
}

// SetSourceApp gets a reference to the given UserAppSourceApp and assigns it to the SourceApp field.
func (o *UserApp) SetSourceApp(v UserAppSourceApp) {
	o.SourceApp = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *UserApp) GetSource() UserAppSource {
	if o == nil || IsNil(o.Source) {
		var ret UserAppSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetSourceOk() (*UserAppSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *UserApp) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given UserAppSource and assigns it to the Source field.
func (o *UserApp) SetSource(v UserAppSource) {
	o.Source = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UserApp) GetAccount() UserAppAccount {
	if o == nil || IsNil(o.Account) {
		var ret UserAppAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetAccountOk() (*UserAppAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *UserApp) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given UserAppAccount and assigns it to the Account field.
func (o *UserApp) SetAccount(v UserAppAccount) {
	o.Account = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *UserApp) GetOwner() UserAppOwner {
	if o == nil || IsNil(o.Owner) {
		var ret UserAppOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApp) GetOwnerOk() (*UserAppOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *UserApp) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given UserAppOwner and assigns it to the Owner field.
func (o *UserApp) SetOwner(v UserAppOwner) {
	o.Owner = &v
}

func (o UserApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.HasMultipleAccounts) {
		toSerialize["hasMultipleAccounts"] = o.HasMultipleAccounts
	}
	if !IsNil(o.UseForPasswordManagement) {
		toSerialize["useForPasswordManagement"] = o.UseForPasswordManagement
	}
	if !IsNil(o.ProvisionRequestEnabled) {
		toSerialize["provisionRequestEnabled"] = o.ProvisionRequestEnabled
	}
	if !IsNil(o.AppCenterEnabled) {
		toSerialize["appCenterEnabled"] = o.AppCenterEnabled
	}
	if !IsNil(o.SourceApp) {
		toSerialize["sourceApp"] = o.SourceApp
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserApp) UnmarshalJSON(data []byte) (err error) {
	varUserApp := _UserApp{}

	err = json.Unmarshal(data, &varUserApp)

	if err != nil {
		return err
	}

	*o = UserApp(varUserApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "hasMultipleAccounts")
		delete(additionalProperties, "useForPasswordManagement")
		delete(additionalProperties, "provisionRequestEnabled")
		delete(additionalProperties, "appCenterEnabled")
		delete(additionalProperties, "sourceApp")
		delete(additionalProperties, "source")
		delete(additionalProperties, "account")
		delete(additionalProperties, "owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserApp struct {
	value *UserApp
	isSet bool
}

func (v NullableUserApp) Get() *UserApp {
	return v.value
}

func (v *NullableUserApp) Set(val *UserApp) {
	v.value = val
	v.isSet = true
}

func (v NullableUserApp) IsSet() bool {
	return v.isSet
}

func (v *NullableUserApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserApp(val *UserApp) *NullableUserApp {
	return &NullableUserApp{value: val, isSet: true}
}

func (v NullableUserApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


