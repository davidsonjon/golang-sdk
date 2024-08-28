/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateFormInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFormInstanceRequest{}

// CreateFormInstanceRequest struct for CreateFormInstanceRequest
type CreateFormInstanceRequest struct {
	CreatedBy FormInstanceCreatedBy `json:"createdBy"`
	// Expire is required
	Expire string `json:"expire"`
	// FormDefinitionID is the id of the form definition that created this form
	FormDefinitionId string `json:"formDefinitionId"`
	// FormInput is an object of form input labels to value
	FormInput map[string]interface{} `json:"formInput,omitempty"`
	// Recipients is required
	Recipients []FormInstanceRecipient `json:"recipients"`
	// StandAloneForm is a boolean flag to indicate if this form should be available for users to complete via the standalone form UI or should this only be available to be completed by as an embedded form
	StandAloneForm *bool `json:"standAloneForm,omitempty"`
	// State is required, if not present initial state is FormInstanceStateAssigned ASSIGNED FormInstanceStateAssigned IN_PROGRESS FormInstanceStateInProgress SUBMITTED FormInstanceStateSubmitted COMPLETED FormInstanceStateCompleted CANCELLED FormInstanceStateCancelled
	State *string `json:"state,omitempty"`
	// TTL an epoch timestamp in seconds, it most be in seconds or dynamodb will ignore it SEE: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/time-to-live-ttl-before-you-start.html
	Ttl *int64 `json:"ttl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFormInstanceRequest CreateFormInstanceRequest

// NewCreateFormInstanceRequest instantiates a new CreateFormInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFormInstanceRequest(createdBy FormInstanceCreatedBy, expire string, formDefinitionId string, recipients []FormInstanceRecipient) *CreateFormInstanceRequest {
	this := CreateFormInstanceRequest{}
	this.CreatedBy = createdBy
	this.Expire = expire
	this.FormDefinitionId = formDefinitionId
	this.Recipients = recipients
	var standAloneForm bool = false
	this.StandAloneForm = &standAloneForm
	return &this
}

// NewCreateFormInstanceRequestWithDefaults instantiates a new CreateFormInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFormInstanceRequestWithDefaults() *CreateFormInstanceRequest {
	this := CreateFormInstanceRequest{}
	var standAloneForm bool = false
	this.StandAloneForm = &standAloneForm
	return &this
}

// GetCreatedBy returns the CreatedBy field value
func (o *CreateFormInstanceRequest) GetCreatedBy() FormInstanceCreatedBy {
	if o == nil {
		var ret FormInstanceCreatedBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetCreatedByOk() (*FormInstanceCreatedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *CreateFormInstanceRequest) SetCreatedBy(v FormInstanceCreatedBy) {
	o.CreatedBy = v
}

// GetExpire returns the Expire field value
func (o *CreateFormInstanceRequest) GetExpire() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expire
}

// GetExpireOk returns a tuple with the Expire field value
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetExpireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expire, true
}

// SetExpire sets field value
func (o *CreateFormInstanceRequest) SetExpire(v string) {
	o.Expire = v
}

// GetFormDefinitionId returns the FormDefinitionId field value
func (o *CreateFormInstanceRequest) GetFormDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormDefinitionId
}

// GetFormDefinitionIdOk returns a tuple with the FormDefinitionId field value
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetFormDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormDefinitionId, true
}

// SetFormDefinitionId sets field value
func (o *CreateFormInstanceRequest) SetFormDefinitionId(v string) {
	o.FormDefinitionId = v
}

// GetFormInput returns the FormInput field value if set, zero value otherwise.
func (o *CreateFormInstanceRequest) GetFormInput() map[string]interface{} {
	if o == nil || IsNil(o.FormInput) {
		var ret map[string]interface{}
		return ret
	}
	return o.FormInput
}

// GetFormInputOk returns a tuple with the FormInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetFormInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FormInput) {
		return map[string]interface{}{}, false
	}
	return o.FormInput, true
}

// HasFormInput returns a boolean if a field has been set.
func (o *CreateFormInstanceRequest) HasFormInput() bool {
	if o != nil && !IsNil(o.FormInput) {
		return true
	}

	return false
}

// SetFormInput gets a reference to the given map[string]interface{} and assigns it to the FormInput field.
func (o *CreateFormInstanceRequest) SetFormInput(v map[string]interface{}) {
	o.FormInput = v
}

// GetRecipients returns the Recipients field value
func (o *CreateFormInstanceRequest) GetRecipients() []FormInstanceRecipient {
	if o == nil {
		var ret []FormInstanceRecipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetRecipientsOk() ([]FormInstanceRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateFormInstanceRequest) SetRecipients(v []FormInstanceRecipient) {
	o.Recipients = v
}

// GetStandAloneForm returns the StandAloneForm field value if set, zero value otherwise.
func (o *CreateFormInstanceRequest) GetStandAloneForm() bool {
	if o == nil || IsNil(o.StandAloneForm) {
		var ret bool
		return ret
	}
	return *o.StandAloneForm
}

// GetStandAloneFormOk returns a tuple with the StandAloneForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetStandAloneFormOk() (*bool, bool) {
	if o == nil || IsNil(o.StandAloneForm) {
		return nil, false
	}
	return o.StandAloneForm, true
}

// HasStandAloneForm returns a boolean if a field has been set.
func (o *CreateFormInstanceRequest) HasStandAloneForm() bool {
	if o != nil && !IsNil(o.StandAloneForm) {
		return true
	}

	return false
}

// SetStandAloneForm gets a reference to the given bool and assigns it to the StandAloneForm field.
func (o *CreateFormInstanceRequest) SetStandAloneForm(v bool) {
	o.StandAloneForm = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateFormInstanceRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateFormInstanceRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CreateFormInstanceRequest) SetState(v string) {
	o.State = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CreateFormInstanceRequest) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormInstanceRequest) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CreateFormInstanceRequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *CreateFormInstanceRequest) SetTtl(v int64) {
	o.Ttl = &v
}

func (o CreateFormInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFormInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["expire"] = o.Expire
	toSerialize["formDefinitionId"] = o.FormDefinitionId
	if !IsNil(o.FormInput) {
		toSerialize["formInput"] = o.FormInput
	}
	toSerialize["recipients"] = o.Recipients
	if !IsNil(o.StandAloneForm) {
		toSerialize["standAloneForm"] = o.StandAloneForm
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFormInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdBy",
		"expire",
		"formDefinitionId",
		"recipients",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateFormInstanceRequest := _CreateFormInstanceRequest{}

	err = json.Unmarshal(data, &varCreateFormInstanceRequest)

	if err != nil {
		return err
	}

	*o = CreateFormInstanceRequest(varCreateFormInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "expire")
		delete(additionalProperties, "formDefinitionId")
		delete(additionalProperties, "formInput")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "standAloneForm")
		delete(additionalProperties, "state")
		delete(additionalProperties, "ttl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFormInstanceRequest struct {
	value *CreateFormInstanceRequest
	isSet bool
}

func (v NullableCreateFormInstanceRequest) Get() *CreateFormInstanceRequest {
	return v.value
}

func (v *NullableCreateFormInstanceRequest) Set(val *CreateFormInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFormInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFormInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFormInstanceRequest(val *CreateFormInstanceRequest) *NullableCreateFormInstanceRequest {
	return &NullableCreateFormInstanceRequest{value: val, isSet: true}
}

func (v NullableCreateFormInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFormInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


