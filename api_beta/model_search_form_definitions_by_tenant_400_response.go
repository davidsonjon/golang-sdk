/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the SearchFormDefinitionsByTenant400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchFormDefinitionsByTenant400Response{}

// SearchFormDefinitionsByTenant400Response struct for SearchFormDefinitionsByTenant400Response
type SearchFormDefinitionsByTenant400Response struct {
	DetailCode *string `json:"detailCode,omitempty"`
	Messages []ErrorMessage `json:"messages,omitempty"`
	StatusCode *int64 `json:"statusCode,omitempty"`
	TrackingId *string `json:"trackingId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchFormDefinitionsByTenant400Response SearchFormDefinitionsByTenant400Response

// NewSearchFormDefinitionsByTenant400Response instantiates a new SearchFormDefinitionsByTenant400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchFormDefinitionsByTenant400Response() *SearchFormDefinitionsByTenant400Response {
	this := SearchFormDefinitionsByTenant400Response{}
	return &this
}

// NewSearchFormDefinitionsByTenant400ResponseWithDefaults instantiates a new SearchFormDefinitionsByTenant400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchFormDefinitionsByTenant400ResponseWithDefaults() *SearchFormDefinitionsByTenant400Response {
	this := SearchFormDefinitionsByTenant400Response{}
	return &this
}

// GetDetailCode returns the DetailCode field value if set, zero value otherwise.
func (o *SearchFormDefinitionsByTenant400Response) GetDetailCode() string {
	if o == nil || isNil(o.DetailCode) {
		var ret string
		return ret
	}
	return *o.DetailCode
}

// GetDetailCodeOk returns a tuple with the DetailCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchFormDefinitionsByTenant400Response) GetDetailCodeOk() (*string, bool) {
	if o == nil || isNil(o.DetailCode) {
		return nil, false
	}
	return o.DetailCode, true
}

// HasDetailCode returns a boolean if a field has been set.
func (o *SearchFormDefinitionsByTenant400Response) HasDetailCode() bool {
	if o != nil && !isNil(o.DetailCode) {
		return true
	}

	return false
}

// SetDetailCode gets a reference to the given string and assigns it to the DetailCode field.
func (o *SearchFormDefinitionsByTenant400Response) SetDetailCode(v string) {
	o.DetailCode = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *SearchFormDefinitionsByTenant400Response) GetMessages() []ErrorMessage {
	if o == nil || isNil(o.Messages) {
		var ret []ErrorMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchFormDefinitionsByTenant400Response) GetMessagesOk() ([]ErrorMessage, bool) {
	if o == nil || isNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *SearchFormDefinitionsByTenant400Response) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ErrorMessage and assigns it to the Messages field.
func (o *SearchFormDefinitionsByTenant400Response) SetMessages(v []ErrorMessage) {
	o.Messages = v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SearchFormDefinitionsByTenant400Response) GetStatusCode() int64 {
	if o == nil || isNil(o.StatusCode) {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchFormDefinitionsByTenant400Response) GetStatusCodeOk() (*int64, bool) {
	if o == nil || isNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SearchFormDefinitionsByTenant400Response) HasStatusCode() bool {
	if o != nil && !isNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *SearchFormDefinitionsByTenant400Response) SetStatusCode(v int64) {
	o.StatusCode = &v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *SearchFormDefinitionsByTenant400Response) GetTrackingId() string {
	if o == nil || isNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchFormDefinitionsByTenant400Response) GetTrackingIdOk() (*string, bool) {
	if o == nil || isNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *SearchFormDefinitionsByTenant400Response) HasTrackingId() bool {
	if o != nil && !isNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *SearchFormDefinitionsByTenant400Response) SetTrackingId(v string) {
	o.TrackingId = &v
}

func (o SearchFormDefinitionsByTenant400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchFormDefinitionsByTenant400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DetailCode) {
		toSerialize["detailCode"] = o.DetailCode
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !isNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchFormDefinitionsByTenant400Response) UnmarshalJSON(bytes []byte) (err error) {
	varSearchFormDefinitionsByTenant400Response := _SearchFormDefinitionsByTenant400Response{}

	if err = json.Unmarshal(bytes, &varSearchFormDefinitionsByTenant400Response); err == nil {
	*o = SearchFormDefinitionsByTenant400Response(varSearchFormDefinitionsByTenant400Response)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "detailCode")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "statusCode")
		delete(additionalProperties, "trackingId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchFormDefinitionsByTenant400Response struct {
	value *SearchFormDefinitionsByTenant400Response
	isSet bool
}

func (v NullableSearchFormDefinitionsByTenant400Response) Get() *SearchFormDefinitionsByTenant400Response {
	return v.value
}

func (v *NullableSearchFormDefinitionsByTenant400Response) Set(val *SearchFormDefinitionsByTenant400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchFormDefinitionsByTenant400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchFormDefinitionsByTenant400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchFormDefinitionsByTenant400Response(val *SearchFormDefinitionsByTenant400Response) *NullableSearchFormDefinitionsByTenant400Response {
	return &NullableSearchFormDefinitionsByTenant400Response{value: val, isSet: true}
}

func (v NullableSearchFormDefinitionsByTenant400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchFormDefinitionsByTenant400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

