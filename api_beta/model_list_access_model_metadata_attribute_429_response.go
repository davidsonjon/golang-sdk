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

// checks if the ListAccessModelMetadataAttribute429Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccessModelMetadataAttribute429Response{}

// ListAccessModelMetadataAttribute429Response struct for ListAccessModelMetadataAttribute429Response
type ListAccessModelMetadataAttribute429Response struct {
	// A message describing the error
	Message map[string]interface{} `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAccessModelMetadataAttribute429Response ListAccessModelMetadataAttribute429Response

// NewListAccessModelMetadataAttribute429Response instantiates a new ListAccessModelMetadataAttribute429Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccessModelMetadataAttribute429Response() *ListAccessModelMetadataAttribute429Response {
	this := ListAccessModelMetadataAttribute429Response{}
	return &this
}

// NewListAccessModelMetadataAttribute429ResponseWithDefaults instantiates a new ListAccessModelMetadataAttribute429Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccessModelMetadataAttribute429ResponseWithDefaults() *ListAccessModelMetadataAttribute429Response {
	this := ListAccessModelMetadataAttribute429Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ListAccessModelMetadataAttribute429Response) GetMessage() map[string]interface{} {
	if o == nil || isNil(o.Message) {
		var ret map[string]interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccessModelMetadataAttribute429Response) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Message) {
		return map[string]interface{}{}, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ListAccessModelMetadataAttribute429Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given map[string]interface{} and assigns it to the Message field.
func (o *ListAccessModelMetadataAttribute429Response) SetMessage(v map[string]interface{}) {
	o.Message = v
}

func (o ListAccessModelMetadataAttribute429Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAccessModelMetadataAttribute429Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAccessModelMetadataAttribute429Response) UnmarshalJSON(bytes []byte) (err error) {
	varListAccessModelMetadataAttribute429Response := _ListAccessModelMetadataAttribute429Response{}

	if err = json.Unmarshal(bytes, &varListAccessModelMetadataAttribute429Response); err == nil {
	*o = ListAccessModelMetadataAttribute429Response(varListAccessModelMetadataAttribute429Response)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAccessModelMetadataAttribute429Response struct {
	value *ListAccessModelMetadataAttribute429Response
	isSet bool
}

func (v NullableListAccessModelMetadataAttribute429Response) Get() *ListAccessModelMetadataAttribute429Response {
	return v.value
}

func (v *NullableListAccessModelMetadataAttribute429Response) Set(val *ListAccessModelMetadataAttribute429Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccessModelMetadataAttribute429Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccessModelMetadataAttribute429Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccessModelMetadataAttribute429Response(val *ListAccessModelMetadataAttribute429Response) *NullableListAccessModelMetadataAttribute429Response {
	return &NullableListAccessModelMetadataAttribute429Response{value: val, isSet: true}
}

func (v NullableListAccessModelMetadataAttribute429Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccessModelMetadataAttribute429Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

