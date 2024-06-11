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

// checks if the SodPolicyOwnerRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodPolicyOwnerRef{}

// SodPolicyOwnerRef The owner of the SOD policy.
type SodPolicyOwnerRef struct {
	// Owner type.
	Type *string `json:"type,omitempty"`
	// Owner's ID.
	Id *string `json:"id,omitempty"`
	// Owner's name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodPolicyOwnerRef SodPolicyOwnerRef

// NewSodPolicyOwnerRef instantiates a new SodPolicyOwnerRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodPolicyOwnerRef() *SodPolicyOwnerRef {
	this := SodPolicyOwnerRef{}
	return &this
}

// NewSodPolicyOwnerRefWithDefaults instantiates a new SodPolicyOwnerRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodPolicyOwnerRefWithDefaults() *SodPolicyOwnerRef {
	this := SodPolicyOwnerRef{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SodPolicyOwnerRef) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicyOwnerRef) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SodPolicyOwnerRef) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SodPolicyOwnerRef) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SodPolicyOwnerRef) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicyOwnerRef) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SodPolicyOwnerRef) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SodPolicyOwnerRef) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SodPolicyOwnerRef) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicyOwnerRef) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SodPolicyOwnerRef) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SodPolicyOwnerRef) SetName(v string) {
	o.Name = &v
}

func (o SodPolicyOwnerRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodPolicyOwnerRef) ToMap() (map[string]interface{}, error) {
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

func (o *SodPolicyOwnerRef) UnmarshalJSON(bytes []byte) (err error) {
	varSodPolicyOwnerRef := _SodPolicyOwnerRef{}

	if err = json.Unmarshal(bytes, &varSodPolicyOwnerRef); err == nil {
	*o = SodPolicyOwnerRef(varSodPolicyOwnerRef)
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

type NullableSodPolicyOwnerRef struct {
	value *SodPolicyOwnerRef
	isSet bool
}

func (v NullableSodPolicyOwnerRef) Get() *SodPolicyOwnerRef {
	return v.value
}

func (v *NullableSodPolicyOwnerRef) Set(val *SodPolicyOwnerRef) {
	v.value = val
	v.isSet = true
}

func (v NullableSodPolicyOwnerRef) IsSet() bool {
	return v.isSet
}

func (v *NullableSodPolicyOwnerRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodPolicyOwnerRef(val *SodPolicyOwnerRef) *NullableSodPolicyOwnerRef {
	return &NullableSodPolicyOwnerRef{value: val, isSet: true}
}

func (v NullableSodPolicyOwnerRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodPolicyOwnerRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


