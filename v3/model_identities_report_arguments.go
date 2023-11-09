/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the IdentitiesReportArguments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitiesReportArguments{}

// IdentitiesReportArguments Arguments for Identities report (IDENTITIES)
type IdentitiesReportArguments struct {
	// Boolean FLAG to specify if only correlated identities should be used in report processing
	CorrelatedOnly *bool `json:"correlatedOnly,omitempty"`
	// Use it to set default s3 bucket where generated report will be saved.  In case this argument is false and 's3Bucket' argument is null or absent there will be default s3Bucket assigned to the report.
	DefaultS3Bucket bool `json:"defaultS3Bucket"`
	// If you want to be specific you could use this argument with defaultS3Bucket = false.
	S3Bucket *string `json:"s3Bucket,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitiesReportArguments IdentitiesReportArguments

// NewIdentitiesReportArguments instantiates a new IdentitiesReportArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitiesReportArguments(defaultS3Bucket bool) *IdentitiesReportArguments {
	this := IdentitiesReportArguments{}
	var correlatedOnly bool = false
	this.CorrelatedOnly = &correlatedOnly
	this.DefaultS3Bucket = defaultS3Bucket
	return &this
}

// NewIdentitiesReportArgumentsWithDefaults instantiates a new IdentitiesReportArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitiesReportArgumentsWithDefaults() *IdentitiesReportArguments {
	this := IdentitiesReportArguments{}
	var correlatedOnly bool = false
	this.CorrelatedOnly = &correlatedOnly
	return &this
}

// GetCorrelatedOnly returns the CorrelatedOnly field value if set, zero value otherwise.
func (o *IdentitiesReportArguments) GetCorrelatedOnly() bool {
	if o == nil || isNil(o.CorrelatedOnly) {
		var ret bool
		return ret
	}
	return *o.CorrelatedOnly
}

// GetCorrelatedOnlyOk returns a tuple with the CorrelatedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitiesReportArguments) GetCorrelatedOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.CorrelatedOnly) {
		return nil, false
	}
	return o.CorrelatedOnly, true
}

// HasCorrelatedOnly returns a boolean if a field has been set.
func (o *IdentitiesReportArguments) HasCorrelatedOnly() bool {
	if o != nil && !isNil(o.CorrelatedOnly) {
		return true
	}

	return false
}

// SetCorrelatedOnly gets a reference to the given bool and assigns it to the CorrelatedOnly field.
func (o *IdentitiesReportArguments) SetCorrelatedOnly(v bool) {
	o.CorrelatedOnly = &v
}

// GetDefaultS3Bucket returns the DefaultS3Bucket field value
func (o *IdentitiesReportArguments) GetDefaultS3Bucket() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultS3Bucket
}

// GetDefaultS3BucketOk returns a tuple with the DefaultS3Bucket field value
// and a boolean to check if the value has been set.
func (o *IdentitiesReportArguments) GetDefaultS3BucketOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultS3Bucket, true
}

// SetDefaultS3Bucket sets field value
func (o *IdentitiesReportArguments) SetDefaultS3Bucket(v bool) {
	o.DefaultS3Bucket = v
}

// GetS3Bucket returns the S3Bucket field value if set, zero value otherwise.
func (o *IdentitiesReportArguments) GetS3Bucket() string {
	if o == nil || isNil(o.S3Bucket) {
		var ret string
		return ret
	}
	return *o.S3Bucket
}

// GetS3BucketOk returns a tuple with the S3Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitiesReportArguments) GetS3BucketOk() (*string, bool) {
	if o == nil || isNil(o.S3Bucket) {
		return nil, false
	}
	return o.S3Bucket, true
}

// HasS3Bucket returns a boolean if a field has been set.
func (o *IdentitiesReportArguments) HasS3Bucket() bool {
	if o != nil && !isNil(o.S3Bucket) {
		return true
	}

	return false
}

// SetS3Bucket gets a reference to the given string and assigns it to the S3Bucket field.
func (o *IdentitiesReportArguments) SetS3Bucket(v string) {
	o.S3Bucket = &v
}

func (o IdentitiesReportArguments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitiesReportArguments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CorrelatedOnly) {
		toSerialize["correlatedOnly"] = o.CorrelatedOnly
	}
	toSerialize["defaultS3Bucket"] = o.DefaultS3Bucket
	if !isNil(o.S3Bucket) {
		toSerialize["s3Bucket"] = o.S3Bucket
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitiesReportArguments) UnmarshalJSON(bytes []byte) (err error) {
	varIdentitiesReportArguments := _IdentitiesReportArguments{}

	if err = json.Unmarshal(bytes, &varIdentitiesReportArguments); err == nil {
		*o = IdentitiesReportArguments(varIdentitiesReportArguments)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "correlatedOnly")
		delete(additionalProperties, "defaultS3Bucket")
		delete(additionalProperties, "s3Bucket")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitiesReportArguments struct {
	value *IdentitiesReportArguments
	isSet bool
}

func (v NullableIdentitiesReportArguments) Get() *IdentitiesReportArguments {
	return v.value
}

func (v *NullableIdentitiesReportArguments) Set(val *IdentitiesReportArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitiesReportArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitiesReportArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitiesReportArguments(val *IdentitiesReportArguments) *NullableIdentitiesReportArguments {
	return &NullableIdentitiesReportArguments{value: val, isSet: true}
}

func (v NullableIdentitiesReportArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitiesReportArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


