/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"time"
)

// checks if the FormInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormInstanceResponse{}

// FormInstanceResponse struct for FormInstanceResponse
type FormInstanceResponse struct {
	// Created is the date the form instance was assigned
	Created *time.Time `json:"created,omitempty"`
	CreatedBy *FormInstanceCreatedBy `json:"createdBy,omitempty"`
	// Expire is the maximum amount of time that a form can be in progress. After this time is reached then the form will be moved to a CANCELED state automatically. The user will no longer be able to complete the submission. When a form instance is expires an audit log will be generated for that record
	Expire *string `json:"expire,omitempty"`
	// FormConditions is the conditional logic that modify the form dynamically modify the form as the recipient is interacting out the form
	FormConditions []FormCondition `json:"formConditions,omitempty"`
	// FormData is the data provided by the form on submit. The data is in a key -> value map
	FormData map[string]map[string]interface{} `json:"formData,omitempty"`
	// FormDefinitionID is the id of the form definition that created this form
	FormDefinitionId *string `json:"formDefinitionId,omitempty"`
	// FormElements is the configuration of the form, this would be a repeat of the fields from the form-config
	FormElements []FormElement `json:"formElements,omitempty"`
	// FormErrors is an array of form validation errors from the last time the form instance was transitioned to the SUBMITTED state. If the form instance had validation errors then it would be moved to the IN PROGRESS state where the client can retrieve these errors
	FormErrors []FormError `json:"formErrors,omitempty"`
	// FormInput is an object of form input labels to value
	FormInput map[string]map[string]interface{} `json:"formInput,omitempty"`
	// FormInstanceID is a unique guid identifying this form instance
	Id *string `json:"id,omitempty"`
	// Modified is the last date the form instance was modified
	Modified *time.Time `json:"modified,omitempty"`
	// Recipients references to the recipient of a form. The recipients are those who are responsible for filling out a form and completing it
	Recipients []FormInstanceRecipient `json:"recipients,omitempty"`
	// StandAloneForm is a boolean flag to indicate if this form should be available for users to complete via the standalone form UI or should this only be available to be completed by as an embedded form
	StandAloneForm *bool `json:"standAloneForm,omitempty"`
	// StandAloneFormURL is the URL where this form may be completed by the designated recipients using the standalone form UI
	StandAloneFormUrl *string `json:"standAloneFormUrl,omitempty"`
	// State the state of the form instance ASSIGNED FormInstanceStateAssigned IN_PROGRESS FormInstanceStateInProgress SUBMITTED FormInstanceStateSubmitted COMPLETED FormInstanceStateCompleted CANCELLED FormInstanceStateCancelled
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormInstanceResponse FormInstanceResponse

// NewFormInstanceResponse instantiates a new FormInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormInstanceResponse() *FormInstanceResponse {
	this := FormInstanceResponse{}
	var standAloneForm bool = false
	this.StandAloneForm = &standAloneForm
	return &this
}

// NewFormInstanceResponseWithDefaults instantiates a new FormInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormInstanceResponseWithDefaults() *FormInstanceResponse {
	this := FormInstanceResponse{}
	var standAloneForm bool = false
	this.StandAloneForm = &standAloneForm
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FormInstanceResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetCreatedBy() FormInstanceCreatedBy {
	if o == nil || isNil(o.CreatedBy) {
		var ret FormInstanceCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetCreatedByOk() (*FormInstanceCreatedBy, bool) {
	if o == nil || isNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given FormInstanceCreatedBy and assigns it to the CreatedBy field.
func (o *FormInstanceResponse) SetCreatedBy(v FormInstanceCreatedBy) {
	o.CreatedBy = &v
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetExpire() string {
	if o == nil || isNil(o.Expire) {
		var ret string
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetExpireOk() (*string, bool) {
	if o == nil || isNil(o.Expire) {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasExpire() bool {
	if o != nil && !isNil(o.Expire) {
		return true
	}

	return false
}

// SetExpire gets a reference to the given string and assigns it to the Expire field.
func (o *FormInstanceResponse) SetExpire(v string) {
	o.Expire = &v
}

// GetFormConditions returns the FormConditions field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetFormConditions() []FormCondition {
	if o == nil || isNil(o.FormConditions) {
		var ret []FormCondition
		return ret
	}
	return o.FormConditions
}

// GetFormConditionsOk returns a tuple with the FormConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetFormConditionsOk() ([]FormCondition, bool) {
	if o == nil || isNil(o.FormConditions) {
		return nil, false
	}
	return o.FormConditions, true
}

// HasFormConditions returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasFormConditions() bool {
	if o != nil && !isNil(o.FormConditions) {
		return true
	}

	return false
}

// SetFormConditions gets a reference to the given []FormCondition and assigns it to the FormConditions field.
func (o *FormInstanceResponse) SetFormConditions(v []FormCondition) {
	o.FormConditions = v
}

// GetFormData returns the FormData field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetFormData() map[string]map[string]interface{} {
	if o == nil || isNil(o.FormData) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.FormData
}

// GetFormDataOk returns a tuple with the FormData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetFormDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.FormData) {
		return map[string]map[string]interface{}{}, false
	}
	return o.FormData, true
}

// HasFormData returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasFormData() bool {
	if o != nil && !isNil(o.FormData) {
		return true
	}

	return false
}

// SetFormData gets a reference to the given map[string]map[string]interface{} and assigns it to the FormData field.
func (o *FormInstanceResponse) SetFormData(v map[string]map[string]interface{}) {
	o.FormData = v
}

// GetFormDefinitionId returns the FormDefinitionId field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetFormDefinitionId() string {
	if o == nil || isNil(o.FormDefinitionId) {
		var ret string
		return ret
	}
	return *o.FormDefinitionId
}

