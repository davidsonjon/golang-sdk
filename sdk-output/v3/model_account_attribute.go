/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// AccountAttribute struct for AccountAttribute
type AccountAttribute struct {
	// A reference to the source to search for the account
	SourceName string `json:"sourceName"`
	// The name of the attribute on the account to return. This should match the name of the account attribute name visible in the user interface, or on the source schema.
	AttributeName string `json:"attributeName"`
	// The value of this configuration is a string name of the attribute to use when determining the ordering of returned accounts when there are multiple entries
	AccountSortAttribute *string `json:"accountSortAttribute,omitempty"`
	// The value of this configuration is a boolean (true/false). Controls the order of the sort when there are multiple accounts. If not defined, the transform will default to false (ascending order)
	AccountSortDescending *bool `json:"accountSortDescending,omitempty"`
	// The value of this configuration is a boolean (true/false). Controls which account to source a value from for an attribute.  If this flag is set to true, the transform returns the value from the first account in the list, even if it is null. If it is set to false, the transform returns the first non-null value. If not defined, the transform will default to false
	AccountReturnFirstLink *bool `json:"accountReturnFirstLink,omitempty"`
	// This expression queries the database to narrow search results. The value of this configuration is a sailpoint.object.Filter expression and used when searching against the database.  The default filter will always include the source and identity, and any subsequent expressions will be combined in an AND operation to the existing search criteria. Only certain searchable attributes are available:  - `nativeIdentity` - the Account ID  - `displayName` - the Account Name  - `entitlements` - a boolean value to determine if the account has entitlements
	AccountFilter *string `json:"accountFilter,omitempty"`
	// This expression is used to search and filter accounts in memory. The value of this configuration is a sailpoint.object.Filter expression and used when searching against the returned resultset.  All account attributes are available for filtering as this operation is performed in memory.
	AccountPropertyFilter *string `json:"accountPropertyFilter,omitempty"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountAttribute AccountAttribute

// NewAccountAttribute instantiates a new AccountAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAttribute(sourceName string, attributeName string) *AccountAttribute {
	this := AccountAttribute{}
	this.SourceName = sourceName
	this.AttributeName = attributeName
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewAccountAttributeWithDefaults instantiates a new AccountAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAttributeWithDefaults() *AccountAttribute {
	this := AccountAttribute{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetSourceName returns the SourceName field value
func (o *AccountAttribute) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *AccountAttribute) SetSourceName(v string) {
	o.SourceName = v
}

// GetAttributeName returns the AttributeName field value
func (o *AccountAttribute) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *AccountAttribute) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetAccountSortAttribute returns the AccountSortAttribute field value if set, zero value otherwise.
func (o *AccountAttribute) GetAccountSortAttribute() string {
	if o == nil || isNil(o.AccountSortAttribute) {
		var ret string
		return ret
	}
	return *o.AccountSortAttribute
}

// GetAccountSortAttributeOk returns a tuple with the AccountSortAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetAccountSortAttributeOk() (*string, bool) {
	if o == nil || isNil(o.AccountSortAttribute) {
		return nil, false
	}
	return o.AccountSortAttribute, true
}

// HasAccountSortAttribute returns a boolean if a field has been set.
func (o *AccountAttribute) HasAccountSortAttribute() bool {
	if o != nil && !isNil(o.AccountSortAttribute) {
		return true
	}

	return false
}

// SetAccountSortAttribute gets a reference to the given string and assigns it to the AccountSortAttribute field.
func (o *AccountAttribute) SetAccountSortAttribute(v string) {
	o.AccountSortAttribute = &v
}

// GetAccountSortDescending returns the AccountSortDescending field value if set, zero value otherwise.
func (o *AccountAttribute) GetAccountSortDescending() bool {
	if o == nil || isNil(o.AccountSortDescending) {
		var ret bool
		return ret
	}
	return *o.AccountSortDescending
}

// GetAccountSortDescendingOk returns a tuple with the AccountSortDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetAccountSortDescendingOk() (*bool, bool) {
	if o == nil || isNil(o.AccountSortDescending) {
		return nil, false
	}
	return o.AccountSortDescending, true
}

// HasAccountSortDescending returns a boolean if a field has been set.
func (o *AccountAttribute) HasAccountSortDescending() bool {
	if o != nil && !isNil(o.AccountSortDescending) {
		return true
	}

	return false
}

// SetAccountSortDescending gets a reference to the given bool and assigns it to the AccountSortDescending field.
func (o *AccountAttribute) SetAccountSortDescending(v bool) {
	o.AccountSortDescending = &v
}

// GetAccountReturnFirstLink returns the AccountReturnFirstLink field value if set, zero value otherwise.
func (o *AccountAttribute) GetAccountReturnFirstLink() bool {
	if o == nil || isNil(o.AccountReturnFirstLink) {
		var ret bool
		return ret
	}
	return *o.AccountReturnFirstLink
}

// GetAccountReturnFirstLinkOk returns a tuple with the AccountReturnFirstLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetAccountReturnFirstLinkOk() (*bool, bool) {
	if o == nil || isNil(o.AccountReturnFirstLink) {
		return nil, false
	}
	return o.AccountReturnFirstLink, true
}

// HasAccountReturnFirstLink returns a boolean if a field has been set.
func (o *AccountAttribute) HasAccountReturnFirstLink() bool {
	if o != nil && !isNil(o.AccountReturnFirstLink) {
		return true
	}

	return false
}

// SetAccountReturnFirstLink gets a reference to the given bool and assigns it to the AccountReturnFirstLink field.
func (o *AccountAttribute) SetAccountReturnFirstLink(v bool) {
	o.AccountReturnFirstLink = &v
}

// GetAccountFilter returns the AccountFilter field value if set, zero value otherwise.
func (o *AccountAttribute) GetAccountFilter() string {
	if o == nil || isNil(o.AccountFilter) {
		var ret string
		return ret
	}
	return *o.AccountFilter
}

// GetAccountFilterOk returns a tuple with the AccountFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetAccountFilterOk() (*string, bool) {
	if o == nil || isNil(o.AccountFilter) {
		return nil, false
	}
	return o.AccountFilter, true
}

// HasAccountFilter returns a boolean if a field has been set.
func (o *AccountAttribute) HasAccountFilter() bool {
	if o != nil && !isNil(o.AccountFilter) {
		return true
	}

	return false
}

// SetAccountFilter gets a reference to the given string and assigns it to the AccountFilter field.
func (o *AccountAttribute) SetAccountFilter(v string) {
	o.AccountFilter = &v
}

// GetAccountPropertyFilter returns the AccountPropertyFilter field value if set, zero value otherwise.
func (o *AccountAttribute) GetAccountPropertyFilter() string {
	if o == nil || isNil(o.AccountPropertyFilter) {
		var ret string
		return ret
	}
	return *o.AccountPropertyFilter
}

// GetAccountPropertyFilterOk returns a tuple with the AccountPropertyFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetAccountPropertyFilterOk() (*string, bool) {
	if o == nil || isNil(o.AccountPropertyFilter) {
		return nil, false
	}
	return o.AccountPropertyFilter, true
}

// HasAccountPropertyFilter returns a boolean if a field has been set.
func (o *AccountAttribute) HasAccountPropertyFilter() bool {
	if o != nil && !isNil(o.AccountPropertyFilter) {
		return true
	}

	return false
}

// SetAccountPropertyFilter gets a reference to the given string and assigns it to the AccountPropertyFilter field.
func (o *AccountAttribute) SetAccountPropertyFilter(v string) {
	o.AccountPropertyFilter = &v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *AccountAttribute) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *AccountAttribute) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *AccountAttribute) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *AccountAttribute) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAttribute) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *AccountAttribute) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *AccountAttribute) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o AccountAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceName"] = o.SourceName
	}
	if true {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !isNil(o.AccountSortAttribute) {
		toSerialize["accountSortAttribute"] = o.AccountSortAttribute
	}
	if !isNil(o.AccountSortDescending) {
		toSerialize["accountSortDescending"] = o.AccountSortDescending
	}
	if !isNil(o.AccountReturnFirstLink) {
		toSerialize["accountReturnFirstLink"] = o.AccountReturnFirstLink
	}
	if !isNil(o.AccountFilter) {
		toSerialize["accountFilter"] = o.AccountFilter
	}
	if !isNil(o.AccountPropertyFilter) {
		toSerialize["accountPropertyFilter"] = o.AccountPropertyFilter
	}
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varAccountAttribute := _AccountAttribute{}

	if err = json.Unmarshal(bytes, &varAccountAttribute); err == nil {
		*o = AccountAttribute(varAccountAttribute)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "accountSortAttribute")
		delete(additionalProperties, "accountSortDescending")
		delete(additionalProperties, "accountReturnFirstLink")
		delete(additionalProperties, "accountFilter")
		delete(additionalProperties, "accountPropertyFilter")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAttribute struct {
	value *AccountAttribute
	isSet bool
}

func (v NullableAccountAttribute) Get() *AccountAttribute {
	return v.value
}

func (v *NullableAccountAttribute) Set(val *AccountAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAttribute(val *AccountAttribute) *NullableAccountAttribute {
	return &NullableAccountAttribute{value: val, isSet: true}
}

func (v NullableAccountAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


