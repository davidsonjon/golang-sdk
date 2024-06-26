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

// checks if the TemplateSlackCustomFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateSlackCustomFields{}

// TemplateSlackCustomFields struct for TemplateSlackCustomFields
type TemplateSlackCustomFields struct {
	RequestType NullableString `json:"requestType,omitempty"`
	ContainsDeny NullableString `json:"containsDeny,omitempty"`
	CampaignId NullableString `json:"campaignId,omitempty"`
	CampaignStatus NullableString `json:"campaignStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateSlackCustomFields TemplateSlackCustomFields

// NewTemplateSlackCustomFields instantiates a new TemplateSlackCustomFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSlackCustomFields() *TemplateSlackCustomFields {
	this := TemplateSlackCustomFields{}
	return &this
}

// NewTemplateSlackCustomFieldsWithDefaults instantiates a new TemplateSlackCustomFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSlackCustomFieldsWithDefaults() *TemplateSlackCustomFields {
	this := TemplateSlackCustomFields{}
	return &this
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlackCustomFields) GetRequestType() string {
	if o == nil || isNil(o.RequestType.Get()) {
		var ret string
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlackCustomFields) GetRequestTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *TemplateSlackCustomFields) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableString and assigns it to the RequestType field.
func (o *TemplateSlackCustomFields) SetRequestType(v string) {
	o.RequestType.Set(&v)
}
// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *TemplateSlackCustomFields) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *TemplateSlackCustomFields) UnsetRequestType() {
	o.RequestType.Unset()
}

// GetContainsDeny returns the ContainsDeny field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlackCustomFields) GetContainsDeny() string {
	if o == nil || isNil(o.ContainsDeny.Get()) {
		var ret string
		return ret
	}
	return *o.ContainsDeny.Get()
}

// GetContainsDenyOk returns a tuple with the ContainsDeny field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlackCustomFields) GetContainsDenyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainsDeny.Get(), o.ContainsDeny.IsSet()
}

// HasContainsDeny returns a boolean if a field has been set.
func (o *TemplateSlackCustomFields) HasContainsDeny() bool {
	if o != nil && o.ContainsDeny.IsSet() {
		return true
	}

	return false
}

// SetContainsDeny gets a reference to the given NullableString and assigns it to the ContainsDeny field.
func (o *TemplateSlackCustomFields) SetContainsDeny(v string) {
	o.ContainsDeny.Set(&v)
}
// SetContainsDenyNil sets the value for ContainsDeny to be an explicit nil
func (o *TemplateSlackCustomFields) SetContainsDenyNil() {
	o.ContainsDeny.Set(nil)
}

// UnsetContainsDeny ensures that no value is present for ContainsDeny, not even an explicit nil
func (o *TemplateSlackCustomFields) UnsetContainsDeny() {
	o.ContainsDeny.Unset()
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlackCustomFields) GetCampaignId() string {
	if o == nil || isNil(o.CampaignId.Get()) {
		var ret string
		return ret
	}
	return *o.CampaignId.Get()
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlackCustomFields) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignId.Get(), o.CampaignId.IsSet()
}

// HasCampaignId returns a boolean if a field has been set.
func (o *TemplateSlackCustomFields) HasCampaignId() bool {
	if o != nil && o.CampaignId.IsSet() {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given NullableString and assigns it to the CampaignId field.
func (o *TemplateSlackCustomFields) SetCampaignId(v string) {
	o.CampaignId.Set(&v)
}
// SetCampaignIdNil sets the value for CampaignId to be an explicit nil
func (o *TemplateSlackCustomFields) SetCampaignIdNil() {
	o.CampaignId.Set(nil)
}

// UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
func (o *TemplateSlackCustomFields) UnsetCampaignId() {
	o.CampaignId.Unset()
}

// GetCampaignStatus returns the CampaignStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlackCustomFields) GetCampaignStatus() string {
	if o == nil || isNil(o.CampaignStatus.Get()) {
		var ret string
		return ret
	}
	return *o.CampaignStatus.Get()
}

// GetCampaignStatusOk returns a tuple with the CampaignStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlackCustomFields) GetCampaignStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignStatus.Get(), o.CampaignStatus.IsSet()
}

// HasCampaignStatus returns a boolean if a field has been set.
func (o *TemplateSlackCustomFields) HasCampaignStatus() bool {
	if o != nil && o.CampaignStatus.IsSet() {
		return true
	}

	return false
}

// SetCampaignStatus gets a reference to the given NullableString and assigns it to the CampaignStatus field.
func (o *TemplateSlackCustomFields) SetCampaignStatus(v string) {
	o.CampaignStatus.Set(&v)
}
// SetCampaignStatusNil sets the value for CampaignStatus to be an explicit nil
func (o *TemplateSlackCustomFields) SetCampaignStatusNil() {
	o.CampaignStatus.Set(nil)
}

// UnsetCampaignStatus ensures that no value is present for CampaignStatus, not even an explicit nil
func (o *TemplateSlackCustomFields) UnsetCampaignStatus() {
	o.CampaignStatus.Unset()
}

func (o TemplateSlackCustomFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateSlackCustomFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestType.IsSet() {
		toSerialize["requestType"] = o.RequestType.Get()
	}
	if o.ContainsDeny.IsSet() {
		toSerialize["containsDeny"] = o.ContainsDeny.Get()
	}
	if o.CampaignId.IsSet() {
		toSerialize["campaignId"] = o.CampaignId.Get()
	}
	if o.CampaignStatus.IsSet() {
		toSerialize["campaignStatus"] = o.CampaignStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateSlackCustomFields) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateSlackCustomFields := _TemplateSlackCustomFields{}

	if err = json.Unmarshal(bytes, &varTemplateSlackCustomFields); err == nil {
	*o = TemplateSlackCustomFields(varTemplateSlackCustomFields)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "containsDeny")
		delete(additionalProperties, "campaignId")
		delete(additionalProperties, "campaignStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateSlackCustomFields struct {
	value *TemplateSlackCustomFields
	isSet bool
}

func (v NullableTemplateSlackCustomFields) Get() *TemplateSlackCustomFields {
	return v.value
}

func (v *NullableTemplateSlackCustomFields) Set(val *TemplateSlackCustomFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSlackCustomFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSlackCustomFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSlackCustomFields(val *TemplateSlackCustomFields) *NullableTemplateSlackCustomFields {
	return &NullableTemplateSlackCustomFields{value: val, isSet: true}
}

func (v NullableTemplateSlackCustomFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSlackCustomFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

