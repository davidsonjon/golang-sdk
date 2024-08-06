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

// checks if the NonEmployeeRequestWithoutApprovalItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeRequestWithoutApprovalItem{}

// NonEmployeeRequestWithoutApprovalItem struct for NonEmployeeRequestWithoutApprovalItem
type NonEmployeeRequestWithoutApprovalItem struct {
	// Non-Employee request id.
	Id *string `json:"id,omitempty"`
	Requester *NonEmployeeIdentityReferenceWithId `json:"requester,omitempty"`
	// Requested identity account name.
	AccountName *string `json:"accountName,omitempty"`
	// Non-Employee's first name.
	FirstName *string `json:"firstName,omitempty"`
	// Non-Employee's last name.
	LastName *string `json:"lastName,omitempty"`
	// Non-Employee's email.
	Email *string `json:"email,omitempty"`
	// Non-Employee's phone.
	Phone *string `json:"phone,omitempty"`
	// The account ID of a valid identity to serve as this non-employee's manager.
	Manager *string `json:"manager,omitempty"`
	NonEmployeeSource *NonEmployeeSourceLiteWithSchemaAttributes `json:"nonEmployeeSource,omitempty"`
	// Attribute blob/bag for a non-employee.
	Data *map[string]string `json:"data,omitempty"`
	ApprovalStatus *ApprovalStatus `json:"approvalStatus,omitempty"`
	// comment of requester
	Comment *string `json:"comment,omitempty"`
	// When the request was completely approved.
	CompletionDate *time.Time `json:"completionDate,omitempty"`
	// Non-Employee employment start date.
	StartDate *string `json:"startDate,omitempty"`
	// Non-Employee employment end date.
	EndDate *string `json:"endDate,omitempty"`
	// When the request was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeRequestWithoutApprovalItem NonEmployeeRequestWithoutApprovalItem

// NewNonEmployeeRequestWithoutApprovalItem instantiates a new NonEmployeeRequestWithoutApprovalItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeRequestWithoutApprovalItem() *NonEmployeeRequestWithoutApprovalItem {
	this := NonEmployeeRequestWithoutApprovalItem{}
	return &this
}

// NewNonEmployeeRequestWithoutApprovalItemWithDefaults instantiates a new NonEmployeeRequestWithoutApprovalItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeRequestWithoutApprovalItemWithDefaults() *NonEmployeeRequestWithoutApprovalItem {
	this := NonEmployeeRequestWithoutApprovalItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetId(v string) {
	o.Id = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetRequester() NonEmployeeIdentityReferenceWithId {
	if o == nil || isNil(o.Requester) {
		var ret NonEmployeeIdentityReferenceWithId
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetRequesterOk() (*NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || isNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasRequester() bool {
	if o != nil && !isNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given NonEmployeeIdentityReferenceWithId and assigns it to the Requester field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetRequester(v NonEmployeeIdentityReferenceWithId) {
	o.Requester = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetAccountName(v string) {
	o.AccountName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetPhone(v string) {
	o.Phone = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetManager() string {
	if o == nil || isNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetManagerOk() (*string, bool) {
	if o == nil || isNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasManager() bool {
	if o != nil && !isNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetManager(v string) {
	o.Manager = &v
}

// GetNonEmployeeSource returns the NonEmployeeSource field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetNonEmployeeSource() NonEmployeeSourceLiteWithSchemaAttributes {
	if o == nil || isNil(o.NonEmployeeSource) {
		var ret NonEmployeeSourceLiteWithSchemaAttributes
		return ret
	}
	return *o.NonEmployeeSource
}

// GetNonEmployeeSourceOk returns a tuple with the NonEmployeeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetNonEmployeeSourceOk() (*NonEmployeeSourceLiteWithSchemaAttributes, bool) {
	if o == nil || isNil(o.NonEmployeeSource) {
		return nil, false
	}
	return o.NonEmployeeSource, true
}

// HasNonEmployeeSource returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasNonEmployeeSource() bool {
	if o != nil && !isNil(o.NonEmployeeSource) {
		return true
	}

	return false
}

// SetNonEmployeeSource gets a reference to the given NonEmployeeSourceLiteWithSchemaAttributes and assigns it to the NonEmployeeSource field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetNonEmployeeSource(v NonEmployeeSourceLiteWithSchemaAttributes) {
	o.NonEmployeeSource = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetData() map[string]string {
	if o == nil || isNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetDataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetData(v map[string]string) {
	o.Data = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetApprovalStatus() ApprovalStatus {
	if o == nil || isNil(o.ApprovalStatus) {
		var ret ApprovalStatus
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetApprovalStatusOk() (*ApprovalStatus, bool) {
	if o == nil || isNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasApprovalStatus() bool {
	if o != nil && !isNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given ApprovalStatus and assigns it to the ApprovalStatus field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetApprovalStatus(v ApprovalStatus) {
	o.ApprovalStatus = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetComment(v string) {
	o.Comment = &v
}

// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetCompletionDate() time.Time {
	if o == nil || isNil(o.CompletionDate) {
		var ret time.Time
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetCompletionDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletionDate) {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasCompletionDate() bool {
	if o != nil && !isNil(o.CompletionDate) {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetCompletionDate(v time.Time) {
	o.CompletionDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetEndDate(v string) {
	o.EndDate = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeRequestWithoutApprovalItem) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeRequestWithoutApprovalItem) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeRequestWithoutApprovalItem) SetCreated(v time.Time) {
	o.Created = &v
}

func (o NonEmployeeRequestWithoutApprovalItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeRequestWithoutApprovalItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !isNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !isNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !isNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !isNil(o.NonEmployeeSource) {
		toSerialize["nonEmployeeSource"] = o.NonEmployeeSource
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.CompletionDate) {
		toSerialize["completionDate"] = o.CompletionDate
	}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
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

func (o *NonEmployeeRequestWithoutApprovalItem) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeRequestWithoutApprovalItem := _NonEmployeeRequestWithoutApprovalItem{}

	if err = json.Unmarshal(bytes, &varNonEmployeeRequestWithoutApprovalItem); err == nil {
	*o = NonEmployeeRequestWithoutApprovalItem(varNonEmployeeRequestWithoutApprovalItem)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "nonEmployeeSource")
		delete(additionalProperties, "data")
		delete(additionalProperties, "approvalStatus")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "completionDate")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeRequestWithoutApprovalItem struct {
	value *NonEmployeeRequestWithoutApprovalItem
	isSet bool
}

func (v NullableNonEmployeeRequestWithoutApprovalItem) Get() *NonEmployeeRequestWithoutApprovalItem {
	return v.value
}

func (v *NullableNonEmployeeRequestWithoutApprovalItem) Set(val *NonEmployeeRequestWithoutApprovalItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeRequestWithoutApprovalItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeRequestWithoutApprovalItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeRequestWithoutApprovalItem(val *NonEmployeeRequestWithoutApprovalItem) *NullableNonEmployeeRequestWithoutApprovalItem {
	return &NullableNonEmployeeRequestWithoutApprovalItem{value: val, isSet: true}
}

func (v NullableNonEmployeeRequestWithoutApprovalItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeRequestWithoutApprovalItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


