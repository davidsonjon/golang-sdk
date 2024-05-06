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

// checks if the ApprovalItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalItems{}

// ApprovalItems struct for ApprovalItems
type ApprovalItems struct {
	// The approval item's ID
	Id *string `json:"id,omitempty"`
	// The account referenced by the approval item
	Account *string `json:"account,omitempty"`
	// The name of the application/source
	Application *string `json:"application,omitempty"`
	// The attribute's name
	Name *string `json:"name,omitempty"`
	// The attribute's operation
	Operation *string `json:"operation,omitempty"`
	// The attribute's value
	Value *string `json:"value,omitempty"`
	State *WorkItemState `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalItems ApprovalItems

// NewApprovalItems instantiates a new ApprovalItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalItems() *ApprovalItems {
	this := ApprovalItems{}
	return &this
}

// NewApprovalItemsWithDefaults instantiates a new ApprovalItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalItemsWithDefaults() *ApprovalItems {
	this := ApprovalItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApprovalItems) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApprovalItems) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApprovalItems) SetId(v string) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApprovalItems) GetAccount() string {
	if o == nil || isNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetAccountOk() (*string, bool) {
	if o == nil || isNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApprovalItems) HasAccount() bool {
	if o != nil && !isNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ApprovalItems) SetAccount(v string) {
	o.Account = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApprovalItems) GetApplication() string {
	if o == nil || isNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetApplicationOk() (*string, bool) {
	if o == nil || isNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApprovalItems) HasApplication() bool {
	if o != nil && !isNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ApprovalItems) SetApplication(v string) {
	o.Application = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApprovalItems) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApprovalItems) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApprovalItems) SetName(v string) {
	o.Name = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ApprovalItems) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ApprovalItems) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ApprovalItems) SetOperation(v string) {
	o.Operation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApprovalItems) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApprovalItems) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApprovalItems) SetValue(v string) {
	o.Value = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApprovalItems) GetState() WorkItemState {
	if o == nil || isNil(o.State) {
		var ret WorkItemState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetStateOk() (*WorkItemState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApprovalItems) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given WorkItemState and assigns it to the State field.
func (o *ApprovalItems) SetState(v WorkItemState) {
	o.State = &v
}

func (o ApprovalItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !isNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalItems) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalItems := _ApprovalItems{}

	if err = json.Unmarshal(bytes, &varApprovalItems); err == nil {
	*o = ApprovalItems(varApprovalItems)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "account")
		delete(additionalProperties, "application")
		delete(additionalProperties, "name")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "value")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalItems struct {
	value *ApprovalItems
	isSet bool
}

func (v NullableApprovalItems) Get() *ApprovalItems {
	return v.value
}

func (v *NullableApprovalItems) Set(val *ApprovalItems) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalItems) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalItems(val *ApprovalItems) *NullableApprovalItems {
	return &NullableApprovalItems{value: val, isSet: true}
}

func (v NullableApprovalItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


