/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the CustomPasswordInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPasswordInstruction{}

// CustomPasswordInstruction struct for CustomPasswordInstruction
type CustomPasswordInstruction struct {
	// The page ID that represents the page for forget user name, reset password and unlock account flow.
	PageId *string `json:"pageId,omitempty"`
	// The custom instructions for the specified page. Allow basic HTML format and maximum length is 1000 characters. The custom instructions will be sanitized to avoid attacks. If the customization text includes a link, like <A HREF=\\\"URL\\\">...</A> clicking on this will open the link on the current browser page. If you want your link to be redirected to a different page, please redirect it to \"_blank\" like this: <a href=\\\"URL\" target=\\\"_blank\\\" >link</a>. This will open a new tab when the link is clicked. Notice we're only supporting _blank as the redirection target.
	PageContent *string `json:"pageContent,omitempty"`
	// The locale for the custom instructions, a BCP47 language tag. The default value is \\\"default\\\".
	Locale *string `json:"locale,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomPasswordInstruction CustomPasswordInstruction

// NewCustomPasswordInstruction instantiates a new CustomPasswordInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPasswordInstruction() *CustomPasswordInstruction {
	this := CustomPasswordInstruction{}
	return &this
}

// NewCustomPasswordInstructionWithDefaults instantiates a new CustomPasswordInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPasswordInstructionWithDefaults() *CustomPasswordInstruction {
	this := CustomPasswordInstruction{}
	return &this
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *CustomPasswordInstruction) GetPageId() string {
	if o == nil || IsNil(o.PageId) {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPasswordInstruction) GetPageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PageId) {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *CustomPasswordInstruction) HasPageId() bool {
	if o != nil && !IsNil(o.PageId) {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *CustomPasswordInstruction) SetPageId(v string) {
	o.PageId = &v
}

// GetPageContent returns the PageContent field value if set, zero value otherwise.
func (o *CustomPasswordInstruction) GetPageContent() string {
	if o == nil || IsNil(o.PageContent) {
		var ret string
		return ret
	}
	return *o.PageContent
}

// GetPageContentOk returns a tuple with the PageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPasswordInstruction) GetPageContentOk() (*string, bool) {
	if o == nil || IsNil(o.PageContent) {
		return nil, false
	}
	return o.PageContent, true
}

// HasPageContent returns a boolean if a field has been set.
func (o *CustomPasswordInstruction) HasPageContent() bool {
	if o != nil && !IsNil(o.PageContent) {
		return true
	}

	return false
}

// SetPageContent gets a reference to the given string and assigns it to the PageContent field.
func (o *CustomPasswordInstruction) SetPageContent(v string) {
	o.PageContent = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CustomPasswordInstruction) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPasswordInstruction) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *CustomPasswordInstruction) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CustomPasswordInstruction) SetLocale(v string) {
	o.Locale = &v
}

func (o CustomPasswordInstruction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPasswordInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageId) {
		toSerialize["pageId"] = o.PageId
	}
	if !IsNil(o.PageContent) {
		toSerialize["pageContent"] = o.PageContent
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomPasswordInstruction) UnmarshalJSON(data []byte) (err error) {
	varCustomPasswordInstruction := _CustomPasswordInstruction{}

	err = json.Unmarshal(data, &varCustomPasswordInstruction)

	if err != nil {
		return err
	}

	*o = CustomPasswordInstruction(varCustomPasswordInstruction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pageId")
		delete(additionalProperties, "pageContent")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomPasswordInstruction struct {
	value *CustomPasswordInstruction
	isSet bool
}

func (v NullableCustomPasswordInstruction) Get() *CustomPasswordInstruction {
	return v.value
}

func (v *NullableCustomPasswordInstruction) Set(val *CustomPasswordInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPasswordInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPasswordInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPasswordInstruction(val *CustomPasswordInstruction) *NullableCustomPasswordInstruction {
	return &NullableCustomPasswordInstruction{value: val, isSet: true}
}

func (v NullableCustomPasswordInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPasswordInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


