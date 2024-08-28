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
	"fmt"
)

// checks if the UploadsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadsResponse{}

// UploadsResponse struct for UploadsResponse
type UploadsResponse struct {
	// Unique id assigned to this job.
	JobId string `json:"jobId"`
	// Status of the job.
	Status string `json:"status"`
	// Type of the job, either Backup or Draft.
	Type string `json:"type"`
	// The name of the tenant performing the upload
	Tenant *string `json:"tenant,omitempty"`
	// The name of the requester.
	RequesterName *string `json:"requesterName,omitempty"`
	// The time the job was started.
	Created time.Time `json:"created"`
	// The time of the last update to the job.
	Modified time.Time `json:"modified"`
	// The name assigned to the upload file in the request body.
	Name *string `json:"name,omitempty"`
	// Is the job a regular backup job, if so is the user allowed to delete the backup file. Since this is an upload job it remains as false.
	UserCanDelete *bool `json:"userCanDelete,omitempty"`
	// Is the job a regular backup job, if so is it partial. Since this is an upload job it remains as false.
	IsPartial *bool `json:"isPartial,omitempty"`
	// What kind of backup is this being treated as.
	BackupType *string `json:"backupType,omitempty"`
	// have the objects contained in the upload file been hydrated.
	HydrationStatus *string `json:"hydrationStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UploadsResponse UploadsResponse

// NewUploadsResponse instantiates a new UploadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadsResponse(jobId string, status string, type_ string, created time.Time, modified time.Time) *UploadsResponse {
	this := UploadsResponse{}
	this.JobId = jobId
	this.Status = status
	this.Type = type_
	this.Created = created
	this.Modified = modified
	var userCanDelete bool = true
	this.UserCanDelete = &userCanDelete
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// NewUploadsResponseWithDefaults instantiates a new UploadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadsResponseWithDefaults() *UploadsResponse {
	this := UploadsResponse{}
	var userCanDelete bool = true
	this.UserCanDelete = &userCanDelete
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// GetJobId returns the JobId field value
func (o *UploadsResponse) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *UploadsResponse) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *UploadsResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UploadsResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *UploadsResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UploadsResponse) SetType(v string) {
	o.Type = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *UploadsResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *UploadsResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *UploadsResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetRequesterName returns the RequesterName field value if set, zero value otherwise.
func (o *UploadsResponse) GetRequesterName() string {
	if o == nil || IsNil(o.RequesterName) {
		var ret string
		return ret
	}
	return *o.RequesterName
}

// GetRequesterNameOk returns a tuple with the RequesterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetRequesterNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterName) {
		return nil, false
	}
	return o.RequesterName, true
}

// HasRequesterName returns a boolean if a field has been set.
func (o *UploadsResponse) HasRequesterName() bool {
	if o != nil && !IsNil(o.RequesterName) {
		return true
	}

	return false
}

// SetRequesterName gets a reference to the given string and assigns it to the RequesterName field.
func (o *UploadsResponse) SetRequesterName(v string) {
	o.RequesterName = &v
}

// GetCreated returns the Created field value
func (o *UploadsResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *UploadsResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *UploadsResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *UploadsResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UploadsResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UploadsResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UploadsResponse) SetName(v string) {
	o.Name = &v
}

// GetUserCanDelete returns the UserCanDelete field value if set, zero value otherwise.
func (o *UploadsResponse) GetUserCanDelete() bool {
	if o == nil || IsNil(o.UserCanDelete) {
		var ret bool
		return ret
	}
	return *o.UserCanDelete
}

// GetUserCanDeleteOk returns a tuple with the UserCanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetUserCanDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.UserCanDelete) {
		return nil, false
	}
	return o.UserCanDelete, true
}

// HasUserCanDelete returns a boolean if a field has been set.
func (o *UploadsResponse) HasUserCanDelete() bool {
	if o != nil && !IsNil(o.UserCanDelete) {
		return true
	}

	return false
}

// SetUserCanDelete gets a reference to the given bool and assigns it to the UserCanDelete field.
func (o *UploadsResponse) SetUserCanDelete(v bool) {
	o.UserCanDelete = &v
}

// GetIsPartial returns the IsPartial field value if set, zero value otherwise.
func (o *UploadsResponse) GetIsPartial() bool {
	if o == nil || IsNil(o.IsPartial) {
		var ret bool
		return ret
	}
	return *o.IsPartial
}

// GetIsPartialOk returns a tuple with the IsPartial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetIsPartialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartial) {
		return nil, false
	}
	return o.IsPartial, true
}

// HasIsPartial returns a boolean if a field has been set.
func (o *UploadsResponse) HasIsPartial() bool {
	if o != nil && !IsNil(o.IsPartial) {
		return true
	}

	return false
}

// SetIsPartial gets a reference to the given bool and assigns it to the IsPartial field.
func (o *UploadsResponse) SetIsPartial(v bool) {
	o.IsPartial = &v
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *UploadsResponse) GetBackupType() string {
	if o == nil || IsNil(o.BackupType) {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetBackupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BackupType) {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *UploadsResponse) HasBackupType() bool {
	if o != nil && !IsNil(o.BackupType) {
		return true
	}

	return false
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *UploadsResponse) SetBackupType(v string) {
	o.BackupType = &v
}

// GetHydrationStatus returns the HydrationStatus field value if set, zero value otherwise.
func (o *UploadsResponse) GetHydrationStatus() string {
	if o == nil || IsNil(o.HydrationStatus) {
		var ret string
		return ret
	}
	return *o.HydrationStatus
}

// GetHydrationStatusOk returns a tuple with the HydrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadsResponse) GetHydrationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HydrationStatus) {
		return nil, false
	}
	return o.HydrationStatus, true
}

// HasHydrationStatus returns a boolean if a field has been set.
func (o *UploadsResponse) HasHydrationStatus() bool {
	if o != nil && !IsNil(o.HydrationStatus) {
		return true
	}

	return false
}

// SetHydrationStatus gets a reference to the given string and assigns it to the HydrationStatus field.
func (o *UploadsResponse) SetHydrationStatus(v string) {
	o.HydrationStatus = &v
}

func (o UploadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.RequesterName) {
		toSerialize["requesterName"] = o.RequesterName
	}
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserCanDelete) {
		toSerialize["userCanDelete"] = o.UserCanDelete
	}
	if !IsNil(o.IsPartial) {
		toSerialize["isPartial"] = o.IsPartial
	}
	if !IsNil(o.BackupType) {
		toSerialize["backupType"] = o.BackupType
	}
	if !IsNil(o.HydrationStatus) {
		toSerialize["hydrationStatus"] = o.HydrationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UploadsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobId",
		"status",
		"type",
		"created",
		"modified",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUploadsResponse := _UploadsResponse{}

	err = json.Unmarshal(data, &varUploadsResponse)

	if err != nil {
		return err
	}

	*o = UploadsResponse(varUploadsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "requesterName")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "name")
		delete(additionalProperties, "userCanDelete")
		delete(additionalProperties, "isPartial")
		delete(additionalProperties, "backupType")
		delete(additionalProperties, "hydrationStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadsResponse struct {
	value *UploadsResponse
	isSet bool
}

func (v NullableUploadsResponse) Get() *UploadsResponse {
	return v.value
}

func (v *NullableUploadsResponse) Set(val *UploadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadsResponse(val *UploadsResponse) *NullableUploadsResponse {
	return &NullableUploadsResponse{value: val, isSet: true}
}

func (v NullableUploadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


