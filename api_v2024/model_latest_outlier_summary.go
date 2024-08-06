/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the LatestOutlierSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LatestOutlierSummary{}

// LatestOutlierSummary struct for LatestOutlierSummary
type LatestOutlierSummary struct {
	// The type of outlier summary
	Type *string `json:"type,omitempty"`
	// The date the bulk outlier detection ran/snapshot was created
	SnapshotDate *time.Time `json:"snapshotDate,omitempty"`
	// Total number of outliers for the customer making the request
	TotalOutliers *int32 `json:"totalOutliers,omitempty"`
	// Total number of identities for the customer making the request
	TotalIdentities *int32 `json:"totalIdentities,omitempty"`
	// Total number of ignored outliers
	TotalIgnored *int32 `json:"totalIgnored,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LatestOutlierSummary LatestOutlierSummary

// NewLatestOutlierSummary instantiates a new LatestOutlierSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLatestOutlierSummary() *LatestOutlierSummary {
	this := LatestOutlierSummary{}
	return &this
}

// NewLatestOutlierSummaryWithDefaults instantiates a new LatestOutlierSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLatestOutlierSummaryWithDefaults() *LatestOutlierSummary {
	this := LatestOutlierSummary{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LatestOutlierSummary) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestOutlierSummary) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LatestOutlierSummary) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LatestOutlierSummary) SetType(v string) {
	o.Type = &v
}

// GetSnapshotDate returns the SnapshotDate field value if set, zero value otherwise.
func (o *LatestOutlierSummary) GetSnapshotDate() time.Time {
	if o == nil || isNil(o.SnapshotDate) {
		var ret time.Time
		return ret
	}
	return *o.SnapshotDate
}

// GetSnapshotDateOk returns a tuple with the SnapshotDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestOutlierSummary) GetSnapshotDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.SnapshotDate) {
		return nil, false
	}
	return o.SnapshotDate, true
}

// HasSnapshotDate returns a boolean if a field has been set.
func (o *LatestOutlierSummary) HasSnapshotDate() bool {
	if o != nil && !isNil(o.SnapshotDate) {
		return true
	}

	return false
}

// SetSnapshotDate gets a reference to the given time.Time and assigns it to the SnapshotDate field.
func (o *LatestOutlierSummary) SetSnapshotDate(v time.Time) {
	o.SnapshotDate = &v
}

// GetTotalOutliers returns the TotalOutliers field value if set, zero value otherwise.
func (o *LatestOutlierSummary) GetTotalOutliers() int32 {
	if o == nil || isNil(o.TotalOutliers) {
		var ret int32
		return ret
	}
	return *o.TotalOutliers
}

// GetTotalOutliersOk returns a tuple with the TotalOutliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestOutlierSummary) GetTotalOutliersOk() (*int32, bool) {
	if o == nil || isNil(o.TotalOutliers) {
		return nil, false
	}
	return o.TotalOutliers, true
}

// HasTotalOutliers returns a boolean if a field has been set.
func (o *LatestOutlierSummary) HasTotalOutliers() bool {
	if o != nil && !isNil(o.TotalOutliers) {
		return true
	}

	return false
}

// SetTotalOutliers gets a reference to the given int32 and assigns it to the TotalOutliers field.
func (o *LatestOutlierSummary) SetTotalOutliers(v int32) {
	o.TotalOutliers = &v
}

// GetTotalIdentities returns the TotalIdentities field value if set, zero value otherwise.
func (o *LatestOutlierSummary) GetTotalIdentities() int32 {
	if o == nil || isNil(o.TotalIdentities) {
		var ret int32
		return ret
	}
	return *o.TotalIdentities
}

// GetTotalIdentitiesOk returns a tuple with the TotalIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestOutlierSummary) GetTotalIdentitiesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalIdentities) {
		return nil, false
	}
	return o.TotalIdentities, true
}

// HasTotalIdentities returns a boolean if a field has been set.
func (o *LatestOutlierSummary) HasTotalIdentities() bool {
	if o != nil && !isNil(o.TotalIdentities) {
		return true
	}

	return false
}

// SetTotalIdentities gets a reference to the given int32 and assigns it to the TotalIdentities field.
func (o *LatestOutlierSummary) SetTotalIdentities(v int32) {
	o.TotalIdentities = &v
}

// GetTotalIgnored returns the TotalIgnored field value if set, zero value otherwise.
func (o *LatestOutlierSummary) GetTotalIgnored() int32 {
	if o == nil || isNil(o.TotalIgnored) {
		var ret int32
		return ret
	}
	return *o.TotalIgnored
}

// GetTotalIgnoredOk returns a tuple with the TotalIgnored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestOutlierSummary) GetTotalIgnoredOk() (*int32, bool) {
	if o == nil || isNil(o.TotalIgnored) {
		return nil, false
	}
	return o.TotalIgnored, true
}

// HasTotalIgnored returns a boolean if a field has been set.
func (o *LatestOutlierSummary) HasTotalIgnored() bool {
	if o != nil && !isNil(o.TotalIgnored) {
		return true
	}

	return false
}

// SetTotalIgnored gets a reference to the given int32 and assigns it to the TotalIgnored field.
func (o *LatestOutlierSummary) SetTotalIgnored(v int32) {
	o.TotalIgnored = &v
}

func (o LatestOutlierSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LatestOutlierSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.SnapshotDate) {
		toSerialize["snapshotDate"] = o.SnapshotDate
	}
	if !isNil(o.TotalOutliers) {
		toSerialize["totalOutliers"] = o.TotalOutliers
	}
	if !isNil(o.TotalIdentities) {
		toSerialize["totalIdentities"] = o.TotalIdentities
	}
	if !isNil(o.TotalIgnored) {
		toSerialize["totalIgnored"] = o.TotalIgnored
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LatestOutlierSummary) UnmarshalJSON(bytes []byte) (err error) {
	varLatestOutlierSummary := _LatestOutlierSummary{}

	if err = json.Unmarshal(bytes, &varLatestOutlierSummary); err == nil {
	*o = LatestOutlierSummary(varLatestOutlierSummary)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "snapshotDate")
		delete(additionalProperties, "totalOutliers")
		delete(additionalProperties, "totalIdentities")
		delete(additionalProperties, "totalIgnored")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLatestOutlierSummary struct {
	value *LatestOutlierSummary
	isSet bool
}

func (v NullableLatestOutlierSummary) Get() *LatestOutlierSummary {
	return v.value
}

func (v *NullableLatestOutlierSummary) Set(val *LatestOutlierSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLatestOutlierSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLatestOutlierSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatestOutlierSummary(val *LatestOutlierSummary) *NullableLatestOutlierSummary {
	return &NullableLatestOutlierSummary{value: val, isSet: true}
}

func (v NullableLatestOutlierSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatestOutlierSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


