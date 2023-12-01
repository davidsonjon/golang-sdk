/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the ListWorkgroups200ResponseInnerOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorkgroups200ResponseInnerOwner{}

// ListWorkgroups200ResponseInnerOwner struct for ListWorkgroups200ResponseInnerOwner
type ListWorkgroups200ResponseInnerOwner struct {
	DisplayName *string `json:"displayName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListWorkgroups200ResponseInnerOwner ListWorkgroups200ResponseInnerOwner

// NewListWorkgroups200ResponseInnerOwner instantiates a new ListWorkgroups200ResponseInnerOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkgroups200ResponseInnerOwner() *ListWorkgroups200ResponseInnerOwner {
	this := ListWorkgroups200ResponseInnerOwner{}
	return &this
}

// NewListWorkgroups200ResponseInnerOwnerWithDefaults instantiates a new ListWorkgroups200ResponseInnerOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkgroups200ResponseInnerOwnerWithDefaults() *ListWorkgroups200ResponseInnerOwner {
	this := ListWorkgroups200ResponseInnerOwner{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListWorkgroups200ResponseInnerOwner) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroups200ResponseInnerOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListWorkgroups200ResponseInnerOwner) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListWorkgroups200ResponseInnerOwner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ListWorkgroups200ResponseInnerOwner) GetEmailAddress() string {
	if o == nil || isNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroups200ResponseInnerOwner) GetEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ListWorkgroups200ResponseInnerOwner) HasEmailAddress() bool {
	if o != nil && !isNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ListWorkgroups200ResponseInnerOwner) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListWorkgroups200ResponseInnerOwner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroups200ResponseInnerOwner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListWorkgroups200ResponseInnerOwner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListWorkgroups200ResponseInnerOwner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListWorkgroups200ResponseInnerOwner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroups200ResponseInnerOwner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListWorkgroups200ResponseInnerOwner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListWorkgroups200ResponseInnerOwner) SetName(v string) {
	o.Name = &v
}

func (o ListWorkgroups200ResponseInnerOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorkgroups200ResponseInnerOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWorkgroups200ResponseInnerOwner) UnmarshalJSON(bytes []byte) (err error) {
	varListWorkgroups200ResponseInnerOwner := _ListWorkgroups200ResponseInnerOwner{}

	if err = json.Unmarshal(bytes, &varListWorkgroups200ResponseInnerOwner); err == nil {
	*o = ListWorkgroups200ResponseInnerOwner(varListWorkgroups200ResponseInnerOwner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "emailAddress")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWorkgroups200ResponseInnerOwner struct {
	value *ListWorkgroups200ResponseInnerOwner
	isSet bool
}

func (v NullableListWorkgroups200ResponseInnerOwner) Get() *ListWorkgroups200ResponseInnerOwner {
	return v.value
}

func (v *NullableListWorkgroups200ResponseInnerOwner) Set(val *ListWorkgroups200ResponseInnerOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkgroups200ResponseInnerOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkgroups200ResponseInnerOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkgroups200ResponseInnerOwner(val *ListWorkgroups200ResponseInnerOwner) *NullableListWorkgroups200ResponseInnerOwner {
	return &NullableListWorkgroups200ResponseInnerOwner{value: val, isSet: true}
}

func (v NullableListWorkgroups200ResponseInnerOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkgroups200ResponseInnerOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


