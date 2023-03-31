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

// checks if the ConnectorRuleCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorRuleCreateRequest{}

// ConnectorRuleCreateRequest ConnectorRuleCreateRequest
type ConnectorRuleCreateRequest struct {
	// the name of the rule
	Name string `json:"name"`
	// a description of the rule's purpose
	Description *string `json:"description,omitempty"`
	// the type of rule
	Type string `json:"type"`
	Signature *ConnectorRuleCreateRequestSignature `json:"signature,omitempty"`
	SourceCode SourceCode `json:"sourceCode"`
	// a map of string to objects
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorRuleCreateRequest ConnectorRuleCreateRequest

// NewConnectorRuleCreateRequest instantiates a new ConnectorRuleCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRuleCreateRequest(name string, type_ string, sourceCode SourceCode) *ConnectorRuleCreateRequest {
	this := ConnectorRuleCreateRequest{}
	this.Name = name
	this.Type = type_
	this.SourceCode = sourceCode
	return &this
}

// NewConnectorRuleCreateRequestWithDefaults instantiates a new ConnectorRuleCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRuleCreateRequestWithDefaults() *ConnectorRuleCreateRequest {
	this := ConnectorRuleCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ConnectorRuleCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorRuleCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectorRuleCreateRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRuleCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectorRuleCreateRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectorRuleCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *ConnectorRuleCreateRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleCreateRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectorRuleCreateRequest) SetType(v string) {
	o.Type = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ConnectorRuleCreateRequest) GetSignature() ConnectorRuleCreateRequestSignature {
	if o == nil || isNil(o.Signature) {
		var ret ConnectorRuleCreateRequestSignature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRuleCreateRequest) GetSignatureOk() (*ConnectorRuleCreateRequestSignature, bool) {
	if o == nil || isNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ConnectorRuleCreateRequest) HasSignature() bool {
	if o != nil && !isNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given ConnectorRuleCreateRequestSignature and assigns it to the Signature field.
func (o *ConnectorRuleCreateRequest) SetSignature(v ConnectorRuleCreateRequestSignature) {
	o.Signature = &v
}

// GetSourceCode returns the SourceCode field value
func (o *ConnectorRuleCreateRequest) GetSourceCode() SourceCode {
	if o == nil {
		var ret SourceCode
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleCreateRequest) GetSourceCodeOk() (*SourceCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *ConnectorRuleCreateRequest) SetSourceCode(v SourceCode) {
	o.SourceCode = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ConnectorRuleCreateRequest) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRuleCreateRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ConnectorRuleCreateRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ConnectorRuleCreateRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o ConnectorRuleCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorRuleCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	if !isNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	toSerialize["sourceCode"] = o.SourceCode
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorRuleCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorRuleCreateRequest := _ConnectorRuleCreateRequest{}

	if err = json.Unmarshal(bytes, &varConnectorRuleCreateRequest); err == nil {
		*o = ConnectorRuleCreateRequest(varConnectorRuleCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "sourceCode")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorRuleCreateRequest struct {
	value *ConnectorRuleCreateRequest
	isSet bool
}

func (v NullableConnectorRuleCreateRequest) Get() *ConnectorRuleCreateRequest {
	return v.value
}

func (v *NullableConnectorRuleCreateRequest) Set(val *ConnectorRuleCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRuleCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRuleCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRuleCreateRequest(val *ConnectorRuleCreateRequest) *NullableConnectorRuleCreateRequest {
	return &NullableConnectorRuleCreateRequest{value: val, isSet: true}
}

func (v NullableConnectorRuleCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRuleCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


