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

// checks if the WorkgroupDeleteItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupDeleteItem{}

// WorkgroupDeleteItem struct for WorkgroupDeleteItem
type WorkgroupDeleteItem struct {
	// Id of the Governance Group.
	Id string `json:"id"`
	//  The HTTP response status code returned for an individual Governance Group that is requested for deletion during a bulk delete operation.  > 204   -  Governance Group deleted successfully.  > 409   - Governance Group is in use,hence can not be deleted.  > 404   - Governance Group not found. 
	Status string `json:"status"`
	// Human readable status description and containing additional context information about success or failures etc. 
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupDeleteItem WorkgroupDeleteItem

// NewWorkgroupDeleteItem instantiates a new WorkgroupDeleteItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupDeleteItem(id string, status string) *WorkgroupDeleteItem {
	this := WorkgroupDeleteItem{}
	this.Id = id
	this.Status = status
	return &this
}

// NewWorkgroupDeleteItemWithDefaults instantiates a new WorkgroupDeleteItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupDeleteItemWithDefaults() *WorkgroupDeleteItem {
	this := WorkgroupDeleteItem{}
	return &this
}

// GetId returns the Id field value
func (o *WorkgroupDeleteItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkgroupDeleteItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkgroupDeleteItem) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *WorkgroupDeleteItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WorkgroupDeleteItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WorkgroupDeleteItem) SetStatus(v string) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkgroupDeleteItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDeleteItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkgroupDeleteItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkgroupDeleteItem) SetDescription(v string) {
	o.Description = &v
}

func (o WorkgroupDeleteItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupDeleteItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupDeleteItem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varWorkgroupDeleteItem := _WorkgroupDeleteItem{}

	if err = json.Unmarshal(bytes, &varWorkgroupDeleteItem); err == nil {
	*o = WorkgroupDeleteItem(varWorkgroupDeleteItem)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupDeleteItem struct {
	value *WorkgroupDeleteItem
	isSet bool
}

func (v NullableWorkgroupDeleteItem) Get() *WorkgroupDeleteItem {
	return v.value
}

func (v *NullableWorkgroupDeleteItem) Set(val *WorkgroupDeleteItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupDeleteItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupDeleteItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupDeleteItem(val *WorkgroupDeleteItem) *NullableWorkgroupDeleteItem {
	return &NullableWorkgroupDeleteItem{value: val, isSet: true}
}

func (v NullableWorkgroupDeleteItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupDeleteItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


