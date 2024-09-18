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

// checks if the BackupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupResponse{}

// BackupResponse struct for BackupResponse
type BackupResponse struct {
	// Unique id assigned to this backup.
	JobId *string `json:"jobId,omitempty"`
	// Status of the backup.
	Status *string `json:"status,omitempty"`
	// Type of the job, will always be BACKUP for this type of job.
	Type *string `json:"type,omitempty"`
	// The name of the tenant performing the upload
	Tenant *string `json:"tenant,omitempty"`
	// The name of the requester.
	RequesterName *string `json:"requesterName,omitempty"`
	// Whether or not a file was created and stored for this backup.
	FileExists *bool `json:"fileExists,omitempty"`
	// The time the job was started.
	Created *time.Time `json:"created,omitempty"`
	// The time of the last update to the job.
	Modified *time.Time `json:"modified,omitempty"`
	// The time the job was completed.
	Completed *time.Time `json:"completed,omitempty"`
	// The name assigned to the upload file in the request body.
	Name *string `json:"name,omitempty"`
	// Whether this backup can be deleted by a regular user.
	UserCanDelete *bool `json:"userCanDelete,omitempty"`
	// Whether this backup contains all supported object types or only some of them.
	IsPartial *bool `json:"isPartial,omitempty"`
	// Denotes how this backup was created. - MANUAL - The backup was created by a user. - AUTOMATED - The backup was created by devops. - AUTOMATED_DRAFT - The backup was created during a draft process. - UPLOADED - The backup was created by uploading an existing configuration file.
	BackupType *string `json:"backupType,omitempty"`
	Options NullableBackupOptions `json:"options,omitempty"`
	// Whether the object details of this backup are ready.
	HydrationStatus *string `json:"hydrationStatus,omitempty"`
	// Number of objects contained in this backup.
	TotalObjectCount *int64 `json:"totalObjectCount,omitempty"`
	// Whether this backup has been transferred to a customer storage location.
	CloudStorageStatus *string `json:"cloudStorageStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BackupResponse BackupResponse

// NewBackupResponse instantiates a new BackupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupResponse() *BackupResponse {
	this := BackupResponse{}
	var fileExists bool = true
	this.FileExists = &fileExists
	var userCanDelete bool = true
	this.UserCanDelete = &userCanDelete
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// NewBackupResponseWithDefaults instantiates a new BackupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupResponseWithDefaults() *BackupResponse {
	this := BackupResponse{}
	var fileExists bool = true
	this.FileExists = &fileExists
	var userCanDelete bool = true
	this.UserCanDelete = &userCanDelete
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *BackupResponse) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *BackupResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *BackupResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BackupResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BackupResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BackupResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BackupResponse) SetType(v string) {
	o.Type = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *BackupResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *BackupResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *BackupResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetRequesterName returns the RequesterName field value if set, zero value otherwise.
func (o *BackupResponse) GetRequesterName() string {
	if o == nil || IsNil(o.RequesterName) {
		var ret string
		return ret
	}
	return *o.RequesterName
}

// GetRequesterNameOk returns a tuple with the RequesterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetRequesterNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterName) {
		return nil, false
	}
	return o.RequesterName, true
}

// HasRequesterName returns a boolean if a field has been set.
func (o *BackupResponse) HasRequesterName() bool {
	if o != nil && !IsNil(o.RequesterName) {
		return true
	}

	return false
}

// SetRequesterName gets a reference to the given string and assigns it to the RequesterName field.
func (o *BackupResponse) SetRequesterName(v string) {
	o.RequesterName = &v
}

// GetFileExists returns the FileExists field value if set, zero value otherwise.
func (o *BackupResponse) GetFileExists() bool {
	if o == nil || IsNil(o.FileExists) {
		var ret bool
		return ret
	}
	return *o.FileExists
}

// GetFileExistsOk returns a tuple with the FileExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetFileExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.FileExists) {
		return nil, false
	}
	return o.FileExists, true
}

// HasFileExists returns a boolean if a field has been set.
func (o *BackupResponse) HasFileExists() bool {
	if o != nil && !IsNil(o.FileExists) {
		return true
	}

	return false
}

// SetFileExists gets a reference to the given bool and assigns it to the FileExists field.
func (o *BackupResponse) SetFileExists(v bool) {
	o.FileExists = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *BackupResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *BackupResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *BackupResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *BackupResponse) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *BackupResponse) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *BackupResponse) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *BackupResponse) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *BackupResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *BackupResponse) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupResponse) SetName(v string) {
	o.Name = &v
}

// GetUserCanDelete returns the UserCanDelete field value if set, zero value otherwise.
func (o *BackupResponse) GetUserCanDelete() bool {
	if o == nil || IsNil(o.UserCanDelete) {
		var ret bool
		return ret
	}
	return *o.UserCanDelete
}

// GetUserCanDeleteOk returns a tuple with the UserCanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetUserCanDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.UserCanDelete) {
		return nil, false
	}
	return o.UserCanDelete, true
}

// HasUserCanDelete returns a boolean if a field has been set.
func (o *BackupResponse) HasUserCanDelete() bool {
	if o != nil && !IsNil(o.UserCanDelete) {
		return true
	}

	return false
}

// SetUserCanDelete gets a reference to the given bool and assigns it to the UserCanDelete field.
func (o *BackupResponse) SetUserCanDelete(v bool) {
	o.UserCanDelete = &v
}

// GetIsPartial returns the IsPartial field value if set, zero value otherwise.
func (o *BackupResponse) GetIsPartial() bool {
	if o == nil || IsNil(o.IsPartial) {
		var ret bool
		return ret
	}
	return *o.IsPartial
}

// GetIsPartialOk returns a tuple with the IsPartial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetIsPartialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartial) {
		return nil, false
	}
	return o.IsPartial, true
}

// HasIsPartial returns a boolean if a field has been set.
func (o *BackupResponse) HasIsPartial() bool {
	if o != nil && !IsNil(o.IsPartial) {
		return true
	}

	return false
}

// SetIsPartial gets a reference to the given bool and assigns it to the IsPartial field.
func (o *BackupResponse) SetIsPartial(v bool) {
	o.IsPartial = &v
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *BackupResponse) GetBackupType() string {
	if o == nil || IsNil(o.BackupType) {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetBackupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BackupType) {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *BackupResponse) HasBackupType() bool {
	if o != nil && !IsNil(o.BackupType) {
		return true
	}

	return false
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *BackupResponse) SetBackupType(v string) {
	o.BackupType = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupResponse) GetOptions() BackupOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret BackupOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupResponse) GetOptionsOk() (*BackupOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *BackupResponse) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableBackupOptions and assigns it to the Options field.
func (o *BackupResponse) SetOptions(v BackupOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *BackupResponse) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *BackupResponse) UnsetOptions() {
	o.Options.Unset()
}

// GetHydrationStatus returns the HydrationStatus field value if set, zero value otherwise.
func (o *BackupResponse) GetHydrationStatus() string {
	if o == nil || IsNil(o.HydrationStatus) {
		var ret string
		return ret
	}
	return *o.HydrationStatus
}

// GetHydrationStatusOk returns a tuple with the HydrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetHydrationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HydrationStatus) {
		return nil, false
	}
	return o.HydrationStatus, true
}

// HasHydrationStatus returns a boolean if a field has been set.
func (o *BackupResponse) HasHydrationStatus() bool {
	if o != nil && !IsNil(o.HydrationStatus) {
		return true
	}

	return false
}

// SetHydrationStatus gets a reference to the given string and assigns it to the HydrationStatus field.
func (o *BackupResponse) SetHydrationStatus(v string) {
	o.HydrationStatus = &v
}

// GetTotalObjectCount returns the TotalObjectCount field value if set, zero value otherwise.
func (o *BackupResponse) GetTotalObjectCount() int64 {
	if o == nil || IsNil(o.TotalObjectCount) {
		var ret int64
		return ret
	}
	return *o.TotalObjectCount
}

// GetTotalObjectCountOk returns a tuple with the TotalObjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetTotalObjectCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalObjectCount) {
		return nil, false
	}
	return o.TotalObjectCount, true
}

// HasTotalObjectCount returns a boolean if a field has been set.
func (o *BackupResponse) HasTotalObjectCount() bool {
	if o != nil && !IsNil(o.TotalObjectCount) {
		return true
	}

	return false
}

// SetTotalObjectCount gets a reference to the given int64 and assigns it to the TotalObjectCount field.
func (o *BackupResponse) SetTotalObjectCount(v int64) {
	o.TotalObjectCount = &v
}

// GetCloudStorageStatus returns the CloudStorageStatus field value if set, zero value otherwise.
func (o *BackupResponse) GetCloudStorageStatus() string {
	if o == nil || IsNil(o.CloudStorageStatus) {
		var ret string
		return ret
	}
	return *o.CloudStorageStatus
}

// GetCloudStorageStatusOk returns a tuple with the CloudStorageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetCloudStorageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CloudStorageStatus) {
		return nil, false
	}
	return o.CloudStorageStatus, true
}

// HasCloudStorageStatus returns a boolean if a field has been set.
func (o *BackupResponse) HasCloudStorageStatus() bool {
	if o != nil && !IsNil(o.CloudStorageStatus) {
		return true
	}

	return false
}

// SetCloudStorageStatus gets a reference to the given string and assigns it to the CloudStorageStatus field.
func (o *BackupResponse) SetCloudStorageStatus(v string) {
	o.CloudStorageStatus = &v
}

func (o BackupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.RequesterName) {
		toSerialize["requesterName"] = o.RequesterName
	}
	if !IsNil(o.FileExists) {
		toSerialize["fileExists"] = o.FileExists
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
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
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if !IsNil(o.HydrationStatus) {
		toSerialize["hydrationStatus"] = o.HydrationStatus
	}
	if !IsNil(o.TotalObjectCount) {
		toSerialize["totalObjectCount"] = o.TotalObjectCount
	}
	if !IsNil(o.CloudStorageStatus) {
		toSerialize["cloudStorageStatus"] = o.CloudStorageStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BackupResponse) UnmarshalJSON(data []byte) (err error) {
	varBackupResponse := _BackupResponse{}

	err = json.Unmarshal(data, &varBackupResponse)

	if err != nil {
		return err
	}

	*o = BackupResponse(varBackupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "requesterName")
		delete(additionalProperties, "fileExists")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "name")
		delete(additionalProperties, "userCanDelete")
		delete(additionalProperties, "isPartial")
		delete(additionalProperties, "backupType")
		delete(additionalProperties, "options")
		delete(additionalProperties, "hydrationStatus")
		delete(additionalProperties, "totalObjectCount")
		delete(additionalProperties, "cloudStorageStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBackupResponse struct {
	value *BackupResponse
	isSet bool
}

func (v NullableBackupResponse) Get() *BackupResponse {
	return v.value
}

func (v *NullableBackupResponse) Set(val *BackupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupResponse(val *BackupResponse) *NullableBackupResponse {
	return &NullableBackupResponse{value: val, isSet: true}
}

func (v NullableBackupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


