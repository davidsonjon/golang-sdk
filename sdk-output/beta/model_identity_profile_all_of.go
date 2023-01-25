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

// IdentityProfileAllOf struct for IdentityProfileAllOf
type IdentityProfileAllOf struct {
	// The description of the Identity Profile.
	Description *string `json:"description,omitempty"`
	Owner NullableIdentityProfileAllOfOwner `json:"owner,omitempty"`
	// The priority for an Identity Profile.
	Priority *int64 `json:"priority,omitempty"`
	AuthoritativeSource IdentityProfileAllOfAuthoritativeSource `json:"authoritativeSource"`
	// True if a identity refresh is needed. Typically triggered when a change on the source has been made
	IdentityRefreshRequired *bool `json:"identityRefreshRequired,omitempty"`
	// The number of identities that belong to the Identity Profile.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	IdentityAttributeConfig *IdentityAttributeConfig `json:"identityAttributeConfig,omitempty"`
	IdentityExceptionReportReference *IdentityExceptionReportReference `json:"identityExceptionReportReference,omitempty"`
	// Indicates the value of requiresPeriodicRefresh attribute for the Identity Profile.
	HasTimeBasedAttr *bool `json:"hasTimeBasedAttr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProfileAllOf IdentityProfileAllOf

// NewIdentityProfileAllOf instantiates a new IdentityProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProfileAllOf(authoritativeSource IdentityProfileAllOfAuthoritativeSource) *IdentityProfileAllOf {
	this := IdentityProfileAllOf{}
	this.AuthoritativeSource = authoritativeSource
	return &this
}

// NewIdentityProfileAllOfWithDefaults instantiates a new IdentityProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProfileAllOfWithDefaults() *IdentityProfileAllOf {
	this := IdentityProfileAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityProfileAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProfileAllOf) GetOwner() IdentityProfileAllOfOwner {
	if o == nil || isNil(o.Owner.Get()) {
		var ret IdentityProfileAllOfOwner
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProfileAllOf) GetOwnerOk() (*IdentityProfileAllOfOwner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableIdentityProfileAllOfOwner and assigns it to the Owner field.
func (o *IdentityProfileAllOf) SetOwner(v IdentityProfileAllOfOwner) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *IdentityProfileAllOf) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *IdentityProfileAllOf) UnsetOwner() {
	o.Owner.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetPriority() int64 {
	if o == nil || isNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetPriorityOk() (*int64, bool) {
	if o == nil || isNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *IdentityProfileAllOf) SetPriority(v int64) {
	o.Priority = &v
}

// GetAuthoritativeSource returns the AuthoritativeSource field value
func (o *IdentityProfileAllOf) GetAuthoritativeSource() IdentityProfileAllOfAuthoritativeSource {
	if o == nil {
		var ret IdentityProfileAllOfAuthoritativeSource
		return ret
	}

	return o.AuthoritativeSource
}

// GetAuthoritativeSourceOk returns a tuple with the AuthoritativeSource field value
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetAuthoritativeSourceOk() (*IdentityProfileAllOfAuthoritativeSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthoritativeSource, true
}

// SetAuthoritativeSource sets field value
func (o *IdentityProfileAllOf) SetAuthoritativeSource(v IdentityProfileAllOfAuthoritativeSource) {
	o.AuthoritativeSource = v
}

// GetIdentityRefreshRequired returns the IdentityRefreshRequired field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetIdentityRefreshRequired() bool {
	if o == nil || isNil(o.IdentityRefreshRequired) {
		var ret bool
		return ret
	}
	return *o.IdentityRefreshRequired
}

// GetIdentityRefreshRequiredOk returns a tuple with the IdentityRefreshRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetIdentityRefreshRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IdentityRefreshRequired) {
		return nil, false
	}
	return o.IdentityRefreshRequired, true
}

// HasIdentityRefreshRequired returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasIdentityRefreshRequired() bool {
	if o != nil && !isNil(o.IdentityRefreshRequired) {
		return true
	}

	return false
}

// SetIdentityRefreshRequired gets a reference to the given bool and assigns it to the IdentityRefreshRequired field.
func (o *IdentityProfileAllOf) SetIdentityRefreshRequired(v bool) {
	o.IdentityRefreshRequired = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *IdentityProfileAllOf) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetIdentityAttributeConfig returns the IdentityAttributeConfig field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetIdentityAttributeConfig() IdentityAttributeConfig {
	if o == nil || isNil(o.IdentityAttributeConfig) {
		var ret IdentityAttributeConfig
		return ret
	}
	return *o.IdentityAttributeConfig
}

// GetIdentityAttributeConfigOk returns a tuple with the IdentityAttributeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetIdentityAttributeConfigOk() (*IdentityAttributeConfig, bool) {
	if o == nil || isNil(o.IdentityAttributeConfig) {
		return nil, false
	}
	return o.IdentityAttributeConfig, true
}

// HasIdentityAttributeConfig returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasIdentityAttributeConfig() bool {
	if o != nil && !isNil(o.IdentityAttributeConfig) {
		return true
	}

	return false
}

// SetIdentityAttributeConfig gets a reference to the given IdentityAttributeConfig and assigns it to the IdentityAttributeConfig field.
func (o *IdentityProfileAllOf) SetIdentityAttributeConfig(v IdentityAttributeConfig) {
	o.IdentityAttributeConfig = &v
}

// GetIdentityExceptionReportReference returns the IdentityExceptionReportReference field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetIdentityExceptionReportReference() IdentityExceptionReportReference {
	if o == nil || isNil(o.IdentityExceptionReportReference) {
		var ret IdentityExceptionReportReference
		return ret
	}
	return *o.IdentityExceptionReportReference
}

// GetIdentityExceptionReportReferenceOk returns a tuple with the IdentityExceptionReportReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetIdentityExceptionReportReferenceOk() (*IdentityExceptionReportReference, bool) {
	if o == nil || isNil(o.IdentityExceptionReportReference) {
		return nil, false
	}
	return o.IdentityExceptionReportReference, true
}

// HasIdentityExceptionReportReference returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasIdentityExceptionReportReference() bool {
	if o != nil && !isNil(o.IdentityExceptionReportReference) {
		return true
	}

	return false
}

// SetIdentityExceptionReportReference gets a reference to the given IdentityExceptionReportReference and assigns it to the IdentityExceptionReportReference field.
func (o *IdentityProfileAllOf) SetIdentityExceptionReportReference(v IdentityExceptionReportReference) {
	o.IdentityExceptionReportReference = &v
}

// GetHasTimeBasedAttr returns the HasTimeBasedAttr field value if set, zero value otherwise.
func (o *IdentityProfileAllOf) GetHasTimeBasedAttr() bool {
	if o == nil || isNil(o.HasTimeBasedAttr) {
		var ret bool
		return ret
	}
	return *o.HasTimeBasedAttr
}

// GetHasTimeBasedAttrOk returns a tuple with the HasTimeBasedAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOf) GetHasTimeBasedAttrOk() (*bool, bool) {
	if o == nil || isNil(o.HasTimeBasedAttr) {
		return nil, false
	}
	return o.HasTimeBasedAttr, true
}

// HasHasTimeBasedAttr returns a boolean if a field has been set.
func (o *IdentityProfileAllOf) HasHasTimeBasedAttr() bool {
	if o != nil && !isNil(o.HasTimeBasedAttr) {
		return true
	}

	return false
}

// SetHasTimeBasedAttr gets a reference to the given bool and assigns it to the HasTimeBasedAttr field.
func (o *IdentityProfileAllOf) SetHasTimeBasedAttr(v bool) {
	o.HasTimeBasedAttr = &v
}

func (o IdentityProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["authoritativeSource"] = o.AuthoritativeSource
	}
	if !isNil(o.IdentityRefreshRequired) {
		toSerialize["identityRefreshRequired"] = o.IdentityRefreshRequired
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.IdentityAttributeConfig) {
		toSerialize["identityAttributeConfig"] = o.IdentityAttributeConfig
	}
	if !isNil(o.IdentityExceptionReportReference) {
		toSerialize["identityExceptionReportReference"] = o.IdentityExceptionReportReference
	}
	if !isNil(o.HasTimeBasedAttr) {
		toSerialize["hasTimeBasedAttr"] = o.HasTimeBasedAttr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProfileAllOf := _IdentityProfileAllOf{}

	if err = json.Unmarshal(bytes, &varIdentityProfileAllOf); err == nil {
		*o = IdentityProfileAllOf(varIdentityProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "authoritativeSource")
		delete(additionalProperties, "identityRefreshRequired")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "identityAttributeConfig")
		delete(additionalProperties, "identityExceptionReportReference")
		delete(additionalProperties, "hasTimeBasedAttr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityProfileAllOf struct {
	value *IdentityProfileAllOf
	isSet bool
}

func (v NullableIdentityProfileAllOf) Get() *IdentityProfileAllOf {
	return v.value
}

func (v *NullableIdentityProfileAllOf) Set(val *IdentityProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProfileAllOf(val *IdentityProfileAllOf) *NullableIdentityProfileAllOf {
	return &NullableIdentityProfileAllOf{value: val, isSet: true}
}

func (v NullableIdentityProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


