/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the SpConfigExportResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpConfigExportResults{}

// SpConfigExportResults Response model for config export download response.
type SpConfigExportResults struct {
	// Current version of the export results object.
	Version *int32 `json:"version,omitempty"`
	// Time the export was completed.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Name of the tenant where this export originated.
	Tenant *string `json:"tenant,omitempty"`
	// Optional user defined description/name for export job.
	Description *string `json:"description,omitempty"`
	Options *ExportOptions `json:"options,omitempty"`
	Objects []ConfigObject `json:"objects,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpConfigExportResults SpConfigExportResults

// NewSpConfigExportResults instantiates a new SpConfigExportResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpConfigExportResults() *SpConfigExportResults {
	this := SpConfigExportResults{}
	return &this
}

// NewSpConfigExportResultsWithDefaults instantiates a new SpConfigExportResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpConfigExportResultsWithDefaults() *SpConfigExportResults {
	this := SpConfigExportResults{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SpConfigExportResults) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigExportResults) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SpConfigExportResults) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SpConfigExportResults) SetVersion(v int32) {
	o.Version = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SpConfigExportResults) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigExportResults) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SpConfigExportResults) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SpConfigExportResults) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *SpConfigExportResults) GetTenant() string {
	if o == nil || isNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigExportResults) GetTenantOk() (*string, bool) {
	if o == nil || isNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *SpConfigExportResults) HasTenant() bool {
	if o != nil && !isNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *SpConfigExportResults) SetTenant(v string) {
	o.Tenant = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SpConfigExportResults) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigExportResults) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SpConfigExportResults) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SpConfigExportResults) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SpConfigExportResults) GetOptions() ExportOptions {
	if o == nil || isNil(o.Options) {
		var ret ExportOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigExportResults) GetOptionsOk() (*ExportOptions, bool) {
	if o == nil || isNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SpConfigExportResults) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ExportOptions and assigns it to the Options field.
func (o *SpConfigExportResults) SetOptions(v ExportOptions) {
	o.Options = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *SpConfigExportResults) GetObjects() []ConfigObject {
	if o == nil || isNil(o.Objects) {
		var ret []ConfigObject
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigExportResults) GetObjectsOk() ([]ConfigObject, bool) {
	if o == nil || isNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *SpConfigExportResults) HasObjects() bool {
	if o != nil && !isNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []ConfigObject and assigns it to the Objects field.
func (o *SpConfigExportResults) SetObjects(v []ConfigObject) {
	o.Objects = v
}

func (o SpConfigExportResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpConfigExportResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !isNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpConfigExportResults) UnmarshalJSON(bytes []byte) (err error) {
	varSpConfigExportResults := _SpConfigExportResults{}

	if err = json.Unmarshal(bytes, &varSpConfigExportResults); err == nil {
	*o = SpConfigExportResults(varSpConfigExportResults)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "options")
		delete(additionalProperties, "objects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpConfigExportResults struct {
	value *SpConfigExportResults
	isSet bool
}

func (v NullableSpConfigExportResults) Get() *SpConfigExportResults {
	return v.value
}

func (v *NullableSpConfigExportResults) Set(val *SpConfigExportResults) {
	v.value = val
	v.isSet = true
}

func (v NullableSpConfigExportResults) IsSet() bool {
	return v.isSet
}

func (v *NullableSpConfigExportResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpConfigExportResults(val *SpConfigExportResults) *NullableSpConfigExportResults {
	return &NullableSpConfigExportResults{value: val, isSet: true}
}

func (v NullableSpConfigExportResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpConfigExportResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


