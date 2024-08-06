/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the AccessRequestItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestItemResponse{}

// AccessRequestItemResponse struct for AccessRequestItemResponse
type AccessRequestItemResponse struct {
	// the access request item operation
	Operation *string `json:"operation,omitempty"`
	// the access item type
	AccessItemType *string `json:"accessItemType,omitempty"`
	// the name of access request item
	Name *string `json:"name,omitempty"`
	// the final decision for the access request
	Decision *string `json:"decision,omitempty"`
	// the description of access request item
	Description *string `json:"description,omitempty"`
	// the source id
	SourceId *string `json:"sourceId,omitempty"`
	// the source Name
	SourceName *string `json:"sourceName,omitempty"`
	ApprovalInfos []ApprovalInfoResponse `json:"approvalInfos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestItemResponse AccessRequestItemResponse

// NewAccessRequestItemResponse instantiates a new AccessRequestItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestItemResponse() *AccessRequestItemResponse {
	this := AccessRequestItemResponse{}
	return &this
}

// NewAccessRequestItemResponseWithDefaults instantiates a new AccessRequestItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestItemResponseWithDefaults() *AccessRequestItemResponse {
	this := AccessRequestItemResponse{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *AccessRequestItemResponse) SetOperation(v string) {
	o.Operation = &v
}

// GetAccessItemType returns the AccessItemType field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetAccessItemType() string {
	if o == nil || isNil(o.AccessItemType) {
		var ret string
		return ret
	}
	return *o.AccessItemType
}

// GetAccessItemTypeOk returns a tuple with the AccessItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetAccessItemTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessItemType) {
		return nil, false
	}
	return o.AccessItemType, true
}

// HasAccessItemType returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasAccessItemType() bool {
	if o != nil && !isNil(o.AccessItemType) {
		return true
	}

	return false
}

// SetAccessItemType gets a reference to the given string and assigns it to the AccessItemType field.
func (o *AccessRequestItemResponse) SetAccessItemType(v string) {
	o.AccessItemType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessRequestItemResponse) SetName(v string) {
	o.Name = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetDecision() string {
	if o == nil || isNil(o.Decision) {
		var ret string
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetDecisionOk() (*string, bool) {
	if o == nil || isNil(o.Decision) {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasDecision() bool {
	if o != nil && !isNil(o.Decision) {
		return true
	}

	return false
}

// SetDecision gets a reference to the given string and assigns it to the Decision field.
func (o *AccessRequestItemResponse) SetDecision(v string) {
	o.Decision = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessRequestItemResponse) SetDescription(v string) {
	o.Description = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AccessRequestItemResponse) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessRequestItemResponse) SetSourceName(v string) {
	o.SourceName = &v
}

// GetApprovalInfos returns the ApprovalInfos field value if set, zero value otherwise.
func (o *AccessRequestItemResponse) GetApprovalInfos() []ApprovalInfoResponse {
	if o == nil || isNil(o.ApprovalInfos) {
		var ret []ApprovalInfoResponse
		return ret
	}
	return o.ApprovalInfos
}

// GetApprovalInfosOk returns a tuple with the ApprovalInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItemResponse) GetApprovalInfosOk() ([]ApprovalInfoResponse, bool) {
	if o == nil || isNil(o.ApprovalInfos) {
		return nil, false
	}
	return o.ApprovalInfos, true
}

// HasApprovalInfos returns a boolean if a field has been set.
func (o *AccessRequestItemResponse) HasApprovalInfos() bool {
	if o != nil && !isNil(o.ApprovalInfos) {
		return true
	}

	return false
}

// SetApprovalInfos gets a reference to the given []ApprovalInfoResponse and assigns it to the ApprovalInfos field.
func (o *AccessRequestItemResponse) SetApprovalInfos(v []ApprovalInfoResponse) {
	o.ApprovalInfos = v
}

func (o AccessRequestItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.AccessItemType) {
		toSerialize["accessItemType"] = o.AccessItemType
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Decision) {
		toSerialize["decision"] = o.Decision
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.ApprovalInfos) {
		toSerialize["approvalInfos"] = o.ApprovalInfos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestItemResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccessRequestItemResponse := _AccessRequestItemResponse{}

	if err = json.Unmarshal(bytes, &varAccessRequestItemResponse); err == nil {
	*o = AccessRequestItemResponse(varAccessRequestItemResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "accessItemType")
		delete(additionalProperties, "name")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "description")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "approvalInfos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestItemResponse struct {
	value *AccessRequestItemResponse
	isSet bool
}

func (v NullableAccessRequestItemResponse) Get() *AccessRequestItemResponse {
	return v.value
}

func (v *NullableAccessRequestItemResponse) Set(val *AccessRequestItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestItemResponse(val *AccessRequestItemResponse) *NullableAccessRequestItemResponse {
	return &NullableAccessRequestItemResponse{value: val, isSet: true}
}

func (v NullableAccessRequestItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


