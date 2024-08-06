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

// checks if the NonEmployeeApprovalItemBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeApprovalItemBase{}

// NonEmployeeApprovalItemBase struct for NonEmployeeApprovalItemBase
type NonEmployeeApprovalItemBase struct {
	// Non-Employee approval item id
	Id *string `json:"id,omitempty"`
	Approver *NonEmployeeIdentityReferenceWithId `json:"approver,omitempty"`
	// Requested identity account name
	AccountName *string `json:"accountName,omitempty"`
	ApprovalStatus *ApprovalStatus `json:"approvalStatus,omitempty"`
	// Approval order
	ApprovalOrder *float32 `json:"approvalOrder,omitempty"`
	// comment of approver
	Comment *string `json:"comment,omitempty"`
	// When the request was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeApprovalItemBase NonEmployeeApprovalItemBase

// NewNonEmployeeApprovalItemBase instantiates a new NonEmployeeApprovalItemBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeApprovalItemBase() *NonEmployeeApprovalItemBase {
	this := NonEmployeeApprovalItemBase{}
	return &this
}

// NewNonEmployeeApprovalItemBaseWithDefaults instantiates a new NonEmployeeApprovalItemBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeApprovalItemBaseWithDefaults() *NonEmployeeApprovalItemBase {
	this := NonEmployeeApprovalItemBase{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeApprovalItemBase) SetId(v string) {
	o.Id = &v
}

// GetApprover returns the Approver field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetApprover() NonEmployeeIdentityReferenceWithId {
	if o == nil || isNil(o.Approver) {
		var ret NonEmployeeIdentityReferenceWithId
		return ret
	}
	return *o.Approver
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetApproverOk() (*NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || isNil(o.Approver) {
		return nil, false
	}
	return o.Approver, true
}

// HasApprover returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasApprover() bool {
	if o != nil && !isNil(o.Approver) {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NonEmployeeIdentityReferenceWithId and assigns it to the Approver field.
func (o *NonEmployeeApprovalItemBase) SetApprover(v NonEmployeeIdentityReferenceWithId) {
	o.Approver = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *NonEmployeeApprovalItemBase) SetAccountName(v string) {
	o.AccountName = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetApprovalStatus() ApprovalStatus {
	if o == nil || isNil(o.ApprovalStatus) {
		var ret ApprovalStatus
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetApprovalStatusOk() (*ApprovalStatus, bool) {
	if o == nil || isNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasApprovalStatus() bool {
	if o != nil && !isNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given ApprovalStatus and assigns it to the ApprovalStatus field.
func (o *NonEmployeeApprovalItemBase) SetApprovalStatus(v ApprovalStatus) {
	o.ApprovalStatus = &v
}

// GetApprovalOrder returns the ApprovalOrder field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetApprovalOrder() float32 {
	if o == nil || isNil(o.ApprovalOrder) {
		var ret float32
		return ret
	}
	return *o.ApprovalOrder
}

// GetApprovalOrderOk returns a tuple with the ApprovalOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetApprovalOrderOk() (*float32, bool) {
	if o == nil || isNil(o.ApprovalOrder) {
		return nil, false
	}
	return o.ApprovalOrder, true
}

// HasApprovalOrder returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasApprovalOrder() bool {
	if o != nil && !isNil(o.ApprovalOrder) {
		return true
	}

	return false
}

// SetApprovalOrder gets a reference to the given float32 and assigns it to the ApprovalOrder field.
func (o *NonEmployeeApprovalItemBase) SetApprovalOrder(v float32) {
	o.ApprovalOrder = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NonEmployeeApprovalItemBase) SetComment(v string) {
	o.Comment = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeApprovalItemBase) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemBase) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemBase) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemBase) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeApprovalItemBase) SetCreated(v time.Time) {
	o.Created = &v
}

func (o NonEmployeeApprovalItemBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeApprovalItemBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Approver) {
		toSerialize["approver"] = o.Approver
	}
	if !isNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !isNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if !isNil(o.ApprovalOrder) {
		toSerialize["approvalOrder"] = o.ApprovalOrder
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeApprovalItemBase) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeApprovalItemBase := _NonEmployeeApprovalItemBase{}

	if err = json.Unmarshal(bytes, &varNonEmployeeApprovalItemBase); err == nil {
	*o = NonEmployeeApprovalItemBase(varNonEmployeeApprovalItemBase)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "approver")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "approvalStatus")
		delete(additionalProperties, "approvalOrder")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeApprovalItemBase struct {
	value *NonEmployeeApprovalItemBase
	isSet bool
}

func (v NullableNonEmployeeApprovalItemBase) Get() *NonEmployeeApprovalItemBase {
	return v.value
}

func (v *NullableNonEmployeeApprovalItemBase) Set(val *NonEmployeeApprovalItemBase) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeApprovalItemBase) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeApprovalItemBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeApprovalItemBase(val *NonEmployeeApprovalItemBase) *NullableNonEmployeeApprovalItemBase {
	return &NullableNonEmployeeApprovalItemBase{value: val, isSet: true}
}

func (v NullableNonEmployeeApprovalItemBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeApprovalItemBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


