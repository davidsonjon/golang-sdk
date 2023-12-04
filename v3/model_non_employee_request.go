/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the NonEmployeeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeRequest{}

// NonEmployeeRequest struct for NonEmployeeRequest
type NonEmployeeRequest struct {
	// Non-Employee source id.
	Id *string `json:"id,omitempty"`
	// Source Id associated with this non-employee source.
	SourceId *string `json:"sourceId,omitempty"`
	// Source name associated with this non-employee source.
	Name *string `json:"name,omitempty"`
	// Source description associated with this non-employee source.
	Description *string `json:"description,omitempty"`
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
	NonEmployeeSource *NonEmployeeSourceLite `json:"nonEmployeeSource,omitempty"`
	// Attribute blob/bag for a non-employee.
	Data *map[string]string `json:"data,omitempty"`
	// List of approval item for the request
	ApprovalItems []NonEmployeeApprovalItemBase `json:"approvalItems,omitempty"`
	ApprovalStatus *ApprovalStatus `json:"approvalStatus,omitempty"`
	// comment of requester
	Comment *string `json:"comment,omitempty"`
	// When the request was completely approved.
	CompletionDate *time.Time `json:"completionDate,omitempty"`
	// Non-Employee employment start date.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Non-Employee employment end date.
	EndDate *time.Time `json:"endDate,omitempty"`
	// When the request was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeRequest NonEmployeeRequest

// NewNonEmployeeRequest instantiates a new NonEmployeeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeRequest() *NonEmployeeRequest {
	this := NonEmployeeRequest{}
	return &this
}

// NewNonEmployeeRequestWithDefaults instantiates a new NonEmployeeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeRequestWithDefaults() *NonEmployeeRequest {
	this := NonEmployeeRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeRequest) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeRequest) SetSourceId(v string) {
	o.SourceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NonEmployeeRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonEmployeeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *NonEmployeeRequest) SetAccountName(v string) {
	o.AccountName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *NonEmployeeRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *NonEmployeeRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NonEmployeeRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *NonEmployeeRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetManager() string {
	if o == nil || isNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetManagerOk() (*string, bool) {
	if o == nil || isNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasManager() bool {
	if o != nil && !isNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *NonEmployeeRequest) SetManager(v string) {
	o.Manager = &v
}

// GetNonEmployeeSource returns the NonEmployeeSource field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetNonEmployeeSource() NonEmployeeSourceLite {
	if o == nil || isNil(o.NonEmployeeSource) {
		var ret NonEmployeeSourceLite
		return ret
	}
	return *o.NonEmployeeSource
}

// GetNonEmployeeSourceOk returns a tuple with the NonEmployeeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetNonEmployeeSourceOk() (*NonEmployeeSourceLite, bool) {
	if o == nil || isNil(o.NonEmployeeSource) {
		return nil, false
	}
	return o.NonEmployeeSource, true
}

// HasNonEmployeeSource returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasNonEmployeeSource() bool {
	if o != nil && !isNil(o.NonEmployeeSource) {
		return true
	}

	return false
}

// SetNonEmployeeSource gets a reference to the given NonEmployeeSourceLite and assigns it to the NonEmployeeSource field.
func (o *NonEmployeeRequest) SetNonEmployeeSource(v NonEmployeeSourceLite) {
	o.NonEmployeeSource = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetData() map[string]string {
	if o == nil || isNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetDataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *NonEmployeeRequest) SetData(v map[string]string) {
	o.Data = &v
}

// GetApprovalItems returns the ApprovalItems field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetApprovalItems() []NonEmployeeApprovalItemBase {
	if o == nil || isNil(o.ApprovalItems) {
		var ret []NonEmployeeApprovalItemBase
		return ret
	}
	return o.ApprovalItems
}

// GetApprovalItemsOk returns a tuple with the ApprovalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetApprovalItemsOk() ([]NonEmployeeApprovalItemBase, bool) {
	if o == nil || isNil(o.ApprovalItems) {
		return nil, false
	}
	return o.ApprovalItems, true
}

// HasApprovalItems returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasApprovalItems() bool {
	if o != nil && !isNil(o.ApprovalItems) {
		return true
	}

	return false
}

// SetApprovalItems gets a reference to the given []NonEmployeeApprovalItemBase and assigns it to the ApprovalItems field.
func (o *NonEmployeeRequest) SetApprovalItems(v []NonEmployeeApprovalItemBase) {
	o.ApprovalItems = v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetApprovalStatus() ApprovalStatus {
	if o == nil || isNil(o.ApprovalStatus) {
		var ret ApprovalStatus
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetApprovalStatusOk() (*ApprovalStatus, bool) {
	if o == nil || isNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasApprovalStatus() bool {
	if o != nil && !isNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given ApprovalStatus and assigns it to the ApprovalStatus field.
func (o *NonEmployeeRequest) SetApprovalStatus(v ApprovalStatus) {
	o.ApprovalStatus = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NonEmployeeRequest) SetComment(v string) {
	o.Comment = &v
}

// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetCompletionDate() time.Time {
	if o == nil || isNil(o.CompletionDate) {
		var ret time.Time
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetCompletionDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletionDate) {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasCompletionDate() bool {
	if o != nil && !isNil(o.CompletionDate) {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.
func (o *NonEmployeeRequest) SetCompletionDate(v time.Time) {
	o.CompletionDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NonEmployeeRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *NonEmployeeRequest) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeRequest) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeRequest) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeRequest) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeRequest) SetCreated(v time.Time) {
	o.Created = &v
}

func (o NonEmployeeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !isNil(o.ApprovalItems) {
		toSerialize["approvalItems"] = o.ApprovalItems
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

func (o *NonEmployeeRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeRequest := _NonEmployeeRequest{}

	if err = json.Unmarshal(bytes, &varNonEmployeeRequest); err == nil {
	*o = NonEmployeeRequest(varNonEmployeeRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "nonEmployeeSource")
		delete(additionalProperties, "data")
		delete(additionalProperties, "approvalItems")
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

type NullableNonEmployeeRequest struct {
	value *NonEmployeeRequest
	isSet bool
}

func (v NullableNonEmployeeRequest) Get() *NonEmployeeRequest {
	return v.value
}

func (v *NullableNonEmployeeRequest) Set(val *NonEmployeeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeRequest(val *NonEmployeeRequest) *NullableNonEmployeeRequest {
	return &NullableNonEmployeeRequest{value: val, isSet: true}
}

func (v NullableNonEmployeeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


