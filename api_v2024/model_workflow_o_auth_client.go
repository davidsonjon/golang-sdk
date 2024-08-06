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

// checks if the WorkflowOAuthClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowOAuthClient{}

// WorkflowOAuthClient struct for WorkflowOAuthClient
type WorkflowOAuthClient struct {
	// OAuth client ID for the trigger. This is a UUID generated upon creation.
	Id *string `json:"id,omitempty"`
	// OAuthClient secret.
	Secret *string `json:"secret,omitempty"`
	// URL for the external trigger to invoke
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowOAuthClient WorkflowOAuthClient

// NewWorkflowOAuthClient instantiates a new WorkflowOAuthClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOAuthClient() *WorkflowOAuthClient {
	this := WorkflowOAuthClient{}
	return &this
}

// NewWorkflowOAuthClientWithDefaults instantiates a new WorkflowOAuthClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOAuthClientWithDefaults() *WorkflowOAuthClient {
	this := WorkflowOAuthClient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowOAuthClient) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOAuthClient) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowOAuthClient) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowOAuthClient) SetId(v string) {
	o.Id = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WorkflowOAuthClient) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOAuthClient) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WorkflowOAuthClient) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WorkflowOAuthClient) SetSecret(v string) {
	o.Secret = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WorkflowOAuthClient) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOAuthClient) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WorkflowOAuthClient) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WorkflowOAuthClient) SetUrl(v string) {
	o.Url = &v
}

func (o WorkflowOAuthClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowOAuthClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowOAuthClient) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowOAuthClient := _WorkflowOAuthClient{}

	if err = json.Unmarshal(bytes, &varWorkflowOAuthClient); err == nil {
	*o = WorkflowOAuthClient(varWorkflowOAuthClient)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowOAuthClient struct {
	value *WorkflowOAuthClient
	isSet bool
}

func (v NullableWorkflowOAuthClient) Get() *WorkflowOAuthClient {
	return v.value
}

func (v *NullableWorkflowOAuthClient) Set(val *WorkflowOAuthClient) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOAuthClient) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOAuthClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOAuthClient(val *WorkflowOAuthClient) *NullableWorkflowOAuthClient {
	return &NullableWorkflowOAuthClient{value: val, isSet: true}
}

func (v NullableWorkflowOAuthClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOAuthClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


