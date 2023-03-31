/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the RoleInsightsInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleInsightsInsight{}

// RoleInsightsInsight struct for RoleInsightsInsight
type RoleInsightsInsight struct {
	// The number of identities in this role with the entitlement.
	Type *string `json:"type,omitempty"`
	// The number of identities in this role with the entitlement.
	IdentitiesWithAccess *int32 `json:"identitiesWithAccess,omitempty"`
	// The number of identities in this role that do not have the specified entitlement.
	IdentitiesImpacted *int32 `json:"identitiesImpacted,omitempty"`
	// The total number of identities.
	TotalNumberOfIdentities *int32 `json:"totalNumberOfIdentities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleInsightsInsight RoleInsightsInsight

// NewRoleInsightsInsight instantiates a new RoleInsightsInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleInsightsInsight() *RoleInsightsInsight {
	this := RoleInsightsInsight{}
	return &this
}

// NewRoleInsightsInsightWithDefaults instantiates a new RoleInsightsInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleInsightsInsightWithDefaults() *RoleInsightsInsight {
	this := RoleInsightsInsight{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleInsightsInsight) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsInsight) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleInsightsInsight) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RoleInsightsInsight) SetType(v string) {
	o.Type = &v
}

// GetIdentitiesWithAccess returns the IdentitiesWithAccess field value if set, zero value otherwise.
func (o *RoleInsightsInsight) GetIdentitiesWithAccess() int32 {
	if o == nil || isNil(o.IdentitiesWithAccess) {
		var ret int32
		return ret
	}
	return *o.IdentitiesWithAccess
}

// GetIdentitiesWithAccessOk returns a tuple with the IdentitiesWithAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsInsight) GetIdentitiesWithAccessOk() (*int32, bool) {
	if o == nil || isNil(o.IdentitiesWithAccess) {
		return nil, false
	}
	return o.IdentitiesWithAccess, true
}

// HasIdentitiesWithAccess returns a boolean if a field has been set.
func (o *RoleInsightsInsight) HasIdentitiesWithAccess() bool {
	if o != nil && !isNil(o.IdentitiesWithAccess) {
		return true
	}

	return false
}

// SetIdentitiesWithAccess gets a reference to the given int32 and assigns it to the IdentitiesWithAccess field.
func (o *RoleInsightsInsight) SetIdentitiesWithAccess(v int32) {
	o.IdentitiesWithAccess = &v
}

// GetIdentitiesImpacted returns the IdentitiesImpacted field value if set, zero value otherwise.
func (o *RoleInsightsInsight) GetIdentitiesImpacted() int32 {
	if o == nil || isNil(o.IdentitiesImpacted) {
		var ret int32
		return ret
	}
	return *o.IdentitiesImpacted
}

// GetIdentitiesImpactedOk returns a tuple with the IdentitiesImpacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsInsight) GetIdentitiesImpactedOk() (*int32, bool) {
	if o == nil || isNil(o.IdentitiesImpacted) {
		return nil, false
	}
	return o.IdentitiesImpacted, true
}

// HasIdentitiesImpacted returns a boolean if a field has been set.
func (o *RoleInsightsInsight) HasIdentitiesImpacted() bool {
	if o != nil && !isNil(o.IdentitiesImpacted) {
		return true
	}

	return false
}

// SetIdentitiesImpacted gets a reference to the given int32 and assigns it to the IdentitiesImpacted field.
func (o *RoleInsightsInsight) SetIdentitiesImpacted(v int32) {
	o.IdentitiesImpacted = &v
}

// GetTotalNumberOfIdentities returns the TotalNumberOfIdentities field value if set, zero value otherwise.
func (o *RoleInsightsInsight) GetTotalNumberOfIdentities() int32 {
	if o == nil || isNil(o.TotalNumberOfIdentities) {
		var ret int32
		return ret
	}
	return *o.TotalNumberOfIdentities
}

// GetTotalNumberOfIdentitiesOk returns a tuple with the TotalNumberOfIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsInsight) GetTotalNumberOfIdentitiesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalNumberOfIdentities) {
		return nil, false
	}
	return o.TotalNumberOfIdentities, true
}

// HasTotalNumberOfIdentities returns a boolean if a field has been set.
func (o *RoleInsightsInsight) HasTotalNumberOfIdentities() bool {
	if o != nil && !isNil(o.TotalNumberOfIdentities) {
		return true
	}

	return false
}

// SetTotalNumberOfIdentities gets a reference to the given int32 and assigns it to the TotalNumberOfIdentities field.
func (o *RoleInsightsInsight) SetTotalNumberOfIdentities(v int32) {
	o.TotalNumberOfIdentities = &v
}

func (o RoleInsightsInsight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleInsightsInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.IdentitiesWithAccess) {
		toSerialize["identitiesWithAccess"] = o.IdentitiesWithAccess
	}
	if !isNil(o.IdentitiesImpacted) {
		toSerialize["identitiesImpacted"] = o.IdentitiesImpacted
	}
	if !isNil(o.TotalNumberOfIdentities) {
		toSerialize["totalNumberOfIdentities"] = o.TotalNumberOfIdentities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleInsightsInsight) UnmarshalJSON(bytes []byte) (err error) {
	varRoleInsightsInsight := _RoleInsightsInsight{}

	if err = json.Unmarshal(bytes, &varRoleInsightsInsight); err == nil {
		*o = RoleInsightsInsight(varRoleInsightsInsight)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "identitiesWithAccess")
		delete(additionalProperties, "identitiesImpacted")
		delete(additionalProperties, "totalNumberOfIdentities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleInsightsInsight struct {
	value *RoleInsightsInsight
	isSet bool
}

func (v NullableRoleInsightsInsight) Get() *RoleInsightsInsight {
	return v.value
}

func (v *NullableRoleInsightsInsight) Set(val *RoleInsightsInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleInsightsInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleInsightsInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleInsightsInsight(val *RoleInsightsInsight) *NullableRoleInsightsInsight {
	return &NullableRoleInsightsInsight{value: val, isSet: true}
}

func (v NullableRoleInsightsInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleInsightsInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


