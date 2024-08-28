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

// checks if the ApprovalReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalReference{}

// ApprovalReference Reference objects related to the approval
type ApprovalReference struct {
	// Id of the reference object
	Id *string `json:"id,omitempty"`
	// What reference object does this ID correspond to
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalReference ApprovalReference

// NewApprovalReference instantiates a new ApprovalReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalReference() *ApprovalReference {
	this := ApprovalReference{}
	return &this
}

// NewApprovalReferenceWithDefaults instantiates a new ApprovalReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalReferenceWithDefaults() *ApprovalReference {
	this := ApprovalReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApprovalReference) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalReference) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApprovalReference) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApprovalReference) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApprovalReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApprovalReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApprovalReference) SetType(v string) {
	o.Type = &v
}

func (o ApprovalReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalReference) UnmarshalJSON(data []byte) (err error) {
	varApprovalReference := _ApprovalReference{}

	err = json.Unmarshal(data, &varApprovalReference)

	if err != nil {
		return err
	}

	*o = ApprovalReference(varApprovalReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalReference struct {
	value *ApprovalReference
	isSet bool
}

func (v NullableApprovalReference) Get() *ApprovalReference {
	return v.value
}

func (v *NullableApprovalReference) Set(val *ApprovalReference) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalReference) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalReference(val *ApprovalReference) *NullableApprovalReference {
	return &NullableApprovalReference{value: val, isSet: true}
}

func (v NullableApprovalReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


