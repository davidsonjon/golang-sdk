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

// checks if the FormDefinitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDefinitionResponse{}

// FormDefinitionResponse struct for FormDefinitionResponse
type FormDefinitionResponse struct {
	// Unique guid identifying the form definition.
	Id *string `json:"id,omitempty"`
	// Name of the form definition.
	Name *string `json:"name,omitempty"`
	// Form definition's description.
	Description *string `json:"description,omitempty"`
	Owner *FormOwner `json:"owner,omitempty"`
	// List of objects using the form definition. Whenever a system uses a form, the API reaches out to the form service to record that the system is currently using it.
	UsedBy []FormUsedBy `json:"usedBy,omitempty"`
	// List of form inputs required to create a form-instance object.
	FormInput []FormDefinitionInput `json:"formInput,omitempty"`
	// List of nested form elements.
	FormElements []FormElement `json:"formElements,omitempty"`
	// Conditional logic that can dynamically modify the form as the recipient is interacting with it.
	FormConditions []FormCondition `json:"formConditions,omitempty"`
	// Created is the date the form definition was created
	Created *time.Time `json:"created,omitempty"`
	// Modified is the last date the form definition was modified
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormDefinitionResponse FormDefinitionResponse

// NewFormDefinitionResponse instantiates a new FormDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDefinitionResponse() *FormDefinitionResponse {
	this := FormDefinitionResponse{}
	return &this
}

// NewFormDefinitionResponseWithDefaults instantiates a new FormDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDefinitionResponseWithDefaults() *FormDefinitionResponse {
	this := FormDefinitionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormDefinitionResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormDefinitionResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FormDefinitionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetOwner() FormOwner {
	if o == nil || isNil(o.Owner) {
		var ret FormOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetOwnerOk() (*FormOwner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given FormOwner and assigns it to the Owner field.
func (o *FormDefinitionResponse) SetOwner(v FormOwner) {
	o.Owner = &v
}

// GetUsedBy returns the UsedBy field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetUsedBy() []FormUsedBy {
	if o == nil || isNil(o.UsedBy) {
		var ret []FormUsedBy
		return ret
	}
	return o.UsedBy
}

// GetUsedByOk returns a tuple with the UsedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetUsedByOk() ([]FormUsedBy, bool) {
	if o == nil || isNil(o.UsedBy) {
		return nil, false
	}
	return o.UsedBy, true
}

// HasUsedBy returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasUsedBy() bool {
	if o != nil && !isNil(o.UsedBy) {
		return true
	}

	return false
}

// SetUsedBy gets a reference to the given []FormUsedBy and assigns it to the UsedBy field.
func (o *FormDefinitionResponse) SetUsedBy(v []FormUsedBy) {
	o.UsedBy = v
}

// GetFormInput returns the FormInput field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetFormInput() []FormDefinitionInput {
	if o == nil || isNil(o.FormInput) {
		var ret []FormDefinitionInput
		return ret
	}
	return o.FormInput
}

// GetFormInputOk returns a tuple with the FormInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetFormInputOk() ([]FormDefinitionInput, bool) {
	if o == nil || isNil(o.FormInput) {
		return nil, false
	}
	return o.FormInput, true
}

// HasFormInput returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasFormInput() bool {
	if o != nil && !isNil(o.FormInput) {
		return true
	}

	return false
}

// SetFormInput gets a reference to the given []FormDefinitionInput and assigns it to the FormInput field.
func (o *FormDefinitionResponse) SetFormInput(v []FormDefinitionInput) {
	o.FormInput = v
}

// GetFormElements returns the FormElements field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetFormElements() []FormElement {
	if o == nil || isNil(o.FormElements) {
		var ret []FormElement
		return ret
	}
	return o.FormElements
}

// GetFormElementsOk returns a tuple with the FormElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetFormElementsOk() ([]FormElement, bool) {
	if o == nil || isNil(o.FormElements) {
		return nil, false
	}
	return o.FormElements, true
}

// HasFormElements returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasFormElements() bool {
	if o != nil && !isNil(o.FormElements) {
		return true
	}

	return false
}

// SetFormElements gets a reference to the given []FormElement and assigns it to the FormElements field.
func (o *FormDefinitionResponse) SetFormElements(v []FormElement) {
	o.FormElements = v
}

// GetFormConditions returns the FormConditions field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetFormConditions() []FormCondition {
	if o == nil || isNil(o.FormConditions) {
		var ret []FormCondition
		return ret
	}
	return o.FormConditions
}

// GetFormConditionsOk returns a tuple with the FormConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetFormConditionsOk() ([]FormCondition, bool) {
	if o == nil || isNil(o.FormConditions) {
		return nil, false
	}
	return o.FormConditions, true
}

// HasFormConditions returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasFormConditions() bool {
	if o != nil && !isNil(o.FormConditions) {
		return true
	}

	return false
}

// SetFormConditions gets a reference to the given []FormCondition and assigns it to the FormConditions field.
func (o *FormDefinitionResponse) SetFormConditions(v []FormCondition) {
	o.FormConditions = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FormDefinitionResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *FormDefinitionResponse) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *FormDefinitionResponse) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *FormDefinitionResponse) SetModified(v time.Time) {
	o.Modified = &v
}

func (o FormDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDefinitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.UsedBy) {
		toSerialize["usedBy"] = o.UsedBy
	}
	if !isNil(o.FormInput) {
		toSerialize["formInput"] = o.FormInput
	}
	if !isNil(o.FormElements) {
		toSerialize["formElements"] = o.FormElements
	}
	if !isNil(o.FormConditions) {
		toSerialize["formConditions"] = o.FormConditions
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormDefinitionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varFormDefinitionResponse := _FormDefinitionResponse{}

	if err = json.Unmarshal(bytes, &varFormDefinitionResponse); err == nil {
		*o = FormDefinitionResponse(varFormDefinitionResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "usedBy")
		delete(additionalProperties, "formInput")
		delete(additionalProperties, "formElements")
		delete(additionalProperties, "formConditions")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormDefinitionResponse struct {
	value *FormDefinitionResponse
	isSet bool
}

func (v NullableFormDefinitionResponse) Get() *FormDefinitionResponse {
	return v.value
}

func (v *NullableFormDefinitionResponse) Set(val *FormDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDefinitionResponse(val *FormDefinitionResponse) *NullableFormDefinitionResponse {
	return &NullableFormDefinitionResponse{value: val, isSet: true}
}

func (v NullableFormDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


