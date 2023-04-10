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

// checks if the ServiceDeskIntegrationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDeskIntegrationDto{}

// ServiceDeskIntegrationDto struct for ServiceDeskIntegrationDto
type ServiceDeskIntegrationDto struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// Description of the Service Desk integration
	Description string `json:"description"`
	// Service Desk integration types  - ServiceNowSDIM - ServiceNow 
	Type string `json:"type"`
	// Reference to the identity that is the owner of this Service Desk integration
	OwnerRef *BaseReferenceDto `json:"ownerRef,omitempty"`
	// Reference to the source cluster for this Service Desk integration
	ClusterRef *BaseReferenceDto `json:"clusterRef,omitempty"`
	// ID of the cluster for the Service Desk integration (replaced by clusterRef, retained for backward compatibility)
	// Deprecated
	Cluster *string `json:"cluster,omitempty"`
	// Source IDs for the Service Desk integration (replaced by provisioningConfig.managedSResourceRefs, but retained here for backward compatibility)
	// Deprecated
	ManagedSources []string `json:"managedSources,omitempty"`
	ProvisioningConfig *ProvisioningConfig `json:"provisioningConfig,omitempty"`
	// Attributes of the Service Desk integration.  Validation constraints enforced by the implementation.
	Attributes map[string]interface{} `json:"attributes"`
	// Reference to beforeProvisioningRule for this Service Desk integration
	BeforeProvisioningRule *BaseReferenceDto `json:"beforeProvisioningRule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeskIntegrationDto ServiceDeskIntegrationDto

// NewServiceDeskIntegrationDto instantiates a new ServiceDeskIntegrationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeskIntegrationDto(name string, description string, type_ string, attributes map[string]interface{}) *ServiceDeskIntegrationDto {
	this := ServiceDeskIntegrationDto{}
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewServiceDeskIntegrationDtoWithDefaults instantiates a new ServiceDeskIntegrationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeskIntegrationDtoWithDefaults() *ServiceDeskIntegrationDto {
	this := ServiceDeskIntegrationDto{}
	var type_ string = "ServiceNowSDIM"
	this.Type = type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceDeskIntegrationDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ServiceDeskIntegrationDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceDeskIntegrationDto) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ServiceDeskIntegrationDto) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ServiceDeskIntegrationDto) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value
func (o *ServiceDeskIntegrationDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceDeskIntegrationDto) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ServiceDeskIntegrationDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceDeskIntegrationDto) SetType(v string) {
	o.Type = v
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetOwnerRef() BaseReferenceDto {
	if o == nil || isNil(o.OwnerRef) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.OwnerRef
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetOwnerRefOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.OwnerRef) {
		return nil, false
	}
	return o.OwnerRef, true
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasOwnerRef() bool {
	if o != nil && !isNil(o.OwnerRef) {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given BaseReferenceDto and assigns it to the OwnerRef field.
func (o *ServiceDeskIntegrationDto) SetOwnerRef(v BaseReferenceDto) {
	o.OwnerRef = &v
}

// GetClusterRef returns the ClusterRef field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetClusterRef() BaseReferenceDto {
	if o == nil || isNil(o.ClusterRef) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.ClusterRef
}

// GetClusterRefOk returns a tuple with the ClusterRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetClusterRefOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.ClusterRef) {
		return nil, false
	}
	return o.ClusterRef, true
}

// HasClusterRef returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasClusterRef() bool {
	if o != nil && !isNil(o.ClusterRef) {
		return true
	}

	return false
}

// SetClusterRef gets a reference to the given BaseReferenceDto and assigns it to the ClusterRef field.
func (o *ServiceDeskIntegrationDto) SetClusterRef(v BaseReferenceDto) {
	o.ClusterRef = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
// Deprecated
func (o *ServiceDeskIntegrationDto) GetCluster() string {
	if o == nil || isNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceDeskIntegrationDto) GetClusterOk() (*string, bool) {
	if o == nil || isNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasCluster() bool {
	if o != nil && !isNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
// Deprecated
func (o *ServiceDeskIntegrationDto) SetCluster(v string) {
	o.Cluster = &v
}

// GetManagedSources returns the ManagedSources field value if set, zero value otherwise.
// Deprecated
func (o *ServiceDeskIntegrationDto) GetManagedSources() []string {
	if o == nil || isNil(o.ManagedSources) {
		var ret []string
		return ret
	}
	return o.ManagedSources
}

// GetManagedSourcesOk returns a tuple with the ManagedSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceDeskIntegrationDto) GetManagedSourcesOk() ([]string, bool) {
	if o == nil || isNil(o.ManagedSources) {
		return nil, false
	}
	return o.ManagedSources, true
}

// HasManagedSources returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasManagedSources() bool {
	if o != nil && !isNil(o.ManagedSources) {
		return true
	}

	return false
}

// SetManagedSources gets a reference to the given []string and assigns it to the ManagedSources field.
// Deprecated
func (o *ServiceDeskIntegrationDto) SetManagedSources(v []string) {
	o.ManagedSources = v
}

// GetProvisioningConfig returns the ProvisioningConfig field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetProvisioningConfig() ProvisioningConfig {
	if o == nil || isNil(o.ProvisioningConfig) {
		var ret ProvisioningConfig
		return ret
	}
	return *o.ProvisioningConfig
}

// GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetProvisioningConfigOk() (*ProvisioningConfig, bool) {
	if o == nil || isNil(o.ProvisioningConfig) {
		return nil, false
	}
	return o.ProvisioningConfig, true
}

// HasProvisioningConfig returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasProvisioningConfig() bool {
	if o != nil && !isNil(o.ProvisioningConfig) {
		return true
	}

	return false
}

// SetProvisioningConfig gets a reference to the given ProvisioningConfig and assigns it to the ProvisioningConfig field.
func (o *ServiceDeskIntegrationDto) SetProvisioningConfig(v ProvisioningConfig) {
	o.ProvisioningConfig = &v
}

// GetAttributes returns the Attributes field value
func (o *ServiceDeskIntegrationDto) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ServiceDeskIntegrationDto) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetBeforeProvisioningRule returns the BeforeProvisioningRule field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDto) GetBeforeProvisioningRule() BaseReferenceDto {
	if o == nil || isNil(o.BeforeProvisioningRule) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.BeforeProvisioningRule
}

// GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDto) GetBeforeProvisioningRuleOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.BeforeProvisioningRule) {
		return nil, false
	}
	return o.BeforeProvisioningRule, true
}

// HasBeforeProvisioningRule returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDto) HasBeforeProvisioningRule() bool {
	if o != nil && !isNil(o.BeforeProvisioningRule) {
		return true
	}

	return false
}

// SetBeforeProvisioningRule gets a reference to the given BaseReferenceDto and assigns it to the BeforeProvisioningRule field.
func (o *ServiceDeskIntegrationDto) SetBeforeProvisioningRule(v BaseReferenceDto) {
	o.BeforeProvisioningRule = &v
}

func (o ServiceDeskIntegrationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDeskIntegrationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: created is readOnly
	// skip: modified is readOnly
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

func (o *ServiceDeskIntegrationDto) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDeskIntegrationDto := _ServiceDeskIntegrationDto{}

	if err = json.Unmarshal(bytes, &varServiceDeskIntegrationDto); err == nil {
		*o = ServiceDeskIntegrationDto(varServiceDeskIntegrationDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
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

type NullableServiceDeskIntegrationDto struct {
	value *ServiceDeskIntegrationDto
	isSet bool
}

func (v NullableServiceDeskIntegrationDto) Get() *ServiceDeskIntegrationDto {
	return v.value
}

func (v *NullableServiceDeskIntegrationDto) Set(val *ServiceDeskIntegrationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeskIntegrationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeskIntegrationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeskIntegrationDto(val *ServiceDeskIntegrationDto) *NullableServiceDeskIntegrationDto {
	return &NullableServiceDeskIntegrationDto{value: val, isSet: true}
}

func (v NullableServiceDeskIntegrationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeskIntegrationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


