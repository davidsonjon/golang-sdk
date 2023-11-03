/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the ServiceDeskIntegrationDtoAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDeskIntegrationDtoAllOf{}

// ServiceDeskIntegrationDtoAllOf Specification of a Service Desk integration.
type ServiceDeskIntegrationDtoAllOf struct {
	// Description of the Service Desk integration.
	Description string `json:"description"`
	// Service Desk integration types  - ServiceNowSDIM - ServiceNow 
	Type string `json:"type"`
	OwnerRef *OwnerDto `json:"ownerRef,omitempty"`
	ClusterRef *SourceClusterDto `json:"clusterRef,omitempty"`
	// ID of the cluster for the Service Desk integration (replaced by clusterRef, retained for backward compatibility).
	// Deprecated
	Cluster *string `json:"cluster,omitempty"`
	// Source IDs for the Service Desk integration (replaced by provisioningConfig.managedSResourceRefs, but retained here for backward compatibility).
	// Deprecated
	ManagedSources []string `json:"managedSources,omitempty"`
	ProvisioningConfig *ProvisioningConfig `json:"provisioningConfig,omitempty"`
	// Attributes of the Service Desk integration.  Validation constraints enforced by the implementation.
	Attributes map[string]interface{} `json:"attributes"`
	BeforeProvisioningRule *BeforeProvisioningRuleDto `json:"beforeProvisioningRule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeskIntegrationDtoAllOf ServiceDeskIntegrationDtoAllOf

// NewServiceDeskIntegrationDtoAllOf instantiates a new ServiceDeskIntegrationDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeskIntegrationDtoAllOf(description string, type_ string, attributes map[string]interface{}) *ServiceDeskIntegrationDtoAllOf {
	this := ServiceDeskIntegrationDtoAllOf{}
	this.Description = description
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewServiceDeskIntegrationDtoAllOfWithDefaults instantiates a new ServiceDeskIntegrationDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeskIntegrationDtoAllOfWithDefaults() *ServiceDeskIntegrationDtoAllOf {
	this := ServiceDeskIntegrationDtoAllOf{}
	var type_ string = "ServiceNowSDIM"
	this.Type = type_
	return &this
}

// GetDescription returns the Description field value
func (o *ServiceDeskIntegrationDtoAllOf) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceDeskIntegrationDtoAllOf) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ServiceDeskIntegrationDtoAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceDeskIntegrationDtoAllOf) SetType(v string) {
	o.Type = v
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOf) GetOwnerRef() OwnerDto {
	if o == nil || isNil(o.OwnerRef) {
		var ret OwnerDto
		return ret
	}
	return *o.OwnerRef
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetOwnerRefOk() (*OwnerDto, bool) {
	if o == nil || isNil(o.OwnerRef) {
		return nil, false
	}
	return o.OwnerRef, true
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOf) HasOwnerRef() bool {
	if o != nil && !isNil(o.OwnerRef) {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given OwnerDto and assigns it to the OwnerRef field.
func (o *ServiceDeskIntegrationDtoAllOf) SetOwnerRef(v OwnerDto) {
	o.OwnerRef = &v
}

// GetClusterRef returns the ClusterRef field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOf) GetClusterRef() SourceClusterDto {
	if o == nil || isNil(o.ClusterRef) {
		var ret SourceClusterDto
		return ret
	}
	return *o.ClusterRef
}

// GetClusterRefOk returns a tuple with the ClusterRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetClusterRefOk() (*SourceClusterDto, bool) {
	if o == nil || isNil(o.ClusterRef) {
		return nil, false
	}
	return o.ClusterRef, true
}

// HasClusterRef returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOf) HasClusterRef() bool {
	if o != nil && !isNil(o.ClusterRef) {
		return true
	}

	return false
}

// SetClusterRef gets a reference to the given SourceClusterDto and assigns it to the ClusterRef field.
func (o *ServiceDeskIntegrationDtoAllOf) SetClusterRef(v SourceClusterDto) {
	o.ClusterRef = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
// Deprecated
func (o *ServiceDeskIntegrationDtoAllOf) GetCluster() string {
	if o == nil || isNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceDeskIntegrationDtoAllOf) GetClusterOk() (*string, bool) {
	if o == nil || isNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOf) HasCluster() bool {
	if o != nil && !isNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
// Deprecated
func (o *ServiceDeskIntegrationDtoAllOf) SetCluster(v string) {
	o.Cluster = &v
}

// GetManagedSources returns the ManagedSources field value if set, zero value otherwise.
// Deprecated
func (o *ServiceDeskIntegrationDtoAllOf) GetManagedSources() []string {
	if o == nil || isNil(o.ManagedSources) {
		var ret []string
		return ret
	}
	return o.ManagedSources
}

// GetManagedSourcesOk returns a tuple with the ManagedSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceDeskIntegrationDtoAllOf) GetManagedSourcesOk() ([]string, bool) {
	if o == nil || isNil(o.ManagedSources) {
		return nil, false
	}
	return o.ManagedSources, true
}

// HasManagedSources returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOf) HasManagedSources() bool {
	if o != nil && !isNil(o.ManagedSources) {
		return true
	}

	return false
}

// SetManagedSources gets a reference to the given []string and assigns it to the ManagedSources field.
// Deprecated
func (o *ServiceDeskIntegrationDtoAllOf) SetManagedSources(v []string) {
	o.ManagedSources = v
}

// GetProvisioningConfig returns the ProvisioningConfig field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOf) GetProvisioningConfig() ProvisioningConfig {
	if o == nil || isNil(o.ProvisioningConfig) {
		var ret ProvisioningConfig
		return ret
	}
	return *o.ProvisioningConfig
}

// GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetProvisioningConfigOk() (*ProvisioningConfig, bool) {
	if o == nil || isNil(o.ProvisioningConfig) {
		return nil, false
	}
	return o.ProvisioningConfig, true
}

// HasProvisioningConfig returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOf) HasProvisioningConfig() bool {
	if o != nil && !isNil(o.ProvisioningConfig) {
		return true
	}

	return false
}

// SetProvisioningConfig gets a reference to the given ProvisioningConfig and assigns it to the ProvisioningConfig field.
func (o *ServiceDeskIntegrationDtoAllOf) SetProvisioningConfig(v ProvisioningConfig) {
	o.ProvisioningConfig = &v
}

// GetAttributes returns the Attributes field value
func (o *ServiceDeskIntegrationDtoAllOf) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ServiceDeskIntegrationDtoAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetBeforeProvisioningRule returns the BeforeProvisioningRule field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOf) GetBeforeProvisioningRule() BeforeProvisioningRuleDto {
	if o == nil || isNil(o.BeforeProvisioningRule) {
		var ret BeforeProvisioningRuleDto
		return ret
	}
	return *o.BeforeProvisioningRule
}

// GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOf) GetBeforeProvisioningRuleOk() (*BeforeProvisioningRuleDto, bool) {
	if o == nil || isNil(o.BeforeProvisioningRule) {
		return nil, false
	}
	return o.BeforeProvisioningRule, true
}

// HasBeforeProvisioningRule returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOf) HasBeforeProvisioningRule() bool {
	if o != nil && !isNil(o.BeforeProvisioningRule) {
		return true
	}

	return false
}

// SetBeforeProvisioningRule gets a reference to the given BeforeProvisioningRuleDto and assigns it to the BeforeProvisioningRule field.
func (o *ServiceDeskIntegrationDtoAllOf) SetBeforeProvisioningRule(v BeforeProvisioningRuleDto) {
	o.BeforeProvisioningRule = &v
}

func (o ServiceDeskIntegrationDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDeskIntegrationDtoAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	if !isNil(o.OwnerRef) {
		toSerialize["ownerRef"] = o.OwnerRef
	}
	if !isNil(o.ClusterRef) {
		toSerialize["clusterRef"] = o.ClusterRef
	}
	if !isNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !isNil(o.ManagedSources) {
		toSerialize["managedSources"] = o.ManagedSources
	}
	if !isNil(o.ProvisioningConfig) {
		toSerialize["provisioningConfig"] = o.ProvisioningConfig
	}
	toSerialize["attributes"] = o.Attributes
	if !isNil(o.BeforeProvisioningRule) {
		toSerialize["beforeProvisioningRule"] = o.BeforeProvisioningRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceDeskIntegrationDtoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDeskIntegrationDtoAllOf := _ServiceDeskIntegrationDtoAllOf{}

	if err = json.Unmarshal(bytes, &varServiceDeskIntegrationDtoAllOf); err == nil {
		*o = ServiceDeskIntegrationDtoAllOf(varServiceDeskIntegrationDtoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ownerRef")
		delete(additionalProperties, "clusterRef")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "managedSources")
		delete(additionalProperties, "provisioningConfig")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "beforeProvisioningRule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceDeskIntegrationDtoAllOf struct {
	value *ServiceDeskIntegrationDtoAllOf
	isSet bool
}

func (v NullableServiceDeskIntegrationDtoAllOf) Get() *ServiceDeskIntegrationDtoAllOf {
	return v.value
}

func (v *NullableServiceDeskIntegrationDtoAllOf) Set(val *ServiceDeskIntegrationDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeskIntegrationDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeskIntegrationDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeskIntegrationDtoAllOf(val *ServiceDeskIntegrationDtoAllOf) *NullableServiceDeskIntegrationDtoAllOf {
	return &NullableServiceDeskIntegrationDtoAllOf{value: val, isSet: true}
}

func (v NullableServiceDeskIntegrationDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeskIntegrationDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


