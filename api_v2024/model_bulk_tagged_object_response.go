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

// checks if the BulkTaggedObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkTaggedObjectResponse{}

// BulkTaggedObjectResponse struct for BulkTaggedObjectResponse
type BulkTaggedObjectResponse struct {
	ObjectRefs []TaggedObjectDto `json:"objectRefs,omitempty"`
	// Label to be applied to an Object
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkTaggedObjectResponse BulkTaggedObjectResponse

// NewBulkTaggedObjectResponse instantiates a new BulkTaggedObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkTaggedObjectResponse() *BulkTaggedObjectResponse {
	this := BulkTaggedObjectResponse{}
	return &this
}

// NewBulkTaggedObjectResponseWithDefaults instantiates a new BulkTaggedObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkTaggedObjectResponseWithDefaults() *BulkTaggedObjectResponse {
	this := BulkTaggedObjectResponse{}
	return &this
}

// GetObjectRefs returns the ObjectRefs field value if set, zero value otherwise.
func (o *BulkTaggedObjectResponse) GetObjectRefs() []TaggedObjectDto {
	if o == nil || isNil(o.ObjectRefs) {
		var ret []TaggedObjectDto
		return ret
	}
	return o.ObjectRefs
}

// GetObjectRefsOk returns a tuple with the ObjectRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTaggedObjectResponse) GetObjectRefsOk() ([]TaggedObjectDto, bool) {
	if o == nil || isNil(o.ObjectRefs) {
		return nil, false
	}
	return o.ObjectRefs, true
}

// HasObjectRefs returns a boolean if a field has been set.
func (o *BulkTaggedObjectResponse) HasObjectRefs() bool {
	if o != nil && !isNil(o.ObjectRefs) {
		return true
	}

	return false
}

// SetObjectRefs gets a reference to the given []TaggedObjectDto and assigns it to the ObjectRefs field.
func (o *BulkTaggedObjectResponse) SetObjectRefs(v []TaggedObjectDto) {
	o.ObjectRefs = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkTaggedObjectResponse) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTaggedObjectResponse) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkTaggedObjectResponse) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BulkTaggedObjectResponse) SetTags(v []string) {
	o.Tags = v
}

func (o BulkTaggedObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkTaggedObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectRefs) {
		toSerialize["objectRefs"] = o.ObjectRefs
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkTaggedObjectResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBulkTaggedObjectResponse := _BulkTaggedObjectResponse{}

	if err = json.Unmarshal(bytes, &varBulkTaggedObjectResponse); err == nil {
	*o = BulkTaggedObjectResponse(varBulkTaggedObjectResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "objectRefs")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkTaggedObjectResponse struct {
	value *BulkTaggedObjectResponse
	isSet bool
}

func (v NullableBulkTaggedObjectResponse) Get() *BulkTaggedObjectResponse {
	return v.value
}

func (v *NullableBulkTaggedObjectResponse) Set(val *BulkTaggedObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkTaggedObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkTaggedObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkTaggedObjectResponse(val *BulkTaggedObjectResponse) *NullableBulkTaggedObjectResponse {
	return &NullableBulkTaggedObjectResponse{value: val, isSet: true}
}

func (v NullableBulkTaggedObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkTaggedObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


