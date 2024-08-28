/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the KbaAnswerRequestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KbaAnswerRequestItem{}

// KbaAnswerRequestItem struct for KbaAnswerRequestItem
type KbaAnswerRequestItem struct {
	// Question Id
	Id string `json:"id"`
	// An answer for the KBA question
	Answer string `json:"answer"`
	AdditionalProperties map[string]interface{}
}

type _KbaAnswerRequestItem KbaAnswerRequestItem

// NewKbaAnswerRequestItem instantiates a new KbaAnswerRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKbaAnswerRequestItem(id string, answer string) *KbaAnswerRequestItem {
	this := KbaAnswerRequestItem{}
	this.Id = id
	this.Answer = answer
	return &this
}

// NewKbaAnswerRequestItemWithDefaults instantiates a new KbaAnswerRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKbaAnswerRequestItemWithDefaults() *KbaAnswerRequestItem {
	this := KbaAnswerRequestItem{}
	return &this
}

// GetId returns the Id field value
func (o *KbaAnswerRequestItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KbaAnswerRequestItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KbaAnswerRequestItem) SetId(v string) {
	o.Id = v
}

// GetAnswer returns the Answer field value
func (o *KbaAnswerRequestItem) GetAnswer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value
// and a boolean to check if the value has been set.
func (o *KbaAnswerRequestItem) GetAnswerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Answer, true
}

// SetAnswer sets field value
func (o *KbaAnswerRequestItem) SetAnswer(v string) {
	o.Answer = v
}

func (o KbaAnswerRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KbaAnswerRequestItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["answer"] = o.Answer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KbaAnswerRequestItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"answer",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKbaAnswerRequestItem := _KbaAnswerRequestItem{}

	err = json.Unmarshal(data, &varKbaAnswerRequestItem)

	if err != nil {
		return err
	}

	*o = KbaAnswerRequestItem(varKbaAnswerRequestItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "answer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKbaAnswerRequestItem struct {
	value *KbaAnswerRequestItem
	isSet bool
}

func (v NullableKbaAnswerRequestItem) Get() *KbaAnswerRequestItem {
	return v.value
}

func (v *NullableKbaAnswerRequestItem) Set(val *KbaAnswerRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKbaAnswerRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKbaAnswerRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKbaAnswerRequestItem(val *KbaAnswerRequestItem) *NullableKbaAnswerRequestItem {
	return &NullableKbaAnswerRequestItem{value: val, isSet: true}
}

func (v NullableKbaAnswerRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKbaAnswerRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


