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

// checks if the RoleInsightsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleInsightsEntitlement{}

// RoleInsightsEntitlement struct for RoleInsightsEntitlement
type RoleInsightsEntitlement struct {
	// Name of the entitlement
	Name *string `json:"name,omitempty"`
	// Id of the entitlement
	Id *string `json:"id,omitempty"`
	// Description for the entitlement
	Description *string `json:"description,omitempty"`
	// Source or the application for the entitlement
	Source *string `json:"source,omitempty"`
	// Attribute for the entitlement
	Attribute *string `json:"attribute,omitempty"`
	// Attribute value for the entitlement
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleInsightsEntitlement RoleInsightsEntitlement

// NewRoleInsightsEntitlement instantiates a new RoleInsightsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleInsightsEntitlement() *RoleInsightsEntitlement {
	this := RoleInsightsEntitlement{}
	return &this
}

// NewRoleInsightsEntitlementWithDefaults instantiates a new RoleInsightsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleInsightsEntitlementWithDefaults() *RoleInsightsEntitlement {
	this := RoleInsightsEntitlement{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleInsightsEntitlement) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlement) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleInsightsEntitlement) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleInsightsEntitlement) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleInsightsEntitlement) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlement) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleInsightsEntitlement) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleInsightsEntitlement) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleInsightsEntitlement) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlement) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleInsightsEntitlement) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleInsightsEntitlement) SetDescription(v string) {
	o.Description = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RoleInsightsEntitlement) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlement) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RoleInsightsEntitlement) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RoleInsightsEntitlement) SetSource(v string) {
	o.Source = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *RoleInsightsEntitlement) GetAttribute() string {
	if o == nil || isNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlement) GetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *RoleInsightsEntitlement) HasAttribute() bool {
	if o != nil && !isNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *RoleInsightsEntitlement) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RoleInsightsEntitlement) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlement) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RoleInsightsEntitlement) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RoleInsightsEntitlement) SetValue(v string) {
	o.Value = &v
}

func (o RoleInsightsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleInsightsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleInsightsEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varRoleInsightsEntitlement := _RoleInsightsEntitlement{}

	if err = json.Unmarshal(bytes, &varRoleInsightsEntitlement); err == nil {
	*o = RoleInsightsEntitlement(varRoleInsightsEntitlement)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "source")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleInsightsEntitlement struct {
	value *RoleInsightsEntitlement
	isSet bool
}

func (v NullableRoleInsightsEntitlement) Get() *RoleInsightsEntitlement {
	return v.value
}

func (v *NullableRoleInsightsEntitlement) Set(val *RoleInsightsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleInsightsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleInsightsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleInsightsEntitlement(val *RoleInsightsEntitlement) *NullableRoleInsightsEntitlement {
	return &NullableRoleInsightsEntitlement{value: val, isSet: true}
}

func (v NullableRoleInsightsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleInsightsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


