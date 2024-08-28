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

// checks if the MailFromAttributesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailFromAttributesDto{}

// MailFromAttributesDto MAIL FROM attributes for a domain / identity
type MailFromAttributesDto struct {
	// The identity or domain address
	Identity *string `json:"identity,omitempty"`
	// The new MAIL FROM domain of the identity. Must be a subdomain of the identity.
	MailFromDomain *string `json:"mailFromDomain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MailFromAttributesDto MailFromAttributesDto

// NewMailFromAttributesDto instantiates a new MailFromAttributesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailFromAttributesDto() *MailFromAttributesDto {
	this := MailFromAttributesDto{}
	return &this
}

// NewMailFromAttributesDtoWithDefaults instantiates a new MailFromAttributesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailFromAttributesDtoWithDefaults() *MailFromAttributesDto {
	this := MailFromAttributesDto{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MailFromAttributesDto) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailFromAttributesDto) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MailFromAttributesDto) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *MailFromAttributesDto) SetIdentity(v string) {
	o.Identity = &v
}

// GetMailFromDomain returns the MailFromDomain field value if set, zero value otherwise.
func (o *MailFromAttributesDto) GetMailFromDomain() string {
	if o == nil || IsNil(o.MailFromDomain) {
		var ret string
		return ret
	}
	return *o.MailFromDomain
}

// GetMailFromDomainOk returns a tuple with the MailFromDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailFromAttributesDto) GetMailFromDomainOk() (*string, bool) {
	if o == nil || IsNil(o.MailFromDomain) {
		return nil, false
	}
	return o.MailFromDomain, true
}

// HasMailFromDomain returns a boolean if a field has been set.
func (o *MailFromAttributesDto) HasMailFromDomain() bool {
	if o != nil && !IsNil(o.MailFromDomain) {
		return true
	}

	return false
}

// SetMailFromDomain gets a reference to the given string and assigns it to the MailFromDomain field.
func (o *MailFromAttributesDto) SetMailFromDomain(v string) {
	o.MailFromDomain = &v
}

func (o MailFromAttributesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailFromAttributesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.MailFromDomain) {
		toSerialize["mailFromDomain"] = o.MailFromDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MailFromAttributesDto) UnmarshalJSON(data []byte) (err error) {
	varMailFromAttributesDto := _MailFromAttributesDto{}

	err = json.Unmarshal(data, &varMailFromAttributesDto)

	if err != nil {
		return err
	}

	*o = MailFromAttributesDto(varMailFromAttributesDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "mailFromDomain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMailFromAttributesDto struct {
	value *MailFromAttributesDto
	isSet bool
}

func (v NullableMailFromAttributesDto) Get() *MailFromAttributesDto {
	return v.value
}

func (v *NullableMailFromAttributesDto) Set(val *MailFromAttributesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMailFromAttributesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMailFromAttributesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailFromAttributesDto(val *MailFromAttributesDto) *NullableMailFromAttributesDto {
	return &NullableMailFromAttributesDto{value: val, isSet: true}
}

func (v NullableMailFromAttributesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailFromAttributesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


