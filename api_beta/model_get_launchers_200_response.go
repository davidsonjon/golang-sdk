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

// checks if the GetLaunchers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLaunchers200Response{}

// GetLaunchers200Response struct for GetLaunchers200Response
type GetLaunchers200Response struct {
	// Pagination marker
	Next *string `json:"next,omitempty"`
	Items []Launcher `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLaunchers200Response GetLaunchers200Response

// NewGetLaunchers200Response instantiates a new GetLaunchers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLaunchers200Response() *GetLaunchers200Response {
	this := GetLaunchers200Response{}
	return &this
}

// NewGetLaunchers200ResponseWithDefaults instantiates a new GetLaunchers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLaunchers200ResponseWithDefaults() *GetLaunchers200Response {
	this := GetLaunchers200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *GetLaunchers200Response) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLaunchers200Response) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *GetLaunchers200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *GetLaunchers200Response) SetNext(v string) {
	o.Next = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetLaunchers200Response) GetItems() []Launcher {
	if o == nil || IsNil(o.Items) {
		var ret []Launcher
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLaunchers200Response) GetItemsOk() ([]Launcher, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetLaunchers200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Launcher and assigns it to the Items field.
func (o *GetLaunchers200Response) SetItems(v []Launcher) {
	o.Items = v
}

func (o GetLaunchers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLaunchers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLaunchers200Response) UnmarshalJSON(data []byte) (err error) {
	varGetLaunchers200Response := _GetLaunchers200Response{}

	err = json.Unmarshal(data, &varGetLaunchers200Response)

	if err != nil {
		return err
	}

	*o = GetLaunchers200Response(varGetLaunchers200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLaunchers200Response struct {
	value *GetLaunchers200Response
	isSet bool
}

func (v NullableGetLaunchers200Response) Get() *GetLaunchers200Response {
	return v.value
}

func (v *NullableGetLaunchers200Response) Set(val *GetLaunchers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLaunchers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLaunchers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLaunchers200Response(val *GetLaunchers200Response) *NullableGetLaunchers200Response {
	return &NullableGetLaunchers200Response{value: val, isSet: true}
}

func (v NullableGetLaunchers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLaunchers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


