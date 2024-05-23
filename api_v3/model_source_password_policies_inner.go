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

// checks if the SourcePasswordPoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourcePasswordPoliciesInner{}

// SourcePasswordPoliciesInner struct for SourcePasswordPoliciesInner
type SourcePasswordPoliciesInner struct {
	// Type of object being referenced.
	Type *string `json:"type,omitempty"`
	// Policy ID.
	Id *string `json:"id,omitempty"`
	// Policy's human-readable display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourcePasswordPoliciesInner SourcePasswordPoliciesInner

// NewSourcePasswordPoliciesInner instantiates a new SourcePasswordPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcePasswordPoliciesInner() *SourcePasswordPoliciesInner {
	this := SourcePasswordPoliciesInner{}
	return &this
}

// NewSourcePasswordPoliciesInnerWithDefaults instantiates a new SourcePasswordPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcePasswordPoliciesInnerWithDefaults() *SourcePasswordPoliciesInner {
	this := SourcePasswordPoliciesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourcePasswordPoliciesInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePasswordPoliciesInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourcePasswordPoliciesInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourcePasswordPoliciesInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourcePasswordPoliciesInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePasswordPoliciesInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourcePasswordPoliciesInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourcePasswordPoliciesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourcePasswordPoliciesInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePasswordPoliciesInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourcePasswordPoliciesInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourcePasswordPoliciesInner) SetName(v string) {
	o.Name = &v
}

func (o SourcePasswordPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourcePasswordPoliciesInner) ToMap() (map[string]interface{}, error) {
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

func (o *SourcePasswordPoliciesInner) UnmarshalJSON(bytes []byte) (err error) {
	varSourcePasswordPoliciesInner := _SourcePasswordPoliciesInner{}

	if err = json.Unmarshal(bytes, &varSourcePasswordPoliciesInner); err == nil {
	*o = SourcePasswordPoliciesInner(varSourcePasswordPoliciesInner)
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

type NullableSourcePasswordPoliciesInner struct {
	value *SourcePasswordPoliciesInner
	isSet bool
}

func (v NullableSourcePasswordPoliciesInner) Get() *SourcePasswordPoliciesInner {
	return v.value
}

func (v *NullableSourcePasswordPoliciesInner) Set(val *SourcePasswordPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcePasswordPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcePasswordPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcePasswordPoliciesInner(val *SourcePasswordPoliciesInner) *NullableSourcePasswordPoliciesInner {
	return &NullableSourcePasswordPoliciesInner{value: val, isSet: true}
}

func (v NullableSourcePasswordPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcePasswordPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


