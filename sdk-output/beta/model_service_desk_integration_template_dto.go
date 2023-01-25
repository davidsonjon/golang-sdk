/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"time"
	"encoding/json"
)

// ServiceDeskIntegrationTemplateDto struct for ServiceDeskIntegrationTemplateDto
type ServiceDeskIntegrationTemplateDto struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// The 'type' property specifies the type of the Service Desk integration template.
	Type string `json:"type"`
	// The 'attributes' property value is a map of attributes available for integrations using this Service Desk integration template.
	Attributes map[string]interface{} `json:"attributes"`
	ProvisioningConfig ProvisioningConfig `json:"provisioningConfig"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeskIntegrationTemplateDto ServiceDeskIntegrationTemplateDto

// NewServiceDeskIntegrationTemplateDto instantiates a new ServiceDeskIntegrationTemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeskIntegrationTemplateDto(name string, type_ string, attributes map[string]interface{}, provisioningConfig ProvisioningConfig) *ServiceDeskIntegrationTemplateDto {
	this := ServiceDeskIntegrationTemplateDto{}
	this.Name = name
	this.Type = type_
	this.Attributes = attributes
	this.ProvisioningConfig = provisioningConfig
	return &this
}

// NewServiceDeskIntegrationTemplateDtoWithDefaults instantiates a new ServiceDeskIntegrationTemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeskIntegrationTemplateDtoWithDefaults() *ServiceDeskIntegrationTemplateDto {
	this := ServiceDeskIntegrationTemplateDto{}
	var type_ string = "Web Service SDIM"
	this.Type = type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationTemplateDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationTemplateDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceDeskIntegrationTemplateDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ServiceDeskIntegrationTemplateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceDeskIntegrationTemplateDto) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationTemplateDto) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationTemplateDto) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ServiceDeskIntegrationTemplateDto) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationTemplateDto) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationTemplateDto) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ServiceDeskIntegrationTemplateDto) SetModified(v time.Time) {
	o.Modified = &v
}

// GetType returns the Type field value
func (o *ServiceDeskIntegrationTemplateDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceDeskIntegrationTemplateDto) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ServiceDeskIntegrationTemplateDto) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ServiceDeskIntegrationTemplateDto) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetProvisioningConfig returns the ProvisioningConfig field value
func (o *ServiceDeskIntegrationTemplateDto) GetProvisioningConfig() ProvisioningConfig {
	if o == nil {
		var ret ProvisioningConfig
		return ret
	}

	return o.ProvisioningConfig
}

// GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDto) GetProvisioningConfigOk() (*ProvisioningConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningConfig, true
}

// SetProvisioningConfig sets field value
func (o *ServiceDeskIntegrationTemplateDto) SetProvisioningConfig(v ProvisioningConfig) {
	o.ProvisioningConfig = v
}

func (o ServiceDeskIntegrationTemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["provisioningConfig"] = o.ProvisioningConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceDeskIntegrationTemplateDto) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDeskIntegrationTemplateDto := _ServiceDeskIntegrationTemplateDto{}

	if err = json.Unmarshal(bytes, &varServiceDeskIntegrationTemplateDto); err == nil {
		*o = ServiceDeskIntegrationTemplateDto(varServiceDeskIntegrationTemplateDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "provisioningConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceDeskIntegrationTemplateDto struct {
	value *ServiceDeskIntegrationTemplateDto
	isSet bool
}

func (v NullableServiceDeskIntegrationTemplateDto) Get() *ServiceDeskIntegrationTemplateDto {
	return v.value
}

func (v *NullableServiceDeskIntegrationTemplateDto) Set(val *ServiceDeskIntegrationTemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeskIntegrationTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeskIntegrationTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeskIntegrationTemplateDto(val *ServiceDeskIntegrationTemplateDto) *NullableServiceDeskIntegrationTemplateDto {
	return &NullableServiceDeskIntegrationTemplateDto{value: val, isSet: true}
}

func (v NullableServiceDeskIntegrationTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeskIntegrationTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


