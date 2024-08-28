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

// checks if the TemplateSlack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateSlack{}

// TemplateSlack struct for TemplateSlack
type TemplateSlack struct {
	Key NullableString `json:"key,omitempty"`
	Text *string `json:"text,omitempty"`
	Blocks NullableString `json:"blocks,omitempty"`
	Attachments *string `json:"attachments,omitempty"`
	NotificationType NullableString `json:"notificationType,omitempty"`
	ApprovalId NullableString `json:"approvalId,omitempty"`
	RequestId NullableString `json:"requestId,omitempty"`
	RequestedById NullableString `json:"requestedById,omitempty"`
	IsSubscription NullableBool `json:"isSubscription,omitempty"`
	AutoApprovalData NullableTemplateSlackAutoApprovalData `json:"autoApprovalData,omitempty"`
	CustomFields NullableTemplateSlackCustomFields `json:"customFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateSlack TemplateSlack

// NewTemplateSlack instantiates a new TemplateSlack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSlack() *TemplateSlack {
	this := TemplateSlack{}
	return &this
}

// NewTemplateSlackWithDefaults instantiates a new TemplateSlack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSlackWithDefaults() *TemplateSlack {
	this := TemplateSlack{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *TemplateSlack) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *TemplateSlack) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *TemplateSlack) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *TemplateSlack) UnsetKey() {
	o.Key.Unset()
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TemplateSlack) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSlack) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TemplateSlack) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TemplateSlack) SetText(v string) {
	o.Text = &v
}

// GetBlocks returns the Blocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetBlocks() string {
	if o == nil || IsNil(o.Blocks.Get()) {
		var ret string
		return ret
	}
	return *o.Blocks.Get()
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetBlocksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blocks.Get(), o.Blocks.IsSet()
}

// HasBlocks returns a boolean if a field has been set.
func (o *TemplateSlack) HasBlocks() bool {
	if o != nil && o.Blocks.IsSet() {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given NullableString and assigns it to the Blocks field.
func (o *TemplateSlack) SetBlocks(v string) {
	o.Blocks.Set(&v)
}
// SetBlocksNil sets the value for Blocks to be an explicit nil
func (o *TemplateSlack) SetBlocksNil() {
	o.Blocks.Set(nil)
}

// UnsetBlocks ensures that no value is present for Blocks, not even an explicit nil
func (o *TemplateSlack) UnsetBlocks() {
	o.Blocks.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TemplateSlack) GetAttachments() string {
	if o == nil || IsNil(o.Attachments) {
		var ret string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSlack) GetAttachmentsOk() (*string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TemplateSlack) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given string and assigns it to the Attachments field.
func (o *TemplateSlack) SetAttachments(v string) {
	o.Attachments = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetNotificationType() string {
	if o == nil || IsNil(o.NotificationType.Get()) {
		var ret string
		return ret
	}
	return *o.NotificationType.Get()
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetNotificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationType.Get(), o.NotificationType.IsSet()
}

// HasNotificationType returns a boolean if a field has been set.
func (o *TemplateSlack) HasNotificationType() bool {
	if o != nil && o.NotificationType.IsSet() {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given NullableString and assigns it to the NotificationType field.
func (o *TemplateSlack) SetNotificationType(v string) {
	o.NotificationType.Set(&v)
}
// SetNotificationTypeNil sets the value for NotificationType to be an explicit nil
func (o *TemplateSlack) SetNotificationTypeNil() {
	o.NotificationType.Set(nil)
}

// UnsetNotificationType ensures that no value is present for NotificationType, not even an explicit nil
func (o *TemplateSlack) UnsetNotificationType() {
	o.NotificationType.Unset()
}

// GetApprovalId returns the ApprovalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetApprovalId() string {
	if o == nil || IsNil(o.ApprovalId.Get()) {
		var ret string
		return ret
	}
	return *o.ApprovalId.Get()
}

// GetApprovalIdOk returns a tuple with the ApprovalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetApprovalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalId.Get(), o.ApprovalId.IsSet()
}

// HasApprovalId returns a boolean if a field has been set.
func (o *TemplateSlack) HasApprovalId() bool {
	if o != nil && o.ApprovalId.IsSet() {
		return true
	}

	return false
}

// SetApprovalId gets a reference to the given NullableString and assigns it to the ApprovalId field.
func (o *TemplateSlack) SetApprovalId(v string) {
	o.ApprovalId.Set(&v)
}
// SetApprovalIdNil sets the value for ApprovalId to be an explicit nil
func (o *TemplateSlack) SetApprovalIdNil() {
	o.ApprovalId.Set(nil)
}

// UnsetApprovalId ensures that no value is present for ApprovalId, not even an explicit nil
func (o *TemplateSlack) UnsetApprovalId() {
	o.ApprovalId.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *TemplateSlack) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *TemplateSlack) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *TemplateSlack) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *TemplateSlack) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetRequestedById returns the RequestedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetRequestedById() string {
	if o == nil || IsNil(o.RequestedById.Get()) {
		var ret string
		return ret
	}
	return *o.RequestedById.Get()
}

// GetRequestedByIdOk returns a tuple with the RequestedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetRequestedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedById.Get(), o.RequestedById.IsSet()
}

// HasRequestedById returns a boolean if a field has been set.
func (o *TemplateSlack) HasRequestedById() bool {
	if o != nil && o.RequestedById.IsSet() {
		return true
	}

	return false
}

// SetRequestedById gets a reference to the given NullableString and assigns it to the RequestedById field.
func (o *TemplateSlack) SetRequestedById(v string) {
	o.RequestedById.Set(&v)
}
// SetRequestedByIdNil sets the value for RequestedById to be an explicit nil
func (o *TemplateSlack) SetRequestedByIdNil() {
	o.RequestedById.Set(nil)
}

// UnsetRequestedById ensures that no value is present for RequestedById, not even an explicit nil
func (o *TemplateSlack) UnsetRequestedById() {
	o.RequestedById.Unset()
}

// GetIsSubscription returns the IsSubscription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetIsSubscription() bool {
	if o == nil || IsNil(o.IsSubscription.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSubscription.Get()
}

// GetIsSubscriptionOk returns a tuple with the IsSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetIsSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSubscription.Get(), o.IsSubscription.IsSet()
}

// HasIsSubscription returns a boolean if a field has been set.
func (o *TemplateSlack) HasIsSubscription() bool {
	if o != nil && o.IsSubscription.IsSet() {
		return true
	}

	return false
}

// SetIsSubscription gets a reference to the given NullableBool and assigns it to the IsSubscription field.
func (o *TemplateSlack) SetIsSubscription(v bool) {
	o.IsSubscription.Set(&v)
}
// SetIsSubscriptionNil sets the value for IsSubscription to be an explicit nil
func (o *TemplateSlack) SetIsSubscriptionNil() {
	o.IsSubscription.Set(nil)
}

// UnsetIsSubscription ensures that no value is present for IsSubscription, not even an explicit nil
func (o *TemplateSlack) UnsetIsSubscription() {
	o.IsSubscription.Unset()
}

// GetAutoApprovalData returns the AutoApprovalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetAutoApprovalData() TemplateSlackAutoApprovalData {
	if o == nil || IsNil(o.AutoApprovalData.Get()) {
		var ret TemplateSlackAutoApprovalData
		return ret
	}
	return *o.AutoApprovalData.Get()
}

// GetAutoApprovalDataOk returns a tuple with the AutoApprovalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetAutoApprovalDataOk() (*TemplateSlackAutoApprovalData, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoApprovalData.Get(), o.AutoApprovalData.IsSet()
}

// HasAutoApprovalData returns a boolean if a field has been set.
func (o *TemplateSlack) HasAutoApprovalData() bool {
	if o != nil && o.AutoApprovalData.IsSet() {
		return true
	}

	return false
}

// SetAutoApprovalData gets a reference to the given NullableTemplateSlackAutoApprovalData and assigns it to the AutoApprovalData field.
func (o *TemplateSlack) SetAutoApprovalData(v TemplateSlackAutoApprovalData) {
	o.AutoApprovalData.Set(&v)
}
// SetAutoApprovalDataNil sets the value for AutoApprovalData to be an explicit nil
func (o *TemplateSlack) SetAutoApprovalDataNil() {
	o.AutoApprovalData.Set(nil)
}

// UnsetAutoApprovalData ensures that no value is present for AutoApprovalData, not even an explicit nil
func (o *TemplateSlack) UnsetAutoApprovalData() {
	o.AutoApprovalData.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateSlack) GetCustomFields() TemplateSlackCustomFields {
	if o == nil || IsNil(o.CustomFields.Get()) {
		var ret TemplateSlackCustomFields
		return ret
	}
	return *o.CustomFields.Get()
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateSlack) GetCustomFieldsOk() (*TemplateSlackCustomFields, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFields.Get(), o.CustomFields.IsSet()
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TemplateSlack) HasCustomFields() bool {
	if o != nil && o.CustomFields.IsSet() {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given NullableTemplateSlackCustomFields and assigns it to the CustomFields field.
func (o *TemplateSlack) SetCustomFields(v TemplateSlackCustomFields) {
	o.CustomFields.Set(&v)
}
// SetCustomFieldsNil sets the value for CustomFields to be an explicit nil
func (o *TemplateSlack) SetCustomFieldsNil() {
	o.CustomFields.Set(nil)
}

// UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
func (o *TemplateSlack) UnsetCustomFields() {
	o.CustomFields.Unset()
}

func (o TemplateSlack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateSlack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if o.Blocks.IsSet() {
		toSerialize["blocks"] = o.Blocks.Get()
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if o.NotificationType.IsSet() {
		toSerialize["notificationType"] = o.NotificationType.Get()
	}
	if o.ApprovalId.IsSet() {
		toSerialize["approvalId"] = o.ApprovalId.Get()
	}
	if o.RequestId.IsSet() {
		toSerialize["requestId"] = o.RequestId.Get()
	}
	if o.RequestedById.IsSet() {
		toSerialize["requestedById"] = o.RequestedById.Get()
	}
	if o.IsSubscription.IsSet() {
		toSerialize["isSubscription"] = o.IsSubscription.Get()
	}
	if o.AutoApprovalData.IsSet() {
		toSerialize["autoApprovalData"] = o.AutoApprovalData.Get()
	}
	if o.CustomFields.IsSet() {
		toSerialize["customFields"] = o.CustomFields.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateSlack) UnmarshalJSON(data []byte) (err error) {
	varTemplateSlack := _TemplateSlack{}

	err = json.Unmarshal(data, &varTemplateSlack)

	if err != nil {
		return err
	}

	*o = TemplateSlack(varTemplateSlack)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "text")
		delete(additionalProperties, "blocks")
		delete(additionalProperties, "attachments")
		delete(additionalProperties, "notificationType")
		delete(additionalProperties, "approvalId")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "requestedById")
		delete(additionalProperties, "isSubscription")
		delete(additionalProperties, "autoApprovalData")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateSlack struct {
	value *TemplateSlack
	isSet bool
}

func (v NullableTemplateSlack) Get() *TemplateSlack {
	return v.value
}

func (v *NullableTemplateSlack) Set(val *TemplateSlack) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSlack) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSlack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSlack(val *TemplateSlack) *NullableTemplateSlack {
	return &NullableTemplateSlack{value: val, isSet: true}
}

func (v NullableTemplateSlack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSlack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


