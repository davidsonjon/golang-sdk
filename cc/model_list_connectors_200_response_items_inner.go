/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the ListConnectors200ResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectors200ResponseItemsInner{}

// ListConnectors200ResponseItemsInner struct for ListConnectors200ResponseItemsInner
type ListConnectors200ResponseItemsInner struct {
	ApplicationXml NullableString `json:"applicationXml,omitempty"`
	ClassName NullableString `json:"className,omitempty"`
	ConnectorMetadata map[string]interface{} `json:"connectorMetadata,omitempty"`
	CorrelationConfigXml NullableString `json:"correlationConfigXml,omitempty"`
	DirectConnect *bool `json:"directConnect,omitempty"`
	FileUpload *bool `json:"fileUpload,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	S3Location NullableString `json:"s3Location,omitempty"`
	Scope *string `json:"scope,omitempty"`
	ScriptName *string `json:"scriptName,omitempty"`
	SourceConfig NullableString `json:"sourceConfig,omitempty"`
	SourceConfigFrom NullableString `json:"sourceConfigFrom,omitempty"`
	SourceConfigXml NullableString `json:"sourceConfigXml,omitempty"`
	Status *string `json:"status,omitempty"`
	TranslationProperties map[string]interface{} `json:"translationProperties,omitempty"`
	Type *string `json:"type,omitempty"`
	UploadedFiles []map[string]interface{} `json:"uploadedFiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListConnectors200ResponseItemsInner ListConnectors200ResponseItemsInner

// NewListConnectors200ResponseItemsInner instantiates a new ListConnectors200ResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectors200ResponseItemsInner() *ListConnectors200ResponseItemsInner {
	this := ListConnectors200ResponseItemsInner{}
	return &this
}

// NewListConnectors200ResponseItemsInnerWithDefaults instantiates a new ListConnectors200ResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectors200ResponseItemsInnerWithDefaults() *ListConnectors200ResponseItemsInner {
	this := ListConnectors200ResponseItemsInner{}
	return &this
}

// GetApplicationXml returns the ApplicationXml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetApplicationXml() string {
	if o == nil || isNil(o.ApplicationXml.Get()) {
		var ret string
		return ret
	}
	return *o.ApplicationXml.Get()
}

// GetApplicationXmlOk returns a tuple with the ApplicationXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetApplicationXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationXml.Get(), o.ApplicationXml.IsSet()
}

// HasApplicationXml returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasApplicationXml() bool {
	if o != nil && o.ApplicationXml.IsSet() {
		return true
	}

	return false
}

// SetApplicationXml gets a reference to the given NullableString and assigns it to the ApplicationXml field.
func (o *ListConnectors200ResponseItemsInner) SetApplicationXml(v string) {
	o.ApplicationXml.Set(&v)
}
// SetApplicationXmlNil sets the value for ApplicationXml to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetApplicationXmlNil() {
	o.ApplicationXml.Set(nil)
}

// UnsetApplicationXml ensures that no value is present for ApplicationXml, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetApplicationXml() {
	o.ApplicationXml.Unset()
}

// GetClassName returns the ClassName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetClassName() string {
	if o == nil || isNil(o.ClassName.Get()) {
		var ret string
		return ret
	}
	return *o.ClassName.Get()
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClassName.Get(), o.ClassName.IsSet()
}

// HasClassName returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasClassName() bool {
	if o != nil && o.ClassName.IsSet() {
		return true
	}

	return false
}

// SetClassName gets a reference to the given NullableString and assigns it to the ClassName field.
func (o *ListConnectors200ResponseItemsInner) SetClassName(v string) {
	o.ClassName.Set(&v)
}
// SetClassNameNil sets the value for ClassName to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetClassNameNil() {
	o.ClassName.Set(nil)
}

// UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetClassName() {
	o.ClassName.Unset()
}

// GetConnectorMetadata returns the ConnectorMetadata field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetConnectorMetadata() map[string]interface{} {
	if o == nil || isNil(o.ConnectorMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectorMetadata
}

// GetConnectorMetadataOk returns a tuple with the ConnectorMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetConnectorMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ConnectorMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ConnectorMetadata, true
}

// HasConnectorMetadata returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasConnectorMetadata() bool {
	if o != nil && !isNil(o.ConnectorMetadata) {
		return true
	}

	return false
}

// SetConnectorMetadata gets a reference to the given map[string]interface{} and assigns it to the ConnectorMetadata field.
func (o *ListConnectors200ResponseItemsInner) SetConnectorMetadata(v map[string]interface{}) {
	o.ConnectorMetadata = v
}

