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

// checks if the SavedSearchCompleteSearchResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedSearchCompleteSearchResults{}

// SavedSearchCompleteSearchResults A preview of the search results for each object type. This includes a count as well as headers, and the first several rows of data, per object type.
type SavedSearchCompleteSearchResults struct {
	Account NullableSavedSearchCompleteSearchResultsAccount `json:"Account,omitempty"`
	Entitlement NullableSavedSearchCompleteSearchResultsEntitlement `json:"Entitlement,omitempty"`
	Identity NullableSavedSearchCompleteSearchResultsIdentity `json:"Identity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearchCompleteSearchResults SavedSearchCompleteSearchResults

// NewSavedSearchCompleteSearchResults instantiates a new SavedSearchCompleteSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchCompleteSearchResults() *SavedSearchCompleteSearchResults {
	this := SavedSearchCompleteSearchResults{}
	return &this
}

// NewSavedSearchCompleteSearchResultsWithDefaults instantiates a new SavedSearchCompleteSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchCompleteSearchResultsWithDefaults() *SavedSearchCompleteSearchResults {
	this := SavedSearchCompleteSearchResults{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearchCompleteSearchResults) GetAccount() SavedSearchCompleteSearchResultsAccount {
	if o == nil || isNil(o.Account.Get()) {
		var ret SavedSearchCompleteSearchResultsAccount
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearchCompleteSearchResults) GetAccountOk() (*SavedSearchCompleteSearchResultsAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *SavedSearchCompleteSearchResults) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableSavedSearchCompleteSearchResultsAccount and assigns it to the Account field.
func (o *SavedSearchCompleteSearchResults) SetAccount(v SavedSearchCompleteSearchResultsAccount) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *SavedSearchCompleteSearchResults) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *SavedSearchCompleteSearchResults) UnsetAccount() {
	o.Account.Unset()
}

// GetEntitlement returns the Entitlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearchCompleteSearchResults) GetEntitlement() SavedSearchCompleteSearchResultsEntitlement {
	if o == nil || isNil(o.Entitlement.Get()) {
		var ret SavedSearchCompleteSearchResultsEntitlement
		return ret
	}
	return *o.Entitlement.Get()
}

// GetEntitlementOk returns a tuple with the Entitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearchCompleteSearchResults) GetEntitlementOk() (*SavedSearchCompleteSearchResultsEntitlement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlement.Get(), o.Entitlement.IsSet()
}

// HasEntitlement returns a boolean if a field has been set.
func (o *SavedSearchCompleteSearchResults) HasEntitlement() bool {
	if o != nil && o.Entitlement.IsSet() {
		return true
	}

	return false
}

// SetEntitlement gets a reference to the given NullableSavedSearchCompleteSearchResultsEntitlement and assigns it to the Entitlement field.
func (o *SavedSearchCompleteSearchResults) SetEntitlement(v SavedSearchCompleteSearchResultsEntitlement) {
	o.Entitlement.Set(&v)
}
// SetEntitlementNil sets the value for Entitlement to be an explicit nil
func (o *SavedSearchCompleteSearchResults) SetEntitlementNil() {
	o.Entitlement.Set(nil)
}

// UnsetEntitlement ensures that no value is present for Entitlement, not even an explicit nil
func (o *SavedSearchCompleteSearchResults) UnsetEntitlement() {
	o.Entitlement.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearchCompleteSearchResults) GetIdentity() SavedSearchCompleteSearchResultsIdentity {
	if o == nil || isNil(o.Identity.Get()) {
		var ret SavedSearchCompleteSearchResultsIdentity
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearchCompleteSearchResults) GetIdentityOk() (*SavedSearchCompleteSearchResultsIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *SavedSearchCompleteSearchResults) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableSavedSearchCompleteSearchResultsIdentity and assigns it to the Identity field.
func (o *SavedSearchCompleteSearchResults) SetIdentity(v SavedSearchCompleteSearchResultsIdentity) {
	o.Identity.Set(&v)
}
// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *SavedSearchCompleteSearchResults) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *SavedSearchCompleteSearchResults) UnsetIdentity() {
	o.Identity.Unset()
}

func (o SavedSearchCompleteSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedSearchCompleteSearchResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.Entitlement.IsSet() {
		toSerialize["Entitlement"] = o.Entitlement.Get()
	}
	if o.Identity.IsSet() {
		toSerialize["Identity"] = o.Identity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavedSearchCompleteSearchResults) UnmarshalJSON(bytes []byte) (err error) {
	varSavedSearchCompleteSearchResults := _SavedSearchCompleteSearchResults{}

	if err = json.Unmarshal(bytes, &varSavedSearchCompleteSearchResults); err == nil {
	*o = SavedSearchCompleteSearchResults(varSavedSearchCompleteSearchResults)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Entitlement")
		delete(additionalProperties, "Identity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearchCompleteSearchResults struct {
	value *SavedSearchCompleteSearchResults
	isSet bool
}

func (v NullableSavedSearchCompleteSearchResults) Get() *SavedSearchCompleteSearchResults {
	return v.value
}

func (v *NullableSavedSearchCompleteSearchResults) Set(val *SavedSearchCompleteSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchCompleteSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchCompleteSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchCompleteSearchResults(val *SavedSearchCompleteSearchResults) *NullableSavedSearchCompleteSearchResults {
	return &NullableSavedSearchCompleteSearchResults{value: val, isSet: true}
}

func (v NullableSavedSearchCompleteSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchCompleteSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


