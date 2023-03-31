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

// checks if the AccessItemAppResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemAppResponse{}

// AccessItemAppResponse struct for AccessItemAppResponse
type AccessItemAppResponse struct {
	// the access item type. entitlement in this case
	AccessType *string `json:"accessType,omitempty"`
	// the access item id
	Id *string `json:"id,omitempty"`
	// the access profile display name
	DisplayName *string `json:"displayName,omitempty"`
	// the associated source name if it exists
	SourceName *string `json:"sourceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemAppResponse AccessItemAppResponse

// NewAccessItemAppResponse instantiates a new AccessItemAppResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemAppResponse() *AccessItemAppResponse {
	this := AccessItemAppResponse{}
	return &this
}

// NewAccessItemAppResponseWithDefaults instantiates a new AccessItemAppResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemAppResponseWithDefaults() *AccessItemAppResponse {
	this := AccessItemAppResponse{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AccessItemAppResponse) GetAccessType() string {
	if o == nil || isNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAppResponse) GetAccessTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessItemAppResponse) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *AccessItemAppResponse) SetAccessType(v string) {
	o.AccessType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemAppResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAppResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemAppResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemAppResponse) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AccessItemAppResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAppResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessItemAppResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AccessItemAppResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessItemAppResponse) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAppResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessItemAppResponse) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessItemAppResponse) SetSourceName(v string) {
	o.SourceName = &v
}

func (o AccessItemAppResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemAppResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemAppResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccessItemAppResponse := _AccessItemAppResponse{}

	if err = json.Unmarshal(bytes, &varAccessItemAppResponse); err == nil {
		*o = AccessItemAppResponse(varAccessItemAppResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "sourceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemAppResponse struct {
	value *AccessItemAppResponse
	isSet bool
}

func (v NullableAccessItemAppResponse) Get() *AccessItemAppResponse {
	return v.value
}

func (v *NullableAccessItemAppResponse) Set(val *AccessItemAppResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemAppResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemAppResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemAppResponse(val *AccessItemAppResponse) *NullableAccessItemAppResponse {
	return &NullableAccessItemAppResponse{value: val, isSet: true}
}

func (v NullableAccessItemAppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemAppResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


