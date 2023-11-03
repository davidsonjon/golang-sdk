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

// checks if the WorkgroupMemberAddItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupMemberAddItem{}

// WorkgroupMemberAddItem struct for WorkgroupMemberAddItem
type WorkgroupMemberAddItem struct {
	// Identifier of identity in bulk member add request.
	Id string `json:"id"`
	//  The HTTP response status code returned for an individual member that is requested for addition during a bulk add operation.   The HTTP response status code returned for an individual Governance Group is requested for deletion.   > 201   - Identity is added into Governance Group members list.  > 409   - Identity is already member of  Governance Group. 
	Status string `json:"status"`
	// Human readable status description and containing additional context information about success or failures etc. 
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupMemberAddItem WorkgroupMemberAddItem

// NewWorkgroupMemberAddItem instantiates a new WorkgroupMemberAddItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupMemberAddItem(id string, status string) *WorkgroupMemberAddItem {
	this := WorkgroupMemberAddItem{}
	this.Id = id
	this.Status = status
	return &this
}

// NewWorkgroupMemberAddItemWithDefaults instantiates a new WorkgroupMemberAddItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupMemberAddItemWithDefaults() *WorkgroupMemberAddItem {
	this := WorkgroupMemberAddItem{}
	return &this
}

// GetId returns the Id field value
func (o *WorkgroupMemberAddItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkgroupMemberAddItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkgroupMemberAddItem) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *WorkgroupMemberAddItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WorkgroupMemberAddItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WorkgroupMemberAddItem) SetStatus(v string) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkgroupMemberAddItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupMemberAddItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkgroupMemberAddItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkgroupMemberAddItem) SetDescription(v string) {
	o.Description = &v
}

func (o WorkgroupMemberAddItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupMemberAddItem) ToMap() (map[string]interface{}, error) {
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

func (o *WorkgroupMemberAddItem) UnmarshalJSON(bytes []byte) (err error) {
	varWorkgroupMemberAddItem := _WorkgroupMemberAddItem{}

	if err = json.Unmarshal(bytes, &varWorkgroupMemberAddItem); err == nil {
		*o = WorkgroupMemberAddItem(varWorkgroupMemberAddItem)
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

type NullableWorkgroupMemberAddItem struct {
	value *WorkgroupMemberAddItem
	isSet bool
}

func (v NullableWorkgroupMemberAddItem) Get() *WorkgroupMemberAddItem {
	return v.value
}

func (v *NullableWorkgroupMemberAddItem) Set(val *WorkgroupMemberAddItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupMemberAddItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupMemberAddItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupMemberAddItem(val *WorkgroupMemberAddItem) *NullableWorkgroupMemberAddItem {
	return &NullableWorkgroupMemberAddItem{value: val, isSet: true}
}

func (v NullableWorkgroupMemberAddItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupMemberAddItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

