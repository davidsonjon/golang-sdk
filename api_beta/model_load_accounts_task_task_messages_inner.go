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

// checks if the LoadAccountsTaskTaskMessagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadAccountsTaskTaskMessagesInner{}

// LoadAccountsTaskTaskMessagesInner struct for LoadAccountsTaskTaskMessagesInner
type LoadAccountsTaskTaskMessagesInner struct {
	// Type of the message.
	Type *string `json:"type,omitempty"`
	// Flag whether message is an error.
	Error *bool `json:"error,omitempty"`
	// Flag whether message is a warning.
	Warning *bool `json:"warning,omitempty"`
	// Message string identifier.
	Key *string `json:"key,omitempty"`
	// Message context with the locale based language.
	LocalizedText *string `json:"localizedText,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadAccountsTaskTaskMessagesInner LoadAccountsTaskTaskMessagesInner

// NewLoadAccountsTaskTaskMessagesInner instantiates a new LoadAccountsTaskTaskMessagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadAccountsTaskTaskMessagesInner() *LoadAccountsTaskTaskMessagesInner {
	this := LoadAccountsTaskTaskMessagesInner{}
	var error_ bool = false
	this.Error = &error_
	var warning bool = false
	this.Warning = &warning
	return &this
}

// NewLoadAccountsTaskTaskMessagesInnerWithDefaults instantiates a new LoadAccountsTaskTaskMessagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadAccountsTaskTaskMessagesInnerWithDefaults() *LoadAccountsTaskTaskMessagesInner {
	this := LoadAccountsTaskTaskMessagesInner{}
	var error_ bool = false
	this.Error = &error_
	var warning bool = false
	this.Warning = &warning
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskMessagesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskMessagesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskMessagesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LoadAccountsTaskTaskMessagesInner) SetType(v string) {
	o.Type = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskMessagesInner) GetError() bool {
	if o == nil || IsNil(o.Error) {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskMessagesInner) GetErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskMessagesInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *LoadAccountsTaskTaskMessagesInner) SetError(v bool) {
	o.Error = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskMessagesInner) GetWarning() bool {
	if o == nil || IsNil(o.Warning) {
		var ret bool
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskMessagesInner) GetWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskMessagesInner) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given bool and assigns it to the Warning field.
func (o *LoadAccountsTaskTaskMessagesInner) SetWarning(v bool) {
	o.Warning = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskMessagesInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskMessagesInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskMessagesInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LoadAccountsTaskTaskMessagesInner) SetKey(v string) {
	o.Key = &v
}

// GetLocalizedText returns the LocalizedText field value if set, zero value otherwise.
func (o *LoadAccountsTaskTaskMessagesInner) GetLocalizedText() string {
	if o == nil || IsNil(o.LocalizedText) {
		var ret string
		return ret
	}
	return *o.LocalizedText
}

// GetLocalizedTextOk returns a tuple with the LocalizedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTaskTaskMessagesInner) GetLocalizedTextOk() (*string, bool) {
	if o == nil || IsNil(o.LocalizedText) {
		return nil, false
	}
	return o.LocalizedText, true
}

// HasLocalizedText returns a boolean if a field has been set.
func (o *LoadAccountsTaskTaskMessagesInner) HasLocalizedText() bool {
	if o != nil && !IsNil(o.LocalizedText) {
		return true
	}

	return false
}

// SetLocalizedText gets a reference to the given string and assigns it to the LocalizedText field.
func (o *LoadAccountsTaskTaskMessagesInner) SetLocalizedText(v string) {
	o.LocalizedText = &v
}

func (o LoadAccountsTaskTaskMessagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadAccountsTaskTaskMessagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.LocalizedText) {
		toSerialize["localizedText"] = o.LocalizedText
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadAccountsTaskTaskMessagesInner) UnmarshalJSON(data []byte) (err error) {
	varLoadAccountsTaskTaskMessagesInner := _LoadAccountsTaskTaskMessagesInner{}

	err = json.Unmarshal(data, &varLoadAccountsTaskTaskMessagesInner)

	if err != nil {
		return err
	}

	*o = LoadAccountsTaskTaskMessagesInner(varLoadAccountsTaskTaskMessagesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "error")
		delete(additionalProperties, "warning")
		delete(additionalProperties, "key")
		delete(additionalProperties, "localizedText")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadAccountsTaskTaskMessagesInner struct {
	value *LoadAccountsTaskTaskMessagesInner
	isSet bool
}

func (v NullableLoadAccountsTaskTaskMessagesInner) Get() *LoadAccountsTaskTaskMessagesInner {
	return v.value
}

func (v *NullableLoadAccountsTaskTaskMessagesInner) Set(val *LoadAccountsTaskTaskMessagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadAccountsTaskTaskMessagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadAccountsTaskTaskMessagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadAccountsTaskTaskMessagesInner(val *LoadAccountsTaskTaskMessagesInner) *NullableLoadAccountsTaskTaskMessagesInner {
	return &NullableLoadAccountsTaskTaskMessagesInner{value: val, isSet: true}
}

func (v NullableLoadAccountsTaskTaskMessagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadAccountsTaskTaskMessagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


