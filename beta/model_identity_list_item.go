/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the IdentityListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityListItem{}

// IdentityListItem struct for IdentityListItem
type IdentityListItem struct {
	// the identity ID
	Id *string `json:"id,omitempty"`
	// the display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// the first name of the identity
	FirstName *string `json:"firstName,omitempty"`
	// the last name of the identity
	LastName *string `json:"lastName,omitempty"`
	// indicates if an identity is active or not
	Active *bool `json:"active,omitempty"`
	// the date when the identity was deleted
	DeletedDate NullableString `json:"deletedDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityListItem IdentityListItem

// NewIdentityListItem instantiates a new IdentityListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityListItem() *IdentityListItem {
	this := IdentityListItem{}
	var active bool = true
	this.Active = &active
	return &this
}

// NewIdentityListItemWithDefaults instantiates a new IdentityListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityListItemWithDefaults() *IdentityListItem {
	this := IdentityListItem{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityListItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityListItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityListItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityListItem) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentityListItem) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityListItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityListItem) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityListItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IdentityListItem) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityListItem) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IdentityListItem) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IdentityListItem) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *IdentityListItem) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityListItem) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *IdentityListItem) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *IdentityListItem) SetLastName(v string) {
	o.LastName = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *IdentityListItem) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityListItem) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *IdentityListItem) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *IdentityListItem) SetActive(v bool) {
	o.Active = &v
}

// GetDeletedDate returns the DeletedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityListItem) GetDeletedDate() string {
	if o == nil || isNil(o.DeletedDate.Get()) {
		var ret string
		return ret
	}
	return *o.DeletedDate.Get()
}

// GetDeletedDateOk returns a tuple with the DeletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityListItem) GetDeletedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedDate.Get(), o.DeletedDate.IsSet()
}

// HasDeletedDate returns a boolean if a field has been set.
func (o *IdentityListItem) HasDeletedDate() bool {
	if o != nil && o.DeletedDate.IsSet() {
		return true
	}

	return false
}

// SetDeletedDate gets a reference to the given NullableString and assigns it to the DeletedDate field.
func (o *IdentityListItem) SetDeletedDate(v string) {
	o.DeletedDate.Set(&v)
}
// SetDeletedDateNil sets the value for DeletedDate to be an explicit nil
func (o *IdentityListItem) SetDeletedDateNil() {
	o.DeletedDate.Set(nil)
}

// UnsetDeletedDate ensures that no value is present for DeletedDate, not even an explicit nil
func (o *IdentityListItem) UnsetDeletedDate() {
	o.DeletedDate.Unset()
}

func (o IdentityListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !isNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if o.DeletedDate.IsSet() {
		toSerialize["deletedDate"] = o.DeletedDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityListItem) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityListItem := _IdentityListItem{}

	if err = json.Unmarshal(bytes, &varIdentityListItem); err == nil {
	*o = IdentityListItem(varIdentityListItem)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "active")
		delete(additionalProperties, "deletedDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityListItem struct {
	value *IdentityListItem
	isSet bool
}

func (v NullableIdentityListItem) Get() *IdentityListItem {
	return v.value
}

func (v *NullableIdentityListItem) Set(val *IdentityListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityListItem(val *IdentityListItem) *NullableIdentityListItem {
	return &NullableIdentityListItem{value: val, isSet: true}
}

func (v NullableIdentityListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


