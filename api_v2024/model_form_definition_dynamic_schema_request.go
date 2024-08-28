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

// checks if the FormDefinitionDynamicSchemaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDefinitionDynamicSchemaRequest{}

// FormDefinitionDynamicSchemaRequest struct for FormDefinitionDynamicSchemaRequest
type FormDefinitionDynamicSchemaRequest struct {
	Attributes *FormDefinitionDynamicSchemaRequestAttributes `json:"attributes,omitempty"`
	// Description is the form definition dynamic schema description text
	Description *string `json:"description,omitempty"`
	// ID is a unique identifier
	Id *string `json:"id,omitempty"`
	// Type is the form definition dynamic schema type
	Type *string `json:"type,omitempty"`
	// VersionNumber is the form definition dynamic schema version number
	VersionNumber *int64 `json:"versionNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormDefinitionDynamicSchemaRequest FormDefinitionDynamicSchemaRequest

// NewFormDefinitionDynamicSchemaRequest instantiates a new FormDefinitionDynamicSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDefinitionDynamicSchemaRequest() *FormDefinitionDynamicSchemaRequest {
	this := FormDefinitionDynamicSchemaRequest{}
	return &this
}

// NewFormDefinitionDynamicSchemaRequestWithDefaults instantiates a new FormDefinitionDynamicSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDefinitionDynamicSchemaRequestWithDefaults() *FormDefinitionDynamicSchemaRequest {
	this := FormDefinitionDynamicSchemaRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FormDefinitionDynamicSchemaRequest) GetAttributes() FormDefinitionDynamicSchemaRequestAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret FormDefinitionDynamicSchemaRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionDynamicSchemaRequest) GetAttributesOk() (*FormDefinitionDynamicSchemaRequestAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FormDefinitionDynamicSchemaRequest) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FormDefinitionDynamicSchemaRequestAttributes and assigns it to the Attributes field.
func (o *FormDefinitionDynamicSchemaRequest) SetAttributes(v FormDefinitionDynamicSchemaRequestAttributes) {
	o.Attributes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FormDefinitionDynamicSchemaRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionDynamicSchemaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FormDefinitionDynamicSchemaRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FormDefinitionDynamicSchemaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormDefinitionDynamicSchemaRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionDynamicSchemaRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormDefinitionDynamicSchemaRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormDefinitionDynamicSchemaRequest) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FormDefinitionDynamicSchemaRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionDynamicSchemaRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FormDefinitionDynamicSchemaRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FormDefinitionDynamicSchemaRequest) SetType(v string) {
	o.Type = &v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *FormDefinitionDynamicSchemaRequest) GetVersionNumber() int64 {
	if o == nil || IsNil(o.VersionNumber) {
		var ret int64
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionDynamicSchemaRequest) GetVersionNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.VersionNumber) {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *FormDefinitionDynamicSchemaRequest) HasVersionNumber() bool {
	if o != nil && !IsNil(o.VersionNumber) {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given int64 and assigns it to the VersionNumber field.
func (o *FormDefinitionDynamicSchemaRequest) SetVersionNumber(v int64) {
	o.VersionNumber = &v
}

func (o FormDefinitionDynamicSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDefinitionDynamicSchemaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.VersionNumber) {
		toSerialize["versionNumber"] = o.VersionNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormDefinitionDynamicSchemaRequest) UnmarshalJSON(data []byte) (err error) {
	varFormDefinitionDynamicSchemaRequest := _FormDefinitionDynamicSchemaRequest{}

	err = json.Unmarshal(data, &varFormDefinitionDynamicSchemaRequest)

	if err != nil {
		return err
	}

	*o = FormDefinitionDynamicSchemaRequest(varFormDefinitionDynamicSchemaRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "versionNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormDefinitionDynamicSchemaRequest struct {
	value *FormDefinitionDynamicSchemaRequest
	isSet bool
}

func (v NullableFormDefinitionDynamicSchemaRequest) Get() *FormDefinitionDynamicSchemaRequest {
	return v.value
}

func (v *NullableFormDefinitionDynamicSchemaRequest) Set(val *FormDefinitionDynamicSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDefinitionDynamicSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDefinitionDynamicSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDefinitionDynamicSchemaRequest(val *FormDefinitionDynamicSchemaRequest) *NullableFormDefinitionDynamicSchemaRequest {
	return &NullableFormDefinitionDynamicSchemaRequest{value: val, isSet: true}
}

func (v NullableFormDefinitionDynamicSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDefinitionDynamicSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


