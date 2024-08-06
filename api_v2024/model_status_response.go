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

// checks if the StatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponse{}

// StatusResponse Response model for connection check, configuration test and ping of source connectors.
type StatusResponse struct {
	// ID of the source
	Id *string `json:"id,omitempty"`
	// Name of the source
	Name *string `json:"name,omitempty"`
	// The status of the health check.
	Status *string `json:"status,omitempty"`
	// The number of milliseconds spent on the entire request.
	ElapsedMillis *int32 `json:"elapsedMillis,omitempty"`
	// The document contains the results of the health check. The schema of this document depends on the type of source used. 
	Details map[string]interface{} `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StatusResponse StatusResponse

// NewStatusResponse instantiates a new StatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponse() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// NewStatusResponseWithDefaults instantiates a new StatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseWithDefaults() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StatusResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StatusResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StatusResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatusResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatusResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StatusResponse) SetStatus(v string) {
	o.Status = &v
}

// GetElapsedMillis returns the ElapsedMillis field value if set, zero value otherwise.
func (o *StatusResponse) GetElapsedMillis() int32 {
	if o == nil || isNil(o.ElapsedMillis) {
		var ret int32
		return ret
	}
	return *o.ElapsedMillis
}

// GetElapsedMillisOk returns a tuple with the ElapsedMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetElapsedMillisOk() (*int32, bool) {
	if o == nil || isNil(o.ElapsedMillis) {
		return nil, false
	}
	return o.ElapsedMillis, true
}

// HasElapsedMillis returns a boolean if a field has been set.
func (o *StatusResponse) HasElapsedMillis() bool {
	if o != nil && !isNil(o.ElapsedMillis) {
		return true
	}

	return false
}

// SetElapsedMillis gets a reference to the given int32 and assigns it to the ElapsedMillis field.
func (o *StatusResponse) SetElapsedMillis(v int32) {
	o.ElapsedMillis = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *StatusResponse) GetDetails() map[string]interface{} {
	if o == nil || isNil(o.Details) {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Details) {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *StatusResponse) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *StatusResponse) SetDetails(v map[string]interface{}) {
	o.Details = v
}

func (o StatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: name is readOnly
	// skip: status is readOnly
	// skip: elapsedMillis is readOnly
	// skip: details is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StatusResponse) UnmarshalJSON(bytes []byte) (err error) {
	varStatusResponse := _StatusResponse{}

	if err = json.Unmarshal(bytes, &varStatusResponse); err == nil {
	*o = StatusResponse(varStatusResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "elapsedMillis")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatusResponse struct {
	value *StatusResponse
	isSet bool
}

func (v NullableStatusResponse) Get() *StatusResponse {
	return v.value
}

func (v *NullableStatusResponse) Set(val *StatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponse(val *StatusResponse) *NullableStatusResponse {
	return &NullableStatusResponse{value: val, isSet: true}
}

func (v NullableStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


