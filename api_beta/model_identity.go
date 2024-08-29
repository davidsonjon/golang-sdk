/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Identity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Identity{}

// Identity struct for Identity
type Identity struct {
	// System-generated unique ID of the identity
	Id *string `json:"id,omitempty"`
	// The identity's name is equivalent to its Display Name attribute.
	Name string `json:"name"`
	// Creation date of the identity
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the identity
	Modified *time.Time `json:"modified,omitempty"`
	// The identity's alternate unique identifier is equivalent to its Account Name on the authoritative source account schema.
	Alias *string `json:"alias,omitempty"`
	// The email address of the identity
	EmailAddress NullableString `json:"emailAddress,omitempty"`
	// The processing state of the identity
	ProcessingState NullableString `json:"processingState,omitempty"`
	// The identity's status in the system
	IdentityStatus *string `json:"identityStatus,omitempty"`
	ManagerRef NullableIdentityManagerRef `json:"managerRef,omitempty"`
	// Whether this identity is a manager of another identity
	IsManager *bool `json:"isManager,omitempty"`
	// The last time the identity was refreshed by the system
	LastRefresh *time.Time `json:"lastRefresh,omitempty"`
	// A map with the identity attributes for the identity
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	LifecycleState *IdentityLifecycleState `json:"lifecycleState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Identity Identity

// NewIdentity instantiates a new Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity(name string) *Identity {
	this := Identity{}
	this.Name = name
	var isManager bool = false
	this.IsManager = &isManager
	return &this
}

// NewIdentityWithDefaults instantiates a new Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithDefaults() *Identity {
	this := Identity{}
	var isManager bool = false
	this.IsManager = &isManager
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Identity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Identity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Identity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Identity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Identity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Identity) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Identity) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Identity) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Identity) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Identity) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Identity) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Identity) SetModified(v time.Time) {
	o.Modified = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *Identity) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *Identity) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *Identity) SetAlias(v string) {
	o.Alias = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress.Get()) {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *Identity) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *Identity) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *Identity) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *Identity) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetProcessingState returns the ProcessingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetProcessingState() string {
	if o == nil || IsNil(o.ProcessingState.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessingState.Get()
}

// GetProcessingStateOk returns a tuple with the ProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetProcessingStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessingState.Get(), o.ProcessingState.IsSet()
}

// HasProcessingState returns a boolean if a field has been set.
func (o *Identity) HasProcessingState() bool {
	if o != nil && o.ProcessingState.IsSet() {
		return true
	}

	return false
}

// SetProcessingState gets a reference to the given NullableString and assigns it to the ProcessingState field.
func (o *Identity) SetProcessingState(v string) {
	o.ProcessingState.Set(&v)
}
// SetProcessingStateNil sets the value for ProcessingState to be an explicit nil
func (o *Identity) SetProcessingStateNil() {
	o.ProcessingState.Set(nil)
}

// UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
func (o *Identity) UnsetProcessingState() {
	o.ProcessingState.Unset()
}

// GetIdentityStatus returns the IdentityStatus field value if set, zero value otherwise.
func (o *Identity) GetIdentityStatus() string {
	if o == nil || IsNil(o.IdentityStatus) {
		var ret string
		return ret
	}
	return *o.IdentityStatus
}

// GetIdentityStatusOk returns a tuple with the IdentityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetIdentityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityStatus) {
		return nil, false
	}
	return o.IdentityStatus, true
}

// HasIdentityStatus returns a boolean if a field has been set.
func (o *Identity) HasIdentityStatus() bool {
	if o != nil && !IsNil(o.IdentityStatus) {
		return true
	}

	return false
}

// SetIdentityStatus gets a reference to the given string and assigns it to the IdentityStatus field.
func (o *Identity) SetIdentityStatus(v string) {
	o.IdentityStatus = &v
}

// GetManagerRef returns the ManagerRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetManagerRef() IdentityManagerRef {
	if o == nil || IsNil(o.ManagerRef.Get()) {
		var ret IdentityManagerRef
		return ret
	}
	return *o.ManagerRef.Get()
}

// GetManagerRefOk returns a tuple with the ManagerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetManagerRefOk() (*IdentityManagerRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagerRef.Get(), o.ManagerRef.IsSet()
}

// HasManagerRef returns a boolean if a field has been set.
func (o *Identity) HasManagerRef() bool {
	if o != nil && o.ManagerRef.IsSet() {
		return true
	}

	return false
}

// SetManagerRef gets a reference to the given NullableIdentityManagerRef and assigns it to the ManagerRef field.
func (o *Identity) SetManagerRef(v IdentityManagerRef) {
	o.ManagerRef.Set(&v)
}
// SetManagerRefNil sets the value for ManagerRef to be an explicit nil
func (o *Identity) SetManagerRefNil() {
	o.ManagerRef.Set(nil)
}

// UnsetManagerRef ensures that no value is present for ManagerRef, not even an explicit nil
func (o *Identity) UnsetManagerRef() {
	o.ManagerRef.Unset()
}

// GetIsManager returns the IsManager field value if set, zero value otherwise.
func (o *Identity) GetIsManager() bool {
	if o == nil || IsNil(o.IsManager) {
		var ret bool
		return ret
	}
	return *o.IsManager
}

// GetIsManagerOk returns a tuple with the IsManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetIsManagerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManager) {
		return nil, false
	}
	return o.IsManager, true
}

// HasIsManager returns a boolean if a field has been set.
func (o *Identity) HasIsManager() bool {
	if o != nil && !IsNil(o.IsManager) {
		return true
	}

	return false
}

// SetIsManager gets a reference to the given bool and assigns it to the IsManager field.
func (o *Identity) SetIsManager(v bool) {
	o.IsManager = &v
}

// GetLastRefresh returns the LastRefresh field value if set, zero value otherwise.
func (o *Identity) GetLastRefresh() time.Time {
	if o == nil || IsNil(o.LastRefresh) {
		var ret time.Time
		return ret
	}
	return *o.LastRefresh
}

// GetLastRefreshOk returns a tuple with the LastRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetLastRefreshOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRefresh) {
		return nil, false
	}
	return o.LastRefresh, true
}

// HasLastRefresh returns a boolean if a field has been set.
func (o *Identity) HasLastRefresh() bool {
	if o != nil && !IsNil(o.LastRefresh) {
		return true
	}

	return false
}

// SetLastRefresh gets a reference to the given time.Time and assigns it to the LastRefresh field.
func (o *Identity) SetLastRefresh(v time.Time) {
	o.LastRefresh = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Identity) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Identity) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Identity) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetLifecycleState returns the LifecycleState field value if set, zero value otherwise.
func (o *Identity) GetLifecycleState() IdentityLifecycleState {
	if o == nil || IsNil(o.LifecycleState) {
		var ret IdentityLifecycleState
		return ret
	}
	return *o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetLifecycleStateOk() (*IdentityLifecycleState, bool) {
	if o == nil || IsNil(o.LifecycleState) {
		return nil, false
	}
	return o.LifecycleState, true
}

// HasLifecycleState returns a boolean if a field has been set.
func (o *Identity) HasLifecycleState() bool {
	if o != nil && !IsNil(o.LifecycleState) {
		return true
	}

	return false
}

// SetLifecycleState gets a reference to the given IdentityLifecycleState and assigns it to the LifecycleState field.
func (o *Identity) SetLifecycleState(v IdentityLifecycleState) {
	o.LifecycleState = &v
}

func (o Identity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Identity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if o.EmailAddress.IsSet() {
		toSerialize["emailAddress"] = o.EmailAddress.Get()
	}
	if o.ProcessingState.IsSet() {
		toSerialize["processingState"] = o.ProcessingState.Get()
	}
	if !IsNil(o.IdentityStatus) {
		toSerialize["identityStatus"] = o.IdentityStatus
	}
	if o.ManagerRef.IsSet() {
		toSerialize["managerRef"] = o.ManagerRef.Get()
	}
	if !IsNil(o.IsManager) {
		toSerialize["isManager"] = o.IsManager
	}
	if !IsNil(o.LastRefresh) {
		toSerialize["lastRefresh"] = o.LastRefresh
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.LifecycleState) {
		toSerialize["lifecycleState"] = o.LifecycleState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Identity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varIdentity := _Identity{}

	err = json.Unmarshal(data, &varIdentity)

	if err != nil {
		return err
	}

	*o = Identity(varIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "alias")
		delete(additionalProperties, "emailAddress")
		delete(additionalProperties, "processingState")
		delete(additionalProperties, "identityStatus")
		delete(additionalProperties, "managerRef")
		delete(additionalProperties, "isManager")
		delete(additionalProperties, "lastRefresh")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "lifecycleState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentity struct {
	value *Identity
	isSet bool
}

func (v NullableIdentity) Get() *Identity {
	return v.value
}

func (v *NullableIdentity) Set(val *Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity(val *Identity) *NullableIdentity {
	return &NullableIdentity{value: val, isSet: true}
}

func (v NullableIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


