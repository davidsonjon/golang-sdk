/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the IdentityProfile1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProfile1{}

// IdentityProfile1 struct for IdentityProfile1
type IdentityProfile1 struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// The description of the Identity Profile.
	Description NullableString `json:"description,omitempty"`
	Owner NullableIdentityProfileAllOfOwner `json:"owner,omitempty"`
	// The priority for an Identity Profile.
	Priority *int64 `json:"priority,omitempty"`
	AuthoritativeSource IdentityProfile1AllOfAuthoritativeSource `json:"authoritativeSource"`
	// True if a identity refresh is needed. Typically triggered when a change on the source has been made.
	IdentityRefreshRequired *bool `json:"identityRefreshRequired,omitempty"`
	// The number of identities that belong to the Identity Profile.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	IdentityAttributeConfig *IdentityAttributeConfig1 `json:"identityAttributeConfig,omitempty"`
	IdentityExceptionReportReference NullableIdentityExceptionReportReference1 `json:"identityExceptionReportReference,omitempty"`
	// Indicates the value of requiresPeriodicRefresh attribute for the Identity Profile.
	HasTimeBasedAttr *bool `json:"hasTimeBasedAttr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProfile1 IdentityProfile1

// NewIdentityProfile1 instantiates a new IdentityProfile1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProfile1(name string, authoritativeSource IdentityProfile1AllOfAuthoritativeSource) *IdentityProfile1 {
	this := IdentityProfile1{}
	this.Name = name
	this.AuthoritativeSource = authoritativeSource
	var identityRefreshRequired bool = false
	this.IdentityRefreshRequired = &identityRefreshRequired
	var hasTimeBasedAttr bool = false
	this.HasTimeBasedAttr = &hasTimeBasedAttr
	return &this
}

// NewIdentityProfile1WithDefaults instantiates a new IdentityProfile1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProfile1WithDefaults() *IdentityProfile1 {
	this := IdentityProfile1{}
	var identityRefreshRequired bool = false
	this.IdentityRefreshRequired = &identityRefreshRequired
	var hasTimeBasedAttr bool = false
	this.HasTimeBasedAttr = &hasTimeBasedAttr
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProfile1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProfile1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProfile1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *IdentityProfile1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityProfile1) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IdentityProfile1) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IdentityProfile1) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IdentityProfile1) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IdentityProfile1) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IdentityProfile1) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IdentityProfile1) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProfile1) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProfile1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProfile1) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IdentityProfile1) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IdentityProfile1) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IdentityProfile1) UnsetDescription() {
	o.Description.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProfile1) GetOwner() IdentityProfileAllOfOwner {
	if o == nil || isNil(o.Owner.Get()) {
		var ret IdentityProfileAllOfOwner
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProfile1) GetOwnerOk() (*IdentityProfileAllOfOwner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *IdentityProfile1) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableIdentityProfileAllOfOwner and assigns it to the Owner field.
func (o *IdentityProfile1) SetOwner(v IdentityProfileAllOfOwner) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *IdentityProfile1) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *IdentityProfile1) UnsetOwner() {
	o.Owner.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IdentityProfile1) GetPriority() int64 {
	if o == nil || isNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetPriorityOk() (*int64, bool) {
	if o == nil || isNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IdentityProfile1) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *IdentityProfile1) SetPriority(v int64) {
	o.Priority = &v
}

// GetAuthoritativeSource returns the AuthoritativeSource field value
func (o *IdentityProfile1) GetAuthoritativeSource() IdentityProfile1AllOfAuthoritativeSource {
	if o == nil {
		var ret IdentityProfile1AllOfAuthoritativeSource
		return ret
	}

	return o.AuthoritativeSource
}

// GetAuthoritativeSourceOk returns a tuple with the AuthoritativeSource field value
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetAuthoritativeSourceOk() (*IdentityProfile1AllOfAuthoritativeSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthoritativeSource, true
}

// SetAuthoritativeSource sets field value
func (o *IdentityProfile1) SetAuthoritativeSource(v IdentityProfile1AllOfAuthoritativeSource) {
	o.AuthoritativeSource = v
}

// GetIdentityRefreshRequired returns the IdentityRefreshRequired field value if set, zero value otherwise.
func (o *IdentityProfile1) GetIdentityRefreshRequired() bool {
	if o == nil || isNil(o.IdentityRefreshRequired) {
		var ret bool
		return ret
	}
	return *o.IdentityRefreshRequired
}

// GetIdentityRefreshRequiredOk returns a tuple with the IdentityRefreshRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetIdentityRefreshRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IdentityRefreshRequired) {
		return nil, false
	}
	return o.IdentityRefreshRequired, true
}

// HasIdentityRefreshRequired returns a boolean if a field has been set.
func (o *IdentityProfile1) HasIdentityRefreshRequired() bool {
	if o != nil && !isNil(o.IdentityRefreshRequired) {
		return true
	}

	return false
}

// SetIdentityRefreshRequired gets a reference to the given bool and assigns it to the IdentityRefreshRequired field.
func (o *IdentityProfile1) SetIdentityRefreshRequired(v bool) {
	o.IdentityRefreshRequired = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *IdentityProfile1) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *IdentityProfile1) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *IdentityProfile1) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetIdentityAttributeConfig returns the IdentityAttributeConfig field value if set, zero value otherwise.
func (o *IdentityProfile1) GetIdentityAttributeConfig() IdentityAttributeConfig1 {
	if o == nil || isNil(o.IdentityAttributeConfig) {
		var ret IdentityAttributeConfig1
		return ret
	}
	return *o.IdentityAttributeConfig
}

// GetIdentityAttributeConfigOk returns a tuple with the IdentityAttributeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetIdentityAttributeConfigOk() (*IdentityAttributeConfig1, bool) {
	if o == nil || isNil(o.IdentityAttributeConfig) {
		return nil, false
	}
	return o.IdentityAttributeConfig, true
}

// HasIdentityAttributeConfig returns a boolean if a field has been set.
func (o *IdentityProfile1) HasIdentityAttributeConfig() bool {
	if o != nil && !isNil(o.IdentityAttributeConfig) {
		return true
	}

	return false
}

// SetIdentityAttributeConfig gets a reference to the given IdentityAttributeConfig1 and assigns it to the IdentityAttributeConfig field.
func (o *IdentityProfile1) SetIdentityAttributeConfig(v IdentityAttributeConfig1) {
	o.IdentityAttributeConfig = &v
}

// GetIdentityExceptionReportReference returns the IdentityExceptionReportReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProfile1) GetIdentityExceptionReportReference() IdentityExceptionReportReference1 {
	if o == nil || isNil(o.IdentityExceptionReportReference.Get()) {
		var ret IdentityExceptionReportReference1
		return ret
	}
	return *o.IdentityExceptionReportReference.Get()
}

// GetIdentityExceptionReportReferenceOk returns a tuple with the IdentityExceptionReportReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProfile1) GetIdentityExceptionReportReferenceOk() (*IdentityExceptionReportReference1, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityExceptionReportReference.Get(), o.IdentityExceptionReportReference.IsSet()
}

// HasIdentityExceptionReportReference returns a boolean if a field has been set.
func (o *IdentityProfile1) HasIdentityExceptionReportReference() bool {
	if o != nil && o.IdentityExceptionReportReference.IsSet() {
		return true
	}

	return false
}

// SetIdentityExceptionReportReference gets a reference to the given NullableIdentityExceptionReportReference1 and assigns it to the IdentityExceptionReportReference field.
func (o *IdentityProfile1) SetIdentityExceptionReportReference(v IdentityExceptionReportReference1) {
	o.IdentityExceptionReportReference.Set(&v)
}
// SetIdentityExceptionReportReferenceNil sets the value for IdentityExceptionReportReference to be an explicit nil
func (o *IdentityProfile1) SetIdentityExceptionReportReferenceNil() {
	o.IdentityExceptionReportReference.Set(nil)
}

// UnsetIdentityExceptionReportReference ensures that no value is present for IdentityExceptionReportReference, not even an explicit nil
func (o *IdentityProfile1) UnsetIdentityExceptionReportReference() {
	o.IdentityExceptionReportReference.Unset()
}

// GetHasTimeBasedAttr returns the HasTimeBasedAttr field value if set, zero value otherwise.
func (o *IdentityProfile1) GetHasTimeBasedAttr() bool {
	if o == nil || isNil(o.HasTimeBasedAttr) {
		var ret bool
		return ret
	}
	return *o.HasTimeBasedAttr
}

// GetHasTimeBasedAttrOk returns a tuple with the HasTimeBasedAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfile1) GetHasTimeBasedAttrOk() (*bool, bool) {
	if o == nil || isNil(o.HasTimeBasedAttr) {
		return nil, false
	}
	return o.HasTimeBasedAttr, true
}

// HasHasTimeBasedAttr returns a boolean if a field has been set.
func (o *IdentityProfile1) HasHasTimeBasedAttr() bool {
	if o != nil && !isNil(o.HasTimeBasedAttr) {
		return true
	}

	return false
}

// SetHasTimeBasedAttr gets a reference to the given bool and assigns it to the HasTimeBasedAttr field.
func (o *IdentityProfile1) SetHasTimeBasedAttr(v bool) {
	o.HasTimeBasedAttr = &v
}

func (o IdentityProfile1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProfile1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: created is readOnly
	// skip: modified is readOnly
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["authoritativeSource"] = o.AuthoritativeSource
	if !isNil(o.IdentityRefreshRequired) {
		toSerialize["identityRefreshRequired"] = o.IdentityRefreshRequired
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.IdentityAttributeConfig) {
		toSerialize["identityAttributeConfig"] = o.IdentityAttributeConfig
	}
	if o.IdentityExceptionReportReference.IsSet() {
		toSerialize["identityExceptionReportReference"] = o.IdentityExceptionReportReference.Get()
	}
	if !isNil(o.HasTimeBasedAttr) {
		toSerialize["hasTimeBasedAttr"] = o.HasTimeBasedAttr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityProfile1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"authoritativeSource",
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

	varIdentityProfile1 := _IdentityProfile1{}

	if err = json.Unmarshal(bytes, &varIdentityProfile1); err == nil {
	*o = IdentityProfile1(varIdentityProfile1)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
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

type NullableIdentityProfile1 struct {
	value *IdentityProfile1
	isSet bool
}

func (v NullableIdentityProfile1) Get() *IdentityProfile1 {
	return v.value
}

func (v *NullableIdentityProfile1) Set(val *IdentityProfile1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProfile1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProfile1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProfile1(val *IdentityProfile1) *NullableIdentityProfile1 {
	return &NullableIdentityProfile1{value: val, isSet: true}
}

func (v NullableIdentityProfile1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProfile1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


