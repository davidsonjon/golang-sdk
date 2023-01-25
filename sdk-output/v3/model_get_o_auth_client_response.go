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

// GetOAuthClientResponse struct for GetOAuthClientResponse
type GetOAuthClientResponse struct {
	// ID of the OAuth client
	Id string `json:"id"`
	// The name of the business the API Client should belong to
	BusinessName NullableString `json:"businessName"`
	// The homepage URL associated with the owner of the API Client
	HomepageUrl NullableString `json:"homepageUrl"`
	// A human-readable name for the API Client
	Name string `json:"name"`
	// A description of the API Client
	Description NullableString `json:"description"`
	// The number of seconds an access token generated for this API Client is valid for
	AccessTokenValiditySeconds int32 `json:"accessTokenValiditySeconds"`
	// The number of seconds a refresh token generated for this API Client is valid for
	RefreshTokenValiditySeconds int32 `json:"refreshTokenValiditySeconds"`
	// A list of the approved redirect URIs used with the authorization_code flow
	RedirectUris []string `json:"redirectUris"`
	// A list of OAuth 2.0 grant types this API Client can be used with
	GrantTypes []GrantType `json:"grantTypes"`
	AccessType AccessType `json:"accessType"`
	Type ClientType `json:"type"`
	// An indicator of whether the API Client can be used for requests internal to IDN
	Internal bool `json:"internal"`
	// An indicator of whether the API Client is enabled for use
	Enabled bool `json:"enabled"`
	// An indicator of whether the API Client supports strong authentication
	StrongAuthSupported bool `json:"strongAuthSupported"`
	// An indicator of whether the API Client supports the serialization of SAML claims when used with the authorization_code flow
	ClaimsSupported bool `json:"claimsSupported"`
	// The date and time, down to the millisecond, when the API Client was created
	Created time.Time `json:"created"`
	// The date and time, down to the millisecond, when the API Client was last updated
	Modified time.Time `json:"modified"`
	// Scopes of the API Client.
	Scope []string `json:"scope"`
	AdditionalProperties map[string]interface{}
}

type _GetOAuthClientResponse GetOAuthClientResponse

// NewGetOAuthClientResponse instantiates a new GetOAuthClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOAuthClientResponse(id string, businessName NullableString, homepageUrl NullableString, name string, description NullableString, accessTokenValiditySeconds int32, refreshTokenValiditySeconds int32, redirectUris []string, grantTypes []GrantType, accessType AccessType, type_ ClientType, internal bool, enabled bool, strongAuthSupported bool, claimsSupported bool, created time.Time, modified time.Time, scope []string) *GetOAuthClientResponse {
	this := GetOAuthClientResponse{}
	this.Id = id
	this.BusinessName = businessName
	this.HomepageUrl = homepageUrl
	this.Name = name
	this.Description = description
	this.AccessTokenValiditySeconds = accessTokenValiditySeconds
	this.RefreshTokenValiditySeconds = refreshTokenValiditySeconds
	this.RedirectUris = redirectUris
	this.GrantTypes = grantTypes
	this.AccessType = accessType
	this.Type = type_
	this.Internal = internal
	this.Enabled = enabled
	this.StrongAuthSupported = strongAuthSupported
	this.ClaimsSupported = claimsSupported
	this.Created = created
	this.Modified = modified
	this.Scope = scope
	return &this
}

// NewGetOAuthClientResponseWithDefaults instantiates a new GetOAuthClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOAuthClientResponseWithDefaults() *GetOAuthClientResponse {
	this := GetOAuthClientResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GetOAuthClientResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetOAuthClientResponse) SetId(v string) {
	o.Id = v
}

// GetBusinessName returns the BusinessName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetOAuthClientResponse) GetBusinessName() string {
	if o == nil || o.BusinessName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BusinessName.Get()
}

// GetBusinessNameOk returns a tuple with the BusinessName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOAuthClientResponse) GetBusinessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessName.Get(), o.BusinessName.IsSet()
}

// SetBusinessName sets field value
func (o *GetOAuthClientResponse) SetBusinessName(v string) {
	o.BusinessName.Set(&v)
}

// GetHomepageUrl returns the HomepageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetOAuthClientResponse) GetHomepageUrl() string {
	if o == nil || o.HomepageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.HomepageUrl.Get()
}

// GetHomepageUrlOk returns a tuple with the HomepageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOAuthClientResponse) GetHomepageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HomepageUrl.Get(), o.HomepageUrl.IsSet()
}

// SetHomepageUrl sets field value
func (o *GetOAuthClientResponse) SetHomepageUrl(v string) {
	o.HomepageUrl.Set(&v)
}

// GetName returns the Name field value
func (o *GetOAuthClientResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetOAuthClientResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetOAuthClientResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOAuthClientResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *GetOAuthClientResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field value
func (o *GetOAuthClientResponse) GetAccessTokenValiditySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccessTokenValiditySeconds
}

// GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetAccessTokenValiditySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenValiditySeconds, true
}

// SetAccessTokenValiditySeconds sets field value
func (o *GetOAuthClientResponse) SetAccessTokenValiditySeconds(v int32) {
	o.AccessTokenValiditySeconds = v
}

// GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field value
func (o *GetOAuthClientResponse) GetRefreshTokenValiditySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RefreshTokenValiditySeconds
}

// GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetRefreshTokenValiditySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshTokenValiditySeconds, true
}

// SetRefreshTokenValiditySeconds sets field value
func (o *GetOAuthClientResponse) SetRefreshTokenValiditySeconds(v int32) {
	o.RefreshTokenValiditySeconds = v
}

// GetRedirectUris returns the RedirectUris field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GetOAuthClientResponse) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOAuthClientResponse) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || isNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *GetOAuthClientResponse) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetGrantTypes returns the GrantTypes field value
func (o *GetOAuthClientResponse) GetGrantTypes() []GrantType {
	if o == nil {
		var ret []GrantType
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetGrantTypesOk() ([]GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *GetOAuthClientResponse) SetGrantTypes(v []GrantType) {
	o.GrantTypes = v
}

// GetAccessType returns the AccessType field value
func (o *GetOAuthClientResponse) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *GetOAuthClientResponse) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetType returns the Type field value
func (o *GetOAuthClientResponse) GetType() ClientType {
	if o == nil {
		var ret ClientType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetTypeOk() (*ClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetOAuthClientResponse) SetType(v ClientType) {
	o.Type = v
}

// GetInternal returns the Internal field value
func (o *GetOAuthClientResponse) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *GetOAuthClientResponse) SetInternal(v bool) {
	o.Internal = v
}

// GetEnabled returns the Enabled field value
func (o *GetOAuthClientResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetOAuthClientResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetStrongAuthSupported returns the StrongAuthSupported field value
func (o *GetOAuthClientResponse) GetStrongAuthSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StrongAuthSupported
}

// GetStrongAuthSupportedOk returns a tuple with the StrongAuthSupported field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetStrongAuthSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrongAuthSupported, true
}

// SetStrongAuthSupported sets field value
func (o *GetOAuthClientResponse) SetStrongAuthSupported(v bool) {
	o.StrongAuthSupported = v
}

// GetClaimsSupported returns the ClaimsSupported field value
func (o *GetOAuthClientResponse) GetClaimsSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClaimsSupported
}

// GetClaimsSupportedOk returns a tuple with the ClaimsSupported field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetClaimsSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimsSupported, true
}

// SetClaimsSupported sets field value
func (o *GetOAuthClientResponse) SetClaimsSupported(v bool) {
	o.ClaimsSupported = v
}

// GetCreated returns the Created field value
func (o *GetOAuthClientResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GetOAuthClientResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *GetOAuthClientResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *GetOAuthClientResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *GetOAuthClientResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GetOAuthClientResponse) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOAuthClientResponse) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *GetOAuthClientResponse) SetScope(v []string) {
	o.Scope = v
}

func (o GetOAuthClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["businessName"] = o.BusinessName.Get()
	}
	if true {
		toSerialize["homepageUrl"] = o.HomepageUrl.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["accessTokenValiditySeconds"] = o.AccessTokenValiditySeconds
	}
	if true {
		toSerialize["refreshTokenValiditySeconds"] = o.RefreshTokenValiditySeconds
	}
	if o.RedirectUris != nil {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if true {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["internal"] = o.Internal
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["strongAuthSupported"] = o.StrongAuthSupported
	}
	if true {
		toSerialize["claimsSupported"] = o.ClaimsSupported
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetOAuthClientResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetOAuthClientResponse := _GetOAuthClientResponse{}

	if err = json.Unmarshal(bytes, &varGetOAuthClientResponse); err == nil {
		*o = GetOAuthClientResponse(varGetOAuthClientResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "businessName")
		delete(additionalProperties, "homepageUrl")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "accessTokenValiditySeconds")
		delete(additionalProperties, "refreshTokenValiditySeconds")
		delete(additionalProperties, "redirectUris")
		delete(additionalProperties, "grantTypes")
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "type")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "strongAuthSupported")
		delete(additionalProperties, "claimsSupported")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOAuthClientResponse struct {
	value *GetOAuthClientResponse
	isSet bool
}

func (v NullableGetOAuthClientResponse) Get() *GetOAuthClientResponse {
	return v.value
}

func (v *NullableGetOAuthClientResponse) Set(val *GetOAuthClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOAuthClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOAuthClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOAuthClientResponse(val *GetOAuthClientResponse) *NullableGetOAuthClientResponse {
	return &NullableGetOAuthClientResponse{value: val, isSet: true}
}

func (v NullableGetOAuthClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOAuthClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


