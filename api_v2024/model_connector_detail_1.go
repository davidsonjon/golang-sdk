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

// checks if the ConnectorDetail1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorDetail1{}

// ConnectorDetail1 struct for ConnectorDetail1
type ConnectorDetail1 struct {
	// The connector name
	Name *string `json:"name,omitempty"`
	// XML representation of the source config data
	SourceConfigXml *string `json:"sourceConfigXml,omitempty"`
	// JSON representation of the source config data
	SourceConfig *string `json:"sourceConfig,omitempty"`
	// true if the source is a direct connect source
	DirectConnect *bool `json:"directConnect,omitempty"`
	// Connector config's file upload attribute, false if not there
	FileUpload *bool `json:"fileUpload,omitempty"`
	// List of uploaded file strings for the connector
	UploadedFiles *string `json:"uploadedFiles,omitempty"`
	// Object containing metadata pertinent to the UI to be used
	ConnectorMetadata map[string]interface{} `json:"connectorMetadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorDetail1 ConnectorDetail1

// NewConnectorDetail1 instantiates a new ConnectorDetail1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorDetail1() *ConnectorDetail1 {
	this := ConnectorDetail1{}
	return &this
}

// NewConnectorDetail1WithDefaults instantiates a new ConnectorDetail1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorDetail1WithDefaults() *ConnectorDetail1 {
	this := ConnectorDetail1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorDetail1) SetName(v string) {
	o.Name = &v
}

// GetSourceConfigXml returns the SourceConfigXml field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetSourceConfigXml() string {
	if o == nil || isNil(o.SourceConfigXml) {
		var ret string
		return ret
	}
	return *o.SourceConfigXml
}

// GetSourceConfigXmlOk returns a tuple with the SourceConfigXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetSourceConfigXmlOk() (*string, bool) {
	if o == nil || isNil(o.SourceConfigXml) {
		return nil, false
	}
	return o.SourceConfigXml, true
}

// HasSourceConfigXml returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasSourceConfigXml() bool {
	if o != nil && !isNil(o.SourceConfigXml) {
		return true
	}

	return false
}

// SetSourceConfigXml gets a reference to the given string and assigns it to the SourceConfigXml field.
func (o *ConnectorDetail1) SetSourceConfigXml(v string) {
	o.SourceConfigXml = &v
}

// GetSourceConfig returns the SourceConfig field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetSourceConfig() string {
	if o == nil || isNil(o.SourceConfig) {
		var ret string
		return ret
	}
	return *o.SourceConfig
}

// GetSourceConfigOk returns a tuple with the SourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetSourceConfigOk() (*string, bool) {
	if o == nil || isNil(o.SourceConfig) {
		return nil, false
	}
	return o.SourceConfig, true
}

// HasSourceConfig returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasSourceConfig() bool {
	if o != nil && !isNil(o.SourceConfig) {
		return true
	}

	return false
}

// SetSourceConfig gets a reference to the given string and assigns it to the SourceConfig field.
func (o *ConnectorDetail1) SetSourceConfig(v string) {
	o.SourceConfig = &v
}

// GetDirectConnect returns the DirectConnect field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetDirectConnect() bool {
	if o == nil || isNil(o.DirectConnect) {
		var ret bool
		return ret
	}
	return *o.DirectConnect
}

// GetDirectConnectOk returns a tuple with the DirectConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetDirectConnectOk() (*bool, bool) {
	if o == nil || isNil(o.DirectConnect) {
		return nil, false
	}
	return o.DirectConnect, true
}

// HasDirectConnect returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasDirectConnect() bool {
	if o != nil && !isNil(o.DirectConnect) {
		return true
	}

	return false
}

// SetDirectConnect gets a reference to the given bool and assigns it to the DirectConnect field.
func (o *ConnectorDetail1) SetDirectConnect(v bool) {
	o.DirectConnect = &v
}

// GetFileUpload returns the FileUpload field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetFileUpload() bool {
	if o == nil || isNil(o.FileUpload) {
		var ret bool
		return ret
	}
	return *o.FileUpload
}

// GetFileUploadOk returns a tuple with the FileUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetFileUploadOk() (*bool, bool) {
	if o == nil || isNil(o.FileUpload) {
		return nil, false
	}
	return o.FileUpload, true
}

// HasFileUpload returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasFileUpload() bool {
	if o != nil && !isNil(o.FileUpload) {
		return true
	}

	return false
}

// SetFileUpload gets a reference to the given bool and assigns it to the FileUpload field.
func (o *ConnectorDetail1) SetFileUpload(v bool) {
	o.FileUpload = &v
}

// GetUploadedFiles returns the UploadedFiles field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetUploadedFiles() string {
	if o == nil || isNil(o.UploadedFiles) {
		var ret string
		return ret
	}
	return *o.UploadedFiles
}

// GetUploadedFilesOk returns a tuple with the UploadedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetUploadedFilesOk() (*string, bool) {
	if o == nil || isNil(o.UploadedFiles) {
		return nil, false
	}
	return o.UploadedFiles, true
}

// HasUploadedFiles returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasUploadedFiles() bool {
	if o != nil && !isNil(o.UploadedFiles) {
		return true
	}

	return false
}

// SetUploadedFiles gets a reference to the given string and assigns it to the UploadedFiles field.
func (o *ConnectorDetail1) SetUploadedFiles(v string) {
	o.UploadedFiles = &v
}

// GetConnectorMetadata returns the ConnectorMetadata field value if set, zero value otherwise.
func (o *ConnectorDetail1) GetConnectorMetadata() map[string]interface{} {
	if o == nil || isNil(o.ConnectorMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectorMetadata
}

// GetConnectorMetadataOk returns a tuple with the ConnectorMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDetail1) GetConnectorMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ConnectorMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ConnectorMetadata, true
}

// HasConnectorMetadata returns a boolean if a field has been set.
func (o *ConnectorDetail1) HasConnectorMetadata() bool {
	if o != nil && !isNil(o.ConnectorMetadata) {
		return true
	}

	return false
}

// SetConnectorMetadata gets a reference to the given map[string]interface{} and assigns it to the ConnectorMetadata field.
func (o *ConnectorDetail1) SetConnectorMetadata(v map[string]interface{}) {
	o.ConnectorMetadata = v
}

func (o ConnectorDetail1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorDetail1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SourceConfigXml) {
		toSerialize["sourceConfigXml"] = o.SourceConfigXml
	}
	if !isNil(o.SourceConfig) {
		toSerialize["sourceConfig"] = o.SourceConfig
	}
	if !isNil(o.DirectConnect) {
		toSerialize["directConnect"] = o.DirectConnect
	}
	if !isNil(o.FileUpload) {
		toSerialize["fileUpload"] = o.FileUpload
	}
	if !isNil(o.UploadedFiles) {
		toSerialize["uploadedFiles"] = o.UploadedFiles
	}
	if !isNil(o.ConnectorMetadata) {
		toSerialize["connectorMetadata"] = o.ConnectorMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorDetail1) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorDetail1 := _ConnectorDetail1{}

	if err = json.Unmarshal(bytes, &varConnectorDetail1); err == nil {
	*o = ConnectorDetail1(varConnectorDetail1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "sourceConfigXml")
		delete(additionalProperties, "sourceConfig")
		delete(additionalProperties, "directConnect")
		delete(additionalProperties, "fileUpload")
		delete(additionalProperties, "uploadedFiles")
		delete(additionalProperties, "connectorMetadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorDetail1 struct {
	value *ConnectorDetail1
	isSet bool
}

func (v NullableConnectorDetail1) Get() *ConnectorDetail1 {
	return v.value
}

func (v *NullableConnectorDetail1) Set(val *ConnectorDetail1) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorDetail1) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorDetail1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorDetail1(val *ConnectorDetail1) *NullableConnectorDetail1 {
	return &NullableConnectorDetail1{value: val, isSet: true}
}

func (v NullableConnectorDetail1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorDetail1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


