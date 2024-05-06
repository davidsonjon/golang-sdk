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

// checks if the SodExemptCriteria1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodExemptCriteria1{}

// SodExemptCriteria1 Details of the Entitlement criteria
type SodExemptCriteria1 struct {
	// If the entitlement already belonged to the user or not.
	Existing *bool `json:"existing,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	// Entitlement ID
	Id *string `json:"id,omitempty"`
	// Entitlement name
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodExemptCriteria1 SodExemptCriteria1

// NewSodExemptCriteria1 instantiates a new SodExemptCriteria1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodExemptCriteria1() *SodExemptCriteria1 {
	this := SodExemptCriteria1{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// NewSodExemptCriteria1WithDefaults instantiates a new SodExemptCriteria1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodExemptCriteria1WithDefaults() *SodExemptCriteria1 {
	this := SodExemptCriteria1{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// GetExisting returns the Existing field value if set, zero value otherwise.
func (o *SodExemptCriteria1) GetExisting() bool {
	if o == nil || isNil(o.Existing) {
		var ret bool
		return ret
	}
	return *o.Existing
}

// GetExistingOk returns a tuple with the Existing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria1) GetExistingOk() (*bool, bool) {
	if o == nil || isNil(o.Existing) {
		return nil, false
	}
	return o.Existing, true
}

// HasExisting returns a boolean if a field has been set.
func (o *SodExemptCriteria1) HasExisting() bool {
	if o != nil && !isNil(o.Existing) {
		return true
	}

	return false
}

// SetExisting gets a reference to the given bool and assigns it to the Existing field.
func (o *SodExemptCriteria1) SetExisting(v bool) {
	o.Existing = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SodExemptCriteria1) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria1) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SodExemptCriteria1) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *SodExemptCriteria1) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SodExemptCriteria1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SodExemptCriteria1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SodExemptCriteria1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SodExemptCriteria1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SodExemptCriteria1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SodExemptCriteria1) SetName(v string) {
	o.Name = &v
}

func (o SodExemptCriteria1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodExemptCriteria1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Existing) {
		toSerialize["existing"] = o.Existing
	}
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

func (o *SodExemptCriteria1) UnmarshalJSON(bytes []byte) (err error) {
	varSodExemptCriteria1 := _SodExemptCriteria1{}

	if err = json.Unmarshal(bytes, &varSodExemptCriteria1); err == nil {
	*o = SodExemptCriteria1(varSodExemptCriteria1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "existing")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodExemptCriteria1 struct {
	value *SodExemptCriteria1
	isSet bool
}

func (v NullableSodExemptCriteria1) Get() *SodExemptCriteria1 {
	return v.value
}

func (v *NullableSodExemptCriteria1) Set(val *SodExemptCriteria1) {
	v.value = val
	v.isSet = true
}

func (v NullableSodExemptCriteria1) IsSet() bool {
	return v.isSet
}

func (v *NullableSodExemptCriteria1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodExemptCriteria1(val *SodExemptCriteria1) *NullableSodExemptCriteria1 {
	return &NullableSodExemptCriteria1{value: val, isSet: true}
}

func (v NullableSodExemptCriteria1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodExemptCriteria1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


