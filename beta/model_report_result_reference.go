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

// checks if the ReportResultReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportResultReference{}

// ReportResultReference struct for ReportResultReference
type ReportResultReference struct {
	Type *DtoType `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReportResultReference ReportResultReference

// NewReportResultReference instantiates a new ReportResultReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResultReference() *ReportResultReference {
	this := ReportResultReference{}
	return &this
}

// NewReportResultReferenceWithDefaults instantiates a new ReportResultReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResultReferenceWithDefaults() *ReportResultReference {
	this := ReportResultReference{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReportResultReference) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultReference) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReportResultReference) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *ReportResultReference) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportResultReference) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultReference) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportResultReference) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportResultReference) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportResultReference) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultReference) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportResultReference) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportResultReference) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportResultReference) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultReference) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportResultReference) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportResultReference) SetStatus(v string) {
	o.Status = &v
}

func (o ReportResultReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportResultReference) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportResultReference) UnmarshalJSON(bytes []byte) (err error) {
	varReportResultReference := _ReportResultReference{}

	if err = json.Unmarshal(bytes, &varReportResultReference); err == nil {
		*o = ReportResultReference(varReportResultReference)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportResultReference struct {
	value *ReportResultReference
	isSet bool
}

func (v NullableReportResultReference) Get() *ReportResultReference {
	return v.value
}

func (v *NullableReportResultReference) Set(val *ReportResultReference) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResultReference) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResultReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResultReference(val *ReportResultReference) *NullableReportResultReference {
	return &NullableReportResultReference{value: val, isSet: true}
}

func (v NullableReportResultReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportResultReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


