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

// checks if the RoleMiningSessionParametersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningSessionParametersDto{}

// RoleMiningSessionParametersDto struct for RoleMiningSessionParametersDto
type RoleMiningSessionParametersDto struct {
	// The ID of the role mining session
	Id *string `json:"id,omitempty"`
	// The session's saved name
	Name NullableString `json:"name,omitempty"`
	// Minimum number of identities in a potential role
	MinNumIdentitiesInPotentialRole NullableInt32 `json:"minNumIdentitiesInPotentialRole,omitempty"`
	// The prune threshold to be used or null to calculate prescribedPruneThreshold
	PruneThreshold NullableInt32 `json:"pruneThreshold,omitempty"`
	// The session's saved status
	Saved *bool `json:"saved,omitempty"`
	Scope *RoleMiningSessionScope `json:"scope,omitempty"`
	Type *RoleMiningRoleType `json:"type,omitempty"`
	State *RoleMiningSessionState `json:"state,omitempty"`
	ScopingMethod *RoleMiningSessionScopingMethod `json:"scopingMethod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningSessionParametersDto RoleMiningSessionParametersDto

// NewRoleMiningSessionParametersDto instantiates a new RoleMiningSessionParametersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningSessionParametersDto() *RoleMiningSessionParametersDto {
	this := RoleMiningSessionParametersDto{}
	var saved bool = true
	this.Saved = &saved
	return &this
}

// NewRoleMiningSessionParametersDtoWithDefaults instantiates a new RoleMiningSessionParametersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningSessionParametersDtoWithDefaults() *RoleMiningSessionParametersDto {
	this := RoleMiningSessionParametersDto{}
	var saved bool = true
	this.Saved = &saved
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMiningSessionParametersDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionParametersDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMiningSessionParametersDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningSessionParametersDto) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningSessionParametersDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RoleMiningSessionParametersDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RoleMiningSessionParametersDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RoleMiningSessionParametersDto) UnsetName() {
	o.Name.Unset()
}

// GetMinNumIdentitiesInPotentialRole returns the MinNumIdentitiesInPotentialRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningSessionParametersDto) GetMinNumIdentitiesInPotentialRole() int32 {
	if o == nil || isNil(o.MinNumIdentitiesInPotentialRole.Get()) {
		var ret int32
		return ret
	}
	return *o.MinNumIdentitiesInPotentialRole.Get()
}

// GetMinNumIdentitiesInPotentialRoleOk returns a tuple with the MinNumIdentitiesInPotentialRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningSessionParametersDto) GetMinNumIdentitiesInPotentialRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinNumIdentitiesInPotentialRole.Get(), o.MinNumIdentitiesInPotentialRole.IsSet()
}

// HasMinNumIdentitiesInPotentialRole returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasMinNumIdentitiesInPotentialRole() bool {
	if o != nil && o.MinNumIdentitiesInPotentialRole.IsSet() {
		return true
	}

	return false
}

// SetMinNumIdentitiesInPotentialRole gets a reference to the given NullableInt32 and assigns it to the MinNumIdentitiesInPotentialRole field.
func (o *RoleMiningSessionParametersDto) SetMinNumIdentitiesInPotentialRole(v int32) {
	o.MinNumIdentitiesInPotentialRole.Set(&v)
}
// SetMinNumIdentitiesInPotentialRoleNil sets the value for MinNumIdentitiesInPotentialRole to be an explicit nil
func (o *RoleMiningSessionParametersDto) SetMinNumIdentitiesInPotentialRoleNil() {
	o.MinNumIdentitiesInPotentialRole.Set(nil)
}

// UnsetMinNumIdentitiesInPotentialRole ensures that no value is present for MinNumIdentitiesInPotentialRole, not even an explicit nil
func (o *RoleMiningSessionParametersDto) UnsetMinNumIdentitiesInPotentialRole() {
	o.MinNumIdentitiesInPotentialRole.Unset()
}

// GetPruneThreshold returns the PruneThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMiningSessionParametersDto) GetPruneThreshold() int32 {
	if o == nil || isNil(o.PruneThreshold.Get()) {
		var ret int32
		return ret
	}
	return *o.PruneThreshold.Get()
}

// GetPruneThresholdOk returns a tuple with the PruneThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMiningSessionParametersDto) GetPruneThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PruneThreshold.Get(), o.PruneThreshold.IsSet()
}

// HasPruneThreshold returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasPruneThreshold() bool {
	if o != nil && o.PruneThreshold.IsSet() {
		return true
	}

	return false
}

// SetPruneThreshold gets a reference to the given NullableInt32 and assigns it to the PruneThreshold field.
func (o *RoleMiningSessionParametersDto) SetPruneThreshold(v int32) {
	o.PruneThreshold.Set(&v)
}
// SetPruneThresholdNil sets the value for PruneThreshold to be an explicit nil
func (o *RoleMiningSessionParametersDto) SetPruneThresholdNil() {
	o.PruneThreshold.Set(nil)
}

// UnsetPruneThreshold ensures that no value is present for PruneThreshold, not even an explicit nil
func (o *RoleMiningSessionParametersDto) UnsetPruneThreshold() {
	o.PruneThreshold.Unset()
}

// GetSaved returns the Saved field value if set, zero value otherwise.
func (o *RoleMiningSessionParametersDto) GetSaved() bool {
	if o == nil || isNil(o.Saved) {
		var ret bool
		return ret
	}
	return *o.Saved
}

// GetSavedOk returns a tuple with the Saved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionParametersDto) GetSavedOk() (*bool, bool) {
	if o == nil || isNil(o.Saved) {
		return nil, false
	}
	return o.Saved, true
}

// HasSaved returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasSaved() bool {
	if o != nil && !isNil(o.Saved) {
		return true
	}

	return false
}

// SetSaved gets a reference to the given bool and assigns it to the Saved field.
func (o *RoleMiningSessionParametersDto) SetSaved(v bool) {
	o.Saved = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *RoleMiningSessionParametersDto) GetScope() RoleMiningSessionScope {
	if o == nil || isNil(o.Scope) {
		var ret RoleMiningSessionScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionParametersDto) GetScopeOk() (*RoleMiningSessionScope, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given RoleMiningSessionScope and assigns it to the Scope field.
func (o *RoleMiningSessionParametersDto) SetScope(v RoleMiningSessionScope) {
	o.Scope = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleMiningSessionParametersDto) GetType() RoleMiningRoleType {
	if o == nil || isNil(o.Type) {
		var ret RoleMiningRoleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionParametersDto) GetTypeOk() (*RoleMiningRoleType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RoleMiningRoleType and assigns it to the Type field.
func (o *RoleMiningSessionParametersDto) SetType(v RoleMiningRoleType) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RoleMiningSessionParametersDto) GetState() RoleMiningSessionState {
	if o == nil || isNil(o.State) {
		var ret RoleMiningSessionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionParametersDto) GetStateOk() (*RoleMiningSessionState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given RoleMiningSessionState and assigns it to the State field.
func (o *RoleMiningSessionParametersDto) SetState(v RoleMiningSessionState) {
	o.State = &v
}

// GetScopingMethod returns the ScopingMethod field value if set, zero value otherwise.
func (o *RoleMiningSessionParametersDto) GetScopingMethod() RoleMiningSessionScopingMethod {
	if o == nil || isNil(o.ScopingMethod) {
		var ret RoleMiningSessionScopingMethod
		return ret
	}
	return *o.ScopingMethod
}

// GetScopingMethodOk returns a tuple with the ScopingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionParametersDto) GetScopingMethodOk() (*RoleMiningSessionScopingMethod, bool) {
	if o == nil || isNil(o.ScopingMethod) {
		return nil, false
	}
	return o.ScopingMethod, true
}

// HasScopingMethod returns a boolean if a field has been set.
func (o *RoleMiningSessionParametersDto) HasScopingMethod() bool {
	if o != nil && !isNil(o.ScopingMethod) {
		return true
	}

	return false
}

// SetScopingMethod gets a reference to the given RoleMiningSessionScopingMethod and assigns it to the ScopingMethod field.
func (o *RoleMiningSessionParametersDto) SetScopingMethod(v RoleMiningSessionScopingMethod) {
	o.ScopingMethod = &v
}

func (o RoleMiningSessionParametersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningSessionParametersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.MinNumIdentitiesInPotentialRole.IsSet() {
		toSerialize["minNumIdentitiesInPotentialRole"] = o.MinNumIdentitiesInPotentialRole.Get()
	}
	if o.PruneThreshold.IsSet() {
		toSerialize["pruneThreshold"] = o.PruneThreshold.Get()
	}
	if !isNil(o.Saved) {
		toSerialize["saved"] = o.Saved
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.ScopingMethod) {
		toSerialize["scopingMethod"] = o.ScopingMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningSessionParametersDto) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningSessionParametersDto := _RoleMiningSessionParametersDto{}

	if err = json.Unmarshal(bytes, &varRoleMiningSessionParametersDto); err == nil {
		*o = RoleMiningSessionParametersDto(varRoleMiningSessionParametersDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "minNumIdentitiesInPotentialRole")
		delete(additionalProperties, "pruneThreshold")
		delete(additionalProperties, "saved")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "type")
		delete(additionalProperties, "state")
		delete(additionalProperties, "scopingMethod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningSessionParametersDto struct {
	value *RoleMiningSessionParametersDto
	isSet bool
}

func (v NullableRoleMiningSessionParametersDto) Get() *RoleMiningSessionParametersDto {
	return v.value
}

func (v *NullableRoleMiningSessionParametersDto) Set(val *RoleMiningSessionParametersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningSessionParametersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningSessionParametersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningSessionParametersDto(val *RoleMiningSessionParametersDto) *NullableRoleMiningSessionParametersDto {
	return &NullableRoleMiningSessionParametersDto{value: val, isSet: true}
}

func (v NullableRoleMiningSessionParametersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningSessionParametersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


