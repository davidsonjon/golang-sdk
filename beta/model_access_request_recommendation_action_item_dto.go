/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// checks if the AccessRequestRecommendationActionItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestRecommendationActionItemDto{}

// AccessRequestRecommendationActionItemDto struct for AccessRequestRecommendationActionItemDto
type AccessRequestRecommendationActionItemDto struct {
	// The identity ID taking the action.
	IdentityId string `json:"identityId"`
	Access AccessRequestRecommendationItem `json:"access"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestRecommendationActionItemDto AccessRequestRecommendationActionItemDto

// NewAccessRequestRecommendationActionItemDto instantiates a new AccessRequestRecommendationActionItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestRecommendationActionItemDto(identityId string, access AccessRequestRecommendationItem) *AccessRequestRecommendationActionItemDto {
	this := AccessRequestRecommendationActionItemDto{}
	this.IdentityId = identityId
	this.Access = access
	return &this
}

// NewAccessRequestRecommendationActionItemDtoWithDefaults instantiates a new AccessRequestRecommendationActionItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestRecommendationActionItemDtoWithDefaults() *AccessRequestRecommendationActionItemDto {
	this := AccessRequestRecommendationActionItemDto{}
	return &this
}

// GetIdentityId returns the IdentityId field value
func (o *AccessRequestRecommendationActionItemDto) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationActionItemDto) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *AccessRequestRecommendationActionItemDto) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetAccess returns the Access field value
func (o *AccessRequestRecommendationActionItemDto) GetAccess() AccessRequestRecommendationItem {
	if o == nil {
		var ret AccessRequestRecommendationItem
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationActionItemDto) GetAccessOk() (*AccessRequestRecommendationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *AccessRequestRecommendationActionItemDto) SetAccess(v AccessRequestRecommendationItem) {
	o.Access = v
}

func (o AccessRequestRecommendationActionItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestRecommendationActionItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identityId"] = o.IdentityId
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestRecommendationActionItemDto) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identityId",
		"access",
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

	varAccessRequestRecommendationActionItemDto := _AccessRequestRecommendationActionItemDto{}

	if err = json.Unmarshal(bytes, &varAccessRequestRecommendationActionItemDto); err == nil {
	*o = AccessRequestRecommendationActionItemDto(varAccessRequestRecommendationActionItemDto)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "access")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestRecommendationActionItemDto struct {
	value *AccessRequestRecommendationActionItemDto
	isSet bool
}

func (v NullableAccessRequestRecommendationActionItemDto) Get() *AccessRequestRecommendationActionItemDto {
	return v.value
}

func (v *NullableAccessRequestRecommendationActionItemDto) Set(val *AccessRequestRecommendationActionItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestRecommendationActionItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestRecommendationActionItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestRecommendationActionItemDto(val *AccessRequestRecommendationActionItemDto) *NullableAccessRequestRecommendationActionItemDto {
	return &NullableAccessRequestRecommendationActionItemDto{value: val, isSet: true}
}

func (v NullableAccessRequestRecommendationActionItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestRecommendationActionItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


