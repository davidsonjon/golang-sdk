/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the ServiceDeskIntegrationTemplateType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDeskIntegrationTemplateType{}

// ServiceDeskIntegrationTemplateType This represents a Service Desk Integration template type.
type ServiceDeskIntegrationTemplateType struct {
	// This is the name of the type.
	Name *string `json:"name,omitempty"`
	// This is the type value for the type.
	Type string `json:"type"`
	// This is the scriptName attribute value for the type.
	ScriptName string `json:"scriptName"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeskIntegrationTemplateType ServiceDeskIntegrationTemplateType

// NewServiceDeskIntegrationTemplateType instantiates a new ServiceDeskIntegrationTemplateType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeskIntegrationTemplateType(type_ string, scriptName string) *ServiceDeskIntegrationTemplateType {
	this := ServiceDeskIntegrationTemplateType{}
	this.Type = type_
	this.ScriptName = scriptName
	return &this
}

// NewServiceDeskIntegrationTemplateTypeWithDefaults instantiates a new ServiceDeskIntegrationTemplateType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeskIntegrationTemplateTypeWithDefaults() *ServiceDeskIntegrationTemplateType {
	this := ServiceDeskIntegrationTemplateType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationTemplateType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationTemplateType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceDeskIntegrationTemplateType) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *ServiceDeskIntegrationTemplateType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceDeskIntegrationTemplateType) SetType(v string) {
	o.Type = v
}

// GetScriptName returns the ScriptName field value
func (o *ServiceDeskIntegrationTemplateType) GetScriptName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateType) GetScriptNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptName, true
}

// SetScriptName sets field value
func (o *ServiceDeskIntegrationTemplateType) SetScriptName(v string) {
	o.ScriptName = v
}

func (o ServiceDeskIntegrationTemplateType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDeskIntegrationTemplateType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	toSerialize["scriptName"] = o.ScriptName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceDeskIntegrationTemplateType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"scriptName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceDeskIntegrationTemplateType := _ServiceDeskIntegrationTemplateType{}

	err = json.Unmarshal(data, &varServiceDeskIntegrationTemplateType)

	if err != nil {
		return err
	}

	*o = ServiceDeskIntegrationTemplateType(varServiceDeskIntegrationTemplateType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "scriptName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceDeskIntegrationTemplateType struct {
	value *ServiceDeskIntegrationTemplateType
	isSet bool
}

func (v NullableServiceDeskIntegrationTemplateType) Get() *ServiceDeskIntegrationTemplateType {
	return v.value
}

func (v *NullableServiceDeskIntegrationTemplateType) Set(val *ServiceDeskIntegrationTemplateType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeskIntegrationTemplateType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeskIntegrationTemplateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeskIntegrationTemplateType(val *ServiceDeskIntegrationTemplateType) *NullableServiceDeskIntegrationTemplateType {
	return &NullableServiceDeskIntegrationTemplateType{value: val, isSet: true}
}

func (v NullableServiceDeskIntegrationTemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeskIntegrationTemplateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


