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

// checks if the SourceAccountCorrelationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAccountCorrelationRule{}

// SourceAccountCorrelationRule Reference to a Rule that can do COMPLEX correlation, should only be used when accountCorrelationConfig can't be used.
type SourceAccountCorrelationRule struct {
	// The type of object being referenced
	Type *string `json:"type,omitempty"`
	// ID of the rule
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the rule
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceAccountCorrelationRule SourceAccountCorrelationRule

// NewSourceAccountCorrelationRule instantiates a new SourceAccountCorrelationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccountCorrelationRule() *SourceAccountCorrelationRule {
	this := SourceAccountCorrelationRule{}
	return &this
}

// NewSourceAccountCorrelationRuleWithDefaults instantiates a new SourceAccountCorrelationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountCorrelationRuleWithDefaults() *SourceAccountCorrelationRule {
	this := SourceAccountCorrelationRule{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceAccountCorrelationRule) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountCorrelationRule) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceAccountCorrelationRule) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceAccountCorrelationRule) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceAccountCorrelationRule) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountCorrelationRule) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceAccountCorrelationRule) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceAccountCorrelationRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceAccountCorrelationRule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountCorrelationRule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceAccountCorrelationRule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceAccountCorrelationRule) SetName(v string) {
	o.Name = &v
}

func (o SourceAccountCorrelationRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAccountCorrelationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceAccountCorrelationRule) UnmarshalJSON(bytes []byte) (err error) {
	varSourceAccountCorrelationRule := _SourceAccountCorrelationRule{}

	if err = json.Unmarshal(bytes, &varSourceAccountCorrelationRule); err == nil {
		*o = SourceAccountCorrelationRule(varSourceAccountCorrelationRule)
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

type NullableSourceAccountCorrelationRule struct {
	value *SourceAccountCorrelationRule
	isSet bool
}

func (v NullableSourceAccountCorrelationRule) Get() *SourceAccountCorrelationRule {
	return v.value
}

func (v *NullableSourceAccountCorrelationRule) Set(val *SourceAccountCorrelationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccountCorrelationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccountCorrelationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccountCorrelationRule(val *SourceAccountCorrelationRule) *NullableSourceAccountCorrelationRule {
	return &NullableSourceAccountCorrelationRule{value: val, isSet: true}
}

func (v NullableSourceAccountCorrelationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccountCorrelationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


