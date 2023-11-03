/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the ListAccounts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccounts200ResponseInner{}

// ListAccounts200ResponseInner struct for ListAccounts200ResponseInner
type ListAccounts200ResponseInner struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Username *string `json:"username,omitempty"`
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	PasswordProvided *bool `json:"passwordProvided,omitempty"`
	Apps []map[string]interface{} `json:"apps,omitempty"`
	SsoMethod *string `json:"ssoMethod,omitempty"`
	IdEncryption *string `json:"idEncryption,omitempty"`
	PasswordEncryption *string `json:"passwordEncryption,omitempty"`
	LastPasswdChange NullableString `json:"lastPasswdChange,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
	DateDisabled NullableString `json:"dateDisabled,omitempty"`
	AccountServiceId *int32 `json:"accountServiceId,omitempty"`
	ServiceId *int32 `json:"serviceId,omitempty"`
	PendingPasswordRequestId NullableString `json:"pendingPasswordRequestId,omitempty"`
	PasswordChangeStatus *string `json:"passwordChangeStatus,omitempty"`
	PasswordChangeResult *ListAccounts200ResponseInnerPasswordChangeResult `json:"passwordChangeResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAccounts200ResponseInner ListAccounts200ResponseInner

// NewListAccounts200ResponseInner instantiates a new ListAccounts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccounts200ResponseInner() *ListAccounts200ResponseInner {
	this := ListAccounts200ResponseInner{}
	return &this
}

// NewListAccounts200ResponseInnerWithDefaults instantiates a new ListAccounts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccounts200ResponseInnerWithDefaults() *ListAccounts200ResponseInner {
	this := ListAccounts200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListAccounts200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListAccounts200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListAccounts200ResponseInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ListAccounts200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

// GetPasswordRequired returns the PasswordRequired field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetPasswordRequired() bool {
	if o == nil || isNil(o.PasswordRequired) {
		var ret bool
		return ret
	}
	return *o.PasswordRequired
}

// GetPasswordRequiredOk returns a tuple with the PasswordRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetPasswordRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.PasswordRequired) {
		return nil, false
	}
	return o.PasswordRequired, true
}

// HasPasswordRequired returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasPasswordRequired() bool {
	if o != nil && !isNil(o.PasswordRequired) {
		return true
	}

	return false
}

// SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.
func (o *ListAccounts200ResponseInner) SetPasswordRequired(v bool) {
	o.PasswordRequired = &v
}

// GetPasswordProvided returns the PasswordProvided field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetPasswordProvided() bool {
	if o == nil || isNil(o.PasswordProvided) {
		var ret bool
		return ret
	}
	return *o.PasswordProvided
}

// GetPasswordProvidedOk returns a tuple with the PasswordProvided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetPasswordProvidedOk() (*bool, bool) {
	if o == nil || isNil(o.PasswordProvided) {
		return nil, false
	}
	return o.PasswordProvided, true
}

// HasPasswordProvided returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasPasswordProvided() bool {
	if o != nil && !isNil(o.PasswordProvided) {
		return true
	}

	return false
}

// SetPasswordProvided gets a reference to the given bool and assigns it to the PasswordProvided field.
func (o *ListAccounts200ResponseInner) SetPasswordProvided(v bool) {
	o.PasswordProvided = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetApps() []map[string]interface{} {
	if o == nil || isNil(o.Apps) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetAppsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasApps() bool {
	if o != nil && !isNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []map[string]interface{} and assigns it to the Apps field.
func (o *ListAccounts200ResponseInner) SetApps(v []map[string]interface{}) {
	o.Apps = v
}

// GetSsoMethod returns the SsoMethod field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetSsoMethod() string {
	if o == nil || isNil(o.SsoMethod) {
		var ret string
		return ret
	}
	return *o.SsoMethod
}

// GetSsoMethodOk returns a tuple with the SsoMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetSsoMethodOk() (*string, bool) {
	if o == nil || isNil(o.SsoMethod) {
		return nil, false
	}
	return o.SsoMethod, true
}

// HasSsoMethod returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasSsoMethod() bool {
	if o != nil && !isNil(o.SsoMethod) {
		return true
	}

	return false
}