// GetCorrelationConfigXml returns the CorrelationConfigXml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetCorrelationConfigXml() string {
	if o == nil || isNil(o.CorrelationConfigXml.Get()) {
		var ret string
		return ret
	}
	return *o.CorrelationConfigXml.Get()
}

// GetCorrelationConfigXmlOk returns a tuple with the CorrelationConfigXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetCorrelationConfigXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrelationConfigXml.Get(), o.CorrelationConfigXml.IsSet()
}

// HasCorrelationConfigXml returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasCorrelationConfigXml() bool {
	if o != nil && o.CorrelationConfigXml.IsSet() {
		return true
	}

	return false
}

// SetCorrelationConfigXml gets a reference to the given NullableString and assigns it to the CorrelationConfigXml field.
func (o *ListConnectors200ResponseItemsInner) SetCorrelationConfigXml(v string) {
	o.CorrelationConfigXml.Set(&v)
}
// SetCorrelationConfigXmlNil sets the value for CorrelationConfigXml to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetCorrelationConfigXmlNil() {
	o.CorrelationConfigXml.Set(nil)
}

// UnsetCorrelationConfigXml ensures that no value is present for CorrelationConfigXml, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetCorrelationConfigXml() {
	o.CorrelationConfigXml.Unset()
}

// GetDirectConnect returns the DirectConnect field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetDirectConnect() bool {
	if o == nil || isNil(o.DirectConnect) {
		var ret bool
		return ret
	}
	return *o.DirectConnect
}

// GetDirectConnectOk returns a tuple with the DirectConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetDirectConnectOk() (*bool, bool) {
	if o == nil || isNil(o.DirectConnect) {
		return nil, false
	}
	return o.DirectConnect, true
}

// HasDirectConnect returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasDirectConnect() bool {
	if o != nil && !isNil(o.DirectConnect) {
		return true
	}

	return false
}

// SetDirectConnect gets a reference to the given bool and assigns it to the DirectConnect field.
func (o *ListConnectors200ResponseItemsInner) SetDirectConnect(v bool) {
	o.DirectConnect = &v
}

// GetFileUpload returns the FileUpload field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetFileUpload() bool {
	if o == nil || isNil(o.FileUpload) {
		var ret bool
		return ret
	}
	return *o.FileUpload
}

// GetFileUploadOk returns a tuple with the FileUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetFileUploadOk() (*bool, bool) {
	if o == nil || isNil(o.FileUpload) {
		return nil, false
	}
	return o.FileUpload, true
}

// HasFileUpload returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasFileUpload() bool {
	if o != nil && !isNil(o.FileUpload) {
		return true
	}

	return false
}

// SetFileUpload gets a reference to the given bool and assigns it to the FileUpload field.
func (o *ListConnectors200ResponseItemsInner) SetFileUpload(v bool) {
	o.FileUpload = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListConnectors200ResponseItemsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListConnectors200ResponseItemsInner) SetName(v string) {
	o.Name = &v
}

// GetS3Location returns the S3Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetS3Location() string {
	if o == nil || isNil(o.S3Location.Get()) {
		var ret string
		return ret
	}
	return *o.S3Location.Get()
}

// GetS3LocationOk returns a tuple with the S3Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetS3LocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Location.Get(), o.S3Location.IsSet()
}

// HasS3Location returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasS3Location() bool {
	if o != nil && o.S3Location.IsSet() {
		return true
	}

	return false
}

// SetS3Location gets a reference to the given NullableString and assigns it to the S3Location field.
func (o *ListConnectors200ResponseItemsInner) SetS3Location(v string) {
	o.S3Location.Set(&v)
}
// SetS3LocationNil sets the value for S3Location to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetS3LocationNil() {
	o.S3Location.Set(nil)
}

// UnsetS3Location ensures that no value is present for S3Location, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetS3Location() {
	o.S3Location.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ListConnectors200ResponseItemsInner) SetScope(v string) {
	o.Scope = &v
}

// GetScriptName returns the ScriptName field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetScriptName() string {
	if o == nil || isNil(o.ScriptName) {
		var ret string
		return ret
	}
	return *o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetScriptNameOk() (*string, bool) {
	if o == nil || isNil(o.ScriptName) {
		return nil, false
	}
	return o.ScriptName, true
}

// HasScriptName returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasScriptName() bool {
	if o != nil && !isNil(o.ScriptName) {
		return true
	}

	return false
}

