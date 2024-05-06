/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the DiscoveredApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveredApplicationsInner{}

// DiscoveredApplicationsInner struct for DiscoveredApplicationsInner
type DiscoveredApplicationsInner struct {
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
	// The timestamp when the application was discovered, in ISO 8601 format.
	DiscoveredTimestamp *time.Time `json:"discoveredTimestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveredApplicationsInner DiscoveredApplicationsInner

// NewDiscoveredApplicationsInner instantiates a new DiscoveredApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredApplicationsInner() *DiscoveredApplicationsInner {
	this := DiscoveredApplicationsInner{}
	return &this
}

// NewDiscoveredApplicationsInnerWithDefaults instantiates a new DiscoveredApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredApplicationsInnerWithDefaults() *DiscoveredApplicationsInner {
	this := DiscoveredApplicationsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiscoveredApplicationsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiscoveredApplicationsInner) SetName(v string) {
	o.Name = &v
}

// GetDiscoverySource returns the DiscoverySource field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetDiscoverySource() string {
	if o == nil || isNil(o.DiscoverySource) {
		var ret string
		return ret
	}
	return *o.DiscoverySource
}

// GetDiscoverySourceOk returns a tuple with the DiscoverySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetDiscoverySourceOk() (*string, bool) {
	if o == nil || isNil(o.DiscoverySource) {
		return nil, false
	}
	return o.DiscoverySource, true
}

// HasDiscoverySource returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasDiscoverySource() bool {
	if o != nil && !isNil(o.DiscoverySource) {
		return true
	}

	return false
}

// SetDiscoverySource gets a reference to the given string and assigns it to the DiscoverySource field.
func (o *DiscoveredApplicationsInner) SetDiscoverySource(v string) {
	o.DiscoverySource = &v
}

// GetDiscoveredVendor returns the DiscoveredVendor field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetDiscoveredVendor() string {
	if o == nil || isNil(o.DiscoveredVendor) {
		var ret string
		return ret
	}
	return *o.DiscoveredVendor
}

// GetDiscoveredVendorOk returns a tuple with the DiscoveredVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetDiscoveredVendorOk() (*string, bool) {
	if o == nil || isNil(o.DiscoveredVendor) {
		return nil, false
	}
	return o.DiscoveredVendor, true
}

// HasDiscoveredVendor returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasDiscoveredVendor() bool {
	if o != nil && !isNil(o.DiscoveredVendor) {
		return true
	}

	return false
}

// SetDiscoveredVendor gets a reference to the given string and assigns it to the DiscoveredVendor field.
func (o *DiscoveredApplicationsInner) SetDiscoveredVendor(v string) {
	o.DiscoveredVendor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DiscoveredApplicationsInner) SetDescription(v string) {
	o.Description = &v
}

// GetRecommendedConnectors returns the RecommendedConnectors field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetRecommendedConnectors() []string {
	if o == nil || isNil(o.RecommendedConnectors) {
		var ret []string
		return ret
	}
	return o.RecommendedConnectors
}

// GetRecommendedConnectorsOk returns a tuple with the RecommendedConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetRecommendedConnectorsOk() ([]string, bool) {
	if o == nil || isNil(o.RecommendedConnectors) {
		return nil, false
	}
	return o.RecommendedConnectors, true
}

// HasRecommendedConnectors returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasRecommendedConnectors() bool {
	if o != nil && !isNil(o.RecommendedConnectors) {
		return true
	}

	return false
}

// SetRecommendedConnectors gets a reference to the given []string and assigns it to the RecommendedConnectors field.
func (o *DiscoveredApplicationsInner) SetRecommendedConnectors(v []string) {
	o.RecommendedConnectors = v
}

// GetDiscoveredTimestamp returns the DiscoveredTimestamp field value if set, zero value otherwise.
func (o *DiscoveredApplicationsInner) GetDiscoveredTimestamp() time.Time {
	if o == nil || isNil(o.DiscoveredTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.DiscoveredTimestamp
}

// GetDiscoveredTimestampOk returns a tuple with the DiscoveredTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredApplicationsInner) GetDiscoveredTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.DiscoveredTimestamp) {
		return nil, false
	}
	return o.DiscoveredTimestamp, true
}

// HasDiscoveredTimestamp returns a boolean if a field has been set.
func (o *DiscoveredApplicationsInner) HasDiscoveredTimestamp() bool {
	if o != nil && !isNil(o.DiscoveredTimestamp) {
		return true
	}

	return false
}

// SetDiscoveredTimestamp gets a reference to the given time.Time and assigns it to the DiscoveredTimestamp field.
func (o *DiscoveredApplicationsInner) SetDiscoveredTimestamp(v time.Time) {
	o.DiscoveredTimestamp = &v
}

func (o DiscoveredApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveredApplicationsInner) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.DiscoveredTimestamp) {
		toSerialize["discoveredTimestamp"] = o.DiscoveredTimestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveredApplicationsInner) UnmarshalJSON(bytes []byte) (err error) {
	varDiscoveredApplicationsInner := _DiscoveredApplicationsInner{}

	if err = json.Unmarshal(bytes, &varDiscoveredApplicationsInner); err == nil {
	*o = DiscoveredApplicationsInner(varDiscoveredApplicationsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "discoverySource")
		delete(additionalProperties, "discoveredVendor")
		delete(additionalProperties, "description")
		delete(additionalProperties, "recommendedConnectors")
		delete(additionalProperties, "discoveredTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveredApplicationsInner struct {
	value *DiscoveredApplicationsInner
	isSet bool
}

func (v NullableDiscoveredApplicationsInner) Get() *DiscoveredApplicationsInner {
	return v.value
}

func (v *NullableDiscoveredApplicationsInner) Set(val *DiscoveredApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredApplicationsInner(val *DiscoveredApplicationsInner) *NullableDiscoveredApplicationsInner {
	return &NullableDiscoveredApplicationsInner{value: val, isSet: true}
}

func (v NullableDiscoveredApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


