/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the AccountAggregationCompletedStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAggregationCompletedStats{}

// AccountAggregationCompletedStats Overall statistics about the account aggregation.
type AccountAggregationCompletedStats struct {
	// The number of accounts which were scanned / iterated over.
	Scanned int32 `json:"scanned"`
	// The number of accounts which existed before, but had no changes.
	Unchanged int32 `json:"unchanged"`
	// The number of accounts which existed before, but had changes.
	Changed int32 `json:"changed"`
	// The number of accounts which are new - have not existed before.
	Added int32 `json:"added"`
	// The number accounts which existed before, but no longer exist (thus getting removed).
	Removed int32 `json:"removed"`
	AdditionalProperties map[string]interface{}
}

type _AccountAggregationCompletedStats AccountAggregationCompletedStats

// NewAccountAggregationCompletedStats instantiates a new AccountAggregationCompletedStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAggregationCompletedStats(scanned int32, unchanged int32, changed int32, added int32, removed int32) *AccountAggregationCompletedStats {
	this := AccountAggregationCompletedStats{}
	this.Scanned = scanned
	this.Unchanged = unchanged
	this.Changed = changed
	this.Added = added
	this.Removed = removed
	return &this
}

// NewAccountAggregationCompletedStatsWithDefaults instantiates a new AccountAggregationCompletedStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAggregationCompletedStatsWithDefaults() *AccountAggregationCompletedStats {
	this := AccountAggregationCompletedStats{}
	return &this
}

// GetScanned returns the Scanned field value
func (o *AccountAggregationCompletedStats) GetScanned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scanned
}

// GetScannedOk returns a tuple with the Scanned field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedStats) GetScannedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scanned, true
}

// SetScanned sets field value
func (o *AccountAggregationCompletedStats) SetScanned(v int32) {
	o.Scanned = v
}

// GetUnchanged returns the Unchanged field value
func (o *AccountAggregationCompletedStats) GetUnchanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Unchanged
}

// GetUnchangedOk returns a tuple with the Unchanged field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedStats) GetUnchangedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unchanged, true
}

// SetUnchanged sets field value
func (o *AccountAggregationCompletedStats) SetUnchanged(v int32) {
	o.Unchanged = v
}

// GetChanged returns the Changed field value
func (o *AccountAggregationCompletedStats) GetChanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Changed
}

// GetChangedOk returns a tuple with the Changed field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedStats) GetChangedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Changed, true
}

// SetChanged sets field value
func (o *AccountAggregationCompletedStats) SetChanged(v int32) {
	o.Changed = v
}

// GetAdded returns the Added field value
func (o *AccountAggregationCompletedStats) GetAdded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedStats) GetAddedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *AccountAggregationCompletedStats) SetAdded(v int32) {
	o.Added = v
}

// GetRemoved returns the Removed field value
func (o *AccountAggregationCompletedStats) GetRemoved() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompletedStats) GetRemovedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *AccountAggregationCompletedStats) SetRemoved(v int32) {
	o.Removed = v
}

func (o AccountAggregationCompletedStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAggregationCompletedStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scanned"] = o.Scanned
	toSerialize["unchanged"] = o.Unchanged
	toSerialize["changed"] = o.Changed
	toSerialize["added"] = o.Added
	toSerialize["removed"] = o.Removed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountAggregationCompletedStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scanned",
		"unchanged",
		"changed",
		"added",
		"removed",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountAggregationCompletedStats := _AccountAggregationCompletedStats{}

	err = json.Unmarshal(data, &varAccountAggregationCompletedStats)

	if err != nil {
		return err
	}

	*o = AccountAggregationCompletedStats(varAccountAggregationCompletedStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scanned")
		delete(additionalProperties, "unchanged")
		delete(additionalProperties, "changed")
		delete(additionalProperties, "added")
		delete(additionalProperties, "removed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAggregationCompletedStats struct {
	value *AccountAggregationCompletedStats
	isSet bool
}

func (v NullableAccountAggregationCompletedStats) Get() *AccountAggregationCompletedStats {
	return v.value
}

func (v *NullableAccountAggregationCompletedStats) Set(val *AccountAggregationCompletedStats) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAggregationCompletedStats) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAggregationCompletedStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAggregationCompletedStats(val *AccountAggregationCompletedStats) *NullableAccountAggregationCompletedStats {
	return &NullableAccountAggregationCompletedStats{value: val, isSet: true}
}

func (v NullableAccountAggregationCompletedStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAggregationCompletedStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


