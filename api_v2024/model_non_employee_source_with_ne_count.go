/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the NonEmployeeSourceWithNECount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeSourceWithNECount{}

// NonEmployeeSourceWithNECount struct for NonEmployeeSourceWithNECount
type NonEmployeeSourceWithNECount struct {
	// Non-Employee source id.
	Id *string `json:"id,omitempty"`
	// Source Id associated with this non-employee source.
	SourceId *string `json:"sourceId,omitempty"`
	// Source name associated with this non-employee source.
	Name *string `json:"name,omitempty"`
	// Source description associated with this non-employee source.
	Description *string `json:"description,omitempty"`
	// List of approvers
	Approvers []NonEmployeeIdentityReferenceWithId `json:"approvers,omitempty"`
	// List of account managers
	AccountManagers []NonEmployeeIdentityReferenceWithId `json:"accountManagers,omitempty"`
	// When the request was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	// Number of non-employee records associated with this source.
	NonEmployeeCount NullableInt32 `json:"nonEmployeeCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceWithNECount NonEmployeeSourceWithNECount

// NewNonEmployeeSourceWithNECount instantiates a new NonEmployeeSourceWithNECount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceWithNECount() *NonEmployeeSourceWithNECount {
	this := NonEmployeeSourceWithNECount{}
	return &this
}

// NewNonEmployeeSourceWithNECountWithDefaults instantiates a new NonEmployeeSourceWithNECount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceWithNECountWithDefaults() *NonEmployeeSourceWithNECount {
	this := NonEmployeeSourceWithNECount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeSourceWithNECount) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeSourceWithNECount) SetSourceId(v string) {
	o.SourceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NonEmployeeSourceWithNECount) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonEmployeeSourceWithNECount) SetDescription(v string) {
	o.Description = &v
}

// GetApprovers returns the Approvers field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetApprovers() []NonEmployeeIdentityReferenceWithId {
	if o == nil || IsNil(o.Approvers) {
		var ret []NonEmployeeIdentityReferenceWithId
		return ret
	}
	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetApproversOk() ([]NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || IsNil(o.Approvers) {
		return nil, false
	}
	return o.Approvers, true
}

// HasApprovers returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasApprovers() bool {
	if o != nil && !IsNil(o.Approvers) {
		return true
	}

	return false
}

// SetApprovers gets a reference to the given []NonEmployeeIdentityReferenceWithId and assigns it to the Approvers field.
func (o *NonEmployeeSourceWithNECount) SetApprovers(v []NonEmployeeIdentityReferenceWithId) {
	o.Approvers = v
}

// GetAccountManagers returns the AccountManagers field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetAccountManagers() []NonEmployeeIdentityReferenceWithId {
	if o == nil || IsNil(o.AccountManagers) {
		var ret []NonEmployeeIdentityReferenceWithId
		return ret
	}
	return o.AccountManagers
}

// GetAccountManagersOk returns a tuple with the AccountManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetAccountManagersOk() ([]NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || IsNil(o.AccountManagers) {
		return nil, false
	}
	return o.AccountManagers, true
}

// HasAccountManagers returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasAccountManagers() bool {
	if o != nil && !IsNil(o.AccountManagers) {
		return true
	}

	return false
}

// SetAccountManagers gets a reference to the given []NonEmployeeIdentityReferenceWithId and assigns it to the AccountManagers field.
func (o *NonEmployeeSourceWithNECount) SetAccountManagers(v []NonEmployeeIdentityReferenceWithId) {
	o.AccountManagers = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeSourceWithNECount) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECount) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECount) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeSourceWithNECount) SetCreated(v time.Time) {
	o.Created = &v
}

// GetNonEmployeeCount returns the NonEmployeeCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NonEmployeeSourceWithNECount) GetNonEmployeeCount() int32 {
	if o == nil || IsNil(o.NonEmployeeCount.Get()) {
		var ret int32
		return ret
	}
	return *o.NonEmployeeCount.Get()
}

// GetNonEmployeeCountOk returns a tuple with the NonEmployeeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NonEmployeeSourceWithNECount) GetNonEmployeeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonEmployeeCount.Get(), o.NonEmployeeCount.IsSet()
}

// HasNonEmployeeCount returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECount) HasNonEmployeeCount() bool {
	if o != nil && o.NonEmployeeCount.IsSet() {
		return true
	}

	return false
}

// SetNonEmployeeCount gets a reference to the given NullableInt32 and assigns it to the NonEmployeeCount field.
func (o *NonEmployeeSourceWithNECount) SetNonEmployeeCount(v int32) {
	o.NonEmployeeCount.Set(&v)
}
// SetNonEmployeeCountNil sets the value for NonEmployeeCount to be an explicit nil
func (o *NonEmployeeSourceWithNECount) SetNonEmployeeCountNil() {
	o.NonEmployeeCount.Set(nil)
}

// UnsetNonEmployeeCount ensures that no value is present for NonEmployeeCount, not even an explicit nil
func (o *NonEmployeeSourceWithNECount) UnsetNonEmployeeCount() {
	o.NonEmployeeCount.Unset()
}

func (o NonEmployeeSourceWithNECount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeSourceWithNECount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Approvers) {
		toSerialize["approvers"] = o.Approvers
	}
	if !IsNil(o.AccountManagers) {
		toSerialize["accountManagers"] = o.AccountManagers
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.NonEmployeeCount.IsSet() {
		toSerialize["nonEmployeeCount"] = o.NonEmployeeCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeSourceWithNECount) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeSourceWithNECount := _NonEmployeeSourceWithNECount{}

	err = json.Unmarshal(data, &varNonEmployeeSourceWithNECount)

	if err != nil {
		return err
	}

	*o = NonEmployeeSourceWithNECount(varNonEmployeeSourceWithNECount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approvers")
		delete(additionalProperties, "accountManagers")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		delete(additionalProperties, "nonEmployeeCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceWithNECount struct {
	value *NonEmployeeSourceWithNECount
	isSet bool
}

func (v NullableNonEmployeeSourceWithNECount) Get() *NonEmployeeSourceWithNECount {
	return v.value
}

func (v *NullableNonEmployeeSourceWithNECount) Set(val *NonEmployeeSourceWithNECount) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceWithNECount) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceWithNECount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceWithNECount(val *NonEmployeeSourceWithNECount) *NullableNonEmployeeSourceWithNECount {
	return &NullableNonEmployeeSourceWithNECount{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceWithNECount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceWithNECount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


