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

// checks if the CommonAccessItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAccessItemResponse{}

// CommonAccessItemResponse struct for CommonAccessItemResponse
type CommonAccessItemResponse struct {
	// Common Access Item ID
	Id *string `json:"id,omitempty"`
	Access *CommonAccessItemAccess `json:"access,omitempty"`
	Status *CommonAccessItemState `json:"status,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	ReviewedByUser *bool `json:"reviewedByUser,omitempty"`
	LastReviewed *string `json:"lastReviewed,omitempty"`
	CreatedByUser *string `json:"createdByUser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonAccessItemResponse CommonAccessItemResponse

// NewCommonAccessItemResponse instantiates a new CommonAccessItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAccessItemResponse() *CommonAccessItemResponse {
	this := CommonAccessItemResponse{}
	return &this
}

// NewCommonAccessItemResponseWithDefaults instantiates a new CommonAccessItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAccessItemResponseWithDefaults() *CommonAccessItemResponse {
	this := CommonAccessItemResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommonAccessItemResponse) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetAccess() CommonAccessItemAccess {
	if o == nil || isNil(o.Access) {
		var ret CommonAccessItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetAccessOk() (*CommonAccessItemAccess, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given CommonAccessItemAccess and assigns it to the Access field.
func (o *CommonAccessItemResponse) SetAccess(v CommonAccessItemAccess) {
	o.Access = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetStatus() CommonAccessItemState {
	if o == nil || isNil(o.Status) {
		var ret CommonAccessItemState
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetStatusOk() (*CommonAccessItemState, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CommonAccessItemState and assigns it to the Status field.
func (o *CommonAccessItemResponse) SetStatus(v CommonAccessItemState) {
	o.Status = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetLastUpdated() string {
	if o == nil || isNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *CommonAccessItemResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetReviewedByUser returns the ReviewedByUser field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetReviewedByUser() bool {
	if o == nil || isNil(o.ReviewedByUser) {
		var ret bool
		return ret
	}
	return *o.ReviewedByUser
}

// GetReviewedByUserOk returns a tuple with the ReviewedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetReviewedByUserOk() (*bool, bool) {
	if o == nil || isNil(o.ReviewedByUser) {
		return nil, false
	}
	return o.ReviewedByUser, true
}

// HasReviewedByUser returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasReviewedByUser() bool {
	if o != nil && !isNil(o.ReviewedByUser) {
		return true
	}

	return false
}

// SetReviewedByUser gets a reference to the given bool and assigns it to the ReviewedByUser field.
func (o *CommonAccessItemResponse) SetReviewedByUser(v bool) {
	o.ReviewedByUser = &v
}

// GetLastReviewed returns the LastReviewed field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetLastReviewed() string {
	if o == nil || isNil(o.LastReviewed) {
		var ret string
		return ret
	}
	return *o.LastReviewed
}

// GetLastReviewedOk returns a tuple with the LastReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetLastReviewedOk() (*string, bool) {
	if o == nil || isNil(o.LastReviewed) {
		return nil, false
	}
	return o.LastReviewed, true
}

// HasLastReviewed returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasLastReviewed() bool {
	if o != nil && !isNil(o.LastReviewed) {
		return true
	}

	return false
}

// SetLastReviewed gets a reference to the given string and assigns it to the LastReviewed field.
func (o *CommonAccessItemResponse) SetLastReviewed(v string) {
	o.LastReviewed = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *CommonAccessItemResponse) GetCreatedByUser() string {
	if o == nil || isNil(o.CreatedByUser) {
		var ret string
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessItemResponse) GetCreatedByUserOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByUser) {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *CommonAccessItemResponse) HasCreatedByUser() bool {
	if o != nil && !isNil(o.CreatedByUser) {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given string and assigns it to the CreatedByUser field.
func (o *CommonAccessItemResponse) SetCreatedByUser(v string) {
	o.CreatedByUser = &v
}

func (o CommonAccessItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAccessItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.ReviewedByUser) {
		toSerialize["reviewedByUser"] = o.ReviewedByUser
	}
	if !isNil(o.LastReviewed) {
		toSerialize["lastReviewed"] = o.LastReviewed
	}
	if !isNil(o.CreatedByUser) {
		toSerialize["createdByUser"] = o.CreatedByUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonAccessItemResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCommonAccessItemResponse := _CommonAccessItemResponse{}

	if err = json.Unmarshal(bytes, &varCommonAccessItemResponse); err == nil {
	*o = CommonAccessItemResponse(varCommonAccessItemResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "access")
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "reviewedByUser")
		delete(additionalProperties, "lastReviewed")
		delete(additionalProperties, "createdByUser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAccessItemResponse struct {
	value *CommonAccessItemResponse
	isSet bool
}

func (v NullableCommonAccessItemResponse) Get() *CommonAccessItemResponse {
	return v.value
}

func (v *NullableCommonAccessItemResponse) Set(val *CommonAccessItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAccessItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAccessItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAccessItemResponse(val *CommonAccessItemResponse) *NullableCommonAccessItemResponse {
	return &NullableCommonAccessItemResponse{value: val, isSet: true}
}

func (v NullableCommonAccessItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAccessItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


