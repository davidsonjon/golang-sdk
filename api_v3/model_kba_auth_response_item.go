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

// checks if the KbaAuthResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KbaAuthResponseItem{}

// KbaAuthResponseItem struct for KbaAuthResponseItem
type KbaAuthResponseItem struct {
	// The KBA question id
	QuestionId NullableString `json:"questionId,omitempty"`
	// Return true if verified
	IsVerified NullableBool `json:"isVerified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KbaAuthResponseItem KbaAuthResponseItem

// NewKbaAuthResponseItem instantiates a new KbaAuthResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKbaAuthResponseItem() *KbaAuthResponseItem {
	this := KbaAuthResponseItem{}
	return &this
}

// NewKbaAuthResponseItemWithDefaults instantiates a new KbaAuthResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKbaAuthResponseItemWithDefaults() *KbaAuthResponseItem {
	this := KbaAuthResponseItem{}
	return &this
}

// GetQuestionId returns the QuestionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KbaAuthResponseItem) GetQuestionId() string {
	if o == nil || isNil(o.QuestionId.Get()) {
		var ret string
		return ret
	}
	return *o.QuestionId.Get()
}

// GetQuestionIdOk returns a tuple with the QuestionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KbaAuthResponseItem) GetQuestionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuestionId.Get(), o.QuestionId.IsSet()
}

// HasQuestionId returns a boolean if a field has been set.
func (o *KbaAuthResponseItem) HasQuestionId() bool {
	if o != nil && o.QuestionId.IsSet() {
		return true
	}

	return false
}

// SetQuestionId gets a reference to the given NullableString and assigns it to the QuestionId field.
func (o *KbaAuthResponseItem) SetQuestionId(v string) {
	o.QuestionId.Set(&v)
}
// SetQuestionIdNil sets the value for QuestionId to be an explicit nil
func (o *KbaAuthResponseItem) SetQuestionIdNil() {
	o.QuestionId.Set(nil)
}

// UnsetQuestionId ensures that no value is present for QuestionId, not even an explicit nil
func (o *KbaAuthResponseItem) UnsetQuestionId() {
	o.QuestionId.Unset()
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KbaAuthResponseItem) GetIsVerified() bool {
	if o == nil || isNil(o.IsVerified.Get()) {
		var ret bool
		return ret
	}
	return *o.IsVerified.Get()
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KbaAuthResponseItem) GetIsVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsVerified.Get(), o.IsVerified.IsSet()
}

// HasIsVerified returns a boolean if a field has been set.
func (o *KbaAuthResponseItem) HasIsVerified() bool {
	if o != nil && o.IsVerified.IsSet() {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given NullableBool and assigns it to the IsVerified field.
func (o *KbaAuthResponseItem) SetIsVerified(v bool) {
	o.IsVerified.Set(&v)
}
// SetIsVerifiedNil sets the value for IsVerified to be an explicit nil
func (o *KbaAuthResponseItem) SetIsVerifiedNil() {
	o.IsVerified.Set(nil)
}

// UnsetIsVerified ensures that no value is present for IsVerified, not even an explicit nil
func (o *KbaAuthResponseItem) UnsetIsVerified() {
	o.IsVerified.Unset()
}

func (o KbaAuthResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KbaAuthResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.QuestionId.IsSet() {
		toSerialize["questionId"] = o.QuestionId.Get()
	}
	if o.IsVerified.IsSet() {
		toSerialize["isVerified"] = o.IsVerified.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KbaAuthResponseItem) UnmarshalJSON(bytes []byte) (err error) {
	varKbaAuthResponseItem := _KbaAuthResponseItem{}

	if err = json.Unmarshal(bytes, &varKbaAuthResponseItem); err == nil {
	*o = KbaAuthResponseItem(varKbaAuthResponseItem)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "questionId")
		delete(additionalProperties, "isVerified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKbaAuthResponseItem struct {
	value *KbaAuthResponseItem
	isSet bool
}

func (v NullableKbaAuthResponseItem) Get() *KbaAuthResponseItem {
	return v.value
}

func (v *NullableKbaAuthResponseItem) Set(val *KbaAuthResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKbaAuthResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKbaAuthResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKbaAuthResponseItem(val *KbaAuthResponseItem) *NullableKbaAuthResponseItem {
	return &NullableKbaAuthResponseItem{value: val, isSet: true}
}

func (v NullableKbaAuthResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKbaAuthResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


