/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the SetIconRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetIconRequest{}

// SetIconRequest struct for SetIconRequest
type SetIconRequest struct {
	// file with icon. Allowed mime-types ['image/png', 'image/jpeg']
	Image *os.File `json:"image"`
	AdditionalProperties map[string]interface{}
}

type _SetIconRequest SetIconRequest

// NewSetIconRequest instantiates a new SetIconRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetIconRequest(image *os.File) *SetIconRequest {
	this := SetIconRequest{}
	this.Image = image
	return &this
}

// NewSetIconRequestWithDefaults instantiates a new SetIconRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetIconRequestWithDefaults() *SetIconRequest {
	this := SetIconRequest{}
	return &this
}

// GetImage returns the Image field value
func (o *SetIconRequest) GetImage() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *SetIconRequest) GetImageOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *SetIconRequest) SetImage(v *os.File) {
	o.Image = v
}

func (o SetIconRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetIconRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetIconRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSetIconRequest := _SetIconRequest{}

	if err = json.Unmarshal(bytes, &varSetIconRequest); err == nil {
	*o = SetIconRequest(varSetIconRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetIconRequest struct {
	value *SetIconRequest
	isSet bool
}

func (v NullableSetIconRequest) Get() *SetIconRequest {
	return v.value
}

func (v *NullableSetIconRequest) Set(val *SetIconRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetIconRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetIconRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetIconRequest(val *SetIconRequest) *NullableSetIconRequest {
	return &NullableSetIconRequest{value: val, isSet: true}
}

func (v NullableSetIconRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetIconRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