// SetScriptName gets a reference to the given string and assigns it to the ScriptName field.
func (o *ListConnectors200ResponseItemsInner) SetScriptName(v string) {
	o.ScriptName = &v
}

// GetSourceConfig returns the SourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetSourceConfig() string {
	if o == nil || isNil(o.SourceConfig.Get()) {
		var ret string
		return ret
	}
	return *o.SourceConfig.Get()
}

// GetSourceConfigOk returns a tuple with the SourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetSourceConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceConfig.Get(), o.SourceConfig.IsSet()
}

// HasSourceConfig returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasSourceConfig() bool {
	if o != nil && o.SourceConfig.IsSet() {
		return true
	}

	return false
}

// SetSourceConfig gets a reference to the given NullableString and assigns it to the SourceConfig field.
func (o *ListConnectors200ResponseItemsInner) SetSourceConfig(v string) {
	o.SourceConfig.Set(&v)
}
// SetSourceConfigNil sets the value for SourceConfig to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetSourceConfigNil() {
	o.SourceConfig.Set(nil)
}

// UnsetSourceConfig ensures that no value is present for SourceConfig, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetSourceConfig() {
	o.SourceConfig.Unset()
}

// GetSourceConfigFrom returns the SourceConfigFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetSourceConfigFrom() string {
	if o == nil || isNil(o.SourceConfigFrom.Get()) {
		var ret string
		return ret
	}
	return *o.SourceConfigFrom.Get()
}

// GetSourceConfigFromOk returns a tuple with the SourceConfigFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetSourceConfigFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceConfigFrom.Get(), o.SourceConfigFrom.IsSet()
}

// HasSourceConfigFrom returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasSourceConfigFrom() bool {
	if o != nil && o.SourceConfigFrom.IsSet() {
		return true
	}

	return false
}

// SetSourceConfigFrom gets a reference to the given NullableString and assigns it to the SourceConfigFrom field.
func (o *ListConnectors200ResponseItemsInner) SetSourceConfigFrom(v string) {
	o.SourceConfigFrom.Set(&v)
}
// SetSourceConfigFromNil sets the value for SourceConfigFrom to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetSourceConfigFromNil() {
	o.SourceConfigFrom.Set(nil)
}

// UnsetSourceConfigFrom ensures that no value is present for SourceConfigFrom, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetSourceConfigFrom() {
	o.SourceConfigFrom.Unset()
}

// GetSourceConfigXml returns the SourceConfigXml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseItemsInner) GetSourceConfigXml() string {
	if o == nil || isNil(o.SourceConfigXml.Get()) {
		var ret string
		return ret
	}
	return *o.SourceConfigXml.Get()
}

// GetSourceConfigXmlOk returns a tuple with the SourceConfigXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseItemsInner) GetSourceConfigXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceConfigXml.Get(), o.SourceConfigXml.IsSet()
}

// HasSourceConfigXml returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasSourceConfigXml() bool {
	if o != nil && o.SourceConfigXml.IsSet() {
		return true
	}

	return false
}

// SetSourceConfigXml gets a reference to the given NullableString and assigns it to the SourceConfigXml field.
func (o *ListConnectors200ResponseItemsInner) SetSourceConfigXml(v string) {
	o.SourceConfigXml.Set(&v)
}
// SetSourceConfigXmlNil sets the value for SourceConfigXml to be an explicit nil
func (o *ListConnectors200ResponseItemsInner) SetSourceConfigXmlNil() {
	o.SourceConfigXml.Set(nil)
}

// UnsetSourceConfigXml ensures that no value is present for SourceConfigXml, not even an explicit nil
func (o *ListConnectors200ResponseItemsInner) UnsetSourceConfigXml() {
	o.SourceConfigXml.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListConnectors200ResponseItemsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTranslationProperties returns the TranslationProperties field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetTranslationProperties() map[string]interface{} {
	if o == nil || isNil(o.TranslationProperties) {
		var ret map[string]interface{}
		return ret
	}
	return o.TranslationProperties
}

// GetTranslationPropertiesOk returns a tuple with the TranslationProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetTranslationPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.TranslationProperties) {
		return map[string]interface{}{}, false
	}
	return o.TranslationProperties, true
}

// HasTranslationProperties returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasTranslationProperties() bool {
	if o != nil && !isNil(o.TranslationProperties) {
		return true
	}

	return false
}

