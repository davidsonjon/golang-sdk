/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the CampaignFilterDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignFilterDetails{}

// CampaignFilterDetails Campaign Filter Details
type CampaignFilterDetails struct {
	// Campaign filter name.
	Name string `json:"name"`
	// Campaign filter description.
	Description *string `json:"description,omitempty"`
	// Owner of the filter. This field automatically populates at creation time with the current user.
	Owner NullableString `json:"owner"`
	// Mode/type of filter, either the INCLUSION or EXCLUSION type. The INCLUSION type includes the data in generated campaigns  as per specified in the criteria, whereas the EXCLUSION type excludes the data in generated campaigns as per specified in criteria.
	Mode map[string]interface{} `json:"mode"`
	// List of criteria.
	CriteriaList []CampaignFilterDetailsCriteriaListInner `json:"criteriaList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignFilterDetails CampaignFilterDetails

// NewCampaignFilterDetails instantiates a new CampaignFilterDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignFilterDetails(name string, owner NullableString, mode map[string]interface{}) *CampaignFilterDetails {
	this := CampaignFilterDetails{}
	this.Name = name
	this.Owner = owner
	this.Mode = mode
	return &this
}

// NewCampaignFilterDetailsWithDefaults instantiates a new CampaignFilterDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignFilterDetailsWithDefaults() *CampaignFilterDetails {
	this := CampaignFilterDetails{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignFilterDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignFilterDetails) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignFilterDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignFilterDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignFilterDetails) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CampaignFilterDetails) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}

	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignFilterDetails) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// SetOwner sets field value
func (o *CampaignFilterDetails) SetOwner(v string) {
	o.Owner.Set(&v)
}

// GetMode returns the Mode field value
func (o *CampaignFilterDetails) GetMode() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetails) GetModeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Mode, true
}

// SetMode sets field value
func (o *CampaignFilterDetails) SetMode(v map[string]interface{}) {
	o.Mode = v
}

// GetCriteriaList returns the CriteriaList field value if set, zero value otherwise.
func (o *CampaignFilterDetails) GetCriteriaList() []CampaignFilterDetailsCriteriaListInner {
	if o == nil || IsNil(o.CriteriaList) {
		var ret []CampaignFilterDetailsCriteriaListInner
		return ret
	}
	return o.CriteriaList
}

// GetCriteriaListOk returns a tuple with the CriteriaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetails) GetCriteriaListOk() ([]CampaignFilterDetailsCriteriaListInner, bool) {
	if o == nil || IsNil(o.CriteriaList) {
		return nil, false
	}
	return o.CriteriaList, true
}

// HasCriteriaList returns a boolean if a field has been set.
func (o *CampaignFilterDetails) HasCriteriaList() bool {
	if o != nil && !IsNil(o.CriteriaList) {
		return true
	}

	return false
}

// SetCriteriaList gets a reference to the given []CampaignFilterDetailsCriteriaListInner and assigns it to the CriteriaList field.
func (o *CampaignFilterDetails) SetCriteriaList(v []CampaignFilterDetailsCriteriaListInner) {
	o.CriteriaList = v
}

func (o CampaignFilterDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignFilterDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["owner"] = o.Owner.Get()
	toSerialize["mode"] = o.Mode
	if !IsNil(o.CriteriaList) {
		toSerialize["criteriaList"] = o.CriteriaList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignFilterDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner",
		"mode",
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

	varCampaignFilterDetails := _CampaignFilterDetails{}

	err = json.Unmarshal(data, &varCampaignFilterDetails)

	if err != nil {
		return err
	}

	*o = CampaignFilterDetails(varCampaignFilterDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "criteriaList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignFilterDetails struct {
	value *CampaignFilterDetails
	isSet bool
}

func (v NullableCampaignFilterDetails) Get() *CampaignFilterDetails {
	return v.value
}

func (v *NullableCampaignFilterDetails) Set(val *CampaignFilterDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignFilterDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignFilterDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignFilterDetails(val *CampaignFilterDetails) *NullableCampaignFilterDetails {
	return &NullableCampaignFilterDetails{value: val, isSet: true}
}

func (v NullableCampaignFilterDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignFilterDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


