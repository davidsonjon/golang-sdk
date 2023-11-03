/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the IdentityPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityPreviewResponse{}

// IdentityPreviewResponse struct for IdentityPreviewResponse
type IdentityPreviewResponse struct {
	Identity *IdentityPreviewResponseIdentity `json:"identity,omitempty"`
	PreviewAttributes []IdentityAttributePreview `json:"previewAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityPreviewResponse IdentityPreviewResponse

// NewIdentityPreviewResponse instantiates a new IdentityPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPreviewResponse() *IdentityPreviewResponse {
	this := IdentityPreviewResponse{}
	return &this
}

// NewIdentityPreviewResponseWithDefaults instantiates a new IdentityPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPreviewResponseWithDefaults() *IdentityPreviewResponse {
	this := IdentityPreviewResponse{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *IdentityPreviewResponse) GetIdentity() IdentityPreviewResponseIdentity {
	if o == nil || isNil(o.Identity) {
		var ret IdentityPreviewResponseIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewResponse) GetIdentityOk() (*IdentityPreviewResponseIdentity, bool) {
	if o == nil || isNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *IdentityPreviewResponse) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given IdentityPreviewResponseIdentity and assigns it to the Identity field.
func (o *IdentityPreviewResponse) SetIdentity(v IdentityPreviewResponseIdentity) {
	o.Identity = &v
}

// GetPreviewAttributes returns the PreviewAttributes field value if set, zero value otherwise.
func (o *IdentityPreviewResponse) GetPreviewAttributes() []IdentityAttributePreview {
	if o == nil || isNil(o.PreviewAttributes) {
		var ret []IdentityAttributePreview
		return ret
	}
	return o.PreviewAttributes
}

// GetPreviewAttributesOk returns a tuple with the PreviewAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewResponse) GetPreviewAttributesOk() ([]IdentityAttributePreview, bool) {
	if o == nil || isNil(o.PreviewAttributes) {
		return nil, false
	}
	return o.PreviewAttributes, true
}

// HasPreviewAttributes returns a boolean if a field has been set.
func (o *IdentityPreviewResponse) HasPreviewAttributes() bool {
	if o != nil && !isNil(o.PreviewAttributes) {
		return true
	}

	return false
}

// SetPreviewAttributes gets a reference to the given []IdentityAttributePreview and assigns it to the PreviewAttributes field.
func (o *IdentityPreviewResponse) SetPreviewAttributes(v []IdentityAttributePreview) {
	o.PreviewAttributes = v
}

func (o IdentityPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.PreviewAttributes) {
		toSerialize["previewAttributes"] = o.PreviewAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityPreviewResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityPreviewResponse := _IdentityPreviewResponse{}

	if err = json.Unmarshal(bytes, &varIdentityPreviewResponse); err == nil {
		*o = IdentityPreviewResponse(varIdentityPreviewResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "previewAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityPreviewResponse struct {
	value *IdentityPreviewResponse
	isSet bool
}

func (v NullableIdentityPreviewResponse) Get() *IdentityPreviewResponse {
	return v.value
}

func (v *NullableIdentityPreviewResponse) Set(val *IdentityPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPreviewResponse(val *IdentityPreviewResponse) *NullableIdentityPreviewResponse {
	return &NullableIdentityPreviewResponse{value: val, isSet: true}
}

func (v NullableIdentityPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


