/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the OrgConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgConfig{}

// OrgConfig DTO class for OrgConfig data accessible by customer external org admin (\"ORG_ADMIN\") users
type OrgConfig struct {
	// The name of the org.
	OrgName *string `json:"orgName,omitempty"`
	// The selected time zone which is to be used for the org.  This directly affects when scheduled tasks are executed.  Valid options can be found at /beta/org-config/valid-time-zones
	TimeZone *string `json:"timeZone,omitempty"`
	// Flag to determine whether the LCS_CHANGE_HONORS_SOURCE_ENABLE_FEATURE flag is enabled for the current org.
	LcsChangeHonorsSourceEnableFeature *bool `json:"lcsChangeHonorsSourceEnableFeature,omitempty"`
	// ARM Customer ID
	ArmCustomerId NullableString `json:"armCustomerId,omitempty"`
	// A list of IDN::sourceId to ARM::systemId mappings.
	ArmSapSystemIdMappings NullableString `json:"armSapSystemIdMappings,omitempty"`
	// ARM authentication string
	ArmAuth NullableString `json:"armAuth,omitempty"`
	// ARM database name
	ArmDb NullableString `json:"armDb,omitempty"`
	// ARM SSO URL
	ArmSsoUrl NullableString `json:"armSsoUrl,omitempty"`
	// Flag to determine whether IAI Certification Recommendations are enabled for the current org
	IaiEnableCertificationRecommendations *bool `json:"iaiEnableCertificationRecommendations,omitempty"`
	SodReportConfigs []ReportConfigDTO `json:"sodReportConfigs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgConfig OrgConfig

// NewOrgConfig instantiates a new OrgConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgConfig() *OrgConfig {
	this := OrgConfig{}
	return &this
}

// NewOrgConfigWithDefaults instantiates a new OrgConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgConfigWithDefaults() *OrgConfig {
	this := OrgConfig{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *OrgConfig) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConfig) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *OrgConfig) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *OrgConfig) SetOrgName(v string) {
	o.OrgName = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *OrgConfig) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConfig) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *OrgConfig) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *OrgConfig) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetLcsChangeHonorsSourceEnableFeature returns the LcsChangeHonorsSourceEnableFeature field value if set, zero value otherwise.
func (o *OrgConfig) GetLcsChangeHonorsSourceEnableFeature() bool {
	if o == nil || IsNil(o.LcsChangeHonorsSourceEnableFeature) {
		var ret bool
		return ret
	}
	return *o.LcsChangeHonorsSourceEnableFeature
}

// GetLcsChangeHonorsSourceEnableFeatureOk returns a tuple with the LcsChangeHonorsSourceEnableFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConfig) GetLcsChangeHonorsSourceEnableFeatureOk() (*bool, bool) {
	if o == nil || IsNil(o.LcsChangeHonorsSourceEnableFeature) {
		return nil, false
	}
	return o.LcsChangeHonorsSourceEnableFeature, true
}

// HasLcsChangeHonorsSourceEnableFeature returns a boolean if a field has been set.
func (o *OrgConfig) HasLcsChangeHonorsSourceEnableFeature() bool {
	if o != nil && !IsNil(o.LcsChangeHonorsSourceEnableFeature) {
		return true
	}

	return false
}

// SetLcsChangeHonorsSourceEnableFeature gets a reference to the given bool and assigns it to the LcsChangeHonorsSourceEnableFeature field.
func (o *OrgConfig) SetLcsChangeHonorsSourceEnableFeature(v bool) {
	o.LcsChangeHonorsSourceEnableFeature = &v
}

// GetArmCustomerId returns the ArmCustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgConfig) GetArmCustomerId() string {
	if o == nil || IsNil(o.ArmCustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.ArmCustomerId.Get()
}

// GetArmCustomerIdOk returns a tuple with the ArmCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgConfig) GetArmCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArmCustomerId.Get(), o.ArmCustomerId.IsSet()
}

// HasArmCustomerId returns a boolean if a field has been set.
func (o *OrgConfig) HasArmCustomerId() bool {
	if o != nil && o.ArmCustomerId.IsSet() {
		return true
	}

	return false
}

// SetArmCustomerId gets a reference to the given NullableString and assigns it to the ArmCustomerId field.
func (o *OrgConfig) SetArmCustomerId(v string) {
	o.ArmCustomerId.Set(&v)
}
// SetArmCustomerIdNil sets the value for ArmCustomerId to be an explicit nil
func (o *OrgConfig) SetArmCustomerIdNil() {
	o.ArmCustomerId.Set(nil)
}

// UnsetArmCustomerId ensures that no value is present for ArmCustomerId, not even an explicit nil
func (o *OrgConfig) UnsetArmCustomerId() {
	o.ArmCustomerId.Unset()
}

// GetArmSapSystemIdMappings returns the ArmSapSystemIdMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgConfig) GetArmSapSystemIdMappings() string {
	if o == nil || IsNil(o.ArmSapSystemIdMappings.Get()) {
		var ret string
		return ret
	}
	return *o.ArmSapSystemIdMappings.Get()
}

// GetArmSapSystemIdMappingsOk returns a tuple with the ArmSapSystemIdMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgConfig) GetArmSapSystemIdMappingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArmSapSystemIdMappings.Get(), o.ArmSapSystemIdMappings.IsSet()
}

// HasArmSapSystemIdMappings returns a boolean if a field has been set.
func (o *OrgConfig) HasArmSapSystemIdMappings() bool {
	if o != nil && o.ArmSapSystemIdMappings.IsSet() {
		return true
	}

	return false
}

// SetArmSapSystemIdMappings gets a reference to the given NullableString and assigns it to the ArmSapSystemIdMappings field.
func (o *OrgConfig) SetArmSapSystemIdMappings(v string) {
	o.ArmSapSystemIdMappings.Set(&v)
}
// SetArmSapSystemIdMappingsNil sets the value for ArmSapSystemIdMappings to be an explicit nil
func (o *OrgConfig) SetArmSapSystemIdMappingsNil() {
	o.ArmSapSystemIdMappings.Set(nil)
}

// UnsetArmSapSystemIdMappings ensures that no value is present for ArmSapSystemIdMappings, not even an explicit nil
func (o *OrgConfig) UnsetArmSapSystemIdMappings() {
	o.ArmSapSystemIdMappings.Unset()
}

// GetArmAuth returns the ArmAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgConfig) GetArmAuth() string {
	if o == nil || IsNil(o.ArmAuth.Get()) {
		var ret string
		return ret
	}
	return *o.ArmAuth.Get()
}

// GetArmAuthOk returns a tuple with the ArmAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgConfig) GetArmAuthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArmAuth.Get(), o.ArmAuth.IsSet()
}

// HasArmAuth returns a boolean if a field has been set.
func (o *OrgConfig) HasArmAuth() bool {
	if o != nil && o.ArmAuth.IsSet() {
		return true
	}

	return false
}

// SetArmAuth gets a reference to the given NullableString and assigns it to the ArmAuth field.
func (o *OrgConfig) SetArmAuth(v string) {
	o.ArmAuth.Set(&v)
}
// SetArmAuthNil sets the value for ArmAuth to be an explicit nil
func (o *OrgConfig) SetArmAuthNil() {
	o.ArmAuth.Set(nil)
}

// UnsetArmAuth ensures that no value is present for ArmAuth, not even an explicit nil
func (o *OrgConfig) UnsetArmAuth() {
	o.ArmAuth.Unset()
}

// GetArmDb returns the ArmDb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgConfig) GetArmDb() string {
	if o == nil || IsNil(o.ArmDb.Get()) {
		var ret string
		return ret
	}
	return *o.ArmDb.Get()
}

// GetArmDbOk returns a tuple with the ArmDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgConfig) GetArmDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArmDb.Get(), o.ArmDb.IsSet()
}

// HasArmDb returns a boolean if a field has been set.
func (o *OrgConfig) HasArmDb() bool {
	if o != nil && o.ArmDb.IsSet() {
		return true
	}

	return false
}

// SetArmDb gets a reference to the given NullableString and assigns it to the ArmDb field.
func (o *OrgConfig) SetArmDb(v string) {
	o.ArmDb.Set(&v)
}
// SetArmDbNil sets the value for ArmDb to be an explicit nil
func (o *OrgConfig) SetArmDbNil() {
	o.ArmDb.Set(nil)
}

// UnsetArmDb ensures that no value is present for ArmDb, not even an explicit nil
func (o *OrgConfig) UnsetArmDb() {
	o.ArmDb.Unset()
}

// GetArmSsoUrl returns the ArmSsoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgConfig) GetArmSsoUrl() string {
	if o == nil || IsNil(o.ArmSsoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ArmSsoUrl.Get()
}

// GetArmSsoUrlOk returns a tuple with the ArmSsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgConfig) GetArmSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArmSsoUrl.Get(), o.ArmSsoUrl.IsSet()
}

// HasArmSsoUrl returns a boolean if a field has been set.
func (o *OrgConfig) HasArmSsoUrl() bool {
	if o != nil && o.ArmSsoUrl.IsSet() {
		return true
	}

	return false
}

// SetArmSsoUrl gets a reference to the given NullableString and assigns it to the ArmSsoUrl field.
func (o *OrgConfig) SetArmSsoUrl(v string) {
	o.ArmSsoUrl.Set(&v)
}
// SetArmSsoUrlNil sets the value for ArmSsoUrl to be an explicit nil
func (o *OrgConfig) SetArmSsoUrlNil() {
	o.ArmSsoUrl.Set(nil)
}

// UnsetArmSsoUrl ensures that no value is present for ArmSsoUrl, not even an explicit nil
func (o *OrgConfig) UnsetArmSsoUrl() {
	o.ArmSsoUrl.Unset()
}

// GetIaiEnableCertificationRecommendations returns the IaiEnableCertificationRecommendations field value if set, zero value otherwise.
func (o *OrgConfig) GetIaiEnableCertificationRecommendations() bool {
	if o == nil || IsNil(o.IaiEnableCertificationRecommendations) {
		var ret bool
		return ret
	}
	return *o.IaiEnableCertificationRecommendations
}

// GetIaiEnableCertificationRecommendationsOk returns a tuple with the IaiEnableCertificationRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConfig) GetIaiEnableCertificationRecommendationsOk() (*bool, bool) {
	if o == nil || IsNil(o.IaiEnableCertificationRecommendations) {
		return nil, false
	}
	return o.IaiEnableCertificationRecommendations, true
}

// HasIaiEnableCertificationRecommendations returns a boolean if a field has been set.
func (o *OrgConfig) HasIaiEnableCertificationRecommendations() bool {
	if o != nil && !IsNil(o.IaiEnableCertificationRecommendations) {
		return true
	}

	return false
}

// SetIaiEnableCertificationRecommendations gets a reference to the given bool and assigns it to the IaiEnableCertificationRecommendations field.
func (o *OrgConfig) SetIaiEnableCertificationRecommendations(v bool) {
	o.IaiEnableCertificationRecommendations = &v
}

// GetSodReportConfigs returns the SodReportConfigs field value if set, zero value otherwise.
func (o *OrgConfig) GetSodReportConfigs() []ReportConfigDTO {
	if o == nil || IsNil(o.SodReportConfigs) {
		var ret []ReportConfigDTO
		return ret
	}
	return o.SodReportConfigs
}

// GetSodReportConfigsOk returns a tuple with the SodReportConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConfig) GetSodReportConfigsOk() ([]ReportConfigDTO, bool) {
	if o == nil || IsNil(o.SodReportConfigs) {
		return nil, false
	}
	return o.SodReportConfigs, true
}

// HasSodReportConfigs returns a boolean if a field has been set.
func (o *OrgConfig) HasSodReportConfigs() bool {
	if o != nil && !IsNil(o.SodReportConfigs) {
		return true
	}

	return false
}

// SetSodReportConfigs gets a reference to the given []ReportConfigDTO and assigns it to the SodReportConfigs field.
func (o *OrgConfig) SetSodReportConfigs(v []ReportConfigDTO) {
	o.SodReportConfigs = v
}

func (o OrgConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.LcsChangeHonorsSourceEnableFeature) {
		toSerialize["lcsChangeHonorsSourceEnableFeature"] = o.LcsChangeHonorsSourceEnableFeature
	}
	if o.ArmCustomerId.IsSet() {
		toSerialize["armCustomerId"] = o.ArmCustomerId.Get()
	}
	if o.ArmSapSystemIdMappings.IsSet() {
		toSerialize["armSapSystemIdMappings"] = o.ArmSapSystemIdMappings.Get()
	}
	if o.ArmAuth.IsSet() {
		toSerialize["armAuth"] = o.ArmAuth.Get()
	}
	if o.ArmDb.IsSet() {
		toSerialize["armDb"] = o.ArmDb.Get()
	}
	if o.ArmSsoUrl.IsSet() {
		toSerialize["armSsoUrl"] = o.ArmSsoUrl.Get()
	}
	if !IsNil(o.IaiEnableCertificationRecommendations) {
		toSerialize["iaiEnableCertificationRecommendations"] = o.IaiEnableCertificationRecommendations
	}
	if !IsNil(o.SodReportConfigs) {
		toSerialize["sodReportConfigs"] = o.SodReportConfigs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgConfig) UnmarshalJSON(data []byte) (err error) {
	varOrgConfig := _OrgConfig{}

	err = json.Unmarshal(data, &varOrgConfig)

	if err != nil {
		return err
	}

	*o = OrgConfig(varOrgConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "lcsChangeHonorsSourceEnableFeature")
		delete(additionalProperties, "armCustomerId")
		delete(additionalProperties, "armSapSystemIdMappings")
		delete(additionalProperties, "armAuth")
		delete(additionalProperties, "armDb")
		delete(additionalProperties, "armSsoUrl")
		delete(additionalProperties, "iaiEnableCertificationRecommendations")
		delete(additionalProperties, "sodReportConfigs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgConfig struct {
	value *OrgConfig
	isSet bool
}

func (v NullableOrgConfig) Get() *OrgConfig {
	return v.value
}

func (v *NullableOrgConfig) Set(val *OrgConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgConfig(val *OrgConfig) *NullableOrgConfig {
	return &NullableOrgConfig{value: val, isSet: true}
}

func (v NullableOrgConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


