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

// checks if the FullcampaignAllOfSourcesWithOrphanEntitlements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullcampaignAllOfSourcesWithOrphanEntitlements{}

// FullcampaignAllOfSourcesWithOrphanEntitlements struct for FullcampaignAllOfSourcesWithOrphanEntitlements
type FullcampaignAllOfSourcesWithOrphanEntitlements struct {
	// Id of the source
	Id *string `json:"id,omitempty"`
	// Type
	Type *string `json:"type,omitempty"`
	// Name of the source
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullcampaignAllOfSourcesWithOrphanEntitlements FullcampaignAllOfSourcesWithOrphanEntitlements

// NewFullcampaignAllOfSourcesWithOrphanEntitlements instantiates a new FullcampaignAllOfSourcesWithOrphanEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullcampaignAllOfSourcesWithOrphanEntitlements() *FullcampaignAllOfSourcesWithOrphanEntitlements {
	this := FullcampaignAllOfSourcesWithOrphanEntitlements{}
	return &this
}

// NewFullcampaignAllOfSourcesWithOrphanEntitlementsWithDefaults instantiates a new FullcampaignAllOfSourcesWithOrphanEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullcampaignAllOfSourcesWithOrphanEntitlementsWithDefaults() *FullcampaignAllOfSourcesWithOrphanEntitlements {
	this := FullcampaignAllOfSourcesWithOrphanEntitlements{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) SetName(v string) {
	o.Name = &v
}

func (o FullcampaignAllOfSourcesWithOrphanEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullcampaignAllOfSourcesWithOrphanEntitlements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullcampaignAllOfSourcesWithOrphanEntitlements) UnmarshalJSON(bytes []byte) (err error) {
	varFullcampaignAllOfSourcesWithOrphanEntitlements := _FullcampaignAllOfSourcesWithOrphanEntitlements{}

	if err = json.Unmarshal(bytes, &varFullcampaignAllOfSourcesWithOrphanEntitlements); err == nil {
	*o = FullcampaignAllOfSourcesWithOrphanEntitlements(varFullcampaignAllOfSourcesWithOrphanEntitlements)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullcampaignAllOfSourcesWithOrphanEntitlements struct {
	value *FullcampaignAllOfSourcesWithOrphanEntitlements
	isSet bool
}

func (v NullableFullcampaignAllOfSourcesWithOrphanEntitlements) Get() *FullcampaignAllOfSourcesWithOrphanEntitlements {
	return v.value
}

func (v *NullableFullcampaignAllOfSourcesWithOrphanEntitlements) Set(val *FullcampaignAllOfSourcesWithOrphanEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableFullcampaignAllOfSourcesWithOrphanEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableFullcampaignAllOfSourcesWithOrphanEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullcampaignAllOfSourcesWithOrphanEntitlements(val *FullcampaignAllOfSourcesWithOrphanEntitlements) *NullableFullcampaignAllOfSourcesWithOrphanEntitlements {
	return &NullableFullcampaignAllOfSourcesWithOrphanEntitlements{value: val, isSet: true}
}

func (v NullableFullcampaignAllOfSourcesWithOrphanEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullcampaignAllOfSourcesWithOrphanEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