// SetSsoMethod gets a reference to the given string and assigns it to the SsoMethod field.
func (o *ListAccounts200ResponseInner) SetSsoMethod(v string) {
	o.SsoMethod = &v
}

// GetIdEncryption returns the IdEncryption field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetIdEncryption() string {
	if o == nil || isNil(o.IdEncryption) {
		var ret string
		return ret
	}
	return *o.IdEncryption
}

// GetIdEncryptionOk returns a tuple with the IdEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetIdEncryptionOk() (*string, bool) {
	if o == nil || isNil(o.IdEncryption) {
		return nil, false
	}
	return o.IdEncryption, true
}

// HasIdEncryption returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasIdEncryption() bool {
	if o != nil && !isNil(o.IdEncryption) {
		return true
	}

	return false
}

// SetIdEncryption gets a reference to the given string and assigns it to the IdEncryption field.
func (o *ListAccounts200ResponseInner) SetIdEncryption(v string) {
	o.IdEncryption = &v
}

// GetPasswordEncryption returns the PasswordEncryption field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetPasswordEncryption() string {
	if o == nil || isNil(o.PasswordEncryption) {
		var ret string
		return ret
	}
	return *o.PasswordEncryption
}

// GetPasswordEncryptionOk returns a tuple with the PasswordEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetPasswordEncryptionOk() (*string, bool) {
	if o == nil || isNil(o.PasswordEncryption) {
		return nil, false
	}
	return o.PasswordEncryption, true
}

// HasPasswordEncryption returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasPasswordEncryption() bool {
	if o != nil && !isNil(o.PasswordEncryption) {
		return true
	}

	return false
}

// SetPasswordEncryption gets a reference to the given string and assigns it to the PasswordEncryption field.
func (o *ListAccounts200ResponseInner) SetPasswordEncryption(v string) {
	o.PasswordEncryption = &v
}

// GetLastPasswdChange returns the LastPasswdChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAccounts200ResponseInner) GetLastPasswdChange() string {
	if o == nil || isNil(o.LastPasswdChange.Get()) {
		var ret string
		return ret
	}
	return *o.LastPasswdChange.Get()
}

// GetLastPasswdChangeOk returns a tuple with the LastPasswdChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAccounts200ResponseInner) GetLastPasswdChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPasswdChange.Get(), o.LastPasswdChange.IsSet()
}

// HasLastPasswdChange returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasLastPasswdChange() bool {
	if o != nil && o.LastPasswdChange.IsSet() {
		return true
	}

	return false
}

// SetLastPasswdChange gets a reference to the given NullableString and assigns it to the LastPasswdChange field.
func (o *ListAccounts200ResponseInner) SetLastPasswdChange(v string) {
	o.LastPasswdChange.Set(&v)
}
// SetLastPasswdChangeNil sets the value for LastPasswdChange to be an explicit nil
func (o *ListAccounts200ResponseInner) SetLastPasswdChangeNil() {
	o.LastPasswdChange.Set(nil)
}

// UnsetLastPasswdChange ensures that no value is present for LastPasswdChange, not even an explicit nil
func (o *ListAccounts200ResponseInner) UnsetLastPasswdChange() {
	o.LastPasswdChange.Unset()
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetServiceName() string {
	if o == nil || isNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetServiceNameOk() (*string, bool) {
	if o == nil || isNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasServiceName() bool {
	if o != nil && !isNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ListAccounts200ResponseInner) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetDateDisabled returns the DateDisabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAccounts200ResponseInner) GetDateDisabled() string {
	if o == nil || isNil(o.DateDisabled.Get()) {
		var ret string
		return ret
	}
	return *o.DateDisabled.Get()
}

// GetDateDisabledOk returns a tuple with the DateDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAccounts200ResponseInner) GetDateDisabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateDisabled.Get(), o.DateDisabled.IsSet()
}

// HasDateDisabled returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasDateDisabled() bool {
	if o != nil && o.DateDisabled.IsSet() {
		return true
	}

	return false
}

