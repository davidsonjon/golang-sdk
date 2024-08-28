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

// checks if the NonEmployeeSourceWithCloudExternalId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeSourceWithCloudExternalId{}

// NonEmployeeSourceWithCloudExternalId struct for NonEmployeeSourceWithCloudExternalId
type NonEmployeeSourceWithCloudExternalId struct {
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
	// Legacy ID used for sources from the V1 API. This attribute will be removed from a future version of the API and will not be considered a breaking change. No clients should rely on this ID always being present.
	CloudExternalId *string `json:"cloudExternalId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceWithCloudExternalId NonEmployeeSourceWithCloudExternalId

// NewNonEmployeeSourceWithCloudExternalId instantiates a new NonEmployeeSourceWithCloudExternalId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceWithCloudExternalId() *NonEmployeeSourceWithCloudExternalId {
	this := NonEmployeeSourceWithCloudExternalId{}
	return &this
}

// NewNonEmployeeSourceWithCloudExternalIdWithDefaults instantiates a new NonEmployeeSourceWithCloudExternalId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceWithCloudExternalIdWithDefaults() *NonEmployeeSourceWithCloudExternalId {
	this := NonEmployeeSourceWithCloudExternalId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeSourceWithCloudExternalId) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeSourceWithCloudExternalId) SetSourceId(v string) {
	o.SourceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NonEmployeeSourceWithCloudExternalId) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonEmployeeSourceWithCloudExternalId) SetDescription(v string) {
	o.Description = &v
}

// GetApprovers returns the Approvers field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetApprovers() []NonEmployeeIdentityReferenceWithId {
	if o == nil || IsNil(o.Approvers) {
		var ret []NonEmployeeIdentityReferenceWithId
		return ret
	}
	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetApproversOk() ([]NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || IsNil(o.Approvers) {
		return nil, false
	}
	return o.Approvers, true
}

// HasApprovers returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasApprovers() bool {
	if o != nil && !IsNil(o.Approvers) {
		return true
	}

	return false
}

// SetApprovers gets a reference to the given []NonEmployeeIdentityReferenceWithId and assigns it to the Approvers field.
func (o *NonEmployeeSourceWithCloudExternalId) SetApprovers(v []NonEmployeeIdentityReferenceWithId) {
	o.Approvers = v
}

// GetAccountManagers returns the AccountManagers field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetAccountManagers() []NonEmployeeIdentityReferenceWithId {
	if o == nil || IsNil(o.AccountManagers) {
		var ret []NonEmployeeIdentityReferenceWithId
		return ret
	}
	return o.AccountManagers
}

// GetAccountManagersOk returns a tuple with the AccountManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetAccountManagersOk() ([]NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || IsNil(o.AccountManagers) {
		return nil, false
	}
	return o.AccountManagers, true
}

// HasAccountManagers returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasAccountManagers() bool {
	if o != nil && !IsNil(o.AccountManagers) {
		return true
	}

	return false
}

// SetAccountManagers gets a reference to the given []NonEmployeeIdentityReferenceWithId and assigns it to the AccountManagers field.
func (o *NonEmployeeSourceWithCloudExternalId) SetAccountManagers(v []NonEmployeeIdentityReferenceWithId) {
	o.AccountManagers = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeSourceWithCloudExternalId) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeSourceWithCloudExternalId) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCloudExternalId returns the CloudExternalId field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithCloudExternalId) GetCloudExternalId() string {
	if o == nil || IsNil(o.CloudExternalId) {
		var ret string
		return ret
	}
	return *o.CloudExternalId
}

// GetCloudExternalIdOk returns a tuple with the CloudExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithCloudExternalId) GetCloudExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudExternalId) {
		return nil, false
	}
	return o.CloudExternalId, true
}

// HasCloudExternalId returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithCloudExternalId) HasCloudExternalId() bool {
	if o != nil && !IsNil(o.CloudExternalId) {
		return true
	}

	return false
}

// SetCloudExternalId gets a reference to the given string and assigns it to the CloudExternalId field.
func (o *NonEmployeeSourceWithCloudExternalId) SetCloudExternalId(v string) {
	o.CloudExternalId = &v
}

func (o NonEmployeeSourceWithCloudExternalId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeSourceWithCloudExternalId) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CloudExternalId) {
		toSerialize["cloudExternalId"] = o.CloudExternalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeSourceWithCloudExternalId) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeSourceWithCloudExternalId := _NonEmployeeSourceWithCloudExternalId{}

	err = json.Unmarshal(data, &varNonEmployeeSourceWithCloudExternalId)

	if err != nil {
		return err
	}

	*o = NonEmployeeSourceWithCloudExternalId(varNonEmployeeSourceWithCloudExternalId)

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
		delete(additionalProperties, "cloudExternalId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceWithCloudExternalId struct {
	value *NonEmployeeSourceWithCloudExternalId
	isSet bool
}

func (v NullableNonEmployeeSourceWithCloudExternalId) Get() *NonEmployeeSourceWithCloudExternalId {
	return v.value
}

func (v *NullableNonEmployeeSourceWithCloudExternalId) Set(val *NonEmployeeSourceWithCloudExternalId) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceWithCloudExternalId) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceWithCloudExternalId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceWithCloudExternalId(val *NonEmployeeSourceWithCloudExternalId) *NullableNonEmployeeSourceWithCloudExternalId {
	return &NullableNonEmployeeSourceWithCloudExternalId{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceWithCloudExternalId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceWithCloudExternalId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


