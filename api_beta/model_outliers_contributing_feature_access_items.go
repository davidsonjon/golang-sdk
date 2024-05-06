/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the OutliersContributingFeatureAccessItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutliersContributingFeatureAccessItems{}

// OutliersContributingFeatureAccessItems struct for OutliersContributingFeatureAccessItems
type OutliersContributingFeatureAccessItems struct {
	// The ID of the access item
	Id *string `json:"id,omitempty"`
	// the display name of the access item
	DisplayName *string `json:"displayName,omitempty"`
	// Description of the access item.
	Description *string `json:"description,omitempty"`
	// The type of the access item.
	AccessType *string `json:"accessType,omitempty"`
	// the associated source name if it exists
	SourceName *string `json:"sourceName,omitempty"`
	// rarest access
	ExtremelyRare *bool `json:"extremelyRare,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutliersContributingFeatureAccessItems OutliersContributingFeatureAccessItems

// NewOutliersContributingFeatureAccessItems instantiates a new OutliersContributingFeatureAccessItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutliersContributingFeatureAccessItems() *OutliersContributingFeatureAccessItems {
	this := OutliersContributingFeatureAccessItems{}
	var extremelyRare bool = false
	this.ExtremelyRare = &extremelyRare
	return &this
}

// NewOutliersContributingFeatureAccessItemsWithDefaults instantiates a new OutliersContributingFeatureAccessItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutliersContributingFeatureAccessItemsWithDefaults() *OutliersContributingFeatureAccessItems {
	this := OutliersContributingFeatureAccessItems{}
	var extremelyRare bool = false
	this.ExtremelyRare = &extremelyRare
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OutliersContributingFeatureAccessItems) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutliersContributingFeatureAccessItems) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OutliersContributingFeatureAccessItems) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OutliersContributingFeatureAccessItems) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OutliersContributingFeatureAccessItems) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutliersContributingFeatureAccessItems) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OutliersContributingFeatureAccessItems) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OutliersContributingFeatureAccessItems) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutliersContributingFeatureAccessItems) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutliersContributingFeatureAccessItems) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutliersContributingFeatureAccessItems) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutliersContributingFeatureAccessItems) SetDescription(v string) {
	o.Description = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *OutliersContributingFeatureAccessItems) GetAccessType() string {
	if o == nil || isNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutliersContributingFeatureAccessItems) GetAccessTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *OutliersContributingFeatureAccessItems) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *OutliersContributingFeatureAccessItems) SetAccessType(v string) {
	o.AccessType = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *OutliersContributingFeatureAccessItems) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutliersContributingFeatureAccessItems) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *OutliersContributingFeatureAccessItems) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *OutliersContributingFeatureAccessItems) SetSourceName(v string) {
	o.SourceName = &v
}

// GetExtremelyRare returns the ExtremelyRare field value if set, zero value otherwise.
func (o *OutliersContributingFeatureAccessItems) GetExtremelyRare() bool {
	if o == nil || isNil(o.ExtremelyRare) {
		var ret bool
		return ret
	}
	return *o.ExtremelyRare
}

// GetExtremelyRareOk returns a tuple with the ExtremelyRare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutliersContributingFeatureAccessItems) GetExtremelyRareOk() (*bool, bool) {
	if o == nil || isNil(o.ExtremelyRare) {
		return nil, false
	}
	return o.ExtremelyRare, true
}

// HasExtremelyRare returns a boolean if a field has been set.
func (o *OutliersContributingFeatureAccessItems) HasExtremelyRare() bool {
	if o != nil && !isNil(o.ExtremelyRare) {
		return true
	}

	return false
}

// SetExtremelyRare gets a reference to the given bool and assigns it to the ExtremelyRare field.
func (o *OutliersContributingFeatureAccessItems) SetExtremelyRare(v bool) {
	o.ExtremelyRare = &v
}

func (o OutliersContributingFeatureAccessItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutliersContributingFeatureAccessItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.ExtremelyRare) {
		toSerialize["extremelyRare"] = o.ExtremelyRare
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutliersContributingFeatureAccessItems) UnmarshalJSON(bytes []byte) (err error) {
	varOutliersContributingFeatureAccessItems := _OutliersContributingFeatureAccessItems{}

	if err = json.Unmarshal(bytes, &varOutliersContributingFeatureAccessItems); err == nil {
	*o = OutliersContributingFeatureAccessItems(varOutliersContributingFeatureAccessItems)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "extremelyRare")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutliersContributingFeatureAccessItems struct {
	value *OutliersContributingFeatureAccessItems
	isSet bool
}

func (v NullableOutliersContributingFeatureAccessItems) Get() *OutliersContributingFeatureAccessItems {
	return v.value
}

func (v *NullableOutliersContributingFeatureAccessItems) Set(val *OutliersContributingFeatureAccessItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOutliersContributingFeatureAccessItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOutliersContributingFeatureAccessItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutliersContributingFeatureAccessItems(val *OutliersContributingFeatureAccessItems) *NullableOutliersContributingFeatureAccessItems {
	return &NullableOutliersContributingFeatureAccessItems{value: val, isSet: true}
}

func (v NullableOutliersContributingFeatureAccessItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutliersContributingFeatureAccessItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


