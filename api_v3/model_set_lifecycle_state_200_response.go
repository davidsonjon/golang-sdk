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

// checks if the SetLifecycleState200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetLifecycleState200Response{}

// SetLifecycleState200Response struct for SetLifecycleState200Response
type SetLifecycleState200Response struct {
	// The ID of the IdentityRequest object that is generated when the workflow launches. To follow the IdentityRequest, you can provide this ID with a [Get Account Activity request](https://developer.sailpoint.com/docs/api/v3/get-account-activity/). The response will contain relevant information about the IdentityRequest, such as its status.
	AccountActivityId *string `json:"accountActivityId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetLifecycleState200Response SetLifecycleState200Response

// NewSetLifecycleState200Response instantiates a new SetLifecycleState200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetLifecycleState200Response() *SetLifecycleState200Response {
	this := SetLifecycleState200Response{}
	return &this
}

// NewSetLifecycleState200ResponseWithDefaults instantiates a new SetLifecycleState200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetLifecycleState200ResponseWithDefaults() *SetLifecycleState200Response {
	this := SetLifecycleState200Response{}
	return &this
}

// GetAccountActivityId returns the AccountActivityId field value if set, zero value otherwise.
func (o *SetLifecycleState200Response) GetAccountActivityId() string {
	if o == nil || isNil(o.AccountActivityId) {
		var ret string
		return ret
	}
	return *o.AccountActivityId
}

// GetAccountActivityIdOk returns a tuple with the AccountActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLifecycleState200Response) GetAccountActivityIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountActivityId) {
		return nil, false
	}
	return o.AccountActivityId, true
}

// HasAccountActivityId returns a boolean if a field has been set.
func (o *SetLifecycleState200Response) HasAccountActivityId() bool {
	if o != nil && !isNil(o.AccountActivityId) {
		return true
	}

	return false
}

// SetAccountActivityId gets a reference to the given string and assigns it to the AccountActivityId field.
func (o *SetLifecycleState200Response) SetAccountActivityId(v string) {
	o.AccountActivityId = &v
}

func (o SetLifecycleState200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetLifecycleState200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountActivityId) {
		toSerialize["accountActivityId"] = o.AccountActivityId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetLifecycleState200Response) UnmarshalJSON(bytes []byte) (err error) {
	varSetLifecycleState200Response := _SetLifecycleState200Response{}

	if err = json.Unmarshal(bytes, &varSetLifecycleState200Response); err == nil {
	*o = SetLifecycleState200Response(varSetLifecycleState200Response)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountActivityId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetLifecycleState200Response struct {
	value *SetLifecycleState200Response
	isSet bool
}

func (v NullableSetLifecycleState200Response) Get() *SetLifecycleState200Response {
	return v.value
}

func (v *NullableSetLifecycleState200Response) Set(val *SetLifecycleState200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLifecycleState200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLifecycleState200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLifecycleState200Response(val *SetLifecycleState200Response) *NullableSetLifecycleState200Response {
	return &NullableSetLifecycleState200Response{value: val, isSet: true}
}

func (v NullableSetLifecycleState200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLifecycleState200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


