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

// CommonAccessItemAccess struct for CommonAccessItemAccess
type CommonAccessItemAccess struct {
	// Common access ID
	Id *string `json:"id,omitempty"`
	Type *CommonAccessType `json:"type,omitempty"`
	// Common access name
	Name *string `json:"name,omitempty"`
	// Common access description
	Description *string `json:"description,omitempty"`
	// Common access owner name
	OwnerName *string `json:"ownerName,omitempty"`
	// Common access owner ID
	OwnerId *string `json:"ownerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonAccessItemAccess CommonAccessItemAccess

// NewCommonAccessItemAccess instantiates a new CommonAccessItemAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAccessItemAccess() *CommonAccessItemAccess {
	this := CommonAccessItemAccess{}
	return &this
}

// NewCommonAccessItemAccessWithDefaults instantiates a new CommonAccessItemAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAccessItemAccessWithDefaults() *CommonAccessItemAccess {
	this := CommonAccessItemAccess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommonAccessItemAccess) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemAccess) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommonAccessItemAccess) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommonAccessItemAccess) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CommonAccessItemAccess) GetType() CommonAccessType {
	if o == nil || isNil(o.Type) {
		var ret CommonAccessType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemAccess) GetTypeOk() (*CommonAccessType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CommonAccessItemAccess) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CommonAccessType and assigns it to the Type field.
func (o *CommonAccessItemAccess) SetType(v CommonAccessType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommonAccessItemAccess) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemAccess) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommonAccessItemAccess) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommonAccessItemAccess) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommonAccessItemAccess) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemAccess) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonAccessItemAccess) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommonAccessItemAccess) SetDescription(v string) {
	o.Description = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *CommonAccessItemAccess) GetOwnerName() string {
	if o == nil || isNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemAccess) GetOwnerNameOk() (*string, bool) {
	if o == nil || isNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *CommonAccessItemAccess) HasOwnerName() bool {
	if o != nil && !isNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *CommonAccessItemAccess) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *CommonAccessItemAccess) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemAccess) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *CommonAccessItemAccess) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *CommonAccessItemAccess) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o CommonAccessItemAccess) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommonAccessItemAccess) UnmarshalJSON(bytes []byte) (err error) {
	varCommonAccessItemAccess := _CommonAccessItemAccess{}

	if err = json.Unmarshal(bytes, &varCommonAccessItemAccess); err == nil {
		*o = CommonAccessItemAccess(varCommonAccessItemAccess)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ownerName")
		delete(additionalProperties, "ownerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAccessItemAccess struct {
	value *CommonAccessItemAccess
	isSet bool
}

func (v NullableCommonAccessItemAccess) Get() *CommonAccessItemAccess {
	return v.value
}

func (v *NullableCommonAccessItemAccess) Set(val *CommonAccessItemAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAccessItemAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAccessItemAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAccessItemAccess(val *CommonAccessItemAccess) *NullableCommonAccessItemAccess {
	return &NullableCommonAccessItemAccess{value: val, isSet: true}
}

func (v NullableCommonAccessItemAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAccessItemAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


