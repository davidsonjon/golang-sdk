/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the RefreshIdentitiesRequestRefreshArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshIdentitiesRequestRefreshArgs{}

// RefreshIdentitiesRequestRefreshArgs struct for RefreshIdentitiesRequestRefreshArgs
type RefreshIdentitiesRequestRefreshArgs struct {
	// This will refresh entitlement, role, and access profile calculations.
	CorrelateEntitlements *bool `json:"correlateEntitlements,omitempty"`
	// This will calculate identity attributes.
	PromoteAttributes *bool `json:"promoteAttributes,omitempty"`
	// This recalculates manager correlation and manager status. Note: This is computationally expensive to run. 
	RefreshManagerStatus *bool `json:"refreshManagerStatus,omitempty"`
	// Enables attribute synchronization.
	SynchronizeAttributes *bool `json:"synchronizeAttributes,omitempty"`
	// Option that will enable deletion of an identity objects if there are no account objects. Note: This is not typically used in IdentityNow, except by guidance from SailPoint. 
	PruneIdentities *bool `json:"pruneIdentities,omitempty"`
	// Enables provisioning of role assignments with entitlements that are not currently fulfilled.
	Provision *bool `json:"provision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RefreshIdentitiesRequestRefreshArgs RefreshIdentitiesRequestRefreshArgs

// NewRefreshIdentitiesRequestRefreshArgs instantiates a new RefreshIdentitiesRequestRefreshArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshIdentitiesRequestRefreshArgs() *RefreshIdentitiesRequestRefreshArgs {
	this := RefreshIdentitiesRequestRefreshArgs{}
	return &this
}

// NewRefreshIdentitiesRequestRefreshArgsWithDefaults instantiates a new RefreshIdentitiesRequestRefreshArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshIdentitiesRequestRefreshArgsWithDefaults() *RefreshIdentitiesRequestRefreshArgs {
	this := RefreshIdentitiesRequestRefreshArgs{}
	return &this
}

// GetCorrelateEntitlements returns the CorrelateEntitlements field value if set, zero value otherwise.
func (o *RefreshIdentitiesRequestRefreshArgs) GetCorrelateEntitlements() bool {
	if o == nil || isNil(o.CorrelateEntitlements) {
		var ret bool
		return ret
	}
	return *o.CorrelateEntitlements
}

// GetCorrelateEntitlementsOk returns a tuple with the CorrelateEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) GetCorrelateEntitlementsOk() (*bool, bool) {
	if o == nil || isNil(o.CorrelateEntitlements) {
		return nil, false
	}
	return o.CorrelateEntitlements, true
}

// HasCorrelateEntitlements returns a boolean if a field has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) HasCorrelateEntitlements() bool {
	if o != nil && !isNil(o.CorrelateEntitlements) {
		return true
	}

	return false
}

// SetCorrelateEntitlements gets a reference to the given bool and assigns it to the CorrelateEntitlements field.
func (o *RefreshIdentitiesRequestRefreshArgs) SetCorrelateEntitlements(v bool) {
	o.CorrelateEntitlements = &v
}

// GetPromoteAttributes returns the PromoteAttributes field value if set, zero value otherwise.
func (o *RefreshIdentitiesRequestRefreshArgs) GetPromoteAttributes() bool {
	if o == nil || isNil(o.PromoteAttributes) {
		var ret bool
		return ret
	}
	return *o.PromoteAttributes
}

// GetPromoteAttributesOk returns a tuple with the PromoteAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) GetPromoteAttributesOk() (*bool, bool) {
	if o == nil || isNil(o.PromoteAttributes) {
		return nil, false
	}
	return o.PromoteAttributes, true
}

// HasPromoteAttributes returns a boolean if a field has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) HasPromoteAttributes() bool {
	if o != nil && !isNil(o.PromoteAttributes) {
		return true
	}

	return false
}

// SetPromoteAttributes gets a reference to the given bool and assigns it to the PromoteAttributes field.
func (o *RefreshIdentitiesRequestRefreshArgs) SetPromoteAttributes(v bool) {
	o.PromoteAttributes = &v
}

// GetRefreshManagerStatus returns the RefreshManagerStatus field value if set, zero value otherwise.
func (o *RefreshIdentitiesRequestRefreshArgs) GetRefreshManagerStatus() bool {
	if o == nil || isNil(o.RefreshManagerStatus) {
		var ret bool
		return ret
	}
	return *o.RefreshManagerStatus
}

// GetRefreshManagerStatusOk returns a tuple with the RefreshManagerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) GetRefreshManagerStatusOk() (*bool, bool) {
	if o == nil || isNil(o.RefreshManagerStatus) {
		return nil, false
	}
	return o.RefreshManagerStatus, true
}

// HasRefreshManagerStatus returns a boolean if a field has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) HasRefreshManagerStatus() bool {
	if o != nil && !isNil(o.RefreshManagerStatus) {
		return true
	}

	return false
}

// SetRefreshManagerStatus gets a reference to the given bool and assigns it to the RefreshManagerStatus field.
func (o *RefreshIdentitiesRequestRefreshArgs) SetRefreshManagerStatus(v bool) {
	o.RefreshManagerStatus = &v
}

// GetSynchronizeAttributes returns the SynchronizeAttributes field value if set, zero value otherwise.
func (o *RefreshIdentitiesRequestRefreshArgs) GetSynchronizeAttributes() bool {
	if o == nil || isNil(o.SynchronizeAttributes) {
		var ret bool
		return ret
	}
	return *o.SynchronizeAttributes
}

// GetSynchronizeAttributesOk returns a tuple with the SynchronizeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) GetSynchronizeAttributesOk() (*bool, bool) {
	if o == nil || isNil(o.SynchronizeAttributes) {
		return nil, false
	}
	return o.SynchronizeAttributes, true
}

// HasSynchronizeAttributes returns a boolean if a field has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) HasSynchronizeAttributes() bool {
	if o != nil && !isNil(o.SynchronizeAttributes) {
		return true
	}

	return false
}

// SetSynchronizeAttributes gets a reference to the given bool and assigns it to the SynchronizeAttributes field.
func (o *RefreshIdentitiesRequestRefreshArgs) SetSynchronizeAttributes(v bool) {
	o.SynchronizeAttributes = &v
}

// GetPruneIdentities returns the PruneIdentities field value if set, zero value otherwise.
func (o *RefreshIdentitiesRequestRefreshArgs) GetPruneIdentities() bool {
	if o == nil || isNil(o.PruneIdentities) {
		var ret bool
		return ret
	}
	return *o.PruneIdentities
}

// GetPruneIdentitiesOk returns a tuple with the PruneIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) GetPruneIdentitiesOk() (*bool, bool) {
	if o == nil || isNil(o.PruneIdentities) {
		return nil, false
	}
	return o.PruneIdentities, true
}

// HasPruneIdentities returns a boolean if a field has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) HasPruneIdentities() bool {
	if o != nil && !isNil(o.PruneIdentities) {
		return true
	}

	return false
}

// SetPruneIdentities gets a reference to the given bool and assigns it to the PruneIdentities field.
func (o *RefreshIdentitiesRequestRefreshArgs) SetPruneIdentities(v bool) {
	o.PruneIdentities = &v
}

// GetProvision returns the Provision field value if set, zero value otherwise.
func (o *RefreshIdentitiesRequestRefreshArgs) GetProvision() bool {
	if o == nil || isNil(o.Provision) {
		var ret bool
		return ret
	}
	return *o.Provision
}

// GetProvisionOk returns a tuple with the Provision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) GetProvisionOk() (*bool, bool) {
	if o == nil || isNil(o.Provision) {
		return nil, false
	}
	return o.Provision, true
}

// HasProvision returns a boolean if a field has been set.
func (o *RefreshIdentitiesRequestRefreshArgs) HasProvision() bool {
	if o != nil && !isNil(o.Provision) {
		return true
	}

	return false
}

// SetProvision gets a reference to the given bool and assigns it to the Provision field.
func (o *RefreshIdentitiesRequestRefreshArgs) SetProvision(v bool) {
	o.Provision = &v
}

func (o RefreshIdentitiesRequestRefreshArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshIdentitiesRequestRefreshArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CorrelateEntitlements) {
		toSerialize["correlateEntitlements"] = o.CorrelateEntitlements
	}
	if !isNil(o.PromoteAttributes) {
		toSerialize["promoteAttributes"] = o.PromoteAttributes
	}
	if !isNil(o.RefreshManagerStatus) {
		toSerialize["refreshManagerStatus"] = o.RefreshManagerStatus
	}
	if !isNil(o.SynchronizeAttributes) {
		toSerialize["synchronizeAttributes"] = o.SynchronizeAttributes
	}
	if !isNil(o.PruneIdentities) {
		toSerialize["pruneIdentities"] = o.PruneIdentities
	}
	if !isNil(o.Provision) {
		toSerialize["provision"] = o.Provision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefreshIdentitiesRequestRefreshArgs) UnmarshalJSON(bytes []byte) (err error) {
	varRefreshIdentitiesRequestRefreshArgs := _RefreshIdentitiesRequestRefreshArgs{}

	if err = json.Unmarshal(bytes, &varRefreshIdentitiesRequestRefreshArgs); err == nil {
		*o = RefreshIdentitiesRequestRefreshArgs(varRefreshIdentitiesRequestRefreshArgs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "correlateEntitlements")
		delete(additionalProperties, "promoteAttributes")
		delete(additionalProperties, "refreshManagerStatus")
		delete(additionalProperties, "synchronizeAttributes")
		delete(additionalProperties, "pruneIdentities")
		delete(additionalProperties, "provision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefreshIdentitiesRequestRefreshArgs struct {
	value *RefreshIdentitiesRequestRefreshArgs
	isSet bool
}

func (v NullableRefreshIdentitiesRequestRefreshArgs) Get() *RefreshIdentitiesRequestRefreshArgs {
	return v.value
}

func (v *NullableRefreshIdentitiesRequestRefreshArgs) Set(val *RefreshIdentitiesRequestRefreshArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshIdentitiesRequestRefreshArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshIdentitiesRequestRefreshArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshIdentitiesRequestRefreshArgs(val *RefreshIdentitiesRequestRefreshArgs) *NullableRefreshIdentitiesRequestRefreshArgs {
	return &NullableRefreshIdentitiesRequestRefreshArgs{value: val, isSet: true}
}

func (v NullableRefreshIdentitiesRequestRefreshArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshIdentitiesRequestRefreshArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