// SetTranslationProperties gets a reference to the given map[string]interface{} and assigns it to the TranslationProperties field.
func (o *ListConnectors200ResponseItemsInner) SetTranslationProperties(v map[string]interface{}) {
	o.TranslationProperties = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListConnectors200ResponseItemsInner) SetType(v string) {
	o.Type = &v
}

// GetUploadedFiles returns the UploadedFiles field value if set, zero value otherwise.
func (o *ListConnectors200ResponseItemsInner) GetUploadedFiles() []map[string]interface{} {
	if o == nil || isNil(o.UploadedFiles) {
		var ret []map[string]interface{}
		return ret
	}
	return o.UploadedFiles
}

// GetUploadedFilesOk returns a tuple with the UploadedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseItemsInner) GetUploadedFilesOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.UploadedFiles) {
		return nil, false
	}
	return o.UploadedFiles, true
}

// HasUploadedFiles returns a boolean if a field has been set.
func (o *ListConnectors200ResponseItemsInner) HasUploadedFiles() bool {
	if o != nil && !isNil(o.UploadedFiles) {
		return true
	}

	return false
}

// SetUploadedFiles gets a reference to the given []map[string]interface{} and assigns it to the UploadedFiles field.
func (o *ListConnectors200ResponseItemsInner) SetUploadedFiles(v []map[string]interface{}) {
	o.UploadedFiles = v
}

func (o ListConnectors200ResponseItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectors200ResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationXml.IsSet() {
		toSerialize["applicationXml"] = o.ApplicationXml.Get()
	}
	if o.ClassName.IsSet() {
		toSerialize["className"] = o.ClassName.Get()
	}
	if !isNil(o.ConnectorMetadata) {
		toSerialize["connectorMetadata"] = o.ConnectorMetadata
	}
	if o.CorrelationConfigXml.IsSet() {
		toSerialize["correlationConfigXml"] = o.CorrelationConfigXml.Get()
	}
	if !isNil(o.DirectConnect) {
		toSerialize["directConnect"] = o.DirectConnect
	}
	if !isNil(o.FileUpload) {
		toSerialize["fileUpload"] = o.FileUpload
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.S3Location.IsSet() {
		toSerialize["s3Location"] = o.S3Location.Get()
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.ScriptName) {
		toSerialize["scriptName"] = o.ScriptName
	}
	if o.SourceConfig.IsSet() {
		toSerialize["sourceConfig"] = o.SourceConfig.Get()
	}
	if o.SourceConfigFrom.IsSet() {
		toSerialize["sourceConfigFrom"] = o.SourceConfigFrom.Get()
	}
	if o.SourceConfigXml.IsSet() {
		toSerialize["sourceConfigXml"] = o.SourceConfigXml.Get()
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TranslationProperties) {
		toSerialize["translationProperties"] = o.TranslationProperties
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UploadedFiles) {
		toSerialize["uploadedFiles"] = o.UploadedFiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListConnectors200ResponseItemsInner) UnmarshalJSON(bytes []byte) (err error) {
	varListConnectors200ResponseItemsInner := _ListConnectors200ResponseItemsInner{}

	if err = json.Unmarshal(bytes, &varListConnectors200ResponseItemsInner); err == nil {
	*o = ListConnectors200ResponseItemsInner(varListConnectors200ResponseItemsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "applicationXml")
		delete(additionalProperties, "className")
		delete(additionalProperties, "connectorMetadata")
		delete(additionalProperties, "correlationConfigXml")
		delete(additionalProperties, "directConnect")
		delete(additionalProperties, "fileUpload")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "s3Location")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "scriptName")
		delete(additionalProperties, "sourceConfig")
		delete(additionalProperties, "sourceConfigFrom")
		delete(additionalProperties, "sourceConfigXml")
		delete(additionalProperties, "status")
		delete(additionalProperties, "translationProperties")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uploadedFiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListConnectors200ResponseItemsInner struct {
	value *ListConnectors200ResponseItemsInner
	isSet bool
}

func (v NullableListConnectors200ResponseItemsInner) Get() *ListConnectors200ResponseItemsInner {
	return v.value
}

func (v *NullableListConnectors200ResponseItemsInner) Set(val *ListConnectors200ResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectors200ResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectors200ResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectors200ResponseItemsInner(val *ListConnectors200ResponseItemsInner) *NullableListConnectors200ResponseItemsInner {
	return &NullableListConnectors200ResponseItemsInner{value: val, isSet: true}
}

func (v NullableListConnectors200ResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectors200ResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


