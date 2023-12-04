/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the CreatePersonalAccessTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePersonalAccessTokenRequest{}

// CreatePersonalAccessTokenRequest Object for specifying the name of a personal access token to create
type CreatePersonalAccessTokenRequest struct {
	// The name of the personal access token (PAT) to be created. Cannot be the same as another PAT owned by the user for whom this PAT is being created.
	Name string `json:"name"`
	// Scopes of the personal  access token. If no scope is specified, the token will be created with the default scope \"sp:scopes:all\". This means the personal access token will have all the rights of the owner who created it.
	Scope []string `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreatePersonalAccessTokenRequest CreatePersonalAccessTokenRequest

// NewCreatePersonalAccessTokenRequest instantiates a new CreatePersonalAccessTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePersonalAccessTokenRequest(name string) *CreatePersonalAccessTokenRequest {
	this := CreatePersonalAccessTokenRequest{}
	this.Name = name
	return &this
}

// NewCreatePersonalAccessTokenRequestWithDefaults instantiates a new CreatePersonalAccessTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePersonalAccessTokenRequestWithDefaults() *CreatePersonalAccessTokenRequest {
	this := CreatePersonalAccessTokenRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePersonalAccessTokenRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePersonalAccessTokenRequest) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePersonalAccessTokenRequest) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePersonalAccessTokenRequest) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CreatePersonalAccessTokenRequest) HasScope() bool {
	if o != nil && isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *CreatePersonalAccessTokenRequest) SetScope(v []string) {
	o.Scope = v
}

func (o CreatePersonalAccessTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePersonalAccessTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreatePersonalAccessTokenRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreatePersonalAccessTokenRequest := _CreatePersonalAccessTokenRequest{}

	if err = json.Unmarshal(bytes, &varCreatePersonalAccessTokenRequest); err == nil {
	*o = CreatePersonalAccessTokenRequest(varCreatePersonalAccessTokenRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePersonalAccessTokenRequest struct {
	value *CreatePersonalAccessTokenRequest
	isSet bool
}

func (v NullableCreatePersonalAccessTokenRequest) Get() *CreatePersonalAccessTokenRequest {
	return v.value
}

func (v *NullableCreatePersonalAccessTokenRequest) Set(val *CreatePersonalAccessTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePersonalAccessTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePersonalAccessTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePersonalAccessTokenRequest(val *CreatePersonalAccessTokenRequest) *NullableCreatePersonalAccessTokenRequest {
	return &NullableCreatePersonalAccessTokenRequest{value: val, isSet: true}
}

func (v NullableCreatePersonalAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePersonalAccessTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


