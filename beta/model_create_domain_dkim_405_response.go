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

// checks if the CreateDomainDkim405Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDomainDkim405Response{}

// CreateDomainDkim405Response struct for CreateDomainDkim405Response
type CreateDomainDkim405Response struct {
	// A message describing the error
	ErrorName map[string]interface{} `json:"errorName,omitempty"`
	// Description of the error
	ErrorMessage map[string]interface{} `json:"errorMessage,omitempty"`
	// Unique tracking id for the error.
	TrackingId *string `json:"trackingId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDomainDkim405Response CreateDomainDkim405Response

// NewCreateDomainDkim405Response instantiates a new CreateDomainDkim405Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomainDkim405Response() *CreateDomainDkim405Response {
	this := CreateDomainDkim405Response{}
	return &this
}

// NewCreateDomainDkim405ResponseWithDefaults instantiates a new CreateDomainDkim405Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainDkim405ResponseWithDefaults() *CreateDomainDkim405Response {
	this := CreateDomainDkim405Response{}
	return &this
}

// GetErrorName returns the ErrorName field value if set, zero value otherwise.
func (o *CreateDomainDkim405Response) GetErrorName() map[string]interface{} {
	if o == nil || isNil(o.ErrorName) {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorName
}

// GetErrorNameOk returns a tuple with the ErrorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainDkim405Response) GetErrorNameOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ErrorName) {
		return map[string]interface{}{}, false
	}
	return o.ErrorName, true
}

// HasErrorName returns a boolean if a field has been set.
func (o *CreateDomainDkim405Response) HasErrorName() bool {
	if o != nil && !isNil(o.ErrorName) {
		return true
	}

	return false
}

// SetErrorName gets a reference to the given map[string]interface{} and assigns it to the ErrorName field.
func (o *CreateDomainDkim405Response) SetErrorName(v map[string]interface{}) {
	o.ErrorName = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CreateDomainDkim405Response) GetErrorMessage() map[string]interface{} {
	if o == nil || isNil(o.ErrorMessage) {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainDkim405Response) GetErrorMessageOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return map[string]interface{}{}, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CreateDomainDkim405Response) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given map[string]interface{} and assigns it to the ErrorMessage field.
func (o *CreateDomainDkim405Response) SetErrorMessage(v map[string]interface{}) {
	o.ErrorMessage = v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *CreateDomainDkim405Response) GetTrackingId() string {
	if o == nil || isNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainDkim405Response) GetTrackingIdOk() (*string, bool) {
	if o == nil || isNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *CreateDomainDkim405Response) HasTrackingId() bool {
	if o != nil && !isNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *CreateDomainDkim405Response) SetTrackingId(v string) {
	o.TrackingId = &v
}

func (o CreateDomainDkim405Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDomainDkim405Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ErrorName) {
		toSerialize["errorName"] = o.ErrorName
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDomainDkim405Response) UnmarshalJSON(bytes []byte) (err error) {
	varCreateDomainDkim405Response := _CreateDomainDkim405Response{}

	if err = json.Unmarshal(bytes, &varCreateDomainDkim405Response); err == nil {
	*o = CreateDomainDkim405Response(varCreateDomainDkim405Response)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errorName")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "trackingId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDomainDkim405Response struct {
	value *CreateDomainDkim405Response
	isSet bool
}

func (v NullableCreateDomainDkim405Response) Get() *CreateDomainDkim405Response {
	return v.value
}

func (v *NullableCreateDomainDkim405Response) Set(val *CreateDomainDkim405Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomainDkim405Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomainDkim405Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomainDkim405Response(val *CreateDomainDkim405Response) *NullableCreateDomainDkim405Response {
	return &NullableCreateDomainDkim405Response{value: val, isSet: true}
}

func (v NullableCreateDomainDkim405Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomainDkim405Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


