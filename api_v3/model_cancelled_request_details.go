/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
)

// checks if the CancelledRequestDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelledRequestDetails{}

// CancelledRequestDetails Provides additional details for a request that has been cancelled.
type CancelledRequestDetails struct {
	// Comment made by the owner when cancelling the associated request.
	Comment *string `json:"comment,omitempty"`
	Owner *OwnerDto `json:"owner,omitempty"`
	// Date comment was added by the owner when cancelling the associated request.
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelledRequestDetails CancelledRequestDetails

// NewCancelledRequestDetails instantiates a new CancelledRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelledRequestDetails() *CancelledRequestDetails {
	this := CancelledRequestDetails{}
	return &this
}

// NewCancelledRequestDetailsWithDefaults instantiates a new CancelledRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelledRequestDetailsWithDefaults() *CancelledRequestDetails {
	this := CancelledRequestDetails{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CancelledRequestDetails) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelledRequestDetails) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CancelledRequestDetails) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CancelledRequestDetails) SetComment(v string) {
	o.Comment = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CancelledRequestDetails) GetOwner() OwnerDto {
	if o == nil || IsNil(o.Owner) {
		var ret OwnerDto
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelledRequestDetails) GetOwnerOk() (*OwnerDto, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CancelledRequestDetails) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given OwnerDto and assigns it to the Owner field.
func (o *CancelledRequestDetails) SetOwner(v OwnerDto) {
	o.Owner = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *CancelledRequestDetails) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelledRequestDetails) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *CancelledRequestDetails) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *CancelledRequestDetails) SetModified(v time.Time) {
	o.Modified = &v
}

func (o CancelledRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelledRequestDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelledRequestDetails) UnmarshalJSON(data []byte) (err error) {
	varCancelledRequestDetails := _CancelledRequestDetails{}

	err = json.Unmarshal(data, &varCancelledRequestDetails)

	if err != nil {
		return err
	}

	*o = CancelledRequestDetails(varCancelledRequestDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelledRequestDetails struct {
	value *CancelledRequestDetails
	isSet bool
}

func (v NullableCancelledRequestDetails) Get() *CancelledRequestDetails {
	return v.value
}

func (v *NullableCancelledRequestDetails) Set(val *CancelledRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelledRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelledRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelledRequestDetails(val *CancelledRequestDetails) *NullableCancelledRequestDetails {
	return &NullableCancelledRequestDetails{value: val, isSet: true}
}

func (v NullableCancelledRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelledRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


