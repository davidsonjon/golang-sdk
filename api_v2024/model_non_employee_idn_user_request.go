/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the NonEmployeeIdnUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeIdnUserRequest{}

// NonEmployeeIdnUserRequest struct for NonEmployeeIdnUserRequest
type NonEmployeeIdnUserRequest struct {
	// Identity id.
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeIdnUserRequest NonEmployeeIdnUserRequest

// NewNonEmployeeIdnUserRequest instantiates a new NonEmployeeIdnUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeIdnUserRequest(id string) *NonEmployeeIdnUserRequest {
	this := NonEmployeeIdnUserRequest{}
	this.Id = id
	return &this
}

// NewNonEmployeeIdnUserRequestWithDefaults instantiates a new NonEmployeeIdnUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeIdnUserRequestWithDefaults() *NonEmployeeIdnUserRequest {
	this := NonEmployeeIdnUserRequest{}
	return &this
}

// GetId returns the Id field value
func (o *NonEmployeeIdnUserRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeIdnUserRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NonEmployeeIdnUserRequest) SetId(v string) {
	o.Id = v
}

func (o NonEmployeeIdnUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeIdnUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeIdnUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varNonEmployeeIdnUserRequest := _NonEmployeeIdnUserRequest{}

	err = json.Unmarshal(data, &varNonEmployeeIdnUserRequest)

	if err != nil {
		return err
	}

	*o = NonEmployeeIdnUserRequest(varNonEmployeeIdnUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeIdnUserRequest struct {
	value *NonEmployeeIdnUserRequest
	isSet bool
}

func (v NullableNonEmployeeIdnUserRequest) Get() *NonEmployeeIdnUserRequest {
	return v.value
}

func (v *NullableNonEmployeeIdnUserRequest) Set(val *NonEmployeeIdnUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeIdnUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeIdnUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeIdnUserRequest(val *NonEmployeeIdnUserRequest) *NullableNonEmployeeIdnUserRequest {
	return &NullableNonEmployeeIdnUserRequest{value: val, isSet: true}
}

func (v NullableNonEmployeeIdnUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeIdnUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


