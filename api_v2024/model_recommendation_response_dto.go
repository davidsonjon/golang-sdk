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

// checks if the RecommendationResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationResponseDto{}

// RecommendationResponseDto struct for RecommendationResponseDto
type RecommendationResponseDto struct {
	Response []RecommendationResponse `json:"response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationResponseDto RecommendationResponseDto

// NewRecommendationResponseDto instantiates a new RecommendationResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationResponseDto() *RecommendationResponseDto {
	this := RecommendationResponseDto{}
	return &this
}

// NewRecommendationResponseDtoWithDefaults instantiates a new RecommendationResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationResponseDtoWithDefaults() *RecommendationResponseDto {
	this := RecommendationResponseDto{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *RecommendationResponseDto) GetResponse() []RecommendationResponse {
	if o == nil || IsNil(o.Response) {
		var ret []RecommendationResponse
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationResponseDto) GetResponseOk() ([]RecommendationResponse, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RecommendationResponseDto) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []RecommendationResponse and assigns it to the Response field.
func (o *RecommendationResponseDto) SetResponse(v []RecommendationResponse) {
	o.Response = v
}

func (o RecommendationResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecommendationResponseDto) UnmarshalJSON(data []byte) (err error) {
	varRecommendationResponseDto := _RecommendationResponseDto{}

	err = json.Unmarshal(data, &varRecommendationResponseDto)

	if err != nil {
		return err
	}

	*o = RecommendationResponseDto(varRecommendationResponseDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationResponseDto struct {
	value *RecommendationResponseDto
	isSet bool
}

func (v NullableRecommendationResponseDto) Get() *RecommendationResponseDto {
	return v.value
}

func (v *NullableRecommendationResponseDto) Set(val *RecommendationResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationResponseDto(val *RecommendationResponseDto) *NullableRecommendationResponseDto {
	return &NullableRecommendationResponseDto{value: val, isSet: true}
}

func (v NullableRecommendationResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


