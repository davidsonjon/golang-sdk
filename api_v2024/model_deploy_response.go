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

// checks if the DeployResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployResponse{}

// DeployResponse struct for DeployResponse
type DeployResponse struct {
	// Unique id assigned to this job.
	JobId *string `json:"jobId,omitempty"`
	// Status of the job.
	Status *string `json:"status,omitempty"`
	// Type of the job, will always be CONFIG_DEPLOY_DRAFT for this type of job.
	Type *string `json:"type,omitempty"`
	// Message providing information about the outcome of the deploy process.
	Message *string `json:"message,omitempty"`
	// The name of the user that initiated the deploy process.
	RequesterName *string `json:"requesterName,omitempty"`
	// Whether or not a results file was created and stored for this deploy.
	FileExists *bool `json:"fileExists,omitempty"`
	// The time the job was started.
	Created *time.Time `json:"created,omitempty"`
	// The time of the last update to the job.
	Modified *time.Time `json:"modified,omitempty"`
	// The time the job was completed.
	Completed *time.Time `json:"completed,omitempty"`
	// The id of the draft that was used for this deploy.
	DraftId *string `json:"draftId,omitempty"`
	// The name of the draft that was used for this deploy.
	DraftName *string `json:"draftName,omitempty"`
	// Whether this deploy results file has been transferred to a customer storage location.
	CloudStorageStatus *string `json:"cloudStorageStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeployResponse DeployResponse

// NewDeployResponse instantiates a new DeployResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployResponse() *DeployResponse {
	this := DeployResponse{}
	var fileExists bool = true
	this.FileExists = &fileExists
	return &this
}

// NewDeployResponseWithDefaults instantiates a new DeployResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployResponseWithDefaults() *DeployResponse {
	this := DeployResponse{}
	var fileExists bool = true
	this.FileExists = &fileExists
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *DeployResponse) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *DeployResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *DeployResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeployResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeployResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeployResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeployResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeployResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeployResponse) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeployResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeployResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeployResponse) SetMessage(v string) {
	o.Message = &v
}

// GetRequesterName returns the RequesterName field value if set, zero value otherwise.
func (o *DeployResponse) GetRequesterName() string {
	if o == nil || IsNil(o.RequesterName) {
		var ret string
		return ret
	}
	return *o.RequesterName
}

// GetRequesterNameOk returns a tuple with the RequesterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetRequesterNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterName) {
		return nil, false
	}
	return o.RequesterName, true
}

// HasRequesterName returns a boolean if a field has been set.
func (o *DeployResponse) HasRequesterName() bool {
	if o != nil && !IsNil(o.RequesterName) {
		return true
	}

	return false
}

// SetRequesterName gets a reference to the given string and assigns it to the RequesterName field.
func (o *DeployResponse) SetRequesterName(v string) {
	o.RequesterName = &v
}

// GetFileExists returns the FileExists field value if set, zero value otherwise.
func (o *DeployResponse) GetFileExists() bool {
	if o == nil || IsNil(o.FileExists) {
		var ret bool
		return ret
	}
	return *o.FileExists
}

// GetFileExistsOk returns a tuple with the FileExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetFileExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.FileExists) {
		return nil, false
	}
	return o.FileExists, true
}

// HasFileExists returns a boolean if a field has been set.
func (o *DeployResponse) HasFileExists() bool {
	if o != nil && !IsNil(o.FileExists) {
		return true
	}

	return false
}

// SetFileExists gets a reference to the given bool and assigns it to the FileExists field.
func (o *DeployResponse) SetFileExists(v bool) {
	o.FileExists = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DeployResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DeployResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DeployResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *DeployResponse) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *DeployResponse) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *DeployResponse) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *DeployResponse) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *DeployResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *DeployResponse) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetDraftId returns the DraftId field value if set, zero value otherwise.
func (o *DeployResponse) GetDraftId() string {
	if o == nil || IsNil(o.DraftId) {
		var ret string
		return ret
	}
	return *o.DraftId
}

// GetDraftIdOk returns a tuple with the DraftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetDraftIdOk() (*string, bool) {
	if o == nil || IsNil(o.DraftId) {
		return nil, false
	}
	return o.DraftId, true
}

// HasDraftId returns a boolean if a field has been set.
func (o *DeployResponse) HasDraftId() bool {
	if o != nil && !IsNil(o.DraftId) {
		return true
	}

	return false
}

// SetDraftId gets a reference to the given string and assigns it to the DraftId field.
func (o *DeployResponse) SetDraftId(v string) {
	o.DraftId = &v
}

// GetDraftName returns the DraftName field value if set, zero value otherwise.
func (o *DeployResponse) GetDraftName() string {
	if o == nil || IsNil(o.DraftName) {
		var ret string
		return ret
	}
	return *o.DraftName
}

// GetDraftNameOk returns a tuple with the DraftName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetDraftNameOk() (*string, bool) {
	if o == nil || IsNil(o.DraftName) {
		return nil, false
	}
	return o.DraftName, true
}

// HasDraftName returns a boolean if a field has been set.
func (o *DeployResponse) HasDraftName() bool {
	if o != nil && !IsNil(o.DraftName) {
		return true
	}

	return false
}

// SetDraftName gets a reference to the given string and assigns it to the DraftName field.
func (o *DeployResponse) SetDraftName(v string) {
	o.DraftName = &v
}

// GetCloudStorageStatus returns the CloudStorageStatus field value if set, zero value otherwise.
func (o *DeployResponse) GetCloudStorageStatus() string {
	if o == nil || IsNil(o.CloudStorageStatus) {
		var ret string
		return ret
	}
	return *o.CloudStorageStatus
}

// GetCloudStorageStatusOk returns a tuple with the CloudStorageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployResponse) GetCloudStorageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CloudStorageStatus) {
		return nil, false
	}
	return o.CloudStorageStatus, true
}

// HasCloudStorageStatus returns a boolean if a field has been set.
func (o *DeployResponse) HasCloudStorageStatus() bool {
	if o != nil && !IsNil(o.CloudStorageStatus) {
		return true
	}

	return false
}

// SetCloudStorageStatus gets a reference to the given string and assigns it to the CloudStorageStatus field.
func (o *DeployResponse) SetCloudStorageStatus(v string) {
	o.CloudStorageStatus = &v
}

func (o DeployResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
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
	if !IsNil(o.DraftId) {
		toSerialize["draftId"] = o.DraftId
	}
	if !IsNil(o.DraftName) {
		toSerialize["draftName"] = o.DraftName
	}
	if !IsNil(o.CloudStorageStatus) {
		toSerialize["cloudStorageStatus"] = o.CloudStorageStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeployResponse) UnmarshalJSON(data []byte) (err error) {
	varDeployResponse := _DeployResponse{}

	err = json.Unmarshal(data, &varDeployResponse)

	if err != nil {
		return err
	}

	*o = DeployResponse(varDeployResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "message")
		delete(additionalProperties, "requesterName")
		delete(additionalProperties, "fileExists")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "draftId")
		delete(additionalProperties, "draftName")
		delete(additionalProperties, "cloudStorageStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeployResponse struct {
	value *DeployResponse
	isSet bool
}

func (v NullableDeployResponse) Get() *DeployResponse {
	return v.value
}

func (v *NullableDeployResponse) Set(val *DeployResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployResponse(val *DeployResponse) *NullableDeployResponse {
	return &NullableDeployResponse{value: val, isSet: true}
}

func (v NullableDeployResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


