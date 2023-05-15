/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the SodViolationCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationCheck{}

// SodViolationCheck An object referencing an SOD violation check
type SodViolationCheck struct {
	// The id of the original request
	RequestId string `json:"requestId"`
	// The date-time when this request was created.
	Created *time.Time `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationCheck SodViolationCheck

// NewSodViolationCheck instantiates a new SodViolationCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationCheck(requestId string) *SodViolationCheck {
	this := SodViolationCheck{}
	this.RequestId = requestId
	return &this
}

// NewSodViolationCheckWithDefaults instantiates a new SodViolationCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationCheckWithDefaults() *SodViolationCheck {
	this := SodViolationCheck{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SodViolationCheck) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SodViolationCheck) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SodViolationCheck) SetRequestId(v string) {
	o.RequestId = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SodViolationCheck) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheck) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SodViolationCheck) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SodViolationCheck) SetCreated(v time.Time) {
	o.Created = &v
}

func (o SodViolationCheck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestId"] = o.RequestId
	// skip: created is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationCheck) UnmarshalJSON(bytes []byte) (err error) {
	varSodViolationCheck := _SodViolationCheck{}

	if err = json.Unmarshal(bytes, &varSodViolationCheck); err == nil {
		*o = SodViolationCheck(varSodViolationCheck)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationCheck struct {
	value *SodViolationCheck
	isSet bool
}

func (v NullableSodViolationCheck) Get() *SodViolationCheck {
	return v.value
}

func (v *NullableSodViolationCheck) Set(val *SodViolationCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationCheck(val *SodViolationCheck) *NullableSodViolationCheck {
	return &NullableSodViolationCheck{value: val, isSet: true}
}

func (v NullableSodViolationCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

