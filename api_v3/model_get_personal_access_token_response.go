/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the GetPersonalAccessTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPersonalAccessTokenResponse{}

// GetPersonalAccessTokenResponse struct for GetPersonalAccessTokenResponse
type GetPersonalAccessTokenResponse struct {
	// The ID of the personal access token (to be used as the username for Basic Auth).
	Id string `json:"id"`
	// The name of the personal access token. Cannot be the same as other personal access tokens owned by a user.
	Name string `json:"name"`
	// Scopes of the personal  access token.
	Scope []string `json:"scope"`
	Owner PatOwner `json:"owner"`
	// The date and time, down to the millisecond, when this personal access token was created.
	Created time.Time `json:"created"`
	// The date and time, down to the millisecond, when this personal access token was last used to generate an access token. This timestamp does not get updated on every PAT usage, but only once a day. This property can be useful for identifying which PATs are no longer actively used and can be removed.
	LastUsed NullableTime `json:"lastUsed,omitempty"`
	// If true, this token is managed by the SailPoint platform, and is not visible in the user interface. For example, Workflows will create managed personal access tokens for users who create workflows.
	Managed *bool `json:"managed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPersonalAccessTokenResponse GetPersonalAccessTokenResponse

// NewGetPersonalAccessTokenResponse instantiates a new GetPersonalAccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPersonalAccessTokenResponse(id string, name string, scope []string, owner PatOwner, created time.Time) *GetPersonalAccessTokenResponse {
	this := GetPersonalAccessTokenResponse{}
	this.Id = id
	this.Name = name
	this.Scope = scope
	this.Owner = owner
	this.Created = created
	var managed bool = false
	this.Managed = &managed
	return &this
}

// NewGetPersonalAccessTokenResponseWithDefaults instantiates a new GetPersonalAccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPersonalAccessTokenResponseWithDefaults() *GetPersonalAccessTokenResponse {
	this := GetPersonalAccessTokenResponse{}
	var managed bool = false
	this.Managed = &managed
	return &this
}

// GetId returns the Id field value
func (o *GetPersonalAccessTokenResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokenResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetPersonalAccessTokenResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetPersonalAccessTokenResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokenResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetPersonalAccessTokenResponse) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GetPersonalAccessTokenResponse) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPersonalAccessTokenResponse) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *GetPersonalAccessTokenResponse) SetScope(v []string) {
	o.Scope = v
}

// GetOwner returns the Owner field value
func (o *GetPersonalAccessTokenResponse) GetOwner() PatOwner {
	if o == nil {
		var ret PatOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokenResponse) GetOwnerOk() (*PatOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *GetPersonalAccessTokenResponse) SetOwner(v PatOwner) {
	o.Owner = v
}

// GetCreated returns the Created field value
func (o *GetPersonalAccessTokenResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokenResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GetPersonalAccessTokenResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPersonalAccessTokenResponse) GetLastUsed() time.Time {
	if o == nil || isNil(o.LastUsed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPersonalAccessTokenResponse) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// HasLastUsed returns a boolean if a field has been set.
func (o *GetPersonalAccessTokenResponse) HasLastUsed() bool {
	if o != nil && o.LastUsed.IsSet() {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given NullableTime and assigns it to the LastUsed field.
func (o *GetPersonalAccessTokenResponse) SetLastUsed(v time.Time) {
	o.LastUsed.Set(&v)
}
// SetLastUsedNil sets the value for LastUsed to be an explicit nil
func (o *GetPersonalAccessTokenResponse) SetLastUsedNil() {
	o.LastUsed.Set(nil)
}

// UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
func (o *GetPersonalAccessTokenResponse) UnsetLastUsed() {
	o.LastUsed.Unset()
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *GetPersonalAccessTokenResponse) GetManaged() bool {
	if o == nil || isNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokenResponse) GetManagedOk() (*bool, bool) {
	if o == nil || isNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *GetPersonalAccessTokenResponse) HasManaged() bool {
	if o != nil && !isNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *GetPersonalAccessTokenResponse) SetManaged(v bool) {
	o.Managed = &v
}

func (o GetPersonalAccessTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPersonalAccessTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["owner"] = o.Owner
	toSerialize["created"] = o.Created
	if o.LastUsed.IsSet() {
		toSerialize["lastUsed"] = o.LastUsed.Get()
	}
	if !isNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPersonalAccessTokenResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"scope",
		"owner",
		"created",
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

	varGetPersonalAccessTokenResponse := _GetPersonalAccessTokenResponse{}

	if err = json.Unmarshal(bytes, &varGetPersonalAccessTokenResponse); err == nil {
	*o = GetPersonalAccessTokenResponse(varGetPersonalAccessTokenResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUsed")
		delete(additionalProperties, "managed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPersonalAccessTokenResponse struct {
	value *GetPersonalAccessTokenResponse
	isSet bool
}

func (v NullableGetPersonalAccessTokenResponse) Get() *GetPersonalAccessTokenResponse {
	return v.value
}

func (v *NullableGetPersonalAccessTokenResponse) Set(val *GetPersonalAccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPersonalAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPersonalAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPersonalAccessTokenResponse(val *GetPersonalAccessTokenResponse) *NullableGetPersonalAccessTokenResponse {
	return &NullableGetPersonalAccessTokenResponse{value: val, isSet: true}
}

func (v NullableGetPersonalAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPersonalAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


