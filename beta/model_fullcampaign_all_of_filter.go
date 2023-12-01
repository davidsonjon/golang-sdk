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

// checks if the FullcampaignAllOfFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullcampaignAllOfFilter{}

// FullcampaignAllOfFilter Determines which items will be included in this campaign. The default campaign filter is used if this field is left blank.
type FullcampaignAllOfFilter struct {
	// The ID of whatever type of filter is being used.
	Id *string `json:"id,omitempty"`
	// Type of the filter
	Type *string `json:"type,omitempty"`
	// Name of the filter
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullcampaignAllOfFilter FullcampaignAllOfFilter

// NewFullcampaignAllOfFilter instantiates a new FullcampaignAllOfFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullcampaignAllOfFilter() *FullcampaignAllOfFilter {
	this := FullcampaignAllOfFilter{}
	return &this
}

// NewFullcampaignAllOfFilterWithDefaults instantiates a new FullcampaignAllOfFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullcampaignAllOfFilterWithDefaults() *FullcampaignAllOfFilter {
	this := FullcampaignAllOfFilter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullcampaignAllOfFilter) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfFilter) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullcampaignAllOfFilter) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullcampaignAllOfFilter) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FullcampaignAllOfFilter) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfFilter) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FullcampaignAllOfFilter) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FullcampaignAllOfFilter) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullcampaignAllOfFilter) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfFilter) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullcampaignAllOfFilter) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullcampaignAllOfFilter) SetName(v string) {
	o.Name = &v
}

func (o FullcampaignAllOfFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullcampaignAllOfFilter) ToMap() (map[string]interface{}, error) {
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

func (o *FullcampaignAllOfFilter) UnmarshalJSON(bytes []byte) (err error) {
	varFullcampaignAllOfFilter := _FullcampaignAllOfFilter{}

	if err = json.Unmarshal(bytes, &varFullcampaignAllOfFilter); err == nil {
	*o = FullcampaignAllOfFilter(varFullcampaignAllOfFilter)
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

type NullableFullcampaignAllOfFilter struct {
	value *FullcampaignAllOfFilter
	isSet bool
}

func (v NullableFullcampaignAllOfFilter) Get() *FullcampaignAllOfFilter {
	return v.value
}

func (v *NullableFullcampaignAllOfFilter) Set(val *FullcampaignAllOfFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableFullcampaignAllOfFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFullcampaignAllOfFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullcampaignAllOfFilter(val *FullcampaignAllOfFilter) *NullableFullcampaignAllOfFilter {
	return &NullableFullcampaignAllOfFilter{value: val, isSet: true}
}

func (v NullableFullcampaignAllOfFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullcampaignAllOfFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


