/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the SedAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SedAssignment{}

// SedAssignment Sed Assignment
type SedAssignment struct {
	Assignee *SedAssignee `json:"assignee,omitempty"`
	// List of SED id's
	Items []string `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SedAssignment SedAssignment

// NewSedAssignment instantiates a new SedAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSedAssignment() *SedAssignment {
	this := SedAssignment{}
	return &this
}

// NewSedAssignmentWithDefaults instantiates a new SedAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSedAssignmentWithDefaults() *SedAssignment {
	this := SedAssignment{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *SedAssignment) GetAssignee() SedAssignee {
	if o == nil || isNil(o.Assignee) {
		var ret SedAssignee
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SedAssignment) GetAssigneeOk() (*SedAssignee, bool) {
	if o == nil || isNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *SedAssignment) HasAssignee() bool {
	if o != nil && !isNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given SedAssignee and assigns it to the Assignee field.
func (o *SedAssignment) SetAssignee(v SedAssignee) {
	o.Assignee = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SedAssignment) GetItems() []string {
	if o == nil || isNil(o.Items) {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SedAssignment) GetItemsOk() ([]string, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SedAssignment) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *SedAssignment) SetItems(v []string) {
	o.Items = v
}

func (o SedAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SedAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SedAssignment) UnmarshalJSON(bytes []byte) (err error) {
	varSedAssignment := _SedAssignment{}

	if err = json.Unmarshal(bytes, &varSedAssignment); err == nil {
	*o = SedAssignment(varSedAssignment)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSedAssignment struct {
	value *SedAssignment
	isSet bool
}

func (v NullableSedAssignment) Get() *SedAssignment {
	return v.value
}

func (v *NullableSedAssignment) Set(val *SedAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableSedAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableSedAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSedAssignment(val *SedAssignment) *NullableSedAssignment {
	return &NullableSedAssignment{value: val, isSet: true}
}

func (v NullableSedAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSedAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


