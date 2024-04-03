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

// checks if the Invocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invocation{}

// Invocation struct for Invocation
type Invocation struct {
	// Invocation ID
	Id *string `json:"id,omitempty"`
	// Trigger ID
	TriggerId *string `json:"triggerId,omitempty"`
	// Unique invocation secret.
	Secret *string `json:"secret,omitempty"`
	// JSON map of invocation metadata.
	ContentJson map[string]interface{} `json:"contentJson,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Invocation Invocation

// NewInvocation instantiates a new Invocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvocation() *Invocation {
	this := Invocation{}
	return &this
}

// NewInvocationWithDefaults instantiates a new Invocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvocationWithDefaults() *Invocation {
	this := Invocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invocation) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invocation) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invocation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invocation) SetId(v string) {
	o.Id = &v
}

// GetTriggerId returns the TriggerId field value if set, zero value otherwise.
func (o *Invocation) GetTriggerId() string {
	if o == nil || isNil(o.TriggerId) {
		var ret string
		return ret
	}
	return *o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invocation) GetTriggerIdOk() (*string, bool) {
	if o == nil || isNil(o.TriggerId) {
		return nil, false
	}
	return o.TriggerId, true
}

// HasTriggerId returns a boolean if a field has been set.
func (o *Invocation) HasTriggerId() bool {
	if o != nil && !isNil(o.TriggerId) {
		return true
	}

	return false
}

// SetTriggerId gets a reference to the given string and assigns it to the TriggerId field.
func (o *Invocation) SetTriggerId(v string) {
	o.TriggerId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Invocation) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invocation) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Invocation) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *Invocation) SetSecret(v string) {
	o.Secret = &v
}

// GetContentJson returns the ContentJson field value if set, zero value otherwise.
func (o *Invocation) GetContentJson() map[string]interface{} {
	if o == nil || isNil(o.ContentJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.ContentJson
}

// GetContentJsonOk returns a tuple with the ContentJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invocation) GetContentJsonOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ContentJson) {
		return map[string]interface{}{}, false
	}
	return o.ContentJson, true
}

// HasContentJson returns a boolean if a field has been set.
func (o *Invocation) HasContentJson() bool {
	if o != nil && !isNil(o.ContentJson) {
		return true
	}

	return false
}

// SetContentJson gets a reference to the given map[string]interface{} and assigns it to the ContentJson field.
func (o *Invocation) SetContentJson(v map[string]interface{}) {
	o.ContentJson = v
}

func (o Invocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.TriggerId) {
		toSerialize["triggerId"] = o.TriggerId
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.ContentJson) {
		toSerialize["contentJson"] = o.ContentJson
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Invocation) UnmarshalJSON(bytes []byte) (err error) {
	varInvocation := _Invocation{}

	if err = json.Unmarshal(bytes, &varInvocation); err == nil {
	*o = Invocation(varInvocation)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "triggerId")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "contentJson")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvocation struct {
	value *Invocation
	isSet bool
}

func (v NullableInvocation) Get() *Invocation {
	return v.value
}

func (v *NullableInvocation) Set(val *Invocation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvocation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvocation(val *Invocation) *NullableInvocation {
	return &NullableInvocation{value: val, isSet: true}
}

func (v NullableInvocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