// GetFormDefinitionIdOk returns a tuple with the FormDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetFormDefinitionIdOk() (*string, bool) {
	if o == nil || isNil(o.FormDefinitionId) {
		return nil, false
	}
	return o.FormDefinitionId, true
}

// HasFormDefinitionId returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasFormDefinitionId() bool {
	if o != nil && !isNil(o.FormDefinitionId) {
		return true
	}

	return false
}

// SetFormDefinitionId gets a reference to the given string and assigns it to the FormDefinitionId field.
func (o *FormInstanceResponse) SetFormDefinitionId(v string) {
	o.FormDefinitionId = &v
}

// GetFormElements returns the FormElements field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetFormElements() []FormElement {
	if o == nil || isNil(o.FormElements) {
		var ret []FormElement
		return ret
	}
	return o.FormElements
}

// GetFormElementsOk returns a tuple with the FormElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetFormElementsOk() ([]FormElement, bool) {
	if o == nil || isNil(o.FormElements) {
		return nil, false
	}
	return o.FormElements, true
}

// HasFormElements returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasFormElements() bool {
	if o != nil && !isNil(o.FormElements) {
		return true
	}

	return false
}

// SetFormElements gets a reference to the given []FormElement and assigns it to the FormElements field.
func (o *FormInstanceResponse) SetFormElements(v []FormElement) {
	o.FormElements = v
}

// GetFormErrors returns the FormErrors field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetFormErrors() []FormError {
	if o == nil || isNil(o.FormErrors) {
		var ret []FormError
		return ret
	}
	return o.FormErrors
}

// GetFormErrorsOk returns a tuple with the FormErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetFormErrorsOk() ([]FormError, bool) {
	if o == nil || isNil(o.FormErrors) {
		return nil, false
	}
	return o.FormErrors, true
}

// HasFormErrors returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasFormErrors() bool {
	if o != nil && !isNil(o.FormErrors) {
		return true
	}

	return false
}

// SetFormErrors gets a reference to the given []FormError and assigns it to the FormErrors field.
func (o *FormInstanceResponse) SetFormErrors(v []FormError) {
	o.FormErrors = v
}

// GetFormInput returns the FormInput field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetFormInput() map[string]map[string]interface{} {
	if o == nil || isNil(o.FormInput) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.FormInput
}

// GetFormInputOk returns a tuple with the FormInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetFormInputOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.FormInput) {
		return map[string]map[string]interface{}{}, false
	}
	return o.FormInput, true
}

// HasFormInput returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasFormInput() bool {
	if o != nil && !isNil(o.FormInput) {
		return true
	}

	return false
}

