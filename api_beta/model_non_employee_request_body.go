/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the NonEmployeeRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeRequestBody{}

// NonEmployeeRequestBody struct for NonEmployeeRequestBody
type NonEmployeeRequestBody struct {
	// Requested identity account name.
	AccountName string `json:"accountName"`
	// Non-Employee's first name.
	FirstName string `json:"firstName"`
	// Non-Employee's last name.
	LastName string `json:"lastName"`
	// Non-Employee's email.
	Email string `json:"email"`
	// Non-Employee's phone.
	Phone string `json:"phone"`
	// The account ID of a valid identity to serve as this non-employee's manager.
	Manager string `json:"manager"`
	// Non-Employee's source id.
	SourceId string `json:"sourceId"`
	// Attribute blob/bag for a non-employee, 10 attributes is the maximum size supported.
	Data *map[string]string `json:"data,omitempty"`
	// Non-Employee employment start date.
	StartDate time.Time `json:"startDate"`
	// Non-Employee employment end date.
	EndDate time.Time `json:"endDate"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeRequestBody NonEmployeeRequestBody

// NewNonEmployeeRequestBody instantiates a new NonEmployeeRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeRequestBody(accountName string, firstName string, lastName string, email string, phone string, manager string, sourceId string, startDate time.Time, endDate time.Time) *NonEmployeeRequestBody {
	this := NonEmployeeRequestBody{}
	this.AccountName = accountName
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Phone = phone
	this.Manager = manager
	this.SourceId = sourceId
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewNonEmployeeRequestBodyWithDefaults instantiates a new NonEmployeeRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeRequestBodyWithDefaults() *NonEmployeeRequestBody {
	this := NonEmployeeRequestBody{}
	return &this
}

// GetAccountName returns the AccountName field value
func (o *NonEmployeeRequestBody) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *NonEmployeeRequestBody) SetAccountName(v string) {
	o.AccountName = v
}

// GetFirstName returns the FirstName field value
func (o *NonEmployeeRequestBody) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *NonEmployeeRequestBody) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *NonEmployeeRequestBody) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *NonEmployeeRequestBody) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *NonEmployeeRequestBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *NonEmployeeRequestBody) SetEmail(v string) {
	o.Email = v
}

// GetPhone returns the Phone field value
func (o *NonEmployeeRequestBody) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *NonEmployeeRequestBody) SetPhone(v string) {
	o.Phone = v
}

// GetManager returns the Manager field value
func (o *NonEmployeeRequestBody) GetManager() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Manager
}

// GetManagerOk returns a tuple with the Manager field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetManagerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manager, true
}

// SetManager sets field value
func (o *NonEmployeeRequestBody) SetManager(v string) {
	o.Manager = v
}

// GetSourceId returns the SourceId field value
func (o *NonEmployeeRequestBody) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *NonEmployeeRequestBody) SetSourceId(v string) {
	o.SourceId = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NonEmployeeRequestBody) GetData() map[string]string {
	if o == nil || isNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetDataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NonEmployeeRequestBody) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *NonEmployeeRequestBody) SetData(v map[string]string) {
	o.Data = &v
}

// GetStartDate returns the StartDate field value
func (o *NonEmployeeRequestBody) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *NonEmployeeRequestBody) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *NonEmployeeRequestBody) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestBody) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *NonEmployeeRequestBody) SetEndDate(v time.Time) {
	o.EndDate = v
}

func (o NonEmployeeRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountName"] = o.AccountName
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["email"] = o.Email
	toSerialize["phone"] = o.Phone
	toSerialize["manager"] = o.Manager
	toSerialize["sourceId"] = o.SourceId
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeRequestBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountName",
		"firstName",
		"lastName",
		"email",
		"phone",
		"manager",
		"sourceId",
		"startDate",
		"endDate",
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

	varNonEmployeeRequestBody := _NonEmployeeRequestBody{}

	if err = json.Unmarshal(bytes, &varNonEmployeeRequestBody); err == nil {
	*o = NonEmployeeRequestBody(varNonEmployeeRequestBody)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "data")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeRequestBody struct {
	value *NonEmployeeRequestBody
	isSet bool
}

func (v NullableNonEmployeeRequestBody) Get() *NonEmployeeRequestBody {
	return v.value
}

func (v *NullableNonEmployeeRequestBody) Set(val *NonEmployeeRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeRequestBody(val *NonEmployeeRequestBody) *NullableNonEmployeeRequestBody {
	return &NullableNonEmployeeRequestBody{value: val, isSet: true}
}

func (v NullableNonEmployeeRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


