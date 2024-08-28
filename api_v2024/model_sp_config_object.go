/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the SpConfigObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpConfigObject{}

// SpConfigObject Response model for get object configuration.
type SpConfigObject struct {
	// The object type this configuration is for.
	ObjectType *string `json:"objectType,omitempty"`
	ResolveByIdUrl *SpConfigUrl `json:"resolveByIdUrl,omitempty"`
	// Url and query parameters to be used to resolve this type of object by name.
	ResolveByNameUrl []SpConfigUrl `json:"resolveByNameUrl,omitempty"`
	ExportUrl *SpConfigUrl `json:"exportUrl,omitempty"`
	// Rights needed by the invoker of sp-config/export in order to export this type of object.
	ExportRight *string `json:"exportRight,omitempty"`
	// Pagination limit imposed by the target service for this object type.
	ExportLimit *int32 `json:"exportLimit,omitempty"`
	ImportUrl *SpConfigUrl `json:"importUrl,omitempty"`
	// Rights needed by the invoker of sp-config/import in order to import this type of object.
	ImportRight *string `json:"importRight,omitempty"`
	// Pagination limit imposed by the target service for this object type.
	ImportLimit *int32 `json:"importLimit,omitempty"`
	// List of json paths within an exported object of this type that represent references that need to be resolved.
	ReferenceExtractors []string `json:"referenceExtractors,omitempty"`
	// If true, this type of object will be JWS signed and cannot be modified before import.
	SignatureRequired *bool `json:"signatureRequired,omitempty"`
	LegacyObject *bool `json:"legacyObject,omitempty"`
	OnePerTenant *bool `json:"onePerTenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpConfigObject SpConfigObject

// NewSpConfigObject instantiates a new SpConfigObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpConfigObject() *SpConfigObject {
	this := SpConfigObject{}
	var signatureRequired bool = false
	this.SignatureRequired = &signatureRequired
	var legacyObject bool = false
	this.LegacyObject = &legacyObject
	var onePerTenant bool = false
	this.OnePerTenant = &onePerTenant
	return &this
}

// NewSpConfigObjectWithDefaults instantiates a new SpConfigObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpConfigObjectWithDefaults() *SpConfigObject {
	this := SpConfigObject{}
	var signatureRequired bool = false
	this.SignatureRequired = &signatureRequired
	var legacyObject bool = false
	this.LegacyObject = &legacyObject
	var onePerTenant bool = false
	this.OnePerTenant = &onePerTenant
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *SpConfigObject) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *SpConfigObject) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *SpConfigObject) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetResolveByIdUrl returns the ResolveByIdUrl field value if set, zero value otherwise.
func (o *SpConfigObject) GetResolveByIdUrl() SpConfigUrl {
	if o == nil || IsNil(o.ResolveByIdUrl) {
		var ret SpConfigUrl
		return ret
	}
	return *o.ResolveByIdUrl
}

// GetResolveByIdUrlOk returns a tuple with the ResolveByIdUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetResolveByIdUrlOk() (*SpConfigUrl, bool) {
	if o == nil || IsNil(o.ResolveByIdUrl) {
		return nil, false
	}
	return o.ResolveByIdUrl, true
}

// HasResolveByIdUrl returns a boolean if a field has been set.
func (o *SpConfigObject) HasResolveByIdUrl() bool {
	if o != nil && !IsNil(o.ResolveByIdUrl) {
		return true
	}

	return false
}

// SetResolveByIdUrl gets a reference to the given SpConfigUrl and assigns it to the ResolveByIdUrl field.
func (o *SpConfigObject) SetResolveByIdUrl(v SpConfigUrl) {
	o.ResolveByIdUrl = &v
}

// GetResolveByNameUrl returns the ResolveByNameUrl field value if set, zero value otherwise.
func (o *SpConfigObject) GetResolveByNameUrl() []SpConfigUrl {
	if o == nil || IsNil(o.ResolveByNameUrl) {
		var ret []SpConfigUrl
		return ret
	}
	return o.ResolveByNameUrl
}

// GetResolveByNameUrlOk returns a tuple with the ResolveByNameUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetResolveByNameUrlOk() ([]SpConfigUrl, bool) {
	if o == nil || IsNil(o.ResolveByNameUrl) {
		return nil, false
	}
	return o.ResolveByNameUrl, true
}

// HasResolveByNameUrl returns a boolean if a field has been set.
func (o *SpConfigObject) HasResolveByNameUrl() bool {
	if o != nil && !IsNil(o.ResolveByNameUrl) {
		return true
	}

	return false
}

// SetResolveByNameUrl gets a reference to the given []SpConfigUrl and assigns it to the ResolveByNameUrl field.
func (o *SpConfigObject) SetResolveByNameUrl(v []SpConfigUrl) {
	o.ResolveByNameUrl = v
}

// GetExportUrl returns the ExportUrl field value if set, zero value otherwise.
func (o *SpConfigObject) GetExportUrl() SpConfigUrl {
	if o == nil || IsNil(o.ExportUrl) {
		var ret SpConfigUrl
		return ret
	}
	return *o.ExportUrl
}

// GetExportUrlOk returns a tuple with the ExportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetExportUrlOk() (*SpConfigUrl, bool) {
	if o == nil || IsNil(o.ExportUrl) {
		return nil, false
	}
	return o.ExportUrl, true
}

// HasExportUrl returns a boolean if a field has been set.
func (o *SpConfigObject) HasExportUrl() bool {
	if o != nil && !IsNil(o.ExportUrl) {
		return true
	}

	return false
}

// SetExportUrl gets a reference to the given SpConfigUrl and assigns it to the ExportUrl field.
func (o *SpConfigObject) SetExportUrl(v SpConfigUrl) {
	o.ExportUrl = &v
}

// GetExportRight returns the ExportRight field value if set, zero value otherwise.
func (o *SpConfigObject) GetExportRight() string {
	if o == nil || IsNil(o.ExportRight) {
		var ret string
		return ret
	}
	return *o.ExportRight
}

// GetExportRightOk returns a tuple with the ExportRight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetExportRightOk() (*string, bool) {
	if o == nil || IsNil(o.ExportRight) {
		return nil, false
	}
	return o.ExportRight, true
}

// HasExportRight returns a boolean if a field has been set.
func (o *SpConfigObject) HasExportRight() bool {
	if o != nil && !IsNil(o.ExportRight) {
		return true
	}

	return false
}

// SetExportRight gets a reference to the given string and assigns it to the ExportRight field.
func (o *SpConfigObject) SetExportRight(v string) {
	o.ExportRight = &v
}

// GetExportLimit returns the ExportLimit field value if set, zero value otherwise.
func (o *SpConfigObject) GetExportLimit() int32 {
	if o == nil || IsNil(o.ExportLimit) {
		var ret int32
		return ret
	}
	return *o.ExportLimit
}

// GetExportLimitOk returns a tuple with the ExportLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetExportLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ExportLimit) {
		return nil, false
	}
	return o.ExportLimit, true
}

// HasExportLimit returns a boolean if a field has been set.
func (o *SpConfigObject) HasExportLimit() bool {
	if o != nil && !IsNil(o.ExportLimit) {
		return true
	}

	return false
}

// SetExportLimit gets a reference to the given int32 and assigns it to the ExportLimit field.
func (o *SpConfigObject) SetExportLimit(v int32) {
	o.ExportLimit = &v
}

// GetImportUrl returns the ImportUrl field value if set, zero value otherwise.
func (o *SpConfigObject) GetImportUrl() SpConfigUrl {
	if o == nil || IsNil(o.ImportUrl) {
		var ret SpConfigUrl
		return ret
	}
	return *o.ImportUrl
}

// GetImportUrlOk returns a tuple with the ImportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetImportUrlOk() (*SpConfigUrl, bool) {
	if o == nil || IsNil(o.ImportUrl) {
		return nil, false
	}
	return o.ImportUrl, true
}

// HasImportUrl returns a boolean if a field has been set.
func (o *SpConfigObject) HasImportUrl() bool {
	if o != nil && !IsNil(o.ImportUrl) {
		return true
	}

	return false
}

// SetImportUrl gets a reference to the given SpConfigUrl and assigns it to the ImportUrl field.
func (o *SpConfigObject) SetImportUrl(v SpConfigUrl) {
	o.ImportUrl = &v
}

// GetImportRight returns the ImportRight field value if set, zero value otherwise.
func (o *SpConfigObject) GetImportRight() string {
	if o == nil || IsNil(o.ImportRight) {
		var ret string
		return ret
	}
	return *o.ImportRight
}

// GetImportRightOk returns a tuple with the ImportRight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetImportRightOk() (*string, bool) {
	if o == nil || IsNil(o.ImportRight) {
		return nil, false
	}
	return o.ImportRight, true
}

// HasImportRight returns a boolean if a field has been set.
func (o *SpConfigObject) HasImportRight() bool {
	if o != nil && !IsNil(o.ImportRight) {
		return true
	}

	return false
}

// SetImportRight gets a reference to the given string and assigns it to the ImportRight field.
func (o *SpConfigObject) SetImportRight(v string) {
	o.ImportRight = &v
}

// GetImportLimit returns the ImportLimit field value if set, zero value otherwise.
func (o *SpConfigObject) GetImportLimit() int32 {
	if o == nil || IsNil(o.ImportLimit) {
		var ret int32
		return ret
	}
	return *o.ImportLimit
}

// GetImportLimitOk returns a tuple with the ImportLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetImportLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ImportLimit) {
		return nil, false
	}
	return o.ImportLimit, true
}

// HasImportLimit returns a boolean if a field has been set.
func (o *SpConfigObject) HasImportLimit() bool {
	if o != nil && !IsNil(o.ImportLimit) {
		return true
	}

	return false
}

// SetImportLimit gets a reference to the given int32 and assigns it to the ImportLimit field.
func (o *SpConfigObject) SetImportLimit(v int32) {
	o.ImportLimit = &v
}

// GetReferenceExtractors returns the ReferenceExtractors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpConfigObject) GetReferenceExtractors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ReferenceExtractors
}

// GetReferenceExtractorsOk returns a tuple with the ReferenceExtractors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpConfigObject) GetReferenceExtractorsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReferenceExtractors) {
		return nil, false
	}
	return o.ReferenceExtractors, true
}

// HasReferenceExtractors returns a boolean if a field has been set.
func (o *SpConfigObject) HasReferenceExtractors() bool {
	if o != nil && !IsNil(o.ReferenceExtractors) {
		return true
	}

	return false
}

// SetReferenceExtractors gets a reference to the given []string and assigns it to the ReferenceExtractors field.
func (o *SpConfigObject) SetReferenceExtractors(v []string) {
	o.ReferenceExtractors = v
}

// GetSignatureRequired returns the SignatureRequired field value if set, zero value otherwise.
func (o *SpConfigObject) GetSignatureRequired() bool {
	if o == nil || IsNil(o.SignatureRequired) {
		var ret bool
		return ret
	}
	return *o.SignatureRequired
}

// GetSignatureRequiredOk returns a tuple with the SignatureRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetSignatureRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SignatureRequired) {
		return nil, false
	}
	return o.SignatureRequired, true
}

// HasSignatureRequired returns a boolean if a field has been set.
func (o *SpConfigObject) HasSignatureRequired() bool {
	if o != nil && !IsNil(o.SignatureRequired) {
		return true
	}

	return false
}

// SetSignatureRequired gets a reference to the given bool and assigns it to the SignatureRequired field.
func (o *SpConfigObject) SetSignatureRequired(v bool) {
	o.SignatureRequired = &v
}

// GetLegacyObject returns the LegacyObject field value if set, zero value otherwise.
func (o *SpConfigObject) GetLegacyObject() bool {
	if o == nil || IsNil(o.LegacyObject) {
		var ret bool
		return ret
	}
	return *o.LegacyObject
}

// GetLegacyObjectOk returns a tuple with the LegacyObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetLegacyObjectOk() (*bool, bool) {
	if o == nil || IsNil(o.LegacyObject) {
		return nil, false
	}
	return o.LegacyObject, true
}

// HasLegacyObject returns a boolean if a field has been set.
func (o *SpConfigObject) HasLegacyObject() bool {
	if o != nil && !IsNil(o.LegacyObject) {
		return true
	}

	return false
}

// SetLegacyObject gets a reference to the given bool and assigns it to the LegacyObject field.
func (o *SpConfigObject) SetLegacyObject(v bool) {
	o.LegacyObject = &v
}

// GetOnePerTenant returns the OnePerTenant field value if set, zero value otherwise.
func (o *SpConfigObject) GetOnePerTenant() bool {
	if o == nil || IsNil(o.OnePerTenant) {
		var ret bool
		return ret
	}
	return *o.OnePerTenant
}

// GetOnePerTenantOk returns a tuple with the OnePerTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConfigObject) GetOnePerTenantOk() (*bool, bool) {
	if o == nil || IsNil(o.OnePerTenant) {
		return nil, false
	}
	return o.OnePerTenant, true
}

// HasOnePerTenant returns a boolean if a field has been set.
func (o *SpConfigObject) HasOnePerTenant() bool {
	if o != nil && !IsNil(o.OnePerTenant) {
		return true
	}

	return false
}

// SetOnePerTenant gets a reference to the given bool and assigns it to the OnePerTenant field.
func (o *SpConfigObject) SetOnePerTenant(v bool) {
	o.OnePerTenant = &v
}

func (o SpConfigObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpConfigObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.ResolveByIdUrl) {
		toSerialize["resolveByIdUrl"] = o.ResolveByIdUrl
	}
	if !IsNil(o.ResolveByNameUrl) {
		toSerialize["resolveByNameUrl"] = o.ResolveByNameUrl
	}
	if !IsNil(o.ExportUrl) {
		toSerialize["exportUrl"] = o.ExportUrl
	}
	if !IsNil(o.ExportRight) {
		toSerialize["exportRight"] = o.ExportRight
	}
	if !IsNil(o.ExportLimit) {
		toSerialize["exportLimit"] = o.ExportLimit
	}
	if !IsNil(o.ImportUrl) {
		toSerialize["importUrl"] = o.ImportUrl
	}
	if !IsNil(o.ImportRight) {
		toSerialize["importRight"] = o.ImportRight
	}
	if !IsNil(o.ImportLimit) {
		toSerialize["importLimit"] = o.ImportLimit
	}
	if o.ReferenceExtractors != nil {
		toSerialize["referenceExtractors"] = o.ReferenceExtractors
	}
	if !IsNil(o.SignatureRequired) {
		toSerialize["signatureRequired"] = o.SignatureRequired
	}
	if !IsNil(o.LegacyObject) {
		toSerialize["legacyObject"] = o.LegacyObject
	}
	if !IsNil(o.OnePerTenant) {
		toSerialize["onePerTenant"] = o.OnePerTenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpConfigObject) UnmarshalJSON(data []byte) (err error) {
	varSpConfigObject := _SpConfigObject{}

	err = json.Unmarshal(data, &varSpConfigObject)

	if err != nil {
		return err
	}

	*o = SpConfigObject(varSpConfigObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "resolveByIdUrl")
		delete(additionalProperties, "resolveByNameUrl")
		delete(additionalProperties, "exportUrl")
		delete(additionalProperties, "exportRight")
		delete(additionalProperties, "exportLimit")
		delete(additionalProperties, "importUrl")
		delete(additionalProperties, "importRight")
		delete(additionalProperties, "importLimit")
		delete(additionalProperties, "referenceExtractors")
		delete(additionalProperties, "signatureRequired")
		delete(additionalProperties, "legacyObject")
		delete(additionalProperties, "onePerTenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpConfigObject struct {
	value *SpConfigObject
	isSet bool
}

func (v NullableSpConfigObject) Get() *SpConfigObject {
	return v.value
}

func (v *NullableSpConfigObject) Set(val *SpConfigObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSpConfigObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSpConfigObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpConfigObject(val *SpConfigObject) *NullableSpConfigObject {
	return &NullableSpConfigObject{value: val, isSet: true}
}

func (v NullableSpConfigObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpConfigObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


