/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ViolationContextPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolationContextPolicy{}

// ViolationContextPolicy The types of objects supported for SOD violations
type ViolationContextPolicy struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ViolationContextPolicy ViolationContextPolicy

// NewViolationContextPolicy instantiates a new ViolationContextPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolationContextPolicy() *ViolationContextPolicy {
	this := ViolationContextPolicy{}
	return &this
}

// NewViolationContextPolicyWithDefaults instantiates a new ViolationContextPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationContextPolicyWithDefaults() *ViolationContextPolicy {
	this := ViolationContextPolicy{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ViolationContextPolicy) GetType() map[string]interface{} {
	if o == nil || isNil(o.Type) {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationContextPolicy) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ViolationContextPolicy) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *ViolationContextPolicy) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViolationContextPolicy) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationContextPolicy) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViolationContextPolicy) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ViolationContextPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViolationContextPolicy) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationContextPolicy) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViolationContextPolicy) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ViolationContextPolicy) SetName(v string) {
	o.Name = &v
}

func (o ViolationContextPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViolationContextPolicy) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ViolationContextPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varViolationContextPolicy := _ViolationContextPolicy{}

	if err = json.Unmarshal(bytes, &varViolationContextPolicy); err == nil {
		*o = ViolationContextPolicy(varViolationContextPolicy)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableViolationContextPolicy struct {
	value *ViolationContextPolicy
	isSet bool
}

func (v NullableViolationContextPolicy) Get() *ViolationContextPolicy {
	return v.value
}

func (v *NullableViolationContextPolicy) Set(val *ViolationContextPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableViolationContextPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableViolationContextPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolationContextPolicy(val *ViolationContextPolicy) *NullableViolationContextPolicy {
	return &NullableViolationContextPolicy{value: val, isSet: true}
}

func (v NullableViolationContextPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolationContextPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


