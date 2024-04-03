/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the ConnectedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedObject{}

// ConnectedObject struct for ConnectedObject
type ConnectedObject struct {
	Type *ConnectedObjectType `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable name of Connected object
	Name *string `json:"name,omitempty"`
	// Description of the Connected object.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectedObject ConnectedObject

// NewConnectedObject instantiates a new ConnectedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedObject() *ConnectedObject {
	this := ConnectedObject{}
	return &this
}

// NewConnectedObjectWithDefaults instantiates a new ConnectedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedObjectWithDefaults() *ConnectedObject {
	this := ConnectedObject{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectedObject) GetType() ConnectedObjectType {
	if o == nil || isNil(o.Type) {
		var ret ConnectedObjectType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedObject) GetTypeOk() (*ConnectedObjectType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectedObject) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConnectedObjectType and assigns it to the Type field.
func (o *ConnectedObject) SetType(v ConnectedObjectType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectedObject) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedObject) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectedObject) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectedObject) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectedObject) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedObject) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectedObject) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectedObject) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectedObject) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedObject) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectedObject) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectedObject) SetDescription(v string) {
	o.Description = &v
}

func (o ConnectedObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectedObject) UnmarshalJSON(bytes []byte) (err error) {
	varConnectedObject := _ConnectedObject{}

	if err = json.Unmarshal(bytes, &varConnectedObject); err == nil {
	*o = ConnectedObject(varConnectedObject)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectedObject struct {
	value *ConnectedObject
	isSet bool
}

func (v NullableConnectedObject) Get() *ConnectedObject {
	return v.value
}

func (v *NullableConnectedObject) Set(val *ConnectedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedObject(val *ConnectedObject) *NullableConnectedObject {
	return &NullableConnectedObject{value: val, isSet: true}
}

func (v NullableConnectedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


