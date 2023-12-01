/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// checks if the FullcampaignAllOfSearchCampaignInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullcampaignAllOfSearchCampaignInfo{}

// FullcampaignAllOfSearchCampaignInfo Must be set only if the campaign type is SEARCH.
type FullcampaignAllOfSearchCampaignInfo struct {
	// The type of search campaign represented.
	Type string `json:"type"`
	// Describes this search campaign. Intended for storing the query used, and possibly the number of identities selected/available.
	Description *string `json:"description,omitempty"`
	Reviewer *FullcampaignAllOfSearchCampaignInfoReviewer `json:"reviewer,omitempty"`
	// The scope for the campaign. The campaign will cover identities returned by the query and identities that have access items returned by the query. One of `query` or `identityIds` must be set.
	Query *string `json:"query,omitempty"`
	// A direct list of identities to include in this campaign. One of `identityIds` or `query` must be set.
	IdentityIds []string `json:"identityIds,omitempty"`
	// Further reduces the scope of the campaign by excluding identities (from `query` or `identityIds`) that do not have this access.
	AccessConstraints []AccessConstraint `json:"accessConstraints,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullcampaignAllOfSearchCampaignInfo FullcampaignAllOfSearchCampaignInfo

// NewFullcampaignAllOfSearchCampaignInfo instantiates a new FullcampaignAllOfSearchCampaignInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullcampaignAllOfSearchCampaignInfo(type_ string) *FullcampaignAllOfSearchCampaignInfo {
	this := FullcampaignAllOfSearchCampaignInfo{}
	this.Type = type_
	return &this
}

// NewFullcampaignAllOfSearchCampaignInfoWithDefaults instantiates a new FullcampaignAllOfSearchCampaignInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullcampaignAllOfSearchCampaignInfoWithDefaults() *FullcampaignAllOfSearchCampaignInfo {
	this := FullcampaignAllOfSearchCampaignInfo{}
	return &this
}

// GetType returns the Type field value
func (o *FullcampaignAllOfSearchCampaignInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FullcampaignAllOfSearchCampaignInfo) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullcampaignAllOfSearchCampaignInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullcampaignAllOfSearchCampaignInfo) SetDescription(v string) {
	o.Description = &v
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *FullcampaignAllOfSearchCampaignInfo) GetReviewer() FullcampaignAllOfSearchCampaignInfoReviewer {
	if o == nil || isNil(o.Reviewer) {
		var ret FullcampaignAllOfSearchCampaignInfoReviewer
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) GetReviewerOk() (*FullcampaignAllOfSearchCampaignInfoReviewer, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given FullcampaignAllOfSearchCampaignInfoReviewer and assigns it to the Reviewer field.
func (o *FullcampaignAllOfSearchCampaignInfo) SetReviewer(v FullcampaignAllOfSearchCampaignInfoReviewer) {
	o.Reviewer = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FullcampaignAllOfSearchCampaignInfo) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FullcampaignAllOfSearchCampaignInfo) SetQuery(v string) {
	o.Query = &v
}

// GetIdentityIds returns the IdentityIds field value if set, zero value otherwise.
func (o *FullcampaignAllOfSearchCampaignInfo) GetIdentityIds() []string {
	if o == nil || isNil(o.IdentityIds) {
		var ret []string
		return ret
	}
	return o.IdentityIds
}

// GetIdentityIdsOk returns a tuple with the IdentityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) GetIdentityIdsOk() ([]string, bool) {
	if o == nil || isNil(o.IdentityIds) {
		return nil, false
	}
	return o.IdentityIds, true
}

// HasIdentityIds returns a boolean if a field has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) HasIdentityIds() bool {
	if o != nil && !isNil(o.IdentityIds) {
		return true
	}

	return false
}

// SetIdentityIds gets a reference to the given []string and assigns it to the IdentityIds field.
func (o *FullcampaignAllOfSearchCampaignInfo) SetIdentityIds(v []string) {
	o.IdentityIds = v
}

// GetAccessConstraints returns the AccessConstraints field value if set, zero value otherwise.
func (o *FullcampaignAllOfSearchCampaignInfo) GetAccessConstraints() []AccessConstraint {
	if o == nil || isNil(o.AccessConstraints) {
		var ret []AccessConstraint
		return ret
	}
	return o.AccessConstraints
}

// GetAccessConstraintsOk returns a tuple with the AccessConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) GetAccessConstraintsOk() ([]AccessConstraint, bool) {
	if o == nil || isNil(o.AccessConstraints) {
		return nil, false
	}
	return o.AccessConstraints, true
}

// HasAccessConstraints returns a boolean if a field has been set.
func (o *FullcampaignAllOfSearchCampaignInfo) HasAccessConstraints() bool {
	if o != nil && !isNil(o.AccessConstraints) {
		return true
	}

	return false
}

// SetAccessConstraints gets a reference to the given []AccessConstraint and assigns it to the AccessConstraints field.
func (o *FullcampaignAllOfSearchCampaignInfo) SetAccessConstraints(v []AccessConstraint) {
	o.AccessConstraints = v
}

func (o FullcampaignAllOfSearchCampaignInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullcampaignAllOfSearchCampaignInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.IdentityIds) {
		toSerialize["identityIds"] = o.IdentityIds
	}
	if !isNil(o.AccessConstraints) {
		toSerialize["accessConstraints"] = o.AccessConstraints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullcampaignAllOfSearchCampaignInfo) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varFullcampaignAllOfSearchCampaignInfo := _FullcampaignAllOfSearchCampaignInfo{}

	if err = json.Unmarshal(bytes, &varFullcampaignAllOfSearchCampaignInfo); err == nil {
	*o = FullcampaignAllOfSearchCampaignInfo(varFullcampaignAllOfSearchCampaignInfo)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "query")
		delete(additionalProperties, "identityIds")
		delete(additionalProperties, "accessConstraints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullcampaignAllOfSearchCampaignInfo struct {
	value *FullcampaignAllOfSearchCampaignInfo
	isSet bool
}

func (v NullableFullcampaignAllOfSearchCampaignInfo) Get() *FullcampaignAllOfSearchCampaignInfo {
	return v.value
}

func (v *NullableFullcampaignAllOfSearchCampaignInfo) Set(val *FullcampaignAllOfSearchCampaignInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFullcampaignAllOfSearchCampaignInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFullcampaignAllOfSearchCampaignInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullcampaignAllOfSearchCampaignInfo(val *FullcampaignAllOfSearchCampaignInfo) *NullableFullcampaignAllOfSearchCampaignInfo {
	return &NullableFullcampaignAllOfSearchCampaignInfo{value: val, isSet: true}
}

func (v NullableFullcampaignAllOfSearchCampaignInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullcampaignAllOfSearchCampaignInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


