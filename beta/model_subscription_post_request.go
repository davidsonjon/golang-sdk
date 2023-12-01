/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// checks if the SubscriptionPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPostRequest{}

// SubscriptionPostRequest struct for SubscriptionPostRequest
type SubscriptionPostRequest struct {
	// Subscription name.
	Name string `json:"name"`
	// Subscription description.
	Description *string `json:"description,omitempty"`
	// ID of trigger subscribed to.
	TriggerId string `json:"triggerId"`
	Type SubscriptionType `json:"type"`
	// Deadline for completing REQUEST_RESPONSE trigger invocation, represented in ISO-8601 duration format.
	ResponseDeadline *string `json:"responseDeadline,omitempty"`
	HttpConfig *HttpConfig `json:"httpConfig,omitempty"`
	EventBridgeConfig *EventBridgeConfig `json:"eventBridgeConfig,omitempty"`
	// Whether subscription should receive real-time trigger invocations or not.  Test trigger invocations are always enabled regardless of this option.
	Enabled *bool `json:"enabled,omitempty"`
	// JSONPath filter to conditionally invoke trigger when expression evaluates to true.
	Filter *string `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionPostRequest SubscriptionPostRequest

// NewSubscriptionPostRequest instantiates a new SubscriptionPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPostRequest(name string, triggerId string, type_ SubscriptionType) *SubscriptionPostRequest {
	this := SubscriptionPostRequest{}
	this.Name = name
	this.TriggerId = triggerId
	this.Type = type_
	var responseDeadline string = "PT1H"
	this.ResponseDeadline = &responseDeadline
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewSubscriptionPostRequestWithDefaults instantiates a new SubscriptionPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPostRequestWithDefaults() *SubscriptionPostRequest {
	this := SubscriptionPostRequest{}
	var responseDeadline string = "PT1H"
	this.ResponseDeadline = &responseDeadline
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value
func (o *SubscriptionPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionPostRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubscriptionPostRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubscriptionPostRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubscriptionPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTriggerId returns the TriggerId field value
func (o *SubscriptionPostRequest) GetTriggerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetTriggerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *SubscriptionPostRequest) SetTriggerId(v string) {
	o.TriggerId = v
}

// GetType returns the Type field value
func (o *SubscriptionPostRequest) GetType() SubscriptionType {
	if o == nil {
		var ret SubscriptionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetTypeOk() (*SubscriptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPostRequest) SetType(v SubscriptionType) {
	o.Type = v
}

// GetResponseDeadline returns the ResponseDeadline field value if set, zero value otherwise.
func (o *SubscriptionPostRequest) GetResponseDeadline() string {
	if o == nil || isNil(o.ResponseDeadline) {
		var ret string
		return ret
	}
	return *o.ResponseDeadline
}

// GetResponseDeadlineOk returns a tuple with the ResponseDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetResponseDeadlineOk() (*string, bool) {
	if o == nil || isNil(o.ResponseDeadline) {
		return nil, false
	}
	return o.ResponseDeadline, true
}

// HasResponseDeadline returns a boolean if a field has been set.
func (o *SubscriptionPostRequest) HasResponseDeadline() bool {
	if o != nil && !isNil(o.ResponseDeadline) {
		return true
	}

	return false
}

// SetResponseDeadline gets a reference to the given string and assigns it to the ResponseDeadline field.
func (o *SubscriptionPostRequest) SetResponseDeadline(v string) {
	o.ResponseDeadline = &v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *SubscriptionPostRequest) GetHttpConfig() HttpConfig {
	if o == nil || isNil(o.HttpConfig) {
		var ret HttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetHttpConfigOk() (*HttpConfig, bool) {
	if o == nil || isNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *SubscriptionPostRequest) HasHttpConfig() bool {
	if o != nil && !isNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given HttpConfig and assigns it to the HttpConfig field.
func (o *SubscriptionPostRequest) SetHttpConfig(v HttpConfig) {
	o.HttpConfig = &v
}

// GetEventBridgeConfig returns the EventBridgeConfig field value if set, zero value otherwise.
func (o *SubscriptionPostRequest) GetEventBridgeConfig() EventBridgeConfig {
	if o == nil || isNil(o.EventBridgeConfig) {
		var ret EventBridgeConfig
		return ret
	}
	return *o.EventBridgeConfig
}

// GetEventBridgeConfigOk returns a tuple with the EventBridgeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetEventBridgeConfigOk() (*EventBridgeConfig, bool) {
	if o == nil || isNil(o.EventBridgeConfig) {
		return nil, false
	}
	return o.EventBridgeConfig, true
}

// HasEventBridgeConfig returns a boolean if a field has been set.
func (o *SubscriptionPostRequest) HasEventBridgeConfig() bool {
	if o != nil && !isNil(o.EventBridgeConfig) {
		return true
	}

	return false
}

// SetEventBridgeConfig gets a reference to the given EventBridgeConfig and assigns it to the EventBridgeConfig field.
func (o *SubscriptionPostRequest) SetEventBridgeConfig(v EventBridgeConfig) {
	o.EventBridgeConfig = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SubscriptionPostRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SubscriptionPostRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SubscriptionPostRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SubscriptionPostRequest) GetFilter() string {
	if o == nil || isNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPostRequest) GetFilterOk() (*string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SubscriptionPostRequest) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *SubscriptionPostRequest) SetFilter(v string) {
	o.Filter = &v
}

func (o SubscriptionPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["triggerId"] = o.TriggerId
	toSerialize["type"] = o.Type
	if !isNil(o.ResponseDeadline) {
		toSerialize["responseDeadline"] = o.ResponseDeadline
	}
	if !isNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	if !isNil(o.EventBridgeConfig) {
		toSerialize["eventBridgeConfig"] = o.EventBridgeConfig
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionPostRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"triggerId",
		"type",
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

	varSubscriptionPostRequest := _SubscriptionPostRequest{}

	if err = json.Unmarshal(bytes, &varSubscriptionPostRequest); err == nil {
	*o = SubscriptionPostRequest(varSubscriptionPostRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "triggerId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "responseDeadline")
		delete(additionalProperties, "httpConfig")
		delete(additionalProperties, "eventBridgeConfig")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionPostRequest struct {
	value *SubscriptionPostRequest
	isSet bool
}

func (v NullableSubscriptionPostRequest) Get() *SubscriptionPostRequest {
	return v.value
}

func (v *NullableSubscriptionPostRequest) Set(val *SubscriptionPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPostRequest(val *SubscriptionPostRequest) *NullableSubscriptionPostRequest {
	return &NullableSubscriptionPostRequest{value: val, isSet: true}
}

func (v NullableSubscriptionPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


