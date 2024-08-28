/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the WorkgroupMemberDeleteItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupMemberDeleteItem{}

// WorkgroupMemberDeleteItem struct for WorkgroupMemberDeleteItem
type WorkgroupMemberDeleteItem struct {
	// Identifier of identity in bulk member add /remove request.
	Id string `json:"id"`
	//  The HTTP response status code returned for an individual  member that is requested for deletion during a bulk delete operation.  > 204   - Identity is removed from Governance Group members list.  > 404   - Identity is not member of Governance Group. 
	Status string `json:"status"`
	// Human readable status description and containing additional context information about success or failures etc. 
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupMemberDeleteItem WorkgroupMemberDeleteItem

// NewWorkgroupMemberDeleteItem instantiates a new WorkgroupMemberDeleteItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupMemberDeleteItem(id string, status string) *WorkgroupMemberDeleteItem {
	this := WorkgroupMemberDeleteItem{}
	this.Id = id
	this.Status = status
	return &this
}

// NewWorkgroupMemberDeleteItemWithDefaults instantiates a new WorkgroupMemberDeleteItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupMemberDeleteItemWithDefaults() *WorkgroupMemberDeleteItem {
	this := WorkgroupMemberDeleteItem{}
	return &this
}

// GetId returns the Id field value
func (o *WorkgroupMemberDeleteItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkgroupMemberDeleteItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkgroupMemberDeleteItem) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *WorkgroupMemberDeleteItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WorkgroupMemberDeleteItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WorkgroupMemberDeleteItem) SetStatus(v string) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkgroupMemberDeleteItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupMemberDeleteItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkgroupMemberDeleteItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkgroupMemberDeleteItem) SetDescription(v string) {
	o.Description = &v
}

func (o WorkgroupMemberDeleteItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupMemberDeleteItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupMemberDeleteItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varWorkgroupMemberDeleteItem := _WorkgroupMemberDeleteItem{}

	err = json.Unmarshal(data, &varWorkgroupMemberDeleteItem)

	if err != nil {
		return err
	}

	*o = WorkgroupMemberDeleteItem(varWorkgroupMemberDeleteItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupMemberDeleteItem struct {
	value *WorkgroupMemberDeleteItem
	isSet bool
}

func (v NullableWorkgroupMemberDeleteItem) Get() *WorkgroupMemberDeleteItem {
	return v.value
}

func (v *NullableWorkgroupMemberDeleteItem) Set(val *WorkgroupMemberDeleteItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupMemberDeleteItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupMemberDeleteItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupMemberDeleteItem(val *WorkgroupMemberDeleteItem) *NullableWorkgroupMemberDeleteItem {
	return &NullableWorkgroupMemberDeleteItem{value: val, isSet: true}
}

func (v NullableWorkgroupMemberDeleteItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupMemberDeleteItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