// SetDateDisabled gets a reference to the given NullableString and assigns it to the DateDisabled field.
func (o *ListAccounts200ResponseInner) SetDateDisabled(v string) {
	o.DateDisabled.Set(&v)
}
// SetDateDisabledNil sets the value for DateDisabled to be an explicit nil
func (o *ListAccounts200ResponseInner) SetDateDisabledNil() {
	o.DateDisabled.Set(nil)
}

// UnsetDateDisabled ensures that no value is present for DateDisabled, not even an explicit nil
func (o *ListAccounts200ResponseInner) UnsetDateDisabled() {
	o.DateDisabled.Unset()
}

// GetAccountServiceId returns the AccountServiceId field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetAccountServiceId() int32 {
	if o == nil || isNil(o.AccountServiceId) {
		var ret int32
		return ret
	}
	return *o.AccountServiceId
}

// GetAccountServiceIdOk returns a tuple with the AccountServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetAccountServiceIdOk() (*int32, bool) {
	if o == nil || isNil(o.AccountServiceId) {
		return nil, false
	}
	return o.AccountServiceId, true
}

// HasAccountServiceId returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasAccountServiceId() bool {
	if o != nil && !isNil(o.AccountServiceId) {
		return true
	}

	return false
}

// SetAccountServiceId gets a reference to the given int32 and assigns it to the AccountServiceId field.
func (o *ListAccounts200ResponseInner) SetAccountServiceId(v int32) {
	o.AccountServiceId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetServiceId() int32 {
	if o == nil || isNil(o.ServiceId) {
		var ret int32
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetServiceIdOk() (*int32, bool) {
	if o == nil || isNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int32 and assigns it to the ServiceId field.
func (o *ListAccounts200ResponseInner) SetServiceId(v int32) {
	o.ServiceId = &v
}

// GetPendingPasswordRequestId returns the PendingPasswordRequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAccounts200ResponseInner) GetPendingPasswordRequestId() string {
	if o == nil || isNil(o.PendingPasswordRequestId.Get()) {
		var ret string
		return ret
	}
	return *o.PendingPasswordRequestId.Get()
}

// GetPendingPasswordRequestIdOk returns a tuple with the PendingPasswordRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAccounts200ResponseInner) GetPendingPasswordRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingPasswordRequestId.Get(), o.PendingPasswordRequestId.IsSet()
}

// HasPendingPasswordRequestId returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasPendingPasswordRequestId() bool {
	if o != nil && o.PendingPasswordRequestId.IsSet() {
		return true
	}

	return false
}

// SetPendingPasswordRequestId gets a reference to the given NullableString and assigns it to the PendingPasswordRequestId field.
func (o *ListAccounts200ResponseInner) SetPendingPasswordRequestId(v string) {
	o.PendingPasswordRequestId.Set(&v)
}
// SetPendingPasswordRequestIdNil sets the value for PendingPasswordRequestId to be an explicit nil
func (o *ListAccounts200ResponseInner) SetPendingPasswordRequestIdNil() {
	o.PendingPasswordRequestId.Set(nil)
}

// UnsetPendingPasswordRequestId ensures that no value is present for PendingPasswordRequestId, not even an explicit nil
func (o *ListAccounts200ResponseInner) UnsetPendingPasswordRequestId() {
	o.PendingPasswordRequestId.Unset()
}

// GetPasswordChangeStatus returns the PasswordChangeStatus field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetPasswordChangeStatus() string {
	if o == nil || isNil(o.PasswordChangeStatus) {
		var ret string
		return ret
	}
	return *o.PasswordChangeStatus
}

// GetPasswordChangeStatusOk returns a tuple with the PasswordChangeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetPasswordChangeStatusOk() (*string, bool) {
	if o == nil || isNil(o.PasswordChangeStatus) {
		return nil, false
	}
	return o.PasswordChangeStatus, true
}

// HasPasswordChangeStatus returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasPasswordChangeStatus() bool {
	if o != nil && !isNil(o.PasswordChangeStatus) {
		return true
	}

	return false
}

// SetPasswordChangeStatus gets a reference to the given string and assigns it to the PasswordChangeStatus field.
func (o *ListAccounts200ResponseInner) SetPasswordChangeStatus(v string) {
	o.PasswordChangeStatus = &v
}

