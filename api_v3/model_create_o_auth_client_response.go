/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CreateOAuthClientResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOAuthClientResponse{}

// CreateOAuthClientResponse struct for CreateOAuthClientResponse
type CreateOAuthClientResponse struct {
	// ID of the OAuth client
	Id string `json:"id"`
	// Secret of the OAuth client (This field is only returned on the intial create call.)
	Secret string `json:"secret"`
	// The name of the business the API Client should belong to
	BusinessName string `json:"businessName"`
	// The homepage URL associated with the owner of the API Client
	HomepageUrl string `json:"homepageUrl"`
	// A human-readable name for the API Client
	Name string `json:"name"`
	// A description of the API Client
	Description string `json:"description"`
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

type _CreateOAuthClientResponse CreateOAuthClientResponse

// NewCreateOAuthClientResponse instantiates a new CreateOAuthClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOAuthClientResponse(id string, secret string, businessName string, homepageUrl string, name string, description string, accessTokenValiditySeconds int32, refreshTokenValiditySeconds int32, redirectUris []string, grantTypes []GrantType, accessType AccessType, type_ ClientType, internal bool, enabled bool, strongAuthSupported bool, claimsSupported bool, created time.Time, modified time.Time, scope []string) *CreateOAuthClientResponse {
	this := CreateOAuthClientResponse{}
	this.Id = id
	this.Secret = secret
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

// NewCreateOAuthClientResponseWithDefaults instantiates a new CreateOAuthClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOAuthClientResponseWithDefaults() *CreateOAuthClientResponse {
	this := CreateOAuthClientResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateOAuthClientResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateOAuthClientResponse) SetId(v string) {
	o.Id = v
}

// GetSecret returns the Secret field value
func (o *CreateOAuthClientResponse) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateOAuthClientResponse) SetSecret(v string) {
	o.Secret = v
}

// GetBusinessName returns the BusinessName field value
func (o *CreateOAuthClientResponse) GetBusinessName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetBusinessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessName, true
}

// SetBusinessName sets field value
func (o *CreateOAuthClientResponse) SetBusinessName(v string) {
	o.BusinessName = v
}

// GetHomepageUrl returns the HomepageUrl field value
func (o *CreateOAuthClientResponse) GetHomepageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomepageUrl
}

// GetHomepageUrlOk returns a tuple with the HomepageUrl field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetHomepageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomepageUrl, true
}

// SetHomepageUrl sets field value
func (o *CreateOAuthClientResponse) SetHomepageUrl(v string) {
	o.HomepageUrl = v
}

// GetName returns the Name field value
func (o *CreateOAuthClientResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOAuthClientResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateOAuthClientResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateOAuthClientResponse) SetDescription(v string) {
	o.Description = v
}

// GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field value
func (o *CreateOAuthClientResponse) GetAccessTokenValiditySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccessTokenValiditySeconds
}

// GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetAccessTokenValiditySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenValiditySeconds, true
}

// SetAccessTokenValiditySeconds sets field value
func (o *CreateOAuthClientResponse) SetAccessTokenValiditySeconds(v int32) {
	o.AccessTokenValiditySeconds = v
}

// GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field value
func (o *CreateOAuthClientResponse) GetRefreshTokenValiditySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RefreshTokenValiditySeconds
}

// GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetRefreshTokenValiditySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshTokenValiditySeconds, true
}

// SetRefreshTokenValiditySeconds sets field value
func (o *CreateOAuthClientResponse) SetRefreshTokenValiditySeconds(v int32) {
	o.RefreshTokenValiditySeconds = v
}

