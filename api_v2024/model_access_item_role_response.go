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
	// the date the role is no longer assigned to the specified identity
	RemoveDate *string `json:"removeDate,omitempty"`
	// indicates whether the role is revocable
	Revocable bool `json:"revocable"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRoleResponse AccessItemRoleResponse

// NewAccessItemRoleResponse instantiates a new AccessItemRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRoleResponse(revocable bool) *AccessItemRoleResponse {
	this := AccessItemRoleResponse{}
	this.Revocable = revocable
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
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessItemRoleResponse) SetSourceName(v string) {
	o.SourceName = &v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise.
func (o *AccessItemRoleResponse) GetRemoveDate() string {
	if o == nil || IsNil(o.RemoveDate) {
		var ret string
		return ret
	}
	return *o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetRemoveDateOk() (*string, bool) {
	if o == nil || IsNil(o.RemoveDate) {
		return nil, false
	}
	return o.RemoveDate, true
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *AccessItemRoleResponse) HasRemoveDate() bool {
	if o != nil && !IsNil(o.RemoveDate) {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given string and assigns it to the RemoveDate field.
func (o *AccessItemRoleResponse) SetRemoveDate(v string) {
	o.RemoveDate = &v
}

// GetRevocable returns the Revocable field value
func (o *AccessItemRoleResponse) GetRevocable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value
// and a boolean to check if the value has been set.
func (o *AccessItemRoleResponse) GetRevocableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revocable, true
}

// SetRevocable sets field value
func (o *AccessItemRoleResponse) SetRevocable(v bool) {
	o.Revocable = v
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
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !IsNil(o.RemoveDate) {
		toSerialize["removeDate"] = o.RemoveDate
	}
	toSerialize["revocable"] = o.Revocable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemRoleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revocable",
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

	varAccessItemRoleResponse := _AccessItemRoleResponse{}

	err = json.Unmarshal(data, &varAccessItemRoleResponse)

	if err != nil {
		return err
	}

	*o = AccessItemRoleResponse(varAccessItemRoleResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "removeDate")
		delete(additionalProperties, "revocable")
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


