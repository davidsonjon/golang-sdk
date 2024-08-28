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

// checks if the ManagedClusterQueue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedClusterQueue{}

// ManagedClusterQueue Managed Cluster key pair for Cluster
type ManagedClusterQueue struct {
	// ManagedCluster queue name
	Name *string `json:"name,omitempty"`
	// ManagedCluster queue aws region
	Region *string `json:"region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClusterQueue ManagedClusterQueue

// NewManagedClusterQueue instantiates a new ManagedClusterQueue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClusterQueue() *ManagedClusterQueue {
	this := ManagedClusterQueue{}
	return &this
}

// NewManagedClusterQueueWithDefaults instantiates a new ManagedClusterQueue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClusterQueueWithDefaults() *ManagedClusterQueue {
	this := ManagedClusterQueue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagedClusterQueue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedClusterQueue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagedClusterQueue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagedClusterQueue) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ManagedClusterQueue) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedClusterQueue) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ManagedClusterQueue) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ManagedClusterQueue) SetRegion(v string) {
	o.Region = &v
}

func (o ManagedClusterQueue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedClusterQueue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedClusterQueue) UnmarshalJSON(data []byte) (err error) {
	varManagedClusterQueue := _ManagedClusterQueue{}

	err = json.Unmarshal(data, &varManagedClusterQueue)

	if err != nil {
		return err
	}

	*o = ManagedClusterQueue(varManagedClusterQueue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedClusterQueue struct {
	value *ManagedClusterQueue
	isSet bool
}

func (v NullableManagedClusterQueue) Get() *ManagedClusterQueue {
	return v.value
}

func (v *NullableManagedClusterQueue) Set(val *ManagedClusterQueue) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClusterQueue) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClusterQueue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClusterQueue(val *ManagedClusterQueue) *NullableManagedClusterQueue {
	return &NullableManagedClusterQueue{value: val, isSet: true}
}

func (v NullableManagedClusterQueue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClusterQueue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


