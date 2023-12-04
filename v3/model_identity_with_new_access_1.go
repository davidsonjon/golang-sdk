/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the IdentityWithNewAccess1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityWithNewAccess1{}

// IdentityWithNewAccess1 An identity with a set of access to be added
type IdentityWithNewAccess1 struct {
	// Set of identity IDs to be checked.
	IdentityId string `json:"identityId"`
	// The bundle of access profiles to be added to the identities specified. All references must be ENTITLEMENT type.
	AccessRefs []EntitlementRef `json:"accessRefs"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on completion of the violation check.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityWithNewAccess1 IdentityWithNewAccess1

// NewIdentityWithNewAccess1 instantiates a new IdentityWithNewAccess1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithNewAccess1(identityId string, accessRefs []EntitlementRef) *IdentityWithNewAccess1 {
	this := IdentityWithNewAccess1{}
	this.IdentityId = identityId
	this.AccessRefs = accessRefs
	return &this
}

// NewIdentityWithNewAccess1WithDefaults instantiates a new IdentityWithNewAccess1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithNewAccess1WithDefaults() *IdentityWithNewAccess1 {
	this := IdentityWithNewAccess1{}
	return &this
}

// GetIdentityId returns the IdentityId field value
func (o *IdentityWithNewAccess1) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *IdentityWithNewAccess1) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *IdentityWithNewAccess1) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetAccessRefs returns the AccessRefs field value
func (o *IdentityWithNewAccess1) GetAccessRefs() []EntitlementRef {
	if o == nil {
		var ret []EntitlementRef
		return ret
	}

	return o.AccessRefs
}

// GetAccessRefsOk returns a tuple with the AccessRefs field value
// and a boolean to check if the value has been set.
func (o *IdentityWithNewAccess1) GetAccessRefsOk() ([]EntitlementRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessRefs, true
}

// SetAccessRefs sets field value
func (o *IdentityWithNewAccess1) SetAccessRefs(v []EntitlementRef) {
	o.AccessRefs = v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *IdentityWithNewAccess1) GetClientMetadata() map[string]string {
	if o == nil || isNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithNewAccess1) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *IdentityWithNewAccess1) HasClientMetadata() bool {
	if o != nil && !isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *IdentityWithNewAccess1) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

func (o IdentityWithNewAccess1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityWithNewAccess1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identityId"] = o.IdentityId
	toSerialize["accessRefs"] = o.AccessRefs
	if !isNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityWithNewAccess1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identityId",
		"accessRefs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIdentityWithNewAccess1 := _IdentityWithNewAccess1{}

	if err = json.Unmarshal(bytes, &varIdentityWithNewAccess1); err == nil {
	*o = IdentityWithNewAccess1(varIdentityWithNewAccess1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "accessRefs")
		delete(additionalProperties, "clientMetadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityWithNewAccess1 struct {
	value *IdentityWithNewAccess1
	isSet bool
}

func (v NullableIdentityWithNewAccess1) Get() *IdentityWithNewAccess1 {
	return v.value
}

func (v *NullableIdentityWithNewAccess1) Set(val *IdentityWithNewAccess1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithNewAccess1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithNewAccess1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithNewAccess1(val *IdentityWithNewAccess1) *NullableIdentityWithNewAccess1 {
	return &NullableIdentityWithNewAccess1{value: val, isSet: true}
}

func (v NullableIdentityWithNewAccess1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithNewAccess1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


