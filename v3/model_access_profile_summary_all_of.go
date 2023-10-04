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

// checks if the AccessProfileSummaryAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfileSummaryAllOf{}

// AccessProfileSummaryAllOf struct for AccessProfileSummaryAllOf
type AccessProfileSummaryAllOf struct {
	Source *Reference `json:"source,omitempty"`
	Owner *DisplayReference `json:"owner,omitempty"`
	Revocable *bool `json:"revocable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileSummaryAllOf AccessProfileSummaryAllOf

// NewAccessProfileSummaryAllOf instantiates a new AccessProfileSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileSummaryAllOf() *AccessProfileSummaryAllOf {
	this := AccessProfileSummaryAllOf{}
	return &this
}

// NewAccessProfileSummaryAllOfWithDefaults instantiates a new AccessProfileSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileSummaryAllOfWithDefaults() *AccessProfileSummaryAllOf {
	this := AccessProfileSummaryAllOf{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AccessProfileSummaryAllOf) GetSource() Reference {
	if o == nil || isNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummaryAllOf) GetSourceOk() (*Reference, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AccessProfileSummaryAllOf) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *AccessProfileSummaryAllOf) SetSource(v Reference) {
	o.Source = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccessProfileSummaryAllOf) GetOwner() DisplayReference {
	if o == nil || isNil(o.Owner) {
		var ret DisplayReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummaryAllOf) GetOwnerOk() (*DisplayReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccessProfileSummaryAllOf) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DisplayReference and assigns it to the Owner field.
func (o *AccessProfileSummaryAllOf) SetOwner(v DisplayReference) {
	o.Owner = &v
}

// GetRevocable returns the Revocable field value if set, zero value otherwise.
func (o *AccessProfileSummaryAllOf) GetRevocable() bool {
	if o == nil || isNil(o.Revocable) {
		var ret bool
		return ret
	}
	return *o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummaryAllOf) GetRevocableOk() (*bool, bool) {
	if o == nil || isNil(o.Revocable) {
		return nil, false
	}
	return o.Revocable, true
}

// HasRevocable returns a boolean if a field has been set.
func (o *AccessProfileSummaryAllOf) HasRevocable() bool {
	if o != nil && !isNil(o.Revocable) {
		return true
	}

	return false
}

// SetRevocable gets a reference to the given bool and assigns it to the Revocable field.
func (o *AccessProfileSummaryAllOf) SetRevocable(v bool) {
	o.Revocable = &v
}

func (o AccessProfileSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfileSummaryAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Revocable) {
		toSerialize["revocable"] = o.Revocable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessProfileSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessProfileSummaryAllOf := _AccessProfileSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varAccessProfileSummaryAllOf); err == nil {
		*o = AccessProfileSummaryAllOf(varAccessProfileSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "revocable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileSummaryAllOf struct {
	value *AccessProfileSummaryAllOf
	isSet bool
}

func (v NullableAccessProfileSummaryAllOf) Get() *AccessProfileSummaryAllOf {
	return v.value
}

func (v *NullableAccessProfileSummaryAllOf) Set(val *AccessProfileSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileSummaryAllOf(val *AccessProfileSummaryAllOf) *NullableAccessProfileSummaryAllOf {
	return &NullableAccessProfileSummaryAllOf{value: val, isSet: true}
}

func (v NullableAccessProfileSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


