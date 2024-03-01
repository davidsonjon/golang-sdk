/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
)

// checks if the ReviewableRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewableRole{}

// ReviewableRole struct for ReviewableRole
type ReviewableRole struct {
	// The id for the Role
	Id *string `json:"id,omitempty"`
	// The name of the Role
	Name *string `json:"name,omitempty"`
	// Information about the Role
	Description *string `json:"description,omitempty"`
	// Indicates if the entitlement is a privileged entitlement
	Privileged *bool `json:"privileged,omitempty"`
	Owner NullableIdentityReferenceWithNameAndEmail `json:"owner,omitempty"`
	// Indicates whether the Role can be revoked or requested
	Revocable *bool `json:"revocable,omitempty"`
	// The date when a user's access expires.
	EndDate *time.Time `json:"endDate,omitempty"`
	// The list of Access Profiles associated with this Role
	AccessProfiles []ReviewableAccessProfile `json:"accessProfiles,omitempty"`
	// The list of entitlements associated with this Role
	Entitlements []ReviewableEntitlement `json:"entitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewableRole ReviewableRole

// NewReviewableRole instantiates a new ReviewableRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewableRole() *ReviewableRole {
	this := ReviewableRole{}
	return &this
}

// NewReviewableRoleWithDefaults instantiates a new ReviewableRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewableRoleWithDefaults() *ReviewableRole {
	this := ReviewableRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReviewableRole) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReviewableRole) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReviewableRole) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewableRole) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewableRole) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewableRole) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReviewableRole) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReviewableRole) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReviewableRole) SetDescription(v string) {
	o.Description = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *ReviewableRole) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *ReviewableRole) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *ReviewableRole) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableRole) GetOwner() IdentityReferenceWithNameAndEmail {
	if o == nil || isNil(o.Owner.Get()) {
		var ret IdentityReferenceWithNameAndEmail
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableRole) GetOwnerOk() (*IdentityReferenceWithNameAndEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ReviewableRole) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableIdentityReferenceWithNameAndEmail and assigns it to the Owner field.
func (o *ReviewableRole) SetOwner(v IdentityReferenceWithNameAndEmail) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ReviewableRole) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ReviewableRole) UnsetOwner() {
	o.Owner.Unset()
}

// GetRevocable returns the Revocable field value if set, zero value otherwise.
func (o *ReviewableRole) GetRevocable() bool {
	if o == nil || isNil(o.Revocable) {
		var ret bool
		return ret
	}
	return *o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetRevocableOk() (*bool, bool) {
	if o == nil || isNil(o.Revocable) {
		return nil, false
	}
	return o.Revocable, true
}

// HasRevocable returns a boolean if a field has been set.
func (o *ReviewableRole) HasRevocable() bool {
	if o != nil && !isNil(o.Revocable) {
		return true
	}

	return false
}

// SetRevocable gets a reference to the given bool and assigns it to the Revocable field.
func (o *ReviewableRole) SetRevocable(v bool) {
	o.Revocable = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ReviewableRole) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetEndDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ReviewableRole) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ReviewableRole) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetAccessProfiles returns the AccessProfiles field value if set, zero value otherwise.
func (o *ReviewableRole) GetAccessProfiles() []ReviewableAccessProfile {
	if o == nil || isNil(o.AccessProfiles) {
		var ret []ReviewableAccessProfile
		return ret
	}
	return o.AccessProfiles
}

// GetAccessProfilesOk returns a tuple with the AccessProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetAccessProfilesOk() ([]ReviewableAccessProfile, bool) {
	if o == nil || isNil(o.AccessProfiles) {
		return nil, false
	}
	return o.AccessProfiles, true
}

// HasAccessProfiles returns a boolean if a field has been set.
func (o *ReviewableRole) HasAccessProfiles() bool {
	if o != nil && !isNil(o.AccessProfiles) {
		return true
	}

	return false
}

// SetAccessProfiles gets a reference to the given []ReviewableAccessProfile and assigns it to the AccessProfiles field.
func (o *ReviewableRole) SetAccessProfiles(v []ReviewableAccessProfile) {
	o.AccessProfiles = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *ReviewableRole) GetEntitlements() []ReviewableEntitlement {
	if o == nil || isNil(o.Entitlements) {
		var ret []ReviewableEntitlement
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableRole) GetEntitlementsOk() ([]ReviewableEntitlement, bool) {
	if o == nil || isNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *ReviewableRole) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []ReviewableEntitlement and assigns it to the Entitlements field.
func (o *ReviewableRole) SetEntitlements(v []ReviewableEntitlement) {
	o.Entitlements = v
}

func (o ReviewableRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewableRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if !isNil(o.Revocable) {
		toSerialize["revocable"] = o.Revocable
	}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !isNil(o.AccessProfiles) {
		toSerialize["accessProfiles"] = o.AccessProfiles
	}
	if !isNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewableRole) UnmarshalJSON(bytes []byte) (err error) {
	varReviewableRole := _ReviewableRole{}

	if err = json.Unmarshal(bytes, &varReviewableRole); err == nil {
	*o = ReviewableRole(varReviewableRole)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "revocable")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "accessProfiles")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewableRole struct {
	value *ReviewableRole
	isSet bool
}

func (v NullableReviewableRole) Get() *ReviewableRole {
	return v.value
}

func (v *NullableReviewableRole) Set(val *ReviewableRole) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewableRole(val *ReviewableRole) *NullableReviewableRole {
	return &NullableReviewableRole{value: val, isSet: true}
}

func (v NullableReviewableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

