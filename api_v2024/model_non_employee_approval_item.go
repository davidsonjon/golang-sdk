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

// checks if the NonEmployeeApprovalItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeApprovalItem{}

// NonEmployeeApprovalItem struct for NonEmployeeApprovalItem
type NonEmployeeApprovalItem struct {
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
	NonEmployeeRequest *NonEmployeeRequestLite `json:"nonEmployeeRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeApprovalItem NonEmployeeApprovalItem

// NewNonEmployeeApprovalItem instantiates a new NonEmployeeApprovalItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeApprovalItem() *NonEmployeeApprovalItem {
	this := NonEmployeeApprovalItem{}
	return &this
}

// NewNonEmployeeApprovalItemWithDefaults instantiates a new NonEmployeeApprovalItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeApprovalItemWithDefaults() *NonEmployeeApprovalItem {
	this := NonEmployeeApprovalItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeApprovalItem) SetId(v string) {
	o.Id = &v
}

// GetApprover returns the Approver field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetApprover() NonEmployeeIdentityReferenceWithId {
	if o == nil || IsNil(o.Approver) {
		var ret NonEmployeeIdentityReferenceWithId
		return ret
	}
	return *o.Approver
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetApproverOk() (*NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || IsNil(o.Approver) {
		return nil, false
	}
	return o.Approver, true
}

// HasApprover returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasApprover() bool {
	if o != nil && !IsNil(o.Approver) {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NonEmployeeIdentityReferenceWithId and assigns it to the Approver field.
func (o *NonEmployeeApprovalItem) SetApprover(v NonEmployeeIdentityReferenceWithId) {
	o.Approver = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *NonEmployeeApprovalItem) SetAccountName(v string) {
	o.AccountName = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetApprovalStatus() ApprovalStatus {
	if o == nil || IsNil(o.ApprovalStatus) {
		var ret ApprovalStatus
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetApprovalStatusOk() (*ApprovalStatus, bool) {
	if o == nil || IsNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasApprovalStatus() bool {
	if o != nil && !IsNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given ApprovalStatus and assigns it to the ApprovalStatus field.
func (o *NonEmployeeApprovalItem) SetApprovalStatus(v ApprovalStatus) {
	o.ApprovalStatus = &v
}

// GetApprovalOrder returns the ApprovalOrder field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetApprovalOrder() float32 {
	if o == nil || IsNil(o.ApprovalOrder) {
		var ret float32
		return ret
	}
	return *o.ApprovalOrder
}

// GetApprovalOrderOk returns a tuple with the ApprovalOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetApprovalOrderOk() (*float32, bool) {
	if o == nil || IsNil(o.ApprovalOrder) {
		return nil, false
	}
	return o.ApprovalOrder, true
}

// HasApprovalOrder returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasApprovalOrder() bool {
	if o != nil && !IsNil(o.ApprovalOrder) {
		return true
	}

	return false
}

// SetApprovalOrder gets a reference to the given float32 and assigns it to the ApprovalOrder field.
func (o *NonEmployeeApprovalItem) SetApprovalOrder(v float32) {
	o.ApprovalOrder = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NonEmployeeApprovalItem) SetComment(v string) {
	o.Comment = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeApprovalItem) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeApprovalItem) SetCreated(v time.Time) {
	o.Created = &v
}

// GetNonEmployeeRequest returns the NonEmployeeRequest field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItem) GetNonEmployeeRequest() NonEmployeeRequestLite {
	if o == nil || IsNil(o.NonEmployeeRequest) {
		var ret NonEmployeeRequestLite
		return ret
	}
	return *o.NonEmployeeRequest
}

// GetNonEmployeeRequestOk returns a tuple with the NonEmployeeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItem) GetNonEmployeeRequestOk() (*NonEmployeeRequestLite, bool) {
	if o == nil || IsNil(o.NonEmployeeRequest) {
		return nil, false
	}
	return o.NonEmployeeRequest, true
}

// HasNonEmployeeRequest returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItem) HasNonEmployeeRequest() bool {
	if o != nil && !IsNil(o.NonEmployeeRequest) {
		return true
	}

	return false
}

// SetNonEmployeeRequest gets a reference to the given NonEmployeeRequestLite and assigns it to the NonEmployeeRequest field.
func (o *NonEmployeeApprovalItem) SetNonEmployeeRequest(v NonEmployeeRequestLite) {
	o.NonEmployeeRequest = &v
}

func (o NonEmployeeApprovalItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeApprovalItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Approver) {
		toSerialize["approver"] = o.Approver
	}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if !IsNil(o.ApprovalOrder) {
		toSerialize["approvalOrder"] = o.ApprovalOrder
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.NonEmployeeRequest) {
		toSerialize["nonEmployeeRequest"] = o.NonEmployeeRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeApprovalItem) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeApprovalItem := _NonEmployeeApprovalItem{}

	err = json.Unmarshal(data, &varNonEmployeeApprovalItem)

	if err != nil {
		return err
	}

	*o = NonEmployeeApprovalItem(varNonEmployeeApprovalItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "approver")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "approvalStatus")
		delete(additionalProperties, "approvalOrder")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		delete(additionalProperties, "nonEmployeeRequest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeApprovalItem struct {
	value *NonEmployeeApprovalItem
	isSet bool
}

func (v NullableNonEmployeeApprovalItem) Get() *NonEmployeeApprovalItem {
	return v.value
}

func (v *NullableNonEmployeeApprovalItem) Set(val *NonEmployeeApprovalItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeApprovalItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeApprovalItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeApprovalItem(val *NonEmployeeApprovalItem) *NullableNonEmployeeApprovalItem {
	return &NullableNonEmployeeApprovalItem{value: val, isSet: true}
}

func (v NullableNonEmployeeApprovalItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeApprovalItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


