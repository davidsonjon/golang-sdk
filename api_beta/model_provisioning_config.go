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

// checks if the ProvisioningConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConfig{}

// ProvisioningConfig Specification of a Service Desk integration provisioning configuration.
type ProvisioningConfig struct {
	// Specifies whether this configuration is used to manage provisioning requests for all sources from the org.  If true, no managedResourceRefs are allowed.
	UniversalManager *bool `json:"universalManager,omitempty"`
	// References to sources for the Service Desk integration template.  May only be specified if universalManager is false.
	ManagedResourceRefs []ProvisioningConfigManagedResourceRefsInner `json:"managedResourceRefs,omitempty"`
	PlanInitializerScript *ProvisioningConfigPlanInitializerScript `json:"planInitializerScript,omitempty"`
	// Name of an attribute that when true disables the saving of ProvisioningRequest objects whenever plans are sent through this integration.
	NoProvisioningRequests *bool `json:"noProvisioningRequests,omitempty"`
	// When saving pending requests is enabled, this defines the number of hours the request is allowed to live before it is considered expired and no longer affects plan compilation.
	ProvisioningRequestExpiration *int32 `json:"provisioningRequestExpiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConfig ProvisioningConfig

// NewProvisioningConfig instantiates a new ProvisioningConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConfig() *ProvisioningConfig {
	this := ProvisioningConfig{}
	var noProvisioningRequests bool = false
	this.NoProvisioningRequests = &noProvisioningRequests
	return &this
}

// NewProvisioningConfigWithDefaults instantiates a new ProvisioningConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConfigWithDefaults() *ProvisioningConfig {
	this := ProvisioningConfig{}
	var noProvisioningRequests bool = false
	this.NoProvisioningRequests = &noProvisioningRequests
	return &this
}

// GetUniversalManager returns the UniversalManager field value if set, zero value otherwise.
func (o *ProvisioningConfig) GetUniversalManager() bool {
	if o == nil || isNil(o.UniversalManager) {
		var ret bool
		return ret
	}
	return *o.UniversalManager
}

// GetUniversalManagerOk returns a tuple with the UniversalManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfig) GetUniversalManagerOk() (*bool, bool) {
	if o == nil || isNil(o.UniversalManager) {
		return nil, false
	}
	return o.UniversalManager, true
}

// HasUniversalManager returns a boolean if a field has been set.
func (o *ProvisioningConfig) HasUniversalManager() bool {
	if o != nil && !isNil(o.UniversalManager) {
		return true
	}

	return false
}

// SetUniversalManager gets a reference to the given bool and assigns it to the UniversalManager field.
func (o *ProvisioningConfig) SetUniversalManager(v bool) {
	o.UniversalManager = &v
}

// GetManagedResourceRefs returns the ManagedResourceRefs field value if set, zero value otherwise.
func (o *ProvisioningConfig) GetManagedResourceRefs() []ProvisioningConfigManagedResourceRefsInner {
	if o == nil || isNil(o.ManagedResourceRefs) {
		var ret []ProvisioningConfigManagedResourceRefsInner
		return ret
	}
	return o.ManagedResourceRefs
}

// GetManagedResourceRefsOk returns a tuple with the ManagedResourceRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfig) GetManagedResourceRefsOk() ([]ProvisioningConfigManagedResourceRefsInner, bool) {
	if o == nil || isNil(o.ManagedResourceRefs) {
		return nil, false
	}
	return o.ManagedResourceRefs, true
}

// HasManagedResourceRefs returns a boolean if a field has been set.
func (o *ProvisioningConfig) HasManagedResourceRefs() bool {
	if o != nil && !isNil(o.ManagedResourceRefs) {
		return true
	}

	return false
}

// SetManagedResourceRefs gets a reference to the given []ProvisioningConfigManagedResourceRefsInner and assigns it to the ManagedResourceRefs field.
func (o *ProvisioningConfig) SetManagedResourceRefs(v []ProvisioningConfigManagedResourceRefsInner) {
	o.ManagedResourceRefs = v
}

// GetPlanInitializerScript returns the PlanInitializerScript field value if set, zero value otherwise.
func (o *ProvisioningConfig) GetPlanInitializerScript() ProvisioningConfigPlanInitializerScript {
	if o == nil || isNil(o.PlanInitializerScript) {
		var ret ProvisioningConfigPlanInitializerScript
		return ret
	}
	return *o.PlanInitializerScript
}

// GetPlanInitializerScriptOk returns a tuple with the PlanInitializerScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfig) GetPlanInitializerScriptOk() (*ProvisioningConfigPlanInitializerScript, bool) {
	if o == nil || isNil(o.PlanInitializerScript) {
		return nil, false
	}
	return o.PlanInitializerScript, true
}

// HasPlanInitializerScript returns a boolean if a field has been set.
func (o *ProvisioningConfig) HasPlanInitializerScript() bool {
	if o != nil && !isNil(o.PlanInitializerScript) {
		return true
	}

	return false
}

// SetPlanInitializerScript gets a reference to the given ProvisioningConfigPlanInitializerScript and assigns it to the PlanInitializerScript field.
func (o *ProvisioningConfig) SetPlanInitializerScript(v ProvisioningConfigPlanInitializerScript) {
	o.PlanInitializerScript = &v
}

// GetNoProvisioningRequests returns the NoProvisioningRequests field value if set, zero value otherwise.
func (o *ProvisioningConfig) GetNoProvisioningRequests() bool {
	if o == nil || isNil(o.NoProvisioningRequests) {
		var ret bool
		return ret
	}
	return *o.NoProvisioningRequests
}

// GetNoProvisioningRequestsOk returns a tuple with the NoProvisioningRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfig) GetNoProvisioningRequestsOk() (*bool, bool) {
	if o == nil || isNil(o.NoProvisioningRequests) {
		return nil, false
	}
	return o.NoProvisioningRequests, true
}

// HasNoProvisioningRequests returns a boolean if a field has been set.
func (o *ProvisioningConfig) HasNoProvisioningRequests() bool {
	if o != nil && !isNil(o.NoProvisioningRequests) {
		return true
	}

	return false
}

// SetNoProvisioningRequests gets a reference to the given bool and assigns it to the NoProvisioningRequests field.
func (o *ProvisioningConfig) SetNoProvisioningRequests(v bool) {
	o.NoProvisioningRequests = &v
}

// GetProvisioningRequestExpiration returns the ProvisioningRequestExpiration field value if set, zero value otherwise.
func (o *ProvisioningConfig) GetProvisioningRequestExpiration() int32 {
	if o == nil || isNil(o.ProvisioningRequestExpiration) {
		var ret int32
		return ret
	}
	return *o.ProvisioningRequestExpiration
}

// GetProvisioningRequestExpirationOk returns a tuple with the ProvisioningRequestExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfig) GetProvisioningRequestExpirationOk() (*int32, bool) {
	if o == nil || isNil(o.ProvisioningRequestExpiration) {
		return nil, false
	}
	return o.ProvisioningRequestExpiration, true
}

// HasProvisioningRequestExpiration returns a boolean if a field has been set.
func (o *ProvisioningConfig) HasProvisioningRequestExpiration() bool {
	if o != nil && !isNil(o.ProvisioningRequestExpiration) {
		return true
	}

	return false
}

// SetProvisioningRequestExpiration gets a reference to the given int32 and assigns it to the ProvisioningRequestExpiration field.
func (o *ProvisioningConfig) SetProvisioningRequestExpiration(v int32) {
	o.ProvisioningRequestExpiration = &v
}

func (o ProvisioningConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: universalManager is readOnly
	if !isNil(o.ManagedResourceRefs) {
		toSerialize["managedResourceRefs"] = o.ManagedResourceRefs
	}
	if !isNil(o.PlanInitializerScript) {
		toSerialize["planInitializerScript"] = o.PlanInitializerScript
	}
	if !isNil(o.NoProvisioningRequests) {
		toSerialize["noProvisioningRequests"] = o.NoProvisioningRequests
	}
	if !isNil(o.ProvisioningRequestExpiration) {
		toSerialize["provisioningRequestExpiration"] = o.ProvisioningRequestExpiration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConfig) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConfig := _ProvisioningConfig{}

	if err = json.Unmarshal(bytes, &varProvisioningConfig); err == nil {
	*o = ProvisioningConfig(varProvisioningConfig)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "universalManager")
		delete(additionalProperties, "managedResourceRefs")
		delete(additionalProperties, "planInitializerScript")
		delete(additionalProperties, "noProvisioningRequests")
		delete(additionalProperties, "provisioningRequestExpiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConfig struct {
	value *ProvisioningConfig
	isSet bool
}

func (v NullableProvisioningConfig) Get() *ProvisioningConfig {
	return v.value
}

func (v *NullableProvisioningConfig) Set(val *ProvisioningConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConfig(val *ProvisioningConfig) *NullableProvisioningConfig {
	return &NullableProvisioningConfig{value: val, isSet: true}
}

func (v NullableProvisioningConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


