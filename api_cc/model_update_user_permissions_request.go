/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_cc

import (
	"encoding/json"
)

// checks if the UpdateUserPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserPermissionsRequest{}

// UpdateUserPermissionsRequest struct for UpdateUserPermissionsRequest
type UpdateUserPermissionsRequest struct {
	Ids *string `json:"ids,omitempty"`
	// Indicates if user should be an IDN Admin.  \"0\" for false, \"1\" for true.
	IsAdmin *string `json:"isAdmin,omitempty"`
	AdminType *string `json:"adminType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserPermissionsRequest UpdateUserPermissionsRequest

// NewUpdateUserPermissionsRequest instantiates a new UpdateUserPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserPermissionsRequest() *UpdateUserPermissionsRequest {
	this := UpdateUserPermissionsRequest{}
	return &this
}

// NewUpdateUserPermissionsRequestWithDefaults instantiates a new UpdateUserPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserPermissionsRequestWithDefaults() *UpdateUserPermissionsRequest {
	this := UpdateUserPermissionsRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *UpdateUserPermissionsRequest) GetIds() string {
	if o == nil || isNil(o.Ids) {
		var ret string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPermissionsRequest) GetIdsOk() (*string, bool) {
	if o == nil || isNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *UpdateUserPermissionsRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given string and assigns it to the Ids field.
func (o *UpdateUserPermissionsRequest) SetIds(v string) {
	o.Ids = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *UpdateUserPermissionsRequest) GetIsAdmin() string {
	if o == nil || isNil(o.IsAdmin) {
		var ret string
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPermissionsRequest) GetIsAdminOk() (*string, bool) {
	if o == nil || isNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *UpdateUserPermissionsRequest) HasIsAdmin() bool {
	if o != nil && !isNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given string and assigns it to the IsAdmin field.
func (o *UpdateUserPermissionsRequest) SetIsAdmin(v string) {
	o.IsAdmin = &v
}

// GetAdminType returns the AdminType field value if set, zero value otherwise.
func (o *UpdateUserPermissionsRequest) GetAdminType() string {
	if o == nil || isNil(o.AdminType) {
		var ret string
		return ret
	}
	return *o.AdminType
}

// GetAdminTypeOk returns a tuple with the AdminType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPermissionsRequest) GetAdminTypeOk() (*string, bool) {
	if o == nil || isNil(o.AdminType) {
		return nil, false
	}
	return o.AdminType, true
}

// HasAdminType returns a boolean if a field has been set.
func (o *UpdateUserPermissionsRequest) HasAdminType() bool {
	if o != nil && !isNil(o.AdminType) {
		return true
	}

	return false
}

// SetAdminType gets a reference to the given string and assigns it to the AdminType field.
func (o *UpdateUserPermissionsRequest) SetAdminType(v string) {
	o.AdminType = &v
}

func (o UpdateUserPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if !isNil(o.AdminType) {
		toSerialize["adminType"] = o.AdminType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateUserPermissionsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateUserPermissionsRequest := _UpdateUserPermissionsRequest{}

	if err = json.Unmarshal(bytes, &varUpdateUserPermissionsRequest); err == nil {
	*o = UpdateUserPermissionsRequest(varUpdateUserPermissionsRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "isAdmin")
		delete(additionalProperties, "adminType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserPermissionsRequest struct {
	value *UpdateUserPermissionsRequest
	isSet bool
}

func (v NullableUpdateUserPermissionsRequest) Get() *UpdateUserPermissionsRequest {
	return v.value
}

func (v *NullableUpdateUserPermissionsRequest) Set(val *UpdateUserPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserPermissionsRequest(val *UpdateUserPermissionsRequest) *NullableUpdateUserPermissionsRequest {
	return &NullableUpdateUserPermissionsRequest{value: val, isSet: true}
}

func (v NullableUpdateUserPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


