/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the Reassignment1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reassignment1{}

// Reassignment1 struct for Reassignment1
type Reassignment1 struct {
	From *CertificationReference1 `json:"from,omitempty"`
	// Comments from the previous reviewer.
	Comment *string `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Reassignment1 Reassignment1

// NewReassignment1 instantiates a new Reassignment1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReassignment1() *Reassignment1 {
	this := Reassignment1{}
	return &this
}

// NewReassignment1WithDefaults instantiates a new Reassignment1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReassignment1WithDefaults() *Reassignment1 {
	this := Reassignment1{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Reassignment1) GetFrom() CertificationReference1 {
	if o == nil || IsNil(o.From) {
		var ret CertificationReference1
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reassignment1) GetFromOk() (*CertificationReference1, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Reassignment1) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given CertificationReference1 and assigns it to the From field.
func (o *Reassignment1) SetFrom(v CertificationReference1) {
	o.From = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Reassignment1) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reassignment1) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Reassignment1) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Reassignment1) SetComment(v string) {
	o.Comment = &v
}

func (o Reassignment1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reassignment1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Reassignment1) UnmarshalJSON(data []byte) (err error) {
	varReassignment1 := _Reassignment1{}

	err = json.Unmarshal(data, &varReassignment1)

	if err != nil {
		return err
	}

	*o = Reassignment1(varReassignment1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReassignment1 struct {
	value *Reassignment1
	isSet bool
}

func (v NullableReassignment1) Get() *Reassignment1 {
	return v.value
}

func (v *NullableReassignment1) Set(val *Reassignment1) {
	v.value = val
	v.isSet = true
}

func (v NullableReassignment1) IsSet() bool {
	return v.isSet
}

func (v *NullableReassignment1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReassignment1(val *Reassignment1) *NullableReassignment1 {
	return &NullableReassignment1{value: val, isSet: true}
}

func (v NullableReassignment1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReassignment1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


