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

// checks if the AuthProfileSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthProfileSummary{}

// AuthProfileSummary struct for AuthProfileSummary
type AuthProfileSummary struct {
	// Tenant name.
	Tenant *string `json:"tenant,omitempty"`
	// Identity ID.
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthProfileSummary AuthProfileSummary

// NewAuthProfileSummary instantiates a new AuthProfileSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthProfileSummary() *AuthProfileSummary {
	this := AuthProfileSummary{}
	return &this
}

// NewAuthProfileSummaryWithDefaults instantiates a new AuthProfileSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthProfileSummaryWithDefaults() *AuthProfileSummary {
	this := AuthProfileSummary{}
	return &this
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AuthProfileSummary) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfileSummary) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AuthProfileSummary) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AuthProfileSummary) SetTenant(v string) {
	o.Tenant = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthProfileSummary) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProfileSummary) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthProfileSummary) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthProfileSummary) SetId(v string) {
	o.Id = &v
}

func (o AuthProfileSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthProfileSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthProfileSummary) UnmarshalJSON(data []byte) (err error) {
	varAuthProfileSummary := _AuthProfileSummary{}

	err = json.Unmarshal(data, &varAuthProfileSummary)

	if err != nil {
		return err
	}

	*o = AuthProfileSummary(varAuthProfileSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthProfileSummary struct {
	value *AuthProfileSummary
	isSet bool
}

func (v NullableAuthProfileSummary) Get() *AuthProfileSummary {
	return v.value
}

func (v *NullableAuthProfileSummary) Set(val *AuthProfileSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthProfileSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthProfileSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthProfileSummary(val *AuthProfileSummary) *NullableAuthProfileSummary {
	return &NullableAuthProfileSummary{value: val, isSet: true}
}

func (v NullableAuthProfileSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthProfileSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


