/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the ListWorkgroupMembers200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorkgroupMembers200ResponseInner{}

// ListWorkgroupMembers200ResponseInner Identity of workgroup member.
type ListWorkgroupMembers200ResponseInner struct {
	// Workgroup member identity DTO type.
	Type *string `json:"type,omitempty"`
	// Workgroup member identity ID.
	Id *string `json:"id,omitempty"`
	// Workgroup member identity display name.
	Name *string `json:"name,omitempty"`
	// Workgroup member identity email.
	Email *string `json:"email,omitempty"`
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

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListWorkgroupMembers200ResponseInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListWorkgroupMembers200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListWorkgroupMembers200ResponseInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkgroupMembers200ResponseInner) GetIdOk() (*string, bool) {
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

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListWorkgroupMembers200ResponseInner) SetId(v string) {
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

func (o ListWorkgroupMembers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorkgroupMembers200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
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
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
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


