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

// checks if the SodReportResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodReportResultDto{}

// SodReportResultDto SOD policy violation report result.
type SodReportResultDto struct {
	// SOD policy violation report result DTO type.
	Type *string `json:"type,omitempty"`
	// SOD policy violation report result ID.
	Id *string `json:"id,omitempty"`
	// Human-readable name of the SOD policy violation report result.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodReportResultDto SodReportResultDto

// NewSodReportResultDto instantiates a new SodReportResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodReportResultDto() *SodReportResultDto {
	this := SodReportResultDto{}
	return &this
}

// NewSodReportResultDtoWithDefaults instantiates a new SodReportResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodReportResultDtoWithDefaults() *SodReportResultDto {
	this := SodReportResultDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SodReportResultDto) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodReportResultDto) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SodReportResultDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SodReportResultDto) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SodReportResultDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodReportResultDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SodReportResultDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SodReportResultDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SodReportResultDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodReportResultDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SodReportResultDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SodReportResultDto) SetName(v string) {
	o.Name = &v
}

func (o SodReportResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodReportResultDto) ToMap() (map[string]interface{}, error) {
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

func (o *SodReportResultDto) UnmarshalJSON(bytes []byte) (err error) {
	varSodReportResultDto := _SodReportResultDto{}

	if err = json.Unmarshal(bytes, &varSodReportResultDto); err == nil {
		*o = SodReportResultDto(varSodReportResultDto)
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

type NullableSodReportResultDto struct {
	value *SodReportResultDto
	isSet bool
}

func (v NullableSodReportResultDto) Get() *SodReportResultDto {
	return v.value
}

func (v *NullableSodReportResultDto) Set(val *SodReportResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSodReportResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSodReportResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodReportResultDto(val *SodReportResultDto) *NullableSodReportResultDto {
	return &NullableSodReportResultDto{value: val, isSet: true}
}

func (v NullableSodReportResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodReportResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


