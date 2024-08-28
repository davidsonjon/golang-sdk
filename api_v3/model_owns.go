/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the Owns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Owns{}

// Owns struct for Owns
type Owns struct {
	Sources []Reference `json:"sources,omitempty"`
	Entitlements []Reference `json:"entitlements,omitempty"`
	AccessProfiles []Reference `json:"accessProfiles,omitempty"`
	Roles []Reference `json:"roles,omitempty"`
	Apps []Reference `json:"apps,omitempty"`
	GovernanceGroups []Reference `json:"governanceGroups,omitempty"`
	FallbackApprover *bool `json:"fallbackApprover,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Owns Owns

// NewOwns instantiates a new Owns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwns() *Owns {
	this := Owns{}
	return &this
}

// NewOwnsWithDefaults instantiates a new Owns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnsWithDefaults() *Owns {
	this := Owns{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *Owns) GetSources() []Reference {
	if o == nil || IsNil(o.Sources) {
		var ret []Reference
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetSourcesOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *Owns) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []Reference and assigns it to the Sources field.
func (o *Owns) SetSources(v []Reference) {
	o.Sources = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *Owns) GetEntitlements() []Reference {
	if o == nil || IsNil(o.Entitlements) {
		var ret []Reference
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetEntitlementsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *Owns) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []Reference and assigns it to the Entitlements field.
func (o *Owns) SetEntitlements(v []Reference) {
	o.Entitlements = v
}

// GetAccessProfiles returns the AccessProfiles field value if set, zero value otherwise.
func (o *Owns) GetAccessProfiles() []Reference {
	if o == nil || IsNil(o.AccessProfiles) {
		var ret []Reference
		return ret
	}
	return o.AccessProfiles
}

// GetAccessProfilesOk returns a tuple with the AccessProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetAccessProfilesOk() ([]Reference, bool) {
	if o == nil || IsNil(o.AccessProfiles) {
		return nil, false
	}
	return o.AccessProfiles, true
}

// HasAccessProfiles returns a boolean if a field has been set.
func (o *Owns) HasAccessProfiles() bool {
	if o != nil && !IsNil(o.AccessProfiles) {
		return true
	}

	return false
}

// SetAccessProfiles gets a reference to the given []Reference and assigns it to the AccessProfiles field.
func (o *Owns) SetAccessProfiles(v []Reference) {
	o.AccessProfiles = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Owns) GetRoles() []Reference {
	if o == nil || IsNil(o.Roles) {
		var ret []Reference
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetRolesOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Owns) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Reference and assigns it to the Roles field.
func (o *Owns) SetRoles(v []Reference) {
	o.Roles = v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *Owns) GetApps() []Reference {
	if o == nil || IsNil(o.Apps) {
		var ret []Reference
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetAppsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *Owns) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []Reference and assigns it to the Apps field.
func (o *Owns) SetApps(v []Reference) {
	o.Apps = v
}

// GetGovernanceGroups returns the GovernanceGroups field value if set, zero value otherwise.
func (o *Owns) GetGovernanceGroups() []Reference {
	if o == nil || IsNil(o.GovernanceGroups) {
		var ret []Reference
		return ret
	}
	return o.GovernanceGroups
}

// GetGovernanceGroupsOk returns a tuple with the GovernanceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetGovernanceGroupsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.GovernanceGroups) {
		return nil, false
	}
	return o.GovernanceGroups, true
}

// HasGovernanceGroups returns a boolean if a field has been set.
func (o *Owns) HasGovernanceGroups() bool {
	if o != nil && !IsNil(o.GovernanceGroups) {
		return true
	}

	return false
}

// SetGovernanceGroups gets a reference to the given []Reference and assigns it to the GovernanceGroups field.
func (o *Owns) SetGovernanceGroups(v []Reference) {
	o.GovernanceGroups = v
}

// GetFallbackApprover returns the FallbackApprover field value if set, zero value otherwise.
func (o *Owns) GetFallbackApprover() bool {
	if o == nil || IsNil(o.FallbackApprover) {
		var ret bool
		return ret
	}
	return *o.FallbackApprover
}

// GetFallbackApproverOk returns a tuple with the FallbackApprover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owns) GetFallbackApproverOk() (*bool, bool) {
	if o == nil || IsNil(o.FallbackApprover) {
		return nil, false
	}
	return o.FallbackApprover, true
}

// HasFallbackApprover returns a boolean if a field has been set.
func (o *Owns) HasFallbackApprover() bool {
	if o != nil && !IsNil(o.FallbackApprover) {
		return true
	}

	return false
}

// SetFallbackApprover gets a reference to the given bool and assigns it to the FallbackApprover field.
func (o *Owns) SetFallbackApprover(v bool) {
	o.FallbackApprover = &v
}

func (o Owns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Owns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.AccessProfiles) {
		toSerialize["accessProfiles"] = o.AccessProfiles
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.GovernanceGroups) {
		toSerialize["governanceGroups"] = o.GovernanceGroups
	}
	if !IsNil(o.FallbackApprover) {
		toSerialize["fallbackApprover"] = o.FallbackApprover
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Owns) UnmarshalJSON(data []byte) (err error) {
	varOwns := _Owns{}

	err = json.Unmarshal(data, &varOwns)

	if err != nil {
		return err
	}

	*o = Owns(varOwns)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sources")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "accessProfiles")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "governanceGroups")
		delete(additionalProperties, "fallbackApprover")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOwns struct {
	value *Owns
	isSet bool
}

func (v NullableOwns) Get() *Owns {
	return v.value
}

func (v *NullableOwns) Set(val *Owns) {
	v.value = val
	v.isSet = true
}

func (v NullableOwns) IsSet() bool {
	return v.isSet
}

func (v *NullableOwns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwns(val *Owns) *NullableOwns {
	return &NullableOwns{value: val, isSet: true}
}

func (v NullableOwns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


