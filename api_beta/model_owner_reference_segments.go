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

// checks if the OwnerReferenceSegments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerReferenceSegments{}

// OwnerReferenceSegments The owner of this object.
type OwnerReferenceSegments struct {
	// Owner type. This field must be either left null or set to 'IDENTITY' on input, otherwise a 400 Bad Request error will result.
	Type *string `json:"type,omitempty"`
	// Identity id
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the owner. It may be left null or omitted in a POST or PATCH. If set, it must match the current value of the owner's display name, otherwise a 400 Bad Request error will result.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OwnerReferenceSegments OwnerReferenceSegments

// NewOwnerReferenceSegments instantiates a new OwnerReferenceSegments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerReferenceSegments() *OwnerReferenceSegments {
	this := OwnerReferenceSegments{}
	return &this
}

// NewOwnerReferenceSegmentsWithDefaults instantiates a new OwnerReferenceSegments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerReferenceSegmentsWithDefaults() *OwnerReferenceSegments {
	this := OwnerReferenceSegments{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OwnerReferenceSegments) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReferenceSegments) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OwnerReferenceSegments) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OwnerReferenceSegments) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OwnerReferenceSegments) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReferenceSegments) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OwnerReferenceSegments) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OwnerReferenceSegments) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OwnerReferenceSegments) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerReferenceSegments) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OwnerReferenceSegments) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OwnerReferenceSegments) SetName(v string) {
	o.Name = &v
}

func (o OwnerReferenceSegments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerReferenceSegments) ToMap() (map[string]interface{}, error) {
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

func (o *OwnerReferenceSegments) UnmarshalJSON(bytes []byte) (err error) {
	varOwnerReferenceSegments := _OwnerReferenceSegments{}

	if err = json.Unmarshal(bytes, &varOwnerReferenceSegments); err == nil {
	*o = OwnerReferenceSegments(varOwnerReferenceSegments)
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

type NullableOwnerReferenceSegments struct {
	value *OwnerReferenceSegments
	isSet bool
}

func (v NullableOwnerReferenceSegments) Get() *OwnerReferenceSegments {
	return v.value
}

func (v *NullableOwnerReferenceSegments) Set(val *OwnerReferenceSegments) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerReferenceSegments) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerReferenceSegments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerReferenceSegments(val *OwnerReferenceSegments) *NullableOwnerReferenceSegments {
	return &NullableOwnerReferenceSegments{value: val, isSet: true}
}

func (v NullableOwnerReferenceSegments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerReferenceSegments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


