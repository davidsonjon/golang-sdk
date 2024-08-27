/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the AuthProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthProfile{}

// AuthProfile struct for AuthProfile
type AuthProfile struct {
	// Authentication Profile name.
	Name *string `json:"name,omitempty"`
	// Use it to block access from off network.
	OffNetwork *bool `json:"offNetwork,omitempty"`
	// Use it to block access from untrusted geoographies.
	UntrustedGeography *bool `json:"untrustedGeography,omitempty"`
	// Application ID.
	ApplicationId *string `json:"applicationId,omitempty"`
	// Application name.
	ApplicationName *string `json:"applicationName,omitempty"`
	// Type of the Authentication Profile.
	Type *string `json:"type,omitempty"`
	// Use it to enable strong authentication.
	StrongAuthLogin *bool `json:"strongAuthLogin,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthProfile AuthProfile

// NewAuthProfile instantiates a new AuthProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthProfile() *AuthProfile {
	this := AuthProfile{}
	var offNetwork bool = false
	this.OffNetwork = &offNetwork
	var untrustedGeography bool = false
	this.UntrustedGeography = &untrustedGeography
	var strongAuthLogin bool = false
	this.StrongAuthLogin = &strongAuthLogin
	return &this
}

// NewAuthProfileWithDefaults instantiates a new AuthProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthProfileWithDefaults() *AuthProfile {
	this := AuthProfile{}
	var offNetwork bool = false
	this.OffNetwork = &offNetwork
	var untrustedGeography bool = false
	this.UntrustedGeography = &untrustedGeography
	var strongAuthLogin bool = false
	this.StrongAuthLogin = &strongAuthLogin
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthProfile) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthProfile) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthProfile) SetName(v string) {
	o.Name = &v
}

// GetOffNetwork returns the OffNetwork field value if set, zero value otherwise.
func (o *AuthProfile) GetOffNetwork() bool {
	if o == nil || isNil(o.OffNetwork) {
		var ret bool
		return ret
	}
	return *o.OffNetwork
}

// GetOffNetworkOk returns a tuple with the OffNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetOffNetworkOk() (*bool, bool) {
	if o == nil || isNil(o.OffNetwork) {
		return nil, false
	}
	return o.OffNetwork, true
}

// HasOffNetwork returns a boolean if a field has been set.
func (o *AuthProfile) HasOffNetwork() bool {
	if o != nil && !isNil(o.OffNetwork) {
		return true
	}

	return false
}

// SetOffNetwork gets a reference to the given bool and assigns it to the OffNetwork field.
func (o *AuthProfile) SetOffNetwork(v bool) {
	o.OffNetwork = &v
}

// GetUntrustedGeography returns the UntrustedGeography field value if set, zero value otherwise.
func (o *AuthProfile) GetUntrustedGeography() bool {
	if o == nil || isNil(o.UntrustedGeography) {
		var ret bool
		return ret
	}
	return *o.UntrustedGeography
}

// GetUntrustedGeographyOk returns a tuple with the UntrustedGeography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetUntrustedGeographyOk() (*bool, bool) {
	if o == nil || isNil(o.UntrustedGeography) {
		return nil, false
	}
	return o.UntrustedGeography, true
}

// HasUntrustedGeography returns a boolean if a field has been set.
func (o *AuthProfile) HasUntrustedGeography() bool {
	if o != nil && !isNil(o.UntrustedGeography) {
		return true
	}

	return false
}

// SetUntrustedGeography gets a reference to the given bool and assigns it to the UntrustedGeography field.
func (o *AuthProfile) SetUntrustedGeography(v bool) {
	o.UntrustedGeography = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AuthProfile) GetApplicationId() string {
	if o == nil || isNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetApplicationIdOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AuthProfile) HasApplicationId() bool {
	if o != nil && !isNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *AuthProfile) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AuthProfile) GetApplicationName() string {
	if o == nil || isNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetApplicationNameOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AuthProfile) HasApplicationName() bool {
	if o != nil && !isNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AuthProfile) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthProfile) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthProfile) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthProfile) SetType(v string) {
	o.Type = &v
}

// GetStrongAuthLogin returns the StrongAuthLogin field value if set, zero value otherwise.
func (o *AuthProfile) GetStrongAuthLogin() bool {
	if o == nil || isNil(o.StrongAuthLogin) {
		var ret bool
		return ret
	}
	return *o.StrongAuthLogin
}

// GetStrongAuthLoginOk returns a tuple with the StrongAuthLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfile) GetStrongAuthLoginOk() (*bool, bool) {
	if o == nil || isNil(o.StrongAuthLogin) {
		return nil, false
	}
	return o.StrongAuthLogin, true
}

// HasStrongAuthLogin returns a boolean if a field has been set.
func (o *AuthProfile) HasStrongAuthLogin() bool {
	if o != nil && !isNil(o.StrongAuthLogin) {
		return true
	}

	return false
}

// SetStrongAuthLogin gets a reference to the given bool and assigns it to the StrongAuthLogin field.
func (o *AuthProfile) SetStrongAuthLogin(v bool) {
	o.StrongAuthLogin = &v
}

func (o AuthProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OffNetwork) {
		toSerialize["offNetwork"] = o.OffNetwork
	}
	if !isNil(o.UntrustedGeography) {
		toSerialize["untrustedGeography"] = o.UntrustedGeography
	}
	if !isNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !isNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.StrongAuthLogin) {
		toSerialize["strongAuthLogin"] = o.StrongAuthLogin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthProfile) UnmarshalJSON(bytes []byte) (err error) {
	varAuthProfile := _AuthProfile{}

	if err = json.Unmarshal(bytes, &varAuthProfile); err == nil {
	*o = AuthProfile(varAuthProfile)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "offNetwork")
		delete(additionalProperties, "untrustedGeography")
		delete(additionalProperties, "applicationId")
		delete(additionalProperties, "applicationName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "strongAuthLogin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthProfile struct {
	value *AuthProfile
	isSet bool
}

func (v NullableAuthProfile) Get() *AuthProfile {
	return v.value
}

func (v *NullableAuthProfile) Set(val *AuthProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthProfile(val *AuthProfile) *NullableAuthProfile {
	return &NullableAuthProfile{value: val, isSet: true}
}

func (v NullableAuthProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

