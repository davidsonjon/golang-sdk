/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the IdentityOwnershipAssociationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityOwnershipAssociationDetails{}

// IdentityOwnershipAssociationDetails struct for IdentityOwnershipAssociationDetails
type IdentityOwnershipAssociationDetails struct {
	// list of all the resource associations for the identity
	AssociationDetails []IdentityOwnershipAssociationDetailsAssociationDetailsInner `json:"associationDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityOwnershipAssociationDetails IdentityOwnershipAssociationDetails

// NewIdentityOwnershipAssociationDetails instantiates a new IdentityOwnershipAssociationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityOwnershipAssociationDetails() *IdentityOwnershipAssociationDetails {
	this := IdentityOwnershipAssociationDetails{}
	return &this
}

// NewIdentityOwnershipAssociationDetailsWithDefaults instantiates a new IdentityOwnershipAssociationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityOwnershipAssociationDetailsWithDefaults() *IdentityOwnershipAssociationDetails {
	this := IdentityOwnershipAssociationDetails{}
	return &this
}

// GetAssociationDetails returns the AssociationDetails field value if set, zero value otherwise.
func (o *IdentityOwnershipAssociationDetails) GetAssociationDetails() []IdentityOwnershipAssociationDetailsAssociationDetailsInner {
	if o == nil || IsNil(o.AssociationDetails) {
		var ret []IdentityOwnershipAssociationDetailsAssociationDetailsInner
		return ret
	}
	return o.AssociationDetails
}

// GetAssociationDetailsOk returns a tuple with the AssociationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOwnershipAssociationDetails) GetAssociationDetailsOk() ([]IdentityOwnershipAssociationDetailsAssociationDetailsInner, bool) {
	if o == nil || IsNil(o.AssociationDetails) {
		return nil, false
	}
	return o.AssociationDetails, true
}

// HasAssociationDetails returns a boolean if a field has been set.
func (o *IdentityOwnershipAssociationDetails) HasAssociationDetails() bool {
	if o != nil && !IsNil(o.AssociationDetails) {
		return true
	}

	return false
}

// SetAssociationDetails gets a reference to the given []IdentityOwnershipAssociationDetailsAssociationDetailsInner and assigns it to the AssociationDetails field.
func (o *IdentityOwnershipAssociationDetails) SetAssociationDetails(v []IdentityOwnershipAssociationDetailsAssociationDetailsInner) {
	o.AssociationDetails = v
}

func (o IdentityOwnershipAssociationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityOwnershipAssociationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociationDetails) {
		toSerialize["associationDetails"] = o.AssociationDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityOwnershipAssociationDetails) UnmarshalJSON(data []byte) (err error) {
	varIdentityOwnershipAssociationDetails := _IdentityOwnershipAssociationDetails{}

	err = json.Unmarshal(data, &varIdentityOwnershipAssociationDetails)

	if err != nil {
		return err
	}

	*o = IdentityOwnershipAssociationDetails(varIdentityOwnershipAssociationDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associationDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityOwnershipAssociationDetails struct {
	value *IdentityOwnershipAssociationDetails
	isSet bool
}

func (v NullableIdentityOwnershipAssociationDetails) Get() *IdentityOwnershipAssociationDetails {
	return v.value
}

func (v *NullableIdentityOwnershipAssociationDetails) Set(val *IdentityOwnershipAssociationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityOwnershipAssociationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityOwnershipAssociationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityOwnershipAssociationDetails(val *IdentityOwnershipAssociationDetails) *NullableIdentityOwnershipAssociationDetails {
	return &NullableIdentityOwnershipAssociationDetails{value: val, isSet: true}
}

func (v NullableIdentityOwnershipAssociationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityOwnershipAssociationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


