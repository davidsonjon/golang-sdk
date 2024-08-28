/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the ConnectorRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorRuleResponse{}

// ConnectorRuleResponse ConnectorRuleResponse
type ConnectorRuleResponse struct {
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
	// the ID of the rule
	Id string `json:"id"`
	// an ISO 8601 UTC timestamp when this rule was created
	Created string `json:"created"`
	// an ISO 8601 UTC timestamp when this rule was last modified
	Modified NullableString `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorRuleResponse ConnectorRuleResponse

// NewConnectorRuleResponse instantiates a new ConnectorRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRuleResponse(name string, type_ string, sourceCode SourceCode, id string, created string) *ConnectorRuleResponse {
	this := ConnectorRuleResponse{}
	this.Name = name
	this.Type = type_
	this.SourceCode = sourceCode
	this.Id = id
	this.Created = created
	return &this
}

// NewConnectorRuleResponseWithDefaults instantiates a new ConnectorRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRuleResponseWithDefaults() *ConnectorRuleResponse {
	this := ConnectorRuleResponse{}
	return &this
}

// GetName returns the Name field value
func (o *ConnectorRuleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorRuleResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectorRuleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectorRuleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectorRuleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *ConnectorRuleResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectorRuleResponse) SetType(v string) {
	o.Type = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ConnectorRuleResponse) GetSignature() ConnectorRuleCreateRequestSignature {
	if o == nil || IsNil(o.Signature) {
		var ret ConnectorRuleCreateRequestSignature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetSignatureOk() (*ConnectorRuleCreateRequestSignature, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ConnectorRuleResponse) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given ConnectorRuleCreateRequestSignature and assigns it to the Signature field.
func (o *ConnectorRuleResponse) SetSignature(v ConnectorRuleCreateRequestSignature) {
	o.Signature = &v
}

// GetSourceCode returns the SourceCode field value
func (o *ConnectorRuleResponse) GetSourceCode() SourceCode {
	if o == nil {
		var ret SourceCode
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetSourceCodeOk() (*SourceCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *ConnectorRuleResponse) SetSourceCode(v SourceCode) {
	o.SourceCode = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorRuleResponse) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorRuleResponse) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ConnectorRuleResponse) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ConnectorRuleResponse) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetId returns the Id field value
func (o *ConnectorRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectorRuleResponse) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *ConnectorRuleResponse) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleResponse) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ConnectorRuleResponse) SetCreated(v string) {
	o.Created = v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorRuleResponse) GetModified() string {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret string
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorRuleResponse) GetModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *ConnectorRuleResponse) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableString and assigns it to the Modified field.
func (o *ConnectorRuleResponse) SetModified(v string) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *ConnectorRuleResponse) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *ConnectorRuleResponse) UnsetModified() {
	o.Modified.Unset()
}

func (o ConnectorRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	toSerialize["sourceCode"] = o.SourceCode
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorRuleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"sourceCode",
		"id",
		"created",
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

	varConnectorRuleResponse := _ConnectorRuleResponse{}

	err = json.Unmarshal(data, &varConnectorRuleResponse)

	if err != nil {
		return err
	}

	*o = ConnectorRuleResponse(varConnectorRuleResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "sourceCode")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorRuleResponse struct {
	value *ConnectorRuleResponse
	isSet bool
}

func (v NullableConnectorRuleResponse) Get() *ConnectorRuleResponse {
	return v.value
}

func (v *NullableConnectorRuleResponse) Set(val *ConnectorRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRuleResponse(val *ConnectorRuleResponse) *NullableConnectorRuleResponse {
	return &NullableConnectorRuleResponse{value: val, isSet: true}
}

func (v NullableConnectorRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


