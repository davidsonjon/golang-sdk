/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the ServiceProviderConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProviderConfiguration{}

// ServiceProviderConfiguration Represents the IdentityNow as Service Provider Configuration allowing customers to log into IDN via an Identity Provider
type ServiceProviderConfiguration struct {
	// This determines whether or not the SAML authentication flow is enabled for an org
	Enabled *bool `json:"enabled,omitempty"`
	// This allows basic login with the parameter prompt=true. This is often toggled on when debugging SAML authentication setup. When false, only org admins with MFA-enabled can bypass the IDP.
	BypassIdp *bool `json:"bypassIdp,omitempty"`
	// This indicates whether or not the SAML configuration is valid.
	SamlConfigurationValid *bool `json:"samlConfigurationValid,omitempty"`
	// A list of the abstract implementations of the Federation Protocol details. Typically, this will include on SpDetails object and one IdpDetails object used in tandem to define a SAML integration between a customer's identity provider and a customer's SailPoint instance (i.e., the service provider).
	FederationProtocolDetails []ServiceProviderConfigurationFederationProtocolDetailsInner `json:"federationProtocolDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProviderConfiguration ServiceProviderConfiguration

// NewServiceProviderConfiguration instantiates a new ServiceProviderConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProviderConfiguration() *ServiceProviderConfiguration {
	this := ServiceProviderConfiguration{}
	var enabled bool = false
	this.Enabled = &enabled
	var bypassIdp bool = false
	this.BypassIdp = &bypassIdp
	var samlConfigurationValid bool = false
	this.SamlConfigurationValid = &samlConfigurationValid
	return &this
}

// NewServiceProviderConfigurationWithDefaults instantiates a new ServiceProviderConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProviderConfigurationWithDefaults() *ServiceProviderConfiguration {
	this := ServiceProviderConfiguration{}
	var enabled bool = false
	this.Enabled = &enabled
	var bypassIdp bool = false
	this.BypassIdp = &bypassIdp
	var samlConfigurationValid bool = false
	this.SamlConfigurationValid = &samlConfigurationValid
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceProviderConfiguration) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceProviderConfiguration) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceProviderConfiguration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetBypassIdp returns the BypassIdp field value if set, zero value otherwise.
func (o *ServiceProviderConfiguration) GetBypassIdp() bool {
	if o == nil || isNil(o.BypassIdp) {
		var ret bool
		return ret
	}
	return *o.BypassIdp
}

// GetBypassIdpOk returns a tuple with the BypassIdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderConfiguration) GetBypassIdpOk() (*bool, bool) {
	if o == nil || isNil(o.BypassIdp) {
		return nil, false
	}
	return o.BypassIdp, true
}

// HasBypassIdp returns a boolean if a field has been set.
func (o *ServiceProviderConfiguration) HasBypassIdp() bool {
	if o != nil && !isNil(o.BypassIdp) {
		return true
	}

	return false
}

// SetBypassIdp gets a reference to the given bool and assigns it to the BypassIdp field.
func (o *ServiceProviderConfiguration) SetBypassIdp(v bool) {
	o.BypassIdp = &v
}

// GetSamlConfigurationValid returns the SamlConfigurationValid field value if set, zero value otherwise.
func (o *ServiceProviderConfiguration) GetSamlConfigurationValid() bool {
	if o == nil || isNil(o.SamlConfigurationValid) {
		var ret bool
		return ret
	}
	return *o.SamlConfigurationValid
}

// GetSamlConfigurationValidOk returns a tuple with the SamlConfigurationValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderConfiguration) GetSamlConfigurationValidOk() (*bool, bool) {
	if o == nil || isNil(o.SamlConfigurationValid) {
		return nil, false
	}
	return o.SamlConfigurationValid, true
}

// HasSamlConfigurationValid returns a boolean if a field has been set.
func (o *ServiceProviderConfiguration) HasSamlConfigurationValid() bool {
	if o != nil && !isNil(o.SamlConfigurationValid) {
		return true
	}

	return false
}

// SetSamlConfigurationValid gets a reference to the given bool and assigns it to the SamlConfigurationValid field.
func (o *ServiceProviderConfiguration) SetSamlConfigurationValid(v bool) {
	o.SamlConfigurationValid = &v
}

// GetFederationProtocolDetails returns the FederationProtocolDetails field value if set, zero value otherwise.
func (o *ServiceProviderConfiguration) GetFederationProtocolDetails() []ServiceProviderConfigurationFederationProtocolDetailsInner {
	if o == nil || isNil(o.FederationProtocolDetails) {
		var ret []ServiceProviderConfigurationFederationProtocolDetailsInner
		return ret
	}
	return o.FederationProtocolDetails
}

// GetFederationProtocolDetailsOk returns a tuple with the FederationProtocolDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderConfiguration) GetFederationProtocolDetailsOk() ([]ServiceProviderConfigurationFederationProtocolDetailsInner, bool) {
	if o == nil || isNil(o.FederationProtocolDetails) {
		return nil, false
	}
	return o.FederationProtocolDetails, true
}

// HasFederationProtocolDetails returns a boolean if a field has been set.
func (o *ServiceProviderConfiguration) HasFederationProtocolDetails() bool {
	if o != nil && !isNil(o.FederationProtocolDetails) {
		return true
	}

	return false
}

// SetFederationProtocolDetails gets a reference to the given []ServiceProviderConfigurationFederationProtocolDetailsInner and assigns it to the FederationProtocolDetails field.
func (o *ServiceProviderConfiguration) SetFederationProtocolDetails(v []ServiceProviderConfigurationFederationProtocolDetailsInner) {
	o.FederationProtocolDetails = v
}

func (o ServiceProviderConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProviderConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.BypassIdp) {
		toSerialize["bypassIdp"] = o.BypassIdp
	}
	if !isNil(o.SamlConfigurationValid) {
		toSerialize["samlConfigurationValid"] = o.SamlConfigurationValid
	}
	if !isNil(o.FederationProtocolDetails) {
		toSerialize["federationProtocolDetails"] = o.FederationProtocolDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProviderConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varServiceProviderConfiguration := _ServiceProviderConfiguration{}

	if err = json.Unmarshal(bytes, &varServiceProviderConfiguration); err == nil {
	*o = ServiceProviderConfiguration(varServiceProviderConfiguration)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "bypassIdp")
		delete(additionalProperties, "samlConfigurationValid")
		delete(additionalProperties, "federationProtocolDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProviderConfiguration struct {
	value *ServiceProviderConfiguration
	isSet bool
}

func (v NullableServiceProviderConfiguration) Get() *ServiceProviderConfiguration {
	return v.value
}

func (v *NullableServiceProviderConfiguration) Set(val *ServiceProviderConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProviderConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProviderConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProviderConfiguration(val *ServiceProviderConfiguration) *NullableServiceProviderConfiguration {
	return &NullableServiceProviderConfiguration{value: val, isSet: true}
}

func (v NullableServiceProviderConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProviderConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