// GetPasswordChangeResult returns the PasswordChangeResult field value if set, zero value otherwise.
func (o *ListAccounts200ResponseInner) GetPasswordChangeResult() ListAccounts200ResponseInnerPasswordChangeResult {
	if o == nil || isNil(o.PasswordChangeResult) {
		var ret ListAccounts200ResponseInnerPasswordChangeResult
		return ret
	}
	return *o.PasswordChangeResult
}

// GetPasswordChangeResultOk returns a tuple with the PasswordChangeResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseInner) GetPasswordChangeResultOk() (*ListAccounts200ResponseInnerPasswordChangeResult, bool) {
	if o == nil || isNil(o.PasswordChangeResult) {
		return nil, false
	}
	return o.PasswordChangeResult, true
}

// HasPasswordChangeResult returns a boolean if a field has been set.
func (o *ListAccounts200ResponseInner) HasPasswordChangeResult() bool {
	if o != nil && !isNil(o.PasswordChangeResult) {
		return true
	}

	return false
}

// SetPasswordChangeResult gets a reference to the given ListAccounts200ResponseInnerPasswordChangeResult and assigns it to the PasswordChangeResult field.
func (o *ListAccounts200ResponseInner) SetPasswordChangeResult(v ListAccounts200ResponseInnerPasswordChangeResult) {
	o.PasswordChangeResult = &v
}

func (o ListAccounts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAccounts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.PasswordRequired) {
		toSerialize["passwordRequired"] = o.PasswordRequired
	}
	if !isNil(o.PasswordProvided) {
		toSerialize["passwordProvided"] = o.PasswordProvided
	}
	if !isNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !isNil(o.SsoMethod) {
		toSerialize["ssoMethod"] = o.SsoMethod
	}
	if !isNil(o.IdEncryption) {
		toSerialize["idEncryption"] = o.IdEncryption
	}
	if !isNil(o.PasswordEncryption) {
		toSerialize["passwordEncryption"] = o.PasswordEncryption
	}
	if o.LastPasswdChange.IsSet() {
		toSerialize["lastPasswdChange"] = o.LastPasswdChange.Get()
	}
	if !isNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.DateDisabled.IsSet() {
		toSerialize["dateDisabled"] = o.DateDisabled.Get()
	}
	if !isNil(o.AccountServiceId) {
		toSerialize["accountServiceId"] = o.AccountServiceId
	}
	if !isNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if o.PendingPasswordRequestId.IsSet() {
		toSerialize["pendingPasswordRequestId"] = o.PendingPasswordRequestId.Get()
	}
	if !isNil(o.PasswordChangeStatus) {
		toSerialize["passwordChangeStatus"] = o.PasswordChangeStatus
	}
	if !isNil(o.PasswordChangeResult) {
		toSerialize["passwordChangeResult"] = o.PasswordChangeResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAccounts200ResponseInner) UnmarshalJSON(bytes []byte) (err error) {
	varListAccounts200ResponseInner := _ListAccounts200ResponseInner{}

	if err = json.Unmarshal(bytes, &varListAccounts200ResponseInner); err == nil {
		*o = ListAccounts200ResponseInner(varListAccounts200ResponseInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "username")
		delete(additionalProperties, "passwordRequired")
		delete(additionalProperties, "passwordProvided")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "ssoMethod")
		delete(additionalProperties, "idEncryption")
		delete(additionalProperties, "passwordEncryption")
		delete(additionalProperties, "lastPasswdChange")
		delete(additionalProperties, "serviceName")
		delete(additionalProperties, "dateDisabled")
		delete(additionalProperties, "accountServiceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "pendingPasswordRequestId")
		delete(additionalProperties, "passwordChangeStatus")
		delete(additionalProperties, "passwordChangeResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAccounts200ResponseInner struct {
	value *ListAccounts200ResponseInner
	isSet bool
}

func (v NullableListAccounts200ResponseInner) Get() *ListAccounts200ResponseInner {
	return v.value
}

func (v *NullableListAccounts200ResponseInner) Set(val *ListAccounts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccounts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccounts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccounts200ResponseInner(val *ListAccounts200ResponseInner) *NullableListAccounts200ResponseInner {
	return &NullableListAccounts200ResponseInner{value: val, isSet: true}
}

func (v NullableListAccounts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccounts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


