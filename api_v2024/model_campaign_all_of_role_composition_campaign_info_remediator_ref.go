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

// checks if the CampaignAllOfRoleCompositionCampaignInfoRemediatorRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}

// CampaignAllOfRoleCompositionCampaignInfoRemediatorRef This determines who remediation tasks will be assigned to. Remediation tasks are created for each revoke decision on items in the campaign. The only legal remediator type is 'IDENTITY', and the chosen identity must be a Role Admin or Org Admin.
type CampaignAllOfRoleCompositionCampaignInfoRemediatorRef struct {
	// Legal Remediator Type
	Type string `json:"type"`
	// The ID of the remediator.
	Id string `json:"id"`
	// The name of the remediator.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAllOfRoleCompositionCampaignInfoRemediatorRef CampaignAllOfRoleCompositionCampaignInfoRemediatorRef

// NewCampaignAllOfRoleCompositionCampaignInfoRemediatorRef instantiates a new CampaignAllOfRoleCompositionCampaignInfoRemediatorRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOfRoleCompositionCampaignInfoRemediatorRef(type_ string, id string) *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	this := CampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCampaignAllOfRoleCompositionCampaignInfoRemediatorRefWithDefaults instantiates a new CampaignAllOfRoleCompositionCampaignInfoRemediatorRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfRoleCompositionCampaignInfoRemediatorRefWithDefaults() *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	this := CampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) SetName(v string) {
	o.Name = &v
}

func (o CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varCampaignAllOfRoleCompositionCampaignInfoRemediatorRef := _CampaignAllOfRoleCompositionCampaignInfoRemediatorRef{}

	err = json.Unmarshal(data, &varCampaignAllOfRoleCompositionCampaignInfoRemediatorRef)

	if err != nil {
		return err
	}

	*o = CampaignAllOfRoleCompositionCampaignInfoRemediatorRef(varCampaignAllOfRoleCompositionCampaignInfoRemediatorRef)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef struct {
	value *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef
	isSet bool
}

func (v NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef) Get() *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	return v.value
}

func (v *NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef) Set(val *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef(val *CampaignAllOfRoleCompositionCampaignInfoRemediatorRef) *NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	return &NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef{value: val, isSet: true}
}

func (v NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOfRoleCompositionCampaignInfoRemediatorRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


