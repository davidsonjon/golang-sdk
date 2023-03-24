/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// CreateConnectorRequest struct for CreateConnectorRequest
type CreateConnectorRequest struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ClassName *string `json:"className,omitempty"`
	DirectConnect *bool `json:"directConnect,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateConnectorRequest CreateConnectorRequest

// NewCreateConnectorRequest instantiates a new CreateConnectorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectorRequest() *CreateConnectorRequest {
	this := CreateConnectorRequest{}
	return &this
}

// NewCreateConnectorRequestWithDefaults instantiates a new CreateConnectorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectorRequestWithDefaults() *CreateConnectorRequest {
	this := CreateConnectorRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateConnectorRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateConnectorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetClassName returns the ClassName field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetClassName() string {
	if o == nil || isNil(o.ClassName) {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetClassNameOk() (*string, bool) {
	if o == nil || isNil(o.ClassName) {
		return nil, false
	}
	return o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasClassName() bool {
	if o != nil && !isNil(o.ClassName) {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *CreateConnectorRequest) SetClassName(v string) {
	o.ClassName = &v
}

// GetDirectConnect returns the DirectConnect field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetDirectConnect() bool {
	if o == nil || isNil(o.DirectConnect) {
		var ret bool
		return ret
	}
	return *o.DirectConnect
}

// GetDirectConnectOk returns a tuple with the DirectConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetDirectConnectOk() (*bool, bool) {
	if o == nil || isNil(o.DirectConnect) {
		return nil, false
	}
	return o.DirectConnect, true
}

// HasDirectConnect returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasDirectConnect() bool {
	if o != nil && !isNil(o.DirectConnect) {
		return true
	}

	return false
}

// SetDirectConnect gets a reference to the given bool and assigns it to the DirectConnect field.
func (o *CreateConnectorRequest) SetDirectConnect(v bool) {
	o.DirectConnect = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateConnectorRequest) SetStatus(v string) {
	o.Status = &v
}

func (o CreateConnectorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ClassName) {
		toSerialize["className"] = o.ClassName
	}
	if !isNil(o.DirectConnect) {
		toSerialize["directConnect"] = o.DirectConnect
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateConnectorRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateConnectorRequest := _CreateConnectorRequest{}

	if err = json.Unmarshal(bytes, &varCreateConnectorRequest); err == nil {
		*o = CreateConnectorRequest(varCreateConnectorRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "className")
		delete(additionalProperties, "directConnect")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateConnectorRequest struct {
	value *CreateConnectorRequest
	isSet bool
}

func (v NullableCreateConnectorRequest) Get() *CreateConnectorRequest {
	return v.value
}

func (v *NullableCreateConnectorRequest) Set(val *CreateConnectorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectorRequest(val *CreateConnectorRequest) *NullableCreateConnectorRequest {
	return &NullableCreateConnectorRequest{value: val, isSet: true}
}

func (v NullableCreateConnectorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

