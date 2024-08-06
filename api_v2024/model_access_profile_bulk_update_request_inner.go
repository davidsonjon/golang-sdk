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

// checks if the AccessProfileBulkUpdateRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfileBulkUpdateRequestInner{}

// AccessProfileBulkUpdateRequestInner Access Profile's basic details.
type AccessProfileBulkUpdateRequestInner struct {
	// Access Profile ID.
	Id *string `json:"id,omitempty"`
	// Access Profile is requestable or not.
	Requestable *bool `json:"requestable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileBulkUpdateRequestInner AccessProfileBulkUpdateRequestInner

// NewAccessProfileBulkUpdateRequestInner instantiates a new AccessProfileBulkUpdateRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileBulkUpdateRequestInner() *AccessProfileBulkUpdateRequestInner {
	this := AccessProfileBulkUpdateRequestInner{}
	return &this
}

// NewAccessProfileBulkUpdateRequestInnerWithDefaults instantiates a new AccessProfileBulkUpdateRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileBulkUpdateRequestInnerWithDefaults() *AccessProfileBulkUpdateRequestInner {
	this := AccessProfileBulkUpdateRequestInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessProfileBulkUpdateRequestInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileBulkUpdateRequestInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessProfileBulkUpdateRequestInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessProfileBulkUpdateRequestInner) SetId(v string) {
	o.Id = &v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *AccessProfileBulkUpdateRequestInner) GetRequestable() bool {
	if o == nil || isNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileBulkUpdateRequestInner) GetRequestableOk() (*bool, bool) {
	if o == nil || isNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *AccessProfileBulkUpdateRequestInner) HasRequestable() bool {
	if o != nil && !isNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *AccessProfileBulkUpdateRequestInner) SetRequestable(v bool) {
	o.Requestable = &v
}

func (o AccessProfileBulkUpdateRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfileBulkUpdateRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessProfileBulkUpdateRequestInner) UnmarshalJSON(bytes []byte) (err error) {
	varAccessProfileBulkUpdateRequestInner := _AccessProfileBulkUpdateRequestInner{}

	if err = json.Unmarshal(bytes, &varAccessProfileBulkUpdateRequestInner); err == nil {
	*o = AccessProfileBulkUpdateRequestInner(varAccessProfileBulkUpdateRequestInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "requestable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileBulkUpdateRequestInner struct {
	value *AccessProfileBulkUpdateRequestInner
	isSet bool
}

func (v NullableAccessProfileBulkUpdateRequestInner) Get() *AccessProfileBulkUpdateRequestInner {
	return v.value
}

func (v *NullableAccessProfileBulkUpdateRequestInner) Set(val *AccessProfileBulkUpdateRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileBulkUpdateRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileBulkUpdateRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileBulkUpdateRequestInner(val *AccessProfileBulkUpdateRequestInner) *NullableAccessProfileBulkUpdateRequestInner {
	return &NullableAccessProfileBulkUpdateRequestInner{value: val, isSet: true}
}

func (v NullableAccessProfileBulkUpdateRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileBulkUpdateRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


