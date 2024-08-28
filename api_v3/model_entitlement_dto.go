/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the EntitlementDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDto{}

// EntitlementDto struct for EntitlementDto
type EntitlementDto struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// Name of the entitlement attribute
	Attribute *string `json:"attribute,omitempty"`
	// Raw value of the entitlement
	Value *string `json:"value,omitempty"`
	// Entitlment description
	Description *string `json:"description,omitempty"`
	// Entitlement attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Schema objectType on the given application that maps to an Account Group
	SourceSchemaObjectType *string `json:"sourceSchemaObjectType,omitempty"`
	// Determines if this Entitlement is privileged.
	Privileged *bool `json:"privileged,omitempty"`
	// Determines if this Entitlement is goverened in the cloud.
	CloudGoverned *bool `json:"cloudGoverned,omitempty"`
	Source *EntitlementSource `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDto EntitlementDto

// NewEntitlementDto instantiates a new EntitlementDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDto(name string) *EntitlementDto {
	this := EntitlementDto{}
	this.Name = name
	return &this
}

// NewEntitlementDtoWithDefaults instantiates a new EntitlementDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDtoWithDefaults() *EntitlementDto {
	this := EntitlementDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *EntitlementDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementDto) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EntitlementDto) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EntitlementDto) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EntitlementDto) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *EntitlementDto) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *EntitlementDto) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *EntitlementDto) SetModified(v time.Time) {
	o.Modified = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *EntitlementDto) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *EntitlementDto) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *EntitlementDto) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementDto) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementDto) SetDescription(v string) {
	o.Description = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntitlementDto) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntitlementDto) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *EntitlementDto) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetSourceSchemaObjectType returns the SourceSchemaObjectType field value if set, zero value otherwise.
func (o *EntitlementDto) GetSourceSchemaObjectType() string {
	if o == nil || IsNil(o.SourceSchemaObjectType) {
		var ret string
		return ret
	}
	return *o.SourceSchemaObjectType
}

// GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetSourceSchemaObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceSchemaObjectType) {
		return nil, false
	}
	return o.SourceSchemaObjectType, true
}

// HasSourceSchemaObjectType returns a boolean if a field has been set.
func (o *EntitlementDto) HasSourceSchemaObjectType() bool {
	if o != nil && !IsNil(o.SourceSchemaObjectType) {
		return true
	}

	return false
}

// SetSourceSchemaObjectType gets a reference to the given string and assigns it to the SourceSchemaObjectType field.
func (o *EntitlementDto) SetSourceSchemaObjectType(v string) {
	o.SourceSchemaObjectType = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *EntitlementDto) GetPrivileged() bool {
	if o == nil || IsNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetPrivilegedOk() (*bool, bool) {
	if o == nil || IsNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *EntitlementDto) HasPrivileged() bool {
	if o != nil && !IsNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *EntitlementDto) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetCloudGoverned returns the CloudGoverned field value if set, zero value otherwise.
func (o *EntitlementDto) GetCloudGoverned() bool {
	if o == nil || IsNil(o.CloudGoverned) {
		var ret bool
		return ret
	}
	return *o.CloudGoverned
}

// GetCloudGovernedOk returns a tuple with the CloudGoverned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetCloudGovernedOk() (*bool, bool) {
	if o == nil || IsNil(o.CloudGoverned) {
		return nil, false
	}
	return o.CloudGoverned, true
}

// HasCloudGoverned returns a boolean if a field has been set.
func (o *EntitlementDto) HasCloudGoverned() bool {
	if o != nil && !IsNil(o.CloudGoverned) {
		return true
	}

	return false
}

// SetCloudGoverned gets a reference to the given bool and assigns it to the CloudGoverned field.
func (o *EntitlementDto) SetCloudGoverned(v bool) {
	o.CloudGoverned = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntitlementDto) GetSource() EntitlementSource {
	if o == nil || IsNil(o.Source) {
		var ret EntitlementSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDto) GetSourceOk() (*EntitlementSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntitlementDto) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given EntitlementSource and assigns it to the Source field.
func (o *EntitlementDto) SetSource(v EntitlementSource) {
	o.Source = &v
}

func (o EntitlementDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.SourceSchemaObjectType) {
		toSerialize["sourceSchemaObjectType"] = o.SourceSchemaObjectType
	}
	if !IsNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !IsNil(o.CloudGoverned) {
		toSerialize["cloudGoverned"] = o.CloudGoverned
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varEntitlementDto := _EntitlementDto{}

	err = json.Unmarshal(data, &varEntitlementDto)

	if err != nil {
		return err
	}

	*o = EntitlementDto(varEntitlementDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "sourceSchemaObjectType")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "cloudGoverned")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDto struct {
	value *EntitlementDto
	isSet bool
}

func (v NullableEntitlementDto) Get() *EntitlementDto {
	return v.value
}

func (v *NullableEntitlementDto) Set(val *EntitlementDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDto(val *EntitlementDto) *NullableEntitlementDto {
	return &NullableEntitlementDto{value: val, isSet: true}
}

func (v NullableEntitlementDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


