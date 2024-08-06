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

// checks if the WorkgroupDtoOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupDtoOwner{}

// WorkgroupDtoOwner struct for WorkgroupDtoOwner
type WorkgroupDtoOwner struct {
	// Owner's DTO type.
	Type *string `json:"type,omitempty"`
	// Owner's identity ID.
	Id *string `json:"id,omitempty"`
	// Owner's name.
	Name *string `json:"name,omitempty"`
	// The display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// The primary email address of the identity
	EmailAddress *string `json:"emailAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupDtoOwner WorkgroupDtoOwner

// NewWorkgroupDtoOwner instantiates a new WorkgroupDtoOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupDtoOwner() *WorkgroupDtoOwner {
	this := WorkgroupDtoOwner{}
	return &this
}

// NewWorkgroupDtoOwnerWithDefaults instantiates a new WorkgroupDtoOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupDtoOwnerWithDefaults() *WorkgroupDtoOwner {
	this := WorkgroupDtoOwner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkgroupDtoOwner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDtoOwner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkgroupDtoOwner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkgroupDtoOwner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkgroupDtoOwner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDtoOwner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkgroupDtoOwner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkgroupDtoOwner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkgroupDtoOwner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDtoOwner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkgroupDtoOwner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkgroupDtoOwner) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *WorkgroupDtoOwner) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDtoOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *WorkgroupDtoOwner) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *WorkgroupDtoOwner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *WorkgroupDtoOwner) GetEmailAddress() string {
	if o == nil || isNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDtoOwner) GetEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *WorkgroupDtoOwner) HasEmailAddress() bool {
	if o != nil && !isNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *WorkgroupDtoOwner) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

func (o WorkgroupDtoOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupDtoOwner) ToMap() (map[string]interface{}, error) {
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
	// skip: displayName is readOnly
	// skip: emailAddress is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupDtoOwner) UnmarshalJSON(bytes []byte) (err error) {
	varWorkgroupDtoOwner := _WorkgroupDtoOwner{}

	if err = json.Unmarshal(bytes, &varWorkgroupDtoOwner); err == nil {
	*o = WorkgroupDtoOwner(varWorkgroupDtoOwner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "emailAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupDtoOwner struct {
	value *WorkgroupDtoOwner
	isSet bool
}

func (v NullableWorkgroupDtoOwner) Get() *WorkgroupDtoOwner {
	return v.value
}

func (v *NullableWorkgroupDtoOwner) Set(val *WorkgroupDtoOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupDtoOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupDtoOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupDtoOwner(val *WorkgroupDtoOwner) *NullableWorkgroupDtoOwner {
	return &NullableWorkgroupDtoOwner{value: val, isSet: true}
}

func (v NullableWorkgroupDtoOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupDtoOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


