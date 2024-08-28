/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the RoleInsightsEntitlementChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleInsightsEntitlementChanges{}

// RoleInsightsEntitlementChanges struct for RoleInsightsEntitlementChanges
type RoleInsightsEntitlementChanges struct {
	// Name of the entitlement
	Name *string `json:"name,omitempty"`
	// Id of the entitlement
	Id *string `json:"id,omitempty"`
	// Description for the entitlement
	Description *string `json:"description,omitempty"`
	// Attribute for the entitlement
	Attribute *string `json:"attribute,omitempty"`
	// Attribute value for the entitlement
	Value *string `json:"value,omitempty"`
	// Source or the application for the entitlement
	Source *string `json:"source,omitempty"`
	Insight *RoleInsightsInsight `json:"insight,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleInsightsEntitlementChanges RoleInsightsEntitlementChanges

// NewRoleInsightsEntitlementChanges instantiates a new RoleInsightsEntitlementChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleInsightsEntitlementChanges() *RoleInsightsEntitlementChanges {
	this := RoleInsightsEntitlementChanges{}
	return &this
}

// NewRoleInsightsEntitlementChangesWithDefaults instantiates a new RoleInsightsEntitlementChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleInsightsEntitlementChangesWithDefaults() *RoleInsightsEntitlementChanges {
	this := RoleInsightsEntitlementChanges{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleInsightsEntitlementChanges) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleInsightsEntitlementChanges) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleInsightsEntitlementChanges) SetDescription(v string) {
	o.Description = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *RoleInsightsEntitlementChanges) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RoleInsightsEntitlementChanges) SetValue(v string) {
	o.Value = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RoleInsightsEntitlementChanges) SetSource(v string) {
	o.Source = &v
}

// GetInsight returns the Insight field value if set, zero value otherwise.
func (o *RoleInsightsEntitlementChanges) GetInsight() RoleInsightsInsight {
	if o == nil || IsNil(o.Insight) {
		var ret RoleInsightsInsight
		return ret
	}
	return *o.Insight
}

// GetInsightOk returns a tuple with the Insight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsEntitlementChanges) GetInsightOk() (*RoleInsightsInsight, bool) {
	if o == nil || IsNil(o.Insight) {
		return nil, false
	}
	return o.Insight, true
}

// HasInsight returns a boolean if a field has been set.
func (o *RoleInsightsEntitlementChanges) HasInsight() bool {
	if o != nil && !IsNil(o.Insight) {
		return true
	}

	return false
}

// SetInsight gets a reference to the given RoleInsightsInsight and assigns it to the Insight field.
func (o *RoleInsightsEntitlementChanges) SetInsight(v RoleInsightsInsight) {
	o.Insight = &v
}

func (o RoleInsightsEntitlementChanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleInsightsEntitlementChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Insight) {
		toSerialize["insight"] = o.Insight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleInsightsEntitlementChanges) UnmarshalJSON(data []byte) (err error) {
	varRoleInsightsEntitlementChanges := _RoleInsightsEntitlementChanges{}

	err = json.Unmarshal(data, &varRoleInsightsEntitlementChanges)

	if err != nil {
		return err
	}

	*o = RoleInsightsEntitlementChanges(varRoleInsightsEntitlementChanges)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "source")
		delete(additionalProperties, "insight")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleInsightsEntitlementChanges struct {
	value *RoleInsightsEntitlementChanges
	isSet bool
}

func (v NullableRoleInsightsEntitlementChanges) Get() *RoleInsightsEntitlementChanges {
	return v.value
}

func (v *NullableRoleInsightsEntitlementChanges) Set(val *RoleInsightsEntitlementChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleInsightsEntitlementChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleInsightsEntitlementChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleInsightsEntitlementChanges(val *RoleInsightsEntitlementChanges) *NullableRoleInsightsEntitlementChanges {
	return &NullableRoleInsightsEntitlementChanges{value: val, isSet: true}
}

func (v NullableRoleInsightsEntitlementChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleInsightsEntitlementChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


