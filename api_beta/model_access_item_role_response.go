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

// checks if the AccessItemRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemRoleResponse{}

// AccessItemRoleResponse struct for AccessItemRoleResponse
type AccessItemRoleResponse struct {
	// the access item type. role in this case
	AccessType *string `json:"accessType,omitempty"`
	// the access item id
	Id *string `json:"id,omitempty"`
	// the role display name
	DisplayName *string `json:"displayName,omitempty"`
	// the description for the role
	Description *string `json:"description,omitempty"`
	// the associated source name if it exists
	SourceName *string `json:"sourceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRoleResponse AccessItemRoleResponse

// NewAccessItemRoleResponse instantiates a new AccessItemRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRoleResponse() *AccessItemRoleResponse {
	this := AccessItemRoleResponse{}
	return &this
}

// NewAccessItemRoleResponseWithDefaults instantiates a new AccessItemRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemRoleResponseWithDefaults() *AccessItemRoleResponse {
	this := AccessItemRoleResponse{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AccessItemRoleResponse) GetAccessType() string {
	if o == nil || isNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetAccessTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *AccessItemRoleResponse) SetAccessType(v string) {
	o.AccessType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemRoleResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemRoleResponse) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AccessItemRoleResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AccessItemRoleResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessItemRoleResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessItemRoleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessItemRoleResponse) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessItemRoleResponse) SetSourceName(v string) {
	o.SourceName = &v
}

func (o AccessItemRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemRoleResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemRoleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccessItemRoleResponse := _AccessItemRoleResponse{}

	if err = json.Unmarshal(bytes, &varAccessItemRoleResponse); err == nil {
	*o = AccessItemRoleResponse(varAccessItemRoleResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "sourceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemRoleResponse struct {
	value *AccessItemRoleResponse
	isSet bool
}

func (v NullableAccessItemRoleResponse) Get() *AccessItemRoleResponse {
	return v.value
}

func (v *NullableAccessItemRoleResponse) Set(val *AccessItemRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemRoleResponse(val *AccessItemRoleResponse) *NullableAccessItemRoleResponse {
	return &NullableAccessItemRoleResponse{value: val, isSet: true}
}

func (v NullableAccessItemRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


