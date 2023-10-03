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

// checks if the GetReferenceIdentityAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReferenceIdentityAttribute{}

// GetReferenceIdentityAttribute struct for GetReferenceIdentityAttribute
type GetReferenceIdentityAttribute struct {
	// This must always be set to \"Cloud Services Deployment Utility\"
	Name string `json:"name"`
	// The operation to perform `getReferenceIdentityAttribute`
	Operation string `json:"operation"`
	// This is the SailPoint User Name (uid) value of the identity whose attribute is desired  As a convenience feature, you can use the `manager` keyword to dynamically look up the user's manager and then get that manager's identity attribute. 
	Uid string `json:"uid"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetReferenceIdentityAttribute GetReferenceIdentityAttribute

// NewGetReferenceIdentityAttribute instantiates a new GetReferenceIdentityAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReferenceIdentityAttribute(name string, operation string, uid string) *GetReferenceIdentityAttribute {
	this := GetReferenceIdentityAttribute{}
	this.Name = name
	this.Operation = operation
	this.Uid = uid
	return &this
}

// NewGetReferenceIdentityAttributeWithDefaults instantiates a new GetReferenceIdentityAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReferenceIdentityAttributeWithDefaults() *GetReferenceIdentityAttribute {
	this := GetReferenceIdentityAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *GetReferenceIdentityAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetReferenceIdentityAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetReferenceIdentityAttribute) SetName(v string) {
	o.Name = v
}

// GetOperation returns the Operation field value
func (o *GetReferenceIdentityAttribute) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *GetReferenceIdentityAttribute) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *GetReferenceIdentityAttribute) SetOperation(v string) {
	o.Operation = v
}

// GetUid returns the Uid field value
func (o *GetReferenceIdentityAttribute) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *GetReferenceIdentityAttribute) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *GetReferenceIdentityAttribute) SetUid(v string) {
	o.Uid = v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *GetReferenceIdentityAttribute) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReferenceIdentityAttribute) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *GetReferenceIdentityAttribute) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *GetReferenceIdentityAttribute) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

func (o GetReferenceIdentityAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReferenceIdentityAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["operation"] = o.Operation
	toSerialize["uid"] = o.Uid
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetReferenceIdentityAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varGetReferenceIdentityAttribute := _GetReferenceIdentityAttribute{}

	if err = json.Unmarshal(bytes, &varGetReferenceIdentityAttribute); err == nil {
		*o = GetReferenceIdentityAttribute(varGetReferenceIdentityAttribute)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "uid")
		delete(additionalProperties, "requiresPeriodicRefresh")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetReferenceIdentityAttribute struct {
	value *GetReferenceIdentityAttribute
	isSet bool
}

func (v NullableGetReferenceIdentityAttribute) Get() *GetReferenceIdentityAttribute {
	return v.value
}

func (v *NullableGetReferenceIdentityAttribute) Set(val *GetReferenceIdentityAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReferenceIdentityAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReferenceIdentityAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReferenceIdentityAttribute(val *GetReferenceIdentityAttribute) *NullableGetReferenceIdentityAttribute {
	return &NullableGetReferenceIdentityAttribute{value: val, isSet: true}
}

func (v NullableGetReferenceIdentityAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReferenceIdentityAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


