/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// ProvisioningPolicy struct for ProvisioningPolicy
type ProvisioningPolicy struct {
	// the provisioning policy name
	Name string `json:"name"`
	// the description of the provisioning policy
	Description *string `json:"description,omitempty"`
	UsageType *UsageType `json:"usageType,omitempty"`
	Fields []FieldDetailsDto `json:"fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningPolicy ProvisioningPolicy

// NewProvisioningPolicy instantiates a new ProvisioningPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningPolicy(name string) *ProvisioningPolicy {
	this := ProvisioningPolicy{}
	this.Name = name
	return &this
}

// NewProvisioningPolicyWithDefaults instantiates a new ProvisioningPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningPolicyWithDefaults() *ProvisioningPolicy {
	this := ProvisioningPolicy{}
	return &this
}

// GetName returns the Name field value
func (o *ProvisioningPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProvisioningPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProvisioningPolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProvisioningPolicy) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProvisioningPolicy) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProvisioningPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *ProvisioningPolicy) GetUsageType() UsageType {
	if o == nil || isNil(o.UsageType) {
		var ret UsageType
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningPolicy) GetUsageTypeOk() (*UsageType, bool) {
	if o == nil || isNil(o.UsageType) {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *ProvisioningPolicy) HasUsageType() bool {
	if o != nil && !isNil(o.UsageType) {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given UsageType and assigns it to the UsageType field.
func (o *ProvisioningPolicy) SetUsageType(v UsageType) {
	o.UsageType = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ProvisioningPolicy) GetFields() []FieldDetailsDto {
	if o == nil || isNil(o.Fields) {
		var ret []FieldDetailsDto
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningPolicy) GetFieldsOk() ([]FieldDetailsDto, bool) {
	if o == nil || isNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ProvisioningPolicy) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []FieldDetailsDto and assigns it to the Fields field.
func (o *ProvisioningPolicy) SetFields(v []FieldDetailsDto) {
	o.Fields = v
}

func (o ProvisioningPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.UsageType) {
		toSerialize["usageType"] = o.UsageType
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningPolicy := _ProvisioningPolicy{}

	if err = json.Unmarshal(bytes, &varProvisioningPolicy); err == nil {
		*o = ProvisioningPolicy(varProvisioningPolicy)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "usageType")
		delete(additionalProperties, "fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningPolicy struct {
	value *ProvisioningPolicy
	isSet bool
}

func (v NullableProvisioningPolicy) Get() *ProvisioningPolicy {
	return v.value
}

func (v *NullableProvisioningPolicy) Set(val *ProvisioningPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningPolicy(val *ProvisioningPolicy) *NullableProvisioningPolicy {
	return &NullableProvisioningPolicy{value: val, isSet: true}
}

func (v NullableProvisioningPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


