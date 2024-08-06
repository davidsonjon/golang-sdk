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

// checks if the BaseEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEntitlement{}

// BaseEntitlement struct for BaseEntitlement
type BaseEntitlement struct {
	// Indicates whether the entitlement has permissions.
	HasPermissions *bool `json:"hasPermissions,omitempty"`
	// Entitlement's description.
	Description *string `json:"description,omitempty"`
	// Entitlement attribute's name.
	Attribute *string `json:"attribute,omitempty"`
	// Entitlement's value.
	Value *string `json:"value,omitempty"`
	// Entitlement's schema.
	Schema *string `json:"schema,omitempty"`
	// Indicates whether the entitlement is privileged.
	Privileged *bool `json:"privileged,omitempty"`
	// Entitlement's ID.
	Id *string `json:"id,omitempty"`
	// Entitlement's name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseEntitlement BaseEntitlement

// NewBaseEntitlement instantiates a new BaseEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntitlement() *BaseEntitlement {
	this := BaseEntitlement{}
	var hasPermissions bool = false
	this.HasPermissions = &hasPermissions
	var privileged bool = false
	this.Privileged = &privileged
	return &this
}

// NewBaseEntitlementWithDefaults instantiates a new BaseEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntitlementWithDefaults() *BaseEntitlement {
	this := BaseEntitlement{}
	var hasPermissions bool = false
	this.HasPermissions = &hasPermissions
	var privileged bool = false
	this.Privileged = &privileged
	return &this
}

// GetHasPermissions returns the HasPermissions field value if set, zero value otherwise.
func (o *BaseEntitlement) GetHasPermissions() bool {
	if o == nil || isNil(o.HasPermissions) {
		var ret bool
		return ret
	}
	return *o.HasPermissions
}

// GetHasPermissionsOk returns a tuple with the HasPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetHasPermissionsOk() (*bool, bool) {
	if o == nil || isNil(o.HasPermissions) {
		return nil, false
	}
	return o.HasPermissions, true
}

// HasHasPermissions returns a boolean if a field has been set.
func (o *BaseEntitlement) HasHasPermissions() bool {
	if o != nil && !isNil(o.HasPermissions) {
		return true
	}

	return false
}

// SetHasPermissions gets a reference to the given bool and assigns it to the HasPermissions field.
func (o *BaseEntitlement) SetHasPermissions(v bool) {
	o.HasPermissions = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseEntitlement) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseEntitlement) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseEntitlement) SetDescription(v string) {
	o.Description = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *BaseEntitlement) GetAttribute() string {
	if o == nil || isNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *BaseEntitlement) HasAttribute() bool {
	if o != nil && !isNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *BaseEntitlement) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BaseEntitlement) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BaseEntitlement) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BaseEntitlement) SetValue(v string) {
	o.Value = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *BaseEntitlement) GetSchema() string {
	if o == nil || isNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetSchemaOk() (*string, bool) {
	if o == nil || isNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *BaseEntitlement) HasSchema() bool {
	if o != nil && !isNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *BaseEntitlement) SetSchema(v string) {
	o.Schema = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *BaseEntitlement) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *BaseEntitlement) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *BaseEntitlement) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseEntitlement) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseEntitlement) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseEntitlement) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseEntitlement) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntitlement) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseEntitlement) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseEntitlement) SetName(v string) {
	o.Name = &v
}

func (o BaseEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasPermissions) {
		toSerialize["hasPermissions"] = o.HasPermissions
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varBaseEntitlement := _BaseEntitlement{}

	if err = json.Unmarshal(bytes, &varBaseEntitlement); err == nil {
	*o = BaseEntitlement(varBaseEntitlement)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "hasPermissions")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseEntitlement struct {
	value *BaseEntitlement
	isSet bool
}

func (v NullableBaseEntitlement) Get() *BaseEntitlement {
	return v.value
}

func (v *NullableBaseEntitlement) Set(val *BaseEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntitlement(val *BaseEntitlement) *NullableBaseEntitlement {
	return &NullableBaseEntitlement{value: val, isSet: true}
}

func (v NullableBaseEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


