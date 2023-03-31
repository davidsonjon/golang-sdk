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

// checks if the CampaignTemplateOwnerRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignTemplateOwnerRef{}

// CampaignTemplateOwnerRef The owner of this template, and the owner of campaigns generated from this template via a schedule. This field is automatically populated at creation time with the current user.
type CampaignTemplateOwnerRef struct {
	// Id of the owner
	Id *string `json:"id,omitempty"`
	// Type of the owner
	Type *string `json:"type,omitempty"`
	// Name of the owner
	Name *string `json:"name,omitempty"`
	// Email of the owner
	Email *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignTemplateOwnerRef CampaignTemplateOwnerRef

// NewCampaignTemplateOwnerRef instantiates a new CampaignTemplateOwnerRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignTemplateOwnerRef() *CampaignTemplateOwnerRef {
	this := CampaignTemplateOwnerRef{}
	return &this
}

// NewCampaignTemplateOwnerRefWithDefaults instantiates a new CampaignTemplateOwnerRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignTemplateOwnerRefWithDefaults() *CampaignTemplateOwnerRef {
	this := CampaignTemplateOwnerRef{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CampaignTemplateOwnerRef) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplateOwnerRef) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CampaignTemplateOwnerRef) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CampaignTemplateOwnerRef) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CampaignTemplateOwnerRef) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplateOwnerRef) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CampaignTemplateOwnerRef) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CampaignTemplateOwnerRef) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignTemplateOwnerRef) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplateOwnerRef) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignTemplateOwnerRef) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignTemplateOwnerRef) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CampaignTemplateOwnerRef) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplateOwnerRef) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CampaignTemplateOwnerRef) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CampaignTemplateOwnerRef) SetEmail(v string) {
	o.Email = &v
}

func (o CampaignTemplateOwnerRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignTemplateOwnerRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignTemplateOwnerRef) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignTemplateOwnerRef := _CampaignTemplateOwnerRef{}

	if err = json.Unmarshal(bytes, &varCampaignTemplateOwnerRef); err == nil {
		*o = CampaignTemplateOwnerRef(varCampaignTemplateOwnerRef)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignTemplateOwnerRef struct {
	value *CampaignTemplateOwnerRef
	isSet bool
}

func (v NullableCampaignTemplateOwnerRef) Get() *CampaignTemplateOwnerRef {
	return v.value
}

func (v *NullableCampaignTemplateOwnerRef) Set(val *CampaignTemplateOwnerRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignTemplateOwnerRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignTemplateOwnerRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignTemplateOwnerRef(val *CampaignTemplateOwnerRef) *NullableCampaignTemplateOwnerRef {
	return &NullableCampaignTemplateOwnerRef{value: val, isSet: true}
}

func (v NullableCampaignTemplateOwnerRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignTemplateOwnerRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


