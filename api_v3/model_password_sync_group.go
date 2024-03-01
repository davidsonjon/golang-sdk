/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the PasswordSyncGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordSyncGroup{}

// PasswordSyncGroup struct for PasswordSyncGroup
type PasswordSyncGroup struct {
	// ID of the sync group
	Id *string `json:"id,omitempty"`
	// Name of the sync group
	Name *string `json:"name,omitempty"`
	// ID of the password policy
	PasswordPolicyId *string `json:"passwordPolicyId,omitempty"`
	// List of password managed sources IDs
	SourceIds []string `json:"sourceIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordSyncGroup PasswordSyncGroup

// NewPasswordSyncGroup instantiates a new PasswordSyncGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordSyncGroup() *PasswordSyncGroup {
	this := PasswordSyncGroup{}
	return &this
}

// NewPasswordSyncGroupWithDefaults instantiates a new PasswordSyncGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordSyncGroupWithDefaults() *PasswordSyncGroup {
	this := PasswordSyncGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PasswordSyncGroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSyncGroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PasswordSyncGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PasswordSyncGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PasswordSyncGroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSyncGroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PasswordSyncGroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PasswordSyncGroup) SetName(v string) {
	o.Name = &v
}

// GetPasswordPolicyId returns the PasswordPolicyId field value if set, zero value otherwise.
func (o *PasswordSyncGroup) GetPasswordPolicyId() string {
	if o == nil || isNil(o.PasswordPolicyId) {
		var ret string
		return ret
	}
	return *o.PasswordPolicyId
}

// GetPasswordPolicyIdOk returns a tuple with the PasswordPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSyncGroup) GetPasswordPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.PasswordPolicyId) {
		return nil, false
	}
	return o.PasswordPolicyId, true
}

// HasPasswordPolicyId returns a boolean if a field has been set.
func (o *PasswordSyncGroup) HasPasswordPolicyId() bool {
	if o != nil && !isNil(o.PasswordPolicyId) {
		return true
	}

	return false
}

// SetPasswordPolicyId gets a reference to the given string and assigns it to the PasswordPolicyId field.
func (o *PasswordSyncGroup) SetPasswordPolicyId(v string) {
	o.PasswordPolicyId = &v
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise.
func (o *PasswordSyncGroup) GetSourceIds() []string {
	if o == nil || isNil(o.SourceIds) {
		var ret []string
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSyncGroup) GetSourceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SourceIds) {
		return nil, false
	}
	return o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *PasswordSyncGroup) HasSourceIds() bool {
	if o != nil && !isNil(o.SourceIds) {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []string and assigns it to the SourceIds field.
func (o *PasswordSyncGroup) SetSourceIds(v []string) {
	o.SourceIds = v
}

func (o PasswordSyncGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordSyncGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PasswordPolicyId) {
		toSerialize["passwordPolicyId"] = o.PasswordPolicyId
	}
	if !isNil(o.SourceIds) {
		toSerialize["sourceIds"] = o.SourceIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordSyncGroup) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordSyncGroup := _PasswordSyncGroup{}

	if err = json.Unmarshal(bytes, &varPasswordSyncGroup); err == nil {
	*o = PasswordSyncGroup(varPasswordSyncGroup)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "passwordPolicyId")
		delete(additionalProperties, "sourceIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordSyncGroup struct {
	value *PasswordSyncGroup
	isSet bool
}

func (v NullablePasswordSyncGroup) Get() *PasswordSyncGroup {
	return v.value
}

func (v *NullablePasswordSyncGroup) Set(val *PasswordSyncGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordSyncGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordSyncGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordSyncGroup(val *PasswordSyncGroup) *NullablePasswordSyncGroup {
	return &NullablePasswordSyncGroup{value: val, isSet: true}
}

func (v NullablePasswordSyncGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordSyncGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

