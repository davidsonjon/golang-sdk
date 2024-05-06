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

// checks if the IdentityHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityHistoryResponse{}

// IdentityHistoryResponse struct for IdentityHistoryResponse
type IdentityHistoryResponse struct {
	// the identity ID
	Id *string `json:"id,omitempty"`
	// the display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// the date when the identity record was created
	Snapshot *string `json:"snapshot,omitempty"`
	// the date when the identity was deleted
	DeletedDate *string `json:"deletedDate,omitempty"`
	// A map containing the count of each access item
	AccessItemCount *map[string]string `json:"accessItemCount,omitempty"`
	// A map containing the identity attributes
	Attributes *map[string]string `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityHistoryResponse IdentityHistoryResponse

// NewIdentityHistoryResponse instantiates a new IdentityHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityHistoryResponse() *IdentityHistoryResponse {
	this := IdentityHistoryResponse{}
	return &this
}

// NewIdentityHistoryResponseWithDefaults instantiates a new IdentityHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityHistoryResponseWithDefaults() *IdentityHistoryResponse {
	this := IdentityHistoryResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityHistoryResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHistoryResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityHistoryResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityHistoryResponse) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentityHistoryResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHistoryResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityHistoryResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityHistoryResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *IdentityHistoryResponse) GetSnapshot() string {
	if o == nil || isNil(o.Snapshot) {
		var ret string
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHistoryResponse) GetSnapshotOk() (*string, bool) {
	if o == nil || isNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *IdentityHistoryResponse) HasSnapshot() bool {
	if o != nil && !isNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *IdentityHistoryResponse) SetSnapshot(v string) {
	o.Snapshot = &v
}

// GetDeletedDate returns the DeletedDate field value if set, zero value otherwise.
func (o *IdentityHistoryResponse) GetDeletedDate() string {
	if o == nil || isNil(o.DeletedDate) {
		var ret string
		return ret
	}
	return *o.DeletedDate
}

// GetDeletedDateOk returns a tuple with the DeletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHistoryResponse) GetDeletedDateOk() (*string, bool) {
	if o == nil || isNil(o.DeletedDate) {
		return nil, false
	}
	return o.DeletedDate, true
}

// HasDeletedDate returns a boolean if a field has been set.
func (o *IdentityHistoryResponse) HasDeletedDate() bool {
	if o != nil && !isNil(o.DeletedDate) {
		return true
	}

	return false
}

// SetDeletedDate gets a reference to the given string and assigns it to the DeletedDate field.
func (o *IdentityHistoryResponse) SetDeletedDate(v string) {
	o.DeletedDate = &v
}

// GetAccessItemCount returns the AccessItemCount field value if set, zero value otherwise.
func (o *IdentityHistoryResponse) GetAccessItemCount() map[string]string {
	if o == nil || isNil(o.AccessItemCount) {
		var ret map[string]string
		return ret
	}
	return *o.AccessItemCount
}

// GetAccessItemCountOk returns a tuple with the AccessItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHistoryResponse) GetAccessItemCountOk() (*map[string]string, bool) {
	if o == nil || isNil(o.AccessItemCount) {
		return nil, false
	}
	return o.AccessItemCount, true
}

// HasAccessItemCount returns a boolean if a field has been set.
func (o *IdentityHistoryResponse) HasAccessItemCount() bool {
	if o != nil && !isNil(o.AccessItemCount) {
		return true
	}

	return false
}

// SetAccessItemCount gets a reference to the given map[string]string and assigns it to the AccessItemCount field.
func (o *IdentityHistoryResponse) SetAccessItemCount(v map[string]string) {
	o.AccessItemCount = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IdentityHistoryResponse) GetAttributes() map[string]string {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHistoryResponse) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IdentityHistoryResponse) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *IdentityHistoryResponse) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o IdentityHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !isNil(o.DeletedDate) {
		toSerialize["deletedDate"] = o.DeletedDate
	}
	if !isNil(o.AccessItemCount) {
		toSerialize["accessItemCount"] = o.AccessItemCount
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityHistoryResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityHistoryResponse := _IdentityHistoryResponse{}

	if err = json.Unmarshal(bytes, &varIdentityHistoryResponse); err == nil {
	*o = IdentityHistoryResponse(varIdentityHistoryResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "snapshot")
		delete(additionalProperties, "deletedDate")
		delete(additionalProperties, "accessItemCount")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityHistoryResponse struct {
	value *IdentityHistoryResponse
	isSet bool
}

func (v NullableIdentityHistoryResponse) Get() *IdentityHistoryResponse {
	return v.value
}

func (v *NullableIdentityHistoryResponse) Set(val *IdentityHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityHistoryResponse(val *IdentityHistoryResponse) *NullableIdentityHistoryResponse {
	return &NullableIdentityHistoryResponse{value: val, isSet: true}
}

func (v NullableIdentityHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


