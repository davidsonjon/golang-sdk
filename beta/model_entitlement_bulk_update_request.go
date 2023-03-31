/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the EntitlementBulkUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementBulkUpdateRequest{}

// EntitlementBulkUpdateRequest struct for EntitlementBulkUpdateRequest
type EntitlementBulkUpdateRequest struct {
	// List of entitlement ids to update
	EntitlementIds []string `json:"entitlementIds"`
	JsonPatch []JsonPatchOperation `json:"jsonPatch"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementBulkUpdateRequest EntitlementBulkUpdateRequest

// NewEntitlementBulkUpdateRequest instantiates a new EntitlementBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBulkUpdateRequest(entitlementIds []string, jsonPatch []JsonPatchOperation) *EntitlementBulkUpdateRequest {
	this := EntitlementBulkUpdateRequest{}
	this.EntitlementIds = entitlementIds
	this.JsonPatch = jsonPatch
	return &this
}

// NewEntitlementBulkUpdateRequestWithDefaults instantiates a new EntitlementBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBulkUpdateRequestWithDefaults() *EntitlementBulkUpdateRequest {
	this := EntitlementBulkUpdateRequest{}
	return &this
}

// GetEntitlementIds returns the EntitlementIds field value
func (o *EntitlementBulkUpdateRequest) GetEntitlementIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EntitlementIds
}

// GetEntitlementIdsOk returns a tuple with the EntitlementIds field value
// and a boolean to check if the value has been set.
func (o *EntitlementBulkUpdateRequest) GetEntitlementIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementIds, true
}

// SetEntitlementIds sets field value
func (o *EntitlementBulkUpdateRequest) SetEntitlementIds(v []string) {
	o.EntitlementIds = v
}

// GetJsonPatch returns the JsonPatch field value
func (o *EntitlementBulkUpdateRequest) GetJsonPatch() []JsonPatchOperation {
	if o == nil {
		var ret []JsonPatchOperation
		return ret
	}

	return o.JsonPatch
}

// GetJsonPatchOk returns a tuple with the JsonPatch field value
// and a boolean to check if the value has been set.
func (o *EntitlementBulkUpdateRequest) GetJsonPatchOk() ([]JsonPatchOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonPatch, true
}

// SetJsonPatch sets field value
func (o *EntitlementBulkUpdateRequest) SetJsonPatch(v []JsonPatchOperation) {
	o.JsonPatch = v
}

func (o EntitlementBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementBulkUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entitlementIds"] = o.EntitlementIds
	toSerialize["jsonPatch"] = o.JsonPatch

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementBulkUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementBulkUpdateRequest := _EntitlementBulkUpdateRequest{}

	if err = json.Unmarshal(bytes, &varEntitlementBulkUpdateRequest); err == nil {
		*o = EntitlementBulkUpdateRequest(varEntitlementBulkUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementIds")
		delete(additionalProperties, "jsonPatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementBulkUpdateRequest struct {
	value *EntitlementBulkUpdateRequest
	isSet bool
}

func (v NullableEntitlementBulkUpdateRequest) Get() *EntitlementBulkUpdateRequest {
	return v.value
}

func (v *NullableEntitlementBulkUpdateRequest) Set(val *EntitlementBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBulkUpdateRequest(val *EntitlementBulkUpdateRequest) *NullableEntitlementBulkUpdateRequest {
	return &NullableEntitlementBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableEntitlementBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


