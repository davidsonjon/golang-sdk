/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the NonEmployeeIdentityReferenceWithId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeIdentityReferenceWithId{}

// NonEmployeeIdentityReferenceWithId struct for NonEmployeeIdentityReferenceWithId
type NonEmployeeIdentityReferenceWithId struct {
	Type *NonEmployeeIdentityDtoType `json:"type,omitempty"`
	// Identity id
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeIdentityReferenceWithId NonEmployeeIdentityReferenceWithId

// NewNonEmployeeIdentityReferenceWithId instantiates a new NonEmployeeIdentityReferenceWithId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeIdentityReferenceWithId() *NonEmployeeIdentityReferenceWithId {
	this := NonEmployeeIdentityReferenceWithId{}
	return &this
}

// NewNonEmployeeIdentityReferenceWithIdWithDefaults instantiates a new NonEmployeeIdentityReferenceWithId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeIdentityReferenceWithIdWithDefaults() *NonEmployeeIdentityReferenceWithId {
	this := NonEmployeeIdentityReferenceWithId{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NonEmployeeIdentityReferenceWithId) GetType() NonEmployeeIdentityDtoType {
	if o == nil || isNil(o.Type) {
		var ret NonEmployeeIdentityDtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeIdentityReferenceWithId) GetTypeOk() (*NonEmployeeIdentityDtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NonEmployeeIdentityReferenceWithId) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NonEmployeeIdentityDtoType and assigns it to the Type field.
func (o *NonEmployeeIdentityReferenceWithId) SetType(v NonEmployeeIdentityDtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeIdentityReferenceWithId) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeIdentityReferenceWithId) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeIdentityReferenceWithId) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeIdentityReferenceWithId) SetId(v string) {
	o.Id = &v
}

func (o NonEmployeeIdentityReferenceWithId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeIdentityReferenceWithId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeIdentityReferenceWithId) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeIdentityReferenceWithId := _NonEmployeeIdentityReferenceWithId{}

	if err = json.Unmarshal(bytes, &varNonEmployeeIdentityReferenceWithId); err == nil {
	*o = NonEmployeeIdentityReferenceWithId(varNonEmployeeIdentityReferenceWithId)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeIdentityReferenceWithId struct {
	value *NonEmployeeIdentityReferenceWithId
	isSet bool
}

func (v NullableNonEmployeeIdentityReferenceWithId) Get() *NonEmployeeIdentityReferenceWithId {
	return v.value
}

func (v *NullableNonEmployeeIdentityReferenceWithId) Set(val *NonEmployeeIdentityReferenceWithId) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeIdentityReferenceWithId) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeIdentityReferenceWithId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeIdentityReferenceWithId(val *NonEmployeeIdentityReferenceWithId) *NullableNonEmployeeIdentityReferenceWithId {
	return &NullableNonEmployeeIdentityReferenceWithId{value: val, isSet: true}
}

func (v NullableNonEmployeeIdentityReferenceWithId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeIdentityReferenceWithId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


