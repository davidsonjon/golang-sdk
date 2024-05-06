/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the ObjectMappingBulkPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectMappingBulkPatchRequest{}

// ObjectMappingBulkPatchRequest struct for ObjectMappingBulkPatchRequest
type ObjectMappingBulkPatchRequest struct {
	// Map of id of the object mapping to a JsonPatchOperation describing what to patch on that object mapping.
	Patches map[string][]JsonPatchOperation `json:"patches"`
	AdditionalProperties map[string]interface{}
}

type _ObjectMappingBulkPatchRequest ObjectMappingBulkPatchRequest

// NewObjectMappingBulkPatchRequest instantiates a new ObjectMappingBulkPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectMappingBulkPatchRequest(patches map[string][]JsonPatchOperation) *ObjectMappingBulkPatchRequest {
	this := ObjectMappingBulkPatchRequest{}
	this.Patches = patches
	return &this
}

// NewObjectMappingBulkPatchRequestWithDefaults instantiates a new ObjectMappingBulkPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectMappingBulkPatchRequestWithDefaults() *ObjectMappingBulkPatchRequest {
	this := ObjectMappingBulkPatchRequest{}
	return &this
}

// GetPatches returns the Patches field value
func (o *ObjectMappingBulkPatchRequest) GetPatches() map[string][]JsonPatchOperation {
	if o == nil {
		var ret map[string][]JsonPatchOperation
		return ret
	}

	return o.Patches
}

// GetPatchesOk returns a tuple with the Patches field value
// and a boolean to check if the value has been set.
func (o *ObjectMappingBulkPatchRequest) GetPatchesOk() (*map[string][]JsonPatchOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patches, true
}

// SetPatches sets field value
func (o *ObjectMappingBulkPatchRequest) SetPatches(v map[string][]JsonPatchOperation) {
	o.Patches = v
}

func (o ObjectMappingBulkPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectMappingBulkPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["patches"] = o.Patches

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObjectMappingBulkPatchRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"patches",
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

	varObjectMappingBulkPatchRequest := _ObjectMappingBulkPatchRequest{}

	if err = json.Unmarshal(bytes, &varObjectMappingBulkPatchRequest); err == nil {
	*o = ObjectMappingBulkPatchRequest(varObjectMappingBulkPatchRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "patches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectMappingBulkPatchRequest struct {
	value *ObjectMappingBulkPatchRequest
	isSet bool
}

func (v NullableObjectMappingBulkPatchRequest) Get() *ObjectMappingBulkPatchRequest {
	return v.value
}

func (v *NullableObjectMappingBulkPatchRequest) Set(val *ObjectMappingBulkPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectMappingBulkPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectMappingBulkPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectMappingBulkPatchRequest(val *ObjectMappingBulkPatchRequest) *NullableObjectMappingBulkPatchRequest {
	return &NullableObjectMappingBulkPatchRequest{value: val, isSet: true}
}

func (v NullableObjectMappingBulkPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectMappingBulkPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


