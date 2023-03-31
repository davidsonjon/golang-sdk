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

// checks if the ListWorkgroupMembers200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorkgroupMembers200ResponseInner{}

// ListWorkgroupMembers200ResponseInner struct for ListWorkgroupMembers200ResponseInner
type ListWorkgroupMembers200ResponseInner struct {
	Alias *string `json:"alias,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListWorkgroupMembers200ResponseInner ListWorkgroupMembers200ResponseInner

// NewListWorkgroupMembers200ResponseInner instantiates a new ListWorkgroupMembers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkgroupMembers200ResponseInner() *ListWorkgroupMembers200ResponseInner {
	this := ListWorkgroupMembers200ResponseInner{}
	return &this
}

// NewListWorkgroupMembers200ResponseInnerWithDefaults instantiates a new ListWorkgroupMembers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkgroupMembers200ResponseInnerWithDefaults() *ListWorkgroupMembers200ResponseInner {
	this := ListWorkgroupMembers200ResponseInner{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ListWorkgroupMembers200ResponseInner) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ListWorkgroupMembers200ResponseInner) SetAlias(v string) {
	o.Alias = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ListWorkgroupMembers200ResponseInner) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ListWorkgroupMembers200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetExternalId() string {
	if o == nil || isNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetExternalIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ListWorkgroupMembers200ResponseInner) HasExternalId() bool {
	if o != nil && !isNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ListWorkgroupMembers200ResponseInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListWorkgroupMembers200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListWorkgroupMembers200ResponseInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListWorkgroupMembers200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListWorkgroupMembers200ResponseInner) SetName(v string) {
	o.Name = &v
}

func (o ListWorkgroupMembers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorkgroupMembers200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
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

func (o *ListWorkgroupMembers200ResponseInner) UnmarshalJSON(bytes []byte) (err error) {
	varListWorkgroupMembers200ResponseInner := _ListWorkgroupMembers200ResponseInner{}

	if err = json.Unmarshal(bytes, &varListWorkgroupMembers200ResponseInner); err == nil {
		*o = ListWorkgroupMembers200ResponseInner(varListWorkgroupMembers200ResponseInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "alias")
		delete(additionalProperties, "email")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWorkgroupMembers200ResponseInner struct {
	value *ListWorkgroupMembers200ResponseInner
	isSet bool
}

func (v NullableListWorkgroupMembers200ResponseInner) Get() *ListWorkgroupMembers200ResponseInner {
	return v.value
}

func (v *NullableListWorkgroupMembers200ResponseInner) Set(val *ListWorkgroupMembers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkgroupMembers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkgroupMembers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkgroupMembers200ResponseInner(val *ListWorkgroupMembers200ResponseInner) *NullableListWorkgroupMembers200ResponseInner {
	return &NullableListWorkgroupMembers200ResponseInner{value: val, isSet: true}
}

func (v NullableListWorkgroupMembers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkgroupMembers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


