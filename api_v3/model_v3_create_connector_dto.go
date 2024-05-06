/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the V3CreateConnectorDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3CreateConnectorDto{}

// V3CreateConnectorDto struct for V3CreateConnectorDto
type V3CreateConnectorDto struct {
	// The connector name. Need to be unique per tenant. The name will able be used to derive a url friendly unique scriptname that will be in response. Script name can then be used for all update endpoints
	Name string `json:"name"`
	// The connector type. If not specified will be defaulted to 'custom '+name
	Type *string `json:"type,omitempty"`
	// The connector class name. If you are implementing openconnector standard (what is recommended), then this need to be set to sailpoint.connector.OpenConnectorAdapter
	ClassName string `json:"className"`
	// true if the source is a direct connect source
	DirectConnect *bool `json:"directConnect,omitempty"`
	// The connector status
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V3CreateConnectorDto V3CreateConnectorDto

// NewV3CreateConnectorDto instantiates a new V3CreateConnectorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3CreateConnectorDto(name string, className string) *V3CreateConnectorDto {
	this := V3CreateConnectorDto{}
	this.Name = name
	this.ClassName = className
	var directConnect bool = true
	this.DirectConnect = &directConnect
	return &this
}

// NewV3CreateConnectorDtoWithDefaults instantiates a new V3CreateConnectorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3CreateConnectorDtoWithDefaults() *V3CreateConnectorDto {
	this := V3CreateConnectorDto{}
	var directConnect bool = true
	this.DirectConnect = &directConnect
	return &this
}

// GetName returns the Name field value
func (o *V3CreateConnectorDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V3CreateConnectorDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V3CreateConnectorDto) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V3CreateConnectorDto) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3CreateConnectorDto) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V3CreateConnectorDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V3CreateConnectorDto) SetType(v string) {
	o.Type = &v
}

// GetClassName returns the ClassName field value
func (o *V3CreateConnectorDto) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *V3CreateConnectorDto) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *V3CreateConnectorDto) SetClassName(v string) {
	o.ClassName = v
}

// GetDirectConnect returns the DirectConnect field value if set, zero value otherwise.
func (o *V3CreateConnectorDto) GetDirectConnect() bool {
	if o == nil || isNil(o.DirectConnect) {
		var ret bool
		return ret
	}
	return *o.DirectConnect
}

// GetDirectConnectOk returns a tuple with the DirectConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3CreateConnectorDto) GetDirectConnectOk() (*bool, bool) {
	if o == nil || isNil(o.DirectConnect) {
		return nil, false
	}
	return o.DirectConnect, true
}

// HasDirectConnect returns a boolean if a field has been set.
func (o *V3CreateConnectorDto) HasDirectConnect() bool {
	if o != nil && !isNil(o.DirectConnect) {
		return true
	}

	return false
}

// SetDirectConnect gets a reference to the given bool and assigns it to the DirectConnect field.
func (o *V3CreateConnectorDto) SetDirectConnect(v bool) {
	o.DirectConnect = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V3CreateConnectorDto) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3CreateConnectorDto) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V3CreateConnectorDto) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V3CreateConnectorDto) SetStatus(v string) {
	o.Status = &v
}

func (o V3CreateConnectorDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3CreateConnectorDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["className"] = o.ClassName
	if !isNil(o.DirectConnect) {
		toSerialize["directConnect"] = o.DirectConnect
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V3CreateConnectorDto) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"className",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV3CreateConnectorDto := _V3CreateConnectorDto{}

	if err = json.Unmarshal(bytes, &varV3CreateConnectorDto); err == nil {
	*o = V3CreateConnectorDto(varV3CreateConnectorDto)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "className")
		delete(additionalProperties, "directConnect")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV3CreateConnectorDto struct {
	value *V3CreateConnectorDto
	isSet bool
}

func (v NullableV3CreateConnectorDto) Get() *V3CreateConnectorDto {
	return v.value
}

func (v *NullableV3CreateConnectorDto) Set(val *V3CreateConnectorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableV3CreateConnectorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableV3CreateConnectorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3CreateConnectorDto(val *V3CreateConnectorDto) *NullableV3CreateConnectorDto {
	return &NullableV3CreateConnectorDto{value: val, isSet: true}
}

func (v NullableV3CreateConnectorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3CreateConnectorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


