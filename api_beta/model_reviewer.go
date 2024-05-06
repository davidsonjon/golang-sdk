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

// checks if the Reviewer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reviewer{}

// Reviewer Details of the reviewer for certification.
type Reviewer struct {
	// The reviewer's DTO type.
	Type string `json:"type"`
	// The reviewer's ID.
	Id string `json:"id"`
	// The reviewer's display name.
	Name string `json:"name"`
	// The reviewing identity's email. Only applicable to `IDENTITY`.
	Email NullableString `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Reviewer Reviewer

// NewReviewer instantiates a new Reviewer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewer(type_ string, id string, name string) *Reviewer {
	this := Reviewer{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewReviewerWithDefaults instantiates a new Reviewer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerWithDefaults() *Reviewer {
	this := Reviewer{}
	return &this
}

// GetType returns the Type field value
func (o *Reviewer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Reviewer) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *Reviewer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Reviewer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Reviewer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Reviewer) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Reviewer) GetEmail() string {
	if o == nil || isNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Reviewer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Reviewer) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Reviewer) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Reviewer) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Reviewer) UnsetEmail() {
	o.Email.Unset()
}

func (o Reviewer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reviewer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Reviewer) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReviewer := _Reviewer{}

	if err = json.Unmarshal(bytes, &varReviewer); err == nil {
	*o = Reviewer(varReviewer)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewer struct {
	value *Reviewer
	isSet bool
}

func (v NullableReviewer) Get() *Reviewer {
	return v.value
}

func (v *NullableReviewer) Set(val *Reviewer) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewer) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewer(val *Reviewer) *NullableReviewer {
	return &NullableReviewer{value: val, isSet: true}
}

func (v NullableReviewer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


