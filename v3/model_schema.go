/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the Schema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schema{}

// Schema struct for Schema
type Schema struct {
	// The id of the Schema.
	Id *string `json:"id,omitempty"`
	// The name of the Schema.
	Name *string `json:"name,omitempty"`
	// The name of the object type on the native system that the schema represents.
	NativeObjectType *string `json:"nativeObjectType,omitempty"`
	// The name of the attribute used to calculate the unique identifier for an object in the schema.
	IdentityAttribute *string `json:"identityAttribute,omitempty"`
	// The name of the attribute used to calculate the display value for an object in the schema.
	DisplayAttribute *string `json:"displayAttribute,omitempty"`
	// The name of the attribute whose values represent other objects in a hierarchy. Only relevant to group schemas.
	HierarchyAttribute *string `json:"hierarchyAttribute,omitempty"`
	// Flag indicating whether or not the include permissions with the object data when aggregating the schema.
	IncludePermissions *bool `json:"includePermissions,omitempty"`
	// The features that the schema supports.
	Features []SourceFeature `json:"features,omitempty"`
	// Holds any extra configuration data that the schema may require.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// The attribute definitions which form the schema.
	Attributes []AttributeDefinition `json:"attributes,omitempty"`
	// The date the Schema was created.
	Created *time.Time `json:"created,omitempty"`
	// The date the Schema was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schema Schema

// NewSchema instantiates a new Schema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchema() *Schema {
	this := Schema{}
	return &this
}

// NewSchemaWithDefaults instantiates a new Schema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaWithDefaults() *Schema {
	this := Schema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Schema) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Schema) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Schema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Schema) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Schema) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Schema) SetName(v string) {
	o.Name = &v
}

// GetNativeObjectType returns the NativeObjectType field value if set, zero value otherwise.
func (o *Schema) GetNativeObjectType() string {
	if o == nil || isNil(o.NativeObjectType) {
		var ret string
		return ret
	}
	return *o.NativeObjectType
}

// GetNativeObjectTypeOk returns a tuple with the NativeObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetNativeObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.NativeObjectType) {
		return nil, false
	}
	return o.NativeObjectType, true
}

// HasNativeObjectType returns a boolean if a field has been set.
func (o *Schema) HasNativeObjectType() bool {
	if o != nil && !isNil(o.NativeObjectType) {
		return true
	}

	return false
}

// SetNativeObjectType gets a reference to the given string and assigns it to the NativeObjectType field.
func (o *Schema) SetNativeObjectType(v string) {
	o.NativeObjectType = &v
}

// GetIdentityAttribute returns the IdentityAttribute field value if set, zero value otherwise.
func (o *Schema) GetIdentityAttribute() string {
	if o == nil || isNil(o.IdentityAttribute) {
		var ret string
		return ret
	}
	return *o.IdentityAttribute
}

// GetIdentityAttributeOk returns a tuple with the IdentityAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetIdentityAttributeOk() (*string, bool) {
	if o == nil || isNil(o.IdentityAttribute) {
		return nil, false
	}
	return o.IdentityAttribute, true
}

// HasIdentityAttribute returns a boolean if a field has been set.
func (o *Schema) HasIdentityAttribute() bool {
	if o != nil && !isNil(o.IdentityAttribute) {
		return true
	}

	return false
}

// SetIdentityAttribute gets a reference to the given string and assigns it to the IdentityAttribute field.
func (o *Schema) SetIdentityAttribute(v string) {
	o.IdentityAttribute = &v
}

// GetDisplayAttribute returns the DisplayAttribute field value if set, zero value otherwise.
func (o *Schema) GetDisplayAttribute() string {
	if o == nil || isNil(o.DisplayAttribute) {
		var ret string
		return ret
	}
	return *o.DisplayAttribute
}

// GetDisplayAttributeOk returns a tuple with the DisplayAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetDisplayAttributeOk() (*string, bool) {
	if o == nil || isNil(o.DisplayAttribute) {
		return nil, false
	}
	return o.DisplayAttribute, true
}

// HasDisplayAttribute returns a boolean if a field has been set.
func (o *Schema) HasDisplayAttribute() bool {
	if o != nil && !isNil(o.DisplayAttribute) {
		return true
	}

	return false
}

// SetDisplayAttribute gets a reference to the given string and assigns it to the DisplayAttribute field.
func (o *Schema) SetDisplayAttribute(v string) {
	o.DisplayAttribute = &v
}

