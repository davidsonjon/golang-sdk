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

// checks if the AccountsExportReportArguments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsExportReportArguments{}

// AccountsExportReportArguments Arguments for Account Export (ACCOUNTS)
type AccountsExportReportArguments struct {
	// Id of the authoritative source to export related accounts e.g. identities
	Application string `json:"application"`
	// Name of the authoritative source for accounts export
	SourceName string `json:"sourceName"`
	AdditionalProperties map[string]interface{}
}

type _AccountsExportReportArguments AccountsExportReportArguments

// NewAccountsExportReportArguments instantiates a new AccountsExportReportArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsExportReportArguments(application string, sourceName string) *AccountsExportReportArguments {
	this := AccountsExportReportArguments{}
	this.Application = application
	this.SourceName = sourceName
	return &this
}

// NewAccountsExportReportArgumentsWithDefaults instantiates a new AccountsExportReportArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsExportReportArgumentsWithDefaults() *AccountsExportReportArguments {
	this := AccountsExportReportArguments{}
	return &this
}

// GetApplication returns the Application field value
func (o *AccountsExportReportArguments) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *AccountsExportReportArguments) GetApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *AccountsExportReportArguments) SetApplication(v string) {
	o.Application = v
}

// GetSourceName returns the SourceName field value
func (o *AccountsExportReportArguments) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *AccountsExportReportArguments) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *AccountsExportReportArguments) SetSourceName(v string) {
	o.SourceName = v
}

func (o AccountsExportReportArguments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsExportReportArguments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application"] = o.Application
	toSerialize["sourceName"] = o.SourceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountsExportReportArguments) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"application",
		"sourceName",
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

	varAccountsExportReportArguments := _AccountsExportReportArguments{}

	if err = json.Unmarshal(bytes, &varAccountsExportReportArguments); err == nil {
	*o = AccountsExportReportArguments(varAccountsExportReportArguments)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "application")
		delete(additionalProperties, "sourceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountsExportReportArguments struct {
	value *AccountsExportReportArguments
	isSet bool
}

func (v NullableAccountsExportReportArguments) Get() *AccountsExportReportArguments {
	return v.value
}

func (v *NullableAccountsExportReportArguments) Set(val *AccountsExportReportArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsExportReportArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsExportReportArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsExportReportArguments(val *AccountsExportReportArguments) *NullableAccountsExportReportArguments {
	return &NullableAccountsExportReportArguments{value: val, isSet: true}
}

func (v NullableAccountsExportReportArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsExportReportArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


