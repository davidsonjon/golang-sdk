/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
)

// checks if the BaseAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseAccess{}

// BaseAccess struct for BaseAccess
type BaseAccess struct {
	// The unique ID of the referenced object.
	Id *string `json:"id,omitempty"`
	// The human readable name of the referenced object.
	Name *string `json:"name,omitempty"`
	// Access item's description.
	Description *string `json:"description,omitempty"`
	// ISO-8601 date-time referring to the time when the object was created.
	Created NullableTime `json:"created,omitempty"`
	// ISO-8601 date-time referring to the time when the object was last modified.
	Modified NullableTime `json:"modified,omitempty"`
	// ISO-8601 date-time referring to the date-time when object was queued to be synced into search database for use in the search API.   This date-time changes anytime there is an update to the object, which triggers a synchronization event being sent to the search database.  There may be some delay between the `synced` time and the time when the updated data is actually available in the search API. 
	Synced NullableTime `json:"synced,omitempty"`
	// Indicates whether the access item is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates whether the access item can be requested.
	Requestable *bool `json:"requestable,omitempty"`
	// Indicates whether comments are required for requests to access the item.
	RequestCommentsRequired *bool `json:"requestCommentsRequired,omitempty"`
	Owner *BaseAccessAllOfOwner `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseAccess BaseAccess

// NewBaseAccess instantiates a new BaseAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAccess() *BaseAccess {
	this := BaseAccess{}
	var enabled bool = false
	this.Enabled = &enabled
	var requestable bool = true
	this.Requestable = &requestable
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	return &this
}

// NewBaseAccessWithDefaults instantiates a new BaseAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAccessWithDefaults() *BaseAccess {
	this := BaseAccess{}
	var enabled bool = false
	this.Enabled = &enabled
	var requestable bool = true
	this.Requestable = &requestable
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseAccess) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseAccess) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseAccess) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseAccess) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseAccess) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseAccess) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseAccess) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseAccess) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseAccess) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseAccess) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseAccess) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *BaseAccess) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *BaseAccess) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *BaseAccess) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *BaseAccess) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseAccess) GetModified() time.Time {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseAccess) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *BaseAccess) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *BaseAccess) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *BaseAccess) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *BaseAccess) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseAccess) GetSynced() time.Time {
	if o == nil || IsNil(o.Synced.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseAccess) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *BaseAccess) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *BaseAccess) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *BaseAccess) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *BaseAccess) UnsetSynced() {
	o.Synced.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BaseAccess) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BaseAccess) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BaseAccess) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *BaseAccess) GetRequestable() bool {
	if o == nil || IsNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetRequestableOk() (*bool, bool) {
	if o == nil || IsNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *BaseAccess) HasRequestable() bool {
	if o != nil && !IsNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *BaseAccess) SetRequestable(v bool) {
	o.Requestable = &v
}

// GetRequestCommentsRequired returns the RequestCommentsRequired field value if set, zero value otherwise.
func (o *BaseAccess) GetRequestCommentsRequired() bool {
	if o == nil || IsNil(o.RequestCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentsRequired
}

// GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetRequestCommentsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestCommentsRequired) {
		return nil, false
	}
	return o.RequestCommentsRequired, true
}

// HasRequestCommentsRequired returns a boolean if a field has been set.
func (o *BaseAccess) HasRequestCommentsRequired() bool {
	if o != nil && !IsNil(o.RequestCommentsRequired) {
		return true
	}

	return false
}

// SetRequestCommentsRequired gets a reference to the given bool and assigns it to the RequestCommentsRequired field.
func (o *BaseAccess) SetRequestCommentsRequired(v bool) {
	o.RequestCommentsRequired = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BaseAccess) GetOwner() BaseAccessAllOfOwner {
	if o == nil || IsNil(o.Owner) {
		var ret BaseAccessAllOfOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccess) GetOwnerOk() (*BaseAccessAllOfOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BaseAccess) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BaseAccessAllOfOwner and assigns it to the Owner field.
func (o *BaseAccess) SetOwner(v BaseAccessAllOfOwner) {
	o.Owner = &v
}

func (o BaseAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if !IsNil(o.RequestCommentsRequired) {
		toSerialize["requestCommentsRequired"] = o.RequestCommentsRequired
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseAccess) UnmarshalJSON(data []byte) (err error) {
	varBaseAccess := _BaseAccess{}

	err = json.Unmarshal(data, &varBaseAccess)

	if err != nil {
		return err
	}

	*o = BaseAccess(varBaseAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "requestCommentsRequired")
		delete(additionalProperties, "owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseAccess struct {
	value *BaseAccess
	isSet bool
}

func (v NullableBaseAccess) Get() *BaseAccess {
	return v.value
}

func (v *NullableBaseAccess) Set(val *BaseAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAccess(val *BaseAccess) *NullableBaseAccess {
	return &NullableBaseAccess{value: val, isSet: true}
}

func (v NullableBaseAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