// GetHierarchyAttribute returns the HierarchyAttribute field value if set, zero value otherwise.
func (o *Schema) GetHierarchyAttribute() string {
	if o == nil || isNil(o.HierarchyAttribute) {
		var ret string
		return ret
	}
	return *o.HierarchyAttribute
}

// GetHierarchyAttributeOk returns a tuple with the HierarchyAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetHierarchyAttributeOk() (*string, bool) {
	if o == nil || isNil(o.HierarchyAttribute) {
		return nil, false
	}
	return o.HierarchyAttribute, true
}

// HasHierarchyAttribute returns a boolean if a field has been set.
func (o *Schema) HasHierarchyAttribute() bool {
	if o != nil && !isNil(o.HierarchyAttribute) {
		return true
	}

	return false
}

// SetHierarchyAttribute gets a reference to the given string and assigns it to the HierarchyAttribute field.
func (o *Schema) SetHierarchyAttribute(v string) {
	o.HierarchyAttribute = &v
}

// GetIncludePermissions returns the IncludePermissions field value if set, zero value otherwise.
func (o *Schema) GetIncludePermissions() bool {
	if o == nil || isNil(o.IncludePermissions) {
		var ret bool
		return ret
	}
	return *o.IncludePermissions
}

// GetIncludePermissionsOk returns a tuple with the IncludePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetIncludePermissionsOk() (*bool, bool) {
	if o == nil || isNil(o.IncludePermissions) {
		return nil, false
	}
	return o.IncludePermissions, true
}

// HasIncludePermissions returns a boolean if a field has been set.
func (o *Schema) HasIncludePermissions() bool {
	if o != nil && !isNil(o.IncludePermissions) {
		return true
	}

	return false
}

// SetIncludePermissions gets a reference to the given bool and assigns it to the IncludePermissions field.
func (o *Schema) SetIncludePermissions(v bool) {
	o.IncludePermissions = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Schema) GetFeatures() []SourceFeature {
	if o == nil || isNil(o.Features) {
		var ret []SourceFeature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetFeaturesOk() ([]SourceFeature, bool) {
	if o == nil || isNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Schema) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []SourceFeature and assigns it to the Features field.
func (o *Schema) SetFeatures(v []SourceFeature) {
	o.Features = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Schema) GetConfiguration() map[string]interface{} {
	if o == nil || isNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Schema) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *Schema) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Schema) GetAttributes() []AttributeDefinition {
	if o == nil || isNil(o.Attributes) {
		var ret []AttributeDefinition
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetAttributesOk() ([]AttributeDefinition, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Schema) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []AttributeDefinition and assigns it to the Attributes field.
func (o *Schema) SetAttributes(v []AttributeDefinition) {
	o.Attributes = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Schema) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Schema) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Schema) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Schema) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Schema) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Schema) SetModified(v time.Time) {
	o.Modified = &v
}

func (o Schema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NativeObjectType) {
		toSerialize["nativeObjectType"] = o.NativeObjectType
	}
	if !isNil(o.IdentityAttribute) {
		toSerialize["identityAttribute"] = o.IdentityAttribute
	}
	if !isNil(o.DisplayAttribute) {
		toSerialize["displayAttribute"] = o.DisplayAttribute
	}
	if !isNil(o.HierarchyAttribute) {
		toSerialize["hierarchyAttribute"] = o.HierarchyAttribute
	}
	if !isNil(o.IncludePermissions) {
		toSerialize["includePermissions"] = o.IncludePermissions
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schema) UnmarshalJSON(bytes []byte) (err error) {
	varSchema := _Schema{}

	if err = json.Unmarshal(bytes, &varSchema); err == nil {
		*o = Schema(varSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nativeObjectType")
		delete(additionalProperties, "identityAttribute")
		delete(additionalProperties, "displayAttribute")
		delete(additionalProperties, "hierarchyAttribute")
		delete(additionalProperties, "includePermissions")
		delete(additionalProperties, "features")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchema struct {
	value *Schema
	isSet bool
}

func (v NullableSchema) Get() *Schema {
	return v.value
}

func (v *NullableSchema) Set(val *Schema) {
	v.value = val
	v.isSet = true
}

func (v NullableSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchema(val *Schema) *NullableSchema {
	return &NullableSchema{value: val, isSet: true}
}

func (v NullableSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


