/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the ProvisioningCompletedAccountRequestsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningCompletedAccountRequestsInner{}

// ProvisioningCompletedAccountRequestsInner struct for ProvisioningCompletedAccountRequestsInner
type ProvisioningCompletedAccountRequestsInner struct {
	Source ProvisioningCompletedAccountRequestsInnerSource `json:"source"`
	// The unique idenfier of the account being provisioned.
	AccountId *string `json:"accountId,omitempty"`
	// The provisioning operation; typically Create, Modify, Enable, Disable, Unlock, or Delete.
	AccountOperation string `json:"accountOperation"`
	// The overall result of the provisioning transaction; this could be success, pending, failed, etc.
	ProvisioningResult map[string]interface{} `json:"provisioningResult"`
	// The name of the provisioning channel selected; this could be the same as the source, or could be a Service Desk Integration Module (SDIM).
	ProvisioningTarget string `json:"provisioningTarget"`
	// A reference to a tracking number, if this is sent to a Service Desk Integration Module (SDIM).
	TicketId NullableString `json:"ticketId,omitempty"`
	// A list of attributes as part of the provisioning transaction.
	AttributeRequests []ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner `json:"attributeRequests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningCompletedAccountRequestsInner ProvisioningCompletedAccountRequestsInner

// NewProvisioningCompletedAccountRequestsInner instantiates a new ProvisioningCompletedAccountRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningCompletedAccountRequestsInner(source ProvisioningCompletedAccountRequestsInnerSource, accountOperation string, provisioningResult map[string]interface{}, provisioningTarget string) *ProvisioningCompletedAccountRequestsInner {
	this := ProvisioningCompletedAccountRequestsInner{}
	this.Source = source
	this.AccountOperation = accountOperation
	this.ProvisioningResult = provisioningResult
	this.ProvisioningTarget = provisioningTarget
	return &this
}

// NewProvisioningCompletedAccountRequestsInnerWithDefaults instantiates a new ProvisioningCompletedAccountRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningCompletedAccountRequestsInnerWithDefaults() *ProvisioningCompletedAccountRequestsInner {
	this := ProvisioningCompletedAccountRequestsInner{}
	return &this
}

// GetSource returns the Source field value
func (o *ProvisioningCompletedAccountRequestsInner) GetSource() ProvisioningCompletedAccountRequestsInnerSource {
	if o == nil {
		var ret ProvisioningCompletedAccountRequestsInnerSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInner) GetSourceOk() (*ProvisioningCompletedAccountRequestsInnerSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ProvisioningCompletedAccountRequestsInner) SetSource(v ProvisioningCompletedAccountRequestsInnerSource) {
	o.Source = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ProvisioningCompletedAccountRequestsInner) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInner) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ProvisioningCompletedAccountRequestsInner) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ProvisioningCompletedAccountRequestsInner) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountOperation returns the AccountOperation field value
func (o *ProvisioningCompletedAccountRequestsInner) GetAccountOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountOperation
}

// GetAccountOperationOk returns a tuple with the AccountOperation field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInner) GetAccountOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountOperation, true
}

// SetAccountOperation sets field value
func (o *ProvisioningCompletedAccountRequestsInner) SetAccountOperation(v string) {
	o.AccountOperation = v
}

// GetProvisioningResult returns the ProvisioningResult field value
func (o *ProvisioningCompletedAccountRequestsInner) GetProvisioningResult() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ProvisioningResult
}

// GetProvisioningResultOk returns a tuple with the ProvisioningResult field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInner) GetProvisioningResultOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ProvisioningResult, true
}

// SetProvisioningResult sets field value
func (o *ProvisioningCompletedAccountRequestsInner) SetProvisioningResult(v map[string]interface{}) {
	o.ProvisioningResult = v
}

// GetProvisioningTarget returns the ProvisioningTarget field value
func (o *ProvisioningCompletedAccountRequestsInner) GetProvisioningTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningTarget
}

// GetProvisioningTargetOk returns a tuple with the ProvisioningTarget field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCompletedAccountRequestsInner) GetProvisioningTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningTarget, true
}

// SetProvisioningTarget sets field value
func (o *ProvisioningCompletedAccountRequestsInner) SetProvisioningTarget(v string) {
	o.ProvisioningTarget = v
}

// GetTicketId returns the TicketId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningCompletedAccountRequestsInner) GetTicketId() string {
	if o == nil || isNil(o.TicketId.Get()) {
		var ret string
		return ret
	}
	return *o.TicketId.Get()
}

// GetTicketIdOk returns a tuple with the TicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningCompletedAccountRequestsInner) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketId.Get(), o.TicketId.IsSet()
}

// HasTicketId returns a boolean if a field has been set.
func (o *ProvisioningCompletedAccountRequestsInner) HasTicketId() bool {
	if o != nil && o.TicketId.IsSet() {
		return true
	}

	return false
}

// SetTicketId gets a reference to the given NullableString and assigns it to the TicketId field.
func (o *ProvisioningCompletedAccountRequestsInner) SetTicketId(v string) {
	o.TicketId.Set(&v)
}
// SetTicketIdNil sets the value for TicketId to be an explicit nil
func (o *ProvisioningCompletedAccountRequestsInner) SetTicketIdNil() {
	o.TicketId.Set(nil)
}

// UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
func (o *ProvisioningCompletedAccountRequestsInner) UnsetTicketId() {
	o.TicketId.Unset()
}

// GetAttributeRequests returns the AttributeRequests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningCompletedAccountRequestsInner) GetAttributeRequests() []ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	if o == nil {
		var ret []ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner
		return ret
	}
	return o.AttributeRequests
}

// GetAttributeRequestsOk returns a tuple with the AttributeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningCompletedAccountRequestsInner) GetAttributeRequestsOk() ([]ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner, bool) {
	if o == nil || isNil(o.AttributeRequests) {
		return nil, false
	}
	return o.AttributeRequests, true
}

// HasAttributeRequests returns a boolean if a field has been set.
func (o *ProvisioningCompletedAccountRequestsInner) HasAttributeRequests() bool {
	if o != nil && isNil(o.AttributeRequests) {
		return true
	}

	return false
}

// SetAttributeRequests gets a reference to the given []ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner and assigns it to the AttributeRequests field.
func (o *ProvisioningCompletedAccountRequestsInner) SetAttributeRequests(v []ProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) {
	o.AttributeRequests = v
}

func (o ProvisioningCompletedAccountRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningCompletedAccountRequestsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	toSerialize["accountOperation"] = o.AccountOperation
	toSerialize["provisioningResult"] = o.ProvisioningResult
	toSerialize["provisioningTarget"] = o.ProvisioningTarget
	if o.TicketId.IsSet() {
		toSerialize["ticketId"] = o.TicketId.Get()
	}
	if o.AttributeRequests != nil {
		toSerialize["attributeRequests"] = o.AttributeRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningCompletedAccountRequestsInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"accountOperation",
		"provisioningResult",
		"provisioningTarget",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProvisioningCompletedAccountRequestsInner := _ProvisioningCompletedAccountRequestsInner{}

	if err = json.Unmarshal(bytes, &varProvisioningCompletedAccountRequestsInner); err == nil {
	*o = ProvisioningCompletedAccountRequestsInner(varProvisioningCompletedAccountRequestsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "accountOperation")
		delete(additionalProperties, "provisioningResult")
		delete(additionalProperties, "provisioningTarget")
		delete(additionalProperties, "ticketId")
		delete(additionalProperties, "attributeRequests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningCompletedAccountRequestsInner struct {
	value *ProvisioningCompletedAccountRequestsInner
	isSet bool
}

func (v NullableProvisioningCompletedAccountRequestsInner) Get() *ProvisioningCompletedAccountRequestsInner {
	return v.value
}

func (v *NullableProvisioningCompletedAccountRequestsInner) Set(val *ProvisioningCompletedAccountRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningCompletedAccountRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningCompletedAccountRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningCompletedAccountRequestsInner(val *ProvisioningCompletedAccountRequestsInner) *NullableProvisioningCompletedAccountRequestsInner {
	return &NullableProvisioningCompletedAccountRequestsInner{value: val, isSet: true}
}

func (v NullableProvisioningCompletedAccountRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningCompletedAccountRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


