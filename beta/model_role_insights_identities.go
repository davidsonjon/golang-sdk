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

// checks if the RoleInsightsIdentities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleInsightsIdentities{}

// RoleInsightsIdentities struct for RoleInsightsIdentities
type RoleInsightsIdentities struct {
	// Id for identity
	Id *string `json:"id,omitempty"`
	// Name for identity
	Name *string `json:"name,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleInsightsIdentities RoleInsightsIdentities

// NewRoleInsightsIdentities instantiates a new RoleInsightsIdentities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleInsightsIdentities() *RoleInsightsIdentities {
	this := RoleInsightsIdentities{}
	return &this
}

// NewRoleInsightsIdentitiesWithDefaults instantiates a new RoleInsightsIdentities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleInsightsIdentitiesWithDefaults() *RoleInsightsIdentities {
	this := RoleInsightsIdentities{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleInsightsIdentities) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsIdentities) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleInsightsIdentities) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleInsightsIdentities) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleInsightsIdentities) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsIdentities) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleInsightsIdentities) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleInsightsIdentities) SetName(v string) {
	o.Name = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RoleInsightsIdentities) GetAttributes() map[string]string {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsIdentities) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RoleInsightsIdentities) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *RoleInsightsIdentities) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o RoleInsightsIdentities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleInsightsIdentities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleInsightsIdentities) UnmarshalJSON(bytes []byte) (err error) {
	varRoleInsightsIdentities := _RoleInsightsIdentities{}

	if err = json.Unmarshal(bytes, &varRoleInsightsIdentities); err == nil {
	*o = RoleInsightsIdentities(varRoleInsightsIdentities)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleInsightsIdentities struct {
	value *RoleInsightsIdentities
	isSet bool
}

func (v NullableRoleInsightsIdentities) Get() *RoleInsightsIdentities {
	return v.value
}

func (v *NullableRoleInsightsIdentities) Set(val *RoleInsightsIdentities) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleInsightsIdentities) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleInsightsIdentities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleInsightsIdentities(val *RoleInsightsIdentities) *NullableRoleInsightsIdentities {
	return &NullableRoleInsightsIdentities{value: val, isSet: true}
}

func (v NullableRoleInsightsIdentities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleInsightsIdentities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


