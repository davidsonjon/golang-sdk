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

// checks if the AuthUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthUser{}

// AuthUser struct for AuthUser
type AuthUser struct {
	// Tenant name.
	Tenant *string `json:"tenant,omitempty"`
	// Identity ID.
	Id *string `json:"id,omitempty"`
	// Identity unique identitifier.
	Uid *string `json:"uid,omitempty"`
	// ID of the auth profile associated with this auth user.
	Profile *string `json:"profile,omitempty"`
	// Auth user employee number.
	IdentificationNumber *string `json:"identificationNumber,omitempty"`
	// Auth user's email.
	Email *string `json:"email,omitempty"`
	// Auth user's phone number.
	Phone *string `json:"phone,omitempty"`
	// Auth user's work phone number.
	WorkPhone *string `json:"workPhone,omitempty"`
	// Auth user's personal email.
	PersonalEmail *string `json:"personalEmail,omitempty"`
	// Auth user's first name.
	Firstname *string `json:"firstname,omitempty"`
	// Auth user's last name.
	Lastname *string `json:"lastname,omitempty"`
	// Auth user's name in displayed format.
	DisplayName *string `json:"displayName,omitempty"`
	// Auth user's alias.
	Alias *string `json:"alias,omitempty"`
	// the date of last password change
	LastPasswordChangeDate *string `json:"lastPasswordChangeDate,omitempty"`
	// Timestamp of the last login (long type value).
	LastLoginTimestamp *int64 `json:"lastLoginTimestamp,omitempty"`
	// Timestamp of the current login (long type value).
	CurrentLoginTimestamp *int64 `json:"currentLoginTimestamp,omitempty"`
	// Array of capabilities for this auth user.
	Capabilities []string `json:"capabilities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthUser AuthUser

// NewAuthUser instantiates a new AuthUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthUser() *AuthUser {
	this := AuthUser{}
	return &this
}

// NewAuthUserWithDefaults instantiates a new AuthUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthUserWithDefaults() *AuthUser {
	this := AuthUser{}
	return &this
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AuthUser) GetTenant() string {
	if o == nil || isNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetTenantOk() (*string, bool) {
	if o == nil || isNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AuthUser) HasTenant() bool {
	if o != nil && !isNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AuthUser) SetTenant(v string) {
	o.Tenant = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthUser) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthUser) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthUser) SetId(v string) {
	o.Id = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *AuthUser) GetUid() string {
	if o == nil || isNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetUidOk() (*string, bool) {
	if o == nil || isNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *AuthUser) HasUid() bool {
	if o != nil && !isNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *AuthUser) SetUid(v string) {
	o.Uid = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AuthUser) GetProfile() string {
	if o == nil || isNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetProfileOk() (*string, bool) {
	if o == nil || isNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AuthUser) HasProfile() bool {
	if o != nil && !isNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *AuthUser) SetProfile(v string) {
	o.Profile = &v
}

// GetIdentificationNumber returns the IdentificationNumber field value if set, zero value otherwise.
func (o *AuthUser) GetIdentificationNumber() string {
	if o == nil || isNil(o.IdentificationNumber) {
		var ret string
		return ret
	}
	return *o.IdentificationNumber
}

// GetIdentificationNumberOk returns a tuple with the IdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetIdentificationNumberOk() (*string, bool) {
	if o == nil || isNil(o.IdentificationNumber) {
		return nil, false
	}
	return o.IdentificationNumber, true
}

// HasIdentificationNumber returns a boolean if a field has been set.
func (o *AuthUser) HasIdentificationNumber() bool {
	if o != nil && !isNil(o.IdentificationNumber) {
		return true
	}

	return false
}

// SetIdentificationNumber gets a reference to the given string and assigns it to the IdentificationNumber field.
func (o *AuthUser) SetIdentificationNumber(v string) {
	o.IdentificationNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthUser) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthUser) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthUser) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AuthUser) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AuthUser) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AuthUser) SetPhone(v string) {
	o.Phone = &v
}

// GetWorkPhone returns the WorkPhone field value if set, zero value otherwise.
func (o *AuthUser) GetWorkPhone() string {
	if o == nil || isNil(o.WorkPhone) {
		var ret string
		return ret
	}
	return *o.WorkPhone
}

// GetWorkPhoneOk returns a tuple with the WorkPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetWorkPhoneOk() (*string, bool) {
	if o == nil || isNil(o.WorkPhone) {
		return nil, false
	}
	return o.WorkPhone, true
}

// HasWorkPhone returns a boolean if a field has been set.
func (o *AuthUser) HasWorkPhone() bool {
	if o != nil && !isNil(o.WorkPhone) {
		return true
	}

	return false
}

// SetWorkPhone gets a reference to the given string and assigns it to the WorkPhone field.
func (o *AuthUser) SetWorkPhone(v string) {
	o.WorkPhone = &v
}

// GetPersonalEmail returns the PersonalEmail field value if set, zero value otherwise.
func (o *AuthUser) GetPersonalEmail() string {
	if o == nil || isNil(o.PersonalEmail) {
		var ret string
		return ret
	}
	return *o.PersonalEmail
}

// GetPersonalEmailOk returns a tuple with the PersonalEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetPersonalEmailOk() (*string, bool) {
	if o == nil || isNil(o.PersonalEmail) {
		return nil, false
	}
	return o.PersonalEmail, true
}

// HasPersonalEmail returns a boolean if a field has been set.
func (o *AuthUser) HasPersonalEmail() bool {
	if o != nil && !isNil(o.PersonalEmail) {
		return true
	}

	return false
}

// SetPersonalEmail gets a reference to the given string and assigns it to the PersonalEmail field.
func (o *AuthUser) SetPersonalEmail(v string) {
	o.PersonalEmail = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *AuthUser) GetFirstname() string {
	if o == nil || isNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetFirstnameOk() (*string, bool) {
	if o == nil || isNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *AuthUser) HasFirstname() bool {
	if o != nil && !isNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *AuthUser) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *AuthUser) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *AuthUser) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *AuthUser) SetLastname(v string) {
	o.Lastname = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthUser) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthUser) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *AuthUser) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *AuthUser) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *AuthUser) SetAlias(v string) {
	o.Alias = &v
}

// GetLastPasswordChangeDate returns the LastPasswordChangeDate field value if set, zero value otherwise.
func (o *AuthUser) GetLastPasswordChangeDate() string {
	if o == nil || isNil(o.LastPasswordChangeDate) {
		var ret string
		return ret
	}
	return *o.LastPasswordChangeDate
}

// GetLastPasswordChangeDateOk returns a tuple with the LastPasswordChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetLastPasswordChangeDateOk() (*string, bool) {
	if o == nil || isNil(o.LastPasswordChangeDate) {
		return nil, false
	}
	return o.LastPasswordChangeDate, true
}

// HasLastPasswordChangeDate returns a boolean if a field has been set.
func (o *AuthUser) HasLastPasswordChangeDate() bool {
	if o != nil && !isNil(o.LastPasswordChangeDate) {
		return true
	}

	return false
}

// SetLastPasswordChangeDate gets a reference to the given string and assigns it to the LastPasswordChangeDate field.
func (o *AuthUser) SetLastPasswordChangeDate(v string) {
	o.LastPasswordChangeDate = &v
}

// GetLastLoginTimestamp returns the LastLoginTimestamp field value if set, zero value otherwise.
func (o *AuthUser) GetLastLoginTimestamp() int64 {
	if o == nil || isNil(o.LastLoginTimestamp) {
		var ret int64
		return ret
	}
	return *o.LastLoginTimestamp
}

// GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetLastLoginTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.LastLoginTimestamp) {
		return nil, false
	}
	return o.LastLoginTimestamp, true
}

// HasLastLoginTimestamp returns a boolean if a field has been set.
func (o *AuthUser) HasLastLoginTimestamp() bool {
	if o != nil && !isNil(o.LastLoginTimestamp) {
		return true
	}

	return false
}

// SetLastLoginTimestamp gets a reference to the given int64 and assigns it to the LastLoginTimestamp field.
func (o *AuthUser) SetLastLoginTimestamp(v int64) {
	o.LastLoginTimestamp = &v
}

// GetCurrentLoginTimestamp returns the CurrentLoginTimestamp field value if set, zero value otherwise.
func (o *AuthUser) GetCurrentLoginTimestamp() int64 {
	if o == nil || isNil(o.CurrentLoginTimestamp) {
		var ret int64
		return ret
	}
	return *o.CurrentLoginTimestamp
}

// GetCurrentLoginTimestampOk returns a tuple with the CurrentLoginTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetCurrentLoginTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.CurrentLoginTimestamp) {
		return nil, false
	}
	return o.CurrentLoginTimestamp, true
}

// HasCurrentLoginTimestamp returns a boolean if a field has been set.
func (o *AuthUser) HasCurrentLoginTimestamp() bool {
	if o != nil && !isNil(o.CurrentLoginTimestamp) {
		return true
	}

	return false
}

// SetCurrentLoginTimestamp gets a reference to the given int64 and assigns it to the CurrentLoginTimestamp field.
func (o *AuthUser) SetCurrentLoginTimestamp(v int64) {
	o.CurrentLoginTimestamp = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AuthUser) GetCapabilities() []string {
	if o == nil || isNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUser) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || isNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AuthUser) HasCapabilities() bool {
	if o != nil && !isNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *AuthUser) SetCapabilities(v []string) {
	o.Capabilities = v
}

func (o AuthUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !isNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !isNil(o.IdentificationNumber) {
		toSerialize["identificationNumber"] = o.IdentificationNumber
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.WorkPhone) {
		toSerialize["workPhone"] = o.WorkPhone
	}
	if !isNil(o.PersonalEmail) {
		toSerialize["personalEmail"] = o.PersonalEmail
	}
	if !isNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !isNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !isNil(o.LastPasswordChangeDate) {
		toSerialize["lastPasswordChangeDate"] = o.LastPasswordChangeDate
	}
	if !isNil(o.LastLoginTimestamp) {
		toSerialize["lastLoginTimestamp"] = o.LastLoginTimestamp
	}
	if !isNil(o.CurrentLoginTimestamp) {
		toSerialize["currentLoginTimestamp"] = o.CurrentLoginTimestamp
	}
	if !isNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthUser) UnmarshalJSON(bytes []byte) (err error) {
	varAuthUser := _AuthUser{}

	if err = json.Unmarshal(bytes, &varAuthUser); err == nil {
	*o = AuthUser(varAuthUser)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "id")
		delete(additionalProperties, "uid")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "identificationNumber")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "workPhone")
		delete(additionalProperties, "personalEmail")
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "alias")
		delete(additionalProperties, "lastPasswordChangeDate")
		delete(additionalProperties, "lastLoginTimestamp")
		delete(additionalProperties, "currentLoginTimestamp")
		delete(additionalProperties, "capabilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthUser struct {
	value *AuthUser
	isSet bool
}

func (v NullableAuthUser) Get() *AuthUser {
	return v.value
}

func (v *NullableAuthUser) Set(val *AuthUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthUser(val *AuthUser) *NullableAuthUser {
	return &NullableAuthUser{value: val, isSet: true}
}

func (v NullableAuthUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


