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

// checks if the RecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationRequest{}

// RecommendationRequest struct for RecommendationRequest
type RecommendationRequest struct {
	// The identity ID
	IdentityId *string `json:"identityId,omitempty"`
	Item *AccessItemRef `json:"item,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationRequest RecommendationRequest

// NewRecommendationRequest instantiates a new RecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationRequest() *RecommendationRequest {
	this := RecommendationRequest{}
	return &this
}

// NewRecommendationRequestWithDefaults instantiates a new RecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationRequestWithDefaults() *RecommendationRequest {
	this := RecommendationRequest{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *RecommendationRequest) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *RecommendationRequest) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *RecommendationRequest) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *RecommendationRequest) GetItem() AccessItemRef {
	if o == nil || isNil(o.Item) {
		var ret AccessItemRef
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetItemOk() (*AccessItemRef, bool) {
	if o == nil || isNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *RecommendationRequest) HasItem() bool {
	if o != nil && !isNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given AccessItemRef and assigns it to the Item field.
func (o *RecommendationRequest) SetItem(v AccessItemRef) {
	o.Item = &v
}

func (o RecommendationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.Item) {
		toSerialize["item"] = o.Item
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecommendationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationRequest := _RecommendationRequest{}

	if err = json.Unmarshal(bytes, &varRecommendationRequest); err == nil {
	*o = RecommendationRequest(varRecommendationRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "item")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationRequest struct {
	value *RecommendationRequest
	isSet bool
}

func (v NullableRecommendationRequest) Get() *RecommendationRequest {
	return v.value
}

func (v *NullableRecommendationRequest) Set(val *RecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationRequest(val *RecommendationRequest) *NullableRecommendationRequest {
	return &NullableRecommendationRequest{value: val, isSet: true}
}

func (v NullableRecommendationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


