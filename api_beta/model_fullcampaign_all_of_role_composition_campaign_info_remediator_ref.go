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

// checks if the FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}

// FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef This determines who remediation tasks will be assigned to. Remediation tasks are created for each revoke decision on items in the campaign. The only legal remediator type is 'IDENTITY', and the chosen identity must be a Role Admin or Org Admin.
type FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef struct {
	// Legal Remediator Type
	Type string `json:"type"`
	// The ID of the remediator.
	Id string `json:"id"`
	// The name of the remediator.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef

// NewFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef instantiates a new FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef(type_ string, id string) *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	this := FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRefWithDefaults instantiates a new FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRefWithDefaults() *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	this := FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}
	return &this
}

// GetType returns the Type field value
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) SetName(v string) {
	o.Name = &v
}

func (o FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	// skip: name is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef := _FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}

	if err = json.Unmarshal(bytes, &varFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef); err == nil {
	*o = FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef(varFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef struct {
	value *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef
	isSet bool
}

func (v NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) Get() *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	return v.value
}

func (v *NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) Set(val *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) {
	v.value = val
	v.isSet = true
}

func (v NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) IsSet() bool {
	return v.isSet
}

func (v *NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef(val *FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) *NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	return &NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef{value: val, isSet: true}
}

func (v NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


