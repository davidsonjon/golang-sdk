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

// checks if the RoleMiningEntitlementRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningEntitlementRef{}

// RoleMiningEntitlementRef struct for RoleMiningEntitlementRef
type RoleMiningEntitlementRef struct {
	// Id of the entitlement
	Id *string `json:"id,omitempty"`
	// Name of the entitlement
	Name *string `json:"name,omitempty"`
	// Description forthe entitlement
	Description NullableString `json:"description,omitempty"`
	// The entitlement attribute
	Attribute *string `json:"attribute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningEntitlementRef RoleMiningEntitlementRef

// NewRoleMiningEntitlementRef instantiates a new RoleMiningEntitlementRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningEntitlementRef() *RoleMiningEntitlementRef {
	this := RoleMiningEntitlementRef{}
	return &this
}

// NewRoleMiningEntitlementRefWithDefaults instantiates a new RoleMiningEntitlementRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningEntitlementRefWithDefaults() *RoleMiningEntitlementRef {
	this := RoleMiningEntitlementRef{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMiningEntitlementRef) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlementRef) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMiningEntitlementRef) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMiningEntitlementRef) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleMiningEntitlementRef) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlementRef) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningEntitlementRef) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleMiningEntitlementRef) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningEntitlementRef) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningEntitlementRef) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleMiningEntitlementRef) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RoleMiningEntitlementRef) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RoleMiningEntitlementRef) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RoleMiningEntitlementRef) UnsetDescription() {
	o.Description.Unset()
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *RoleMiningEntitlementRef) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningEntitlementRef) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *RoleMiningEntitlementRef) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *RoleMiningEntitlementRef) SetAttribute(v string) {
	o.Attribute = &v
}

func (o RoleMiningEntitlementRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningEntitlementRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningEntitlementRef) UnmarshalJSON(data []byte) (err error) {
	varRoleMiningEntitlementRef := _RoleMiningEntitlementRef{}

	err = json.Unmarshal(data, &varRoleMiningEntitlementRef)

	if err != nil {
		return err
	}

	*o = RoleMiningEntitlementRef(varRoleMiningEntitlementRef)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningEntitlementRef struct {
	value *RoleMiningEntitlementRef
	isSet bool
}

func (v NullableRoleMiningEntitlementRef) Get() *RoleMiningEntitlementRef {
	return v.value
}

func (v *NullableRoleMiningEntitlementRef) Set(val *RoleMiningEntitlementRef) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningEntitlementRef) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningEntitlementRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningEntitlementRef(val *RoleMiningEntitlementRef) *NullableRoleMiningEntitlementRef {
	return &NullableRoleMiningEntitlementRef{value: val, isSet: true}
}

func (v NullableRoleMiningEntitlementRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningEntitlementRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


