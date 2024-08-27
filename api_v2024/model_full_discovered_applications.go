/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the FullDiscoveredApplications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullDiscoveredApplications{}

// FullDiscoveredApplications Discovered applications with their respective associated sources
type FullDiscoveredApplications struct {
	// Unique identifier for the discovered application.
	Id *string `json:"id,omitempty"`
	// Name of the discovered application.
	Name *string `json:"name,omitempty"`
	// Source from which the application was discovered.
	DiscoverySource *string `json:"discoverySource,omitempty"`
	// The vendor associated with the discovered application.
	DiscoveredVendor *string `json:"discoveredVendor,omitempty"`
	// A brief description of the discovered application.
	Description *string `json:"description,omitempty"`
	// List of recommended connectors for the application.
	RecommendedConnectors []string `json:"recommendedConnectors,omitempty"`
	// The timestamp when the application was last received via an entitlement aggregation invocation  or a manual csv upload, in ISO 8601 format.
	DiscoveredAt *time.Time `json:"discoveredAt,omitempty"`
	// The timestamp when the application was first discovered, in ISO 8601 format.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The status of an application within the discovery source.  By default this field is set to \"ACTIVE\" when the application is discovered.  If an application has been deleted from within the discovery source, the status will be set to \"INACTIVE\".
	Status *string `json:"status,omitempty"`
	// List of associated sources related to this discovered application.
	AssociatedSources []string `json:"associatedSources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullDiscoveredApplications FullDiscoveredApplications

// NewFullDiscoveredApplications instantiates a new FullDiscoveredApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDiscoveredApplications() *FullDiscoveredApplications {
	this := FullDiscoveredApplications{}
	return &this
}

// NewFullDiscoveredApplicationsWithDefaults instantiates a new FullDiscoveredApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDiscoveredApplicationsWithDefaults() *FullDiscoveredApplications {
	this := FullDiscoveredApplications{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullDiscoveredApplications) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullDiscoveredApplications) SetName(v string) {
	o.Name = &v
}

// GetDiscoverySource returns the DiscoverySource field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetDiscoverySource() string {
	if o == nil || isNil(o.DiscoverySource) {
		var ret string
		return ret
	}
	return *o.DiscoverySource
}

// GetDiscoverySourceOk returns a tuple with the DiscoverySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetDiscoverySourceOk() (*string, bool) {
	if o == nil || isNil(o.DiscoverySource) {
		return nil, false
	}
	return o.DiscoverySource, true
}

// HasDiscoverySource returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasDiscoverySource() bool {
	if o != nil && !isNil(o.DiscoverySource) {
		return true
	}

	return false
}

// SetDiscoverySource gets a reference to the given string and assigns it to the DiscoverySource field.
func (o *FullDiscoveredApplications) SetDiscoverySource(v string) {
	o.DiscoverySource = &v
}

// GetDiscoveredVendor returns the DiscoveredVendor field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetDiscoveredVendor() string {
	if o == nil || isNil(o.DiscoveredVendor) {
		var ret string
		return ret
	}
	return *o.DiscoveredVendor
}

// GetDiscoveredVendorOk returns a tuple with the DiscoveredVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetDiscoveredVendorOk() (*string, bool) {
	if o == nil || isNil(o.DiscoveredVendor) {
		return nil, false
	}
	return o.DiscoveredVendor, true
}

// HasDiscoveredVendor returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasDiscoveredVendor() bool {
	if o != nil && !isNil(o.DiscoveredVendor) {
		return true
	}

	return false
}

// SetDiscoveredVendor gets a reference to the given string and assigns it to the DiscoveredVendor field.
func (o *FullDiscoveredApplications) SetDiscoveredVendor(v string) {
	o.DiscoveredVendor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullDiscoveredApplications) SetDescription(v string) {
	o.Description = &v
}

// GetRecommendedConnectors returns the RecommendedConnectors field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetRecommendedConnectors() []string {
	if o == nil || isNil(o.RecommendedConnectors) {
		var ret []string
		return ret
	}
	return o.RecommendedConnectors
}

// GetRecommendedConnectorsOk returns a tuple with the RecommendedConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetRecommendedConnectorsOk() ([]string, bool) {
	if o == nil || isNil(o.RecommendedConnectors) {
		return nil, false
	}
	return o.RecommendedConnectors, true
}

// HasRecommendedConnectors returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasRecommendedConnectors() bool {
	if o != nil && !isNil(o.RecommendedConnectors) {
		return true
	}

	return false
}

// SetRecommendedConnectors gets a reference to the given []string and assigns it to the RecommendedConnectors field.
func (o *FullDiscoveredApplications) SetRecommendedConnectors(v []string) {
	o.RecommendedConnectors = v
}

// GetDiscoveredAt returns the DiscoveredAt field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetDiscoveredAt() time.Time {
	if o == nil || isNil(o.DiscoveredAt) {
		var ret time.Time
		return ret
	}
	return *o.DiscoveredAt
}

// GetDiscoveredAtOk returns a tuple with the DiscoveredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetDiscoveredAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DiscoveredAt) {
		return nil, false
	}
	return o.DiscoveredAt, true
}

// HasDiscoveredAt returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasDiscoveredAt() bool {
	if o != nil && !isNil(o.DiscoveredAt) {
		return true
	}

	return false
}

// SetDiscoveredAt gets a reference to the given time.Time and assigns it to the DiscoveredAt field.
func (o *FullDiscoveredApplications) SetDiscoveredAt(v time.Time) {
	o.DiscoveredAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FullDiscoveredApplications) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FullDiscoveredApplications) SetStatus(v string) {
	o.Status = &v
}

// GetAssociatedSources returns the AssociatedSources field value if set, zero value otherwise.
func (o *FullDiscoveredApplications) GetAssociatedSources() []string {
	if o == nil || isNil(o.AssociatedSources) {
		var ret []string
		return ret
	}
	return o.AssociatedSources
}

// GetAssociatedSourcesOk returns a tuple with the AssociatedSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDiscoveredApplications) GetAssociatedSourcesOk() ([]string, bool) {
	if o == nil || isNil(o.AssociatedSources) {
		return nil, false
	}
	return o.AssociatedSources, true
}

// HasAssociatedSources returns a boolean if a field has been set.
func (o *FullDiscoveredApplications) HasAssociatedSources() bool {
	if o != nil && !isNil(o.AssociatedSources) {
		return true
	}

	return false
}

// SetAssociatedSources gets a reference to the given []string and assigns it to the AssociatedSources field.
func (o *FullDiscoveredApplications) SetAssociatedSources(v []string) {
	o.AssociatedSources = v
}

func (o FullDiscoveredApplications) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullDiscoveredApplications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DiscoverySource) {
		toSerialize["discoverySource"] = o.DiscoverySource
	}
	if !isNil(o.DiscoveredVendor) {
		toSerialize["discoveredVendor"] = o.DiscoveredVendor
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.RecommendedConnectors) {
		toSerialize["recommendedConnectors"] = o.RecommendedConnectors
	}
	if !isNil(o.DiscoveredAt) {
		toSerialize["discoveredAt"] = o.DiscoveredAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.AssociatedSources) {
		toSerialize["associatedSources"] = o.AssociatedSources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullDiscoveredApplications) UnmarshalJSON(bytes []byte) (err error) {
	varFullDiscoveredApplications := _FullDiscoveredApplications{}

	if err = json.Unmarshal(bytes, &varFullDiscoveredApplications); err == nil {
	*o = FullDiscoveredApplications(varFullDiscoveredApplications)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "discoverySource")
		delete(additionalProperties, "discoveredVendor")
		delete(additionalProperties, "description")
		delete(additionalProperties, "recommendedConnectors")
		delete(additionalProperties, "discoveredAt")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "associatedSources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullDiscoveredApplications struct {
	value *FullDiscoveredApplications
	isSet bool
}

func (v NullableFullDiscoveredApplications) Get() *FullDiscoveredApplications {
	return v.value
}

func (v *NullableFullDiscoveredApplications) Set(val *FullDiscoveredApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDiscoveredApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDiscoveredApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDiscoveredApplications(val *FullDiscoveredApplications) *NullableFullDiscoveredApplications {
	return &NullableFullDiscoveredApplications{value: val, isSet: true}
}

func (v NullableFullDiscoveredApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDiscoveredApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