// SetFormInput gets a reference to the given map[string]map[string]interface{} and assigns it to the FormInput field.
func (o *FormInstanceResponse) SetFormInput(v map[string]map[string]interface{}) {
	o.FormInput = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormInstanceResponse) SetId(v string) {
	o.Id = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *FormInstanceResponse) SetModified(v time.Time) {
	o.Modified = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetRecipients() []FormInstanceRecipient {
	if o == nil || isNil(o.Recipients) {
		var ret []FormInstanceRecipient
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetRecipientsOk() ([]FormInstanceRecipient, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []FormInstanceRecipient and assigns it to the Recipients field.
func (o *FormInstanceResponse) SetRecipients(v []FormInstanceRecipient) {
	o.Recipients = v
}

// GetStandAloneForm returns the StandAloneForm field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetStandAloneForm() bool {
	if o == nil || isNil(o.StandAloneForm) {
		var ret bool
		return ret
	}
	return *o.StandAloneForm
}

// GetStandAloneFormOk returns a tuple with the StandAloneForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetStandAloneFormOk() (*bool, bool) {
	if o == nil || isNil(o.StandAloneForm) {
		return nil, false
	}
	return o.StandAloneForm, true
}

// HasStandAloneForm returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasStandAloneForm() bool {
	if o != nil && !isNil(o.StandAloneForm) {
		return true
	}

	return false
}

// SetStandAloneForm gets a reference to the given bool and assigns it to the StandAloneForm field.
func (o *FormInstanceResponse) SetStandAloneForm(v bool) {
	o.StandAloneForm = &v
}

// GetStandAloneFormUrl returns the StandAloneFormUrl field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetStandAloneFormUrl() string {
	if o == nil || isNil(o.StandAloneFormUrl) {
		var ret string
		return ret
	}
	return *o.StandAloneFormUrl
}

// GetStandAloneFormUrlOk returns a tuple with the StandAloneFormUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetStandAloneFormUrlOk() (*string, bool) {
	if o == nil || isNil(o.StandAloneFormUrl) {
		return nil, false
	}
	return o.StandAloneFormUrl, true
}

// HasStandAloneFormUrl returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasStandAloneFormUrl() bool {
	if o != nil && !isNil(o.StandAloneFormUrl) {
		return true
	}

	return false
}

// SetStandAloneFormUrl gets a reference to the given string and assigns it to the StandAloneFormUrl field.
func (o *FormInstanceResponse) SetStandAloneFormUrl(v string) {
	o.StandAloneFormUrl = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FormInstanceResponse) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInstanceResponse) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FormInstanceResponse) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FormInstanceResponse) SetState(v string) {
	o.State = &v
}

func (o FormInstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.Expire) {
		toSerialize["expire"] = o.Expire
	}
	if !isNil(o.FormConditions) {
		toSerialize["formConditions"] = o.FormConditions
	}
	if !isNil(o.FormData) {
		toSerialize["formData"] = o.FormData
	}
	if !isNil(o.FormDefinitionId) {
		toSerialize["formDefinitionId"] = o.FormDefinitionId
	}
	if !isNil(o.FormElements) {
		toSerialize["formElements"] = o.FormElements
	}
	if !isNil(o.FormErrors) {
		toSerialize["formErrors"] = o.FormErrors
	}
	if !isNil(o.FormInput) {
		toSerialize["formInput"] = o.FormInput
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !isNil(o.StandAloneForm) {
		toSerialize["standAloneForm"] = o.StandAloneForm
	}
	if !isNil(o.StandAloneFormUrl) {
		toSerialize["standAloneFormUrl"] = o.StandAloneFormUrl
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormInstanceResponse) UnmarshalJSON(bytes []byte) (err error) {
	varFormInstanceResponse := _FormInstanceResponse{}

	if err = json.Unmarshal(bytes, &varFormInstanceResponse); err == nil {
		*o = FormInstanceResponse(varFormInstanceResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "expire")
		delete(additionalProperties, "formConditions")
		delete(additionalProperties, "formData")
		delete(additionalProperties, "formDefinitionId")
		delete(additionalProperties, "formElements")
		delete(additionalProperties, "formErrors")
		delete(additionalProperties, "formInput")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "standAloneForm")
		delete(additionalProperties, "standAloneFormUrl")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormInstanceResponse struct {
	value *FormInstanceResponse
	isSet bool
}

func (v NullableFormInstanceResponse) Get() *FormInstanceResponse {
	return v.value
}

func (v *NullableFormInstanceResponse) Set(val *FormInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFormInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFormInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormInstanceResponse(val *FormInstanceResponse) *NullableFormInstanceResponse {
	return &NullableFormInstanceResponse{value: val, isSet: true}
}

func (v NullableFormInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