// GetRedirectUris returns the RedirectUris field value
func (o *CreateOAuthClientResponse) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetRedirectUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *CreateOAuthClientResponse) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetGrantTypes returns the GrantTypes field value
func (o *CreateOAuthClientResponse) GetGrantTypes() []GrantType {
	if o == nil {
		var ret []GrantType
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetGrantTypesOk() ([]GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *CreateOAuthClientResponse) SetGrantTypes(v []GrantType) {
	o.GrantTypes = v
}

// GetAccessType returns the AccessType field value
func (o *CreateOAuthClientResponse) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *CreateOAuthClientResponse) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetType returns the Type field value
func (o *CreateOAuthClientResponse) GetType() ClientType {
	if o == nil {
		var ret ClientType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetTypeOk() (*ClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOAuthClientResponse) SetType(v ClientType) {
	o.Type = v
}

// GetInternal returns the Internal field value
func (o *CreateOAuthClientResponse) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *CreateOAuthClientResponse) SetInternal(v bool) {
	o.Internal = v
}

// GetEnabled returns the Enabled field value
func (o *CreateOAuthClientResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateOAuthClientResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetStrongAuthSupported returns the StrongAuthSupported field value
func (o *CreateOAuthClientResponse) GetStrongAuthSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StrongAuthSupported
}

// GetStrongAuthSupportedOk returns a tuple with the StrongAuthSupported field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetStrongAuthSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrongAuthSupported, true
}

// SetStrongAuthSupported sets field value
func (o *CreateOAuthClientResponse) SetStrongAuthSupported(v bool) {
	o.StrongAuthSupported = v
}

// GetClaimsSupported returns the ClaimsSupported field value
func (o *CreateOAuthClientResponse) GetClaimsSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClaimsSupported
}

// GetClaimsSupportedOk returns a tuple with the ClaimsSupported field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetClaimsSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimsSupported, true
}

// SetClaimsSupported sets field value
func (o *CreateOAuthClientResponse) SetClaimsSupported(v bool) {
	o.ClaimsSupported = v
}

// GetCreated returns the Created field value
func (o *CreateOAuthClientResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CreateOAuthClientResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *CreateOAuthClientResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *CreateOAuthClientResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CreateOAuthClientResponse) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOAuthClientResponse) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *CreateOAuthClientResponse) SetScope(v []string) {
	o.Scope = v
}

func (o CreateOAuthClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOAuthClientResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["secret"] = o.Secret
	toSerialize["businessName"] = o.BusinessName
	toSerialize["homepageUrl"] = o.HomepageUrl
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["accessTokenValiditySeconds"] = o.AccessTokenValiditySeconds
	toSerialize["refreshTokenValiditySeconds"] = o.RefreshTokenValiditySeconds
	toSerialize["redirectUris"] = o.RedirectUris
	toSerialize["grantTypes"] = o.GrantTypes
	toSerialize["accessType"] = o.AccessType
	toSerialize["type"] = o.Type
	toSerialize["internal"] = o.Internal
	toSerialize["enabled"] = o.Enabled
	toSerialize["strongAuthSupported"] = o.StrongAuthSupported
	toSerialize["claimsSupported"] = o.ClaimsSupported
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOAuthClientResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"secret",
		"businessName",
		"homepageUrl",
		"name",
		"description",
		"accessTokenValiditySeconds",
		"refreshTokenValiditySeconds",
		"redirectUris",
		"grantTypes",
		"accessType",
		"type",
		"internal",
		"enabled",
		"strongAuthSupported",
		"claimsSupported",
		"created",
		"modified",
		"scope",
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

	varCreateOAuthClientResponse := _CreateOAuthClientResponse{}

	if err = json.Unmarshal(bytes, &varCreateOAuthClientResponse); err == nil {
	*o = CreateOAuthClientResponse(varCreateOAuthClientResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "secret")
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

type NullableCreateOAuthClientResponse struct {
	value *CreateOAuthClientResponse
	isSet bool
}

func (v NullableCreateOAuthClientResponse) Get() *CreateOAuthClientResponse {
	return v.value
}

func (v *NullableCreateOAuthClientResponse) Set(val *CreateOAuthClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOAuthClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOAuthClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOAuthClientResponse(val *CreateOAuthClientResponse) *NullableCreateOAuthClientResponse {
	return &NullableCreateOAuthClientResponse{value: val, isSet: true}
}

func (v NullableCreateOAuthClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOAuthClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


